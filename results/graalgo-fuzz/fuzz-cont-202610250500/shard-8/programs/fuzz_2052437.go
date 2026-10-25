package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint16) uint64 {
    var v6 float32 = 1.5
    if true {
        switch -320 {
        case 727:
            fmt.Println(v6)
        default:
            v7 := []int{128, 723}
            _ = v7
        }
    }
    var v8 uint8 = 56
    v9 := GenS1{F0: -2147483648, F1: (p0 * p0)}
    _ = p0
    _ = v6
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int8) int {
    if ((true && false) && ((false || true) && (((((3.14159 / float32(float64((3.14159 + float32(0.1))))) - float32(-2.25)) / 1.5) / -2.25) == (3.14159 * 3.14159)))) {
        v10 := []int{-29, -3}
        switch 344 {
        case 4294967295:
            var v11 uint64 = 2
            _ = v11
        case 32768:
            var v12 uint64 = 654
            _ = v12
        default:
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: -1, F1: 466}
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 uint = 255
                    p0 -= -1
                    var v17 string = "\n\t"
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi13
                _ = v14
            }
        }
        _ = v10
    } else {
        var v18 uint8 = 103
        switch "\n\t" {
        case "hello":
            fmt.Println(v18, v18)
        case "\n\t":
            v19 := GenS2{F0: false}
            _ = v19
        }
        _ = v18
    }
    v20 := []int{821}
    var v21 int32 = -2147483648
    _ = p0
    _ = v20
    _ = v21
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{492}
    var v3 bool = (255 > 148)
    var v4 uint8 = 127
    v5 := make(map[string]string)
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := v5[("\n\t" + "a")]
        v3 = ((v3 || v3) && ((v3 || v3) && true))
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := GenS1{F0: 1, F1: 65535}
        fmt.Println(v4)
        _ = fi24
        _ = v25
    }
    v26 := v5["a"]
    switch v26 {
    case "\n\t":
        var v27 uint16 = 42
        _ = v27
    default:
        fmt.Println(v4)
    }
    if false {
        delete(v1, 42)
    } else {
        if v3 {
            fmt.Println(v26, v26, v26)
        }
        var v28 float32 = 3.14159
        fmt.Println(v28)
        _ = v28
    }
    v26 = v26
    fmt.Println(v3)
    if ((2147483647 == -1) && (953 <= 42)) {
        v1[-1] = 32768
        for fi29 := 0; fi29 < 17; fi29++ {
            v3 = (float32(-2.25) > 0.0)
            if true {
                v3 = false
                var v30 int16 = -1
                for fi31 := 0; fi31 < 18; fi31++ {
                    if v3 {
                        v4 = 255
                    } else {
                        v32 := GenS3{F0: -128, F1: 883}
                        switch "a" {
                        case "a":
                            fmt.Println(v4, v30)
                        }
                        var v33 uint32 = 2147483647
                        _ = v32
                        _ = v33
                    }
                    _ = fi31
                }
                _ = v30
            } else {
                var v34 uint16 = 434
                _ = v34
            }
            v35 := GenS3{F0: 100, F1: 430}
            _ = fi29
            _ = v35
        }
        if v3 {
            _ = len(v2)
            var v36 bool = v3
            _ = v36
        } else {
            var v37 int32 = -2147483648
            v4 = 127
            var v38 float32 = (float32(1.5) / 1.0)
            _ = v37
            _ = v38
        }
    } else {
        if (!v3) {
            v26 = "a"
            for fi39 := 0; fi39 < 12; fi39++ {
                var v40 string = v26
                _ = fi39
                _ = v40
            }
            for fi41 := 0; fi41 < 5; fi41++ {
                switch v26 {
                case "\n\t":
                    for fi42 := 0; fi42 < 6; fi42++ {
                        var v43 uint64 = 440
                        fmt.Println(fi41)
                        _ = fi42
                        _ = v43
                    }
                case "foo bar":
                    var v44 float64 = 1e10
                    _ = v44
                default:
                    for fi45 := 0; fi45 < 13; fi45++ {
                        var v46 float64 = -2.25
                        v46 = (((1.5 / v46) * 1e10) * (v46 * ((((v46 / -2.25) - v46) + (v46 * 1e10)) / 1.0)))
                        _ = fi45
                        _ = v46
                    }
                }
                _ = fi41
            }
        }
    }
    _ = len(v5)
    for fi47 := 0; fi47 < 19; fi47++ {
        var v48 uint = 65535
        v49 := GenS1{F0: int32(v48), F1: 2}
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println(v26, v4)
    v50 := GenS1{F0: -1, F1: 255}
    var v51 int = 127
    fmt.Println(v4, v4, v4)
    fmt.Println(v3, v3)
    for fi52 := 0; fi52 < 4; fi52++ {
        v53 := GenS3{F0: -128, F1: 0}
        delete(v5, v26)
        var v54 string = v26
        _ = fi52
        _ = v53
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
