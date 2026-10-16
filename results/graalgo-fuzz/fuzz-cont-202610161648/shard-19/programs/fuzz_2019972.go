package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int8; F2 uint }
func genFn0(p0 uint8) uint64 {
    fmt.Println(p0)
    v7 := []int{32767, -195, -341}
    v8 := GenS1{F0: "\n\t", F1: 96, F2: ("hello" + "")}
    _ = len(v7)
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1() int16 {
    var v9 int64 = -207
    switch 9223372036854775807 {
    case -2147483648:
        var v10 string = "hello"
        _ = v10
    case 127:
        v11 := []int{127, -1}
        _ = v11
    default:
        fmt.Println(v9)
    }
    var v12 uint16 = 0
    _ = v12
    _ = v9
    return 32767
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := []uint32{0}
    v5 := make(map[uint32]string)
    var v6 int64 = 1
    fmt.Println(v6)
    _ = len(v2)
    if ((v6 << ((255) & 63)) < v6) {
        v13 := GenS1{F0: (("foo bar" + "\n\t") + "a"), F1: 18446744073709551615, F2: ("foo bar" + (("a" + "foo bar") + ("hello" + "foo bar")))}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v14 float32 = 1.5
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 18; fi15++ {
            fmt.Println(v6, fi15, v1)
            _ = fi15
        }
        v1 += -1
        v1 -= v1
    }
    var v16 bool = false
    v4 = append(v4, 2)
    v1 -= (v1 * v1)
    for fi17 := 0; fi17 < 5; fi17++ {
        switch "hello" {
        case "foo bar":
            var v18 uint8 = 242
            _ = v18
        case "hello":
            switch "\n\t" {
            case "hello":
                var v19 int8 = 1
                _ = v19
            case "":
                if (v16 && v16) {
                    var v20 int32 = (v1 | (v1 | v1))
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v4 = append(v4, 65535)
                        _ = fi21
                    }
                    switch ("\n\t" + (("foo bar" + "\n\t") + ("\n\t" + "\n\t"))) {
                    case "\n\t":
                        var v22 uint = 694
                        _ = v22
                    case "a":
                        var v23 uint32 = 0
                        _ = v23
                    case "hello":
                        for fi24 := 0; fi24 < 20; fi24++ {
                            if (genFn1() <= -32768) {
                                switch (-(fi24)) {
                                case 0:
                                    if (v16 && v16) {
                                        v20 = v20
                                        v1 = 42
                                        if ((-2.25 < float64(0.1)) && (0 <= 0)) {
                                            v20 = (v1 % ((v20) | 1))
                                            var v25 string = ((((("" + "\n\t") + "\n\t") + ((("hello" + "a") + ("a" + "")) + "a")) + ("\n\t" + ("hello" + "foo bar"))) + "")
                                            _ = v25
                                        }
                                        v26 := GenS2{F0: 42, F1: 18446744073709551615, F2: -422}
                                        switch "\n\t" {
                                        case "":
                                            if (922 <= 2) {
                                                var v27 uint8 = 42
                                                _ = v27
                                            }
                                        default:
                                            v1 -= (v20 | (v1 << ((2) & 63)))
                                        }
                                        _ = v26
                                    } else {
                                        fmt.Println(fi24, fi24)
                                    }
                                case -232:
                                    if len(v4) > 0 {
                                        v4[0] = 2
                                    }
                                default:
                                    delete(v3, 4294967295)
                                }
                                var v28 int64 = v6
                                if v16 {
                                    if (!(v16 && false)) {
                                        v16 = (618 != 206)
                                        var v29 int64 = v6
                                        var v30 int = (fi17 >> ((208) & 63))
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 int8 = -1
                                    _ = v31
                                } else {
                                    fmt.Println(v6, v20)
                                    var v32 float64 = 1e10
                                    _ = v32
                                }
                                _ = v28
                            }
                            var v33 int16 = genFn1()
                            var v34 uint64 = 926
                            _ = fi24
                            _ = v33
                            _ = v34
                        }
                    default:
                        var v35 int64 = v6
                        _ = v35
                    }
                    _ = v20
                }
            case "a":
                v36 := GenS3{F0: v1}
                _ = v36
            }
        default:
            if false {
                var v37 int = -558
                v37 = (fi17 ^ v37)
                v1 -= v1
                _ = v37
            }
        }
        v16 = true
        _ = fi17
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 int32 = v1
        v1 = 2147483647
        _ = fi38
        _ = v39
    }
    if true {
        v40 := GenS3{F0: (v1 >> ((167) & 63))}
        _ = v40
    }
    v41 := GenS4{F0: "foo bar", F1: -128, F2: 65535}
    switch ("a" + "\n\t") {
    case "hello":
        v1 -= (^(v1))
    case "\n\t":
        v6 -= v6
    default:
        var v42 uint = 460
        _ = v42
    }
    var v43 bool = true
    v6 = v6
    var v44 int64 = (v6 - (v6 >> ((2) & 63)))
    var v45 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
