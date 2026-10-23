package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0() uint8 {
    var v7 uint8 = 42
    var v8 float64 = (0.0 / 1e10)
    var v9 int64 = 889
    var v10 float64 = (0.1 / 3.14159)
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 44
}
func genFn1(p0 int32) int8 {
    var v11 int64 = 1
    v11 -= -9223372036854775808
    var v12 int32 = p0
    var v13 uint8 = 1
    v13 = 1
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 90
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = -626
    v3 := make(map[uint32]int)
    var v4 int32 = 0
    var v5 int64 = 0
    var v6 uint = 2
    fmt.Println(v6, v4, v5)
    v6 = 42
    v4 = v4
    fmt.Println(v6)
    var v14 float32 = 0.1
    v1 -= (v1 - (v1 / ((2) | 1)))
    var v15 string = ""
    if false {
        var v16 bool = true
        if v16 {
            fmt.Println(v6, v16)
        }
        switch 2147483647 {
        case 4294967295:
            v17 := GenS1{F0: (v5 << (((v6 + (v6 % ((42) | 1)))) & 63)), F1: 1}
            _ = v17
        default:
            delete(v3, (v1 + 411))
        }
        _ = v16
    }
    _ = len(v3)
    if ((v4 == v4) || ((v4 ^ (v4 | v4)) != v4)) {
        if (v4 != v4) {
            v18 := v3[v1]
            _ = v18
        } else {
            switch 128 {
            case -9223372036854775808:
                switch (((v15 + ((v15 + v15) + v15)) + (((v15 + v15) + ("\n\t" + ((((v15 + ("foo bar" + ("\n\t" + ("\n\t" + (((v15 + v15) + v15) + "a"))))) + "hello") + "\n\t") + (((v15 + ("a" + v15)) + (v15 + (v15 + v15))) + (v15 + ((v15 + "foo bar") + v15)))))) + v15)) + v15) {
                case "\n\t":
                    v3[v1] = 863
                default:
                    v19 := []int{}
                    _ = v19
                }
            case -121:
                for fi20 := 0; fi20 < 1; fi20++ {
                    fmt.Println(v1)
                    _ = fi20
                }
            }
            var v21 bool = true
            var v22 uint32 = v1
            _ = v21
            _ = v22
        }
        if false {
            for fi23 := 0; fi23 < 6; fi23++ {
                if (-2.25 >= 1e10) {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v15 = (v15 + v15)
                        v25 := GenS1{F0: (v5 & v5), F1: 255}
                        var v26 int32 = v4
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
                for fi27 := 0; fi27 < 18; fi27++ {
                    switch v15 {
                    case "":
                        if (!false) {
                            v2 = (v2 << ((v6) & 63))
                        } else {
                            v2 += v2
                            v14 += v14
                        }
                    case "a":
                        fmt.Println(v6, v1)
                    }
                    if (uint(v1) < (v6 | (v6 + (v6 * v6)))) {
                        var v28 uint64 = 0
                        v2 += v2
                        _ = v28
                    } else {
                        delete(v3, (v1 & (v1 + (^(v1)))))
                        v29 := GenS1{F0: v5, F1: 0}
                        v30 := GenS1{F0: (v5 - (v5 + 42)), F1: 0}
                        _ = v29
                        _ = v30
                    }
                    var v31 int64 = 749
                    _ = fi27
                    _ = v31
                }
                _ = fi23
            }
            v32 := []int{4294967295}
            _ = v32
        } else {
            delete(v3, 2)
            v33 := GenS1{F0: (-(v5)), F1: 335}
            var v34 uint64 = 4294967295
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            if (255 != 255) {
                v3[v1] = -2147483648
            } else {
                var v36 uint32 = (v1 & uint32(v6))
                v37 := []int{fi35, (fi35 * fi35), 1, fi35}
                _ = v36
                _ = v37
            }
            v6 -= (v6 + v6)
            _ = fi35
        }
    } else {
        for fi38 := 0; fi38 < 7; fi38++ {
            v39 := GenS2{F0: 255, F1: true, F2: fi38}
            var v40 string = "a"
            switch (v40 + "\n\t") {
            case "a":
                fmt.Println(v4)
            case "":
                v6 = v6
            case "hello":
                if (!((v5 & (v5 % (((v5 >> (((v6 ^ (v6 % ((42) | 1)))) & 63))) | 1))) == 355)) {
                    var v41 uint8 = 2
                    _ = v41
                }
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        v42 := GenS3{F0: int64(v4), F1: (-(v4))}
        _ = v42
    }
    fmt.Println(v2)
    v43 := v3[(v1 + (v1 << ((719) & 63)))]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
