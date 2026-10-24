package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 int, p2 int8) uint64 {
    p1 = (-(p1))
    var v4 uint16 = 533
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 uint8, p1 string, p2 int64) string {
    p2 += (p2 << ((970) & 63))
    var v5 float32 = 0.0
    var v6 int = -1
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7, v6)
        fmt.Println(v6, v6, v5)
        _ = fi7
    }
    v8 := GenS1{F0: -1, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v8
    return p1
}

func main() {
    var v1 bool = (("a" <= "hello") && false)
    v2 := make(map[string]uint32)
    var v3 float32 = (0.1 + 3.14159)
    v9 := []int{122, 32767, -9223372036854775808}
    if v1 {
        v3 += float32(0.0)
        if (282 > 42) {
            if (((((0 <= 42) || false) || v1) || v1) || v1) {
                v1 = (-32768 != 1000)
                if (false && false) {
                    for fi10 := 0; fi10 < 17; fi10++ {
                        switch fi10 {
                        case -2147483648:
                            for fi11 := 0; fi11 < 19; fi11++ {
                                var v12 float32 = v3
                                fmt.Println(v1)
                                v13 := GenS1{F0: 1000, F1: 2}
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        case 9223372036854775807:
                            var v14 uint64 = 812
                            _ = v14
                        case 127:
                            if v1 {
                                delete(v2, "foo bar")
                            } else {
                                var v15 int32 = 744
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    v1 = (float64(v3) != 3.14159)
                                    v1 = v1
                                    _ = fi16
                                }
                                for fi17 := 0; fi17 < 8; fi17++ {
                                    var v18 uint8 = 42
                                    _ = fi17
                                    _ = v18
                                }
                                _ = v15
                            }
                        default:
                            v1 = false
                        }
                        _ = fi10
                    }
                } else {
                    var v19 uint64 = genFn0(1, -745, -1)
                    _ = len(v2)
                    _ = v19
                }
                v20 := GenS2{F0: 42, F1: 2}
                _ = v20
            } else {
                var v21 int16 = 0
                _ = v21
            }
            v3 -= float32(0.1)
        }
    } else {
        var v22 string = "foo bar"
        v3 -= (float32(1e10) / float32(3.14159))
        _ = v22
    }
    switch -501 {
    case -266:
        v23 := GenS1{F0: -545, F1: 1}
        _ = v23
    default:
        v1 = true
    }
    var v24 float32 = v3
    if true {
        if ((!(!(!v1))) && v1) {
            if len(v9) > 0 {
                v9[0] = -9223372036854775808
            }
            switch "" {
            case "\n\t":
                switch 9223372036854775807 {
                case -9223372036854775808:
                    var v25 uint = 356
                    _ = v25
                default:
                    v3 = 0.1
                }
            case "foo bar":
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 uint64 = 891
                    _ = fi26
                    _ = v27
                }
            default:
                _ = len(v9)
            }
        }
    }
    v1 = (985 == 18446744073709551615)
    if v1 {
        fmt.Println(v24)
        if true {
            var v28 float64 = 3.14159
            v24 -= v3
            _ = v28
        } else {
            var v29 float32 = float32((float64((((v3 + ((1.5 + 3.14159) / 1.5)) * (0.0 + (v24 / float32((-2.25 / 0.1))))) - v24)) * 3.14159))
            _ = v29
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    var v30 string = ("hello" + ("" + (("hello" + "a") + "hello")))
    for fi31 := 0; fi31 < 9; fi31++ {
        v30 = ("\n\t" + v30)
        var v32 int64 = -311
        _ = fi31
        _ = v32
    }
    if (-1 <= 100) {
        v33 := GenS3{F0: 2}
        var v34 float64 = (((0.1 + 1e10) + 0.0) * 0.1)
        v3 -= v24
        _ = v33
        _ = v34
    } else {
        switch -1 {
        case 42:
            fmt.Println(v3, v1)
        }
    }
    v35 := GenS2{F0: 0, F1: 42}
    var v36 uint32 = 65535
    v24 -= float32(3.14159)
    var v37 bool = (!false)
    var v38 uint64 = 2
    var v39 uint8 = 42
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 uint8 = (v39 - 255)
        _ = fi40
        _ = v41
    }
    v36 = v36
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v9
}
