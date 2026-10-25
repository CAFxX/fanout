package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 int32; F1 int32; F2 uint }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 uint8) uint8 {
    var v7 int64 = 9223372036854775807
    p1 -= 1
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := []int{-499, -2147483648, 4294967295, -1}
    var v2 string = "\n\t"
    var v3 float32 = 3.14159
    var v4 uint8 = 226
    var v5 int16 = 1000
    v6 := make(map[string]int)
    v8 := GenS1{F0: 54, F1: 2, F2: -100}
    fmt.Println(v2)
    v3 += 1.5
    if true {
        v3 += (v3 - v3)
        var v9 uint = 4294967295
        var v10 uint8 = (v4 >> (((v9 * 65535)) & 63))
        _ = v10
        _ = v9
    }
    var v11 int64 = 42
    var v12 uint32 = 42
    v3 += float32(((((float64((float32(-2.25) + float32(((float64((v3 * (1e10 / (1e10 + v3)))) + float64(1.5)) / 3.14159)))) - (((0.1 + (0.1 * -2.25)) - float64((v3 * v3))) + 0.0)) / (-2.25 - 3.14159)) + 3.14159) - float64(0.1)))
    v13 := GenS1{F0: -27, F1: 0, F2: -925}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int = -446
        v16 := GenS2{F0: -1, F1: -222, F2: 2}
        fmt.Println(v4, v2, v4)
        _ = fi14
        _ = v15
        _ = v16
    }
    switch -2147483648 {
    case 1:
        for fi17 := 0; fi17 < 6; fi17++ {
            v5 += -32768
            if false {
                var v18 uint64 = 42
                _ = v18
            } else {
                for fi19 := 0; fi19 < 12; fi19++ {
                    switch "hello" {
                    case "\n\t":
                        switch ((v2 + "") + v2) {
                        case "foo bar":
                            v20 := GenS3{F0: (-(fi19))}
                            _ = v20
                        case "":
                            v3 -= (v3 * (0.0 - ((((v3 * ((((float32((1e10 - ((3.14159 - 0.0) + (3.14159 * 1.5)))) * (v3 * v3)) / v3) * (v3 + float32((0.1 - 1.5)))) + v3)) + 0.0) - 0.0) / v3)))
                        default:
                            var v21 string = "foo bar"
                            _ = v21
                        }
                    case "a":
                        _ = len(v6)
                    default:
                        v4 = v4
                    }
                    fmt.Println(v2, fi17, v3)
                    _ = fi19
                }
                var v22 string = (v2 + v2)
                _ = v22
            }
            var v23 int16 = v5
            _ = fi17
            _ = v23
        }
    case -1:
        v24 := GenS3{F0: 613}
        _ = v24
    }
    if false {
        delete(v6, "")
        switch ("" + ((("hello" + "") + ((("" + v2) + v2) + v2)) + "\n\t")) {
        case "foo bar":
            fmt.Println(v5)
        case "\n\t":
            v25 := GenS2{F0: 0, F1: 1, F2: 255}
            _ = v25
        }
        v5 = 1000
    } else {
        if (-1 <= (v11 / (((v11 * v11)) | 1))) {
            v1 = append(v1, 657)
            var v26 float64 = 3.14159
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := GenS2{F0: 2147483647, F1: 1, F2: 18446744073709551615}
                _ = fi27
                _ = v28
            }
            _ = v26
        }
    }
    var v29 int64 = int64(v12)
    v29 = (v29 | v11)
    var v30 int32 = int32(v4)
    for fi31 := 0; fi31 < 19; fi31++ {
        switch (fi31 * fi31) {
        case 688:
            var v32 int32 = v30
            _ = v32
        case -1:
            for fi33 := 0; fi33 < 10; fi33++ {
                v34 := GenS2{F0: v30, F1: v30, F2: 18446744073709551615}
                v4 -= v4
                _ = fi33
                _ = v34
            }
        case 1:
            var v35 uint64 = 4294967295
            _ = v35
        default:
            v3 = 1e10
        }
        _ = fi31
    }
    v29 = v29
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
