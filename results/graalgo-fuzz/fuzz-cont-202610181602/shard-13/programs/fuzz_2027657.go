package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 bool; F1 int64; F2 string }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint8; F1 int32; F2 string }

func main() {
    var v1 float32 = (float32(0.0) - float32((1e10 * float64(1.5))))
    var v2 float32 = 0.1
    var v3 bool = true
    var v4 float32 = 0.0
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 65535, F1: -128}
    switch ("foo bar" + "hello") {
    case "\n\t":
        if v3 {
            var v7 int8 = 1
            _ = v7
        } else {
            var v8 uint16 = 42
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(fi9, fi9, v1)
                if false {
                    v11 := GenS1{F0: 255, F1: -71}
                    _ = v11
                }
                var v12 int = 702
                _ = fi10
                _ = v12
            }
            var v13 int8 = -1
            v14 := []int{}
            _ = fi9
            _ = v13
            _ = v14
        }
    }
    switch ("" + "") {
    case "foo bar":
        v15 := GenS2{F0: false, F1: -9223372036854775808, F2: "a"}
        _ = v15
    case "a":
        v1 -= (3.14159 * 0.0)
    default:
        var v16 float64 = (3.14159 / ((((0.1 * 0.0) * 1e10) / ((((0.0 * 0.1) - (1.5 - 1.5)) - float64(3.14159)) - 0.1)) - 3.14159))
        _ = v16
    }
    v17 := GenS3{F0: 255, F1: -404}
    _ = len(v5)
    if false {
        fmt.Println(v1, v1, v2)
    } else {
        v18 := []int{}
        v2 = v1
        _ = v18
    }
    fmt.Println(v4, v3)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS1{F0: 255, F1: -72}
        for fi21 := 0; fi21 < 12; fi21++ {
            if true {
                var v22 uint16 = 32767
                _ = v22
            }
            v23 := GenS3{F0: 1, F1: 42}
            _ = fi21
            _ = v23
        }
        fmt.Println(v3, fi19)
        _ = fi19
        _ = v20
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v1 = ((v4 - v1) - v2)
        v3 = (!v3)
        v25 := GenS4{F0: -1, F1: 32767}
        _ = fi24
        _ = v25
    }
    var v26 int32 = 1
    v26 = v26
    switch -103 {
    case -2147483648:
        fmt.Println(v1, v2)
    case 0:
        if (!v3) {
            if (32767 > 0) {
                v27 := []int{0, -327, -351, -904}
                v26 += (v26 >> ((0) & 63))
                var v28 int64 = -1
                _ = v27
                _ = v28
            } else {
                v29 := []int{32767, 105, 2147483647}
                v30 := GenS5{F0: 606, F1: 128}
                _ = v29
                _ = v30
            }
        }
    default:
        var v31 uint = 228
        _ = v31
    }
    v32 := GenS3{F0: 32767, F1: -9223372036854775808}
    var v33 int = -112
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 float32 = 1.5
        var v36 uint64 = 18446744073709551615
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 int64 = -949
    v38 := GenS6{F0: 42, F1: (v26 ^ v26), F2: "a"}
    v3 = ((v4 >= (v2 / v1)) && v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
