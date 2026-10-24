package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int, p1 int8, p2 uint16) uint16 {
    v6 := GenS1{F0: 89, F1: 18446744073709551615}
    var v7 int32 = -1
    for fi8 := 0; fi8 < 6; fi8++ {
        if (false && false) {
            var v9 int32 = (^(v7))
            _ = v9
        } else {
            v10 := []int{p0}
            if (!false) {
                var v11 string = "hello"
                for fi12 := 0; fi12 < 19; fi12++ {
                    if (!true) {
                        var v13 int8 = (p1 * 100)
                        _ = v13
                    }
                    var v14 uint64 = 4294967295
                    fmt.Println(p1, v11, fi8)
                    if (((v14 % ((v14) | 1)) > v14) || true) {
                        v15 := GenS1{F0: -44, F1: (v14 % (((v14 * uint64(p2))) | 1))}
                        _ = v15
                    } else {
                        p2 = (p2 * 772)
                    }
                    _ = len(v10)
                    _ = fi12
                    _ = v14
                }
                _ = v11
            } else {
                p1 += (p1 >> ((42) & 63))
            }
            _ = v10
        }
        _ = fi8
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        switch fi16 {
        case 2147483647:
            v17 := []int{}
            _ = v17
        default:
            v18 := []int{int(v7), fi16, (p0 << ((42) & 63))}
            _ = v18
        }
        var v19 uint32 = 2147483647
        _ = fi16
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    v1 := []string{("hello" + "foo bar"), ((("\n\t" + ("hello" + (("\n\t" + (("" + "a") + "a")) + "a"))) + "a") + "a"), ((((((("" + (("a" + "") + "hello")) + "hello") + ("foo bar" + "foo bar")) + ("" + ("foo bar" + ("" + "a")))) + "\n\t") + "") + "a")}
    var v2 float32 = 0.1
    var v3 int16 = -1
    v4 := []int{0, 574, -2147483648, -519}
    v5 := make(map[string]uint32)
    switch "a" {
    case "foo bar":
        delete(v5, "")
    case "a":
        fmt.Println(v3, v2)
    default:
        v20 := GenS1{F0: 70, F1: 42}
        _ = v20
    }
    switch ("hello" + (("foo bar" + "") + "\n\t")) {
    case "\n\t":
        v2 += v2
    case "a":
        fmt.Println(v3)
    default:
        for fi21 := 0; fi21 < 1; fi21++ {
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 int = fi22
                _ = fi22
                _ = v23
            }
            fmt.Println(v3, v3)
            _ = fi21
        }
    }
    v3 = v3
    switch -168 {
    case 32767:
        _ = len(v1)
    case 672:
        v24 := v5["foo bar"]
        _ = v24
    case 4294967295:
        if (128 != 32767) {
            for fi25 := 0; fi25 < 5; fi25++ {
                v3 += (v3 % (((v3 & (v3 << ((453) & 63)))) | 1))
                v2 = float32(1.5)
                _ = fi25
            }
        } else {
            v3 = (v3 ^ v3)
            if false {
                _ = len(v1)
                v2 += (float32(0.0) - v2)
            }
        }
    default:
        switch "a" {
        case "":
            var v26 int16 = (v3 << ((2) & 63))
            _ = v26
        case "a":
            if (!(((1.5 * ((v2 + v2) * (v2 + 1.5))) - v2) < v2)) {
                v27 := GenS2{F0: 255, F1: 44, F2: -9223372036854775808}
                if true {
                    if ("" >= ("foo bar" + ("a" + ""))) {
                        if false {
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                        }
                    }
                    var v28 uint32 = 2147483648
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if (v2 == 1e10) {
                            if false {
                                var v30 int32 = 1
                                v31 := GenS2{F0: 0, F1: 2, F2: 9223372036854775807}
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    delete(v5, ("" + "foo bar"))
                                    v2 += (0.1 - 3.14159)
                                    _ = fi32
                                }
                                _ = v30
                                _ = v31
                            } else {
                                v2 = -2.25
                                v33 := v5["foo bar"]
                                var v34 uint = 170
                                _ = v33
                                _ = v34
                            }
                        }
                        fmt.Println(v2, fi29, v3)
                        fmt.Println(v3, v3, v2)
                        _ = fi29
                    }
                    fmt.Println(v3, v3)
                }
                for fi35 := 0; fi35 < 4; fi35++ {
                    _ = len(v4)
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v3)
                    default:
                        var v36 string = "foo bar"
                        _ = v36
                    }
                    if true {
                        fmt.Println(v3)
                        v2 -= float32(1.5)
                    } else {
                        v37 := GenS3{F0: 4294967295, F1: -9223372036854775808}
                        _ = v37
                    }
                    _ = fi35
                }
                _ = v27
            } else {
                _ = len(v4)
            }
        case "hello":
            switch 475 {
            case 2147483647:
                v3 += v3
            }
        }
    }
    if (326 == 65535) {
        _ = len(v1)
        if true {
            var v38 uint64 = 316
            _ = v38
        }
        if ((false && true) && (-1 >= 42)) {
            v39 := v5[("" + ("hello" + "a"))]
            v40 := GenS4{F0: 4294967295}
            _ = v39
            _ = v40
        } else {
            fmt.Println(v3)
            var v41 int64 = -9223372036854775808
            fmt.Println(v2, v3)
            _ = v41
        }
    }
    v42 := GenS4{F0: 0}
    var v43 uint32 = 0
    switch 127 {
    case -899:
        v43 += 998
    case 4294967295:
        fmt.Println(v2, v43)
    case 9223372036854775807:
        fmt.Println(v2, v2, v2)
    }
    v2 += -2.25
    fmt.Println(v3, v3)
    fmt.Println(v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
