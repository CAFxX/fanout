package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32; F1 int16; F2 bool }
func genFn0(p0 uint32) int {
    var v5 uint8 = 2
    var v6 float64 = 1.5
    switch "" {
    case "foo bar":
        v7 := GenS1{F0: 9223372036854775807, F1: 0, F2: (v5 / (((v5 % ((127) | 1))) | 1))}
        _ = v7
    default:
        switch 9223372036854775807 {
        case -1:
            v5 -= v5
        default:
            var v8 float64 = v6
            _ = v8
        }
    }
    switch 683 {
    case -2147483648:
        switch 42 {
        case 128:
            if false {
                var v9 int32 = 1
                if false {
                    var v10 uint8 = (v5 >> ((2) & 63))
                    v11 := GenS2{F0: false, F1: 2147483647, F2: uint16(v10)}
                    v12 := []int{2147483647, 42, -9223372036854775808}
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v13 := GenS3{F0: 127}
                _ = v13
                _ = v9
            } else {
                v14 := GenS4{F0: 2, F1: -128}
                v6 = (-2.25 / (((v6 / v6) / v6) + ((0.0 - float64((-2.25 * -2.25))) - v6)))
                var v15 float64 = v6
                _ = v14
                _ = v15
            }
        }
    default:
        fmt.Println(v6)
    }
    _ = p0
    _ = v5
    _ = v6
    return -971
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", ("a" + "foo bar"), ""}
    var v3 uint = 2
    var v4 uint8 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int = genFn0(2147483647)
        _ = fi16
        _ = v17
    }
    _ = len(v1)
    v18 := GenS5{F0: uint32(v4), F1: -32768, F2: ((32768 != -209) && (!(int32(v4) != -2147483648)))}
    if (!(true && (true || (511 < 431)))) {
        var v19 string = ("hello" + "\n\t")
        fmt.Println(v3)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v19
    }
    fmt.Println(v4, v3)
    if (!true) {
        v2 = append(v2, "hello")
        var v20 uint16 = 2
        fmt.Println(v4)
        _ = v20
    }
    fmt.Println(v4)
    v4 -= 22
    fmt.Println(v3, v3)
    v21 := GenS1{F0: 9223372036854775807, F1: 639, F2: (^(v4))}
    v3 = 1
    for fi22 := 0; fi22 < 20; fi22++ {
        v3 += 2
        _ = fi22
    }
    var v23 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
