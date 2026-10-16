package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }
func genFn0() uint64 {
    switch 552 {
    case -289:
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{fi6, (^(fi6))}
            v8 := GenS1{F0: 65535}
            if false {
                v9 := GenS1{F0: 42}
                _ = v9
            } else {
                var v10 string = "foo bar"
                v10 = "foo bar"
                _ = v10
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        var v11 float32 = 0.1
        _ = v11
    }
    var v12 int16 = 32767
    _ = v12
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{4294967295, 906, 2}
    var v3 uint64 = 2
    var v4 int16 = -845
    v5 := []uint32{}
    v4 = v4
    var v13 bool = true
    v14 := GenS1{F0: 2}
    var v15 uint64 = v3
    v16 := GenS1{F0: 65535}
    fmt.Println(v15, v4)
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v4, v15, v3)
        _ = fi17
    }
    fmt.Println(v4, v13, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v13 {
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(v3, fi18, v13)
            v19 := GenS1{F0: 255}
            for fi20 := 0; fi20 < 18; fi20++ {
                if v13 {
                    if v13 {
                        v13 = (!(65535 >= 2))
                        v21 := GenS2{F0: (v3 + (v15 - v3)), F1: fi20, F2: 128}
                        if v13 {
                            v13 = v13
                            var v22 uint8 = 0
                            fmt.Println(v13, v4, v3)
                            _ = v22
                        }
                        _ = v21
                    }
                    v3 = 535
                } else {
                    v15 = v15
                    fmt.Println(v13)
                    v1[(fi20 << ((498) & 63))] = fi20
                    v1[1] = (fi18 ^ fi18)
                    var v23 int16 = (v4 + (v4 / ((v4) | 1)))
                    fmt.Println(fi20)
                    _ = v23
                }
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    var v24 uint64 = genFn0()
    var v25 string = ("\n\t" + ("" + ("" + "hello")))
    var v26 uint = 82
    if v13 {
        if v13 {
            if (-1 >= -499) {
                var v27 float64 = ((0.1 * 0.1) * 0.0)
                _ = v27
            } else {
                if v13 {
                    if (("foo bar" + v25) <= (v25 + (v25 + v25))) {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            var v29 int8 = -128
                            switch (fi28 ^ (fi28 - fi28)) {
                            case 128:
                                v13 = v13
                            case 32768:
                                v30 := v1[fi28]
                                _ = v30
                            default:
                                v31 := GenS1{F0: 4294967295}
                                _ = v31
                            }
                            var v32 int = (fi28 - fi28)
                            _ = fi28
                            _ = v29
                            _ = v32
                        }
                        v26 = (v26 << (((^(v26))) & 63))
                    } else {
                        _ = len(v1)
                        v3 += genFn0()
                    }
                }
                for fi33 := 0; fi33 < 2; fi33++ {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        v4 += v4
                        _ = fi34
                    }
                    for fi35 := 0; fi35 < 9; fi35++ {
                        var v36 int16 = int16(v24)
                        fmt.Println(v15)
                        var v37 int = (fi33 ^ fi33)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = fi33
                }
            }
            var v38 uint32 = 2147483648
            _ = v38
        } else {
            switch v25 {
            case "\n\t":
                v3 += v15
            }
            for fi39 := 0; fi39 < 1; fi39++ {
                if (v25 <= (v25 + "foo bar")) {
                    if v13 {
                        _ = len(v2)
                        var v40 int16 = 1000
                        switch 2147483647 {
                        case 9223372036854775807:
                            var v41 bool = true
                            _ = v41
                        case -1:
                            v42 := GenS1{F0: v26}
                            _ = v42
                        case 4294967295:
                            var v43 uint64 = (v24 & genFn0())
                            _ = v43
                        default:
                            v24 = genFn0()
                        }
                        _ = v40
                    }
                    v5 = append(v5, 2)
                    if ((-2.25 * (1.5 * (float32(3.14159) * (3.14159 + ((0.0 - 1.5) * (0.0 / 1e10)))))) < -2.25) {
                        delete(v1, 0)
                    } else {
                        for fi44 := 0; fi44 < 4; fi44++ {
                            v24 -= (v15 & (v15 << (((v26 ^ 255)) & 63)))
                            v15 += (v15 % ((v24) | 1))
                            _ = fi44
                        }
                        v25 = "a"
                    }
                }
                _ = fi39
            }
            _ = len(v5)
        }
    } else {
        var v45 int16 = v4
        _ = v45
    }
    var v46 int32 = 440
    v26 = v26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
