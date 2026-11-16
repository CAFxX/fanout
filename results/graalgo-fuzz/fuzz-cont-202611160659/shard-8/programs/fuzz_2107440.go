package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 bool, p1 int64) uint16 {
    if (1.5 < ((0.0 + -2.25) / (3.14159 - 1.5))) {
        var v4 uint32 = 4294967295
        if ("a" == ("foo bar" + "a")) {
            if true {
                v5 := GenS1{F0: 742, F1: -663}
                p1 -= p1
                _ = v5
            }
        } else {
            v6 := []int{128, 4294967295, 9223372036854775807, 9223372036854775807}
            for fi7 := 0; fi7 < 13; fi7++ {
                p0 = (p0 && (p0 && (622 < -2147483648)))
                var v8 int16 = 42
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        _ = v4
    } else {
        v9 := []int{203, 9223372036854775807, 32767, -9223372036854775808}
        if true {
            if p0 {
                fmt.Println(p0, p0)
                if p0 {
                    fmt.Println(p1)
                    if (42 >= 707) {
                        if (32767 <= 1) {
                            fmt.Println(p1, p0)
                            v10 := GenS1{F0: (p1 ^ 0), F1: -686}
                            _ = v10
                        }
                        if len(v9) > 0 {
                            v9[0] = 4294967295
                        }
                        p0 = (42 != 391)
                    }
                }
                p0 = false
            }
        }
        v11 := GenS2{F0: 65535}
        _ = v11
        _ = v9
    }
    p0 = (4294967295 <= 2147483648)
    v12 := GenS1{F0: p1, F1: 421}
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{(fi13 + (fi13 | (-(fi13))))}
        for fi15 := 0; fi15 < 7; fi15++ {
            switch fi13 {
            case 127:
                if len(v14) > 0 {
                    v14[0] = fi13
                }
            case -919:
                v14 = append(v14, (fi15 + fi15))
            default:
                if false {
                    v16 := GenS2{F0: 0}
                    fmt.Println(p0, p0)
                    var v17 int8 = -128
                    _ = v16
                    _ = v17
                } else {
                    var v18 int32 = -601
                    var v19 bool = (!((("hello" + ("\n\t" + "hello")) + "hello") > "hello"))
                    _ = v18
                    _ = v19
                }
            }
            _ = fi15
        }
        p1 = p1
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return 255
}
func genFn1(p0 uint8, p1 uint32) uint16 {
    var v20 int8 = -67
    v21 := GenS2{F0: p1}
    v22 := []int{}
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v22
    return 255
}

func main() {
    var v1 uint8 = 120
    var v2 int16 = 42
    var v3 int16 = v2
    if (0 == 896) {
        v23 := []int{350, 128, 393}
        fmt.Println(v1)
        _ = v23
    } else {
        var v24 int = 0
        _ = v24
    }
    var v25 int64 = -221
    if (true || true) {
        for fi26 := 0; fi26 < 13; fi26++ {
            v25 = (-(v25))
            v3 = (v3 ^ v3)
            _ = fi26
        }
        var v27 bool = (4294967295 == 914)
        _ = v27
    } else {
        v2 = v2
        v3 = (v2 | v2)
        var v28 int8 = 1
        _ = v28
    }
    v29 := GenS3{F0: (v3 << ((42) & 63)), F1: 1}
    fmt.Println(v2)
    if ((v3 + v2) > -1) {
        var v30 float64 = 0.1
        v31 := GenS3{F0: v3, F1: v3}
        v2 = (v2 | (-(v3)))
        _ = v30
        _ = v31
    }
    var v32 uint = 4294967295
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 int64 = v25
        switch "hello" {
        case "foo bar":
            var v35 int32 = 2147483647
            _ = v35
        case "hello":
            fmt.Println(v32)
        case "a":
            switch (fi33 % (((fi33 / (((fi33 + (fi33 % (((fi33 & 127)) | 1)))) | 1))) | 1)) {
            case -2147483648:
                v36 := []int{(^(fi33)), fi33, -875, 127}
                _ = v36
            case 9223372036854775807:
                v3 = (^(v2))
            case 127:
                var v37 int32 = -967
                _ = v37
            }
        }
        switch fi33 {
        case 42:
            v3 += v2
        }
        _ = fi33
        _ = v34
    }
    v1 += v1
    if (42 <= 127) {
        v1 += 98
        var v38 int64 = -188
        v25 += v38
        _ = v38
    }
    var v39 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v32
    _ = v39
}
