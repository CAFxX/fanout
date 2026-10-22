package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 string; F2 int16 }
type GenS9 struct { F0 int16 }
func genFn0(p0 int16) int32 {
    p0 = p0
    p0 += p0
    v4 := GenS1{F0: 32767, F1: 235, F2: -2147483648}
    switch 482 {
    case 544:
        switch "hello" {
        case "hello":
            p0 = 42
        case "a":
            var v5 int64 = 1
            _ = v5
        }
    case 0:
        v6 := GenS2{F0: (("hello" + ("" + "a")) + "a")}
        _ = v6
    }
    p0 = p0
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 int64, p1 uint8, p2 uint8) uint32 {
    var v7 int = -1
    v8 := []int{v7, 0, 0}
    var v9 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint64 = 2
    var v2 int32 = 869
    var v3 string = ((("hello" + (("hello" + "a") + "a")) + "hello") + ("a" + "a"))
    for fi10 := 0; fi10 < 14; fi10++ {
        v3 = (v3 + "foo bar")
        _ = fi10
    }
    v3 = (v3 + v3)
    var v11 float32 = -2.25
    switch 552 {
    case -365:
        var v12 float64 = 3.14159
        _ = v12
    case 2147483647:
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                if (54 > 160) {
                    switch fi14 {
                    case -2147483648:
                        var v15 float32 = (v11 - (float32(float64(v11)) - (v11 - v11)))
                        _ = v15
                    case 1:
                        switch (fi13 << ((65535) & 63)) {
                        case 503:
                            switch "\n\t" {
                            case "\n\t":
                                v1 -= (v1 - (v1 << ((18446744073709551615) & 63)))
                            case "hello":
                                v16 := []int{fi14, (fi14 >> ((433) & 63)), (fi13 >> ((65535) & 63))}
                                _ = v16
                            case "foo bar":
                                v11 = float32((((-2.25 - 0.0) - 3.14159) - 0.0))
                            }
                        case -893:
                            v11 = v11
                        }
                        v3 = ""
                        v17 := GenS1{F0: 65535, F1: -9223372036854775808, F2: 42}
                        _ = v17
                    default:
                        v3 = "\n\t"
                    }
                    switch v3 {
                    case "foo bar":
                        var v18 uint = 2
                        _ = v18
                    case "hello":
                        v2 = -416
                    }
                    v19 := GenS3{F0: 64}
                    var v20 uint32 = 805
                    _ = v19
                    _ = v20
                } else {
                    if true {
                        if true {
                            v21 := GenS3{F0: 42}
                            var v22 uint8 = 128
                            _ = v21
                            _ = v22
                        }
                        if (v2 < v2) {
                            fmt.Println(v3, fi13)
                        } else {
                            v1 = v1
                            v23 := GenS4{F0: (v2 / ((v2) | 1)), F1: 4294967295}
                            var v24 int = fi13
                            _ = v23
                            _ = v24
                        }
                        v1 = v1
                    } else {
                        var v25 int64 = -303
                        _ = v25
                    }
                    fmt.Println(v1, fi13, v3)
                }
                _ = fi14
            }
            _ = fi13
        }
    case 984:
        fmt.Println(v11, v3, v2)
    default:
        v26 := []int{4294967295, 32767}
        _ = v26
    }
    var v27 uint64 = (v1 | v1)
    v27 = 2
    for fi28 := 0; fi28 < 9; fi28++ {
        switch (fi28 << ((2) & 63)) {
        case 32768:
            var v29 int64 = -9223372036854775808
            _ = v29
        }
        v11 += (v11 + -2.25)
        for fi30 := 0; fi30 < 9; fi30++ {
            var v31 int32 = genFn0(-32768)
            var v32 uint = 4294967295
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    v3 = v3
    switch "a" {
    case "foo bar":
        v33 := GenS5{F0: -32768}
        _ = v33
    case "":
        var v34 bool = (((false && (956 < 42)) || ((-2.25 > 1e10) || false)) && (!false))
        _ = v34
    case "a":
        switch "foo bar" {
        case "foo bar":
            for fi35 := 0; fi35 < 14; fi35++ {
                if (!((2 < 1) && true)) {
                    v2 = (v2 & v2)
                } else {
                    var v36 int8 = -1
                    _ = v36
                }
                if false {
                    fmt.Println(v3, v11)
                    for fi37 := 0; fi37 < 10; fi37++ {
                        var v38 bool = (127 >= 127)
                        var v39 float32 = (v11 * ((float32(((-2.25 - -2.25) / (0.1 - 3.14159))) - (v11 - (0.1 / (v11 + v11)))) * v11))
                        var v40 uint16 = 1
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                }
                v41 := GenS6{F0: "foo bar", F1: v2}
                _ = fi35
                _ = v41
            }
        case "hello":
            v11 = 3.14159
        case "a":
            v2 = -20
        default:
            if true {
                for fi42 := 0; fi42 < 14; fi42++ {
                    v27 = (v27 | (v1 % ((v27) | 1)))
                    if (((v11 / 0.1) - (3.14159 - v11)) == (1e10 * v11)) {
                        var v43 string = (("foo bar" + v3) + v3)
                        v44 := GenS3{F0: 414}
                        var v45 int8 = 0
                        _ = v43
                        _ = v44
                        _ = v45
                    } else {
                        v46 := []int{fi42, fi42, (-(fi42))}
                        v3 = ("hello" + (v3 + (v3 + v3)))
                        var v47 float64 = 0.1
                        _ = v46
                        _ = v47
                    }
                    v48 := []int{-216}
                    _ = fi42
                    _ = v48
                }
                v11 -= v11
                fmt.Println(v27)
            } else {
                v49 := GenS7{F0: 0}
                if (true && true) {
                    for fi50 := 0; fi50 < 8; fi50++ {
                        v11 += v11
                        v51 := []int{}
                        _ = fi50
                        _ = v51
                    }
                    v52 := []int{127, 4294967295}
                    _ = v52
                } else {
                    v53 := []int{}
                    var v54 uint16 = 42
                    var v55 int = 1
                    _ = v53
                    _ = v54
                    _ = v55
                }
                fmt.Println(v27)
                _ = v49
            }
        }
    }
    v3 = "foo bar"
    if (!(true || false)) {
        if (128 <= 200) {
            var v56 int16 = 1
            switch v3 {
            case "a":
                for fi57 := 0; fi57 < 15; fi57++ {
                    var v58 float32 = v11
                    _ = fi57
                    _ = v58
                }
            case "hello":
                fmt.Println(v27, v11)
            case "foo bar":
                fmt.Println(v3)
            default:
                v59 := GenS4{F0: v2, F1: 58}
                _ = v59
            }
            v1 += 4294967295
            _ = v56
        }
    } else {
        fmt.Println(v2, v11)
    }
    switch -9223372036854775808 {
    case 32768:
        v60 := GenS8{F0: 42, F1: "foo bar", F2: 1000}
        _ = v60
    case -363:
        var v61 int64 = -9223372036854775808
        _ = v61
    case -101:
        v62 := GenS9{F0: -1}
        _ = v62
    }
    v63 := []int{32768, -1}
    if (2 > 127) {
        v1 -= (v27 << ((65535) & 63))
        switch 871 {
        case 4294967295:
            fmt.Println(v2, v3)
        case 9223372036854775807:
            v64 := GenS5{F0: 42}
            _ = v64
        }
    }
    if len(v63) > 0 {
        v63[0] = 128
    }
    v65 := GenS4{F0: v2, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v63:", len(v63), v63)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v63
    _ = v65
}
