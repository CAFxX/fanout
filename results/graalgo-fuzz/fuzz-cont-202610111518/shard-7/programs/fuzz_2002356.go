package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint }
func genFn0() bool {
    for fi4 := 0; fi4 < 2; fi4++ {
        if (-32768 != 858) {
            var v5 uint64 = 1
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 int32 = -2147483648
                _ = fi6
                _ = v7
            }
            var v8 int32 = 0
            _ = v5
            _ = v8
        }
        var v9 uint = 2
        v10 := []int{}
        _ = fi4
        _ = v10
        _ = v9
    }
    var v11 float64 = -2.25
    switch 9223372036854775807 {
    case 2147483647:
        var v12 uint32 = 1
        _ = v12
    default:
        var v13 uint16 = 65535
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(fi14, fi14)
        v11 += 0.0
        _ = fi14
    }
    _ = v11
    return false
}

func main() {
    var v1 int = 42
    v2 := []string{"a", ((("foo bar" + (("" + (("\n\t" + ("hello" + ("hello" + "a"))) + "foo bar")) + (("foo bar" + "foo bar") + ("foo bar" + "foo bar")))) + "\n\t") + "hello")}
    v3 := make(map[int]uint32)
    var v15 uint = 42
    v15 -= (v15 << (((v15 >> (((^(v15))) & 63))) & 63))
    v16 := GenS1{F0: (false && false), F1: 430}
    v1 += -1
    switch ((("a" + "") + "a") + "\n\t") {
    case "foo bar":
        v17 := GenS2{F0: 65535, F1: ("a" + "foo bar"), F2: 42}
        _ = v17
    }
    var v18 uint64 = 42
    switch v1 {
    case 4294967295:
        var v19 int32 = 0
        _ = v19
    case 1:
        var v20 int = (v1 - v1)
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS1{F0: (0.1 > (1e10 - -2.25)), F1: 2}
        if (false || (("" + "a") < (((("hello" + "foo bar") + "a") + "foo bar") + "\n\t"))) {
            var v23 string = "hello"
            var v24 float64 = 0.1
            _ = v23
            _ = v24
        }
        var v25 float64 = 1.5
        _ = fi21
        _ = v22
        _ = v25
    }
    fmt.Println(v18, v1)
    var v26 uint64 = v18
    switch "hello" {
    case "a":
        for fi27 := 0; fi27 < 8; fi27++ {
            if true {
                if ((float32(((0.1 - float64((1e10 + 0.0))) + 3.14159)) * (((-2.25 - 1e10) - 1e10) + (1e10 + 3.14159))) <= -2.25) {
                    v28 := GenS3{F0: (((-54 <= 1) && (v1 == (v1 % (((v1 | -9223372036854775808)) | 1)))) && (0 != 42)), F1: 1, F2: (v15 ^ 557)}
                    var v29 float32 = 0.0
                    for fi30 := 0; fi30 < 18; fi30++ {
                        v31 := GenS1{F0: true, F1: 42}
                        _ = fi30
                        _ = v31
                    }
                    _ = v28
                    _ = v29
                }
                var v32 uint32 = 42
                var v33 uint64 = 751
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 bool = false
                _ = fi34
                _ = v35
            }
            switch v1 {
            case -1:
                fmt.Println(v18, v26)
            }
            _ = fi27
        }
    default:
        var v36 float32 = float32(((float64(float32(1e10)) - 0.1) / (((-2.25 / 0.1) / 1.0) / (((1e10 / -2.25) / 1.0) / 0.1))))
        _ = v36
    }
    var v37 int32 = -973
    v18 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v37
}
