import System.Environment

main = do
    args <- getArgs
    let text = args !! 0
    putStrLn $ text
