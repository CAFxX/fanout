package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint8 }
func genFn0() bool {
    var v4 float64 = ((1.5 / float64(1e10)) - 3.14159)
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 446, F1: true, F2: 65535}
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{fi5, -1, 32768}
            _ = fi7
            _ = v8
        }
        if (-216 >= 32767) {
            var v9 string = "hello"
            var v10 float32 = -2.25
            if (0 < 2) {
                fmt.Println(v10, v10)
                var v11 float32 = float32((v4 * (float64((v10 + (0.1 * (v10 + (((v10 * v10) + float32(((1.5 + ((((((-2.25 * (float64(((v10 + 1.5) / v10)) * v4)) / 1.5) + v4) / v4) + v4) + float64(0.0))) / 1.0))) * 0.0))))) - v4)))
                v12 := GenS2{F0: (436 == 65535), F1: -99}
                _ = v11
                _ = v12
            }
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v13 := GenS2{F0: true, F1: 100}
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: 65535, F1: ((!(!true)) && (720 <= 2)), F2: 26}
        v16 := GenS3{F0: -1, F1: 251}
        v17 := GenS4{F0: (fi14 / ((32768) | 1)), F1: -1}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint64 = 244
        var v20 int8 = 42
        if false {
            v4 = (((-2.25 - v4) - v4) + (v4 * v4))
            var v21 int64 = 288
            _ = v21
        } else {
            var v22 float32 = 1e10
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = v13
    _ = v4
    return false
}
func genFn1(p0 int16, p1 bool) uint32 {
    var v23 float64 = 1.5
    if false {
        var v24 uint32 = 4294967295
        var v25 bool = false
        v26 := GenS5{F0: 42}
        _ = v24
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v23
    return 618
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 100
    var v3 int = 2147483647
    v3 += v3
    v27 := []int{}
    fmt.Println(v2)
    v28 := GenS4{F0: 4294967295, F1: (^(v2))}
    switch "a" {
    case "a":
        v2 = -77
    default:
        v3 -= (v3 << ((42) & 63))
    }
    fmt.Println(v3)
    if (!false) {
        switch "a" {
        case "\n\t":
            v29 := GenS3{F0: 850, F1: 1000}
            _ = v29
        case "a":
            v30 := GenS1{F0: 1, F1: (1 == 2), F2: genFn1(1000, (42 > 330))}
            _ = v30
        default:
            v31 := GenS3{F0: 1, F1: 32767}
            _ = v31
        }
    } else {
        var v32 string = "\n\t"
        _ = v32
    }
    switch "hello" {
    case "foo bar":
        v2 = v2
    }
    var v33 string = ("foo bar" + "")
    if len(v27) > 0 {
        v27[0] = -9223372036854775808
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        for fi35 := 0; fi35 < 19; fi35++ {
            fmt.Println(v3)
            if true {
                v3 = -2147483648
            }
            var v36 string = "\n\t"
            _ = fi35
            _ = v36
        }
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
        var v37 bool = true
        _ = fi34
        _ = v37
    }
    switch v33 {
    case "hello":
        v38 := GenS1{F0: 65535, F1: (!(v2 >= v2)), F2: 735}
        _ = v38
    case "a":
        if (-1 <= 0) {
            var v39 string = ("hello" + v33)
            _ = v39
        } else {
            v40 := GenS6{F0: 199, F1: 469, F2: 1}
            _ = v40
        }
    }
    var v41 uint64 = 42
    fmt.Println(v41)
    v3 = (v3 % (((v3 + v3)) | 1))
    var v42 int16 = 1
    fmt.Println(v42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v41
    _ = v42
}
