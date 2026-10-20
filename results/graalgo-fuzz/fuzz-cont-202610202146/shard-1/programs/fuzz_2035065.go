package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0() int {
    var v7 int16 = 32767
    if false {
        v8 := GenS1{F0: 255}
        var v9 uint32 = 42
        v7 += (-(v7))
        _ = v8
        _ = v9
    } else {
        v7 = v7
        v10 := GenS2{F0: false}
        v11 := GenS1{F0: 127}
        _ = v10
        _ = v11
    }
    v12 := []int{}
    v7 -= (v7 ^ (v7 << ((65535) & 63)))
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int = fi13
        for fi15 := 0; fi15 < 1; fi15++ {
            v14 += (fi13 * fi15)
            _ = fi15
        }
        fmt.Println(v7)
        _ = fi13
        _ = v14
    }
    _ = v12
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 162
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 uint8 = 255
    v5 := make(map[string]int)
    v6 := make(map[string]bool)
    switch "foo bar" {
    case "":
        v4 = v4
    case "foo bar":
        for fi16 := 0; fi16 < 18; fi16++ {
            if (127 != fi16) {
                switch ("hello" + "\n\t") {
                case "hello":
                    var v17 float64 = ((((1e10 + 0.1) * float64(0.0)) - -2.25) + 3.14159)
                    _ = v17
                case "foo bar":
                    v4 += v4
                default:
                    var v18 bool = (127 >= 9)
                    _ = v18
                }
            }
            v19 := []int{-421, (fi16 * (fi16 << ((671) & 63))), (-(fi16))}
            _ = fi16
            _ = v19
        }
    case "a":
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := []int{128, genFn0(), fi20, 42}
            v22 := GenS1{F0: 2}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v4 += 2
    delete(v6, ("hello" + "a"))
    var v23 bool = (!(105 < 60))
    v5["hello"] = 1
    v24 := v2[4294967295]
    _ = len(v3)
    v25 := GenS2{F0: v23}
    var v26 uint16 = 2
    v4 = v1
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 uint32 = 971
        var v29 uint16 = (v26 & (v26 / (((v26 ^ v26)) | 1)))
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
