package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 2147483647
    v2 := []int{497}
    v3 := []uint32{(v1 & v1), v1, (v1 % ((65535) | 1))}
    v4 := []uint32{}
    v5 := []string{"\n\t", (("hello" + (("a" + "\n\t") + "")) + "a"), "a", ""}
    var v6 int16 = 797
    var v7 int16 = (v6 * (v6 - (v6 << ((18446744073709551615) & 63))))
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint64 = 42
        switch fi8 {
        case 4294967295:
            fmt.Println(fi8)
        case 209:
            v7 = 1000
        case 788:
            var v10 uint64 = v9
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    var v11 bool = (1 != 42)
    switch 32767 {
    case -75:
        for fi12 := 0; fi12 < 9; fi12++ {
            switch "" {
            case "":
                var v13 int16 = (v7 * v6)
                _ = v13
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "hello":
                fmt.Println(v11, v7)
            default:
                _ = len(v5)
            }
            _ = fi12
        }
    case 42:
        v7 = (v7 >> ((0) & 63))
    case 0:
        v6 = (v7 % (((v6 * (^(v7)))) | 1))
    default:
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    var v14 uint32 = (v1 * (v1 ^ (v1 + (v1 << ((0) & 63)))))
    var v15 string = ("foo bar" + "a")
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(fi16, v1, v15)
        _ = fi16
    }
    if true {
        fmt.Println(v15)
        fmt.Println(v15, v7)
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v6, v1)
            _ = fi17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
