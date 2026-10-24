package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 32767
    var v3 int = 2147483647
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 17; fi5++ {
        switch 32768 {
        case 4294967295:
            if true {
                fmt.Println(v4, fi5, v3)
            } else {
                var v6 float32 = 1.5
                fmt.Println(v3, v2, fi5)
                _ = v6
            }
        default:
            var v7 int16 = (v2 | v2)
            _ = v7
        }
        v2 += v2
        _ = fi5
    }
    v3 = 0
    v3 += v3
    if (1.5 < -2.25) {
        var v8 int8 = -44
        var v9 uint32 = 2
        var v10 uint64 = 4294967295
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 18446744073709551615, F1: false}
    if false {
        v12 := GenS1{F0: 18446744073709551615, F1: (true || (v2 != v2))}
        if (false && false) {
            if true {
                switch ("foo bar" + "hello") {
                case "foo bar":
                    if (0.0 != 0.0) {
                        v13 := []int{-1}
                        _ = v13
                    }
                case "a":
                    v14 := GenS1{F0: 0, F1: true}
                    _ = v14
                }
            }
            if false {
                v1[(v3 + -1)] = ("hello" + ("\n\t" + "a"))
            }
        } else {
            var v15 int16 = (v2 >> ((255) & 63))
            v3 = 128
            v16 := []int{}
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: 18446744073709551615, F1: (1 >= 0)}
        _ = v12
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(v4, fi18)
        v19 := []int{(v3 - (fi18 - v3)), (-(fi18)), (v3 + fi18), fi18}
        _ = fi18
        _ = v19
    }
    v3 -= v3
    v20 := GenS1{F0: 1, F1: (18446744073709551615 >= 1)}
    v21 := v1[v3]
    v22 := []int{-2147483648}
    var v23 uint8 = v4
    v2 += -689
    if (false || false) {
        v2 += (v2 >> ((0) & 63))
        v23 -= 0
        _ = len(v22)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
