package manual

func Banner() string {
	return `
	$$$$$$$$\ $$\           $$\        $$\           $$$$$$$$\ $$\ $$\                  $$$$$$\                                    
	$$  _____|\__|          $$ |       $$ |          \__$$  __|\__|$$ |                $$  __$$\                                   
	$$ |      $$\  $$$$$$\  $$$$$$$\ $$$$$$\            $$ |   $$\ $$ | $$$$$$\        $$ /  \__| $$$$$$\  $$$$$$\$$$$\   $$$$$$\  
	$$$$$\    $$ |$$  __$$\ $$  __$$\\_$$  _|           $$ |   $$ |$$ |$$  __$$\       $$ |$$$$\  \____$$\ $$  _$$  _$$\ $$  __$$\ 
	$$  __|   $$ |$$ /  $$ |$$ |  $$ | $$ |             $$ |   $$ |$$ |$$$$$$$$ |      $$ |\_$$ | $$$$$$$ |$$ / $$ / $$ |$$$$$$$$ |
	$$ |      $$ |$$ |  $$ |$$ |  $$ | $$ |$$\          $$ |   $$ |$$ |$$   ____|      $$ |  $$ |$$  __$$ |$$ | $$ | $$ |$$   ____|
	$$$$$$$$\ $$ |\$$$$$$$ |$$ |  $$ | \$$$$  |         $$ |   $$ |$$ |\$$$$$$$\       \$$$$$$  |\$$$$$$$ |$$ | $$ | $$ |\$$$$$$$\ 
	\________|\__| \____$$ |\__|  \__|  \____/          \__|   \__|\__| \_______|       \______/  \_______|\__| \__| \__| \_______|
                      $$\   $$ |                                                                                                       
                      \$$$$$$  |                                                                                                       
                       \______/	                                             
	`
}

func Instructions() string {
	return `
	w - Up
	s - Down
	a - Left
	d - Right
	i - Print instructions
	p - Print current board
	b - Print current branch
	t - Print the entire tree
	x - Exit game
	`
}
