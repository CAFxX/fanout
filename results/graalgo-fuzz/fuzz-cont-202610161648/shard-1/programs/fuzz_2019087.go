package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 int32 }
type GenS3 struct { F0 int8; F1 int; F2 string }
type GenS4 struct { F0 bool }
func genFn0(p0 int32, p1 uint8, p2 int8) int {
    p2 += p2
    if true {
        v7 := []int{0, 32768, 431, 32767}
        for fi8 := 0; fi8 < 12; fi8++ {
            if (p1 != 255) {
                v9 := GenS1{F0: 255, F1: -9223372036854775808, F2: 4294967295}
                var v10 string = "foo bar"
                _ = v10
                _ = v9
            }
            p2 += p2
            p2 = p2
            _ = fi8
        }
        _ = v7
    }
    v11 := []int{2147483647, -796, 2147483647}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint = 4294967295
        var v14 bool = false
        _ = fi12
        _ = v13
        _ = v14
    }
    p1 += 128
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := []string{("foo bar" + ("" + "hello"))}
    v2 := make(map[int]string)
    var v3 int64 = 0
    var v4 float32 = 1e10
    var v5 int16 = 1
    var v6 string = ""
    v1 = append(v1, ("a" + ""))
    if (v5 != 486) {
        fmt.Println(v4, v6, v3)
    }
    if (v6 != (v6 + (("\n\t" + "\n\t") + ("a" + ((v6 + v6) + ("" + "a")))))) {
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS2{F0: (^(fi15)), F1: 36, F2: -217}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v17 bool = ((!(842 != -902)) || (!(1 == 2147483648)))
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := GenS2{F0: 32768, F1: 234, F2: 42}
        _ = v18
    }
    v19 := GenS1{F0: 399, F1: (^(v3)), F2: 0}
    v1 = append(v1, v6)
    for fi20 := 0; fi20 < 17; fi20++ {
        if false {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 uint8 = 127
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: 0, F1: (v3 & -386), F2: 42}
            _ = v23
        }
        var v24 uint64 = 18446744073709551615
        _ = fi20
        _ = v24
    }
    var v25 uint32 = 65535
    _ = len(v2)
    var v26 int32 = 42
    v27 := GenS1{F0: 42, F1: (v3 >> ((18446744073709551615) & 63)), F2: 4294967295}
    v28 := GenS1{F0: uint16(v25), F1: v3, F2: 2}
    if true {
        fmt.Println(v25, v25)
        var v29 int = 128
        _ = v29
    } else {
        fmt.Println(v6)
        if ((!(103 >= 200)) || true) {
            var v30 uint64 = 761
            _ = v30
        } else {
            _ = len(v1)
            if ((genFn0(v26, 128, 127) == 32767) || true) {
                fmt.Println(v6, v26, v3)
                fmt.Println(v4, v25)
            } else {
                for fi31 := 0; fi31 < 10; fi31++ {
                    v26 -= -1
                    _ = len(v2)
                    switch (-(fi31)) {
                    case -9223372036854775808:
                        fmt.Println(v5, v25)
                    case -404:
                        if ((true || (938 <= 644)) && true) {
                            v32 := GenS1{F0: 32767, F1: 698, F2: 273}
                            switch (fi31 << ((4294967295) & 63)) {
                            case 127:
                                var v33 float64 = 0.1
                                _ = v33
                            }
                            _ = v32
                        } else {
                            if true {
                                v34 := GenS3{F0: 0, F1: (fi31 >> ((2) & 63)), F2: "hello"}
                                _ = v34
                            } else {
                                switch (fi31 & fi31) {
                                case 9223372036854775807:
                                    v35 := v2[fi31]
                                    _ = v35
                                case -2147483648:
                                    var v36 float32 = 0.1
                                    _ = v36
                                }
                                for fi37 := 0; fi37 < 5; fi37++ {
                                    v26 = v26
                                    v3 += (v3 >> ((4294967295) & 63))
                                    _ = fi37
                                }
                            }
                            var v38 float64 = 0.0
                            v4 = v4
                            _ = v38
                        }
                    case 4294967295:
                        delete(v2, genFn0((v26 | -2147483648), 3, -78))
                    }
                    _ = fi31
                }
            }
        }
    }
    v4 = (((v4 + -2.25) - v4) / ((((v4 / 3.14159) + (v4 - v4)) + (-2.25 - -2.25)) / 3.14159))
    switch v6 {
    case "hello":
        fmt.Println(v6, v3, v6)
    case "\n\t":
        v6 = ""
    default:
        var v39 int16 = (v5 - (v5 << ((541) & 63)))
        _ = v39
    }
    if false {
        var v40 uint16 = uint16(v25)
        _ = v40
    } else {
        var v41 float32 = 0.0
        for fi42 := 0; fi42 < 20; fi42++ {
            var v43 uint64 = 42
            v3 -= 718
            var v44 float32 = (float32(-2.25) / (v4 * (-2.25 + 3.14159)))
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v41
    }
    if len(v1) > 0 {
        v1[0] = (((v6 + "") + v6) + v6)
    }
    if true {
        for fi45 := 0; fi45 < 9; fi45++ {
            v2[(-(fi45))] = (v6 + (v6 + (v6 + v6)))
            for fi46 := 0; fi46 < 8; fi46++ {
                fmt.Println(v26)
                fmt.Println(fi45, v6)
                var v47 bool = false
                _ = fi46
                _ = v47
            }
            var v48 int64 = 258
            _ = fi45
            _ = v48
        }
        v49 := GenS1{F0: 1, F1: (v3 + (v3 + -505)), F2: 18}
        _ = v49
    } else {
        v50 := GenS4{F0: (((v6 + "hello") <= v6) || ((!true) && (!(65535 < 65535))))}
        if true {
            var v51 uint32 = (v25 % ((v25) | 1))
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            var v52 int16 = v5
            _ = v51
            _ = v52
        }
        var v53 string = v6
        _ = v50
        _ = v53
    }
    var v54 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
