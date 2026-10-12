package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0(p0 int16, p1 int8) uint32 {
    var v7 int8 = p1
    fmt.Println(p0, v7, p0)
    _ = p0
    _ = p1
    _ = v7
    return 389
}

func main() {
    v1 := []int{}
    var v2 int64 = -1
    v3 := []int{-9223372036854775808, 42, 1}
    v4 := []int{-649}
    v5 := make(map[string]int)
    v6 := []string{"hello", "foo bar", ("" + "foo bar")}
    v5["foo bar"] = 32768
    if true {
        _ = len(v5)
        if len(v6) > 0 {
            v6[0] = ""
        }
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v2 += v2
        switch "a" {
        case "\n\t":
            v9 := GenS1{F0: 821}
            _ = v9
        }
        var v10 float32 = float32(float64((float32((0.1 / -2.25)) * ((float32(-2.25) - 3.14159) / 3.14159))))
        _ = fi8
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 float64 = (float64((float32(0.0) - 1e10)) / -2.25)
        for fi13 := 0; fi13 < 14; fi13++ {
            v12 = (((0.0 / v12) - (0.1 * v12)) / float64(1.5))
            if false {
                switch "hello" {
                case "\n\t":
                    v14 := GenS2{F0: 822}
                    _ = v14
                case "hello":
                    var v15 uint64 = 0
                    if (!(194 < 71)) {
                        var v16 uint32 = 0
                        var v17 int16 = 42
                        _ = v16
                        _ = v17
                    } else {
                        v12 = v12
                    }
                    v2 += (v2 - v2)
                    v18 := GenS1{F0: 1}
                    _ = v15
                    _ = v18
                }
            } else {
                v19 := v5["foo bar"]
                var v20 uint8 = 0
                fmt.Println(v2, v2)
                _ = v19
                _ = v20
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    if (!true) {
        v6 = append(v6, "")
        if true {
            if len(v4) > 0 {
                v4[0] = 224
            }
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        delete(v5, "foo bar")
        var v22 float32 = float32(0.0)
        _ = fi21
        _ = v22
    }
    var v23 float64 = float64(3.14159)
    if false {
        switch (("a" + "\n\t") + "\n\t") {
        case "foo bar":
            var v24 int32 = 1
            _ = v24
        case "hello":
            v2 = 9223372036854775807
        default:
            v23 = (1.5 - float64((-2.25 - 3.14159)))
        }
    }
    var v25 string = "foo bar"
    v26 := GenS1{F0: -1}
    var v27 int = -2147483648
    if (1 > 1000) {
        var v28 uint64 = 869
        for fi29 := 0; fi29 < 4; fi29++ {
            if (1e10 > ((-2.25 * float32(-2.25)) - (0.0 + ((float32(v23) + ((1.5 / 3.14159) + (1e10 - -2.25))) - (1.5 / float32(float64(0.1))))))) {
                var v30 float32 = (3.14159 / -2.25)
                _ = v30
            } else {
                switch (v27 >> ((2) & 63)) {
                case -528:
                    for fi31 := 0; fi31 < 1; fi31++ {
                        v2 += 140
                        _ = fi31
                    }
                case -9223372036854775808:
                    for fi32 := 0; fi32 < 4; fi32++ {
                        var v33 int8 = 6
                        var v34 int = (fi29 / ((v27) | 1))
                        v2 += (v2 - (v2 + (^(v2))))
                        switch fi32 {
                        case 0:
                            switch -2147483648 {
                            case 4294967295:
                                v35 := GenS1{F0: 42}
                                _ = v35
                            case 127:
                                v2 = (v2 | (v2 % ((9223372036854775807) | 1)))
                            }
                        default:
                            fmt.Println(v23, v25, v34)
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case 9223372036854775807:
                    for fi36 := 0; fi36 < 1; fi36++ {
                        v1 = append(v1, fi36)
                        _ = fi36
                    }
                default:
                    v28 = (v28 + (v28 - v28))
                }
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                v38 := GenS3{F0: 9223372036854775807, F1: (v28 & 0), F2: 103}
                fmt.Println(v23, v23)
                _ = fi37
                _ = v38
            }
            _ = fi29
        }
        _ = len(v6)
        _ = v28
    }
    if (true || (1000 <= 1)) {
        var v39 float64 = 0.1
        v25 = ""
        _ = v39
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v2, v23)
    var v40 int64 = 1
    v41 := GenS1{F0: -1}
    v40 += (v2 * 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
