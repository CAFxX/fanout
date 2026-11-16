package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    v1 := []string{("" + "foo bar"), ("hello" + (("hello" + "") + "hello")), "hello", "a"}
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483647, 492}
    var v4 uint8 = 112
    _ = len(v1)
    var v5 float64 = 3.14159
    fmt.Println(v5, v4)
    v6 := GenS1{F0: v4, F1: 328}
    fmt.Println(v4, v5)
    var v7 uint8 = (v4 ^ (^(v4)))
    v8 := GenS1{F0: v7, F1: 65535}
    v9 := GenS2{F0: (("" + "foo bar") + "a")}
    fmt.Println(v7, v4)
    switch 127 {
    case 128:
        var v10 uint = 18446744073709551615
        _ = v10
    case 366:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 string = ("\n\t" + "foo bar")
            fmt.Println(v7, v12)
            _ = fi11
            _ = v12
        }
    case 32768:
        switch "foo bar" {
        case "foo bar":
            v7 = (v7 | (v7 - v4))
        case "a":
            if (v5 != float64(1e10)) {
                var v13 uint16 = 270
                v5 += v5
                _ = v13
            }
        case "":
            if true {
                v7 += (v4 % (((^(v4))) | 1))
            }
        default:
            if true {
                v14 := GenS1{F0: (v7 / ((v7) | 1)), F1: 255}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v15 uint = 402
                _ = v14
                _ = v15
            }
        }
    default:
        var v16 int32 = -1
        _ = v16
    }
    var v17 float64 = (((v5 / ((v5 - (v5 / (v5 / 1.5))) / (v5 / 0.1))) + 3.14159) - v5)
    var v18 uint16 = 666
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
