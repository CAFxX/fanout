package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint64; F2 uint }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0(p0 int, p1 uint32, p2 bool) bool {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 string = "hello"
        _ = fi5
        _ = v6
    }
    p0 = p0
    fmt.Println(p0, p1, p1)
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return ((!(-740 != -1)) || (-1 >= 32767))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1000
    var v3 uint32 = 0
    var v4 uint16 = 0
    if false {
        v7 := []int{}
        _ = v7
    }
    v8 := []int{127}
    if (32768 != -99) {
        fmt.Println(v3, v3, v3)
        var v9 string = "hello"
        _ = v9
    } else {
        v10 := v1[("" + "foo bar")]
        switch 42 {
        case 2147483647:
            v11 := GenS1{F0: 4, F1: (v3 << ((1) & 63)), F2: uint8(v4)}
            _ = v11
        }
        switch 127 {
        case -97:
            if (8 >= uint8(v10)) {
                fmt.Println(v3, v2)
            } else {
                if false {
                    var v12 int32 = 0
                    if len(v8) > 0 {
                        v8[0] = 4294967295
                    }
                    _ = v12
                }
            }
        case 4294967295:
            switch "" {
            case "":
                v13 := GenS2{F0: 13}
                _ = v13
            }
        case 32767:
            var v14 uint8 = 157
            _ = v14
        }
        _ = v10
    }
    fmt.Println(v3, v2)
    if len(v8) > 0 {
        v8[0] = -838
    }
    v8 = append(v8, 1)
    fmt.Println(v4, v3, v2)
    var v15 float32 = (0.1 * (float32(1e10) * ((0.1 * (1e10 * 1.5)) - -2.25)))
    switch "a" {
    case "hello":
        var v16 int8 = 103
        _ = v16
    case "\n\t":
        if ((v3 * (v3 + v3)) >= v3) {
            v17 := GenS3{F0: 4294967295, F1: 4294967295, F2: 266}
            v2 -= (v2 % (((v2 & v2)) | 1))
            v2 += (v2 >> ((255) & 63))
            _ = v17
        } else {
            v8 = append(v8, 1)
            switch 32768 {
            case 128:
                fmt.Println(v4, v4, v4)
            case 629:
                var v18 float64 = (0.1 / 0.1)
                _ = v18
            case 402:
                switch ("\n\t" + "a") {
                case "foo bar":
                    var v19 int = -997
                    _ = v19
                case "":
                    v20 := GenS4{F0: 42, F1: (^(v4))}
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 bool = (v4 >= 42)
                _ = fi21
                _ = v22
            }
        }
    default:
        var v23 uint64 = 2
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
