package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 float32 = 1e10
    var v2 float64 = 1e10
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    var v6 int32 = 525
    v7 := []int{int(v6), 116, 128, 1}
    _ = len(v7)
    v8 := GenS1{F0: "a"}
    var v9 uint16 = 42
    if false {
        if (false && (1 <= -32768)) {
            v10 := GenS1{F0: "foo bar"}
            _ = v10
        } else {
            switch "hello" {
            case "\n\t":
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(v1, v1)
                    _ = fi11
                }
            case "":
                switch "" {
                case "":
                    v1 -= v1
                case "hello":
                    if len(v7) > 0 {
                        v7[0] = 128
                    }
                default:
                    v6 -= v6
                }
            case "hello":
                fmt.Println(v2, v2)
            }
        }
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    var v12 int = 32768
    fmt.Println(v6, v9)
    if ((("a" + ((("" + "") + ("a" + ("hello" + ((("foo bar" + "foo bar") + "foo bar") + "\n\t")))) + "")) <= "hello") || ((v9 != 108) && (0 >= 0))) {
        v12 = (v12 | v12)
    }
    var v13 uint = 65535
    var v14 uint64 = 459
    v1 += ((0.0 + v1) - ((float32((float64(1e10) / (v2 * (1e10 / v2)))) / (v1 + v1)) * v1))
    v15 := GenS2{F0: v14, F1: (^(v6))}
    _ = len(v7)
    v9 = (v9 - (v9 % (((^(v9))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
