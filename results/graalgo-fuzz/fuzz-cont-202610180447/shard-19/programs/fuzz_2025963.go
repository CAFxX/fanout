package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 string, p1 uint64) int32 {
    var v4 int8 = 42
    v4 = (v4 - v4)
    v4 += (v4 << ((4294967295) & 63))
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    return -2147483648
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v5 float32 = -2.25
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            _ = len(v3)
            _ = fi7
        }
        _ = fi6
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if (-976 != 97) {
            for fi9 := 0; fi9 < 13; fi9++ {
                if (!true) {
                    var v10 int32 = -2147483648
                    _ = len(v3)
                    v11 := []int{fi8, (fi8 | (fi8 * (fi9 & (fi9 >> ((65535) & 63)))))}
                    _ = v10
                    _ = v11
                }
                v5 = (((3.14159 - -2.25) - (((v5 / v5) / (v5 / v5)) / 1e10)) + v5)
                _ = fi9
            }
        } else {
            v12 := GenS1{F0: 0, F1: 765}
            v13 := []int{-67, fi8}
            _ = v12
            _ = v13
        }
        fmt.Println(v1, v5, fi8)
        _ = fi8
    }
    var v14 uint8 = 42
    var v15 bool = false
    for fi16 := 0; fi16 < 20; fi16++ {
        v14 = (v14 * (v14 % ((v14) | 1)))
        _ = fi16
    }
    if v15 {
        v17 := GenS2{F0: 1}
        for fi18 := 0; fi18 < 16; fi18++ {
            switch (("foo bar" + "hello") + (("hello" + "a") + (("\n\t" + ("a" + ((("foo bar" + "a") + "foo bar") + "\n\t"))) + ("foo bar" + "")))) {
            case "\n\t":
                fmt.Println(v1, v14)
            case "hello":
                v14 = 128
            default:
                var v19 int64 = 0
                _ = v19
            }
            _ = fi18
        }
        _ = v17
    }
    v15 = v15
    fmt.Println(v14)
    for fi20 := 0; fi20 < 3; fi20++ {
        switch "foo bar" {
        case "hello":
            for fi21 := 0; fi21 < 3; fi21++ {
                v15 = true
                _ = fi21
            }
        case "foo bar":
            for fi22 := 0; fi22 < 16; fi22++ {
                if (true || ((-1 > -9223372036854775808) && (v15 || false))) {
                    v1 -= v1
                    switch (^(fi22)) {
                    case 42:
                        if v15 {
                            if false {
                                v23 := []int{(fi22 & fi22), (fi20 % (((fi22 * fi20)) | 1)), (fi20 - fi22), fi22}
                                v24 := GenS1{F0: 18446744073709551615, F1: -1}
                                v1 = -32768
                                var v25 uint32 = 460
                                if (false || (!v15)) {
                                    delete(v3, (^(v25)))
                                    delete(v3, v25)
                                }
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                switch fi20 {
                                case 42:
                                    v2[""] = 2147483648
                                default:
                                    if false {
                                        v1 -= (v1 * v1)
                                        delete(v2, ("a" + "a"))
                                        if (0.1 != v5) {
                                            v14 = v14
                                        }
                                    } else {
                                        _ = len(v3)
                                        v15 = ((fi20 != 127) && false)
                                    }
                                }
                                _ = len(v3)
                                var v26 int = fi20
                                _ = v26
                            }
                            if (v15 || ((("foo bar" + ("a" + "hello")) + "foo bar") < "foo bar")) {
                                var v27 int16 = 0
                                _ = v27
                            } else {
                                v28 := []int{fi22, fi20, fi20}
                                var v29 uint32 = 719
                                var v30 int = fi20
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            var v31 uint16 = 265
                            var v32 uint16 = v31
                            _ = v31
                            _ = v32
                        } else {
                            var v33 uint16 = 42
                            v33 -= (v33 ^ 2)
                            v33 += (v33 & 65535)
                            _ = v33
                        }
                    default:
                        _ = len(v3)
                    }
                    fmt.Println(v1, v1, v5)
                    if v15 {
                        v5 = v5
                        v5 = 0.1
                    }
                } else {
                    var v34 int8 = 100
                    v35 := []int{int(v34), (-(fi20)), (-(fi20)), fi22}
                    var v36 uint64 = 238
                    _ = v34
                    _ = v35
                    _ = v36
                }
                v37 := GenS1{F0: 1, F1: -866}
                _ = fi22
                _ = v37
            }
        }
        if true {
            fmt.Println(v5)
            delete(v2, ("" + "\n\t"))
            v38 := []int{}
            _ = v38
        } else {
            var v39 uint16 = 1
            var v40 int32 = 432
            _ = v39
            _ = v40
        }
        _ = fi20
    }
    v41 := []int{32767}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v41
    _ = v5
}
