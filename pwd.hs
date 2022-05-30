import System.Directory

main = do
    currentDir <- getCurrentDirectory
    putStrLn currentDir
    
