#!/bin/bash

PS3="Select operation: "
current_record=""

base64 -d <<<"ICAgX19fX18gICAgICBfICAgICAgICAgICAgICBfICAgICAgIF8gICAgICAgICAgIAogIC8gX19f
X3wgICAgfCB8ICAgICAgICAgICAgfCB8ICAgICB8IHwgICAgICAgICAgCiB8IChfX18gICBfX198
IHxfXyAgIF9fXyAgX198IHxfICAgX3wgfCBfX18gXyBfXyAKICBcX19fIFwgLyBfX3wgJ18gXCAv
IF8gXC8gX2AgfCB8IHwgfCB8LyBfIFwgJ19ffAogIF9fX18pIHwgKF9ffCB8IHwgfCAgX18vIChf
fCB8IHxffCB8IHwgIF9fLyB8ICAgCiB8X19fX18vIFxfX198X3wgfF98XF9fX3xcX18sX3xcX18s
X3xffFxfX198X3wgICAK"
echo

select opt in init_record load_record new_record switch get_current_schedule quit; do
	clear
	base64 -d <<<"ICAgX19fX18gICAgICBfICAgICAgICAgICAgICBfICAgICAgIF8gICAgICAgICAgIAogIC8gX19f
X3wgICAgfCB8ICAgICAgICAgICAgfCB8ICAgICB8IHwgICAgICAgICAgCiB8IChfX18gICBfX198
IHxfXyAgIF9fXyAgX198IHxfICAgX3wgfCBfX18gXyBfXyAKICBcX19fIFwgLyBfX3wgJ18gXCAv
IF8gXC8gX2AgfCB8IHwgfCB8LyBfIFwgJ19ffAogIF9fX18pIHwgKF9ffCB8IHwgfCAgX18vIChf
fCB8IHxffCB8IHwgIF9fLyB8ICAgCiB8X19fX18vIFxfX198X3wgfF98XF9fX3xcX18sX3xcX18s
X3xffFxfX198X3wgICAK"
	echo
	cat<<EOF
1) init_record           3) new_record            5) get_current_schedule
2) load_record           4) switch                6) quit
EOF
	case $opt in
		init_record)
			curl http://localhost:3000/schedule

			read -p "enter schedule id: " scheduleId
			read -p "what you currently doing?: " nowDoing
			
			recordId=`curl -H 'content-type: application/json' \
				-d "{ \"schedule_id\": \"$scheduleId\", \"now_doing\": \"$nowDoing\" }" \
				-X POST \
				--fail http://localhost:3000/record/init-record`

			if [ $? -eq 0 ]; then
				current_record=$recordId
			fi
			;;
		load_record)
			read -p "Enter record id: " recordId
			curl --fail http://localhost:3000/record/remain/$recordId
			if [ $? -eq 0 ]; then
				current_record=$recordId
			fi
			;;
		new_record)
			if [ -z $current_record ]; then
				echo "You need to initize record first [1](or load the record [2])"
				continue
			fi

			curl http://localhost:3000/schedule

			read -p "enter schedule id: " scheduleId
			read -p "what you currently doing?: " nowDoing
			
			recordId=`curl -H 'content-type: application/json' \
				-d "{ \"schedule_id\": \"$scheduleId\", \"now_doing\": \"$nowDoing\" }" \
				-X POST \
				--fail http://localhost:3000/record/$current_record`

			if [ $? -eq 0 ]; then
				current_record=$recordId
			fi
			;;
		switch)
			if [ -z $current_record ]; then
				echo "You need to initize record first [1](or load the record [2])"
				continue
			fi

			read -p "what you currently doing?: " nowDoing

			curl -H 'content-type: application/json' \
				-d "{ \"record_id\": \"$current_record\", \"switch_to\": \"$nowDoing\" }" \
				-X POST \
				--fail http://localhost:3000/record/switch
			;;
		get_current_schedule)
			if [ -z $current_record ]; then
				echo "You need to initize record first [1](or load the record [2])"
				continue
			fi
			
			curl --fail http://localhost:3000/record/remain/$current_record
			;;
		quit)
			break
			;;
		*)
			echo "Invalid operation: $opt"
			;;
	esac
done
