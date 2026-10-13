package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]int)
    var v3 uint16 = 32767
    v4 := GenS1{F0: 718, F1: 32767, F2: uint32(v3)}
    var v5 uint8 = (v1 / ((144) | 1))
    delete(v2, 1)
    var v6 uint = 4294967295
    v6 = v6
    v7 := []int{612, 542, 9223372036854775807}
    if true {
        v7 = append(v7, -9223372036854775808)
        switch ((("a" + "hello") + "foo bar") + ("\n\t" + ("" + "hello"))) {
        case "\n\t":
            var v8 string = (("\n\t" + "hello") + "a")
            _ = v8
        case "":
            if true {
                v5 -= 2
            } else {
                v3 -= (^(v3))
            }
        default:
            v3 = v3
        }
    } else {
        var v9 float64 = (((((float64(0.0) / (float64((1e10 + 0.0)) / 1.5)) + (float64(-2.25) + 1.5)) / (-2.25 + -2.25)) - (1e10 - -2.25)) / float64(3.14159))
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v9
    }
    v10 := GenS1{F0: 114, F1: 720, F2: 4294967295}
    v6 = v6
    fmt.Println(v3, v5, v3)
    for fi11 := 0; fi11 < 3; fi11++ {
        v6 = v6
        if false {
            var v12 int = -196
            for fi13 := 0; fi13 < 2; fi13++ {
                v3 = (v3 % (((v3 & v3)) | 1))
                _ = fi13
            }
            if (!(4294967295 >= 545)) {
                var v14 uint = (^(v6))
                v5 += (v5 << ((4294967295) & 63))
                _ = v14
            }
            _ = v12
        } else {
            v6 += (^(v6))
        }
        _ = fi11
    }
    _ = len(v7)
    var v15 uint64 = 743
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
