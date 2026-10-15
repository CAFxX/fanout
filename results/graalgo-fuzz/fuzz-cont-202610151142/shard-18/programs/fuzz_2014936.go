package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 uint32; F1 string; F2 int }
func genFn0() int64 {
    v4 := GenS1{F0: 18446744073709551615, F1: 0}
    var v5 uint32 = 2147483647
    _ = v4
    _ = v5
    return 552
}
func genFn1(p0 bool) uint {
    var v6 int = 688
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS2{F0: 450}
            fmt.Println(fi8, fi8)
            _ = fi9
            _ = v10
        }
        switch "hello" {
        case "foo bar":
            v11 := GenS3{F0: (fi8 << ((42) & 63))}
            _ = v11
        case "\n\t":
            fmt.Println(fi8, p0, v6)
        }
        fmt.Println(v6, p0)
        _ = fi8
    }
    _ = p0
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 float64 = float64(float32((1e10 - -2.25)))
    var v2 int = 0
    v3 := []int{(v2 << ((255) & 63)), v2, (v2 + (v2 / ((v2) | 1))), -919}
    v2 = v2
    switch ((("\n\t" + "") + (("" + "foo bar") + "")) + "\n\t") {
    case "\n\t":
        v12 := GenS1{F0: 42, F1: v2}
        _ = v12
    case "hello":
        v2 -= (-(v2))
    case "":
        v13 := GenS1{F0: 186, F1: -9223372036854775808}
        _ = v13
    default:
        v14 := GenS1{F0: 1, F1: 32767}
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = 32768
    }
    if true {
        v15 := GenS1{F0: 42, F1: -329}
        v1 -= (float64(-2.25) - v1)
        for fi16 := 0; fi16 < 2; fi16++ {
            if (1 >= 1) {
                v3 = append(v3, fi16)
                v17 := GenS3{F0: (-(fi16))}
                if len(v3) > 0 {
                    v3[0] = v2
                }
                _ = v17
            } else {
                switch "" {
                case "":
                    v2 += fi16
                default:
                    fmt.Println(fi16, v1, v2)
                }
                for fi18 := 0; fi18 < 18; fi18++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (!(-654 > -2147483648)) {
                        v19 := GenS4{F0: 727, F1: fi16}
                        v2 = (fi16 % (((fi18 * (-(fi16)))) | 1))
                        v20 := GenS1{F0: 2, F1: -24}
                        _ = v19
                        _ = v20
                    }
                    if false {
                        v1 = 3.14159
                    }
                    _ = fi18
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v22 := GenS5{F0: 0, F1: 0}
                    v1 = float64((1e10 - (0.1 + float32((v1 / 0.1)))))
                    _ = fi21
                    _ = v22
                }
            }
            fmt.Println(v1, v2)
            switch (-(v2)) {
            case 9223372036854775807:
                for fi23 := 0; fi23 < 15; fi23++ {
                    if (65535 < 1) {
                        var v24 int = (fi16 | (fi23 + fi16))
                        v2 -= (v2 - -356)
                        var v25 bool = (-1 > 42)
                        var v26 uint16 = 65535
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v27 bool = (true && true)
                    switch 32768 {
                    case 127:
                        var v28 uint64 = 367
                        if (!(127 > int8(v28))) {
                            var v29 float64 = v1
                            _ = v29
                        }
                        _ = v28
                    }
                    v1 = ((v1 - (float64(1e10) / v1)) + v1)
                    _ = fi23
                    _ = v27
                }
            case -2147483648:
                v30 := GenS1{F0: 921, F1: fi16}
                _ = v30
            default:
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 int8 = -128
                    fmt.Println(fi31)
                    if len(v3) > 0 {
                        v3[0] = (fi31 | 1)
                    }
                    _ = fi31
                    _ = v32
                }
            }
            _ = fi16
        }
        _ = v15
    }
    v2 = v2
    for fi33 := 0; fi33 < 10; fi33++ {
        if len(v3) > 0 {
            v3[0] = fi33
        }
        if len(v3) > 0 {
            v3[0] = (v2 >> ((16) & 63))
        }
        fmt.Println(v2, v1, v1)
        _ = fi33
    }
    v34 := GenS5{F0: 32767, F1: 773}
    fmt.Println(v2)
    var v35 float64 = v1
    fmt.Println(v2, v35, v35)
    var v36 uint32 = 65535
    var v37 int16 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v38 := GenS6{F0: 2, F1: ("a" + "\n\t"), F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
