package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 bool) int32 {
    v7 := []int{383, -9223372036854775808}
    v8 := GenS1{F0: 197, F1: 1}
    if len(v7) > 0 {
        v7[0] = 1
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 42
    var v3 int8 = -128
    var v4 uint32 = v2
    var v5 int32 = -277
    v6 := make(map[string]string)
    switch "a" {
    case "hello":
        var v9 string = "\n\t"
        _ = v9
    default:
        var v10 int32 = 145
        _ = v10
    }
    v11 := GenS1{F0: 815, F1: -334}
    var v12 string = ("\n\t" + "")
    var v13 float64 = ((float64((-2.25 + 0.1)) / 0.1) - float64(1e10))
    for fi14 := 0; fi14 < 20; fi14++ {
        if ((true || false) || false) {
            fmt.Println(v13)
            var v15 uint16 = 2
            v13 -= v13
            _ = v15
        } else {
            v16 := GenS2{F0: (^(v2))}
            _ = v16
        }
        if (!(((v5 > 2147483647) && true) || ((v3 ^ (v3 << ((uint(v1)) & 63))) != -98))) {
            var v17 uint16 = uint16(v1)
            if true {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := GenS3{F0: (v17 - 503)}
                    fmt.Println(v4)
                    _ = fi18
                    _ = v19
                }
            }
            _ = v17
        }
        var v20 uint = 65535
        _ = fi14
        _ = v20
    }
    var v21 int64 = 777
    switch 522 {
    case 54:
        var v22 uint32 = v4
        _ = v22
    }
    var v23 int64 = (v21 & 9223372036854775807)
    var v24 float64 = ((((((0.1 * (v13 / v13)) * v13) * (3.14159 - v13)) - (v13 - v13)) + v13) * (1.5 - v13))
    _ = len(v6)
    if true {
        for fi25 := 0; fi25 < 17; fi25++ {
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 int16 = -1
                if false {
                    fmt.Println(v21, v5, v4)
                }
                v28 := GenS3{F0: 489}
                var v29 string = ("" + v12)
                v30 := v6[(("a" + "a") + (v12 + v29))]
                var v31 int8 = 127
                if (!false) {
                    var v32 string = ((v29 + v29) + v29)
                    _ = v32
                }
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
            v33 := []int{(fi25 & (fi25 * 2147483647))}
            _ = fi25
            _ = v33
        }
    } else {
        var v34 uint8 = v1
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS1{F0: 1, F1: fi35}
        _ = fi35
        _ = v36
    }
    v12 = "\n\t"
    v37 := []int{127, -2147483648, 539}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
