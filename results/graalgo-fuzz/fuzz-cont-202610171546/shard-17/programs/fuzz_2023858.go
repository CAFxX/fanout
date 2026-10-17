package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint; F2 uint32 }
type GenS5 struct { F0 int16; F1 int32; F2 bool }
type GenS6 struct { F0 uint64 }
func genFn0() int64 {
    var v7 float32 = 1.5
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int = -9223372036854775808
        v9 += 4294967295
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if true {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint64 = 404
                if (false && (("a" >= ("\n\t" + (("a" + "foo bar") + "a"))) || (100 >= 1))) {
                    v13 := GenS1{F0: 4294967295, F1: 32767}
                    _ = v13
                }
                v14 := []int{-878, fi10, (fi10 - -1)}
                _ = fi11
                _ = v12
                _ = v14
            }
        } else {
            switch (^(fi10)) {
            case -24:
                var v15 uint32 = 2147483647
                _ = v15
            default:
                v16 := []int{fi10}
                _ = v16
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := []int{fi17, fi17}
                v7 -= (v7 * 0.1)
                if (78 <= 127) {
                    if true {
                        v18 = append(v18, (fi17 - fi10))
                        switch fi10 {
                        case 640:
                            var v19 uint16 = 938
                            _ = v19
                        }
                    }
                    var v20 int16 = -1
                    _ = v20
                } else {
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    v7 += v7
                }
                _ = fi17
                _ = v18
            }
        }
        switch ("foo bar" + "foo bar") {
        case "":
            fmt.Println(fi10)
        case "a":
            var v21 int64 = -9223372036854775808
            _ = v21
        case "hello":
            v22 := []int{-193, (fi10 << ((775) & 63)), (^(fi10))}
            _ = v22
        default:
            v23 := GenS1{F0: 42, F1: 1000}
            _ = v23
        }
        var v24 uint16 = 0
        _ = fi10
        _ = v24
    }
    var v25 uint64 = 1
    _ = v25
    _ = v7
    return -310
}
func genFn1(p0 int, p1 uint16) uint {
    v26 := GenS1{F0: 2, F1: -32768}
    v27 := GenS2{F0: -1}
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 276
}

func main() {
    var v1 float64 = (-2.25 - ((0.0 - -2.25) / 0.1))
    var v2 uint8 = 127
    v3 := make(map[string]uint32)
    v4 := []uint32{42, 825, 2, 2147483648}
    var v5 uint32 = 1
    var v6 float64 = 0.0
    for fi28 := 0; fi28 < 1; fi28++ {
        v2 = (v2 % ((v2) | 1))
        var v29 uint64 = 18446744073709551615
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v5 = v5
        _ = fi30
    }
    fmt.Println(v1, v6, v5)
    v31 := GenS3{F0: 18446744073709551615, F1: 42, F2: (v5 - 65535)}
    var v32 uint16 = 2
    var v33 uint = genFn1(-9223372036854775808, v32)
    if true {
        for fi34 := 0; fi34 < 19; fi34++ {
            for fi35 := 0; fi35 < 15; fi35++ {
                if (1 <= 9223372036854775807) {
                    var v36 int64 = 54
                    _ = v36
                } else {
                    fmt.Println(v5, fi34, v33)
                    var v37 uint32 = v5
                    var v38 int = fi34
                    _ = v37
                    _ = v38
                }
                v33 = (v33 % ((v33) | 1))
                if (1000 >= -925) {
                    delete(v3, "")
                    v1 = float64(1e10)
                } else {
                    var v39 uint64 = 223
                    fmt.Println(v39, v5)
                    _ = v39
                }
                fmt.Println(fi34, fi34, v2)
                var v40 uint64 = 2
                switch (fi35 * -312) {
                case 128:
                    v6 = (0.1 * 3.14159)
                default:
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            v4 = append(v4, uint32(v32))
                            var v41 float64 = (v1 + v1)
                            v42 := v3[""]
                            _ = v41
                            _ = v42
                        }
                    case "foo bar":
                        fmt.Println(v2, fi35)
                    case "":
                        fmt.Println(v2)
                    default:
                        var v43 string = (((("foo bar" + "foo bar") + "hello") + "foo bar") + "\n\t")
                        _ = v43
                    }
                }
                v44 := GenS4{F0: -128, F1: uint(v40), F2: (v5 & v5)}
                _ = fi35
                _ = v40
                _ = v44
            }
            if true {
                switch "hello" {
                case "foo bar":
                    switch "" {
                    case "":
                        var v45 int16 = -1
                        _ = v45
                    case "foo bar":
                        for fi46 := 0; fi46 < 6; fi46++ {
                            var v47 bool = (65535 == uint32(v32))
                            var v48 int8 = 42
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                    case "hello":
                        var v49 int = (-(fi34))
                        _ = v49
                    default:
                        var v50 uint = genFn1((fi34 + -383), v32)
                        _ = v50
                    }
                case "":
                    v3[("\n\t" + "foo bar")] = 2147483648
                case "\n\t":
                    for fi51 := 0; fi51 < 14; fi51++ {
                        v4 = append(v4, 885)
                        _ = fi51
                    }
                }
                if true {
                    var v52 uint64 = 42
                    _ = v52
                }
            } else {
                if (!true) {
                    v4 = append(v4, v5)
                    _ = len(v4)
                } else {
                    for fi53 := 0; fi53 < 7; fi53++ {
                        switch "foo bar" {
                        case "hello":
                            v54 := GenS5{F0: 42, F1: -1, F2: (false || (!(!(18446744073709551615 <= 4294967295))))}
                            _ = v54
                        case "\n\t":
                            var v55 int = fi34
                            _ = v55
                        }
                        v56 := GenS4{F0: 1, F1: 1, F2: v5}
                        var v57 uint16 = (v32 ^ 2)
                        _ = fi53
                        _ = v56
                        _ = v57
                    }
                    v58 := GenS2{F0: -1}
                    v33 -= v33
                    _ = v58
                }
                var v59 int64 = 1
                v32 = 959
                _ = v59
            }
            v5 = (v5 + uint32(v2))
            _ = fi34
        }
        v1 = 0.0
    }
    delete(v3, ("hello" + ("\n\t" + ("a" + "a"))))
    fmt.Println(v32)
    v3["\n\t"] = (v5 / ((185) | 1))
    var v60 int16 = 32767
    switch (((((("\n\t" + "foo bar") + "") + (("" + "\n\t") + "foo bar")) + (("hello" + "foo bar") + "hello")) + "foo bar") + "a") {
    case "":
        switch 32767 {
        case 32768:
            switch "\n\t" {
            case "a":
                v61 := GenS6{F0: 1}
                _ = v61
            case "hello":
                v62 := GenS4{F0: 100, F1: genFn1(2147483647, uint16(v2)), F2: 15}
                _ = v62
            }
        }
    }
    v1 = ((v1 / v1) + v1)
    fmt.Println(v5)
    var v63 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v63
}
