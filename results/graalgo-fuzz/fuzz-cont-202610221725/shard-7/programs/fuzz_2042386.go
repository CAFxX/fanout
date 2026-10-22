package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8, p1 int8, p2 uint8) uint8 {
    v7 := GenS1{F0: (true || ((-399 == -9223372036854775808) && (-2.25 > (-2.25 + 1.5))))}
    for fi8 := 0; fi8 < 7; fi8++ {
        switch ("a" + ("\n\t" + "")) {
        case "":
            p2 += p2
        case "hello":
            if (!false) {
                v9 := GenS1{F0: false}
                _ = v9
            }
        default:
            p2 -= (^(p2))
        }
        p2 = p2
        p1 = (p0 ^ p0)
        _ = fi8
    }
    var v10 float64 = 1e10
    v11 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 1
    v2 := []int{32768, 127}
    var v3 string = ("a" + "foo bar")
    v4 := make(map[int]int)
    var v5 uint8 = 128
    var v6 uint64 = 4294967295
    v1 = v1
    if (v5 != v5) {
        fmt.Println(v1, v1)
    }
    fmt.Println(v5)
    delete(v4, 4294967295)
    var v12 uint = 1
    var v13 uint16 = 65535
    var v14 uint16 = v13
    switch 127 {
    case 9223372036854775807:
        fmt.Println(v12)
    case -278:
        v15 := GenS2{F0: 233, F1: (^(v1))}
        _ = v15
    case -137:
        var v16 string = v3
        _ = v16
    default:
        v2 = append(v2, 9223372036854775807)
    }
    if false {
        v17 := GenS3{F0: 42}
        switch "\n\t" {
        case "":
            var v18 int16 = 412
            _ = v18
        default:
            v19 := GenS2{F0: uint8(v12), F1: 309}
            _ = v19
        }
        _ = v17
    }
    if (!true) {
        if true {
            if (1 >= -1) {
                _ = len(v4)
                v14 -= v13
            }
            v3 = v3
            if true {
                v20 := GenS3{F0: 2147483647}
                var v21 uint64 = uint64(v1)
                v22 := v4[int(v5)]
                _ = v20
                _ = v21
                _ = v22
            }
        }
        fmt.Println(v1, v14)
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint8 = uint8(v12)
        v24 += v5
        _ = fi23
        _ = v24
    }
    v6 += (v6 / ((v6) | 1))
    var v25 uint = (^(v12))
    _ = len(v2)
    fmt.Println(v25, v12, v25)
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v2)
            _ = fi27
        }
        v28 := GenS3{F0: 2147483647}
        _ = fi26
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
