package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint16 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    var v4 int8 = 1
    v5 := make(map[uint32]bool)
    var v6 float32 = (1e10 * 3.14159)
    v4 = (^(v4))
    if ((0 >= -864) && false) {
        if (false || ((float64(float32(v1)) + (v1 + 1e10)) < float64(3.14159))) {
            var v7 uint64 = 4294967295
            var v8 int = 720
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS1{F0: 0}
                fmt.Println(v4, fi9, v4)
                if (!false) {
                    var v11 string = ("a" + "foo bar")
                    v8 = 128
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
            _ = v7
            _ = v8
        }
        v12 := GenS2{F0: 2, F1: 32767, F2: 65535}
        _ = v12
    } else {
        var v13 uint32 = 482
        if true {
            for fi14 := 0; fi14 < 18; fi14++ {
                if false {
                    v13 = 1
                    fmt.Println(fi14, fi14, fi14)
                }
                switch (fi14 >> ((994) & 63)) {
                case 127:
                    v15 := []int{772}
                    _ = v15
                case 625:
                    delete(v2, v13)
                default:
                    v16 := []int{fi14, (fi14 >> ((46) & 63)), 32768}
                    _ = v16
                }
                _ = fi14
            }
            v4 = v4
        } else {
            v3 = 1
            for fi17 := 0; fi17 < 20; fi17++ {
                if (v13 == 65535) {
                    if true {
                        v18 := GenS2{F0: 23, F1: 2, F2: 2}
                        var v19 uint64 = 0
                        v20 := []int{fi17, (fi17 >> ((4294967295) & 63)), fi17}
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 string = "foo bar"
                        var v22 uint32 = (v13 << ((0) & 63))
                        if (true || false) {
                            v1 -= v1
                        }
                        _ = v21
                        _ = v22
                    }
                    v23 := []int{fi17, int(v4), (fi17 | -1), fi17}
                    v24 := GenS1{F0: 2147483647}
                    _ = v23
                    _ = v24
                }
                _ = fi17
            }
            if (true && (0 < 2)) {
                fmt.Println(v1)
                if ((("\n\t" >= "\n\t") && false) && (504 <= -9223372036854775808)) {
                    v25 := GenS1{F0: 1}
                    switch -216 {
                    case -868:
                        var v26 float64 = (((v1 + (v1 - v1)) / v1) / (float64(1.5) / float64((((v6 / 1e10) + v6) * (1.5 - (v6 * 3.14159))))))
                        _ = v26
                    case 2147483647:
                        var v27 uint32 = v13
                        _ = v27
                    case -605:
                        switch 32767 {
                        case 2147483647:
                            fmt.Println(v3, v1, v13)
                        case 106:
                            var v28 uint16 = 1
                            _ = v28
                        default:
                            var v29 int8 = v4
                            _ = v29
                        }
                    default:
                        var v30 float32 = 0.1
                        _ = v30
                    }
                    _ = v25
                } else {
                    v31 := GenS1{F0: -2147483648}
                    _ = v31
                }
            } else {
                var v32 int16 = v3
                v1 = v1
                _ = v32
            }
        }
        _ = v13
    }
    var v33 float64 = v1
    v4 -= v4
    var v34 int16 = v3
    v35 := GenS1{F0: -2147483648}
    v6 -= ((0.1 * ((3.14159 - (v6 - ((v6 + (-2.25 + (-2.25 * 3.14159))) * (((v6 - v6) / v6) - v6)))) * v6)) + v6)
    var v36 string = ""
    v37 := []int{32768, 32767, 128}
    var v38 float32 = 0.0
    fmt.Println(v1, v33)
    v39 := v5[42]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
