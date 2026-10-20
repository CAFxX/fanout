package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 bool; F2 int16 }
type GenS4 struct { F0 int16 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int = -1
            _ = fi5
            _ = v6
        }
        var v7 int8 = 55
        v7 += 1
        _ = fi4
        _ = v7
    }
    var v8 float32 = (1.5 * 0.0)
    v8 -= (v8 * (0.0 / ((float32(1.5) / (v8 - v8)) / 0.1)))
    v9 := GenS1{F0: 0}
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 int8) int16 {
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := []int{(^(fi10)), fi10, (fi10 << ((4294967295) & 63))}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        p0 = (p0 & p0)
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := []int{fi12}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    switch ("foo bar" + ("a" + "a")) {
    case "a":
        if true {
            v15 := GenS2{F0: 0}
            v16 := []int{9223372036854775807, -1, -2147483648, -1}
            var v17 int16 = -418
            _ = v15
            _ = v16
            _ = v17
        }
    }
    v18 := []int{128, 4294967295, 9223372036854775807, 32767}
    _ = p0
    _ = v18
    return 32767
}

func main() {
    var v1 float32 = (float32(1.5) * 1.5)
    v2 := []int{-1, 32767, 42}
    var v3 uint32 = 2
    _ = len(v2)
    var v19 int64 = 9223372036854775807
    for fi20 := 0; fi20 < 11; fi20++ {
        if true {
            for fi21 := 0; fi21 < 20; fi21++ {
                if len(v2) > 0 {
                    v2[0] = (-(fi21))
                }
                _ = fi21
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                v23 := GenS2{F0: genFn1(100)}
                v24 := GenS2{F0: -688}
                v19 = int64(v3)
                v1 = float32(1.5)
                v19 = v19
                v25 := GenS2{F0: 0}
                if ((((1e10 + 0.1) + ((((0.0 * 0.1) - 0.1) * 0.0) * (((float64(1.5) - -2.25) / ((1e10 * 1e10) * (1.5 + (float64((v1 - (-2.25 - (3.14159 / 0.1)))) + (1e10 / 1.0))))) + ((float64(v1) / 1.0) / 1.5)))) + -2.25) != 0.0) {
                    switch (fi22 + (fi20 % ((0) | 1))) {
                    case -2147483648:
                        v1 = v1
                    case -909:
                        var v26 bool = (false && false)
                        _ = v26
                    case 9223372036854775807:
                        var v27 int64 = (^(v19))
                        _ = v27
                    }
                    var v28 int = (fi20 % (((fi22 + (fi22 ^ (fi20 ^ fi20)))) | 1))
                    v1 += v1
                    fmt.Println(fi22, v1)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v29 := GenS3{F0: "hello", F1: (65535 >= 42), F2: 0}
                    if false {
                        if ("" <= "a") {
                            v28 = fi20
                        } else {
                            v3 += genFn0()
                        }
                        var v30 float64 = ((3.14159 - 1e10) / 3.14159)
                        _ = v30
                    }
                    if false {
                        switch (-(fi22)) {
                        case 256:
                            var v31 int32 = 42
                            _ = v31
                        default:
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = (fi22 / ((v28) | 1))
                                }
                            }
                        }
                        v1 += v1
                    }
                    _ = v28
                    _ = v29
                } else {
                    v19 -= v19
                }
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            v32 := GenS1{F0: (v19 >> ((2) & 63))}
            _ = v32
        } else {
            fmt.Println(v3, v3, v19)
        }
        if (150 == 2) {
            v33 := GenS3{F0: "a", F1: (!(!true)), F2: -868}
            if false {
                v19 -= 1
            } else {
                var v34 uint64 = 18446744073709551615
                _ = v34
            }
            if true {
                fmt.Println(v3)
                var v35 int16 = 32767
                _ = v35
            } else {
                switch (fi20 / (((^(fi20))) | 1)) {
                case 32768:
                    var v36 string = "foo bar"
                    _ = v36
                case -274:
                    v1 += v1
                default:
                    var v37 int8 = -1
                    _ = v37
                }
                if len(v2) > 0 {
                    v2[0] = fi20
                }
                _ = len(v2)
            }
            _ = v33
        }
        if true {
            var v38 int8 = -128
            _ = v38
        } else {
            fmt.Println(v3, v1, fi20)
            v19 += (v19 >> ((0) & 63))
            fmt.Println(v1, fi20, v19)
        }
        _ = fi20
    }
    if ((false && false) || (127 >= 2)) {
        v39 := GenS4{F0: 1000}
        _ = v39
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        v41 := GenS4{F0: 42}
        _ = fi40
        _ = v41
    }
    if len(v2) > 0 {
        v2[0] = 32768
    }
    v3 = genFn0()
    _ = len(v2)
    for fi42 := 0; fi42 < 13; fi42++ {
        v19 = (v19 - v19)
        var v43 string = ("foo bar" + ("a" + "a"))
        if false {
            v44 := GenS1{F0: (v19 << ((739) & 63))}
            _ = len(v2)
            var v45 uint8 = 42
            _ = v44
            _ = v45
        } else {
            v19 += v19
            var v46 float32 = (v1 * 1e10)
            _ = v46
        }
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
