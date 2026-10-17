package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 bool; F1 int; F2 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 int }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int) uint64 {
    switch (^(p0)) {
    case -407:
        var v4 string = ("\n\t" + (("a" + "") + ((("hello" + "") + "a") + "\n\t")))
        _ = v4
    case 1:
        var v5 float32 = ((((0.0 + 1e10) - (-2.25 / 1.0)) * (3.14159 * 1e10)) - ((float32(((-2.25 * 3.14159) - 0.1)) - 1.5) - 1e10))
        _ = v5
    default:
        switch (("hello" + "") + "\n\t") {
        case "\n\t":
            v6 := GenS1{F0: 0, F1: ("hello" + "foo bar"), F2: 618}
            _ = v6
        case "hello":
            p0 = p0
        case "foo bar":
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 uint32 = 2
                v8 -= (v8 % (((v8 % ((893) | 1))) | 1))
                var v9 bool = true
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            v10 := []int{p0, -9223372036854775808, -9223372036854775808, p0}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    _ = p0
    return 166
}
func genFn1() uint64 {
    var v12 float64 = 1e10
    v13 := GenS2{F0: 2147483648, F1: 100, F2: -1}
    v12 = v12
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 bool = (605 >= 2)
    v2 := make(map[int]bool)
    v3 := []string{"a", "hello", "hello"}
    var v14 uint = 111
    fmt.Println(v14, v1)
    if false {
        var v15 string = ""
        _ = v15
    } else {
        delete(v2, 127)
        if ("hello" <= ("hello" + "a")) {
            v14 += (v14 | (v14 << ((v14) & 63)))
            if len(v3) > 0 {
                v3[0] = ("hello" + "a")
            }
        } else {
            fmt.Println(v14)
            v16 := GenS3{F0: (-2.25 <= 1e10), F1: 118, F2: -32768}
            for fi17 := 0; fi17 < 3; fi17++ {
                v1 = (127 <= -128)
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := v2[fi17]
                    _ = fi18
                    _ = v19
                }
                fmt.Println(fi17, v1)
                _ = fi17
            }
            _ = v16
        }
    }
    delete(v2, 571)
    _ = len(v3)
    fmt.Println(v14, v14)
    if v1 {
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint8 = 97
            v21 -= (v21 / (((v21 * 17)) | 1))
            if false {
                var v22 int8 = 1
                _ = v22
            }
            _ = fi20
            _ = v21
        }
        switch ("hello" + "\n\t") {
        case "":
            var v23 float32 = 0.0
            _ = v23
        case "a":
            for fi24 := 0; fi24 < 11; fi24++ {
                if (0 == 1) {
                    var v25 int32 = -1
                    fmt.Println(v14)
                    _ = v25
                }
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 int64 = 0
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        v2[fi26] = (v1 && false)
                    default:
                        switch fi24 {
                        case -527:
                            fmt.Println(fi24, v27, v27)
                        case -1:
                            v28 := GenS2{F0: 259, F1: 0, F2: 100}
                            _ = v28
                        case -344:
                            if (42 < uint8(v14)) {
                                if v1 {
                                    _ = len(v2)
                                    var v29 uint8 = 255
                                    v30 := GenS4{F0: 927, F1: -137, F2: (^(fi26))}
                                    v31 := GenS5{F0: 1}
                                    fmt.Println(v29)
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                _ = len(v2)
                            } else {
                                fmt.Println(v1, fi26, v14)
                                if false {
                                    var v32 uint64 = 2
                                    _ = v32
                                }
                            }
                        }
                    }
                    v33 := GenS4{F0: 592, F1: 1000, F2: fi26}
                    _ = fi26
                    _ = v27
                    _ = v33
                }
                var v34 uint32 = 994
                _ = fi24
                _ = v34
            }
        }
    }
    if ((1 == 0) && (!v1)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        delete(v2, -2147483648)
    }
    var v35 int8 = 0
    for fi36 := 0; fi36 < 8; fi36++ {
        v14 = (^(v14))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi36
    }
    var v37 int32 = -1
    v2[42] = v1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v37
}
