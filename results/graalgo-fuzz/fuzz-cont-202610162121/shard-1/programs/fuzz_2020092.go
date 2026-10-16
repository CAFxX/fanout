package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint32) int16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        switch "foo bar" {
        case "":
            v8 := GenS1{F0: 0}
            _ = v8
        case "hello":
            var v9 float64 = 1.5
            _ = v9
        case "\n\t":
            var v10 uint16 = 746
            _ = v10
        }
        fmt.Println(fi7, p0)
        v11 := GenS1{F0: 1}
        _ = fi7
        _ = v11
    }
    v12 := GenS2{F0: 1, F1: 237, F2: -795}
    _ = p0
    _ = v12
    return -32768
}

func main() {
    var v1 uint64 = 42
    var v2 uint64 = v1
    v3 := make(map[string]bool)
    var v4 uint8 = 255
    var v5 int32 = -1
    var v6 int16 = 516
    v13 := GenS3{F0: 42}
    for fi14 := 0; fi14 < 9; fi14++ {
        if (true || ((-(fi14)) == fi14)) {
            var v15 uint = 1
            fmt.Println(v6, v6)
            _ = v15
        } else {
            fmt.Println(fi14, v1)
            switch 160 {
            case 1:
                var v16 int8 = -1
                _ = v16
            case 0:
                fmt.Println(v6, v4)
            }
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            v5 = (v5 % ((268) | 1))
            _ = fi17
        }
        v18 := GenS1{F0: -50}
        _ = fi14
        _ = v18
    }
    if (!(false && false)) {
        _ = len(v3)
        if (float64(0.0) <= (1.5 + 0.1)) {
            if true {
                v1 = v1
                for fi19 := 0; fi19 < 12; fi19++ {
                    if (false || ((!true) || (100 <= 1))) {
                        var v20 int64 = 1
                        v3["foo bar"] = true
                        _ = v20
                    }
                    _ = len(v3)
                    v21 := GenS4{F0: 0}
                    _ = fi19
                    _ = v21
                }
                var v22 uint8 = (^(v4))
                _ = v22
            } else {
                var v23 uint16 = 1
                _ = v23
            }
        } else {
            var v24 int = -2147483648
            var v25 uint64 = (^(v1))
            if false {
                v1 = v2
            }
            _ = v24
            _ = v25
        }
    }
    switch 42 {
    case 4294967295:
        v26 := GenS3{F0: 42}
        _ = v26
    case -1:
        v4 += (v4 % ((101) | 1))
    default:
        v27 := GenS2{F0: (-(v5)), F1: (v1 << ((0) & 63)), F2: int16(v5)}
        _ = v27
    }
    var v28 bool = false
    var v29 int = 0
    var v30 string = "hello"
    v3[("hello" + v30)] = (v4 <= v4)
    fmt.Println(v1, v1, v5)
    v31 := []int{0, (-(v29)), v29, 128}
    v5 = v5
    if true {
        v32 := GenS4{F0: 4294967295}
        switch v29 {
        case 501:
            if (!false) {
                v1 -= (v2 - (v1 << ((uint(v1)) & 63)))
                v30 = v30
            }
        }
        _ = v32
    }
    v33 := GenS2{F0: 266, F1: (v2 / ((763) | 1)), F2: v6}
    _ = len(v3)
    fmt.Println(v4)
    v6 += (v6 / ((1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
