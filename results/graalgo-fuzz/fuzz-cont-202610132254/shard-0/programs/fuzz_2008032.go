package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }
type GenS7 struct { F0 string; F1 int32; F2 int64 }
func genFn0() int16 {
    v6 := GenS1{F0: 964, F1: 42, F2: 738}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint = 65535
        _ = fi7
        _ = v8
    }
    var v9 int16 = 689
    fmt.Println(v9)
    _ = v6
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 9223372036854775807
    var v3 int16 = 829
    var v4 int8 = 100
    var v5 uint64 = 18446744073709551615
    var v10 uint = 878
    if false {
        fmt.Println(v2, v2, v4)
    } else {
        if true {
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS2{F0: "hello", F1: (v4 >> ((v10) & 63)), F2: (^(v4))}
                v4 += v4
                v13 := GenS3{F0: 505}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v10 = 65535
    }
    _ = len(v1)
    v14 := []int{}
    v15 := GenS4{F0: genFn0()}
    switch v2 {
    case 328:
        switch (-(v2)) {
        case -2147483648:
            _ = len(v14)
        }
    case -1:
        for fi16 := 0; fi16 < 16; fi16++ {
            v3 += genFn0()
            _ = fi16
        }
    case 2147483647:
        if (true && false) {
            if true {
                if true {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        switch "foo bar" {
                        case "hello":
                            v18 := GenS5{F0: int32(v5), F1: 18446744073709551615}
                            _ = v18
                        case "\n\t":
                            var v19 int64 = -9223372036854775808
                            _ = v19
                        }
                        var v20 uint8 = 150
                        _ = fi17
                        _ = v20
                    }
                    delete(v1, 239)
                }
                if (616 == uint16(v10)) {
                    if false {
                        fmt.Println(v3)
                        v1[2147483648] = 2
                        v21 := GenS3{F0: 217}
                        _ = v21
                    }
                    var v22 uint8 = 0
                    if (true || ((v3 & (v3 / (((^(v3))) | 1))) != (v3 * (-(v3))))) {
                        var v23 uint64 = (v5 | v5)
                        v1[583] = 1
                        _ = v23
                    }
                    _ = v22
                } else {
                    switch (v2 + -307) {
                    case -9223372036854775808:
                        if false {
                            if ((v3 + 0) >= (v3 - v3)) {
                                for fi24 := 0; fi24 < 5; fi24++ {
                                    v4 += v4
                                    v25 := GenS5{F0: 2147483647, F1: (v10 | (v10 & 4294967295))}
                                    _ = fi24
                                    _ = v25
                                }
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    v4 = v4
                                    _ = fi26
                                }
                                v27 := GenS3{F0: 448}
                                _ = v27
                            } else {
                                for fi28 := 0; fi28 < 12; fi28++ {
                                    for fi29 := 0; fi29 < 2; fi29++ {
                                        delete(v1, 2147483647)
                                        var v30 int = (fi29 << ((98) & 63))
                                        _ = fi29
                                        _ = v30
                                    }
                                    fmt.Println(v3)
                                    _ = fi28
                                }
                            }
                        } else {
                            if true {
                                var v31 int16 = (^(v3))
                                _ = v31
                            } else {
                                var v32 int32 = 1
                                _ = v32
                            }
                            v4 = 95
                        }
                    }
                    v33 := GenS5{F0: 61, F1: 0}
                    var v34 uint16 = 10
                    _ = v33
                    _ = v34
                }
            } else {
                v3 = genFn0()
            }
        }
    default:
        v35 := GenS6{F0: 1, F1: v10, F2: (v3 | v3)}
        _ = v35
    }
    v3 = (v3 + -65)
    v5 += v5
    switch (v2 % ((v2) | 1)) {
    case 531:
        switch v2 {
        case -601:
            v3 += genFn0()
        case -2147483648:
            switch (v2 >> ((v10) & 63)) {
            case -1:
                v36 := GenS6{F0: v3, F1: 444, F2: v3}
                _ = v36
            case -9223372036854775808:
                switch (^(v2)) {
                case 32768:
                    switch (("a" + ("hello" + ((((("hello" + "a") + "hello") + (("hello" + "\n\t") + "a")) + "\n\t") + ("hello" + "a")))) + ("\n\t" + ("\n\t" + (("a" + (("foo bar" + "a") + "")) + "foo bar")))) {
                    case "a":
                        switch "\n\t" {
                        case "":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        case "a":
                            var v37 uint16 = 65535
                            _ = v37
                        case "\n\t":
                            var v38 uint = v10
                            _ = v38
                        default:
                            var v39 int8 = 100
                            _ = v39
                        }
                    case "foo bar":
                        var v40 float64 = 1.5
                        _ = v40
                    case "\n\t":
                        v10 = 136
                    default:
                        v2 += (v2 ^ (v2 >> ((207) & 63)))
                    }
                case 2147483647:
                    var v41 uint8 = 243
                    _ = v41
                }
            case 936:
                if (false || false) {
                    if true {
                        fmt.Println(v3, v4)
                    } else {
                        var v42 int = (^(v2))
                        fmt.Println(v2, v42)
                        _ = v42
                    }
                    fmt.Println(v10)
                }
            default:
                for fi43 := 0; fi43 < 14; fi43++ {
                    for fi44 := 0; fi44 < 2; fi44++ {
                        var v45 int32 = 0
                        _ = fi44
                        _ = v45
                    }
                    if len(v14) > 0 {
                        v14[0] = v2
                    }
                    _ = fi43
                }
            }
        }
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 uint = v10
        v4 -= v4
        _ = fi46
        _ = v47
    }
    if false {
        v2 -= v2
        if ("foo bar" == "hello") {
            var v48 int8 = 19
            v5 += 42
            switch v2 {
            case 560:
                if false {
                    switch ("" + ("" + "\n\t")) {
                    case "\n\t":
                        var v49 int16 = 32767
                        _ = v49
                    case "":
                        v5 = 226
                    case "foo bar":
                        v2 = (^(v2))
                    default:
                        var v50 int64 = -441
                        _ = v50
                    }
                    var v51 float64 = (1e10 - (1e10 * ((0.0 * float64((((1.5 - float32(1e10)) + ((-2.25 * 3.14159) + 1.5)) - 3.14159))) * 3.14159)))
                    switch (^(v2)) {
                    case 187:
                        for fi52 := 0; fi52 < 7; fi52++ {
                            switch 604 {
                            case 32768:
                                v2 = v2
                            default:
                                var v53 uint64 = (v5 % ((v5) | 1))
                                _ = v53
                            }
                            v3 += genFn0()
                            fmt.Println(v2, v4, v51)
                            _ = fi52
                        }
                    case 697:
                        var v54 uint32 = 42
                        _ = v54
                    case 128:
                        v5 += (^(v5))
                    default:
                        v4 = -128
                    }
                    _ = v51
                }
            case -977:
                fmt.Println(v10)
            default:
                switch v2 {
                case 4294967295:
                    _ = len(v14)
                default:
                    v2 += v2
                }
            }
            _ = v48
        } else {
            v14 = append(v14, v2)
        }
    }
    v55 := GenS7{F0: "", F1: 122, F2: 0}
    v56 := GenS1{F0: 2147483648, F1: v2, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
}
