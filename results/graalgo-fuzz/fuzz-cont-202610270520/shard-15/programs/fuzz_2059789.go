package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0() uint {
    switch 32768 {
    case 32768:
        var v5 float32 = ((3.14159 * 0.0) - 0.1)
        _ = v5
    }
    var v6 bool = (!((1 < 795) || (0 != 1)))
    v7 := GenS1{F0: -1}
    v8 := GenS1{F0: -517}
    _ = v6
    _ = v7
    _ = v8
    return 545
}
func genFn1(p0 bool, p1 uint64, p2 string) int64 {
    p1 = p1
    v9 := []int{-825}
    var v10 uint32 = 814
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int64 = 1
        v10 -= v10
        switch ((p2 + p2) + ((p2 + ((p2 + (((p2 + ("" + "hello")) + p2) + ("a" + p2))) + p2)) + (p2 + ""))) {
        case "":
            if p0 {
                v13 := GenS1{F0: 286}
                v14 := GenS1{F0: 516}
                p0 = (fi11 < fi11)
                _ = v13
                _ = v14
            }
        default:
            switch ((p2 + "hello") + "foo bar") {
            case "foo bar":
                v15 := GenS2{F0: 959, F1: 215}
                _ = v15
            case "a":
                if (473 < 65535) {
                    v9 = append(v9, (fi11 ^ (fi11 | (fi11 ^ (fi11 ^ (-(fi11)))))))
                    for fi16 := 0; fi16 < 18; fi16++ {
                        _ = len(v9)
                        _ = fi16
                    }
                } else {
                    v9 = append(v9, fi11)
                }
            case "hello":
                if p0 {
                    p2 = (("foo bar" + p2) + p2)
                    switch (((p2 + p2) + ((p2 + ("" + (p2 + p2))) + p2)) + p2) {
                    case "hello":
                        var v17 uint32 = 2147483647
                        _ = v17
                    }
                } else {
                    var v18 int8 = -1
                    if (true && (!true)) {
                        v12 = -261
                        for fi19 := 0; fi19 < 17; fi19++ {
                            if len(v9) > 0 {
                                v9[0] = (fi11 / ((0) | 1))
                            }
                            _ = fi19
                        }
                        for fi20 := 0; fi20 < 1; fi20++ {
                            if (!(!((p1 / ((p1) | 1)) <= 4294967295))) {
                                var v21 int = fi20
                                _ = v21
                            }
                            var v22 uint64 = p1
                            _ = fi20
                            _ = v22
                        }
                    } else {
                        v23 := GenS3{F0: v18, F1: genFn0()}
                        _ = len(v9)
                        if (!(!(!p0))) {
                            switch p2 {
                            case "a":
                                if (((!p0) || p0) || (!(0 > 1))) {
                                    v10 = uint32(p1)
                                    var v24 uint32 = v10
                                    _ = v24
                                } else {
                                    var v25 int8 = (v18 % (((v18 % (((v18 * 127)) | 1))) | 1))
                                    _ = v25
                                }
                            default:
                                fmt.Println(fi11, fi11, p0)
                            }
                            v26 := GenS4{F0: int8(p1), F1: -32768}
                            p0 = (p0 && (65535 < 255))
                            _ = v26
                        } else {
                            _ = len(v9)
                        }
                        _ = v23
                    }
                    var v27 uint8 = 120
                    _ = v18
                    _ = v27
                }
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -268
}

func main() {
    var v1 float32 = 0.1
    v2 := []string{}
    var v3 uint64 = 2
    var v4 int16 = 1
    v3 = v3
    var v28 uint64 = v3
    if true {
        var v29 int = 988
        _ = v29
    }
    switch 1 {
    case -918:
        if len(v2) > 0 {
            v2[0] = ("hello" + ((("foo bar" + (("\n\t" + "a") + ((("foo bar" + "a") + "foo bar") + "\n\t"))) + "hello") + "a"))
        }
    case 9223372036854775807:
        v2 = append(v2, "")
    default:
        v30 := GenS2{F0: 1, F1: genFn0()}
        _ = v30
    }
    _ = len(v2)
    v1 += ((float32(((0.1 / (1.5 * -2.25)) - 0.0)) - 1.5) - v1)
    v4 -= v4
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
