package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int, p1 uint8) string {
    var v7 string = ""
    if (2 > 1) {
        v7 = v7
        v8 := []int{p0, p0, (p0 + 127)}
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: 18446744073709551615, F1: 393}
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int64 = 9223372036854775807
        v13 := GenS1{F0: 2, F1: p0}
        for fi14 := 0; fi14 < 11; fi14++ {
            p1 += p1
            v15 := GenS2{F0: 65535, F1: (v12 >> ((4294967295) & 63)), F2: 2}
            v16 := []int{fi14, p0}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return v7
}
func genFn1(p0 int64, p1 uint32) uint {
    if false {
        switch int(p1) {
        case -750:
            p1 = p1
        case 0:
            var v17 bool = false
            _ = v17
        default:
            p0 -= (p0 + -1)
        }
    }
    switch "hello" {
    case "\n\t":
        p1 += (p1 + (p1 * 0))
    }
    v18 := GenS3{F0: 127, F1: 285, F2: 737}
    _ = p0
    _ = p1
    _ = v18
    return 2
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 12
    v3 := []string{"", "", "foo bar"}
    var v4 uint = 449
    var v5 float32 = (float32(-2.25) * float32(3.14159))
    v6 := make(map[uint32]uint32)
    v19 := GenS4{F0: 0, F1: 42}
    v5 = v5
    v20 := GenS5{F0: 1}
    fmt.Println(v1, v5, v5)
    v21 := GenS1{F0: 4294967295, F1: -2147483648}
    if false {
        switch 9223372036854775807 {
        case -1:
            if (v5 >= (1.5 - v5)) {
                var v22 int = 128
                if ((v2 >= 546) || ("hello" == (("foo bar" + "a") + ("" + ("a" + ("foo bar" + "foo bar")))))) {
                    fmt.Println(v1, v22, v22)
                }
                v1 = v1
                _ = v22
            }
        case 4294967295:
            switch "foo bar" {
            case "\n\t":
                for fi23 := 0; fi23 < 16; fi23++ {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v2 += 107
                        _ = fi24
                    }
                    var v25 int32 = int32(v2)
                    switch fi23 {
                    case -504:
                        var v26 uint = (^(v4))
                        _ = v26
                    case -226:
                        v2 = v2
                    case 32767:
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v1 = v1
                            var v28 string = (("foo bar" + "\n\t") + ("hello" + "a"))
                            _ = fi27
                            _ = v28
                        }
                    default:
                        var v29 uint16 = 305
                        _ = v29
                    }
                    _ = fi23
                    _ = v25
                }
            }
        case 796:
            if ((-128 <= 1) && ((1000 > 1000) || true)) {
                v30 := GenS5{F0: 42}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v4, v1, v2)
                _ = v30
            } else {
                var v31 float64 = 1e10
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            if (fi32 <= fi32) {
                var v33 float32 = 0.0
                switch "" {
                case "":
                    var v34 float32 = ((1.5 - -2.25) + float32(((0.1 - ((-2.25 + (0.1 * 0.0)) * 1.5)) - 1e10)))
                    _ = v34
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v33
            }
            if true {
                if len(v3) > 0 {
                    v3[0] = (((("a" + "\n\t") + "foo bar") + ("foo bar" + "\n\t")) + "a")
                }
                for fi35 := 0; fi35 < 7; fi35++ {
                    var v36 uint8 = 1
                    var v37 uint64 = 42
                    fmt.Println(fi32, v36, v5)
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            } else {
                var v38 uint8 = 0
                var v39 bool = true
                _ = v38
                _ = v39
            }
            v5 += 1e10
            _ = fi32
        }
        for fi40 := 0; fi40 < 19; fi40++ {
            var v41 uint = v4
            delete(v6, v2)
            _ = fi40
            _ = v41
        }
    } else {
        _ = len(v3)
    }
    var v42 uint = genFn1(50, v2)
    if true {
        v5 = v5
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        fmt.Println(v1, v1, v2)
        _ = fi43
    }
    var v44 int8 = 21
    switch -2147483648 {
    case 32768:
        v5 += v5
    default:
        var v45 uint32 = (v2 | 42)
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
