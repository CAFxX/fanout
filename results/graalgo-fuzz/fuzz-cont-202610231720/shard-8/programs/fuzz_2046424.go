package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 bool; F1 int; F2 uint }
type GenS3 struct { F0 string; F1 bool; F2 string }
func genFn0(p0 uint, p1 uint64) int16 {
    v7 := []int{-2147483648}
    v8 := GenS1{F0: 0, F1: -1, F2: 18446744073709551615}
    fmt.Println(p1)
    p1 -= p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 float64 = 0.0
    var v2 int = 128
    v3 := make(map[int]string)
    var v4 uint16 = 42
    v5 := []uint32{319, 2147483647, 0, 246}
    v6 := make(map[string]int)
    var v9 uint = 1
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch v2 {
            case 1:
                v5 = append(v5, 2147483648)
            case -2147483648:
                var v12 int = fi10
                _ = v12
            case 128:
                var v13 int64 = 897
                if ((float64((-2.25 * 3.14159)) / (-2.25 * (v1 / v1))) <= v1) {
                    var v14 uint64 = 710
                    _ = v14
                }
                _ = v13
            default:
                v15 := GenS1{F0: 1, F1: 639, F2: (v9 & (v9 - (v9 & uint(v4))))}
                _ = v15
            }
            _ = fi11
        }
        v16 := GenS1{F0: 974, F1: 1, F2: (v9 + v9)}
        v1 -= v1
        _ = fi10
        _ = v16
    }
    fmt.Println(v1)
    v17 := GenS2{F0: false, F1: (v2 + (v2 ^ v2)), F2: (^(v9))}
    v18 := GenS1{F0: 489, F1: -313, F2: 255}
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 float32 = 3.14159
        var v21 bool = false
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint16 = v4
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := GenS1{F0: 4294967295, F1: -464, F2: 1}
        v25 := GenS3{F0: "foo bar", F1: ((623 != 1) || (!(false && (("" + "foo bar") < "hello")))), F2: "hello"}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 bool = false
    v27 := GenS1{F0: 0, F1: -9223372036854775808, F2: 795}
    var v28 bool = v26
    if (32767 == -963) {
        switch v2 {
        case 9223372036854775807:
            v29 := GenS1{F0: 2147483647, F1: 0, F2: (^(v9))}
            _ = v29
        default:
            var v30 int = 128
            _ = v30
        }
        var v31 bool = false
        _ = v31
    } else {
        if (!(v28 || ((((!(-2147483648 > -2147483648)) || (v2 > v2)) && (((v9 - v9) != (v9 + v9)) && (v26 && ((!(!v28)) && true)))) && (0 >= -2147483648)))) {
            if (!(!(v4 >= uint16(v9)))) {
                var v32 string = "a"
                var v33 uint8 = 127
                _ = v32
                _ = v33
            }
            v34 := GenS1{F0: 42, F1: 1, F2: (v9 * (^(v9)))}
            v28 = ((v26 && (42 == 41)) || ((0 != -59) || (1 < -1)))
            _ = v34
        }
        switch ("\n\t" + ("hello" + "foo bar")) {
        case "foo bar":
            var v35 string = "\n\t"
            _ = v35
        case "\n\t":
            var v36 uint64 = 908
            _ = v36
        default:
            if ((-32768 <= -32768) || (int64(v22) == 9223372036854775807)) {
                fmt.Println(v1, v28, v28)
                var v37 float32 = (-2.25 - 1e10)
                _ = v37
            }
        }
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        if (!true) {
            v4 += 255
            v1 = ((((v1 / 0.1) + v1) / (-2.25 - (0.1 / -2.25))) + 1e10)
            if (-47 > -721) {
                var v39 int8 = 100
                fmt.Println(v9, v28, v4)
                v39 -= (v39 << ((uint(v4)) & 63))
                _ = v39
            }
        }
        var v40 uint8 = 146
        delete(v3, v2)
        _ = fi38
        _ = v40
    }
    v1 = float64(((3.14159 - ((0.1 - (float32((v1 - v1)) * -2.25)) - (0.1 / (0.0 + (0.0 + (((1e10 / (float32(0.1) / 0.1)) + 1.5) + (float32(v1) * 0.0))))))) / (((float32(v1) + 3.14159) / float32((v1 + v1))) - ((float32(0.1) + (-2.25 * ((3.14159 / 3.14159) + -2.25))) * 0.1))))
    v9 = (v9 / (((v9 * v9)) | 1))
    for fi41 := 0; fi41 < 14; fi41++ {
        for fi42 := 0; fi42 < 13; fi42++ {
            switch "a" {
            case "a":
                v5 = append(v5, 689)
            case "hello":
                var v43 uint = (v9 / ((v9) | 1))
                _ = v43
            }
            _ = fi42
        }
        var v44 int8 = 1
        var v45 uint = (^(v9))
        _ = fi41
        _ = v44
        _ = v45
    }
    v5 = append(v5, uint32(v4))
    v9 += v9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
