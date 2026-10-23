package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 int8) int8 {
    if false {
        v4 := GenS1{F0: -1, F1: p0, F2: 366}
        v5 := GenS2{F0: 130, F1: 32767}
        _ = v4
        _ = v5
    }
    var v6 uint8 = 0
    _ = p0
    _ = p1
    _ = v6
    return p1
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), "\n\t", "a"}
    var v2 int64 = 1
    var v3 int = 0
    v2 -= v2
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 = 808
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        if false {
            if (!false) {
                var v9 uint8 = 42
                _ = v9
            } else {
                _ = len(v1)
            }
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        } else {
            switch fi8 {
            case -751:
                var v10 int = v3
                _ = v10
            case -5:
                if true {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 float64 = 0.0
                        var v13 bool = (77 <= 4294967295)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(v2, v3, v2)
                    fmt.Println(fi8, v3, v3)
                } else {
                    var v14 uint64 = 0
                    v14 += v14
                    _ = v14
                }
            }
            var v15 uint = 1
            var v16 float32 = 1e10
            _ = v15
            _ = v16
        }
        v17 := GenS3{F0: v2}
        _ = fi8
        _ = v17
    }
    switch "" {
    case "hello":
        var v18 uint = 65535
        _ = v18
    default:
        switch (v3 >> ((18446744073709551615) & 63)) {
        case 905:
            v3 = (v3 * 9223372036854775807)
        default:
            if ((3.14159 >= (1e10 / 1.5)) && (!true)) {
                var v19 int32 = 42
                v3 = v3
                v2 -= (v2 << ((721) & 63))
                _ = v19
            }
        }
    }
    v20 := GenS3{F0: v2}
    if (!false) {
        v21 := GenS3{F0: -1}
        v3 = v3
        fmt.Println(v3, v3)
        _ = v21
    } else {
        v1 = append(v1, ((("\n\t" + "foo bar") + ("a" + ("\n\t" + "hello"))) + "\n\t"))
        v3 = 32767
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 -= (v2 + 0)
    v2 -= (v2 << ((2) & 63))
    for fi22 := 0; fi22 < 13; fi22++ {
        v3 += (fi22 * 768)
        if false {
            var v23 float64 = (float64((0.1 - 3.14159)) * ((-2.25 * 1e10) + 0.0))
            _ = v23
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi22
    }
    v2 += v2
    if (!true) {
        var v24 uint32 = 65535
        for fi25 := 0; fi25 < 12; fi25++ {
            switch fi25 {
            case 1:
                v3 -= -2147483648
            default:
                fmt.Println(v2, v3, v2)
            }
            _ = len(v1)
            if (((!(v2 >= -526)) && true) || true) {
                for fi26 := 0; fi26 < 9; fi26++ {
                    v27 := GenS1{F0: v2, F1: 2, F2: (^(v2))}
                    _ = fi26
                    _ = v27
                }
                v2 = v2
                var v28 float64 = 0.0
                _ = v28
            }
            _ = fi25
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v24
    } else {
        if (!false) {
            var v29 float32 = 3.14159
            fmt.Println(v2, v2)
            fmt.Println(v3, v2)
            _ = v29
        }
        fmt.Println(v2, v3)
        v2 += v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
