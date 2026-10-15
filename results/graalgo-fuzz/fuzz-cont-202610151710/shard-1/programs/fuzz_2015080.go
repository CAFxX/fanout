package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() uint32 {
    var v5 int = 181
    v6 := []int{32768, v5, (v5 - -9223372036854775808), (v5 | (^(v5)))}
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    var v3 int64 = 42
    v4 := []string{("hello" + (("" + "foo bar") + (("a" + "hello") + ("hello" + ((("hello" + "") + ("\n\t" + "\n\t")) + ("a" + "hello")))))), ("a" + "hello")}
    var v7 uint16 = 2
    for fi8 := 0; fi8 < 20; fi8++ {
        if v1 {
            fmt.Println(v7, v3, v1)
        }
        var v9 int8 = -78
        _ = fi8
        _ = v9
    }
    if (2 <= 0) {
        switch -9223372036854775808 {
        case 56:
            var v10 uint16 = (v7 / ((v7) | 1))
            _ = v10
        }
        if (255 == 42) {
            v3 -= (v3 + v3)
            v2[433] = true
        }
        if true {
            switch ("a" + "foo bar") {
            case "":
                v11 := GenS1{F0: 65535}
                _ = v11
            default:
                var v12 uint16 = (v7 & 1)
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int64 = (v3 >> ((4294967295) & 63))
                _ = fi13
                _ = v14
            }
        }
    } else {
        fmt.Println(v7, v3)
    }
    v1 = v1
    fmt.Println(v3, v1, v1)
    v1 = false
    v3 = (v3 >> ((42) & 63))
    var v15 uint = 269
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
