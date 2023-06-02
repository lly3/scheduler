#!/bin/bash

base64 -d <<<"ICAgX19fX18gICAgICBfICAgICAgICAgICAgICBfICAgICAgIF8gICAgICAgICAgIAogIC8gX19f
X3wgICAgfCB8ICAgICAgICAgICAgfCB8ICAgICB8IHwgICAgICAgICAgCiB8IChfX18gICBfX198
IHxfXyAgIF9fXyAgX198IHxfICAgX3wgfCBfX18gXyBfXyAKICBcX19fIFwgLyBfX3wgJ18gXCAv
IF8gXC8gX2AgfCB8IHwgfCB8LyBfIFwgJ19ffAogIF9fX18pIHwgKF9ffCB8IHwgfCAgX18vIChf
fCB8IHxffCB8IHwgIF9fLyB8ICAgCiB8X19fX18vIFxfX198X3wgfF98XF9fX3xcX18sX3xcX18s
X3xffFxfX198X3wgICAK"
echo

PS3="Select operation: "
current_record=""
url="http://localhost:3000"
user=""
curl="curl -k"

recordId=`$curl --fail $url/record/latest`
if [ $? -eq 0 ]
then
	current_record=$recordId
else
	echo -e "Create record [1] to start using the application"
fi

select opt in new_record new_schedule switch get_current_schedule quit; do
	clear
	case $opt in
		new_record)
			$curl $url/schedule

			read -p "enter schedule id: " scheduleId
			read -p "what you currently doing?: " nowDoing
			
			recordId=`$curl -H 'content-type: application/json' \
				-d "{ \"prev_record_id\": \"$current_record\", \"schedule_id\": \"$scheduleId\", \"now_doing\": \"$nowDoing\" }" \
				-X POST \
				--fail $url/record/`

			if [ $? -eq 0 ]; then
				current_record=$recordId
			fi
			;;
		new_schedule)
			todos=""

			while true
			do
				read -p "Enter your todo: " title
				read -p "Enter your duration(10ms,1m,1h): " duration

				todos+="{"
				todos+="\"title\":\"$title\""
				todos+=","
				todos+="\"duration\":\"$duration\""
				todos+="}"

				read -p "Are you done?(y/n): " ex
				if [ $ex == "y" ]; then
					break
				fi

				todos+=","
			done

			echo $todos

			$curl -H 'content-type: application/json' \
				-d "{ \"todos\": [$todos] }" \
				-X POST \
				$url/schedule
			;;
		switch)
			read -p "what you currently doing?: " nowDoing

			$curl -H 'content-type: application/json' \
				-d "{ \"record_id\": \"$current_record\", \"switch_to\": \"$nowDoing\" }" \
				-X POST \
				--fail $url/record/switch
			;;
		get_current_schedule)
			$curl --fail $url/record/remain/$current_record
			;;
		quit)
			break
			;;
		*)
			echo "Invalid operation: $opt"
			;;
	esac
done
