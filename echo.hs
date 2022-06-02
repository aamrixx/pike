import System.Environment

main = do
    args <- getArgs
    
    if length args == 0
        then putStrLn ""
    else do
        let text = args !! 0
        putStrLn $ text
