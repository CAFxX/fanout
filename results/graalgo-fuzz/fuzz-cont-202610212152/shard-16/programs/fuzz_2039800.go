package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int64, p1 string) uint64 {
    var v5 uint32 = 42
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    return 383
}
func genFn1(p0 int16, p1 string, p2 uint8) int16 {
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 += (p0 + (^(p0)))
        var v8 int64 = -9223372036854775808
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 ^ -1)
}

func main() {
    v1 := []uint32{}
    var v2 string = ""
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    fmt.Println(v2)
    if true {
        if true {
            v2 = "\n\t"
            var v10 uint = 107
            v11 := GenS1{F0: -107}
            _ = v10
            _ = v11
        }
    } else {
        v2 = v2
        var v12 int8 = -128
        if true {
            v2 = v2
            v13 := GenS1{F0: v12}
            _ = v13
        } else {
            v12 -= 2
            if true {
                var v14 float32 = 0.1
                _ = v14
            }
            v2 = (v2 + "a")
        }
        _ = v12
    }
    v2 = "foo bar"
    var v15 uint16 = 255
    var v16 int16 = -32768
    for fi17 := 0; fi17 < 20; fi17++ {
        v15 += 824
        for fi18 := 0; fi18 < 16; fi18++ {
            v16 -= -1
            var v19 int64 = 42
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v2 = (v2 + v2)
    v4[128] = 0
    var v20 uint64 = 1
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            if true {
                v23 := GenS2{F0: 255}
                v16 -= (v16 & (v16 << ((1) & 63)))
                if false {
                    _ = len(v3)
                    var v24 int32 = -1
                    _ = v24
                }
                _ = v23
            } else {
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch (fi21 - 42) {
                    case 15:
                        switch 32767 {
                        case 4294967295:
                            var v25 float64 = ((3.14159 + 3.14159) / (0.0 - 3.14159))
                            _ = v25
                        case 2147483647:
                            v26 := v3[fi21]
                            _ = v26
                        case 128:
                            switch 128 {
                            case 559:
                                v2 = ("a" + "hello")
                            case 0:
                                if false {
                                    v27 := GenS3{F0: 42}
                                    v28 := GenS4{F0: 0}
                                    _ = v27
                                    _ = v28
                                }
                            case -1:
                                if false {
                                    var v29 string = v2
                                    v30 := GenS1{F0: 42}
                                    _ = v29
                                    _ = v30
                                }
                            default:
                                v15 += (^(v15))
                            }
                        default:
                            if (1 > 128) {
                                var v31 uint = 18446744073709551615
                                v32 := GenS1{F0: -45}
                                _ = v31
                                _ = v32
                            }
                        }
                    case 32768:
                        _ = len(v1)
                    case -2147483648:
                        var v33 uint32 = 4294967295
                        _ = v33
                    }
                }
                var v34 int8 = 1
                _ = v34
            }
            _ = fi22
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
