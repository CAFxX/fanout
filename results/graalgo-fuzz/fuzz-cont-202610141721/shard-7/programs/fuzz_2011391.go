package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 string }
func genFn0() int64 {
    var v5 int16 = 42
    switch 2147483647 {
    case 42:
        if (float64(3.14159) >= ((1.5 - 3.14159) - 3.14159)) {
            v6 := []int{-155, -439, 128, 732}
            _ = v6
        }
    case -188:
        for fi7 := 0; fi7 < 6; fi7++ {
            switch 1 {
            case -9223372036854775808:
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v5, fi7)
                default:
                    v8 := []int{fi7, 4294967295, (fi7 ^ fi7), -933}
                    _ = v8
                }
            case 4294967295:
                var v9 bool = (false || ((true || true) && false))
                _ = v9
            default:
                v5 = (^(v5))
            }
            if true {
                v10 := []int{fi7, fi7, fi7}
                _ = v10
            }
            _ = fi7
        }
    case -450:
        if false {
            v5 = (v5 * (v5 | 1))
        } else {
            if true {
                v5 -= -879
            } else {
                fmt.Println(v5)
                for fi11 := 0; fi11 < 20; fi11++ {
                    v5 += (^(v5))
                    v5 = -32768
                    var v12 uint = 65535
                    _ = fi11
                    _ = v12
                }
                var v13 bool = ((0.1 / -2.25) <= 3.14159)
                _ = v13
            }
            v5 = (v5 << ((255) & 63))
            if ("" >= ("\n\t" + "")) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS1{F0: 0, F1: -660}
                    _ = fi14
                    _ = v15
                }
                var v16 bool = true
                v17 := []int{int(v5), 81}
                _ = v16
                _ = v17
            }
        }
    }
    v5 += 190
    v18 := GenS2{F0: 18446744073709551615, F1: 1}
    v19 := []int{2147483647}
    _ = v18
    _ = v19
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := []uint32{663}
    var v2 int16 = 144
    v3 := make(map[int]int)
    var v4 int8 = -73
    v2 = v2
    var v20 string = (((("hello" + "hello") + "a") + ("a" + ("a" + "hello"))) + "\n\t")
    if ((!true) && (!((-9223372036854775808 >= 1) && false))) {
        fmt.Println(v20)
        var v21 int16 = (v2 + v2)
        v20 = ((v20 + v20) + v20)
        _ = v21
    }
    var v22 int32 = 0
    var v23 int = int(v22)
    switch "a" {
    case "\n\t":
        switch "foo bar" {
        case "hello":
            var v24 uint16 = 32767
            _ = v24
        }
    case "a":
        v3[(-(v23))] = (v23 ^ (v23 | v23))
    default:
        switch v23 {
        case 593:
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 uint = 18446744073709551615
                _ = fi25
                _ = v26
            }
        case -2147483648:
            v27 := GenS2{F0: 0, F1: 65535}
            _ = v27
        case -208:
            var v28 float32 = -2.25
            _ = v28
        }
    }
    var v29 float64 = float64(1.5)
    var v30 int16 = v2
    switch v20 {
    case "":
        v31 := GenS1{F0: (v30 - (v2 | v2)), F1: (v23 ^ int(v4))}
        _ = v31
    default:
        switch v23 {
        case 965:
            _ = len(v3)
        case 128:
            fmt.Println(v29, v2)
        case 32767:
            v32 := GenS3{F0: 4294967295, F1: "hello"}
            _ = v32
        default:
            fmt.Println(v30)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
