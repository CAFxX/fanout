package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int8; F1 uint16; F2 int32 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint16; F1 uint8 }
type GenS9 struct { F0 int32; F1 uint; F2 int64 }
type GenS10 struct { F0 int32; F1 int16 }
type GenS11 struct { F0 int32 }
func genFn0(p0 string, p1 int32) uint {
    v7 := GenS1{F0: 0}
    p1 += (-(p1))
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS2{F0: -117}
        v10 := GenS1{F0: -1}
        var v11 string = "hello"
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1() uint64 {
    v12 := []int{-1}
    if false {
        var v13 int64 = -892
        switch 257 {
        case 787:
            fmt.Println(v13)
        case 537:
            if len(v12) > 0 {
                v12[0] = 32767
            }
        default:
            v14 := GenS3{F0: 42, F1: 430}
            _ = v14
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            v13 -= v13
            _ = fi15
        }
        _ = v13
    } else {
        var v16 int8 = -49
        _ = v16
    }
    switch -9223372036854775808 {
    case -534:
        var v17 string = ""
        _ = v17
    case -1:
        switch "a" {
        case "\n\t":
            switch 9223372036854775807 {
            case -843:
                var v18 uint8 = 247
                _ = v18
            case 32767:
                if false {
                    var v19 bool = false
                    _ = v19
                }
            }
        case "a":
            var v20 uint = 773
            _ = v20
        }
    case 920:
        var v21 int = 128
        _ = v21
    }
    if true {
        var v22 int8 = 63
        fmt.Println(v22)
        if true {
            for fi23 := 0; fi23 < 20; fi23++ {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 float64 = ((1e10 * (1e10 * (((float64(1e10) * float64(3.14159)) / -2.25) * 1e10))) * (0.0 - 1.5))
                    v22 -= -128
                    fmt.Println(v25, fi23, fi24)
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
            v22 += v22
        }
        _ = v22
    } else {
        var v26 uint = genFn0("\n\t", 42)
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 uint = 42
        v29 := GenS4{F0: -1, F1: v28}
        v30 := GenS5{F0: -72, F1: 42, F2: -474}
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    _ = v12
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", ("" + "\n\t"), ("foo bar" + "\n\t"), (("foo bar" + "hello") + ("" + ("a" + "foo bar")))}
    var v3 int16 = -925
    var v4 int16 = 0
    v5 := []string{"foo bar", ("" + ""), "\n\t", "hello"}
    var v6 float64 = (0.1 * 1e10)
    if ((true && (!(false && false))) || ((float32(v6) * 0.0) > -2.25)) {
        v4 = v3
        v6 = (v6 * v6)
    }
    fmt.Println(v3)
    switch ("foo bar" + "a") {
    case "foo bar":
        var v31 int64 = -825
        _ = v31
    case "a":
        if false {
            var v32 bool = (false && false)
            switch "\n\t" {
            case "\n\t":
                var v33 float64 = v6
                _ = v33
            }
            _ = v32
        } else {
            v34 := GenS4{F0: 0, F1: 0}
            v35 := GenS6{F0: 100, F1: 255}
            if (((255 >= 18446744073709551615) || false) && true) {
                for fi36 := 0; fi36 < 4; fi36++ {
                    v4 -= (v4 << ((genFn0("", -98)) & 63))
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v4, v6)
                    case "\n\t":
                        var v37 uint8 = 72
                        _ = v37
                    case "":
                        for fi38 := 0; fi38 < 4; fi38++ {
                            var v39 uint8 = 227
                            var v40 string = (("hello" + ("a" + "hello")) + "\n\t")
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                    default:
                        var v41 uint32 = 574
                        _ = v41
                    }
                    _ = fi36
                }
            } else {
                if ((42 >= 42) || ((genFn1() >= 2) || false)) {
                    if false {
                        fmt.Println(v3, v4)
                    } else {
                        var v42 float32 = float32(1.5)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v43 := GenS7{F0: (false && (true && (4294967295 == 1))), F1: 2147483647}
                        _ = v42
                        _ = v43
                    }
                }
            }
            _ = v34
            _ = v35
        }
    case "\n\t":
        var v44 uint16 = 65535
        _ = v44
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        v46 := v1[65535]
        _ = fi45
        _ = v46
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    delete(v1, 65535)
    v4 += 934
    switch 4294967295 {
    case 0:
        v47 := GenS4{F0: 127, F1: 1}
        _ = v47
    case -105:
        v48 := v1[2147483647]
        _ = v48
    case 1:
        if false {
            switch ((("a" + "hello") + "\n\t") + "a") {
            case "":
                var v49 uint = 2
                _ = v49
            case "a":
                v6 -= v6
            default:
                v50 := GenS8{F0: 291, F1: 0}
                _ = v50
            }
            v3 = v3
            if len(v5) > 0 {
                v5[0] = (("foo bar" + "") + "")
            }
        } else {
            v51 := GenS8{F0: 841, F1: 1}
            v4 = 0
            _ = v51
        }
    }
    fmt.Println(v3, v6, v3)
    v4 -= v3
    if false {
        if (2 != 1) {
            for fi52 := 0; fi52 < 10; fi52++ {
                fmt.Println(fi52)
                v4 = v4
                _ = fi52
            }
            v53 := GenS9{F0: 1, F1: genFn0("\n\t", 0), F2: 1}
            _ = v53
        }
    } else {
        v54 := GenS10{F0: 42, F1: v4}
        var v55 uint8 = 42
        for fi56 := 0; fi56 < 3; fi56++ {
            _ = len(v2)
            fmt.Println(v55, v4)
            _ = fi56
        }
        _ = v54
        _ = v55
    }
    var v57 int = 2147483647
    v58 := GenS11{F0: 0}
    if false {
        for fi59 := 0; fi59 < 4; fi59++ {
            v60 := GenS1{F0: 877}
            v4 = v4
            var v61 uint64 = 2
            _ = fi59
            _ = v60
            _ = v61
        }
        fmt.Println(v57, v6, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
