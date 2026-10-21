package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []string{("hello" + "hello"), "hello", "\n\t"}
    var v2 int32 = 939
    var v3 uint16 = 0
    v2 = 2147483647
    var v4 bool = true
    for fi5 := 0; fi5 < 4; fi5++ {
        v2 += (v2 << ((950) & 63))
        fmt.Println(fi5, v2)
        v1 = append(v1, ((("foo bar" + ((("" + "a") + "") + "hello")) + ("foo bar" + "")) + "hello"))
        _ = fi5
    }
    var v6 int64 = -1
    if true {
        v3 = (^(v3))
        for fi7 := 0; fi7 < 6; fi7++ {
            v6 -= int64(v2)
            _ = fi7
        }
    }
    v3 -= (v3 << ((385) & 63))
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    var v8 string = "foo bar"
    var v9 uint16 = v3
    v10 := GenS1{F0: v6, F1: 1}
    if ((3.14159 < 3.14159) || (!(1000 <= 776))) {
        if v4 {
            v3 += v9
        } else {
            v9 = (v9 | (v9 ^ v3))
        }
    }
    switch -530 {
    case -755:
        var v11 int = 32767
        _ = v11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
