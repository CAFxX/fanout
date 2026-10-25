package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int8; F1 uint16; F2 int }
type GenS5 struct { F0 int8; F1 int16 }
func genFn0(p0 int16) int64 {
    if false {
        if false {
            if false {
                v5 := GenS1{F0: 576, F1: 65535, F2: 17}
                _ = v5
            }
            var v6 int = 545
            _ = v6
        }
        p0 -= 185
        p0 = p0
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            p0 = 32767
            _ = fi7
        }
    }
    var v8 int = 318
    p0 = (^(p0))
    fmt.Println(p0)
    _ = p0
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("a" + ("hello" + "foo bar")), ("\n\t" + "a"), "\n\t", ""}
    v3 := make(map[string]int)
    var v4 float64 = (1e10 * float64(1e10))
    fmt.Println(v4)
    if ("a" == "hello") {
        switch 0 {
        case 513:
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(fi9)
                _ = fi9
            }
        default:
            v10 := v3["hello"]
            _ = v10
        }
        v11 := GenS1{F0: 1, F1: 378, F2: 0}
        v4 -= 1.5
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        if (!(4294967295 == 1)) {
            var v13 float64 = -2.25
            switch 32768 {
            case 127:
                v13 -= (float64(1e10) * (3.14159 + (1.5 * float64(((-2.25 + (1e10 * float32((0.0 + v13)))) / -2.25)))))
            case 1:
                for fi14 := 0; fi14 < 18; fi14++ {
                    v15 := GenS2{F0: 42}
                    _ = fi14
                    _ = v15
                }
            case -1:
                v13 -= float64(3.14159)
            }
            v16 := GenS1{F0: -26, F1: 182, F2: 0}
            _ = v13
            _ = v16
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            v4 -= 1e10
            if false {
                var v18 int16 = 42
                _ = v18
            }
            _ = fi17
        }
        _ = fi12
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        delete(v3, "")
        v4 = float64(3.14159)
        fmt.Println(v4)
        _ = fi19
    }
    v20 := GenS1{F0: -881, F1: 255, F2: -1}
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := GenS3{F0: "\n\t", F1: genFn0(42)}
        var v23 uint64 = 314
        var v24 int8 = -128
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    if (632 >= 32767) {
        var v25 uint32 = 392
        for fi26 := 0; fi26 < 11; fi26++ {
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 int64 = int64(fi26)
                var v29 uint8 = 1
                v30 := GenS4{F0: 1, F1: 2, F2: (fi27 | (^(fi27)))}
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi26
        }
        _ = v25
    }
    v4 -= float64((1e10 - 1.5))
    if true {
        v2 = append(v2, "hello")
        if len(v2) > 0 {
            v2[0] = ("hello" + "\n\t")
        }
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    }
    var v31 uint32 = 2147483647
    for fi32 := 0; fi32 < 2; fi32++ {
        v4 = (-2.25 / ((v4 / (v4 / v4)) / v4))
        v33 := GenS5{F0: 100, F1: 1000}
        _ = fi32
        _ = v33
    }
    if (453 < 42) {
        v4 = (v4 + (0.1 + (float64(1e10) + float64(0.0))))
        v4 = (float64((0.0 / 1.0)) - (float64(3.14159) + v4))
        var v34 bool = (42 == 0)
        _ = v34
    } else {
        var v35 uint64 = 42
        _ = v35
    }
    _ = len(v2)
    v36 := GenS5{F0: 80, F1: -831}
    for fi37 := 0; fi37 < 16; fi37++ {
        _ = len(v3)
        _ = fi37
    }
    var v38 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v36
    _ = v38
    _ = v4
}
