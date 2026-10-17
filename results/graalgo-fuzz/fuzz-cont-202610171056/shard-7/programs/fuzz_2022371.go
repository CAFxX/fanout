package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int8 }
type GenS5 struct { F0 int; F1 uint32; F2 int8 }
type GenS6 struct { F0 int32; F1 int16 }
func genFn0(p0 int64, p1 uint64, p2 string) uint {
    p1 -= (p1 + p1)
    var v6 bool = (20 < 100)
    fmt.Println(p0, p2)
    var v7 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 699
}

func main() {
    var v1 bool = (18446744073709551615 >= 0)
    var v2 uint = 656
    var v3 string = "a"
    v4 := []int{-9223372036854775808}
    var v5 uint32 = 706
    for fi8 := 0; fi8 < 6; fi8++ {
        if v1 {
            if (31 >= 0) {
                v5 += uint32(v2)
            } else {
                for fi9 := 0; fi9 < 16; fi9++ {
                    switch (fi9 - 9223372036854775807) {
                    case 2147483647:
                        _ = len(v4)
                    default:
                        var v10 uint8 = 183
                        _ = v10
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi9
                }
            }
        }
        var v11 uint8 = 42
        _ = fi8
        _ = v11
    }
    if v1 {
        v12 := GenS1{F0: 0, F1: 84}
        _ = v12
    }
    var v13 int64 = -9223372036854775808
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS2{F0: (-(v13)), F1: 927}
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS3{F0: 245, F1: 1}
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    switch ("" + ("a" + "\n\t")) {
    case "":
        if true {
            if false {
                v13 = 1
                if ((!v1) || true) {
                    var v18 uint32 = 309
                    _ = v18
                } else {
                    _ = len(v4)
                    v3 = ((("" + v3) + "") + v3)
                    var v19 int = 908
                    _ = v19
                }
            } else {
                for fi20 := 0; fi20 < 16; fi20++ {
                    for fi21 := 0; fi21 < 7; fi21++ {
                        if (!v1) {
                            var v22 bool = (false || ((!v1) && (((("hello" + v3) + "\n\t") + (((v3 + v3) + v3) + v3)) == (v3 + ("hello" + v3)))))
                            if (v1 && (0 > 161)) {
                                var v23 int64 = v13
                                var v24 float32 = 0.1
                                var v25 float32 = v24
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            var v26 int16 = 1
                            _ = v22
                            _ = v26
                        } else {
                            v27 := GenS1{F0: 389, F1: 197}
                            if v1 {
                                v3 = "\n\t"
                            } else {
                                var v28 string = "\n\t"
                                _ = v28
                            }
                            if v1 {
                                v5 = 2
                                var v29 float32 = 3.14159
                                v1 = true
                                _ = v29
                            }
                            _ = v27
                        }
                        v3 = (v3 + (("foo bar" + v3) + "a"))
                        var v30 bool = v1
                        _ = fi21
                        _ = v30
                    }
                    _ = fi20
                }
                var v31 uint16 = 2
                _ = v31
            }
            v4 = append(v4, 2147483647)
            fmt.Println(v5, v2)
        }
    case "a":
        for fi32 := 0; fi32 < 8; fi32++ {
            fmt.Println(v5, v5, v3)
            if len(v4) > 0 {
                v4[0] = 74
            }
            _ = fi32
        }
    default:
        fmt.Println(v13)
    }
    v1 = (true || v1)
    switch 42 {
    case 1:
        v3 = (((("hello" + v3) + "") + "hello") + "hello")
    default:
        var v33 uint = v2
        _ = v33
    }
    var v34 string = v3
    var v35 string = ("\n\t" + "")
    for fi36 := 0; fi36 < 14; fi36++ {
        var v37 int16 = 1000
        v38 := GenS1{F0: (-(v37)), F1: 65}
        var v39 uint64 = 1
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        fmt.Println(v35, v35, v13)
        v41 := GenS3{F0: 4294967295, F1: 630}
        if v1 {
            v42 := GenS4{F0: 32767, F1: (v13 ^ v13), F2: -128}
            v43 := GenS5{F0: fi40, F1: v5, F2: 42}
            var v44 uint16 = 42
            _ = v42
            _ = v43
            _ = v44
        }
        _ = fi40
        _ = v41
    }
    switch "\n\t" {
    case "a":
        v2 += v2
    case "hello":
        var v45 uint8 = 2
        _ = v45
    }
    if (0.0 != float32(1e10)) {
        var v46 uint32 = (v5 + (v5 >> ((58) & 63)))
        if true {
            v3 = ""
        } else {
            var v47 uint8 = 100
            _ = v47
        }
        _ = v46
    } else {
        v4 = append(v4, int(v5))
    }
    for fi48 := 0; fi48 < 17; fi48++ {
        var v49 uint64 = 2
        v5 = v5
        _ = fi48
        _ = v49
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi50 := 0; fi50 < 19; fi50++ {
        v51 := GenS6{F0: 42, F1: 32767}
        var v52 uint32 = 967
        _ = fi50
        _ = v51
        _ = v52
    }
    var v53 int = 32768
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v53
}
