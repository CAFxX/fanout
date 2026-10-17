package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 int, p2 uint) int {
    fmt.Println(p2, p2, p1)
    v4 := GenS1{F0: 42}
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{fi5, (fi5 ^ (fi5 | fi5)), fi5, -1}
        _ = fi5
        _ = v6
    }
    fmt.Println(p1, p0, p0)
    if true {
        fmt.Println(p2, p0)
    } else {
        if false {
            v7 := GenS1{F0: 746}
            v8 := GenS1{F0: 42}
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p1 >> ((2) & 63))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 255
    var v3 float32 = (((1.5 * 3.14159) / (-2.25 / 1.0)) - 3.14159)
    for fi9 := 0; fi9 < 12; fi9++ {
        v2 -= v2
        v3 += v3
        fmt.Println(fi9, v3)
        _ = fi9
    }
    fmt.Println(v3, v3)
    v10 := []int{-28, 2147483647, 128, 127}
    for fi11 := 0; fi11 < 12; fi11++ {
        _ = len(v10)
        var v12 uint16 = (v2 % ((v2) | 1))
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v2)
    for fi13 := 0; fi13 < 1; fi13++ {
        delete(v1, 2)
        switch fi13 {
        case -165:
            v10 = append(v10, -317)
        case -357:
            if false {
                delete(v1, 2147483648)
            } else {
                if (false && (-1 < -102)) {
                    v2 = 996
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 bool = (true && (!false))
                        var v16 bool = v15
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v1[42] = "\n\t"
                } else {
                    v17 := GenS2{F0: 2147483647}
                    if (479 == 42) {
                        v2 = v2
                        var v18 int32 = 1
                        var v19 string = "hello"
                        _ = v18
                        _ = v19
                    } else {
                        var v20 string = ("\n\t" + ("foo bar" + "foo bar"))
                        _ = v20
                    }
                    _ = v17
                }
                if false {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        var v22 string = ("hello" + ((("foo bar" + "") + "") + "hello"))
                        if (((0.0 / -2.25) == 1.5) || true) {
                            var v23 uint64 = 236
                            fmt.Println(fi13)
                            _ = v23
                        } else {
                            fmt.Println(v22, v22)
                            if (true && true) {
                                if true {
                                    switch fi21 {
                                    case 637:
                                        v2 = (v2 & v2)
                                    default:
                                        switch (fi13 % ((fi13) | 1)) {
                                        case 2147483647:
                                            v2 = (v2 << ((1) & 63))
                                        case 187:
                                            v2 += (v2 ^ 1)
                                        }
                                    }
                                    if (!false) {
                                        fmt.Println(v22, v22)
                                        if ((-2.25 + ((-2.25 / (((0.1 / (0.0 + float64(float32(-2.25)))) - (1e10 + 1.5)) + 1.5)) + 1.5)) < (1e10 * 3.14159)) {
                                            v24 := GenS1{F0: 32767}
                                            v25 := GenS1{F0: v2}
                                            _ = v24
                                            _ = v25
                                        } else {
                                            _ = len(v10)
                                        }
                                    }
                                }
                                var v26 bool = ((-32768 > -32768) && false)
                                v10 = append(v10, fi21)
                                _ = v26
                            } else {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                        }
                        _ = fi21
                        _ = v22
                    }
                    var v27 bool = false
                    _ = v27
                } else {
                    var v28 uint32 = 549
                    if ((526 != (v2 >> ((0) & 63))) && (("a" + "a") != "foo bar")) {
                        if true {
                            delete(v1, v28)
                            for fi29 := 0; fi29 < 9; fi29++ {
                                fmt.Println(v2, fi13, fi13)
                                _ = fi29
                            }
                        } else {
                            var v30 int64 = 9223372036854775807
                            _ = v30
                        }
                    }
                    _ = v28
                }
            }
        case 32768:
            if false {
                var v31 int = 1
                var v32 uint16 = (v2 & (v2 / ((v2) | 1)))
                _ = v31
                _ = v32
            } else {
                fmt.Println(v2, fi13)
                v2 = (v2 & v2)
                var v33 float32 = 3.14159
                _ = v33
            }
        }
        _ = fi13
    }
    switch "a" {
    case "hello":
        v2 += 1
    default:
        if false {
            if (128 > 1) {
                v34 := GenS1{F0: (v2 | (v2 >> ((255) & 63)))}
                for fi35 := 0; fi35 < 4; fi35++ {
                    v36 := GenS2{F0: -1}
                    _ = fi35
                    _ = v36
                }
                _ = v34
            }
            v3 = (v3 + ((v3 - (v3 - v3)) * v3))
        } else {
            v3 -= v3
        }
    }
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
