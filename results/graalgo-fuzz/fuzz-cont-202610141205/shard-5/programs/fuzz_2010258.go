package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint64; F2 bool }

func main() {
    v1 := []uint32{}
    var v2 uint16 = 2
    var v3 bool = true
    var v4 int8 = -14
    v5 := []uint32{}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (!v3)
    if v3 {
        v3 = (uint32(v2) < 2147483648)
        if ((42 < 0) && v3) {
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
            v6 := GenS1{F0: ("" + ("hello" + ("a" + ("\n\t" + "a"))))}
            _ = v6
        }
    }
    switch ("" + "\n\t") {
    case "\n\t":
        if v3 {
            if v3 {
                v3 = ((42 != 851) || (42 >= -813))
                if false {
                    if len(v5) > 0 {
                        v5[0] = 2147483648
                    }
                } else {
                    var v7 uint64 = 42
                    v8 := GenS2{F0: (^(v7))}
                    _ = v7
                    _ = v8
                }
                v9 := GenS1{F0: ("a" + "foo bar")}
                _ = v9
            } else {
                var v10 int8 = v4
                fmt.Println(v3, v3, v4)
                _ = v10
            }
            v11 := GenS3{F0: 2147483647}
            if len(v5) > 0 {
                v5[0] = 42
            }
            _ = v11
        } else {
            fmt.Println(v2, v3)
            var v12 int64 = 9223372036854775807
            _ = v12
        }
    case "a":
        switch 4294967295 {
        case 228:
            for fi13 := 0; fi13 < 19; fi13++ {
                v4 += v4
                _ = fi13
            }
        }
    case "hello":
        var v14 uint32 = 893
        _ = v14
    default:
        var v15 int = 143
        _ = v15
    }
    v4 += (v4 / (((-(v4))) | 1))
    v16 := GenS3{F0: -380}
    v17 := GenS1{F0: ("\n\t" + "a")}
    v2 -= v2
    fmt.Println(v2, v2, v3)
    var v18 int = 2147483647
    if (!(0.1 <= 0.1)) {
        v3 = true
    } else {
        if v3 {
            if v3 {
                v2 += 770
            } else {
                for fi19 := 0; fi19 < 11; fi19++ {
                    v3 = v3
                    for fi20 := 0; fi20 < 20; fi20++ {
                        _ = len(v5)
                        v18 -= (v18 >> ((42) & 63))
                        fmt.Println(v4, v2)
                        _ = fi20
                    }
                    if (!v3) {
                        v21 := GenS1{F0: "foo bar"}
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v18 = fi22
                            _ = fi22
                        }
                        for fi23 := 0; fi23 < 3; fi23++ {
                            if ((654 > (v2 | v2)) && (!((766 <= 1) && (fi19 > (fi23 << ((824) & 63)))))) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v3 = v3
                            }
                            v18 = -1
                            var v24 int64 = -1
                            v2 += 42
                            var v25 uint8 = 1
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v21
                    } else {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            var v27 int32 = 0
                            v2 = (v2 % ((v2) | 1))
                            if (1000 == 940) {
                                if (!v3) {
                                    switch "hello" {
                                    case "":
                                        v27 -= (^(v27))
                                    case "\n\t":
                                        fmt.Println(v3)
                                    case "hello":
                                        v3 = (!(v3 && false))
                                    }
                                } else {
                                    v3 = (!((((0.1 + (0.1 + float32((0.0 - 3.14159)))) - (float32(1.5) - (1.5 / 1.5))) + 0.0) > 0.1))
                                    fmt.Println(v27)
                                }
                                v4 -= (v4 ^ (v4 / ((42) | 1)))
                            } else {
                                var v28 float32 = 1e10
                                _ = v28
                            }
                            _ = fi26
                            _ = v27
                        }
                    }
                    _ = fi19
                }
                v29 := GenS4{F0: "\n\t", F1: 18446744073709551615, F2: (1 != 898)}
                v18 += (v18 | (v18 ^ 32767))
                _ = v29
            }
        }
    }
    v4 += -128
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 bool = v3
        switch (("hello" + ("a" + ("\n\t" + ("a" + "hello")))) + (("\n\t" + "") + "foo bar")) {
        case "hello":
            v4 = 1
        case "a":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            if v31 {
                v2 = (v2 % (((v2 + v2)) | 1))
                var v32 uint8 = 184
                _ = v32
            }
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 uint8 = 2
            v1 = append(v1, 716)
            _ = fi33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v2 = v2
        var v36 float32 = float32((((((0.0 * 1e10) / 1.5) / (1.5 / (1.5 * 1.5))) + 1.5) * 1e10))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi35
        _ = v36
    }
    fmt.Println(v2, v2, v18)
    v37 := GenS2{F0: 42}
    var v38 uint = 518
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
