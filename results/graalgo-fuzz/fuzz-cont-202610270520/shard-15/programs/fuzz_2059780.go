package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int8; F1 int64 }
func genFn0(p0 int32) string {
    switch "\n\t" {
    case "\n\t":
        v7 := GenS1{F0: -987, F1: -128}
        _ = v7
    case "hello":
        p0 -= 0
    }
    p0 = p0
    _ = p0
    return "foo bar"
}
func genFn1(p0 int) int {
    if ((!(!false)) || (false && (798 != 0))) {
        p0 = 2147483647
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint32 = 134
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := []int{42, p0}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    if (true && false) {
        v12 := GenS1{F0: 0, F1: -128}
        p0 = 9223372036854775807
        if false {
            fmt.Println(p0)
            var v13 int = -9223372036854775808
            _ = v13
        }
        _ = v12
    }
    _ = p0
    return (p0 / ((128) | 1))
}

func main() {
    var v1 int32 = 1
    var v2 int = 567
    v3 := []string{"hello"}
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    v6 := make(map[string]bool)
    fmt.Println(v2, v2)
    var v14 uint32 = 2147483647
    switch "\n\t" {
    case "\n\t":
        var v15 int = (-(v2))
        _ = v15
    case "a":
        fmt.Println(v14, v2, v1)
    }
    _ = len(v5)
    v16 := GenS1{F0: 691, F1: 38}
    var v17 uint64 = 18446744073709551615
    switch "a" {
    case "":
        v18 := v5[(v14 / ((v14) | 1))]
        _ = v18
    case "a":
        delete(v5, (^(v14)))
    default:
        v1 += -554
    }
    if true {
        for fi19 := 0; fi19 < 6; fi19++ {
            switch "hello" {
            case "\n\t":
                v1 = v1
            default:
                v1 += (v1 ^ 2147483647)
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                if true {
                    if ((fi19 - genFn1(-99)) <= (fi19 | (v2 & (v2 << ((440) & 63))))) {
                        v21 := GenS1{F0: -1, F1: -1}
                        var v22 uint32 = (v14 / ((v14) | 1))
                        v3 = append(v3, ((("\n\t" + "") + "foo bar") + ("foo bar" + ("\n\t" + ((((("foo bar" + "foo bar") + "hello") + "\n\t") + ((("" + (("a" + "\n\t") + "\n\t")) + "") + (("" + (("a" + "a") + "hello")) + ("foo bar" + ("a" + "hello"))))) + (("foo bar" + (("\n\t" + "foo bar") + "foo bar")) + "\n\t"))))))
                        if ((0.0 + 0.0) <= 3.14159) {
                            if (!(true || false)) {
                                var v23 string = ("\n\t" + "foo bar")
                                var v24 int8 = 42
                                if true {
                                    var v25 int8 = v24
                                    _ = v25
                                }
                                _ = v23
                                _ = v24
                            }
                            var v26 float32 = -2.25
                            v22 -= (v22 - 2147483647)
                            _ = v26
                        }
                        fmt.Println(fi20, fi19, v2)
                        _ = v21
                        _ = v22
                    } else {
                        v1 = v1
                        if (true || true) {
                            v27 := GenS2{F0: 337}
                            v2 -= v2
                            var v28 bool = false
                            _ = v27
                            _ = v28
                        } else {
                            v14 = (v14 ^ 2147483647)
                            v14 += (v14 >> ((255) & 63))
                            v1 += (^(v1))
                        }
                        v29 := GenS3{F0: genFn1((v2 << ((295) & 63))), F1: v14}
                        _ = v29
                    }
                    v14 = 775
                    switch v2 {
                    case 976:
                        var v30 int32 = v1
                        _ = v30
                    }
                }
                v31 := GenS4{F0: (v17 | (v17 | 30))}
                v17 = 2
                _ = fi20
                _ = v31
            }
            for fi32 := 0; fi32 < 17; fi32++ {
                v33 := GenS3{F0: v2, F1: v14}
                if (1 > 127) {
                    v14 = 2147483647
                    var v34 uint32 = (v14 << ((18446744073709551615) & 63))
                    _ = v34
                } else {
                    var v35 uint8 = 27
                    _ = v35
                }
                v36 := v6["hello"]
                v37 := v4[(fi19 * 817)]
                _ = fi32
                _ = v33
                _ = v36
                _ = v37
            }
            _ = fi19
        }
        if true {
            fmt.Println(v14, v17)
            switch "hello" {
            case "\n\t":
                fmt.Println(v2, v1)
            case "a":
                v38 := GenS1{F0: 32767, F1: -15}
                _ = v38
            case "hello":
                if ((v1 | v1) <= v1) {
                    for fi39 := 0; fi39 < 15; fi39++ {
                        _ = len(v4)
                        _ = fi39
                    }
                }
            default:
                v40 := GenS5{F0: (-(v1)), F1: 1}
                _ = v40
            }
        } else {
            v2 = (^(v2))
            v2 = int(v1)
            if len(v3) > 0 {
                v3[0] = "hello"
            }
        }
    } else {
        v3 = append(v3, "")
    }
    var v41 float64 = 3.14159
    v41 += ((v41 - (((1e10 / v41) - v41) / v41)) - (float64((1.5 * 1.5)) - 1.5))
    if true {
        v41 = v41
        v41 = v41
        v3 = append(v3, ((("a" + "a") + "") + (("a" + ("hello" + "a")) + ("hello" + "foo bar"))))
    }
    delete(v4, genFn1(-2147483648))
    if false {
        fmt.Println(v2, v14, v41)
    }
    v42 := GenS6{F0: -1, F1: 504}
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 bool = (388 <= 32767)
        _ = fi43
        _ = v44
    }
    if (255 < 0) {
        if (!(!false)) {
            v6["a"] = false
            var v45 uint8 = 255
            _ = v45
        }
    } else {
        for fi46 := 0; fi46 < 4; fi46++ {
            fmt.Println(v2)
            v3 = append(v3, ("foo bar" + ""))
            _ = fi46
        }
        v17 -= (v17 & (v17 / ((v17) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
