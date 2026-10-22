package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint; F2 bool }
func genFn0(p0 uint) int16 {
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p0)
    case "hello":
        v4 := []int{714, -901, 692, 32768}
        _ = v4
    default:
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 uint64 = 296
            var v7 float32 = -2.25
            v8 := GenS1{F0: 271, F1: 0, F2: "foo bar"}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    p0 -= (p0 >> (((p0 ^ (p0 | 2))) & 63))
    _ = p0
    return 796
}
func genFn1() int {
    var v9 int8 = 0
    v10 := GenS1{F0: 150, F1: 312, F2: "a"}
    fmt.Println(v9)
    fmt.Println(v9)
    var v11 uint = 65535
    _ = v10
    _ = v11
    _ = v9
    return 452
}

func main() {
    var v1 bool = (true && ((false || (634 < 2)) || (true || (18446744073709551615 != 0))))
    var v2 uint32 = 42
    v3 := make(map[string]int)
    var v12 float64 = ((0.1 - (((1.5 / (float64(float32(3.14159)) * 1.5)) + (1e10 - -2.25)) * 1.5)) + (float64(1e10) + 1e10))
    switch 9223372036854775807 {
    case 2147483647:
        switch "hello" {
        case "hello":
            v1 = ("\n\t" == "\n\t")
        case "\n\t":
            v13 := GenS1{F0: 885, F1: v2, F2: "hello"}
            _ = v13
        }
    default:
        switch "foo bar" {
        case "hello":
            var v14 uint32 = v2
            _ = v14
        default:
            v2 = 1
        }
    }
    v2 = (v2 / (((^(v2))) | 1))
    var v15 float32 = ((1e10 - ((1.5 - 0.1) / (1.5 - -2.25))) + (-2.25 * 3.14159))
    if true {
        fmt.Println(v2)
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int32 = 2147483647
            var v18 float32 = ((float32(float64((v15 / 1.0))) - v15) - 3.14159)
            fmt.Println(v2)
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            var v19 float32 = float32((float64(1.5) * (float64(v15) / (v12 - ((((3.14159 / float64(v15)) / -2.25) * 3.14159) + 0.1)))))
            v19 -= float32((v12 - -2.25))
            _ = v19
        }
    }
    v1 = v1
    if ((!v1) && v1) {
        v15 -= v15
        v2 -= (^(v2))
    } else {
        switch "" {
        case "foo bar":
            if v1 {
                v1 = (false && true)
                var v20 float64 = (((v12 - (v12 * v12)) + v12) / 1e10)
                if false {
                    v2 = (v2 & (^(v2)))
                } else {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 float64 = (v12 - v12)
                        _ = fi21
                        _ = v22
                    }
                    fmt.Println(v20, v20)
                }
                _ = v20
            } else {
                fmt.Println(v1, v2, v2)
                fmt.Println(v2, v12)
            }
        case "":
            switch -1 {
            case -1:
                v15 += ((((3.14159 / v15) / (v15 * 0.0)) + v15) - 0.1)
            default:
                var v23 int32 = -427
                _ = v23
            }
        case "\n\t":
            if (32768 < 771) {
                v24 := GenS2{F0: 2, F1: 42, F2: 9223372036854775807}
                v3["hello"] = 4294967295
                _ = v24
            } else {
                var v25 uint16 = 42
                fmt.Println(v25, v1)
                if ((1 < v2) || (100 < -103)) {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 float32 = 0.1
                        switch 2147483647 {
                        case -678:
                            for fi28 := 0; fi28 < 3; fi28++ {
                                v27 += v15
                                _ = fi28
                            }
                        case -325:
                            for fi29 := 0; fi29 < 12; fi29++ {
                                if false {
                                    switch (fi26 << ((65535) & 63)) {
                                    case -452:
                                        v30 := GenS1{F0: 804, F1: (v2 * (v2 * 42)), F2: "a"}
                                        _ = v30
                                    default:
                                        if true {
                                            var v31 uint = 771
                                            _ = v31
                                        }
                                    }
                                    v32 := GenS3{F0: -32768}
                                    var v33 float64 = v12
                                    v12 = ((v33 - 0.0) / (v12 + v12))
                                    var v34 int16 = genFn0(uint(v2))
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                _ = fi29
                            }
                        case 840:
                            v2 -= 2147483648
                        }
                        _ = fi26
                        _ = v27
                    }
                    v15 = v15
                }
                _ = v25
            }
        }
        var v35 int16 = 676
        var v36 uint16 = 42
        _ = v35
        _ = v36
    }
    if false {
        var v37 int16 = genFn0(4294967295)
        fmt.Println(v2, v12)
        v2 -= 2147483647
        _ = v37
    }
    v1 = v1
    v38 := GenS4{F0: 9223372036854775807, F1: 661, F2: true}
    v15 = v15
    v39 := GenS2{F0: 2, F1: 255, F2: 128}
    v2 = 1
    v40 := GenS3{F0: 643}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}
