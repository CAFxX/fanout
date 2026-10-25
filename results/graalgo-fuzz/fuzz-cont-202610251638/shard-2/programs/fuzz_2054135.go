package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 int8 = 127
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 uint = 608
    var v5 uint16 = 255
    var v6 bool = (4294967295 > uint32(v5))
    v6 = ((^(v5)) == (v5 >> (((v4 - (v4 & 984))) & 63)))
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 = v5
        for fi8 := 0; fi8 < 2; fi8++ {
            if false {
                if true {
                    v4 = 1
                    v1 = 127
                }
            }
            v9 := []int{148, 631, (fi7 | (-(fi7))), -1}
            v5 -= (v5 ^ (v5 & (v5 - (v5 / ((v5) | 1)))))
            switch fi7 {
            case -9223372036854775808:
                if (fi8 != 2147483647) {
                    if ((0.1 + float32(1.5)) != ((1e10 + 3.14159) - (-2.25 / ((((0.1 * 1.5) + -2.25) / 1.5) / (float32((-2.25 / 1.0)) * 1.5))))) {
                        v4 += v4
                        var v10 uint16 = v5
                        v2[""] = (!v6)
                        _ = v10
                    }
                }
            case -2147483648:
                v11 := GenS1{F0: 884, F1: v4, F2: 0}
                _ = v11
            default:
                var v12 int64 = 0
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v13 int = 32767
    var v14 string = "a"
    var v15 int64 = 0
    for fi16 := 0; fi16 < 17; fi16++ {
        v1 -= v1
        v17 := GenS2{F0: 4294967295, F1: ((v4 > v4) && (1 != uint64(v5))), F2: v14}
        _ = fi16
        _ = v17
    }
    v18 := []int{}
    if v6 {
        var v19 uint16 = v5
        if false {
            v13 += (v13 * v13)
            var v20 uint8 = 0
            _ = v20
        } else {
            v13 = v13
            for fi21 := 0; fi21 < 10; fi21++ {
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 uint32 = 551
                    v6 = v6
                    var v24 int16 = -628
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                fmt.Println(fi21, v15, fi21)
                if (!(v15 <= v15)) {
                    v18 = append(v18, v13)
                    for fi25 := 0; fi25 < 14; fi25++ {
                        var v26 uint8 = 2
                        v27 := GenS1{F0: v19, F1: v4, F2: 802}
                        var v28 string = v14
                        switch "" {
                        case "hello":
                            var v29 uint16 = (v5 / ((v5) | 1))
                            _ = v29
                        case "\n\t":
                            v6 = v6
                        case "foo bar":
                            v19 += (v5 / ((v19) | 1))
                        }
                        fmt.Println(v4)
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                } else {
                    var v30 int64 = (v15 % ((v15) | 1))
                    delete(v2, "foo bar")
                    v31 := GenS1{F0: (v5 << (((v4 | v4)) & 63)), F1: (v4 + v4), F2: 42}
                    _ = v30
                    _ = v31
                }
                _ = fi21
            }
        }
        _ = v19
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 int8 = 1
        v5 = (v5 & (v5 * v5))
        _ = fi32
        _ = v33
    }
    switch (v13 << ((v4) & 63)) {
    case 32768:
        for fi34 := 0; fi34 < 1; fi34++ {
            if v6 {
                v4 = (v4 ^ (v4 % ((42) | 1)))
                v4 = 2
                for fi35 := 0; fi35 < 2; fi35++ {
                    var v36 int64 = v15
                    v37 := v3[""]
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            fmt.Println(v13)
            v18 = append(v18, fi34)
            _ = fi34
        }
    }
    switch "foo bar" {
    case "a":
        for fi38 := 0; fi38 < 17; fi38++ {
            switch 42 {
            case 9223372036854775807:
                v13 = v13
            case 2147483647:
                v13 -= (fi38 >> (((v4 & v4)) & 63))
            default:
                var v39 uint64 = 42
                _ = v39
            }
            for fi40 := 0; fi40 < 12; fi40++ {
                fmt.Println(v13, v4, v14)
                var v41 int32 = 2147483647
                _ = fi40
                _ = v41
            }
            fmt.Println(v13)
            _ = fi38
        }
    }
    v42 := v3[v14]
    v5 = 255
    fmt.Println(v15, v4)
    var v43 float64 = (0.0 - (((3.14159 + float64((((((0.0 / 1.0) + (1.5 - -2.25)) + 3.14159) + ((1.5 * (0.0 - 1e10)) * 0.0)) - 1e10))) * 0.1) / 1.0))
    var v44 float32 = 1.5
    v4 -= 255
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
