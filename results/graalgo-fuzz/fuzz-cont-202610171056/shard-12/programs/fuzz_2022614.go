package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 string }
func genFn0() int8 {
    var v6 int64 = -9223372036854775808
    v6 += (v6 >> ((0) & 63))
    v6 = (-(v6))
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{(fi7 + fi7)}
        v9 := GenS1{F0: 42, F1: "a"}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return 42
}
func genFn1(p0 bool) int {
    p0 = (-128 > 1)
    p0 = (p0 && true)
    var v10 int16 = -696
    _ = p0
    _ = v10
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 42
    var v3 uint = 0
    var v4 int = 127
    var v5 uint16 = 32767
    var v11 uint64 = 37
    if (false || (true && true)) {
        switch (-(v4)) {
        case 305:
            v12 := GenS1{F0: 255, F1: ("" + "foo bar")}
            _ = v12
        case -1:
            v1[(("foo bar" + "foo bar") + "\n\t")] = 0
        case -2147483648:
            if (1e10 > 3.14159) {
                delete(v1, "\n\t")
            }
        default:
            var v13 int8 = 100
            _ = v13
        }
        fmt.Println(v2, v3, v11)
    } else {
        if ("" <= "a") {
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{genFn1((!((v11 << ((v3) & 63)) == v11))), genFn1((true && false))}
                _ = fi14
                _ = v15
            }
            v2 = v2
            v16 := GenS1{F0: (v5 / ((v5) | 1)), F1: "foo bar"}
            _ = v16
        }
    }
    var v17 int = genFn1(true)
    var v18 int32 = -1
    v17 += v4
    v19 := v1["\n\t"]
    if (v5 == (v5 / ((v5) | 1))) {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS2{F0: 0, F1: true}
            switch 9223372036854775807 {
            case -370:
                v2 -= (v2 + v2)
            case -812:
                delete(v1, "a")
            case 0:
                v4 = (fi20 | fi20)
            }
            v22 := GenS3{F0: ("\n\t" + "foo bar"), F1: 1}
            _ = fi20
            _ = v21
            _ = v22
        }
        switch int(v11) {
        case 32768:
            var v23 int16 = 234
            _ = v23
        }
    } else {
        if true {
            v18 -= (v18 >> ((1) & 63))
        } else {
            v19 = (v19 % (((v19 - 2147483647)) | 1))
            if false {
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 float32 = (3.14159 - float32(0.0))
                    v18 = -1
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v27 := GenS3{F0: ("hello" + "\n\t"), F1: 42}
                        switch (v4 / ((fi24) | 1)) {
                        case 127:
                            v28 := GenS4{F0: (v2 * v2), F1: (v2 * (^(v2)))}
                            _ = v28
                        case -743:
                            var v29 bool = (((!(1e10 == 0.0)) && false) && (42 > -128))
                            _ = v29
                        default:
                            v30 := v1[("hello" + ("a" + ("hello" + "\n\t")))]
                            _ = v30
                        }
                        fmt.Println(fi26, v3)
                        _ = fi26
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    v32 := GenS3{F0: ("\n\t" + ("\n\t" + "")), F1: 0}
                    if (-519 != -1) {
                        var v33 int8 = -112
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
            }
            v1["\n\t"] = (v19 + (v19 << ((v3) & 63)))
        }
        fmt.Println(v18)
        v1[(("foo bar" + ("\n\t" + "")) + "foo bar")] = v19
    }
    var v34 uint8 = v2
    v34 = v34
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 uint64 = (v11 % ((v11) | 1))
        for fi37 := 0; fi37 < 19; fi37++ {
            fmt.Println(v34, v36, v2)
            v2 = v34
            if true {
                switch fi37 {
                case 127:
                    fmt.Println(fi37, v11)
                    v38 := GenS1{F0: (^(v5)), F1: "foo bar"}
                    v39 := []int{}
                    _ = v38
                    _ = v39
                case -2147483648:
                    fmt.Println(v5)
                case -1:
                    _ = len(v1)
                default:
                    switch fi35 {
                    case 4294967295:
                        fmt.Println(v17, v34)
                    case 128:
                        v40 := GenS5{F0: "foo bar"}
                        _ = v40
                    case -2147483648:
                        v3 += (v3 * 2)
                    default:
                        v11 = v11
                    }
                }
            } else {
                fmt.Println(v36, v19)
                v5 = (^(v5))
                v34 = (^(v2))
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    var v41 uint8 = 90
    switch "hello" {
    case "foo bar":
        var v42 string = "foo bar"
        _ = v42
    }
    var v43 bool = (128 != 812)
    for fi44 := 0; fi44 < 17; fi44++ {
        v11 += v11
        var v45 int64 = int64(v19)
        v17 = genFn1(v43)
        _ = fi44
        _ = v45
    }
    fmt.Println(v17, v4)
    fmt.Println(v11, v18)
    fmt.Println(v41)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v43
    _ = v5
}
