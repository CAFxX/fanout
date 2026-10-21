package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint8 }
func genFn0(p0 int16, p1 string, p2 uint64) uint {
    var v6 string = p1
    var v7 uint16 = 2
    v6 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1() uint16 {
    var v8 bool = (0 == -9223372036854775808)
    for fi9 := 0; fi9 < 18; fi9++ {
        v8 = false
        _ = fi9
    }
    v10 := []int{-2147483648, -1, 127, 635}
    if ((1 >= 0) && true) {
        v11 := GenS1{F0: 0}
        v12 := GenS1{F0: 9223372036854775807}
        _ = v11
        _ = v12
    } else {
        v8 = true
        v8 = (v8 || (2147483648 <= 42))
        var v13 string = (((((("\n\t" + "a") + "") + "") + ((("\n\t" + (("hello" + "hello") + "hello")) + "\n\t") + ("hello" + "foo bar"))) + "hello") + ("\n\t" + "a"))
        _ = v13
    }
    v8 = v8
    _ = v10
    _ = v8
    return 85
}

func main() {
    var v1 int = 4294967295
    var v2 bool = false
    v3 := []uint32{42, 356, 2, 2147483648}
    v4 := make(map[string]uint32)
    v5 := []uint32{2, 2147483647, 65535, 65535}
    v14 := v4[((("" + "a") + "a") + "a")]
    v1 -= v1
    if len(v3) > 0 {
        v3[0] = (v14 ^ v14)
    }
    var v15 int8 = -128
    switch v1 {
    case -9223372036854775808:
        if v2 {
            if (-9223372036854775808 < -9223372036854775808) {
                fmt.Println(v14, v2, v2)
                v16 := GenS2{F0: (!true), F1: 0, F2: 4294967295}
                var v17 int16 = 42
                _ = v16
                _ = v17
            }
            v2 = false
        } else {
            var v18 float32 = 0.1
            v19 := GenS1{F0: 1}
            _ = v18
            _ = v19
        }
    case -1:
        v14 -= 4294967295
    default:
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                if v2 {
                    fmt.Println(v1)
                } else {
                    v22 := GenS2{F0: true, F1: 0, F2: 2}
                    if (3.14159 < 0.1) {
                        v23 := GenS3{F0: 42}
                        _ = v23
                    }
                    var v24 uint = 4294967295
                    _ = v22
                    _ = v24
                }
                var v25 float32 = 0.1
                v2 = (false && false)
                _ = fi21
                _ = v25
            }
            var v26 int = v1
            _ = fi20
            _ = v26
        }
    }
    v1 -= 42
    for fi27 := 0; fi27 < 6; fi27++ {
        fmt.Println(v1, v15, v15)
        v15 += 42
        _ = fi27
    }
    v1 = (v1 % ((v1) | 1))
    if (!v2) {
        v1 = int(v15)
        v28 := GenS4{F0: -32768, F1: 128}
        _ = v28
    } else {
        for fi29 := 0; fi29 < 17; fi29++ {
            fmt.Println(v2, v15)
            fmt.Println(v2)
            switch 32768 {
            case 9223372036854775807:
                v14 = 2147483648
            case -1:
                if v2 {
                    switch (v1 % ((-2147483648) | 1)) {
                    case 2147483647:
                        v15 = v15
                    }
                    var v30 int64 = 0
                    if (1e10 >= 1e10) {
                        if false {
                            var v31 uint64 = 0
                            _ = v31
                        } else {
                            var v32 string = ""
                            var v33 string = (("\n\t" + v32) + (("hello" + v32) + v32))
                            fmt.Println(v30, v14, v33)
                            _ = v32
                            _ = v33
                        }
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            if v2 {
                                var v34 uint8 = 255
                                _ = len(v5)
                                var v35 int64 = (v30 / ((v30) | 1))
                                _ = v34
                                _ = v35
                            }
                        case "a":
                            if v2 {
                                for fi36 := 0; fi36 < 10; fi36++ {
                                    if v2 {
                                        v2 = true
                                        var v37 uint64 = 656
                                        _ = v37
                                    } else {
                                        fmt.Println(v2, v14, v30)
                                        if true {
                                            if true {
                                                v38 := GenS3{F0: -442}
                                                _ = v38
                                            } else {
                                                var v39 float32 = (3.14159 - -2.25)
                                                var v40 int = (fi29 % ((128) | 1))
                                                v1 += fi36
                                                _ = v39
                                                _ = v40
                                            }
                                            v30 -= 42
                                        } else {
                                            v41 := GenS5{F0: (^(v14)), F1: 42, F2: 30}
                                            fmt.Println(fi36, v30)
                                            if v2 {
                                                v4["a"] = (v14 * (v14 ^ v14))
                                                v42 := v4[("" + ("\n\t" + ""))]
                                                var v43 uint32 = v42
                                                _ = v42
                                                _ = v43
                                            }
                                            _ = v41
                                        }
                                        var v44 int8 = (v15 << ((0) & 63))
                                        _ = v44
                                    }
                                    _ = fi36
                                }
                            }
                        case "":
                            var v45 int32 = 166
                            _ = v45
                        }
                        fmt.Println(v2)
                        v2 = (v2 || (0 >= 0))
                    }
                    _ = v30
                } else {
                    var v46 uint64 = 42
                    _ = v46
                }
            case 127:
                v14 = v14
            }
            _ = fi29
        }
        var v47 uint32 = (v14 << ((0) & 63))
        for fi48 := 0; fi48 < 18; fi48++ {
            var v49 uint = 18446744073709551615
            _ = fi48
            _ = v49
        }
        _ = v47
    }
    _ = len(v5)
    switch "hello" {
    case "":
        switch (v1 - v1) {
        case 32767:
            v50 := GenS3{F0: 411}
            _ = v50
        }
    default:
        for fi51 := 0; fi51 < 8; fi51++ {
            var v52 int64 = -9223372036854775808
            _ = fi51
            _ = v52
        }
    }
    delete(v4, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
