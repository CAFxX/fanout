package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]uint32)
    var v3 bool = (42 == -970)
    v4 := []uint32{v1}
    var v5 uint16 = 255
    v6 := make(map[uint32]bool)
    switch 258 {
    case 2147483647:
        v7 := GenS1{F0: 42, F1: -128, F2: false}
        _ = v7
    case 569:
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    delete(v2, 881)
    v9 := v6[uint32(v5)]
    v3 = (v3 && true)
    var v10 int64 = -937
    for fi11 := 0; fi11 < 4; fi11++ {
        delete(v2, (-(fi11)))
        var v12 uint64 = 2
        _ = fi11
        _ = v12
    }
    v9 = ((v10 << ((1) & 63)) >= 192)
    if (!v3) {
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 float64 = (0.0 - (0.0 / (1e10 + 0.1)))
                if (!v9) {
                    var v16 int = fi14
                    var v17 float64 = (v15 / -2.25)
                    switch 870 {
                    case -319:
                        if (((1 != 42) || true) && (!(1.5 > (0.1 * 1e10)))) {
                            v6[(v1 ^ (^(v1)))] = v9
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                            v2[v16] = (v1 % ((v1) | 1))
                            v18 := GenS2{F0: 0}
                            _ = v18
                        }
                    case 2147483647:
                        var v19 float64 = 1e10
                        _ = v19
                    default:
                        v15 = v17
                    }
                    v3 = (3.14159 < -2.25)
                    if v3 {
                        var v20 uint16 = v5
                        fmt.Println(v3, fi14, v3)
                        _ = v20
                    } else {
                        var v21 float64 = float64(((1e10 * (0.1 - 0.1)) * 0.0))
                        _ = v21
                    }
                    _ = v16
                    _ = v17
                } else {
                    var v22 uint16 = (v5 / ((255) | 1))
                    v23 := GenS3{F0: v5, F1: (1 >= 42)}
                    var v24 int16 = 32767
                    _ = v22
                    _ = v23
                    _ = v24
                }
                fmt.Println(v15, v1)
                _ = fi14
                _ = v15
            }
            fmt.Println(v1, v10, fi13)
            var v25 string = "foo bar"
            _ = fi13
            _ = v25
        }
        if v9 {
            v5 = (v5 / (((v5 | (v5 / ((42) | 1)))) | 1))
        }
        var v26 uint64 = 825
        _ = v26
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 bool = true
            for fi29 := 0; fi29 < 17; fi29++ {
                v3 = (!(((0.0 / 1e10) - ((-2.25 / 3.14159) + (-2.25 + 0.1))) <= 1e10))
                if v9 {
                    v5 = (v5 & 65535)
                    v3 = v28
                    var v30 string = (("" + "a") + ((("foo bar" + "\n\t") + "\n\t") + (((("" + "\n\t") + "a") + "a") + "")))
                    v31 := GenS4{F0: 42}
                    _ = v30
                    _ = v31
                } else {
                    fmt.Println(v10)
                }
                delete(v2, -904)
                _ = fi29
            }
            var v32 int16 = 519
            _ = fi27
            _ = v28
            _ = v32
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            switch (("a" + "") + ((("hello" + "hello") + "") + ("a" + "a"))) {
            case "":
                switch "" {
                case "hello":
                    if ((v5 < (v5 ^ (v5 + 2))) && v9) {
                        v10 += v10
                    }
                }
            case "foo bar":
                var v34 uint = 4294967295
                _ = v34
            case "hello":
                fmt.Println(v10, fi33)
            }
            v9 = (-743 >= -561)
            _ = fi33
        }
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        v9 = (((fi35 | fi35) != 127) && true)
        fmt.Println(v10)
        v36 := GenS4{F0: -1}
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
