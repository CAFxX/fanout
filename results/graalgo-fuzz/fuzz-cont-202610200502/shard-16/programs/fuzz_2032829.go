package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 string; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 string; F1 string; F2 int32 }
type GenS7 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS8 struct { F0 string; F1 uint8 }
func genFn0(p0 int16, p1 uint64) int16 {
    v4 := GenS1{F0: (!true), F1: 0}
    v5 := GenS2{F0: 127}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 int) uint64 {
    var v6 int = (p0 << ((255) & 63))
    for fi7 := 0; fi7 < 17; fi7++ {
        if true {
            if false {
                var v8 uint64 = 4294967295
                switch (^(v6)) {
                case -9223372036854775808:
                    for fi9 := 0; fi9 < 17; fi9++ {
                        v6 = 42
                        v10 := []int{}
                        var v11 string = ("foo bar" + "foo bar")
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
                _ = v8
            }
        } else {
            v6 += (v6 << ((2) & 63))
            fmt.Println(fi7, v6)
        }
        _ = fi7
    }
    v12 := GenS3{F0: 987, F1: 247}
    fmt.Println(v6)
    _ = p0
    _ = v12
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]string)
    v3 := []int{0, 128, -798}
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS4{F0: "foo bar", F1: (!true), F2: 2}
        for fi15 := 0; fi15 < 10; fi15++ {
            switch ("a" + "hello") {
            case "\n\t":
                var v16 uint32 = 2147483647
                _ = v16
            case "foo bar":
                if false {
                    var v17 int16 = -1
                    v18 := GenS4{F0: ("" + "\n\t"), F1: (0 != 255), F2: 18446744073709551615}
                    _ = v17
                    _ = v18
                } else {
                    if ((!true) || true) {
                        v19 := GenS5{F0: (v1 ^ v1), F1: v1}
                        fmt.Println(fi13, fi15, fi15)
                        _ = v19
                    }
                    v3 = append(v3, fi15)
                    v20 := GenS6{F0: "a", F1: ((("hello" + "") + (("foo bar" + ("" + ("\n\t" + "a"))) + ("a" + (("foo bar" + "hello") + ("hello" + ((("hello" + "a") + "") + "")))))) + ("a" + ("a" + "hello"))), F2: -2147483648}
                    _ = v20
                }
            default:
                var v21 bool = false
                _ = v21
            }
            v1 = v1
            _ = fi15
        }
        switch -1 {
        case -9223372036854775808:
            v22 := GenS1{F0: false, F1: 18446744073709551615}
            _ = v22
        case 222:
            for fi23 := 0; fi23 < 19; fi23++ {
                if (1 >= 2147483647) {
                    switch "hello" {
                    case "hello":
                        var v24 uint16 = 2
                        _ = v24
                    case "":
                        if false {
                            var v25 bool = false
                            v26 := v2[2]
                            var v27 uint16 = 99
                            v26 = (v26 + "\n\t")
                            v28 := GenS7{F0: (v27 + v27), F1: -2147483648, F2: 0}
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        if (-2147483648 >= 860) {
                            v1 += v1
                            v29 := GenS2{F0: -1}
                            _ = v29
                        }
                        switch fi13 {
                        case 958:
                            v1 = (v1 ^ v1)
                        case -1:
                            _ = len(v3)
                        }
                        switch ("\n\t" + ((("hello" + "foo bar") + ("\n\t" + "hello")) + ((("\n\t" + "foo bar") + "a") + "\n\t"))) {
                        case "a":
                            if false {
                                var v30 uint8 = (^(v1))
                                _ = v30
                            } else {
                                fmt.Println(fi13)
                                if false {
                                    v31 := GenS7{F0: 32767, F1: 730, F2: v1}
                                    switch (fi23 % ((fi23) | 1)) {
                                    case -1:
                                        var v32 uint16 = 806
                                        _ = v32
                                    case 488:
                                        v33 := GenS8{F0: "foo bar", F1: v1}
                                        _ = v33
                                    case 63:
                                        v1 += (^(v1))
                                    }
                                    _ = v31
                                }
                                v1 += (v1 % ((5) | 1))
                            }
                        case "hello":
                            var v34 string = "hello"
                            _ = v34
                        default:
                            v1 = (v1 + 1)
                        }
                    case "\n\t":
                        var v35 bool = (("hello" <= "foo bar") && (true && (true || false)))
                        _ = v35
                    }
                    v3 = append(v3, (fi13 / ((2147483647) | 1)))
                    switch "" {
                    case "hello":
                        v1 += (^(v1))
                    case "foo bar":
                        delete(v2, 4294967295)
                    case "":
                        v1 += (v1 ^ v1)
                    }
                }
                _ = fi23
            }
        default:
            switch "foo bar" {
            case "foo bar":
                var v36 uint = 18446744073709551615
                _ = v36
            case "a":
                v1 = (v1 - (^(v1)))
            case "":
                v1 = (v1 + 255)
            }
        }
        _ = fi13
        _ = v14
    }
    var v37 string = "foo bar"
    var v38 bool = false
    v38 = (18446744073709551615 > 18446744073709551615)
    var v39 float32 = (1.5 + 1.5)
    var v40 uint64 = 496
    var v41 string = (v37 + "\n\t")
    switch 32768 {
    case 2147483647:
        v42 := GenS6{F0: v41, F1: (((("foo bar" + "foo bar") + "") + v37) + (v37 + "a")), F2: -394}
        _ = v42
    case 723:
        v43 := GenS5{F0: (v1 & 42), F1: (^(v1))}
        _ = v43
    default:
        fmt.Println(v41, v41, v1)
    }
    var v44 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v44
}
