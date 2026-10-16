package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 uint8; F1 int32; F2 bool }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0(p0 uint32, p1 uint32, p2 uint32) int {
    var v5 int = -1
    var v6 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v5 ^ v5)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    v4 := make(map[int]uint32)
    if (!(1 >= -397)) {
        switch "a" {
        case "\n\t":
            var v7 int8 = 0
            _ = v7
        }
        v8 := GenS1{F0: 61}
        _ = v8
    } else {
        var v9 int16 = 42
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                for fi11 := 0; fi11 < 2; fi11++ {
                    if ("a" <= "\n\t") {
                        switch "a" {
                        case "\n\t":
                            var v12 bool = true
                            _ = v12
                        case "foo bar":
                            v1[(fi10 & fi11)] = ""
                        case "a":
                            v3 = append(v3, 42)
                        }
                        v13 := GenS2{F0: 42, F1: -86, F2: (v9 + v9)}
                        _ = v13
                    }
                    v14 := GenS3{F0: 240, F1: -1, F2: ((false && false) || true)}
                    var v15 string = "\n\t"
                    _ = fi11
                    _ = v14
                    _ = v15
                }
                var v16 int = 32768
                _ = fi10
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 float64 = float64((0.1 * 1.5))
            v19 := GenS3{F0: 4, F1: 2147483647, F2: (("hello" + "hello") >= "")}
            v20 := v4[(fi17 & 32767)]
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v9
    }
    v21 := GenS3{F0: 147, F1: -286, F2: false}
    switch "" {
    case "hello":
        v22 := GenS3{F0: 128, F1: -2147483648, F2: false}
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 int16 = -1
            var v25 int8 = -128
            _ = fi23
            _ = v24
            _ = v25
        }
    case "a":
        var v26 uint16 = 2
        _ = v26
    }
    v27 := GenS2{F0: -1, F1: 22, F2: 32767}
    v28 := GenS1{F0: 604}
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 float64 = float64(0.0)
        _ = fi29
        _ = v30
    }
    var v31 float32 = 1e10
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 float32 = (v31 - v31)
        switch "" {
        case "\n\t":
            v34 := GenS1{F0: 948}
            _ = v34
        case "a":
            if true {
                v35 := GenS2{F0: 0, F1: 127, F2: -1}
                _ = v35
            }
        default:
            v33 = v31
        }
        if false {
            var v36 int32 = 2147483647
            _ = v36
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v31)
    if false {
        for fi37 := 0; fi37 < 11; fi37++ {
            for fi38 := 0; fi38 < 7; fi38++ {
                v31 = 3.14159
                var v39 uint32 = 4294967295
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    }
    v31 = ((float32((1.5 * ((-2.25 * -2.25) * 1.5))) + (float32(1.5) / (v31 - -2.25))) * v31)
    v31 = (3.14159 + (float32(3.14159) - (v31 / (v31 * (float32((-2.25 - (float64(-2.25) - (1e10 + 3.14159)))) * v31)))))
    v31 = (v31 - float32(-2.25))
    var v40 uint8 = 1
    var v41 uint16 = 2
    for fi42 := 0; fi42 < 16; fi42++ {
        for fi43 := 0; fi43 < 20; fi43++ {
            v41 += v41
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = fi43
        }
        fmt.Println(v40)
        _ = fi42
    }
    switch 42 {
    case -310:
        v44 := GenS4{F0: ("hello" + ("hello" + (((("hello" + "a") + "hello") + "foo bar") + "a"))), F1: 2147483647}
        _ = v44
    case 128:
        v40 = 2
    case -2147483648:
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 uint32 = 2147483648
            _ = fi45
            _ = v46
        }
    default:
        var v47 int8 = 23
        _ = v47
    }
    var v48 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v48
}
