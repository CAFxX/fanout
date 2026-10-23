package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 string; F1 uint8 }
func genFn0(p0 uint8) int32 {
    p0 = (p0 & (^(p0)))
    v5 := GenS1{F0: -128}
    p0 = (p0 + (p0 % ((p0) | 1)))
    if true {
        switch "hello" {
        case "":
            for fi6 := 0; fi6 < 15; fi6++ {
                for fi7 := 0; fi7 < 18; fi7++ {
                    fmt.Println(fi7)
                    p0 = (p0 >> ((255) & 63))
                    _ = fi7
                }
                v8 := []int{(fi6 >> ((1) & 63)), fi6}
                var v9 string = "hello"
                _ = fi6
                _ = v8
                _ = v9
            }
        default:
            var v10 int64 = 0
            _ = v10
        }
        switch ("foo bar" + "hello") {
        case "":
            p0 = p0
        default:
            for fi11 := 0; fi11 < 12; fi11++ {
                if true {
                    v12 := GenS2{F0: -677, F1: 0, F2: int32(fi11)}
                    for fi13 := 0; fi13 < 8; fi13++ {
                        if false {
                            if false {
                                var v14 int = fi13
                                v15 := []int{}
                                v16 := GenS1{F0: -128}
                                v14 = 0
                                _ = v14
                                _ = v15
                                _ = v16
                            } else {
                                switch "" {
                                case "\n\t":
                                    p0 = p0
                                }
                                if true {
                                    v17 := GenS3{F0: 42}
                                    if ((("foo bar" + ("foo bar" + "")) + "foo bar") >= "hello") {
                                        if false {
                                            p0 = (p0 / ((p0) | 1))
                                        }
                                        var v18 uint = 1
                                        switch fi13 {
                                        case 32767:
                                            v19 := GenS4{F0: 1, F1: 42}
                                            _ = v19
                                        case 873:
                                            v20 := GenS4{F0: (fi13 >> (((^(v18))) & 63)), F1: 276}
                                            _ = v20
                                        case 2147483647:
                                            if (!(-258 < 1000)) {
                                                var v21 int = 9223372036854775807
                                                var v22 uint = (v18 - (v18 / ((v18) | 1)))
                                                if ((^(p0)) >= p0) {
                                                    switch ("hello" + "") {
                                                    case "foo bar":
                                                        v23 := GenS1{F0: 0}
                                                        _ = v23
                                                    case "hello":
                                                        v24 := []int{fi11, 2147483647, -1, (fi11 ^ (fi11 << ((65535) & 63)))}
                                                        _ = v24
                                                    case "a":
                                                        var v25 int16 = -1
                                                        _ = v25
                                                    }
                                                    var v26 uint8 = p0
                                                    _ = v26
                                                }
                                                _ = v21
                                                _ = v22
                                            }
                                            var v27 float32 = 0.1
                                            _ = v27
                                        }
                                        _ = v18
                                    } else {
                                        v28 := []int{-1, (fi13 % ((9223372036854775807) | 1)), fi13, (fi13 + 0)}
                                        _ = v28
                                    }
                                    _ = v17
                                }
                            }
                        }
                        p0 = p0
                        _ = fi13
                    }
                    _ = v12
                }
                switch fi11 {
                case -133:
                    if (1 <= 42) {
                        var v29 int16 = 32767
                        fmt.Println(fi11, fi11)
                        _ = v29
                    }
                case 32767:
                    for fi30 := 0; fi30 < 1; fi30++ {
                        var v31 bool = (!(false || (0 != 0)))
                        _ = fi30
                        _ = v31
                    }
                case 0:
                    v32 := GenS5{F0: 67, F1: p0}
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if (42 > 1) {
                            fmt.Println(p0, p0)
                            fmt.Println(p0, fi33, fi11)
                            if false {
                                var v34 int32 = 42
                                switch (("a" + "a") + ("a" + "foo bar")) {
                                case "a":
                                    v34 += v34
                                case "foo bar":
                                    fmt.Println(p0, fi33, fi33)
                                case "":
                                    switch "" {
                                    case "":
                                        switch (fi11 / ((fi33) | 1)) {
                                        case 519:
                                            if false {
                                                var v35 bool = false
                                                v34 -= (v34 / ((v34) | 1))
                                                v35 = false
                                                _ = v35
                                            }
                                        case -2147483648:
                                            p0 += 0
                                        default:
                                            switch (fi33 & -656) {
                                            case -410:
                                                p0 = (^(p0))
                                            }
                                        }
                                        v36 := []int{-343, fi33, fi33, (fi33 - -780)}
                                        if (false || (729 != 9223372036854775807)) {
                                            if false {
                                                p0 += (p0 >> ((4294967295) & 63))
                                            } else {
                                                if (!true) {
                                                    var v37 float32 = -2.25
                                                    v38 := GenS1{F0: -32}
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    var v39 uint16 = 621
                                                    v39 -= 255
                                                    switch ("" + "a") {
                                                    case "hello":
                                                        fmt.Println(fi33, p0)
                                                    case "\n\t":
                                                        v40 := GenS6{F0: ("a" + ""), F1: (p0 ^ 1)}
                                                        _ = v40
                                                    default:
                                                        _ = len(v36)
                                                    }
                                                    var v41 string = ""
                                                    if len(v36) >= 2 {
                                                        v36 = v36[1:len(v36)]
                                                    }
                                                    _ = v39
                                                    _ = v41
                                                }
                                            }
                                            v42 := GenS4{F0: fi33, F1: 0}
                                            fmt.Println(fi11)
                                            _ = v42
                                        }
                                        _ = v36
                                    case "a":
                                        fmt.Println(p0, v34, fi33)
                                    default:
                                        var v43 string = "\n\t"
                                        _ = v43
                                    }
                                default:
                                    p0 = (p0 * p0)
                                }
                                v34 += v34
                                var v44 int64 = -1
                                v44 = v44
                                _ = v34
                                _ = v44
                            }
                            v45 := []int{fi33}
                            _ = v45
                        }
                        p0 += p0
                        if (false && true) {
                            switch (fi11 ^ 245) {
                            case 128:
                                p0 = p0
                            }
                            p0 = (p0 << ((4294967295) & 63))
                            v46 := []int{(fi33 | (^(fi33))), fi33, fi11, fi11}
                            _ = v46
                        } else {
                            var v47 int16 = -1
                            _ = v47
                        }
                        _ = fi33
                    }
                }
                _ = fi11
            }
        }
        p0 = (p0 * (^(p0)))
    } else {
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 uint8 = 192
            var v50 float32 = 1.5
            v50 = 1e10
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    var v51 int32 = 2147483647
    _ = p0
    _ = v5
    _ = v51
    return 0
}

func main() {
    var v1 string = "a"
    var v2 int32 = 1
    var v3 float32 = -2.25
    var v4 int64 = -9223372036854775808
    v4 -= (v4 + (v4 >> ((42) & 63)))
    v4 = v4
    v3 -= float32(1e10)
    for fi52 := 0; fi52 < 2; fi52++ {
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 int32 = (^(v2))
            _ = fi53
            _ = v54
        }
        fmt.Println(v1, v1, fi52)
        _ = fi52
    }
    v55 := []int{0, 4294967295, 127, 1}
    v56 := GenS2{F0: -1, F1: 979, F2: 970}
    var v57 uint16 = 423
    var v58 uint16 = (v57 | 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
