package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 string; F2 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }
func genFn0() uint16 {
    if (true || false) {
        switch "a" {
        case "\n\t":
            v4 := GenS1{F0: -9223372036854775808, F1: 928, F2: 842}
            _ = v4
        default:
            for fi5 := 0; fi5 < 18; fi5++ {
                var v6 uint16 = 42
                if (false && (1000 != 0)) {
                    var v7 float32 = 0.1
                    fmt.Println(v7)
                    var v8 uint = 833
                    _ = v7
                    _ = v8
                }
                _ = fi5
                _ = v6
            }
        }
    } else {
        var v9 uint64 = 2
        _ = v9
    }
    if (false || (42 != -63)) {
        var v10 int8 = -1
        _ = v10
    } else {
        var v11 uint32 = 890
        v11 -= v11
        v12 := []int{2147483647}
        _ = v11
        _ = v12
    }
    var v13 uint32 = 2147483647
    _ = v13
    return 810
}
func genFn1(p0 uint, p1 uint, p2 int8) int64 {
    p0 += 2
    p1 += p1
    v14 := []int{128, -193, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = v1
    v3 := make(map[string]uint32)
    v15 := GenS2{F0: 42, F1: 359}
    var v16 uint16 = 277
    v17 := GenS3{F0: 9223372036854775807, F1: (("hello" + ("" + "\n\t")) + ("hello" + (((("a" + ("hello" + (("" + "hello") + "a"))) + "") + "foo bar") + ("" + ("\n\t" + ""))))), F2: 2147483647}
    if false {
        if ((v1 * 127) != v2) {
            v2 += (v1 % ((v2) | 1))
            var v18 int32 = -2147483648
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            if false {
                for fi20 := 0; fi20 < 3; fi20++ {
                    v2 += 128
                    fmt.Println(v2, fi20, fi20)
                    v21 := GenS4{F0: -289, F1: ("a" + "\n\t"), F2: 374}
                    var v22 int32 = 42
                    if ((v16 + v16) <= (v16 * (^(v16)))) {
                        v1 += (^(v1))
                        fmt.Println(fi19, fi20)
                    } else {
                        switch int(v22) {
                        case 0:
                            v1 += (v2 << ((65535) & 63))
                        case 213:
                            var v23 bool = ((false && false) || (!((fi20 % ((127) | 1)) >= (fi20 | 4294967295))))
                            _ = v23
                        case 32768:
                            v24 := v3[((("a" + "") + "hello") + "")]
                            _ = v24
                        }
                        if ((v22 << ((2) & 63)) == v22) {
                            _ = len(v3)
                            fmt.Println(v2, fi19)
                            v25 := v3["hello"]
                            _ = v25
                        } else {
                            switch int(v16) {
                            case 9223372036854775807:
                                v26 := GenS3{F0: 2147483647, F1: ("hello" + ""), F2: 42}
                                _ = v26
                            case 32768:
                                fmt.Println(fi20)
                            case -1:
                                if (v16 == 255) {
                                    v27 := GenS5{F0: int32(v2)}
                                    v28 := []int{(fi20 + fi20)}
                                    v1 = 127
                                    _ = len(v3)
                                    var v29 string = ""
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            }
                            var v30 int64 = genFn1(22, 349, 1)
                            _ = v30
                        }
                        var v31 uint16 = genFn0()
                        _ = v31
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                if (100 == 49) {
                    if false {
                        fmt.Println(fi19, v1, v1)
                    } else {
                        var v32 int = 128
                        _ = v32
                    }
                    v16 += genFn0()
                    v16 = (v16 | genFn0())
                }
                v33 := []int{}
                _ = v33
            }
            var v34 uint16 = v16
            _ = fi19
            _ = v34
        }
        for fi35 := 0; fi35 < 2; fi35++ {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                fmt.Println(fi35)
            case "foo bar":
                fmt.Println(v16, fi35)
            default:
                v1 = v2
            }
            var v36 uint32 = 4294967295
            _ = fi35
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 uint64 = 684
        for fi39 := 0; fi39 < 3; fi39++ {
            switch "foo bar" {
            case "hello":
                var v40 float64 = float64(0.0)
                _ = v40
            case "foo bar":
                v41 := GenS2{F0: 916, F1: 697}
                _ = v41
            }
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    v1 = (v2 / (((^(v1))) | 1))
    fmt.Println(v2)
    delete(v3, ("" + "foo bar"))
    _ = len(v3)
    if ((1e10 >= 1e10) || (!(!(1e10 > 1.5)))) {
        fmt.Println(v1, v16)
    } else {
        for fi42 := 0; fi42 < 15; fi42++ {
            var v43 uint64 = 395
            _ = fi42
            _ = v43
        }
        var v44 string = "hello"
        switch -2147483648 {
        case 32768:
            fmt.Println(v1, v44, v44)
        case 128:
            if ((0 == -9223372036854775808) || (false && false)) {
                if false {
                    var v45 uint16 = v16
                    _ = v45
                } else {
                    var v46 float64 = 1e10
                    v47 := v3[(v44 + ((v44 + v44) + v44))]
                    v48 := GenS6{F0: true}
                    _ = v46
                    _ = v47
                    _ = v48
                }
                var v49 bool = ((false || (true || (!(3.14159 > 1e10)))) || false)
                fmt.Println(v44)
                _ = v49
            }
        case 462:
            v50 := GenS7{F0: 473}
            _ = v50
        default:
            fmt.Println(v16)
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
