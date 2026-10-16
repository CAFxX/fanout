package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 int32) int8 {
    var v4 int8 = 1
    var v5 int = 32768
    v5 = v5
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 uint8, p1 int8) uint16 {
    v6 := GenS1{F0: -937}
    p1 -= (p1 + (p1 % ((genFn0((p0 - (p0 / ((p0) | 1))), 42)) | 1)))
    _ = p0
    _ = p1
    _ = v6
    return 97
}

func main() {
    v1 := []uint32{2, 65535, 2147483648}
    v2 := []string{("hello" + ((((("\n\t" + (("\n\t" + "a") + (("\n\t" + ("\n\t" + ("" + "foo bar"))) + ((("" + "") + "hello") + "hello")))) + (("foo bar" + ("\n\t" + "a")) + "")) + "\n\t") + ("\n\t" + "foo bar")) + "a")), "\n\t", "hello", ""}
    var v3 int64 = 821
    for fi7 := 0; fi7 < 10; fi7++ {
        if (genFn0(2, 1) > -54) {
            if true {
                fmt.Println(v3, fi7)
                v8 := GenS1{F0: -723}
                v9 := GenS1{F0: 0}
                _ = v8
                _ = v9
            }
            switch ("a" + "a") {
            case "hello":
                switch (fi7 ^ fi7) {
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 42:
                    v10 := GenS1{F0: 810}
                    _ = v10
                default:
                    v3 -= (v3 ^ (v3 >> ((1) & 63)))
                }
            }
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        } else {
            if len(v2) > 0 {
                v2[0] = (("hello" + ("\n\t" + "foo bar")) + "")
            }
            if true {
                var v11 int64 = 0
                _ = len(v1)
                _ = v11
            } else {
                switch -1 {
                case 617:
                    v2 = append(v2, "hello")
                case -356:
                    var v12 float32 = 3.14159
                    _ = v12
                case 128:
                    if true {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch ("" + "\n\t") {
                            case "a":
                                v3 -= (v3 * (v3 >> ((0) & 63)))
                            default:
                                fmt.Println(v3, fi7)
                            }
                            fmt.Println(fi13)
                            var v14 int8 = 1
                            _ = fi13
                            _ = v14
                        }
                        fmt.Println(v3, v3)
                    }
                }
                var v15 int = fi7
                if true {
                    v3 = (^(v3))
                    v3 -= (^(v3))
                } else {
                    _ = len(v2)
                    v2 = append(v2, "\n\t")
                    for fi16 := 0; fi16 < 6; fi16++ {
                        var v17 float64 = 0.0
                        _ = fi16
                        _ = v17
                    }
                }
                _ = v15
            }
            var v18 bool = (!(false && false))
            _ = v18
        }
        var v19 float32 = float32(float64(0.1))
        if len(v2) > 0 {
            v2[0] = "a"
        }
        _ = fi7
        _ = v19
    }
    fmt.Println(v3)
    switch ("hello" + ("a" + "\n\t")) {
    case "":
        v20 := GenS1{F0: -2147483648}
        _ = v20
    default:
        v3 = v3
    }
    fmt.Println(v3)
    fmt.Println(v3)
    for fi21 := 0; fi21 < 10; fi21++ {
        if ("\n\t" <= "foo bar") {
            v3 -= -543
            _ = len(v1)
        } else {
            v3 = -1
            v22 := GenS1{F0: 991}
            _ = v22
        }
        v23 := GenS2{F0: fi21}
        var v24 uint = 1
        _ = fi21
        _ = v23
        _ = v24
    }
    if (!(0 <= 100)) {
        v3 = 1
    }
    v3 += v3
    v25 := GenS1{F0: -1}
    fmt.Println(v3)
    if true {
        if (true || false) {
            v3 = -9223372036854775808
            var v26 uint32 = 42
            v3 += v3
            _ = v26
        } else {
            var v27 uint8 = 45
            _ = v27
        }
    }
    v3 += -1
    if ((0.0 - ((float32(((1.5 * 1.5) / 1.5)) + 0.1) + 0.1)) < 0.1) {
        var v28 bool = (4294967295 <= 0)
        _ = v28
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v29 := GenS2{F0: 86}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
}
