package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int8 }
func genFn0(p0 uint64, p1 int32) uint32 {
    v7 := GenS1{F0: 2, F1: (p0 - 18446744073709551615), F2: ("\n\t" + ("a" + (("" + ("a" + "hello")) + (("" + ("\n\t" + ("" + "foo bar"))) + ("\n\t" + (("\n\t" + "a") + "a"))))))}
    var v8 uint8 = 118
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(p1, v8, fi9)
        v10 := GenS1{F0: uint32(v8), F1: 448, F2: ("a" + "foo bar")}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if true {
            switch 0 {
            case 906:
                var v12 string = ""
                _ = v12
            default:
                switch "hello" {
                case "\n\t":
                    p0 -= (^(p0))
                case "hello":
                    for fi13 := 0; fi13 < 8; fi13++ {
                        fmt.Println(v8, p1)
                        v14 := []int{fi11}
                        if (400 > 537) {
                            fmt.Println(fi13, fi11)
                            var v15 uint8 = (^(v8))
                            _ = v15
                        }
                        _ = fi13
                        _ = v14
                    }
                default:
                    p1 = p1
                }
            }
        }
        _ = fi11
    }
    if false {
        switch int(p0) {
        case -2147483648:
            p1 += (p1 & p1)
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int, p1 int8) uint64 {
    fmt.Println(p0, p1)
    var v16 string = "hello"
    v17 := GenS1{F0: 2147483648, F1: 125, F2: ((("hello" + (v16 + v16)) + (v16 + ("a" + ("\n\t" + (v16 + "foo bar"))))) + v16)}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 2
    var v3 bool = ((false || false) || true)
    var v4 int16 = 330
    v5 := []int{32767, 32768, -1, 4294967295}
    v6 := []string{("a" + "\n\t"), "hello", "foo bar", "\n\t"}
    v5 = append(v5, 2147483647)
    v18 := GenS1{F0: 0, F1: 1, F2: ("\n\t" + (("" + ("a" + "\n\t")) + "\n\t"))}
    if ((-2.25 * 1.5) < float32((v1 * v1))) {
        v19 := GenS1{F0: genFn0(18446744073709551615, 631), F1: 42, F2: "foo bar"}
        v20 := GenS1{F0: 2147483648, F1: 776, F2: "hello"}
        v2 += (v2 | v2)
        _ = v19
        _ = v20
    } else {
        if (v1 < -2.25) {
            switch 846 {
            case 32767:
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS2{F0: 2, F1: 42, F2: 42}
                    if (fi21 <= 127) {
                        if true {
                            switch ("foo bar" + ("a" + "hello")) {
                            case "foo bar":
                                v23 := GenS2{F0: v2, F1: 0, F2: 42}
                                _ = v23
                            case "a":
                                var v24 uint = (v2 * v2)
                                _ = v24
                            case "hello":
                                var v25 uint8 = 215
                                _ = v25
                            default:
                                var v26 int = (fi21 - (fi21 / ((fi21) | 1)))
                                _ = v26
                            }
                        } else {
                            fmt.Println(v4)
                            var v27 int16 = v4
                            _ = v27
                        }
                    } else {
                        var v28 string = (("\n\t" + "hello") + "")
                        v3 = ((0.0 - 1.5) != 0.1)
                        _ = v28
                    }
                    v29 := GenS3{F0: 4294967295, F1: 35, F2: 100}
                    _ = fi21
                    _ = v22
                    _ = v29
                }
            }
        }
    }
    var v30 uint = (v2 % ((18446744073709551615) | 1))
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS1{F0: 2, F1: 738, F2: "hello"}
        _ = fi31
        _ = v32
    }
    var v33 uint64 = 674
    v1 -= v1
    fmt.Println(v2, v33, v4)
    switch 1 {
    case -943:
        var v34 float32 = 1e10
        _ = v34
    case 9223372036854775807:
        if v3 {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v1)
        }
    default:
        var v35 uint32 = 1
        _ = v35
    }
    fmt.Println(v33, v2, v33)
    for fi36 := 0; fi36 < 12; fi36++ {
        fmt.Println(v1)
        if v3 {
            if false {
                var v37 uint16 = 807
                v4 += v4
                _ = v37
            }
        } else {
            _ = len(v6)
            var v38 int32 = 1
            _ = v38
        }
        v33 += (v33 << (((v2 + (^(v2)))) & 63))
        _ = fi36
    }
    var v39 int64 = -497
    fmt.Println(v1, v4, v3)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
