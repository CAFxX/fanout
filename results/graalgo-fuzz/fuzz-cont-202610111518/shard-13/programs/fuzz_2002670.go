package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 string; F1 int64; F2 uint8 }
func genFn0(p0 uint8) uint8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: true, F1: 4294967295}
        p0 = 128
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: true, F1: 65535}
        _ = fi9
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    return (p0 | p0)
}

func main() {
    v1 := []uint32{165, 4294967295, 0, 800}
    v2 := make(map[uint32]uint32)
    var v3 float64 = 1e10
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    var v6 uint = 940
    v6 += (v6 % (((v6 << ((4294967295) & 63))) | 1))
    if true {
        fmt.Println(v6, v3)
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            v6 = (v6 + v6)
            _ = fi11
        }
        v3 = float64(((1e10 / 0.1) - float32(3.14159)))
        var v12 int32 = -2147483648
        _ = v12
    }
    v1 = append(v1, 2147483648)
    if (-311 >= -1) {
        v6 += (^(v6))
        var v13 uint = v6
        var v14 float64 = (0.1 * (v3 + float64(3.14159)))
        _ = v13
        _ = v14
    }
    fmt.Println(v3, v3)
    v15 := GenS1{F0: ((false || (-1 <= 1)) || false), F1: 42}
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            v3 = -2.25
            switch ("hello" + "") {
            case "hello":
                v6 = v6
            case "\n\t":
                if (425 <= 2) {
                    v18 := GenS1{F0: false, F1: 2}
                    _ = v18
                }
            case "":
                var v19 int32 = -1
                v6 -= (v6 | (^(v6)))
                if (v6 > (^(v6))) {
                    v3 = ((1e10 + (v3 * v3)) * v3)
                    v3 = v3
                } else {
                    var v20 float32 = 1e10
                    _ = v20
                }
                var v21 float64 = ((v3 + v3) * v3)
                _ = v19
                _ = v21
            }
            v3 -= 1.5
            _ = fi17
        }
        _ = fi16
    }
    v1 = append(v1, 502)
    v6 = v6
    switch "\n\t" {
    case "hello":
        var v22 int = 4294967295
        _ = v22
    case "":
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 string = ""
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    case "foo bar":
        switch ("\n\t" + ("hello" + "a")) {
        case "a":
            for fi26 := 0; fi26 < 17; fi26++ {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 bool = true
                    _ = fi27
                    _ = v28
                }
                v3 = (float64(0.1) / v3)
                for fi29 := 0; fi29 < 20; fi29++ {
                    var v30 int16 = 32767
                    var v31 uint = v6
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi26
            }
        }
    default:
        var v32 bool = (true && (32767 == 0))
        _ = v32
    }
    v6 -= v6
    var v33 bool = (!true)
    var v34 uint64 = 711
    v2[4294967295] = 0
    if (v6 >= (v6 - 26)) {
        var v35 int8 = 35
        v6 += v6
        _ = v35
    } else {
        switch "\n\t" {
        case "foo bar":
            v3 = float64(0.0)
        default:
            fmt.Println(v6, v34)
        }
        v36 := GenS2{F0: 0}
        for fi37 := 0; fi37 < 10; fi37++ {
            v3 = -2.25
            v38 := GenS3{F0: "hello", F1: 864}
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        v40 := GenS2{F0: 81}
        _ = fi39
        _ = v40
    }
    v41 := GenS4{F0: "", F1: int64(v34), F2: 128}
    switch 1 {
    case 32768:
        _ = len(v5)
    case 32767:
        var v42 uint = (^(v6))
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
