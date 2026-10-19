package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 bool }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 uint32) string {
    p0 = (p0 % (((p0 - p0)) | 1))
    if true {
        for fi5 := 0; fi5 < 18; fi5++ {
            p0 += 2
            if ((((!((!(false && (!true))) && (("hello" != "") && false))) || true) && (!(42 != 1))) || (!true)) {
                var v6 int64 = -9223372036854775808
                _ = v6
            } else {
                var v7 uint16 = 0
                var v8 int8 = 42
                v9 := GenS1{F0: 128, F1: (v8 * (v8 * v8)), F2: (1 < 4294967295)}
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 uint32 = 4294967295
            _ = fi5
            _ = v10
        }
    } else {
        var v11 uint16 = 42
        v11 += v11
        _ = v11
    }
    _ = p0
    return "foo bar"
}

func main() {
    var v1 string = ""
    v2 := []int{128}
    v3 := []string{(("a" + v1) + v1)}
    v4 := []uint32{2147483648, 480, 2147483648}
    fmt.Println(v1)
    if true {
        v1 = ((v1 + v1) + ("foo bar" + (((("" + v1) + (v1 + "")) + (((v1 + v1) + v1) + v1)) + v1)))
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(fi12, fi12)
            fmt.Println(v1, v1)
            for fi13 := 0; fi13 < 14; fi13++ {
                if (!(((-1 == 42) || (1e10 <= float32((-2.25 - float64(float32(float64((((((0.0 + (((0.1 + 0.0) - -2.25) / 0.1)) / float32(1.5)) / 3.14159) - -2.25) - 0.0)))))))) && true)) {
                    fmt.Println(v1)
                    var v14 int16 = -1
                    v15 := GenS1{F0: 42, F1: -128, F2: false}
                    _ = v14
                    _ = v15
                } else {
                    if (401 > 987) {
                        fmt.Println(fi13, fi12)
                    } else {
                        if (1 <= 233) {
                            var v16 float64 = float64(1e10)
                            var v17 bool = true
                            if (v1 > (((((v1 + v1) + (((v1 + (("hello" + (("hello" + (("" + "foo bar") + "a")) + v1)) + v1)) + ((v1 + (v1 + (v1 + ""))) + v1)) + (((((v1 + "foo bar") + "") + "hello") + "") + "\n\t"))) + v1) + (("hello" + "foo bar") + ((v1 + (("foo bar" + (("" + "\n\t") + (v1 + ("a" + "a")))) + (v1 + "foo bar"))) + v1))) + (((v1 + (("" + v1) + (v1 + v1))) + "\n\t") + ""))) {
                                _ = len(v3)
                                v16 = (-2.25 + v16)
                                var v18 float32 = ((0.1 - 0.0) + (((1e10 / 1.0) * 3.14159) * (((((float32(float64(0.0)) * 1.5) * -2.25) - float32(((v16 / (v16 * v16)) + float64(0.1)))) + 0.1) + 1e10)))
                                _ = v18
                            }
                            _ = v16
                            _ = v17
                        }
                        var v19 int64 = 662
                        v1 = (v1 + (v1 + ((((((v1 + (("" + ("\n\t" + v1)) + ("\n\t" + (v1 + v1)))) + (v1 + ("hello" + "foo bar"))) + v1) + "hello") + v1) + "\n\t")))
                        _ = v19
                    }
                    v1 = ("a" + (v1 + (v1 + (((v1 + (v1 + v1)) + "\n\t") + (v1 + v1)))))
                }
                switch fi12 {
                case 42:
                    v1 = (v1 + ((v1 + v1) + ""))
                }
                switch (fi13 + 64) {
                case 1:
                    var v20 uint16 = 1
                    _ = v20
                case 42:
                    fmt.Println(fi12, v1, fi13)
                case 32767:
                    if (!((((0.0 / float64(0.1)) - 0.0) - 0.0) >= (3.14159 + 3.14159))) {
                        var v21 int32 = -226
                        if false {
                            v21 = v21
                            v22 := GenS2{F0: -1, F1: 2}
                            v21 = (-(v21))
                            _ = v22
                        } else {
                            v23 := GenS1{F0: 38, F1: 1, F2: true}
                            _ = v23
                        }
                        _ = v21
                    }
                default:
                    var v24 string = (v1 + (v1 + v1))
                    _ = v24
                }
                v1 = v1
                if ((false && true) && false) {
                    if len(v3) > 0 {
                        v3[0] = v1
                    }
                    var v25 int8 = -17
                    v25 -= (v25 % ((v25) | 1))
                    _ = v25
                }
                _ = fi13
            }
            _ = fi12
        }
        if (2 != 2) {
            v26 := GenS3{F0: 0, F1: 65535, F2: 2147483648}
            _ = v26
        }
    } else {
        v1 = "a"
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 int = fi27
        v28 = (fi27 >> ((65535) & 63))
        v1 = v1
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v1 = v1
        _ = fi29
    }
    v1 = v1
    var v30 int16 = 42
    if (42 > 4294967295) {
        v30 -= -575
        var v31 uint8 = 2
        v4 = append(v4, uint32(v31))
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := GenS1{F0: 127, F1: 100, F2: false}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi32
        _ = v33
    }
    v2 = append(v2, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
