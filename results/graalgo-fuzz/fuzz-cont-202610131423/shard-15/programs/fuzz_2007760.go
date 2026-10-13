package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint8 }
func genFn0(p0 int16) string {
    var v7 int16 = p0
    p0 += 1000
    var v8 uint32 = 4294967295
    for fi9 := 0; fi9 < 15; fi9++ {
        v7 -= (-(v7))
        fmt.Println(v8, p0, p0)
        _ = fi9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        if false {
            var v11 uint8 = 255
            switch 32768 {
            case 42:
                fmt.Println(v11, v11, fi10)
            }
            _ = v11
        }
        v8 = 222
        _ = fi10
    }
    _ = p0
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := []string{("a" + (("" + "hello") + "\n\t")), (((("\n\t" + "foo bar") + ("" + ((("hello" + "hello") + ("a" + ("foo bar" + "\n\t"))) + "hello"))) + "") + "a")}
    var v2 int8 = 1
    var v3 float64 = ((0.0 / 1.0) * 3.14159)
    v4 := []int{130, 546, -2147483648, 32768}
    var v5 int = 32768
    var v6 uint32 = 2147483647
    if (!true) {
        fmt.Println(v5)
        switch ("a" + "hello") {
        case "":
            var v12 int64 = 42
            _ = v12
        case "hello":
            v5 -= v5
        case "a":
            var v13 int8 = 100
            _ = v13
        default:
            fmt.Println(v6, v5, v3)
        }
        if (v6 <= (v6 | (^(v6)))) {
            v1 = append(v1, "a")
            v6 += v6
            v14 := GenS1{F0: false}
            _ = v14
        }
    }
    fmt.Println(v6, v3)
    v15 := GenS2{F0: 1, F1: v2, F2: 1}
    v16 := GenS3{F0: 428, F1: 197, F2: 0}
    switch "hello" {
    case "\n\t":
        v17 := GenS2{F0: -240, F1: (-(v2)), F2: 0}
        _ = v17
    case "foo bar":
        switch "" {
        case "\n\t":
            _ = len(v4)
        default:
            var v18 int64 = -464
            _ = v18
        }
    }
    if true {
        v6 -= v6
        if false {
            if (1000 == int16(v2)) {
                v3 += v3
            } else {
                v2 -= -128
                switch "" {
                case "foo bar":
                    var v19 int64 = -1
                    _ = v19
                default:
                    v3 += 1.5
                }
                v6 = (v6 % (((v6 + (v6 + 4294967295))) | 1))
            }
        } else {
            var v20 uint8 = 255
            var v21 bool = true
            fmt.Println(v2, v5, v6)
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 5; fi22++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi22
        }
    } else {
        _ = len(v1)
        v4 = append(v4, 4294967295)
        v1 = append(v1, "\n\t")
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        if (!((((1e10 + (0.0 / (0.1 + 1e10))) / (float32(0.1) * 1.5)) != 1.5) && true)) {
            fmt.Println(v2, fi23)
        }
        var v24 uint = 385
        fmt.Println(v5, v24, fi23)
        _ = fi23
        _ = v24
    }
    var v25 uint16 = uint16(v6)
    var v26 string = ""
    v3 -= float64(float32(1e10))
    var v27 uint8 = 65
    v27 = 42
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 uint = 1000
        var v30 uint8 = v27
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint16 = (^(v25))
    var v32 float32 = 0.1
    fmt.Println(v3, v32)
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 19; fi34++ {
            v4 = append(v4, (fi34 | 774))
            v27 += (v27 - (v27 ^ v27))
            fmt.Println(v3, v31)
            _ = fi34
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            v36 := GenS3{F0: 1, F1: 2, F2: (v27 * v27)}
            var v37 int32 = 42
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
