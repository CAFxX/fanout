package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 int64, p1 uint32, p2 uint16) uint64 {
    if false {
        p1 = (p1 ^ (p1 >> ((0) & 63)))
    } else {
        var v4 bool = (358 > 173)
        switch ("hello" + "a") {
        case "hello":
            p1 -= 2147483647
        }
        var v5 float64 = 1e10
        _ = v4
        _ = v5
    }
    fmt.Println(p2)
    for fi6 := 0; fi6 < 13; fi6++ {
        p0 = (p0 - p0)
        _ = fi6
    }
    switch 4294967295 {
    case -9223372036854775808:
        fmt.Println(p2)
    case 9223372036854775807:
        if false {
            if ((-1 < -1) && (!true)) {
                p0 -= p0
            } else {
                v7 := GenS1{F0: 18446744073709551615}
                v8 := []int{127, -2147483648, -1, 9223372036854775807}
                for fi9 := 0; fi9 < 19; fi9++ {
                    if len(v8) > 0 {
                        v8[0] = (fi9 ^ (fi9 * fi9))
                    }
                    _ = fi9
                }
                _ = v7
                _ = v8
            }
            var v10 uint32 = p1
            _ = v10
        } else {
            var v11 int32 = 42
            switch "a" {
            case "hello":
                p2 -= 42
            case "a":
                if true {
                    fmt.Println(p1, p0, v11)
                    if ((p1 << ((686) & 63)) == 2147483647) {
                        v12 := []int{759, 32768, -818}
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 4294967295}
                        var v14 float32 = 0.0
                        _ = v13
                        _ = v14
                    }
                } else {
                    switch 900 {
                    case 69:
                        fmt.Println(p2, v11, p1)
                    default:
                        var v15 uint = 4294967295
                        _ = v15
                    }
                    if (("a" >= "hello") || (!false)) {
                        p0 = 9223372036854775807
                        v16 := []int{-1, 128, 32767}
                        fmt.Println(v11)
                        _ = v16
                    } else {
                        p1 = (p1 ^ p1)
                        for fi17 := 0; fi17 < 18; fi17++ {
                            var v18 uint16 = p2
                            _ = fi17
                            _ = v18
                        }
                        var v19 float32 = 0.1
                        _ = v19
                    }
                    switch 127 {
                    case 42:
                        v11 = v11
                    default:
                        v20 := []int{-1}
                        _ = v20
                    }
                }
            default:
                fmt.Println(v11, p2)
            }
            var v21 float32 = 1e10
            _ = v11
            _ = v21
        }
    case 0:
        var v22 int64 = p0
        _ = v22
    }
    if true {
        var v23 int16 = 32767
        for fi24 := 0; fi24 < 14; fi24++ {
            if false {
                if true {
                    if (true && false) {
                        p0 -= p0
                        v25 := GenS2{F0: p0, F1: p0, F2: -2147483648}
                        _ = v25
                    }
                }
                var v26 bool = false
                var v27 int8 = 0
                _ = v26
                _ = v27
            } else {
                v28 := []int{fi24, fi24}
                for fi29 := 0; fi29 < 11; fi29++ {
                    fmt.Println(fi29)
                    _ = fi29
                }
                v23 -= v23
                _ = v28
            }
            _ = fi24
        }
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return uint64(p1)
}

func main() {
    var v1 uint64 = 74
    var v2 float32 = 1.5
    var v3 uint32 = 2147483648
    if (true || false) {
        if (false || true) {
            v2 -= (float32(0.0) + v2)
            if false {
                var v30 int32 = 2147483647
                var v31 int64 = 42
                _ = v30
                _ = v31
            }
            if true {
                v32 := GenS3{F0: 752, F1: "hello"}
                v33 := GenS1{F0: 4294967295}
                _ = v32
                _ = v33
            }
        }
        var v34 string = "\n\t"
        _ = v34
    } else {
        v1 = 42
        var v35 uint8 = 127
        var v36 int = 32768
        _ = v35
        _ = v36
    }
    var v37 int16 = 185
    for fi38 := 0; fi38 < 10; fi38++ {
        v3 += v3
        _ = fi38
    }
    v1 = 445
    v39 := GenS4{F0: 2, F1: (v37 * 1000)}
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        var v40 int64 = 9223372036854775807
        _ = v40
    case "a":
        var v41 int32 = -90
        _ = v41
    case "foo bar":
        var v42 uint64 = (v1 | (v1 & v1))
        _ = v42
    }
    if (100 <= -121) {
        var v43 uint16 = 2
        if (!(v2 >= 1.5)) {
            v43 = 495
            var v44 uint32 = v3
            for fi45 := 0; fi45 < 9; fi45++ {
                v43 = 2
                var v46 uint8 = 42
                _ = fi45
                _ = v46
            }
            _ = v44
        } else {
            v3 = v3
            var v47 uint64 = 42
            _ = v47
        }
        var v48 uint32 = uint32(v43)
        _ = v43
        _ = v48
    } else {
        v1 = genFn0(0, (v3 | (v3 ^ (v3 - v3))), 0)
    }
    var v49 int8 = -99
    v50 := []int{32768, 32767, 1}
    v50 = append(v50, -425)
    v50 = append(v50, 127)
    fmt.Println(v49, v49, v1)
    var v51 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v49
    _ = v50
    _ = v51
}
