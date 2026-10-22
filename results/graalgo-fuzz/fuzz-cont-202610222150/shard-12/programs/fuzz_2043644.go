package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint64, p1 int32) uint {
    fmt.Println(p0)
    if (0.1 != 3.14159) {
        var v5 float32 = (1e10 / 1.5)
        if false {
            switch 32767 {
            case 9223372036854775807:
                v5 += float32(float64(1.5))
            case 128:
                var v6 bool = false
                _ = v6
            default:
                var v7 int64 = 1
                _ = v7
            }
            var v8 int64 = -612
            v9 := GenS1{F0: 42, F1: -1, F2: v8}
            _ = v8
            _ = v9
        } else {
            var v10 uint = 42
            _ = v10
        }
        p0 = (^(p0))
        _ = v5
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            p1 = 1
            var v12 uint64 = (p0 << ((65535) & 63))
            _ = fi11
            _ = v12
        }
        p1 = 511
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 float64 = 1.5
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        if true {
            v16 := []int{(fi15 * 4294967295)}
            if (uint8(p0) > 255) {
                v17 := GenS1{F0: 173, F1: -1, F2: -765}
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                var v18 int64 = 0
                _ = v17
                _ = v18
            }
            _ = v16
        }
        v19 := []int{(fi15 & fi15), (fi15 / (((fi15 >> ((65535) & 63))) | 1)), (fi15 % ((4294967295) | 1))}
        switch ("foo bar" + ("a" + ((("" + "\n\t") + ((("a" + ("a" + "a")) + "foo bar") + ("hello" + "\n\t"))) + "a"))) {
        case "hello":
            p1 += (p1 + p1)
        case "\n\t":
            p0 = p0
        }
        _ = fi15
        _ = v19
    }
    var v20 int16 = 32767
    _ = p0
    _ = p1
    _ = v20
    return 2
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int = 4294967295
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := v1["hello"]
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            for fi23 := 0; fi23 < 8; fi23++ {
                if false {
                    v24 := []int{-150, v4}
                    v25 := GenS1{F0: 0, F1: 349, F2: 0}
                    _ = v24
                    _ = v25
                } else {
                    _ = len(v2)
                    var v26 int64 = 455
                    if ((2 != 4294967295) || false) {
                        v27 := GenS2{F0: 4294967295, F1: 50}
                        _ = v27
                    }
                    v28 := []int{-2147483648, (fi21 | (v22 - v22))}
                    _ = v26
                    _ = v28
                }
                v22 -= 0
                _ = fi23
            }
        case "a":
            var v29 uint32 = 2147483648
            _ = v29
        default:
            var v30 uint32 = 0
            _ = v30
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int32 = 910
            var v33 int = v22
            if ((-2.25 - ((0.0 + 0.0) - 0.0)) < (-2.25 - (-2.25 - float32((1.5 / (0.1 / 1e10)))))) {
                v33 -= (-(fi21))
                v34 := GenS1{F0: 120, F1: 790, F2: 42}
                _ = v34
            }
            fmt.Println(v33, fi31, v22)
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi21
        _ = v22
    }
    v35 := GenS1{F0: 255, F1: 1000, F2: 444}
    v36 := GenS1{F0: 128, F1: 32767, F2: 0}
    var v37 int = v4
    for fi38 := 0; fi38 < 9; fi38++ {
        var v39 float32 = ((float32(0.1) * (-2.25 - 0.0)) * 1.5)
        delete(v3, (("foo bar" + "a") + "hello"))
        v40 := GenS2{F0: 65535, F1: 42}
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 int = v4
    switch "a" {
    case "\n\t":
        delete(v2, ("foo bar" + "hello"))
    case "":
        var v42 string = (("foo bar" + ("hello" + "\n\t")) + ("a" + "foo bar"))
        _ = v42
    case "a":
        v37 = -706
    default:
        var v43 int = v4
        _ = v43
    }
    fmt.Println(v41)
    v41 -= 9223372036854775807
    switch (v37 >> ((112) & 63)) {
    case 4294967295:
        var v44 int8 = -109
        _ = v44
    case -9223372036854775808:
        fmt.Println(v4, v37, v41)
    default:
        v37 = (v41 | 32768)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
}
