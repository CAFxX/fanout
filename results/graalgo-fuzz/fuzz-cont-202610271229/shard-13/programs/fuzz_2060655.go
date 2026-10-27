package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() int {
    v5 := []int{-1, 2147483647}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 uint8 = 2
    var v7 string = ((("foo bar" + "\n\t") + "\n\t") + "a")
    if false {
        var v8 uint16 = 162
        var v9 uint64 = uint64(v8)
        if ((^(v8)) <= (^(v8))) {
            var v10 uint = 65535
            _ = v10
        }
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return -3
}

func main() {
    var v1 int64 = 914
    v2 := []string{}
    var v3 int8 = int8(v1)
    var v4 bool = true
    var v11 uint64 = 109
    v1 = 42
    var v12 uint32 = 42
    var v13 uint64 = (v11 | v11)
    if true {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v11 = (v11 ^ 18446744073709551615)
        v11 = 53
    }
    if true {
        v11 = v13
        v2 = append(v2, ("hello" + (("" + "\n\t") + "")))
    }
    fmt.Println(v1, v3, v13)
    var v14 int16 = 394
    var v15 int = 32768
    v16 := GenS1{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
