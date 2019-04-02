# Array of commands to run in different tabs
commands=(
    #'cd $GOPATH/src/github.com/ofonimefrancis/problemsapp/web && PORT=3500 yarn dev'
    'cd $GOPATH/src/github.com/ofonimefrancis/problemsapp && go run main.go'
)

# Build final command with all the tabs to launch
set finalCommand=""
for (( i = 0; i < ${#commands[@]}; i++ )); do
    export finalCommand+="--tab -e 'bash -c \"${commands[$i]}\"' "
done

# Run the final command
eval "gnome-terminal "$finalCommand