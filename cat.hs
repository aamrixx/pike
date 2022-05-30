-- Functional Coreutils in Haskell
-- cat.hs Print out file contents

import System.Environment

main = do
    args <- getArgs
    let file = args !! 0
    contents <- readFile file
    putStrLn contents
