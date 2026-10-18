package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16, p1 int8, p2 int32) int16 {
    fmt.Println(p2, p0, p0)
    p0 += -920
    v7 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (^(p0))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -1
    var v3 uint16 = 1
    var v4 int16 = 32767
    v5 := []int{v2}
    var v6 float64 = 0.1
    if false {
        v8 := GenS1{F0: "a", F1: 127, F2: 127}
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: ("foo bar" + ""), F1: 0, F2: (v2 + fi9)}
            v5 = append(v5, fi9)
            for fi11 := 0; fi11 < 3; fi11++ {
                v5 = append(v5, (fi9 >> ((65535) & 63)))
                fmt.Println(fi11, fi11)
                var v12 int32 = 497
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        if false {
            v4 = v4
            v6 -= 0.0
            var v13 int64 = 390
            _ = v13
        }
        fmt.Println(v4, v6)
    }
    var v14 bool = (true || true)
    for fi15 := 0; fi15 < 7; fi15++ {
        v14 = true
        _ = fi15
    }
    v3 = (^(v3))
    v16 := GenS1{F0: "\n\t", F1: 1, F2: v2}
    var v17 int = v2
    if (255 <= 42) {
        var v18 uint = 0
        _ = v18
    }
    if v14 {
        if (777 >= 977) {
            var v19 uint16 = (v3 | (v3 / (((v3 | (v3 % ((65535) | 1)))) | 1)))
            _ = v19
        }
        var v20 int8 = 83
        _ = v20
    } else {
        var v21 uint32 = 0
        if (255 >= 207) {
            v22 := GenS1{F0: "\n\t", F1: -1, F2: v17}
            _ = v22
        }
        _ = v21
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v14 = ((0 == 9223372036854775807) && (1 <= 18446744073709551615))
        v14 = (v14 || (4294967295 > 2))
        _ = fi23
    }
    fmt.Println(v6, v6)
    var v24 uint16 = (v3 % (((v3 >> ((250) & 63))) | 1))
    v6 = (v6 / (v6 - (v6 * (((v6 / v6) * (-2.25 + 0.1)) + ((-2.25 + v6) * v6)))))
    var v25 float64 = -2.25
    for fi26 := 0; fi26 < 8; fi26++ {
        switch v2 {
        case -9223372036854775808:
            v27 := GenS2{F0: 497}
            _ = v27
        case 9223372036854775807:
            switch "foo bar" {
            case "foo bar":
                v28 := GenS2{F0: 2147483648}
                _ = v28
            case "hello":
                v29 := v1[2147483647]
                _ = v29
            }
        }
        if (0 >= 255) {
            var v30 uint64 = 0
            _ = v30
        }
        var v31 int64 = 9223372036854775807
        _ = fi26
        _ = v31
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 uint16 = v3
        if (1 > -2147483648) {
            v24 -= 65535
            v14 = (0 >= -1)
            var v34 int = -2147483648
            _ = v34
        } else {
            v33 += (v3 * (^(v24)))
        }
        _ = fi32
        _ = v33
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        if (160 <= 17) {
            var v36 int64 = -9223372036854775808
            _ = v36
        } else {
            switch ("" + "hello") {
            case "foo bar":
                for fi37 := 0; fi37 < 16; fi37++ {
                    if ((v2 ^ fi37) < (-(fi35))) {
                        fmt.Println(v25, v4)
                        if v14 {
                            switch ((("" + "a") + "foo bar") + "hello") {
                            case "":
                                v6 += ((v25 - (((v6 * 1.5) / v25) / v25)) + (v6 / float64(((1.5 + 1e10) * 3.14159))))
                            }
                        } else {
                            var v38 bool = (-23 <= 0)
                            if v38 {
                                v17 = fi37
                            }
                            v39 := GenS2{F0: 2147483647}
                            _ = v38
                            _ = v39
                        }
                        v1[-9223372036854775808] = 790
                    } else {
                        var v40 uint64 = 42
                        delete(v1, (^(v17)))
                        var v41 int = 127
                        _ = v40
                        _ = v41
                    }
                    _ = fi37
                }
            case "a":
                for fi42 := 0; fi42 < 7; fi42++ {
                    if (!v14) {
                        if (v3 == v3) {
                            v14 = true
                        }
                        v43 := GenS2{F0: 1}
                        switch ("hello" + ("foo bar" + "")) {
                        case "a":
                            v44 := GenS1{F0: (("a" + ("\n\t" + (("" + "\n\t") + "foo bar"))) + "foo bar"), F1: 7, F2: fi35}
                            _ = v44
                        default:
                            v45 := GenS2{F0: 47}
                            _ = v45
                        }
                        _ = v43
                    }
                    var v46 uint8 = 255
                    v47 := GenS3{F0: (v4 & 0)}
                    _ = fi42
                    _ = v46
                    _ = v47
                }
            default:
                var v48 int8 = -128
                _ = v48
            }
        }
        v6 -= (v25 * 3.14159)
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
