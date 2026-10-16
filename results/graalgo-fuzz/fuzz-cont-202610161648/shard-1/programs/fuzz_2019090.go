package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 int, p2 int) bool {
    p0 += 42
    v4 := []int{(p2 + 32768), -133}
    var v5 uint16 = 32767
    if false {
        p1 += (p2 / (((p1 | (p1 << ((949) & 63)))) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 339
    var v3 uint = 716
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 1
    v3 -= v3
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: 0, F1: 54, F2: 32767}
        switch (fi9 | 9223372036854775807) {
        case -101:
            if (42 > 128) {
                v1 = v1
                var v11 uint16 = 65535
                v12 := GenS2{F0: ("\n\t" + ("\n\t" + ("" + ("foo bar" + ((("hello" + "") + (("hello" + ("a" + ("" + ("\n\t" + (("a" + "foo bar") + "a"))))) + "foo bar")) + "foo bar")))))}
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: 167, F1: 0, F2: v8}
                _ = v13
            }
        case -677:
            var v14 float32 = (3.14159 / 0.1)
            _ = v14
        case 1:
            v15 := GenS3{F0: 2147483648, F1: 2}
            _ = v15
        default:
            v16 := []int{fi9}
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    v1 = ((float64((1.5 + 0.1)) - v1) + v1)
    var v17 int16 = 999
    v18 := GenS4{F0: (^(v17))}
    v19 := []int{-709}
    if (!true) {
        if (!(false && false)) {
            _ = len(v19)
            v2 += 12
        }
    }
    switch (("" + ((((("foo bar" + "foo bar") + "a") + (("" + ("" + ("hello" + ""))) + "")) + ("hello" + ("hello" + "a"))) + ("hello" + ""))) + "") {
    case "a":
        v3 = v3
    case "":
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 uint8 = 71
            var v22 float64 = v1
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 uint = (v3 | (v3 >> ((281) & 63)))
    if true {
        v2 += 63
        switch 32767 {
        case -949:
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
        fmt.Println(v2)
    } else {
        if (!(("\n\t" + "hello") < ("foo bar" + ""))) {
            if (!false) {
                for fi24 := 0; fi24 < 12; fi24++ {
                    if false {
                        var v25 int16 = 32767
                        if (true && (true && (false && false))) {
                            fmt.Println(v2)
                        }
                        _ = v25
                    }
                    _ = fi24
                }
            }
            var v26 uint = (v3 << (((v3 >> ((v23) & 63))) & 63))
            _ = v26
        } else {
            v19 = append(v19, 32767)
            if (v17 == (v17 / (((v17 >> (((^(v23))) & 63))) | 1))) {
                var v27 uint16 = v8
                _ = v27
            }
        }
    }
    if ((v8 ^ 42) <= v8) {
        v28 := GenS3{F0: 42, F1: 122}
        _ = v28
    } else {
        var v29 uint32 = 42
        v2 -= 2
        if false {
            var v30 int16 = (v17 % (((v17 & -32768)) | 1))
            v30 += 32767
            _ = v30
        }
        _ = v29
    }
    fmt.Println(v2, v1)
    v3 = v3
    v23 -= (v23 * 1)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v8
}
