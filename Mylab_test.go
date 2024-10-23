package main

import "testing"

func TestMyLab(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{{
		name:  "",
		input: "",
		want:  ``,
	}, {
		name:  "\\n",
		input: "\\n",
		want:  "\n",
	}, {
		name:  "Hello\\n",
		input: "Hello\\n",
		want: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
`,
	}, {
		name:  "hello",
		input: "hello",
		want: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
	}, {
		name:  "HeLlO",
		input: "HeLlO",
		want: ` _    _          _        _    ____   
| |  | |        | |      | |  / __ \  
| |__| |   ___  | |      | | | |  | | 
|  __  |  / _ \ | |      | | | |  | | 
| |  | | |  __/ | |____  | | | |__| | 
|_|  |_|  \___| |______| |_|  \____/  
                                      
                                      
`,
	}, {
		name:  "Hello There",
		input: "Hello There",
		want: ` _    _          _   _                 _______   _                           
| |  | |        | | | |               |__   __| | |                          
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| 
                                                                             
                                                                             
`,
	}, {
		name:  "1Hello 2There",
		input: "1Hello 2There",
		want: `     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`,
	}, {
		name:  "{Hello There}",
		input: "{Hello There}",
		want: `   __  _    _          _   _                 _______   _                           __    
  / / | |  | |        | | | |               |__   __| | |                          \ \   
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ 
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  
  \_\                                                                              /_/   
                                                                                         
`,
	}, {
		name:  "Hello\nThere",
		input: "Hello\\nThere",
		want: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
	}, {
		name:  "Hello\n\nThere",
		input: "Hello\\n\\nThere",
		want: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
	}}

	// Iterate through each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(*testing.T) {

			// Run the function with the test input
			result := MyLab(tc.input)

			// Compare the result with the expected output
			if result != tc.want {
				t.Errorf("MyLab(%s): \n%sexpected\n%s", tc.input, result, tc.want)
			}
		})
	}
}
