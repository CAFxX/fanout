package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int) uint32 {
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{fi5, fi5, (fi5 ^ (p0 / ((fi5) | 1)))}
        if false {
            if len(v6) > 0 {
                v6[0] = (p0 - 32767)
            }
            var v7 bool = false
            fmt.Println(fi5, fi5)
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    if false {
        p0 = p0
    }
    v8 := []int{}
    _ = p0
    _ = v8
    return 157
}
func genFn1(p0 bool, p1 bool, p2 int32) uint32 {
    var v9 uint = 818
    v10 := []int{}
    switch "" {
    case "\n\t":
        p0 = p1
    case "foo bar":
        v10 = append(v10, 4294967295)
    }
    v9 -= v9
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]string)
    var v3 int = 42
    var v4 uint = 4294967295
    fmt.Println(v1)
    switch v3 {
    case -706:
        if false {
            if ("\n\t" >= "\n\t") {
                fmt.Println(v3, v1, v3)
                v11 := []int{(v3 | v3), v3, (^(v3))}
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 int64 = 42
                    v3 += (v3 & -9223372036854775808)
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                v3 = v3
                v1 += (v1 << ((2) & 63))
                _ = fi14
            }
        }
    default:
        for fi15 := 0; fi15 < 16; fi15++ {
            v4 = 296
            v3 = 558
            switch (v3 - (fi15 << ((4294967295) & 63))) {
            case -9223372036854775808:
                switch v3 {
                case 32767:
                    delete(v2, "a")
                case -105:
                    var v16 uint64 = 42
                    _ = v16
                }
            case 0:
                v17 := []int{}
                _ = v17
            }
            _ = fi15
        }
    }
    v18 := GenS1{F0: (v4 & v4), F1: (v1 / (((v1 | (v1 % (((v1 << (((^(v4))) & 63))) | 1)))) | 1))}
    if ((^(v4)) <= 65535) {
        if false {
            switch 42 {
            case 2147483647:
                switch ("" + ("foo bar" + "")) {
                case "":
                    v19 := GenS2{F0: 175, F1: (v4 | (v4 >> ((1) & 63))), F2: 42}
                    _ = v19
                default:
                    v20 := []int{v3}
                    _ = v20
                }
            case -2147483648:
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 int32 = -2147483648
                    _ = fi21
                    _ = v22
                }
            case 0:
                var v23 int = (v3 ^ (v3 * v3))
                _ = v23
            }
            v1 -= v1
            v1 = v1
        }
    } else {
        v24 := []int{v3, (v3 - 2147483647)}
        if false {
            if (false || true) {
                v24 = append(v24, v3)
                v25 := GenS2{F0: -2147483648, F1: 65535, F2: 0}
                _ = v25
            } else {
                _ = len(v24)
                v3 += (v3 & v3)
                if (true || (!false)) {
                    v1 = v1
                }
            }
            var v26 string = ((("\n\t" + "foo bar") + ((("a" + "foo bar") + ("hello" + "a")) + "foo bar")) + "a")
            _ = v26
        } else {
            var v27 uint64 = 2
            for fi28 := 0; fi28 < 7; fi28++ {
                if (32767 != 255) {
                    v1 -= 1
                    v4 -= (v4 & v4)
                    v1 = v1
                }
                fmt.Println(v1, fi28, v4)
                for fi29 := 0; fi29 < 5; fi29++ {
                    v30 := GenS3{F0: fi29, F1: v1}
                    var v31 string = ""
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
            _ = v27
        }
        fmt.Println(v4, v4, v4)
        _ = v24
    }
    v1 = 0
    fmt.Println(v1, v4)
    v32 := v2["foo bar"]
    if ((0.0 / (0.0 - (float64(0.0) - (float64(0.1) / 1.0)))) >= 0.1) {
        v1 -= v1
        var v33 uint = v4
        _ = v33
    } else {
        v1 += (v1 << (((^(v4))) & 63))
        if false {
            v34 := v2[(v32 + "a")]
            for fi35 := 0; fi35 < 2; fi35++ {
                if (103 != -128) {
                    if (true && false) {
                        var v36 uint8 = v1
                        fmt.Println(fi35, v3, v4)
                        v37 := GenS1{F0: 65535, F1: (v36 ^ (v1 >> (((^(v4))) & 63)))}
                        _ = v36
                        _ = v37
                    } else {
                        var v38 uint64 = 18446744073709551615
                        _ = v38
                    }
                }
                v39 := []int{}
                _ = fi35
                _ = v39
            }
            delete(v2, v32)
            _ = v34
        }
    }
    var v40 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
}
