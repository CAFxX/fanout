package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0() uint {
    v4 := GenS1{F0: 40, F1: -2147483648}
    v5 := []int{}
    var v6 float64 = float64((1e10 - 3.14159))
    fmt.Println(v6)
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{}
    v3 := make(map[int]int)
    v7 := v3[-9223372036854775808]
    v1 = ((float64((-2.25 + (1e10 * 1.5))) * 1.5) / 1.0)
    v1 += (3.14159 * v1)
    v1 = v1
    switch 1 {
    case 128:
        fmt.Println(v7)
    case -903:
        if (("foo bar" + "hello") == "a") {
            var v8 uint64 = 0
            for fi9 := 0; fi9 < 10; fi9++ {
                if (312 == 42) {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        switch "" {
                        case "\n\t":
                            fmt.Println(v8)
                        case "foo bar":
                            var v11 int32 = -1
                            _ = v11
                        case "":
                            if (518 <= 1) {
                                v7 -= (v7 >> ((788) & 63))
                                var v12 int8 = -53
                                v8 = 130
                                _ = v12
                            } else {
                                v1 -= float64(1e10)
                                var v13 uint32 = 42
                                if ((((v1 / float64(float32(1.5))) - 1e10) <= (0.0 + float64(-2.25))) || (!(0.0 != v1))) {
                                    v14 := GenS2{F0: 109, F1: 623}
                                    if false {
                                        var v15 uint8 = 195
                                        v13 = 0
                                        var v16 float64 = ((v1 * v1) / float64(0.1))
                                        _ = v15
                                        _ = v16
                                    } else {
                                        v17 := GenS2{F0: 64, F1: 2}
                                        v8 -= (v8 - (^(v8)))
                                        _ = v17
                                    }
                                    _ = v14
                                }
                                _ = v13
                            }
                        }
                        v18 := GenS1{F0: 18446744073709551615, F1: -1}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi10
                        _ = v18
                    }
                    v1 += (v1 + 1e10)
                }
                var v19 uint64 = (v8 % ((v8) | 1))
                v2 = append(v2, 600)
                _ = fi9
                _ = v19
            }
            if (false && (v8 <= v8)) {
                v20 := GenS2{F0: 0, F1: 255}
                switch "foo bar" {
                case "hello":
                    if false {
                        _ = len(v3)
                        v21 := GenS1{F0: (v8 << ((255) & 63)), F1: -666}
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                        _ = v21
                    } else {
                        if false {
                            var v22 int16 = -1
                            if true {
                                var v23 int = (v7 + v7)
                                var v24 int32 = 2147483647
                                _ = v23
                                _ = v24
                            } else {
                                _ = len(v2)
                            }
                            v8 = 18446744073709551615
                            _ = v22
                        }
                        var v25 string = "foo bar"
                        _ = v25
                    }
                case "foo bar":
                    v7 = v7
                case "a":
                    v1 += v1
                default:
                    var v26 int16 = 621
                    _ = v26
                }
                _ = v20
            }
            _ = v8
        } else {
            var v27 uint64 = 941
            fmt.Println(v27)
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v7, v7)
            switch (^(fi28)) {
            case 32768:
                for fi29 := 0; fi29 < 10; fi29++ {
                    fmt.Println(fi29, fi28)
                    _ = fi29
                }
            case 127:
                var v30 int8 = 42
                _ = v30
            }
            _ = fi28
        }
    }
    switch (("" + ("\n\t" + "foo bar")) + "foo bar") {
    case "":
        v1 = ((float64(3.14159) * v1) - 0.0)
    default:
        v31 := GenS3{F0: 1, F1: 732}
        _ = v31
    }
    v1 -= v1
    v1 += (v1 + 1e10)
    v7 = v7
    _ = len(v2)
    delete(v3, (v7 - 2147483647))
    var v32 int64 = -1
    if false {
        var v33 uint16 = 32767
        v34 := GenS4{F0: -1, F1: 2147483647, F2: 2}
        _ = v33
        _ = v34
    }
    v35 := GenS3{F0: 1, F1: 42}
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 int = 4294967295
        v37 = 32767
        _ = fi36
        _ = v37
    }
    var v38 uint8 = 42
    if true {
        for fi39 := 0; fi39 < 3; fi39++ {
            var v40 uint32 = 123
            _ = fi39
            _ = v40
        }
    } else {
        if false {
            v1 = v1
            if len(v2) > 0 {
                v2[0] = 42
            }
        }
    }
    v41 := GenS3{F0: 18446744073709551615, F1: 551}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v38
    _ = v41
    _ = v7
}
