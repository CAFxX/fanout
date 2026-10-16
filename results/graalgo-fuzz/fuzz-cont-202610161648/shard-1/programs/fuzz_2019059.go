package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int8 = 15
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v1 = v1
    var v4 uint32 = 92
    var v5 int8 = (v1 * v1)
    var v6 uint32 = v4
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 932}
    v5 -= v5
    if false {
        v9 := []int{389, 127, 127}
        var v10 float64 = 1e10
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int32 = -1
        _ = fi11
        _ = v12
    }
    v13 := []int{127, 1, 598}
    if (127 <= 2) {
        if (3.14159 > 0.1) {
            if false {
                var v14 float32 = float32(1e10)
                v15 := GenS2{F0: (v1 >> ((129) & 63))}
                _ = v14
                _ = v15
            }
            if true {
                switch 944 {
                case -28:
                    if false {
                        v2[327] = v4
                        var v16 int = -336
                        v16 = 42
                        _ = v16
                    }
                case 127:
                    fmt.Println(v5)
                case -551:
                    if false {
                        v17 := GenS2{F0: v5}
                        if true {
                            fmt.Println(v5, v1, v6)
                            if ((18446744073709551615 <= uint(v4)) || (871 <= 539)) {
                                v18 := GenS1{F0: 0}
                                v19 := GenS1{F0: 0}
                                v2[32767] = (v6 * (v4 | v6))
                                _ = v18
                                _ = v19
                            }
                            v7 += (v7 >> ((343) & 63))
                        } else {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                v1 += v1
                                _ = fi20
                            }
                            v21 := v3[128]
                            v6 -= (v4 | v4)
                            _ = v21
                        }
                        switch -1 {
                        case 4294967295:
                            switch "a" {
                            case "\n\t":
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    switch (fi22 ^ fi22) {
                                    case 4294967295:
                                        fmt.Println(v1)
                                    case -1:
                                        v5 = (^(v1))
                                    case 127:
                                        if (42 >= 42) {
                                            if true {
                                                v23 := GenS1{F0: 729}
                                                var v24 uint64 = 1
                                                _ = v23
                                                _ = v24
                                            } else {
                                                _ = len(v13)
                                                delete(v2, (fi22 % (((^(fi22))) | 1)))
                                            }
                                            var v25 uint8 = 221
                                            for fi26 := 0; fi26 < 14; fi26++ {
                                                if (-1 <= 42) {
                                                    var v27 int16 = 32767
                                                    v7 += (v7 << ((42) & 63))
                                                    v28 := v3[32768]
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    v7 -= (v7 % ((v7) | 1))
                                                }
                                                _ = fi26
                                            }
                                            _ = v25
                                        }
                                    }
                                    var v29 uint = 42
                                    var v30 float64 = 3.14159
                                    _ = fi22
                                    _ = v29
                                    _ = v30
                                }
                            case "a":
                                var v31 uint16 = 646
                                _ = v31
                            case "":
                                var v32 string = ("" + ((("a" + (((("hello" + ("" + "\n\t")) + "") + "a") + "a")) + "foo bar") + ("a" + "hello")))
                                _ = v32
                            }
                        case 9223372036854775807:
                            v5 -= v5
                        default:
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        }
                        _ = v17
                    } else {
                        v33 := GenS2{F0: (v5 + (v1 >> ((4294967295) & 63)))}
                        _ = v33
                    }
                default:
                    var v34 float32 = -2.25
                    _ = v34
                }
            }
        }
        v4 += (v6 >> ((42) & 63))
        for fi35 := 0; fi35 < 11; fi35++ {
            for fi36 := 0; fi36 < 5; fi36++ {
                _ = len(v3)
                var v37 uint = 4294967295
                _ = fi36
                _ = v37
            }
            var v38 int = (-(fi35))
            v7 = v7
            _ = fi35
            _ = v38
        }
    } else {
        if (-2.25 == 3.14159) {
            for fi39 := 0; fi39 < 6; fi39++ {
                v3[2147483647] = false
                _ = fi39
            }
        }
        v40 := GenS1{F0: 762}
        _ = v40
    }
    fmt.Println(v6, v4)
    var v41 uint8 = 0
    var v42 uint = 356
    v41 += v41
    var v43 int8 = 18
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
