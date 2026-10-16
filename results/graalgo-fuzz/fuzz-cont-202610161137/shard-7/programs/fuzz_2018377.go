package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint; F1 int64 }
func genFn0() int {
    if (23 == 1) {
        if (1 <= 127) {
            v5 := GenS1{F0: 131}
            var v6 int16 = 947
            _ = v5
            _ = v6
        } else {
            v7 := GenS2{F0: -32768, F1: 42, F2: 1}
            if false {
                var v8 int16 = 32767
                switch ((("foo bar" + ((("" + ("foo bar" + "")) + "\n\t") + ("a" + "a"))) + "foo bar") + "hello") {
                case "\n\t":
                    switch 4294967295 {
                    case 42:
                        for fi9 := 0; fi9 < 12; fi9++ {
                            if true {
                                v10 := GenS3{F0: 1000, F1: 2147483647}
                                v11 := GenS2{F0: v8, F1: 2, F2: v8}
                                _ = v10
                                _ = v11
                            } else {
                                var v12 string = "hello"
                                _ = v12
                            }
                            v8 += 234
                            _ = fi9
                        }
                    case 32767:
                        v13 := []int{-200, 4294967295, -471, 2147483647}
                        _ = v13
                    case 127:
                        var v14 uint16 = 0
                        _ = v14
                    default:
                        var v15 uint64 = 911
                        _ = v15
                    }
                case "foo bar":
                    fmt.Println(v8)
                case "a":
                    if (true && (!false)) {
                        for fi16 := 0; fi16 < 19; fi16++ {
                            v8 -= (v8 - v8)
                            var v17 bool = true
                            _ = fi16
                            _ = v17
                        }
                    }
                default:
                    v18 := GenS3{F0: v8, F1: 2}
                    _ = v18
                }
                fmt.Println(v8)
                _ = v8
            }
            _ = v7
        }
        if (!((1 > 280) || (9223372036854775807 == 135))) {
            var v19 int = -266
            if (0 > 255) {
                var v20 bool = false
                _ = v20
            }
            _ = v19
        } else {
            if (true && false) {
                v21 := GenS1{F0: 620}
                _ = v21
            } else {
                var v22 float32 = ((((0.1 - 3.14159) * 1.5) * (-2.25 + 1e10)) + 0.0)
                _ = v22
            }
            switch -9223372036854775808 {
            case 69:
                var v23 bool = (!(785 >= 1))
                _ = v23
            case 9223372036854775807:
                v24 := GenS3{F0: -32768, F1: 2147483648}
                _ = v24
            default:
                if (-1 == 100) {
                    var v25 uint = 2
                    if ((("a" + ("\n\t" + ("\n\t" + "\n\t"))) + ("foo bar" + "foo bar")) <= "foo bar") {
                        v26 := GenS3{F0: -100, F1: 807}
                        if (true || true) {
                            var v27 float32 = 1e10
                            v28 := GenS4{F0: 65535, F1: (606 >= -1)}
                            _ = v27
                            _ = v28
                        }
                        switch "a" {
                        case "\n\t":
                            v25 = (v25 | (v25 | (v25 + v25)))
                        case "a":
                            v25 = v25
                        case "":
                            v29 := []int{-697}
                            _ = v29
                        }
                        _ = v26
                    } else {
                        if true {
                            var v30 int64 = -430
                            _ = v30
                        }
                        v25 = (v25 >> (((v25 | (^(v25)))) & 63))
                        v31 := []int{}
                        _ = v31
                    }
                    v32 := []int{}
                    _ = v25
                    _ = v32
                } else {
                    var v33 int8 = -128
                    _ = v33
                }
            }
            var v34 int = -2147483648
            _ = v34
        }
    } else {
        for fi35 := 0; fi35 < 19; fi35++ {
            v36 := []int{(fi35 | fi35), 249, (fi35 * fi35)}
            var v37 uint16 = 164
            _ = fi35
            _ = v36
            _ = v37
        }
        switch "" {
        case "\n\t":
            var v38 uint64 = 787
            _ = v38
        default:
            v39 := []int{128, 32767}
            _ = v39
        }
    }
    var v40 uint8 = 0
    _ = v40
    return 128
}
func genFn1() uint16 {
    var v41 uint8 = 2
    v41 -= 84
    if (!true) {
        switch "hello" {
        case "":
            var v42 uint16 = 577
            _ = v42
        default:
            fmt.Println(v41)
        }
    }
    fmt.Println(v41)
    v41 += (v41 % (((^(v41))) | 1))
    _ = v41
    return 513
}

func main() {
    var v1 float64 = -2.25
    var v2 int = 1
    var v3 int8 = 100
    v4 := make(map[uint32]string)
    v43 := []int{v2, v2}
    v2 = (v2 ^ -9223372036854775808)
    v3 += 1
    if (1 == -32768) {
        var v44 bool = false
        v1 = 0.0
        fmt.Println(v1, v2, v44)
        _ = v44
    }
    fmt.Println(v3, v1)
    if false {
        v45 := GenS5{F0: 1, F1: 1}
        var v46 bool = (v2 > genFn0())
        _ = v45
        _ = v46
    } else {
        if false {
            var v47 uint32 = 2147483648
            v1 = v1
            if true {
                var v48 float32 = 0.0
                var v49 string = "a"
                v1 -= v1
                _ = v48
                _ = v49
            }
            _ = v47
        } else {
            var v50 string = ("\n\t" + "hello")
            var v51 float64 = 0.0
            v3 = 42
            _ = v50
            _ = v51
        }
        if len(v43) > 0 {
            v43[0] = (v2 + (v2 * (v2 - (v2 + v2))))
        }
    }
    var v52 bool = (true && true)
    var v53 uint8 = 0
    if (0.1 != -2.25) {
        var v54 int8 = (v3 >> ((0) & 63))
        v4[669] = "foo bar"
        _ = v54
    } else {
        v52 = v52
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
    }
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v52
    _ = v53
}
