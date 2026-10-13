package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int; F2 uint16 }
func genFn0() uint64 {
    var v4 int16 = 32767
    var v5 int32 = 1
    _ = v4
    _ = v5
    return 534
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[int]uint32)
    var v3 int16 = 42
    var v6 int32 = 42
    v2[304] = 1
    fmt.Println(v6)
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 = append(v1, "")
        fmt.Println(v3, v6)
        _ = fi8
    }
    switch 152 {
    case 726:
        if (-128 < 127) {
            v2[-9223372036854775808] = v7
            v9 := GenS1{F0: -231, F1: v3}
            delete(v2, 0)
            _ = v9
        } else {
            if (2147483647 == v7) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    _ = len(v1)
                    _ = fi10
                }
                if (!(false && (2 < 810))) {
                    var v11 uint = 42
                    _ = v11
                }
                var v12 float64 = (-2.25 + 3.14159)
                _ = v12
            } else {
                var v13 int8 = 100
                if true {
                    var v14 int64 = 9223372036854775807
                    var v15 int16 = (v3 & (v3 - v3))
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 13; fi16++ {
                    fmt.Println(v13, v6)
                    _ = fi16
                }
                _ = v13
            }
            var v17 string = "foo bar"
            _ = v17
        }
    case -2147483648:
        fmt.Println(v6)
    case -172:
        if ((("\n\t" + ("" + ("\n\t" + ("a" + ("" + "\n\t"))))) + "foo bar") == "hello") {
            v6 = v6
            v7 = v7
            v18 := GenS2{F0: v6, F1: ""}
            _ = v18
        }
    default:
        v3 += (v3 * -646)
    }
    var v19 int8 = 48
    if (!((true || true) && (("" + "a") < ("foo bar" + "")))) {
        if (-2.25 != 0.0) {
            delete(v2, 32767)
        }
        v6 = v6
        fmt.Println(v6, v3)
    } else {
        var v20 int64 = -9223372036854775808
        fmt.Println(v3, v7)
        _ = v20
    }
    v21 := GenS3{F0: v7, F1: (^(v7)), F2: 42}
    switch 32767 {
    case 9223372036854775807:
        var v22 uint64 = 887
        _ = v22
    case -290:
        v3 -= v3
    case 1:
        if true {
            if (!false) {
                if true {
                    v19 = (v19 + (v19 & -125))
                    v3 = (v3 % ((v3) | 1))
                    v19 += v19
                }
                v7 = v7
            }
            if (33 >= 2) {
                if (((v19 % ((-1) | 1)) > (v19 | v19)) || ((-2.25 > -2.25) && (((v19 % ((v19) | 1)) >= v19) && (!(!(("\n\t" + ("hello" + ("a" + ""))) >= "a")))))) {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        v2[(^(fi23))] = 1
                        _ = fi23
                    }
                    v7 -= v7
                }
                v24 := GenS2{F0: (-(v6)), F1: ("\n\t" + "hello")}
                var v25 int8 = v19
                _ = v24
                _ = v25
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (65535 != 0) {
                    if (2147483648 < (v7 + (v7 + (v7 * (v7 >> ((0) & 63)))))) {
                        var v26 float64 = ((0.1 / 0.1) + float64(-2.25))
                        _ = v26
                    }
                }
                var v27 int8 = v19
                _ = v27
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                v19 += (-(v19))
                var v29 int = fi28
                _ = fi28
                _ = v29
            }
        }
    default:
        v7 -= 458
    }
    v7 += 0
    v6 -= (v6 - v6)
    var v30 uint = 255
    v31 := GenS4{F0: 0, F1: 2147483647, F2: 1}
    var v32 int = 4294967295
    var v33 int64 = -894
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v6
    _ = v7
}
