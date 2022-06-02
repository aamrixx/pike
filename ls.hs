import System.Directory
import System.Environment

def :: String
def = "."

main = do
    args <- getArgs
    
    if length args == 1 then do
        let path = args !! 0
        dirContents <- listDirectory path
        let c = dirContents
        print c
    else do
        dirContents <- listDirectory def
        mapM_ putStrLn $ dirContents
