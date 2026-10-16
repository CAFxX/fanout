package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0(p0 string, p1 int32, p2 uint16) int32 {
    if ((0 > 127) && (!false)) {
        var v6 uint64 = 1
        var v7 int64 = 490
        _ = v6
        _ = v7
    }
    var v8 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}
func genFn1(p0 int) string {
    v9 := []int{32767, (p0 / ((-9223372036854775808) | 1)), p0}
    var v10 uint16 = 1
    _ = p0
    _ = v10
    _ = v9
    return (("a" + "\n\t") + "\n\t")
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]uint32)
    var v3 float64 = (0.0 / -2.25)
    var v4 string = "a"
    v5 := []uint32{669, 2147483647, 832}
    if false {
        v1 = (v1 + (v1 * v1))
    }
    v4 = "a"
    switch "\n\t" {
    case "a":
        var v11 uint = 0
        _ = v11
    default:
        if (((!true) && (0.1 == 1e10)) && true) {
            v12 := GenS1{F0: -610, F1: (^(v1))}
            v13 := GenS1{F0: 1000, F1: v1}
            if false {
                var v14 uint64 = 4294967295
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS1{F0: 0, F1: v1}
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v4, v4, v4)
    fmt.Println(v1, v1)
    v17 := GenS1{F0: -32768, F1: (-(v1))}
    var v18 float64 = 1.5
    _ = len(v2)
    switch 2147483647 {
    case 218:
        if false {
            if false {
                v3 = (-2.25 + float64(1e10))
            } else {
                var v19 float32 = 1.5
                v1 = v1
                _ = v19
            }
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 uint = 1
                v18 -= v18
                _ = fi20
                _ = v21
            }
        }
    }
    switch 127 {
    case 254:
        switch v4 {
        case "foo bar":
            var v22 float32 = 0.0
            _ = v22
        default:
            switch v4 {
            case "":
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := v2[(fi23 % (((fi23 | (^(fi23)))) | 1))]
                    _ = fi23
                    _ = v24
                }
            case "foo bar":
                if (680 < -32768) {
                    if true {
                        delete(v2, 32767)
                        for fi25 := 0; fi25 < 9; fi25++ {
                            v1 += -128
                            v4 = (v4 + "foo bar")
                            _ = fi25
                        }
                        var v26 int = 32767
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 9; fi27++ {
                            var v28 uint8 = 127
                            _ = fi27
                            _ = v28
                        }
                    }
                }
            case "a":
                if (!true) {
                    v1 = (v1 / ((v1) | 1))
                }
            default:
                switch (v4 + v4) {
                case "foo bar":
                    for fi29 := 0; fi29 < 7; fi29++ {
                        var v30 uint32 = 618
                        v31 := GenS1{F0: -32768, F1: 61}
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case "":
                    var v32 uint8 = 13
                    _ = v32
                }
            }
        }
    default:
        if (true || (!false)) {
            var v33 string = "a"
            var v34 bool = (v3 < 1.5)
            var v35 uint = 39
            _ = v33
            _ = v34
            _ = v35
        }
    }
    var v36 int16 = -70
    var v37 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
