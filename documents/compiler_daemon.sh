while true; do
  change=$(inotifywait -e close_write,moved_to,create .)
  change=${change#./ * }
  if [[ "$change" = *".go" ]]
  then
  	echo "######## Killing previous proccess ###########"
	killall raspGo
	echo "######## Starting Compiling ###########"
	echo $'\e]9;Start Compiling\007' && go build && (./raspGo &) && echo $'\e]9;Server Started\007'
	echo "######## Finished Compiling ###########"
   fi
done
