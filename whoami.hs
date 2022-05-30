import System.Posix.User
import System.Environment

main = do
    args <- getArgs
    
    if length args == 0 then do
        user <- getEffectiveUserName
        putStrLn user
    else do
        if args !! 0 == "--version"
            then putStrLn "whoami 0.0.1 by xhci1"
        else putStrLn "usage: whoami [--version]"
        
