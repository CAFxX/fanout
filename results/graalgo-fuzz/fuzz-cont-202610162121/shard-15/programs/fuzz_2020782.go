package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int32; F1 int8; F2 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint; F1 int8; F2 int16 }
type GenS6 struct { F0 int16; F1 string; F2 uint64 }
func genFn0(p0 uint8) uint8 {
    p0 -= p0
    fmt.Println(p0)
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: -9223372036854775808}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: -785}
        fmt.Println(fi6)
        _ = fi6
        _ = v9
    }
    _ = p0
    return 193
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-2147483648, -9223372036854775808, 42}
    var v3 int64 = 300
    var v4 float32 = ((((1.5 / 1.5) / ((-2.25 / 1.5) * 1e10)) + (float32(0.0) * 1e10)) + float32((((((-2.25 + 0.1) + 0.1) * 1.5) / (-2.25 - (0.0 / (float64(1e10) + 0.0)))) - ((float64(0.0) + (1e10 / 1e10)) - float64(0.0)))))
    var v5 int = -2147483648
    switch "" {
    case "\n\t":
        fmt.Println(v3)
    case "a":
        var v10 int = (v5 / ((v5) | 1))
        _ = v10
    case "hello":
        v11 := GenS1{F0: v3}
        _ = v11
    default:
        if (!("foo bar" <= "a")) {
            var v12 uint32 = 4294967295
            var v13 uint64 = 2
            fmt.Println(v12, v13, v5)
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 128, F1: 776}
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 string = ("\n\t" + "a")
                v17 := GenS1{F0: (v3 ^ (v3 / ((-1) | 1)))}
                _ = fi15
                _ = v16
                _ = v17
            }
            switch v5 {
            case 840:
                v18 := GenS3{F0: -1, F1: 100, F2: true}
                _ = v18
            case -1:
                var v19 int16 = 1
                _ = v19
            }
            _ = v14
        }
    }
    var v20 float32 = 1e10
    _ = len(v2)
    if true {
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(v20, v3, fi21)
            v20 = (-2.25 * v4)
            _ = fi21
        }
        var v22 uint8 = 255
        _ = v22
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            v3 -= 647
            v24 := GenS1{F0: int64(v5)}
            v25 := GenS4{F0: 589, F1: 206}
            _ = fi23
            _ = v24
            _ = v25
        }
        v5 = v5
        if (127 <= 42) {
            var v26 int64 = (-(v3))
            switch (v5 << ((0) & 63)) {
            case 2147483647:
                for fi27 := 0; fi27 < 13; fi27++ {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 float64 = 1e10
                        _ = fi28
                        _ = v29
                    }
                    _ = fi27
                }
            default:
                var v30 int = -1
                _ = v30
            }
            _ = v26
        } else {
            if false {
                v3 += int64(v5)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                if false {
                    for fi31 := 0; fi31 < 10; fi31++ {
                        for fi32 := 0; fi32 < 19; fi32++ {
                            var v33 uint = 42
                            var v34 string = "hello"
                            switch (((("a" + "foo bar") + ("a" + v34)) + "foo bar") + v34) {
                            case "hello":
                                if len(v2) > 0 {
                                    v2[0] = v5
                                }
                            }
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = fi31
                    }
                    var v35 int64 = (-(v3))
                    _ = v35
                }
                if true {
                    for fi36 := 0; fi36 < 15; fi36++ {
                        v37 := v1["hello"]
                        _ = fi36
                        _ = v37
                    }
                    v38 := v1["\n\t"]
                    _ = v38
                } else {
                    var v39 int = -97
                    _ = v39
                }
            }
        }
    }
    var v40 int32 = -973
    v3 = (v3 & v3)
    v3 = v3
    switch "foo bar" {
    case "foo bar":
        v4 = ((1e10 - (-2.25 * 1.5)) * v4)
    default:
        _ = len(v1)
    }
    fmt.Println(v20, v40)
    v1["hello"] = v5
    v41 := GenS1{F0: (^(v3))}
    v42 := GenS5{F0: 42, F1: -39, F2: 1000}
    v43 := GenS4{F0: 1, F1: 19}
    if (-100 > v5) {
        v5 = 9223372036854775807
        if ((v3 <= 42) && (-2.25 != 1.5)) {
            v2 = append(v2, (v5 & v5))
        }
        v20 = (v4 - v4)
    }
    v44 := GenS6{F0: 42, F1: "\n\t", F2: 2}
    if (!(65535 <= 1)) {
        v3 = (v3 * -65)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
