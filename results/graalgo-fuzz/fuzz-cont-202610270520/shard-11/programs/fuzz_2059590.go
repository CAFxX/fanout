package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 uint8; F2 int32 }
func genFn0(p0 int64, p1 int) uint64 {
    v6 := []int{(p1 | p1), p1, p1}
    p1 += p1
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 int16, p1 int) uint16 {
    v7 := []int{128, 1, 4294967295, (p1 | p1)}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"hello", (("\n\t" + "foo bar") + "\n\t"), "\n\t"}
    v3 := make(map[int]string)
    var v4 uint64 = 755
    var v5 string = ("foo bar" + "a")
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            v5 = "hello"
            _ = fi8
        }
        var v9 uint8 = 2
        v5 = "\n\t"
        _ = v9
    } else {
        var v10 uint16 = genFn1(1, 2147483647)
        _ = v10
    }
    var v11 float32 = (0.0 - 1e10)
    _ = len(v2)
    if (v11 <= float32(0.0)) {
        v4 += (v4 | 930)
    }
    var v12 int8 = -1
    _ = len(v2)
    v13 := v3[0]
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int = 9223372036854775807
        var v16 int16 = -1
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        var v17 uint64 = 914
        var v18 string = v5
        v19 := v3[32767]
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 float32 = ((((v11 * ((v11 - 0.0) / 1.5)) + (v11 / v11)) + v11) * float32(-2.25))
    if false {
        v21 := v1[(v5 + "hello")]
        switch 0 {
        case 4294967295:
            var v22 uint32 = (v21 >> ((4294967295) & 63))
            _ = v22
        case -2147483648:
            fmt.Println(v13)
        case 127:
            switch v13 {
            case "a":
                if ((v21 / (((v21 - 0)) | 1)) <= (v21 / ((v21) | 1))) {
                    v21 -= v21
                    if false {
                        v23 := GenS1{F0: 128, F1: 42, F2: 42}
                        switch v5 {
                        case "":
                            fmt.Println(v4, v20, v21)
                        case "foo bar":
                            var v24 string = (((v5 + ("a" + "hello")) + v13) + v5)
                            _ = v24
                        default:
                            switch -576 {
                            case 78:
                                var v25 uint = 530
                                _ = v25
                            case 127:
                                delete(v1, v13)
                            case -2147483648:
                                fmt.Println(v20)
                            default:
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    v21 += v21
                                    v2 = append(v2, (v13 + (v5 + "a")))
                                    _ = fi26
                                }
                            }
                        }
                        for fi27 := 0; fi27 < 6; fi27++ {
                            v28 := v1[(("" + (((v13 + ((("hello" + (v5 + ("" + v5))) + v5) + (v13 + (((v13 + v13) + "hello") + v13)))) + ("foo bar" + v5)) + "a")) + (v13 + v13))]
                            fmt.Println(v5, v20, v5)
                            _ = fi27
                            _ = v28
                        }
                        _ = v23
                    }
                    var v29 uint = 209
                    _ = v29
                } else {
                    var v30 uint8 = 2
                    v31 := GenS2{F0: 42, F1: v30, F2: 954}
                    for fi32 := 0; fi32 < 20; fi32++ {
                        var v33 uint16 = 826
                        var v34 int64 = -475
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    _ = v30
                    _ = v31
                }
            default:
                for fi35 := 0; fi35 < 19; fi35++ {
                    v3[fi35] = v5
                    _ = fi35
                }
            }
        }
        _ = v21
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        switch 74 {
        case 127:
            v13 = (v5 + ("a" + ""))
        default:
            if true {
                if (genFn0(42, (fi36 - fi36)) < (v4 << ((0) & 63))) {
                    if (false && true) {
                        v12 = v12
                    }
                    for fi37 := 0; fi37 < 3; fi37++ {
                        fmt.Println(v13, v12)
                        _ = fi37
                    }
                    v12 -= v12
                } else {
                    v38 := v3[fi36]
                    v11 = (v11 * 1e10)
                    var v39 string = v5
                    _ = v38
                    _ = v39
                }
                var v40 uint = 0
                _ = v40
            }
        }
        _ = fi36
    }
    v4 += v4
    v2 = append(v2, "hello")
    v41 := GenS1{F0: 127, F1: 42, F2: -1}
    v4 = 0
    var v42 bool = (0.0 <= float64(((-2.25 / (v20 * v20)) * (v20 / v20))))
    switch v5 {
    case "\n\t":
        if false {
            for fi43 := 0; fi43 < 16; fi43++ {
                var v44 string = v13
                var v45 int16 = 42
                var v46 int32 = -1
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
            fmt.Println(v13, v13)
            v42 = (v42 && ((1 <= -1) && true))
        }
    case "a":
        fmt.Println(v42, v13)
    default:
        fmt.Println(v11, v42, v42)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
