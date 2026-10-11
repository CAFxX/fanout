package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint8, p1 int8) uint64 {
    if (false && (!true)) {
        for fi4 := 0; fi4 < 7; fi4++ {
            var v5 uint64 = 1
            var v6 int = (fi4 / ((fi4) | 1))
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{(fi7 / ((fi7) | 1)), fi7, (-(fi7))}
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 string = ("a" + "a")
            _ = fi9
            _ = v10
        }
        var v11 float64 = ((0.1 * 1.5) * ((float64((1.5 + 0.1)) - 0.1) - float64(0.1)))
        _ = fi7
        _ = v11
        _ = v8
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}
func genFn1(p0 int64, p1 uint16) uint32 {
    p0 = p0
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := []uint32{947, 65535, 4294967295}
    var v2 string = (("" + ("hello" + "")) + ("foo bar" + "\n\t"))
    var v3 int = 32768
    v2 = v2
    switch (v3 - v3) {
    case -789:
        var v12 uint = 42
        _ = v12
    case 4294967295:
        var v13 uint8 = 158
        _ = v13
    case 128:
        v2 = (("foo bar" + v2) + ("a" + v2))
    default:
        var v14 bool = (42 < 9223372036854775807)
        _ = v14
    }
    v3 = (v3 / (((v3 << ((661) & 63))) | 1))
    var v15 float32 = -2.25
    if (!false) {
        if len(v1) > 0 {
            v1[0] = genFn1(30, 0)
        }
        fmt.Println(v15)
    } else {
        var v16 int16 = 32767
        v15 -= float32(3.14159)
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS1{F0: 854}
            for fi19 := 0; fi19 < 14; fi19++ {
                switch (v2 + (v2 + v2)) {
                case "":
                    switch 9223372036854775807 {
                    case -465:
                        v2 = (((v2 + "a") + "\n\t") + v2)
                    case -708:
                        fmt.Println(v15)
                    }
                default:
                    _ = len(v1)
                    v20 := GenS1{F0: -1}
                    _ = v20
                }
                v3 -= (fi19 | 32767)
                _ = fi19
            }
            if (!(128 > v3)) {
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint = 255
                    _ = fi21
                    _ = v22
                }
                switch (v3 * fi17) {
                case 127:
                    v15 -= float32(1.5)
                default:
                    if len(v1) > 0 {
                        v1[0] = 452
                    }
                }
                if (genFn1(42, 537) != 4294967295) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    for fi23 := 0; fi23 < 9; fi23++ {
                        v15 += 0.1
                        _ = fi23
                    }
                    if ((v15 == v15) && false) {
                        v24 := GenS1{F0: -1}
                        for fi25 := 0; fi25 < 3; fi25++ {
                            v26 := GenS2{F0: int8(v3)}
                            _ = len(v1)
                            var v27 uint16 = 0
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                    } else {
                        var v28 int8 = 0
                        for fi29 := 0; fi29 < 18; fi29++ {
                            v15 = float32(-2.25)
                            fmt.Println(v28, fi17)
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                            _ = fi29
                        }
                        _ = v28
                    }
                }
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    fmt.Println(v15, v15, v3)
    v2 = v2
    for fi30 := 0; fi30 < 18; fi30++ {
        _ = len(v1)
        fmt.Println(v3, fi30, v15)
        _ = fi30
    }
    v31 := GenS3{F0: 260}
    v3 -= v3
    for fi32 := 0; fi32 < 17; fi32++ {
        v15 -= float32(-2.25)
        if (!false) {
            _ = len(v1)
            v33 := GenS4{F0: 0}
            _ = v33
        }
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
}
