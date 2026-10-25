package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0() int {
    var v4 int16 = -32768
    v5 := GenS1{F0: -2147483648, F1: 196}
    var v6 int16 = v4
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 += v6
        v8 := []int{fi7}
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return 622
}
func genFn1() uint32 {
    var v9 uint64 = 0
    v9 = (v9 % ((v9) | 1))
    _ = v9
    return 42
}

func main() {
    var v1 int = 32767
    var v2 int = v1
    var v3 bool = true
    v10 := []int{(v1 >> ((2) & 63))}
    var v11 float32 = (1.5 - 3.14159)
    if len(v10) > 0 {
        v10[0] = (-(v2))
    }
    if (-1 != -1) {
        v12 := GenS2{F0: 42, F1: 144}
        _ = v12
    } else {
        var v13 uint32 = 0
        if v3 {
            v3 = ("foo bar" >= ("a" + "hello"))
            v2 -= v1
        } else {
            switch v2 {
            case 475:
                if v3 {
                    if (964 >= v13) {
                        if v3 {
                            switch "\n\t" {
                            case "a":
                                if false {
                                    var v14 uint8 = 204
                                    var v15 int = (v2 & -47)
                                    var v16 int16 = 266
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                            case "foo bar":
                                if (!(true || ((!(!(0 == 68))) || ((0 >= v13) || (v3 && (-9223372036854775808 > -906)))))) {
                                    v3 = (!true)
                                } else {
                                    for fi17 := 0; fi17 < 4; fi17++ {
                                        v13 += (v13 | v13)
                                        switch "" {
                                        case "\n\t":
                                            v10 = append(v10, (v2 << ((2) & 63)))
                                        default:
                                            switch (v1 >> ((796) & 63)) {
                                            case 0:
                                                v18 := GenS3{F0: v3, F1: 0}
                                                _ = v18
                                            case 887:
                                                var v19 bool = v3
                                                _ = v19
                                            case 4294967295:
                                                v2 -= (fi17 * 32767)
                                            }
                                        }
                                        for fi20 := 0; fi20 < 5; fi20++ {
                                            v21 := GenS1{F0: -411, F1: 255}
                                            var v22 int16 = 1
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                        _ = fi17
                                    }
                                }
                            }
                        } else {
                            v23 := GenS2{F0: 127, F1: 1}
                            _ = v23
                        }
                        v10 = append(v10, int(v13))
                    }
                }
            case 32767:
                var v24 uint16 = 65535
                _ = v24
            case 42:
                v2 = v1
            default:
                if false {
                    switch (v2 ^ genFn0()) {
                    case 127:
                        fmt.Println(v13, v3, v3)
                    }
                }
            }
            _ = len(v10)
        }
        var v25 float64 = (1.5 * 0.1)
        _ = v13
        _ = v25
    }
    if len(v10) > 0 {
        v10[0] = (v2 ^ (v1 << ((1) & 63)))
    }
    var v26 uint32 = genFn1()
    v3 = false
    if (!(true || v3)) {
        _ = len(v10)
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS1{F0: -953, F1: 4294967295}
        switch ("\n\t" + ("" + "foo bar")) {
        case "foo bar":
            if (42 <= 4294967295) {
                if (906 == 42) {
                    v1 += v2
                } else {
                    if v3 {
                        var v29 int16 = 32767
                        for fi30 := 0; fi30 < 7; fi30++ {
                            if (v3 || (v3 || true)) {
                                v31 := GenS4{F0: 25, F1: 0}
                                _ = v31
                            } else {
                                if (!(!((9 >= 190) || (255 == 65535)))) {
                                    var v32 bool = (-1 <= fi30)
                                    v10 = append(v10, -1)
                                    var v33 int = v2
                                    var v34 uint = 65535
                                    fmt.Println(v2, v32)
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            }
                            switch v1 {
                            case -1:
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            case -108:
                                v1 = (v2 & genFn0())
                            }
                            var v35 uint32 = 723
                            _ = fi30
                            _ = v35
                        }
                        v3 = (!v3)
                        _ = v29
                    }
                    v1 += (fi27 ^ 2147483647)
                }
                v26 = v26
            } else {
                fmt.Println(fi27)
            }
        case "\n\t":
            v2 = (v2 * v1)
        default:
            v36 := GenS5{F0: v2, F1: (v26 * 2147483647)}
            _ = v36
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
}
