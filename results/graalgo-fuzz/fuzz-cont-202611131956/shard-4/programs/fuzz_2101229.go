package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint64, p1 bool) int32 {
    var v6 string = "\n\t"
    if ((65535 <= 1) || p1) {
        switch "\n\t" {
        case "foo bar":
            var v7 bool = (!(false || true))
            _ = v7
        case "hello":
            v8 := []int{-239, 9223372036854775807}
            _ = v8
        case "\n\t":
            p0 = (p0 << ((255) & 63))
        default:
            if (!p1) {
                fmt.Println(p0)
                switch (((v6 + ("foo bar" + (v6 + "a"))) + ((((v6 + (("" + ((((v6 + v6) + (v6 + "hello")) + v6) + v6)) + (v6 + ("foo bar" + v6)))) + v6) + v6) + (v6 + ("" + v6)))) + "a") {
                case "\n\t":
                    if true {
                        switch v6 {
                        case "foo bar":
                            var v9 float64 = -2.25
                            _ = v9
                        case "hello":
                            var v10 string = "a"
                            _ = v10
                        }
                    }
                case "foo bar":
                    v11 := GenS1{F0: 42}
                    _ = v11
                }
            }
        }
        v12 := []int{1, 2147483647, 4294967295}
        _ = v12
    } else {
        v13 := []int{-320, 1, 42}
        var v14 int16 = -18
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = v15
    _ = v6
    return 341
}
func genFn1(p0 int, p1 uint) uint8 {
    var v16 int8 = 127
    var v17 bool = false
    v18 := []int{(-(p0))}
    if (!((!(1.5 > 0.1)) && (39 < 65535))) {
        p0 -= 1
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]string)
    var v3 float64 = 0.1
    var v4 uint32 = 657
    v5 := make(map[string]string)
    v1 += (-(v1))
    _ = len(v5)
    var v19 uint16 = 944
    v1 -= -9223372036854775808
    _ = len(v5)
    fmt.Println(v4, v19)
    if true {
        if false {
            v1 = v1
        }
    } else {
        switch "a" {
        case "foo bar":
            for fi20 := 0; fi20 < 20; fi20++ {
                for fi21 := 0; fi21 < 12; fi21++ {
                    v3 -= ((v3 / v3) + 0.0)
                    var v22 uint32 = (v4 | v4)
                    var v23 uint16 = 537
                    var v24 int = (fi21 ^ 32768)
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v3 += float64(0.1)
                _ = fi20
            }
        case "hello":
            v4 -= v4
        case "a":
            for fi25 := 0; fi25 < 2; fi25++ {
                switch (^(fi25)) {
                case 337:
                    for fi26 := 0; fi26 < 3; fi26++ {
                        var v27 uint32 = v4
                        fmt.Println(fi26, v19)
                        _ = fi26
                        _ = v27
                    }
                case -347:
                    v3 = float64(-2.25)
                }
                _ = fi25
            }
        default:
            var v28 float32 = 0.0
            _ = v28
        }
        fmt.Println(v4)
        var v29 int64 = v1
        _ = v29
    }
    var v30 int64 = 42
    v31 := GenS2{F0: 2}
    var v32 int = -506
    v33 := []int{v32, (v32 << ((454) & 63)), (v32 / ((v32) | 1))}
    if (true || true) {
        v34 := GenS1{F0: 2147483647}
        if (!(("a" > ("foo bar" + "")) && (false || true))) {
            if (4294967295 == 18446744073709551615) {
                switch v32 {
                case 32768:
                    var v35 uint32 = v4
                    _ = v35
                case -749:
                    v3 = v3
                default:
                    v32 += v32
                }
                delete(v5, ((("a" + "a") + ("a" + "foo bar")) + ""))
                var v36 uint16 = v19
                _ = v36
            } else {
                v30 = v30
                v1 -= v1
                var v37 uint64 = 1
                _ = v37
            }
        } else {
            fmt.Println(v3, v32)
            var v38 uint32 = 2147483648
            for fi39 := 0; fi39 < 6; fi39++ {
                fmt.Println(v4)
                _ = fi39
            }
            _ = v38
        }
        var v40 int64 = (v1 & (v1 * (v30 & v30)))
        _ = v34
        _ = v40
    } else {
        _ = len(v33)
    }
    var v41 uint = 4294967295
    var v42 int8 = -128
    for fi43 := 0; fi43 < 7; fi43++ {
        v32 = (v32 % ((-1) | 1))
        switch fi43 {
        case 1:
            for fi44 := 0; fi44 < 7; fi44++ {
                var v45 uint16 = (^(v19))
                _ = fi44
                _ = v45
            }
        case 32767:
            v30 = -916
        case 9:
            v1 = (-(v30))
        }
        v41 -= (v41 ^ v41)
        _ = fi43
    }
    v46 := GenS2{F0: v41}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v46
    _ = v5
}
