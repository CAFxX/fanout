package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint, p1 int64) int64 {
    var v7 uint32 = 0
    var v8 int32 = 0
    var v9 int = 32768
    if false {
        v7 = 65535
    } else {
        v10 := GenS1{F0: 1}
        for fi11 := 0; fi11 < 1; fi11++ {
            v7 = (v7 >> ((p0) & 63))
            p1 += (p1 % (((p1 - (p1 ^ (^(p1))))) | 1))
            _ = fi11
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 float32 = ((0.1 * float32((float64((-2.25 / ((((0.0 + 3.14159) + float32(0.1)) * 3.14159) - 1e10))) / ((0.1 * 0.1) * 3.14159)))) * 1.5)
    var v2 int = 128
    var v3 uint32 = 0
    v4 := make(map[uint32]int)
    var v5 uint64 = 2
    v6 := make(map[uint32]uint32)
    v12 := GenS1{F0: (v5 ^ (v5 >> ((255) & 63)))}
    if (!false) {
        if true {
            delete(v4, (^(v3)))
            v3 = 431
        } else {
            switch "hello" {
            case "foo bar":
                v3 = v3
            case "a":
                if (!(!true)) {
                    v3 += 2147483647
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v5 += v5
                        v14 := GenS1{F0: (^(v5))}
                        if true {
                            for fi15 := 0; fi15 < 5; fi15++ {
                                v2 -= -1
                                _ = fi15
                            }
                        } else {
                            var v16 float64 = 1e10
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                    if ((^(v3)) > 65535) {
                        fmt.Println(v3, v3)
                    } else {
                        var v17 float64 = (3.14159 - 0.1)
                        var v18 uint16 = 84
                        if false {
                            v19 := []int{9223372036854775807, (v2 << ((884) & 63)), 4294967295, int(v18)}
                            v1 = ((v1 + v1) + 0.0)
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 11; fi20++ {
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    var v22 uint16 = 255
                                    switch ("foo bar" + "hello") {
                                    case "hello":
                                        fmt.Println(v17, v2, fi21)
                                    case "foo bar":
                                        fmt.Println(v1, v2, v1)
                                    }
                                    var v23 string = "a"
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                var v24 string = ("" + ("" + ("foo bar" + "a")))
                                var v25 uint8 = 1
                                _ = fi20
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS1{F0: v5}
                            var v27 uint16 = 88
                            _ = v26
                            _ = v27
                        }
                        _ = v17
                        _ = v18
                    }
                }
            case "hello":
                v28 := GenS1{F0: (v5 & (^(v5)))}
                _ = v28
            default:
                var v29 int = v2
                _ = v29
            }
            v3 = (v3 >> ((255) & 63))
            var v30 int8 = int8(v5)
            _ = v30
        }
    } else {
        var v31 bool = (!true)
        _ = v31
    }
    v32 := []int{(^(v2))}
    v2 -= (v2 - 128)
    var v33 int64 = 1
    if true {
        v4[v3] = 2147483647
        var v34 int32 = 2147483647
        _ = v34
    }
    v5 -= v5
    v1 = v1
    v35 := v4[(v3 / (((v3 & (v3 / ((uint32(v5)) | 1)))) | 1))]
    v1 += 3.14159
    v5 += v5
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    v5 += uint64(v3)
    switch ((("" + "a") + "a") + ("foo bar" + "a")) {
    case "a":
        v5 -= 1
    default:
        var v36 float64 = 1e10
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
