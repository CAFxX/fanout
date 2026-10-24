package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0() uint64 {
    var v5 int = 127
    v5 = (v5 << ((4294967295) & 63))
    v5 = 1
    v6 := []int{(v5 + 32768)}
    var v7 int16 = 8
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 string) bool {
    var v8 float64 = float64(0.0)
    switch -2147483648 {
    case -2147483648:
        var v9 int64 = -9223372036854775808
        _ = v9
    case 32768:
        v10 := []int{-2147483648, 9223372036854775807}
        _ = v10
    default:
        fmt.Println(p0)
    }
    var v11 float32 = 3.14159
    _ = p0
    _ = v11
    _ = v8
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 206
    v3 := make(map[string]string)
    var v4 int8 = 42
    switch (("\n\t" + "foo bar") + "a") {
    case "a":
        v2 = v2
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        var v12 uint8 = 72
        _ = v12
    }
    if false {
        var v13 bool = (3.14159 > ((0.0 - 3.14159) - 1.5))
        switch "foo bar" {
        case "a":
            v13 = true
        }
        _ = v13
    }
    v14 := GenS1{F0: 42}
    v2 = v2
    v15 := []int{v2, v2}
    if ("hello" <= ("hello" + "")) {
        if (!(170 >= -270)) {
            for fi16 := 0; fi16 < 20; fi16++ {
                for fi17 := 0; fi17 < 19; fi17++ {
                    v4 = (v4 % ((-1) | 1))
                    var v18 uint32 = 0
                    if true {
                        fmt.Println(fi17)
                        fmt.Println(v2)
                    }
                    if true {
                        var v19 uint8 = 128
                        _ = v19
                    } else {
                        var v20 uint16 = 321
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                if (0.1 < (1e10 / (0.1 / 1.5))) {
                    v2 = v2
                } else {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 int16 = 1
                        if ((v4 << ((2) & 63)) > 100) {
                            v2 = 9223372036854775807
                            fmt.Println(v23)
                        }
                        v24 := GenS2{F0: -924}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                v2 -= (v2 * v2)
                _ = fi21
            }
            if true {
                fmt.Println(v4)
                v15 = append(v15, (v2 << ((42) & 63)))
            } else {
                v25 := v3[("foo bar" + (("hello" + "hello") + "foo bar"))]
                v4 -= (v4 / (((v4 | (v4 + -29))) | 1))
                if false {
                    v25 = v25
                }
                _ = v25
            }
        }
    } else {
        for fi26 := 0; fi26 < 14; fi26++ {
            for fi27 := 0; fi27 < 4; fi27++ {
                switch v2 {
                case -235:
                    v2 += fi26
                case 0:
                    if true {
                        v28 := GenS1{F0: 127}
                        if false {
                            v4 += (-(v4))
                        } else {
                            if false {
                                var v29 float32 = 0.1
                                _ = v29
                            }
                            var v30 uint = 0
                            _ = v30
                        }
                        v2 += v2
                        _ = v28
                    }
                default:
                    v31 := v1[(v2 >> ((65535) & 63))]
                    _ = v31
                }
                _ = fi27
            }
            v32 := GenS1{F0: 255}
            _ = fi26
            _ = v32
        }
        v2 -= v2
        var v33 int8 = v4
        _ = v33
    }
    if len(v15) > 0 {
        v15[0] = (v2 ^ v2)
    }
    fmt.Println(v4, v2)
    v2 -= v2
    v34 := GenS3{F0: 1, F1: 4294967295}
    if true {
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 uint8 = 1
            v4 -= (v4 & (v4 - 18))
            fmt.Println(fi35)
            _ = fi35
            _ = v36
        }
        if true {
            for fi37 := 0; fi37 < 8; fi37++ {
                fmt.Println(fi37, v2, fi37)
                _ = fi37
            }
            v4 += v4
        } else {
            v38 := GenS1{F0: 42}
            var v39 uint = 1
            for fi40 := 0; fi40 < 11; fi40++ {
                delete(v1, -331)
                fmt.Println(v2, v2, v2)
                _ = fi40
            }
            _ = v38
            _ = v39
        }
        v41 := GenS3{F0: 847, F1: 513}
        _ = v41
    } else {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        v4 += 42
    }
    v4 = -1
    for fi42 := 0; fi42 < 1; fi42++ {
        fmt.Println(v4, fi42)
        switch v2 {
        case 8:
            if (-2147483648 > 379) {
                v1[(fi42 | 445)] = fi42
            }
        case 42:
            for fi43 := 0; fi43 < 1; fi43++ {
                if ((673 == 42) && (!(true || false))) {
                    v15 = append(v15, 1)
                    if len(v15) > 0 {
                        v15[0] = fi42
                    }
                    var v44 int32 = -2147483648
                    _ = v44
                }
                _ = fi43
            }
        }
        _ = fi42
    }
    fmt.Println(v4)
    v45 := GenS3{F0: 1, F1: 65535}
    var v46 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v45
    _ = v46
}
