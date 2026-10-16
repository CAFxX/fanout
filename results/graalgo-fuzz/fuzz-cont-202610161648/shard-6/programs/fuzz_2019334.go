package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
func genFn0(p0 bool, p1 uint32, p2 uint16) uint16 {
    var v7 int = 678
    v8 := GenS1{F0: uint32(p2), F1: p2}
    for fi9 := 0; fi9 < 8; fi9++ {
        if (!(!(537 < 0))) {
            var v10 int16 = 0
            v11 := GenS2{F0: 1, F1: -177, F2: 127}
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: (p1 % ((p1) | 1)), F1: (p2 + (p2 - 481))}
                fmt.Println(p2)
                _ = fi12
                _ = v13
            }
            _ = v10
            _ = v11
        }
        if true {
            for fi14 := 0; fi14 < 8; fi14++ {
                switch 127 {
                case 0:
                    p2 -= 255
                }
                _ = fi14
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := []int{(v7 % ((9223372036854775807) | 1)), (-(v7)), fi15, fi9}
                if false {
                    v16 = append(v16, fi9)
                } else {
                    p0 = p0
                    fmt.Println(p1, v7)
                    var v17 float32 = 3.14159
                    _ = v17
                }
                if false {
                    var v18 int32 = 808
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
        _ = fi9
    }
    v19 := []int{v7, (-(v7)), v7}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v7
    _ = v8
    return p2
}
func genFn1(p0 int32) int32 {
    var v20 bool = (true && ((-128 < 100) && true))
    v21 := []int{9223372036854775807}
    v22 := GenS3{F0: 2, F1: -1}
    _ = p0
    _ = v20
    _ = v21
    _ = v22
    return p0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 uint = 792
    v4 := []string{"hello"}
    var v5 uint64 = 0
    var v6 uint16 = 42
    var v23 int = 127
    v24 := GenS4{F0: (v5 / ((v5) | 1))}
    var v25 uint32 = 1
    for fi26 := 0; fi26 < 1; fi26++ {
        if false {
            v4 = append(v4, (("hello" + "foo bar") + ("" + "\n\t")))
        }
        var v27 int8 = -33
        v28 := v2[(("foo bar" + "hello") + (("" + "hello") + "hello"))]
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int32 = 42
    v30 := GenS5{F0: 1}
    var v31 uint = 1
    switch v23 {
    case 42:
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 bool = true
            _ = fi32
            _ = v33
        }
    }
    if ((!((3.14159 + float32((((1e10 * (-2.25 / -2.25)) * 1.5) / 1.0))) < float32(1e10))) || (3.14159 > (1.5 + 0.0))) {
        fmt.Println(v31)
        if true {
            switch v23 {
            case 42:
                if false {
                    for fi34 := 0; fi34 < 10; fi34++ {
                        if true {
                            if (1000 == 32767) {
                                var v35 uint64 = v5
                                _ = v35
                            } else {
                                if (((!false) || (-2.25 > 0.0)) || (!(1.5 != (float64(3.14159) * (1.5 / 1e10))))) {
                                    fmt.Println(v23, v3)
                                    if false {
                                        for fi36 := 0; fi36 < 10; fi36++ {
                                            switch (fi34 % (((^(fi34))) | 1)) {
                                            case 287:
                                                fmt.Println(v5, v31, v23)
                                            case 9223372036854775807:
                                                _ = len(v4)
                                            case 127:
                                                switch 457 {
                                                case 0:
                                                    v4 = append(v4, "a")
                                                case 127:
                                                    var v37 uint = v31
                                                    _ = v37
                                                default:
                                                    if (false && false) {
                                                        var v38 int64 = 9223372036854775807
                                                        _ = v38
                                                    }
                                                }
                                            default:
                                                var v39 uint = (^(v3))
                                                _ = v39
                                            }
                                            switch ((((("foo bar" + "hello") + "foo bar") + "\n\t") + "") + "\n\t") {
                                            case "\n\t":
                                                fmt.Println(v31, fi34, v31)
                                            case "":
                                                fmt.Println(v29)
                                            case "foo bar":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            _ = fi36
                                        }
                                    } else {
                                        for fi40 := 0; fi40 < 20; fi40++ {
                                            switch "a" {
                                            case "a":
                                                var v41 uint64 = (^(v5))
                                                _ = v41
                                            case "\n\t":
                                                var v42 uint8 = 203
                                                _ = v42
                                            default:
                                                if false {
                                                    v25 = (^(v25))
                                                }
                                            }
                                            if (float32(float64(0.0)) == float32(0.1)) {
                                                v43 := GenS6{F0: (v5 ^ v5)}
                                                _ = v43
                                            }
                                            if true {
                                                var v44 uint32 = (v25 * 42)
                                                v6 += v6
                                                _ = v44
                                            } else {
                                                var v45 int32 = (v29 ^ int32(v31))
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                fmt.Println(v23)
                                                _ = v45
                                            }
                                            _ = fi40
                                        }
                                        v46 := v1[(v23 + fi34)]
                                        _ = v46
                                    }
                                } else {
                                    v1[128] = "a"
                                    v5 += v5
                                }
                                _ = len(v2)
                                var v47 uint = (v3 % ((v31) | 1))
                                _ = v47
                            }
                            switch "\n\t" {
                            case "hello":
                                v48 := GenS7{F0: 176}
                                _ = v48
                            }
                        }
                        v23 += (^(v23))
                        _ = fi34
                    }
                } else {
                    var v49 int32 = v29
                    _ = v49
                }
            }
        }
    } else {
        v50 := GenS8{F0: true}
        _ = v50
    }
    var v51 float64 = (3.14159 - 0.1)
    v23 += 0
    var v52 uint64 = (v5 - v5)
    var v53 int32 = v29
    if len(v4) > 0 {
        v4[0] = ("a" + "")
    }
    v54 := GenS8{F0: true}
    delete(v2, "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
