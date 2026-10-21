package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint64; F2 int32 }
type GenS4 struct { F0 int16; F1 int8; F2 int32 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint) bool {
    p0 = (p0 & 140)
    if (!(!(!(255 == 32767)))) {
        for fi4 := 0; fi4 < 20; fi4++ {
            for fi5 := 0; fi5 < 9; fi5++ {
                var v6 uint32 = 242
                p0 = (p0 ^ (^(p0)))
                _ = fi5
                _ = v6
            }
            var v7 float64 = 1.5
            var v8 int32 = -2147483648
            _ = fi4
            _ = v7
            _ = v8
        }
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(p0, fi9)
            _ = fi9
        }
    }
    _ = p0
    return false
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    var v3 float64 = 0.1
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            delete(v2, 352)
            v3 -= v3
            if false {
                v3 -= (v3 * v3)
                var v12 float32 = (0.1 * (1e10 + 0.1))
                var v13 bool = (v3 <= v3)
                switch ("hello" + "foo bar") {
                case "":
                    v13 = (!false)
                    v14 := GenS1{F0: 0}
                    _ = v14
                case "hello":
                    v15 := GenS2{F0: 42}
                    _ = v15
                default:
                    v3 = 3.14159
                }
                v16 := GenS3{F0: 42, F1: 42, F2: 42}
                _ = v12
                _ = v13
                _ = v16
            }
            _ = fi11
        }
        _ = fi10
    }
    var v17 string = "foo bar"
    v18 := GenS2{F0: 642}
    var v19 int8 = -128
    v3 += (float64((float32((float64((3.14159 - 1e10)) + 1e10)) * 0.0)) + v3)
    v20 := GenS4{F0: -627, F1: (^(v19)), F2: 2147483647}
    switch "a" {
    case "":
        v19 = (v19 & v19)
    case "foo bar":
        for fi21 := 0; fi21 < 5; fi21++ {
            v19 = v19
            fmt.Println(v3, fi21)
            var v22 uint = 386
            _ = fi21
            _ = v22
        }
    }
    v3 = 1e10
    fmt.Println(v17, v17)
    v23 := GenS3{F0: 18446744073709551615, F1: 42, F2: 0}
    var v24 uint8 = 48
    if true {
        switch (v17 + v17) {
        case "\n\t":
            v17 = v17
        case "foo bar":
            if false {
                if true {
                    switch v17 {
                    case "":
                        delete(v2, 4294967295)
                    default:
                        v25 := GenS2{F0: 2147483647}
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 3; fi26++ {
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v19 += 100
                            v19 += (v19 / ((v19) | 1))
                            v28 := GenS1{F0: 746}
                            _ = fi27
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 15; fi29++ {
                            var v30 uint8 = 255
                            v3 += ((v3 * (v3 + (3.14159 / v3))) / 3.14159)
                            _ = fi29
                            _ = v30
                        }
                        v31 := GenS5{F0: "", F1: v24}
                        _ = fi26
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 14; fi32++ {
                        if (false || (false && (!false))) {
                            var v33 uint16 = 0
                            var v34 int8 = 127
                            _ = v33
                            _ = v34
                        } else {
                            var v35 int64 = -9223372036854775808
                            _ = v35
                        }
                        _ = fi32
                    }
                }
                if false {
                    var v36 string = v17
                    v19 = (v19 / (((^(v19))) | 1))
                    v37 := GenS6{F0: 2}
                    _ = v36
                    _ = v37
                }
                v17 = ((((v17 + v17) + (v17 + v17)) + ("a" + (v17 + ""))) + v17)
            } else {
                v19 = v19
                if false {
                    switch -1 {
                    case 929:
                        var v38 uint64 = 2
                        _ = v38
                    case 128:
                        v39 := GenS1{F0: 408}
                        _ = v39
                    }
                } else {
                    if true {
                        v40 := GenS7{F0: 401}
                        var v41 int64 = 0
                        _ = v40
                        _ = v41
                    }
                }
            }
        case "a":
            var v42 int16 = -213
            _ = v42
        default:
            v43 := GenS2{F0: 68}
            _ = v43
        }
        v19 = (v19 & 42)
    }
    switch 0 {
    case -955:
        if (v17 == "a") {
            v44 := GenS3{F0: 4294967295, F1: 2, F2: -49}
            _ = v44
        }
    default:
        var v45 int = 0
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
}
