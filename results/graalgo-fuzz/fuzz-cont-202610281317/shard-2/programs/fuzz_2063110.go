package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 uint8; F2 int32 }
func genFn0(p0 uint8) int {
    p0 -= p0
    v5 := GenS1{F0: p0, F1: 0}
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 231
}
func genFn1() uint16 {
    v6 := []int{-39, -630, genFn0(255)}
    var v7 int32 = 42
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float64 = -2.25
        var v10 int32 = -2147483648
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int16 = 1000
    var v2 int64 = 0
    var v3 uint32 = 14
    var v4 uint = 574
    v11 := GenS2{F0: 1}
    v12 := []int{842}
    var v13 uint32 = v3
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int16 = v1
        _ = fi14
        _ = v15
    }
    if ((78 < 0) && true) {
        var v16 string = "hello"
        _ = v16
    }
    v17 := GenS3{F0: v3, F1: 291, F2: uint16(v4)}
    switch "hello" {
    case "":
        if ((float64(0.0) / (1.5 - (1e10 + 3.14159))) < 0.1) {
            switch -157 {
            case 0:
                if true {
                    v1 = (v1 + v1)
                    if false {
                        fmt.Println(v2, v13, v13)
                        fmt.Println(v13)
                        if len(v12) > 0 {
                            v12[0] = -2147483648
                        }
                    }
                    switch ((("" + ("" + "\n\t")) + "a") + "\n\t") {
                    case "foo bar":
                        v12 = append(v12, int(v4))
                    case "\n\t":
                        v18 := GenS4{F0: (((1e10 == 3.14159) || true) || ((!(true || (true && (((2147483647 == 1) && true) || true)))) && true)), F1: -1}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 bool = (-1 == 1000)
                            v21 := GenS5{F0: v4, F1: 283}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                }
            case -800:
                v2 -= (v2 ^ 42)
            default:
                v22 := GenS6{F0: 42}
                _ = v22
            }
            var v23 uint32 = v3
            if true {
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 uint64 = 144
                    v4 -= v4
                    if false {
                        v1 = v1
                        v13 = 1
                    } else {
                        for fi26 := 0; fi26 < 5; fi26++ {
                            var v27 uint16 = 32767
                            _ = fi26
                            _ = v27
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                if false {
                    if (42 != -14) {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        for fi28 := 0; fi28 < 1; fi28++ {
                            var v29 int32 = 0
                            for fi30 := 0; fi30 < 17; fi30++ {
                                _ = len(v12)
                                var v31 uint32 = v13
                                var v32 uint16 = 255
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            fmt.Println(v13, v29, fi28)
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        var v33 uint = v4
                        _ = v33
                    }
                } else {
                    var v34 bool = false
                    fmt.Println(v1, v3)
                    _ = v34
                }
            }
            _ = v23
        } else {
            var v35 float32 = 0.1
            fmt.Println(v13)
            _ = v35
        }
    case "foo bar":
        fmt.Println(v4)
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v36 string = ""
    if true {
        if false {
            v37 := GenS7{F0: int32(v1), F1: 42, F2: 1}
            _ = v37
        }
        switch 439 {
        case 131:
            var v38 int8 = 100
            _ = v38
        case 0:
            if len(v12) > 0 {
                v12[0] = genFn0(128)
            }
        case -1:
            for fi39 := 0; fi39 < 13; fi39++ {
                for fi40 := 0; fi40 < 16; fi40++ {
                    var v41 uint32 = 42
                    var v42 uint32 = v13
                    fmt.Println(v4)
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = fi39
            }
        default:
            fmt.Println(v3)
        }
        v4 = (v4 / ((v4) | 1))
    } else {
        var v43 int64 = v2
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
