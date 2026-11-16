package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    if (-707 >= -9223372036854775808) {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := []int{(fi6 + -1), (fi6 | 776)}
            p0 = (^(p0))
            _ = fi6
            _ = v7
        }
        var v8 int32 = -1
        _ = v8
    }
    v9 := []int{32767, -967, 127}
    _ = p0
    _ = v9
    return 0
}
func genFn1(p0 uint8, p1 uint, p2 int64) uint {
    var v10 float32 = -2.25
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 string = "foo bar"
        switch v12 {
        case "hello":
            var v13 int32 = 42
            _ = v13
        case "foo bar":
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 uint64 = 2
                _ = fi14
                _ = v15
            }
        }
        if true {
            v16 := []int{(-(fi11))}
            p0 = p0
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = v16
        } else {
            switch fi11 {
            case -1:
                var v17 float32 = (((float32(-2.25) + -2.25) / -2.25) - 0.1)
                _ = v17
            case 649:
                var v18 int16 = 1000
                _ = v18
            default:
                v19 := []int{fi11, 1, 42}
                _ = v19
            }
        }
        _ = fi11
        _ = v12
    }
    v20 := []int{42, 9223372036854775807}
    fmt.Println(p2, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v20
    return p1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]bool)
    var v3 int32 = 308
    var v4 int32 = 42
    var v5 float64 = 1e10
    var v21 uint16 = 42
    var v22 float64 = (v5 * v5)
    for fi23 := 0; fi23 < 14; fi23++ {
        v3 = v4
        v24 := []int{fi23, (fi23 >> ((65535) & 63))}
        v22 += 1.5
        _ = fi23
        _ = v24
    }
    var v25 uint8 = genFn0(449)
    fmt.Println(v1, v3)
    switch "\n\t" {
    case "a":
        var v26 uint = 863
        _ = v26
    case "\n\t":
        if ((float32(v5) / -2.25) >= float32(3.14159)) {
            switch ("" + ("foo bar" + ((("hello" + "foo bar") + "\n\t") + ("a" + ("\n\t" + ("a" + "\n\t")))))) {
            case "foo bar":
                fmt.Println(v25)
            default:
                fmt.Println(v1, v22, v5)
            }
            v27 := v2[307]
            _ = v27
        }
    case "foo bar":
        fmt.Println(v3)
    default:
        var v28 string = "a"
        _ = v28
    }
    var v29 float32 = 0.1
    fmt.Println(v21, v1, v22)
    fmt.Println(v21)
    v21 += v21
    v30 := []int{}
    fmt.Println(v22, v22, v21)
    fmt.Println(v29, v3)
    fmt.Println(v5, v21)
    v5 -= (((v22 * v22) * v5) * 1.5)
    v31 := v2[int(v25)]
    for fi32 := 0; fi32 < 6; fi32++ {
        v29 += (((v29 / (v29 * 1e10)) * 0.0) - float32(1e10))
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 uint32 = 65535
            if v31 {
                if v31 {
                    v35 := GenS1{F0: 331, F1: 409, F2: v1}
                    _ = v35
                }
                var v36 int64 = -1
                var v37 uint32 = v34
                _ = v36
                _ = v37
            } else {
                v1 += 53
                v38 := GenS1{F0: 42, F1: 1, F2: (v25 ^ v1)}
                v3 += 233
                _ = v38
            }
            var v39 uint64 = 42
            _ = fi33
            _ = v34
            _ = v39
        }
        v21 += v21
        _ = fi32
    }
    if false {
        if true {
            _ = len(v2)
            var v40 bool = true
            _ = v40
        }
        switch "" {
        case "a":
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 uint = 203
                _ = fi41
                _ = v42
            }
        case "foo bar":
            fmt.Println(v5, v31)
        }
        v5 = 0.1
    } else {
        var v43 float32 = (v29 + (((v29 + (v29 - (v29 + (v29 + float32(v5))))) / ((1.5 / v29) - (v29 / 0.1))) / 3.14159))
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
