package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int16; F2 int }
func genFn0(p0 int64, p1 int) uint16 {
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 += (p0 ^ (p0 / ((794) | 1)))
        v8 := []int{(^(p1)), 1}
        _ = fi7
        _ = v8
    }
    if (-1 >= 2147483647) {
        v9 := GenS1{F0: (p1 | p1), F1: 42}
        _ = v9
    } else {
        var v10 uint32 = 65535
        fmt.Println(p0)
        fmt.Println(p1, p1, p1)
        _ = v10
    }
    _ = p0
    _ = p1
    return 575
}
func genFn1(p0 uint64, p1 string, p2 uint32) int {
    v11 := GenS2{F0: 425}
    for fi12 := 0; fi12 < 6; fi12++ {
        p1 = p1
        v13 := GenS1{F0: (fi12 ^ fi12), F1: 1}
        _ = fi12
        _ = v13
    }
    fmt.Println(p1, p0)
    for fi14 := 0; fi14 < 2; fi14++ {
        if (0 != 960) {
            var v15 string = ("a" + p1)
            v16 := GenS1{F0: fi14, F1: (p2 * (^(p2)))}
            switch "" {
            case "foo bar":
                v15 = v15
            case "a":
                v17 := []int{-755, fi14, fi14, 4294967295}
                _ = v17
            case "":
                v18 := []int{}
                _ = v18
            default:
                var v19 float32 = 3.14159
                _ = v19
            }
            _ = v15
            _ = v16
        }
        p0 = (p0 + (p0 + 0))
        v20 := GenS3{F0: (p0 ^ (p0 & (p0 ^ (p0 % (((p0 - 0)) | 1))))), F1: 2147483647, F2: p0}
        _ = fi14
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 4294967295
}

func main() {
    var v1 uint16 = 174
    v2 := []string{"", "a"}
    var v3 int8 = 1
    v4 := []uint32{uint32(v1), 42}
    var v5 int16 = -32768
    var v6 float64 = (1e10 - 3.14159)
    switch -9223372036854775808 {
    case 32767:
        if (true || (!false)) {
            _ = len(v2)
            v5 = v5
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v21 string = "foo bar"
            _ = v21
        }
    case 0:
        switch ("hello" + "") {
        case "":
            v5 = (v5 * (^(v5)))
        case "foo bar":
            switch genFn1(18446744073709551615, ("foo bar" + (("hello" + "\n\t") + "a")), 2) {
            case 9223372036854775807:
                var v22 int = 127
                _ = v22
            case 4294967295:
                v1 = (^(v1))
            case 42:
                v5 = v5
            }
        default:
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
    }
    var v23 int = -977
    v5 -= v5
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 string = (("hello" + "a") + "")
        v1 = 915
        _ = fi24
        _ = v25
    }
    v1 += (^(v1))
    var v26 string = ""
    fmt.Println(v5, v5)
    var v27 uint = 65535
    if (!((v27 >> ((397) & 63)) > v27)) {
        if len(v4) > 0 {
            v4[0] = 2
        }
        var v28 float32 = 1.5
        _ = v28
    } else {
        v29 := GenS1{F0: (^(v23)), F1: 33}
        v5 += (v5 / ((v5) | 1))
        _ = v29
    }
    if true {
        v23 -= v23
    } else {
        var v30 int32 = -2147483648
        _ = v30
    }
    switch v23 {
    case 2147483647:
        for fi31 := 0; fi31 < 17; fi31++ {
            v23 = (^(fi31))
            var v32 int8 = 127
            fmt.Println(v3)
            _ = fi31
            _ = v32
        }
    }
    switch ("foo bar" + v26) {
    case "hello":
        v3 = 116
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 uint = (^(v27))
        if false {
            fmt.Println(v34, v27)
            if false {
                v26 = (v26 + v26)
            } else {
                _ = len(v2)
            }
            var v35 int8 = v3
            _ = v35
        } else {
            v36 := GenS1{F0: (v23 << ((42) & 63)), F1: 401}
            if (((437 > (v5 << (((v27 | (v34 << ((v34) & 63)))) & 63))) || true) && ((false && false) && ((v23 | genFn1(42, v26, 65535)) > (fi33 % ((127) | 1))))) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v37 := GenS4{F0: 0}
                _ = v37
            }
            _ = v36
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            var v39 int16 = v5
            v40 := GenS5{F0: (v1 % (((^(v1))) | 1)), F1: -22, F2: fi38}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
