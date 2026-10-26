package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0(p0 int8) int32 {
    if false {
        p0 = 0
        p0 += p0
    }
    v7 := GenS1{F0: 431, F1: "a"}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 int) uint32 {
    v8 := []int{p0, (-(p0)), p0}
    for fi9 := 0; fi9 < 16; fi9++ {
        switch (p0 | fi9) {
        case 1:
            v8 = append(v8, fi9)
        default:
            fmt.Println(p0)
        }
        _ = fi9
    }
    v8 = append(v8, p0)
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 4294967295
}

func main() {
    v1 := []string{"\n\t", ("" + ("\n\t" + "foo bar")), "foo bar", (("\n\t" + "foo bar") + "foo bar")}
    var v2 bool = (42 <= 219)
    v3 := []string{""}
    v4 := make(map[string]string)
    var v5 int8 = 100
    v6 := make(map[string]string)
    v10 := GenS2{F0: 255, F1: 45, F2: -2147483648}
    var v11 uint = 520
    v5 += v5
    v12 := GenS1{F0: 265, F1: ""}
    v13 := GenS2{F0: 32767, F1: 42, F2: 42}
    switch 340 {
    case 1:
        fmt.Println(v5)
    case -976:
        if (true && v2) {
            v5 = -73
        } else {
            if (21 != (v11 >> (((v11 | 245)) & 63))) {
                v14 := GenS1{F0: 0, F1: "\n\t"}
                _ = v14
            }
            fmt.Println(v11, v5, v11)
            v11 += v11
        }
    case -592:
        if false {
            v15 := GenS2{F0: 65535, F1: 679, F2: 1}
            _ = len(v1)
            _ = v15
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if v2 {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v5 += 1
                    _ = len(v1)
                    switch fi16 {
                    case 1:
                        var v17 uint = (v11 | (v11 >> ((243) & 63)))
                        _ = v17
                    case -199:
                        v18 := GenS2{F0: 32767, F1: 259, F2: 2147483647}
                        _ = v18
                    case 127:
                        if (v5 <= 127) {
                            delete(v6, "a")
                            if false {
                                var v19 uint16 = 2
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "\n\t")
                                    }
                                    _ = fi20
                                }
                                _ = v19
                            }
                            if v2 {
                                var v21 uint = v11
                                _ = v21
                            }
                        } else {
                            v22 := GenS2{F0: 4, F1: 4294967295, F2: 1}
                            for fi23 := 0; fi23 < 17; fi23++ {
                                switch "foo bar" {
                                case "foo bar":
                                    if false {
                                        fmt.Println(fi23, v5, v2)
                                        var v24 string = "\n\t"
                                        var v25 string = "hello"
                                        v26 := GenS2{F0: 65535, F1: 54, F2: -58}
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    if (!v2) {
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = ((("hello" + (("" + "a") + "foo bar")) + ("\n\t" + "hello")) + ("hello" + ""))
                                    }
                                    switch fi23 {
                                    case 42:
                                        fmt.Println(fi23)
                                    case 128:
                                        var v27 float32 = 0.0
                                        _ = v27
                                    }
                                    v2 = (603 != 18446744073709551615)
                                case "\n\t":
                                    if (((fi23 * fi16) > fi23) || (2147483648 <= 209)) {
                                        fmt.Println(fi23, v5, v2)
                                        v28 := GenS1{F0: 18446744073709551615, F1: ""}
                                        v11 = (v11 & v11)
                                        _ = v28
                                    }
                                case "":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    var v29 uint32 = 753
                                    _ = v29
                                }
                                _ = fi23
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v22
                        }
                    }
                    _ = fi16
                }
                var v30 uint = (v11 ^ (v11 - (v11 * 732)))
                _ = v30
            } else {
                for fi31 := 0; fi31 < 3; fi31++ {
                    switch (fi31 >> (((v11 | v11)) & 63)) {
                    case -9223372036854775808:
                        v5 += (v5 ^ (v5 >> ((v11) & 63)))
                    case 989:
                        v32 := GenS3{F0: 4294967295, F1: 172, F2: v11}
                        _ = v32
                    case 32767:
                        fmt.Println(v11, v2, v2)
                    default:
                        var v33 float32 = 0.1
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 4; fi34++ {
                        if true {
                            v35 := GenS4{F0: 1}
                            _ = v35
                        }
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v2, v11)
                        case "\n\t":
                            if false {
                                _ = len(v6)
                                v11 = (^(v11))
                                if v2 {
                                    var v36 uint32 = 65535
                                    _ = v36
                                }
                            } else {
                                _ = len(v6)
                            }
                        }
                        _ = fi34
                    }
                    var v37 string = ("\n\t" + "foo bar")
                    _ = fi31
                    _ = v37
                }
                v38 := GenS1{F0: 498, F1: ("hello" + "\n\t")}
                _ = v38
            }
            v2 = (!(0 <= 1))
        }
    default:
        var v39 uint8 = 226
        _ = v39
    }
    fmt.Println(v5)
    if ((!(!v2)) && (v2 || v2)) {
        v2 = (-32768 <= 0)
    } else {
        v40 := GenS4{F0: -332}
        var v41 int16 = 345
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        fmt.Println(v5, v11, fi42)
        if (287 != genFn1((fi42 / (((^(fi42))) | 1)))) {
            var v43 string = "\n\t"
            v44 := GenS5{F0: 130, F1: v2}
            _ = v43
            _ = v44
        }
        v2 = v2
        _ = fi42
    }
    var v45 int16 = 32767
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := GenS4{F0: 503}
        var v48 int64 = -1
        var v49 int16 = 1
        _ = fi46
        _ = v47
        _ = v48
        _ = v49
    }
    var v50 int64 = 1
    _ = len(v3)
    if (42 == v45) {
        var v51 bool = (2147483647 == genFn0(v5))
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v50
    _ = v6
}
