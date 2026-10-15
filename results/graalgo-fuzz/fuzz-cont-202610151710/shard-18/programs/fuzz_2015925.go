package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0() uint32 {
    if (844 != 215) {
        if ((true && (-1 == -45)) && (1.5 > 1.5)) {
            if false {
                var v4 int = -2147483648
                v5 := []int{(v4 / ((9223372036854775807) | 1)), (v4 << ((18446744073709551615) & 63)), -1}
                _ = v4
                _ = v5
            }
            v6 := []int{}
            var v7 int32 = -39
            _ = v6
            _ = v7
        }
        switch 1 {
        case 32767:
            var v8 int16 = -1
            _ = v8
        case -1:
            v9 := []int{-9223372036854775808, -9223372036854775808, 42, 32767}
            _ = v9
        default:
            if true {
                var v10 int16 = 1
                v11 := []int{0, 127, -871, 183}
                v12 := GenS1{F0: 9, F1: 1}
                _ = v10
                _ = v11
                _ = v12
            } else {
                var v13 float32 = ((1.5 + -2.25) * 0.0)
                _ = v13
            }
        }
    } else {
        var v14 bool = true
        fmt.Println(v14)
        if false {
            fmt.Println(v14)
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := []int{}
                var v17 float64 = float64(3.14159)
                v18 := GenS1{F0: 747, F1: 86}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 uint64 = 909
            _ = v19
        } else {
            v14 = (-1 >= -128)
        }
        _ = v14
    }
    if ((!true) || false) {
        v20 := []int{-9223372036854775808}
        var v21 int = 42
        if len(v20) > 0 {
            v20[0] = 42
        }
        _ = v20
        _ = v21
    } else {
        v22 := []int{}
        var v23 uint8 = 127
        v23 -= 0
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := GenS1{F0: 784, F1: 2}
        var v26 int32 = 2147483647
        _ = fi24
        _ = v25
        _ = v26
    }
    return 2147483648
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 2147483647
    v3 := make(map[int]uint32)
    v2 = 819
    _ = len(v1)
    v2 -= (v2 - v2)
    var v27 uint = 383
    if true {
        delete(v3, -1)
        v28 := GenS1{F0: (v2 + 65535), F1: 2}
        _ = v28
    } else {
        v29 := []int{32767, 2147483647}
        if (true || ((324 == -52) && true)) {
            var v30 int8 = -80
            if (((false && (v30 < 127)) && (false && false)) || true) {
                fmt.Println(v27, v30, v2)
                var v31 string = ""
                _ = v31
            }
            v30 += v30
            _ = v30
        } else {
            v2 += (v2 % (((v2 >> ((4294967295) & 63))) | 1))
            v27 = 0
        }
        _ = v29
    }
    var v32 float32 = 1.5
    v27 -= v27
    for fi33 := 0; fi33 < 14; fi33++ {
        v34 := GenS1{F0: 936, F1: 18446744073709551615}
        var v35 float64 = (0.1 + 0.0)
        fmt.Println(v35)
        _ = fi33
        _ = v34
        _ = v35
    }
    v32 -= 0.0
    var v36 int64 = 825
    v32 = (v32 + (float32(-2.25) / 3.14159))
    var v37 string = "\n\t"
    v2 -= v2
    if (32767 >= -57) {
        if ((1 == -722) && (!true)) {
            v2 += (v2 >> ((uint(v2)) & 63))
            var v38 uint32 = (v2 ^ v2)
            var v39 int32 = 2147483647
            _ = v38
            _ = v39
        } else {
            var v40 int16 = int16(v27)
            switch ("a" + (v37 + v37)) {
            case "a":
                var v41 int16 = v40
                _ = v41
            case "foo bar":
                fmt.Println(v32, v32, v2)
            }
            fmt.Println(v37)
            _ = v40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v37
}
