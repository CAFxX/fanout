package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint64, p2 int) string {
    fmt.Println(p2)
    var v6 bool = (1 < 0)
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := []int{-56, fi7, fi7, (fi8 & -442)}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{127, p2}
    for fi11 := 0; fi11 < 1; fi11++ {
        p1 += (p1 | 18446744073709551615)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return "a"
}
func genFn1(p0 int64, p1 uint32, p2 bool) bool {
    p2 = ((((0.1 + 0.0) - float64(float32(0.1))) + (3.14159 - (1.5 + -2.25))) > 0.1)
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    return (p2 && p2)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]uint32)
    var v5 int32 = -2147483648
    var v12 int8 = 1
    v4[(v3 ^ (v3 / ((763) | 1)))] = (^(v3))
    var v13 int8 = (v12 - -76)
    var v14 float64 = 3.14159
    if true {
        if true {
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 19; fi16++ {
                    v17 := GenS1{F0: (-(v13))}
                    if true {
                        v18 := GenS2{F0: 2147483647}
                        v19 := []int{}
                        _ = v18
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    fmt.Println(fi20, v5, v13)
                    v21 := GenS2{F0: v3}
                    var v22 uint32 = (v3 & (v3 & v3))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 uint64 = 4294967295
                _ = fi15
                _ = v23
            }
            v24 := GenS1{F0: (v13 * v12)}
            v25 := GenS3{F0: 32767}
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 uint16 = 65535
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 17; fi29++ {
            v14 += (float64(((1.5 / 3.14159) * (3.14159 / 1.0))) * -2.25)
            fmt.Println(fi29, v5)
            _ = fi29
        }
        _ = fi28
    }
    if (false || ((false || (!((v13 - v13) != v13))) && true)) {
        for fi30 := 0; fi30 < 20; fi30++ {
            switch 42 {
            case 488:
                v1[2147483647] = (v3 | v3)
            case 0:
                var v31 int8 = 121
                _ = v31
            default:
                v32 := []int{-1, (^(fi30)), 128}
                _ = v32
            }
            fmt.Println(v14)
            _ = fi30
        }
        if false {
            delete(v1, (^(v3)))
            var v33 uint = 255
            _ = len(v4)
            _ = v33
        }
    } else {
        switch 872 {
        case 2147483647:
            for fi34 := 0; fi34 < 4; fi34++ {
                v13 += -1
                fmt.Println(v14, v5, v3)
                _ = fi34
            }
        default:
            var v35 uint8 = 42
            _ = v35
        }
    }
    v5 -= v5
    v36 := GenS1{F0: v13}
    v37 := GenS1{F0: v13}
    switch "" {
    case "foo bar":
        for fi38 := 0; fi38 < 20; fi38++ {
            if false {
                if (!(true || true)) {
                    if (v3 != v3) {
                        for fi39 := 0; fi39 < 5; fi39++ {
                            v3 -= v3
                            v40 := GenS2{F0: 923}
                            _ = fi39
                            _ = v40
                        }
                        var v41 int64 = 42
                        for fi42 := 0; fi42 < 2; fi42++ {
                            var v43 int64 = 9223372036854775807
                            fmt.Println(fi42)
                            _ = fi42
                            _ = v43
                        }
                        _ = v41
                    }
                    fmt.Println(v13, v14, v13)
                }
                v2[v3] = true
                v44 := []int{-900, fi38, fi38}
                _ = v44
            } else {
                if (((0.1 + 1.5) >= float32((0.0 / ((1e10 / (v14 * v14)) * v14)))) && (3.14159 <= 3.14159)) {
                    for fi45 := 0; fi45 < 8; fi45++ {
                        if (1 != 275) {
                            fmt.Println(v5, fi45)
                            var v46 float32 = 1.5
                            _ = v46
                        }
                        _ = fi45
                    }
                }
            }
            fmt.Println(v5)
            _ = fi38
        }
    case "hello":
        var v47 int16 = 0
        _ = v47
    default:
        v48 := GenS2{F0: (^(v3))}
        _ = v48
    }
    var v49 float32 = float32(v14)
    var v50 float64 = v14
    var v51 int8 = (^(v13))
    v50 = (1e10 / v50)
    var v52 uint8 = 53
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
