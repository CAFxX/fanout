package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int64; F1 string; F2 bool }
func genFn0(p0 uint16, p1 uint, p2 int) int {
    var v4 int64 = 9223372036854775807
    v5 := GenS1{F0: p0, F1: (v4 % ((v4) | 1)), F2: p0}
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(fi6, v4, p1)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1(p0 uint64) string {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float32 = (1.5 + 1.5)
        fmt.Println(fi7, v8)
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint = 4294967295
            v11 := GenS2{F0: -187}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS2{F0: 32767}
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS3{F0: 1, F1: 1, F2: (-128 == -31)}
        _ = fi13
        _ = v14
    }
    if (!true) {
        p0 += (p0 >> ((4294967295) & 63))
        v15 := []int{32767, genFn0(1, 960, 0)}
        _ = v15
    }
    _ = p0
    _ = v12
    return "\n\t"
}

func main() {
    var v1 float64 = float64(3.14159)
    v2 := make(map[int]string)
    var v3 uint32 = 255
    v16 := GenS4{F0: 0, F1: 2, F2: 0}
    v17 := v2[genFn0(743, 200, 364)]
    v18 := []int{1}
    if false {
        fmt.Println(v1)
        var v19 float32 = (float32((1.5 - ((float64(0.1) - (v1 * v1)) + ((v1 / float64((((-2.25 + 3.14159) / (3.14159 + 0.1)) + (float32(1.5) + 1.5)))) - (v1 + v1))))) - 1e10)
        _ = v19
    } else {
        switch v17 {
        case "foo bar":
            v20 := GenS1{F0: 42, F1: -857, F2: 1}
            _ = v20
        default:
            v1 = 1.5
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            switch fi21 {
            case 128:
                var v22 int64 = -197
                _ = v22
            case 2147483647:
                v1 = v1
            case 767:
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 string = v17
                    var v25 uint32 = 996
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            _ = fi21
        }
        v26 := GenS5{F0: 1, F1: (v17 + "hello"), F2: false}
        _ = v26
    }
    switch 32768 {
    case 349:
        v1 = float64(-2.25)
    }
    var v27 uint32 = 42
    for fi28 := 0; fi28 < 14; fi28++ {
        switch fi28 {
        case 127:
            for fi29 := 0; fi29 < 3; fi29++ {
                switch (fi29 + (fi29 + fi28)) {
                case 9223372036854775807:
                    delete(v2, -9223372036854775808)
                case -403:
                    fmt.Println(v1, v17)
                }
                switch (((((v17 + "foo bar") + v17) + v17) + (v17 + v17)) + "") {
                case "\n\t":
                    v1 = 3.14159
                }
                _ = fi29
            }
        case 32768:
            v17 = v17
        default:
            v3 -= v27
        }
        _ = fi28
    }
    var v30 float32 = (3.14159 * 1e10)
    v31 := v2[32768]
    var v32 uint8 = 0
    switch (("" + (v31 + v31)) + ("\n\t" + "a")) {
    case "hello":
        var v33 bool = false
        _ = v33
    case "\n\t":
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := GenS3{F0: (v32 % ((v32) | 1)), F1: 91, F2: (false && false)}
            _ = fi34
            _ = v35
        }
    }
    var v36 bool = false
    fmt.Println(v32, v30, v31)
    fmt.Println(v17, v32, v30)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    if v36 {
        delete(v2, genFn0(836, 2, -1))
    } else {
        if v36 {
            fmt.Println(v30, v1)
            v37 := GenS1{F0: 0, F1: -1, F2: 65535}
            _ = v37
        } else {
            v30 += (((0.0 / (1e10 / (((float32((0.0 * (0.0 + ((v1 - (float64(0.1) * v1)) - float64(v30))))) + float32(((v1 * 0.0) * v1))) - 0.1) * v30))) - float32(3.14159)) + (0.1 * 1.5))
            var v38 int32 = -470
            var v39 float32 = 3.14159
            _ = v38
            _ = v39
        }
    }
    var v40 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v40
}
