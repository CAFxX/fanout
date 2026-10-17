package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int64; F1 bool }
func genFn0(p0 int16) uint8 {
    v6 := []int{-2147483648, 85}
    switch "" {
    case "hello":
        switch ((((("hello" + "a") + "\n\t") + ("" + ((("" + ("foo bar" + "")) + (("hello" + ((("" + "\n\t") + "foo bar") + "hello")) + "\n\t")) + ""))) + ("a" + "hello")) + "hello") {
        case "":
            fmt.Println(p0)
        case "\n\t":
            v7 := GenS1{F0: 0, F1: -1}
            _ = v7
        default:
            fmt.Println(p0)
        }
    case "foo bar":
        var v8 uint = 488
        _ = v8
    }
    switch "foo bar" {
    case "foo bar":
        if (!(false || (true || (false && true)))) {
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 int64 = 1
                    var v12 bool = ((p0 & p0) <= p0)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: 1, F1: 42}
                var v14 float64 = (((float64((0.0 / float32((3.14159 * float64((3.14159 - float32(float64((1.5 / (float32(((0.0 * (1.5 / 1e10)) - (float64(1e10) * float64(1e10)))) * 1.5)))))))))) / 3.14159) / (float64((((0.0 + 1e10) - ((-2.25 - ((3.14159 / 1.5) - float32(((((0.1 * float64(0.1)) / ((3.14159 + 1e10) / 3.14159)) * (((0.1 - 1.5) / -2.25) / float64(1e10))) - 0.0)))) - (0.0 * (3.14159 * 1e10)))) + (-2.25 + ((((float32((float64(-2.25) + (1e10 - 3.14159))) * -2.25) + 0.0) - 0.1) / 1.0)))) * (((3.14159 * (3.14159 - 3.14159)) + 0.1) + ((((0.0 / 1e10) + (0.1 / float64(float32((0.1 + (1.5 / (1.5 + (0.0 * -2.25)))))))) - 1.5) + float64(3.14159))))) + ((0.1 + -2.25) / 1.5))
                _ = fi9
                _ = v13
                _ = v14
            }
            p0 -= 0
        }
    case "hello":
        v15 := GenS1{F0: 1, F1: 955}
        _ = v15
    case "":
        var v16 uint = 65535
        _ = v16
    }
    v6 = append(v6, 2147483647)
    _ = p0
    _ = v6
    return 238
}

func main() {
    var v1 int = 127
    var v2 bool = (48 < 128)
    var v3 int64 = 9223372036854775807
    var v4 int16 = 32767
    var v5 string = "foo bar"
    var v17 int32 = 481
    v4 += (v4 - (v4 + (v4 - (v4 & -32768))))
    v18 := GenS1{F0: v3, F1: (v17 ^ v17)}
    var v19 uint32 = 2147483647
    v2 = false
    var v20 bool = v2
    switch v5 {
    case "a":
        fmt.Println(v4)
    case "\n\t":
        var v21 int16 = (v4 + 497)
        _ = v21
    }
    v22 := GenS1{F0: (v3 / (((v3 - v3)) | 1)), F1: (v17 | int32(v1))}
    for fi23 := 0; fi23 < 19; fi23++ {
        v17 = v17
        _ = fi23
    }
    var v24 uint16 = 0
    var v25 bool = v20
    fmt.Println(v1)
    if false {
        v25 = (65535 != 4294967295)
        v4 -= v4
        var v26 uint16 = 86
        _ = v26
    } else {
        fmt.Println(v1)
        v27 := []int{}
        _ = v27
    }
    if (v20 && true) {
        if true {
            v28 := []int{32768, v1, v1, -719}
            _ = v28
        }
        if (v1 >= v1) {
            v4 += 185
            var v29 uint64 = 1
            fmt.Println(v25, v17)
            _ = v29
        }
        v30 := GenS2{F0: (255 < 73), F1: (-(v4))}
        _ = v30
    }
    var v31 int = (-(v1))
    var v32 bool = (v4 != (v4 * 14))
    for fi33 := 0; fi33 < 9; fi33++ {
        v25 = v2
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 uint16 = 97
            if (!(((0.1 - 0.1) + ((1e10 - (float32(-2.25) * (0.1 * 0.1))) * float32(0.0))) > ((-2.25 - 1.5) - (((3.14159 / -2.25) - 0.0) + 0.1)))) {
                var v36 uint = 500
                v37 := []int{}
                var v38 int8 = 42
                _ = v36
                _ = v37
                _ = v38
            } else {
                v39 := GenS3{F0: (v3 << ((1) & 63)), F1: v2}
                v40 := GenS2{F0: true, F1: (v4 | (-(v4)))}
                _ = v39
                _ = v40
            }
            _ = fi34
            _ = v35
        }
        v41 := GenS1{F0: (v3 % (((v3 * v3)) | 1)), F1: v17}
        _ = fi33
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
