package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 bool; F2 int8 }
type GenS3 struct { F0 int64; F1 bool; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int64, p2 int16) int8 {
    v7 := []int{4294967295}
    fmt.Println(p2)
    switch 32767 {
    case 127:
        var v8 uint = 42
        _ = v8
    case -2147483648:
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(p0, p2, p1)
            _ = fi9
        }
    case 33:
        switch 128 {
        case -1:
            p1 -= 9223372036854775807
        case 0:
            p1 += p1
        case 891:
            var v10 bool = (!true)
            _ = v10
        default:
            p1 -= p1
        }
    default:
        fmt.Println(p0)
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if ((1e10 != (1e10 * 1.5)) || false) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            fmt.Println(fi11)
            v12 := GenS1{F0: p2}
            _ = v12
        }
        v13 := GenS2{F0: (fi11 + fi11), F1: (p0 <= 127), F2: 93}
        _ = fi11
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint64, p2 int8) string {
    if false {
        fmt.Println(p1, p0)
    }
    v14 := GenS2{F0: 4294967295, F1: false, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return ("" + ("hello" + ""))
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 103
    var v3 bool = false
    v4 := []uint32{113, 42}
    v5 := []uint32{334, 459, 986, 1}
    v6 := make(map[uint32]uint32)
    if (v3 && ((-2.25 + 0.0) != ((0.1 + 3.14159) * (float64(1.5) * float64((float32(-2.25) / -2.25)))))) {
        v3 = false
        v2 = (v2 * 0)
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS3{F0: -784, F1: false, F2: 422}
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint = 1
            v3 = (!v3)
            fmt.Println(v3, v3)
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    if v3 {
        fmt.Println(v2)
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS4{F0: 1}
            for fi21 := 0; fi21 < 7; fi21++ {
                v2 = 1
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        if len(v4) > 0 {
            v4[0] = 611
        }
    }
    switch "foo bar" {
    case "":
        if v3 {
            var v22 bool = (v3 && (!v3))
            var v23 bool = false
            _ = v22
            _ = v23
        } else {
            v24 := GenS3{F0: 42, F1: v3, F2: 32767}
            fmt.Println(v3, v2)
            _ = v24
        }
    default:
        v25 := GenS5{F0: 0}
        _ = v25
    }
    var v26 float32 = 1e10
    var v27 int32 = -150
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := v1["hello"]
        for fi30 := 0; fi30 < 15; fi30++ {
            if (32767 > 498) {
                v31 := GenS6{F0: 2}
                _ = len(v6)
                v26 = v26
                _ = v31
            } else {
                _ = len(v4)
            }
            _ = fi30
        }
        var v32 float32 = v26
        _ = fi28
        _ = v29
        _ = v32
    }
    var v33 uint = 1
    var v34 string = "foo bar"
    fmt.Println(v3, v34, v34)
    var v35 string = v34
    switch 0 {
    case 127:
        v27 += (v27 << ((597) & 63))
    case 0:
        var v36 int64 = 0
        _ = v36
    case 4294967295:
        v2 = 255
    default:
        delete(v1, "")
    }
    _ = len(v6)
    var v37 uint16 = 2
    v38 := v1[(v35 + ("hello" + "a"))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
