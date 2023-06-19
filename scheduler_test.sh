#!/bin/bash

base64 -d <<<"ICAgX19fX18gICAgICBfICAgICAgICAgICAgICBfICAgICAgIF8gICAgICAgICAgIAogIC8gX19f
X3wgICAgfCB8ICAgICAgICAgICAgfCB8ICAgICB8IHwgICAgICAgICAgCiB8IChfX18gICBfX198
IHxfXyAgIF9fXyAgX198IHxfICAgX3wgfCBfX18gXyBfXyAKICBcX19fIFwgLyBfX3wgJ18gXCAv
IF8gXC8gX2AgfCB8IHwgfCB8LyBfIFwgJ19ffAogIF9fX18pIHwgKF9ffCB8IHwgfCAgX18vIChf
fCB8IHxffCB8IHwgIF9fLyB8ICAgCiB8X19fX18vIFxfX198X3wgfF98XF9fX3xcX18sX3xcX18s
X3xffFxfX198X3wgICAK"
echo

PS3="Select operation: "
url="http://localhost:3000"
user=""
curl="curl -k"

select opt in new_record new_schedule switch get_current_schedule quit; do
	clear
	case $opt in
		new_record)
			$curl $url/schedule

			read -p "enter schedule id: " scheduleId
			read -p "what you currently doing?: " nowDoing
			
			$curl -H 'content-type: application/json' \
				-d "{ \"schedule_id\": \"$scheduleId\", \"now_doing\": \"$nowDoing\" }" \
				-X POST \
				--fail $url/record/
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
				-d "{ \"switch_to\": \"$nowDoing\" }" \
				-X POST \
				--fail $url/record/switch
			;;
		get_current_schedule)
			$curl --fail $url/record/remain
			;;
		quit)
			break
			;;
		*)
			echo "Invalid operation: $opt"
			;;
	esac
done
