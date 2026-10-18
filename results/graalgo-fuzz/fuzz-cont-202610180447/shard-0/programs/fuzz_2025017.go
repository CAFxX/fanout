package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0() int8 {
    var v6 string = ("" + "a")
    var v7 uint64 = 732
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := []int{(^(fi8))}
        var v10 int16 = 1
        var v11 string = v6
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int32 = 380
    v13 := GenS1{F0: 32767, F1: -920, F2: 240}
    _ = v12
    _ = v13
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []string{"foo bar", ("a" + "\n\t"), "foo bar", "foo bar"}
    var v2 float64 = float64(0.1)
    v3 := make(map[uint32]int)
    var v4 uint64 = 42
    var v5 int = -2147483648
    var v14 float32 = -2.25
    v3[2147483648] = (v5 << ((4294967295) & 63))
    if true {
        for fi15 := 0; fi15 < 6; fi15++ {
            if (true || false) {
                v2 += 0.1
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 uint8 = 255
                    switch v5 {
                    case 4294967295:
                        v18 := GenS1{F0: (-(fi15)), F1: (fi15 >> ((uint(v17)) & 63)), F2: (^(v17))}
                        _ = v18
                    }
                    switch "hello" {
                    case "\n\t":
                        _ = len(v3)
                    default:
                        var v19 int32 = -735
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi15
        }
        v20 := GenS1{F0: v5, F1: (v5 + v5), F2: 53}
        _ = v20
    } else {
        var v21 float64 = 3.14159
        if true {
            if false {
                switch (v5 + 1) {
                case 1:
                    var v22 int32 = 42
                    _ = v22
                case 127:
                    var v23 float64 = v21
                    _ = v23
                }
                var v24 uint = 255
                v24 = v24
                _ = v24
            } else {
                var v25 uint = 851
                if (959 < 890) {
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    var v26 int16 = 1000
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v3[2] = 32768
                        if true {
                            var v28 uint64 = v4
                            switch 0 {
                            case 128:
                                v21 = v2
                            case 127:
                                v28 += (v28 * (v4 + 614))
                            default:
                                v4 -= 646
                            }
                            _ = v28
                        } else {
                            switch (fi27 << ((v25) & 63)) {
                            case 128:
                                v2 -= (v21 + v2)
                            }
                            v25 += (^(v25))
                        }
                        _ = fi27
                    }
                    _ = v26
                }
                v4 = 437
                _ = v25
            }
            v14 += (v14 - (v14 * v14))
            v4 += (v4 + (v4 & (^(v4))))
        }
        if ((-923 != 1) || false) {
            var v29 int8 = -1
            switch "a" {
            case "a":
                if (!(2 <= 1)) {
                    _ = len(v1)
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v21 = 0.0
                        v5 -= fi30
                        fmt.Println(fi30, fi30, fi30)
                        _ = fi30
                    }
                    v14 = float32(v21)
                }
            case "hello":
                v14 += ((float32(v2) - v14) + v14)
            case "foo bar":
                var v31 uint = 4294967295
                _ = v31
            default:
                v14 = float32(v21)
            }
            var v32 uint8 = 128
            _ = v29
            _ = v32
        }
        _ = v21
    }
    v33 := GenS1{F0: (v5 ^ v5), F1: v5, F2: 2}
    var v34 float32 = v14
    switch "\n\t" {
    case "hello":
        v34 -= (v14 + float32(v2))
    case "a":
        var v35 int = -712
        _ = v35
    case "":
        v36 := GenS2{F0: "\n\t", F1: 2}
        _ = v36
    }
    switch "" {
    case "\n\t":
        v14 += v14
    case "foo bar":
        var v37 uint16 = 2
        _ = v37
    default:
        v38 := v3[2147483647]
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 int16 = 42
        switch "\n\t" {
        case "a":
            fmt.Println(fi39, v14, v4)
        case "hello":
            var v41 int32 = 257
            _ = v41
        }
        _ = fi39
        _ = v40
    }
    var v42 float64 = 3.14159
    switch "foo bar" {
    case "":
        var v43 bool = (65535 < 42)
        _ = v43
    default:
        v44 := GenS2{F0: "\n\t", F1: 2}
        _ = v44
    }
    var v45 uint32 = 2
    v5 += (v5 * v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v45
    _ = v5
}
