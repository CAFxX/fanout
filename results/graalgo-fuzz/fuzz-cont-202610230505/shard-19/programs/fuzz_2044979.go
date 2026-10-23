package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 int16 }
func genFn0(p0 string, p1 uint8, p2 int16) int {
    var v7 string = "a"
    v8 := []int{372, 2147483647, -318, 23}
    if false {
        var v9 uint = 1
        _ = v9
    }
    switch 299 {
    case 1:
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int64 = 974
            fmt.Println(p1)
            if true {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = fi10
            _ = v11
        }
    case 4294967295:
        v12 := GenS1{F0: -81, F1: 18446744073709551615}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 323
}

func main() {
    var v1 int64 = -899
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    v4 := []int{1, 1, -2147483648, 1}
    var v5 uint32 = 0
    var v6 string = ""
    v13 := GenS2{F0: 127}
    if (1 > 1) {
        var v14 int64 = 0
        _ = v14
    }
    var v15 int8 = -6
    v16 := v2[(v6 + "foo bar")]
    v6 = "\n\t"
    for fi17 := 0; fi17 < 20; fi17++ {
        if v16 {
            for fi18 := 0; fi18 < 2; fi18++ {
                v5 = 697
                v16 = (true || ((v15 <= v15) || true))
                _ = len(v3)
                _ = fi18
            }
            var v19 float32 = 3.14159
            fmt.Println(v1)
            _ = v19
        }
        _ = fi17
    }
    var v20 uint16 = 771
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            switch fi22 {
            case 155:
                v23 := GenS2{F0: int8(fi21)}
                _ = v23
            case 32767:
                if (!v16) {
                    fmt.Println(v15)
                } else {
                    v5 += v5
                    v24 := GenS2{F0: (v15 << ((65535) & 63))}
                    v1 += 334
                    _ = v24
                }
            default:
                v25 := GenS1{F0: 0, F1: 4294967295}
                _ = v25
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi22
        }
        var v26 int16 = 969
        v5 = 4294967295
        _ = fi21
        _ = v26
    }
    var v27 float64 = ((-2.25 * float64(3.14159)) - 1.5)
    v5 = v5
    var v28 int32 = 1
    if ((v1 + (v1 - 42)) == (v1 ^ 42)) {
        v4 = append(v4, 42)
        v20 += 719
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        fmt.Println(v16, v1, v15)
        _ = fi29
    }
    if true {
        v30 := GenS2{F0: v15}
        _ = v30
    }
    switch genFn0(v6, 127, -134) {
    case 42:
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    case -2147483648:
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 float64 = ((v27 - 1.5) / float64(-2.25))
            fmt.Println(v15)
            v16 = v16
            _ = fi31
            _ = v32
        }
    default:
        v33 := GenS3{F0: -785, F1: (959 < 65535), F2: 1}
        _ = v33
    }
    v5 = (v5 + v5)
    switch genFn0((v6 + v6), 188, 1000) {
    case 2147483647:
        var v34 int64 = (v1 - -9223372036854775808)
        _ = v34
    case 694:
        _ = len(v4)
    case 32767:
        if v16 {
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS3{F0: -32768, F1: (v27 < (v27 * (float64(float32(v27)) - (v27 - v27)))), F2: -1}
                _ = fi35
                _ = v36
            }
            switch ((v6 + "\n\t") + (((v6 + v6) + v6) + ("" + "foo bar"))) {
            case "\n\t":
                var v37 uint16 = (v20 >> ((1) & 63))
                _ = v37
            case "foo bar":
                if (true && v16) {
                    v15 -= int8(v5)
                    var v38 float32 = 1e10
                    fmt.Println(v15)
                    _ = v38
                } else {
                    var v39 bool = (127 < 336)
                    _ = v39
                }
            default:
                var v40 int64 = -621
                _ = v40
            }
            v15 = (v15 % ((v15) | 1))
        }
    default:
        v16 = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
