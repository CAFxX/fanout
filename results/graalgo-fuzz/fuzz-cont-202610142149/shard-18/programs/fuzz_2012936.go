package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 bool }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint32; F1 string; F2 uint8 }
func genFn0() int {
    v7 := []int{0}
    var v8 int64 = -1
    if (!(false && (287 != 1))) {
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                v8 = (v8 << ((255) & 63))
                fmt.Println(fi10, v8)
                v11 := GenS1{F0: "\n\t", F1: (v8 | 42), F2: (("a" >= (("foo bar" + "hello") + "hello")) || true)}
                v12 := GenS2{F0: 304, F1: (v8 * v8)}
                v13 := GenS3{F0: 2, F1: "\n\t", F2: 102}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            if false {
                v8 = v8
            }
            _ = fi9
        }
        v8 = 1
    } else {
        _ = len(v7)
    }
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    _ = v7
    _ = v8
    return 0
}
func genFn1() int32 {
    v14 := GenS2{F0: 178, F1: 142}
    v15 := []int{42, -795}
    var v16 int32 = -911
    _ = v14
    _ = v15
    _ = v16
    return v16
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]bool)
    var v3 float64 = 3.14159
    var v4 bool = (true && (0 > -683))
    v5 := []int{-785, 32767, -1}
    v6 := []int{-765, 42, int(v1), 1}
    v3 = (((v3 / v3) - 1e10) + ((-2.25 / ((v3 * (v3 * v3)) * 1.5)) * (1.5 + (float64(0.1) / ((v3 / v3) + v3)))))
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 string = "foo bar"
        _ = fi17
        _ = v18
    }
    v3 = (3.14159 - 1e10)
    if v4 {
        v5 = append(v5, genFn0())
        v1 -= 1
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                v3 -= (0.0 - 0.0)
                fmt.Println(v3, v1, v3)
                _ = len(v2)
                _ = fi20
            }
            var v21 int64 = 15
            for fi22 := 0; fi22 < 5; fi22++ {
                v3 += v3
                _ = fi22
            }
            _ = fi19
            _ = v21
        }
        switch ("" + "\n\t") {
        case "\n\t":
            v1 = (v1 >> ((634) & 63))
        }
        if len(v6) > 0 {
            v6[0] = 0
        }
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        if true {
            v4 = true
        } else {
            if (516 > 149) {
                v4 = (95 > 525)
            }
            v4 = (18446744073709551615 <= 2)
            v24 := GenS1{F0: "foo bar", F1: 1, F2: false}
            _ = v24
        }
        if true {
            v1 += 42
        } else {
            if (((("" + ("hello" + "")) + ("" + ("a" + ("\n\t" + ((("hello" + "") + ((("hello" + "hello") + "") + "")) + ("" + ((("\n\t" + "a") + "a") + "hello"))))))) + "foo bar") > "foo bar") {
                if ((float32(float64(float32(0.1))) + (-2.25 * 3.14159)) <= 0.1) {
                    fmt.Println(v4, v4, fi23)
                    for fi25 := 0; fi25 < 10; fi25++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi25
                    }
                } else {
                    fmt.Println(fi23, v1, fi23)
                }
            } else {
                if v4 {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        switch fi26 {
                        case 1:
                            if v4 {
                                if v4 {
                                    v3 -= v3
                                    var v27 int = (-(fi23))
                                    _ = v27
                                }
                                v2[4294967295] = (((v3 + ((0.0 - v3) / ((v3 * v3) * v3))) * -2.25) > 0.1)
                                v4 = true
                            }
                        case -2147483648:
                            var v28 uint = 1
                            _ = v28
                        }
                        fmt.Println(fi26)
                        delete(v2, 189)
                        _ = fi26
                    }
                    v4 = true
                } else {
                    if (!(!(((((-2.25 + 0.0) * -2.25) + (1e10 * -2.25)) - 3.14159) > 3.14159))) {
                        v4 = (!false)
                        if v4 {
                            switch 9223372036854775807 {
                            case 128:
                                var v29 float32 = 1e10
                                _ = v29
                            }
                            v3 += (0.0 * (1.5 / (v3 / v3)))
                        } else {
                            delete(v2, 65535)
                            _ = len(v2)
                        }
                    }
                }
            }
        }
        _ = fi23
    }
    var v30 bool = (!((v4 && v4) && (2 >= 802)))
    if (!((v30 && (!false)) && v30)) {
        v30 = false
    }
    var v31 int16 = (v1 + v1)
    var v32 int8 = -1
    v4 = true
    for fi33 := 0; fi33 < 9; fi33++ {
        v34 := GenS2{F0: 0, F1: 151}
        _ = fi33
        _ = v34
    }
    if v4 {
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
        fmt.Println(v3)
    }
    var v35 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
