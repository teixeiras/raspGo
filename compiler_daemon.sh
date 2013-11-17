while true; do
  change=$(inotifywait -e close_write,moved_to,create .)
  change=${change#./ * }
  if [[ "$change" = *".go" ]]
  then
  	echo "######## Killing previous proccess ###########"
	killall raspGo
	echo "######## Starting Compiling ###########"
	go build && ./raspGo &
	echo "######## Finished Compiling ###########"
   fi
done
