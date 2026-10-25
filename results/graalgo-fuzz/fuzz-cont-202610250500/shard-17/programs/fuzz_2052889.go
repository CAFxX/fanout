package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint8) uint8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int16 = -32768
            var v10 uint32 = 2147483647
            fmt.Println(v9, fi7, fi7)
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{}
        _ = len(v11)
        _ = fi7
        _ = v11
    }
    p0 -= (p0 | 4294967295)
    for fi12 := 0; fi12 < 13; fi12++ {
        p0 += p0
        v13 := []int{883, 2147483647}
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 42}
    v15 := []int{0}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return (p1 * (p1 + p1))
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 1
    v3 := make(map[int]uint32)
    var v4 uint32 = v1
    var v5 uint8 = 121
    var v6 uint = 42
    fmt.Println(v1, v2)
    v16 := []int{370, 4294967295, 4294967295, 782}
    v17 := GenS1{F0: 601}
    v5 -= v5
    switch "a" {
    case "":
        var v18 float64 = ((((float64(0.0) * (3.14159 + (float64(1.5) + (1e10 * 3.14159)))) * 1.5) / 0.1) + 3.14159)
        _ = v18
    }
    if (!true) {
        v2 = (v2 & 1)
        var v19 int8 = 127
        v3[-351] = v1
        _ = v19
    }
    if false {
        v20 := GenS2{F0: -1, F1: v4}
        _ = v20
    } else {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        switch "a" {
        case "a":
            if (!true) {
                var v21 int32 = 2147483647
                _ = v21
            }
        case "":
            fmt.Println(v5)
        default:
            v1 = v4
        }
    }
    var v22 float32 = (1.5 / (3.14159 - 1.5))
    v22 = v22
    var v23 int16 = 1000
    v23 = v23
    switch "\n\t" {
    case "foo bar":
        var v24 int32 = -1
        _ = v24
    default:
        for fi25 := 0; fi25 < 14; fi25++ {
            if len(v16) > 0 {
                v16[0] = 698
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS3{F0: (v6 / ((uint(v1)) | 1)), F1: v23}
                v28 := GenS4{F0: 66}
                _ = fi26
                _ = v27
                _ = v28
            }
            v29 := GenS1{F0: (v23 / ((v23) | 1))}
            _ = fi25
            _ = v29
        }
    }
    v23 -= (v23 - v23)
    for fi30 := 0; fi30 < 14; fi30++ {
        if (genFn0(758, v5) >= 0) {
            v5 = 255
            v31 := GenS5{F0: 0}
            delete(v3, (fi30 % (((fi30 / ((597) | 1))) | 1)))
            _ = v31
        } else {
            var v32 float64 = -2.25
            _ = v32
        }
        _ = fi30
    }
    var v33 bool = true
    var v34 float32 = -2.25
    for fi35 := 0; fi35 < 6; fi35++ {
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 int = (fi36 * int(v23))
            _ = fi36
            _ = v37
        }
        fmt.Println(v6, v4)
        v5 = (v5 - uint8(v6))
        _ = fi35
    }
    if v33 {
        v38 := GenS4{F0: 127}
        _ = len(v16)
        v34 = (((v34 - v22) * (1.5 - (-2.25 / (v22 / v22)))) + ((1e10 - v34) - v34))
        _ = v38
    } else {
        if true {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            switch 42 {
            case -9223372036854775808:
                v39 := v3[int(v5)]
                _ = v39
            }
        } else {
            var v40 int8 = 0
            _ = v40
        }
        v33 = ("a" == "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
