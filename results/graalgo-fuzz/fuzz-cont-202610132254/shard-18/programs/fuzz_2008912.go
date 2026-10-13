package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint, p2 int) uint {
    fmt.Println(p0, p1)
    if (65535 < 0) {
        v5 := GenS1{F0: -2147483648, F1: true}
        var v6 int32 = -2147483648
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 269, F1: ((p2 + 32767) <= p2)}
    v8 := GenS1{F0: -2147483648, F1: true}
    v9 := GenS1{F0: -1, F1: (!false)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p1 + p0)
}
func genFn1(p0 bool) string {
    if p0 {
        p0 = p0
        v10 := []int{32768}
        p0 = (3.14159 == (1.5 / (3.14159 - 0.1)))
        _ = v10
    } else {
        p0 = (32768 >= -634)
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := []int{fi11, (fi11 * 32768), (fi11 << ((299) & 63)), (fi11 + (fi11 ^ 2147483647))}
            switch fi11 {
            case -1:
                for fi13 := 0; fi13 < 10; fi13++ {
                    p0 = p0
                    _ = fi13
                }
            case 0:
                for fi14 := 0; fi14 < 15; fi14++ {
                    switch -2147483648 {
                    case -275:
                        p0 = (1e10 <= (1.5 - 0.1))
                    default:
                        v15 := GenS1{F0: -2147483648, F1: p0}
                        _ = v15
                    }
                    var v16 uint = 4294967295
                    _ = fi14
                    _ = v16
                }
            case -988:
                if (!(42 <= 418)) {
                    fmt.Println(fi11, p0)
                    for fi17 := 0; fi17 < 3; fi17++ {
                        if ((1.5 != ((-2.25 / 1.0) + 0.1)) || p0) {
                            var v18 int64 = 0
                            v18 = v18
                            var v19 int64 = v18
                            if (false || p0) {
                                v20 := GenS1{F0: -416, F1: false}
                                switch fi17 {
                                case 126:
                                    p0 = true
                                case -9223372036854775808:
                                    var v21 int = 32767
                                    _ = v21
                                case 32768:
                                    var v22 int8 = 127
                                    _ = v22
                                default:
                                    if (!(!p0)) {
                                        var v23 float32 = -2.25
                                        _ = v23
                                    }
                                }
                                _ = v20
                            } else {
                                v18 -= (v18 / ((v18) | 1))
                                if true {
                                    switch (fi11 / ((fi11) | 1)) {
                                    case 444:
                                        fmt.Println(fi11, v18, fi11)
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    if len(v12) > 0 {
                                        v12[0] = (^(fi11))
                                    }
                                } else {
                                    v18 = v18
                                    var v24 string = "\n\t"
                                    _ = v24
                                }
                            }
                            if (903 > 464) {
                                v19 = (v18 & -1)
                            } else {
                                v25 := GenS1{F0: -2147483648, F1: p0}
                                v26 := GenS2{F0: 0, F1: (v18 | v18)}
                                fmt.Println(p0, v19)
                                _ = v25
                                _ = v26
                            }
                            _ = len(v12)
                            v19 = 637
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            _ = v18
                            _ = v19
                        } else {
                            var v27 uint32 = 449
                            if ((-63 < 0) || p0) {
                                _ = len(v12)
                            }
                            _ = v27
                        }
                        switch ("a" + "hello") {
                        case "foo bar":
                            var v28 uint = 4294967295
                            _ = v28
                        default:
                            p0 = (!(-1 >= 42))
                        }
                        var v29 int = 9223372036854775807
                        _ = fi17
                        _ = v29
                    }
                    v30 := GenS1{F0: -2147483648, F1: p0}
                    _ = v30
                } else {
                    v31 := GenS2{F0: 65535, F1: 9223372036854775807}
                    _ = v31
                }
            }
            _ = fi11
            _ = v12
        }
    }
    if p0 {
        var v32 int = 9223372036854775807
        _ = v32
    }
    p0 = ((-1 < -653) || true)
    _ = p0
    return (("" + "a") + "\n\t")
}

func main() {
    v1 := []uint32{2147483648, 4294967295, 0, 2147483648}
    var v2 int64 = -573
    v3 := []int{42}
    v4 := make(map[string]string)
    v33 := GenS1{F0: -971, F1: (true || (-2147483648 == 185))}
    fmt.Println(v2)
    _ = len(v4)
    var v34 uint = 572
    fmt.Println(v34, v2)
    v34 -= 0
    v34 = 2
    v35 := GenS3{F0: 2}
    switch "foo bar" {
    case "":
        var v36 uint16 = 42
        _ = v36
    case "a":
        var v37 uint32 = 1
        _ = v37
    }
    var v38 int16 = 749
    v4["foo bar"] = (((("\n\t" + ("a" + "")) + "a") + ("hello" + (("\n\t" + ("foo bar" + "a")) + ("a" + "foo bar")))) + (("\n\t" + "\n\t") + ("a" + "foo bar")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
}
