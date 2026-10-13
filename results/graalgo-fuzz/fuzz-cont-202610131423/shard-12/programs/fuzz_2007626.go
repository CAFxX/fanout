package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0(p0 int64, p1 int) uint64 {
    fmt.Println(p1, p1)
    var v4 float32 = -2.25
    if true {
        switch "a" {
        case "a":
            switch "foo bar" {
            case "hello":
                if false {
                    v5 := GenS1{F0: 1000, F1: 557, F2: 0}
                    _ = v5
                } else {
                    var v6 uint64 = 42
                    p1 = 120
                    _ = v6
                }
            case "":
                switch p1 {
                case 128:
                    switch 127 {
                    case 32767:
                        for fi7 := 0; fi7 < 2; fi7++ {
                            v8 := GenS1{F0: -1, F1: 249, F2: 2}
                            _ = fi7
                            _ = v8
                        }
                    case 502:
                        if (!((p0 & -9223372036854775808) >= p0)) {
                            v9 := []int{p1}
                            if true {
                                if (-128 != 1) {
                                    var v10 float64 = 0.0
                                    fmt.Println(p0, p1)
                                    _ = v10
                                }
                                if false {
                                    v11 := GenS2{F0: int(p0)}
                                    _ = v11
                                }
                            } else {
                                var v12 uint32 = 42
                                var v13 uint8 = 1
                                _ = v12
                                _ = v13
                            }
                            _ = v9
                        } else {
                            v14 := []int{p1, p1}
                            var v15 uint32 = 777
                            _ = v14
                            _ = v15
                        }
                    default:
                        fmt.Println(p0, p0)
                    }
                case 42:
                    p1 = p1
                case -598:
                    for fi16 := 0; fi16 < 6; fi16++ {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            v18 := GenS1{F0: -32768, F1: 774, F2: 255}
                            var v19 string = (("a" + "") + ("" + (("\n\t" + "foo bar") + "hello")))
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        var v20 uint32 = 42
                        v21 := GenS3{F0: 45, F1: 0}
                        _ = fi16
                        _ = v20
                        _ = v21
                    }
                default:
                    var v22 string = "foo bar"
                    _ = v22
                }
            case "\n\t":
                fmt.Println(p0, p1)
            default:
                v23 := GenS3{F0: -128, F1: -1}
                _ = v23
            }
        }
        p1 = 42
        v4 += float32(((3.14159 - ((((-2.25 * (float64(((1.5 - -2.25) * v4)) * ((float64(v4) - 3.14159) - 1e10))) / float64((v4 - v4))) * 0.0) * 1e10)) - 0.0))
    }
    if false {
        var v24 uint32 = 2147483647
        fmt.Println(v4)
        fmt.Println(p0, v4)
        _ = v24
    } else {
        if ("" == "hello") {
            var v25 string = ""
            if ((1 == 42) && (2147483647 == 42)) {
                switch (^(p1)) {
                case -1:
                    for fi26 := 0; fi26 < 19; fi26++ {
                        switch (fi26 / ((0) | 1)) {
                        case 32767:
                            var v27 int32 = -2147483648
                            _ = v27
                        case 2147483647:
                            v28 := []int{fi26}
                            _ = v28
                        default:
                            v4 -= ((0.0 - v4) / v4)
                        }
                        _ = fi26
                    }
                }
            }
            _ = v25
        } else {
            var v29 int8 = 40
            v4 = 0.1
            _ = v29
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 int = -1
    v2 := []string{("a" + "a")}
    v3 := make(map[string]string)
    var v30 string = ""
    _ = len(v3)
    v3[""] = (v30 + ("hello" + v30))
    for fi31 := 0; fi31 < 2; fi31++ {
        fmt.Println(v1, v1, v1)
        v1 = -674
        for fi32 := 0; fi32 < 15; fi32++ {
            if true {
                switch (-(v1)) {
                case 128:
                    v1 = (fi31 * fi31)
                case 2147483647:
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                case -666:
                    if (-1 != 100) {
                        v1 += (fi32 & -713)
                    }
                default:
                    v30 = v30
                }
            }
            _ = fi32
        }
        _ = fi31
    }
    v33 := GenS3{F0: 127, F1: 127}
    var v34 uint32 = 2147483648
    var v35 string = v30
    v34 = 411
    var v36 uint32 = (v34 / ((v34) | 1))
    var v37 int16 = 814
    v36 -= (v34 >> ((18446744073709551615) & 63))
    v34 = (v36 << ((689) & 63))
    v1 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
