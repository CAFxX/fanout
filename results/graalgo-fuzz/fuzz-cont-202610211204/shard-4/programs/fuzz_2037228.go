package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 string; F1 uint32 }
func genFn0(p0 int16, p1 uint) uint64 {
    p0 -= p0
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint32 = 621
        v9 := GenS1{F0: -779}
        switch "a" {
        case "":
            if false {
                var v10 int64 = 438
                var v11 uint64 = 0
                _ = v10
                _ = v11
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v12 float32 = 1e10
    v13 := GenS1{F0: -104}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{327, 127, 141}
    var v3 int = 32768
    v4 := []int{(v3 - v3), v3}
    v5 := make(map[int]bool)
    var v6 float32 = float32(3.14159)
    _ = len(v4)
    fmt.Println(v3)
    v6 = -2.25
    var v14 uint8 = 178
    v15 := GenS1{F0: -394}
    if len(v2) > 0 {
        v2[0] = (v3 - 247)
    }
    var v16 int16 = -32768
    v5[-425] = true
    v17 := GenS2{F0: 255, F1: 614, F2: v14}
    v18 := GenS3{F0: 127, F1: "", F2: (v14 % ((v14) | 1))}
    v6 = v6
    var v19 uint64 = 2
    for fi20 := 0; fi20 < 3; fi20++ {
        _ = len(v1)
        delete(v1, (("foo bar" + "foo bar") + (("\n\t" + "foo bar") + "foo bar")))
        var v21 float64 = (float64(v6) / 1.0)
        _ = fi20
        _ = v21
    }
    if (127 <= int8(v19)) {
        var v22 uint64 = genFn0((v16 & v16), 258)
        if (3.14159 >= v6) {
            var v23 uint16 = 32767
            switch v3 {
            case -235:
                var v24 int = (v3 + v3)
                _ = v24
            case 128:
                v25 := GenS4{F0: v14, F1: v14, F2: 2}
                _ = v25
            case 980:
                v26 := GenS5{F0: ("\n\t" + "a"), F1: 42}
                _ = v26
            }
            switch v3 {
            case 128:
                _ = len(v1)
            case -2147483648:
                v19 += 4294967295
            default:
                delete(v1, "foo bar")
            }
            _ = v23
        }
        _ = v22
    } else {
        for fi27 := 0; fi27 < 18; fi27++ {
            for fi28 := 0; fi28 < 3; fi28++ {
                fmt.Println(v3, v16)
                _ = fi28
            }
            v16 += v16
            _ = fi27
        }
        fmt.Println(v3, v19, v16)
        fmt.Println(v19, v3, v14)
    }
    var v29 float32 = ((v6 + (v6 - ((v6 / 1e10) + float32(-2.25)))) / 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
