package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int16 }
func genFn0(p0 int16, p1 uint16, p2 uint32) int16 {
    v6 := GenS1{F0: 0, F1: (p2 - 1), F2: p0}
    if (p0 < p0) {
        p1 = 186
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int16 = p0
            p1 = uint16(p2)
            switch "\n\t" {
            case "\n\t":
                v9 := []int{fi7}
                _ = v9
            case "":
                v10 := []int{(fi7 | (-(fi7))), -216, fi7}
                _ = v10
            case "hello":
                v11 := []int{-9223372036854775808}
                _ = v11
            }
            _ = fi7
            _ = v8
        }
    }
    var v12 float64 = float64((-2.25 / 1.5))
    v13 := GenS1{F0: 42, F1: (p2 + 2147483648), F2: 0}
    var v14 uint = 181
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    _ = v6
    return p0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 uint = 18446744073709551615
    var v4 bool = ((!(!false)) || (!(!(42 > 0))))
    v5 := []string{("foo bar" + "foo bar")}
    v4 = (32767 > -721)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v15 := GenS1{F0: 133, F1: 1, F2: 284}
    var v16 int32 = -2147483648
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int32 = 0
        _ = fi17
        _ = v18
    }
    var v19 int = 32767
    if v4 {
        fmt.Println(v3)
        var v20 float64 = (1e10 + float64(float32(1.5)))
        _ = v20
    }
    var v21 int = (v19 & (v19 >> (((^(v3))) & 63)))
    v1[v21] = (0 > -58)
    var v22 uint16 = 843
    var v23 uint64 = 4294967295
    v24 := GenS1{F0: 151, F1: 4294967295, F2: 42}
    v1[(-(v19))] = v4
    v19 -= (v21 % (((v19 | -55)) | 1))
    if v4 {
        switch (v21 / (((v21 * (v21 - (v19 * v21)))) | 1)) {
        case -2147483648:
            var v25 uint8 = 1
            _ = v25
        default:
            v23 += 1
        }
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    }
    _ = len(v2)
    if (1 < 128) {
        switch "foo bar" {
        case "foo bar":
            if (int8(v3) < 1) {
                switch v21 {
                case 948:
                    var v26 float32 = 1e10
                    _ = v26
                case 32768:
                    var v27 bool = (v22 == 0)
                    _ = v27
                default:
                    v21 = (v21 - v21)
                }
                v22 = 893
            } else {
                v21 = v21
                var v28 uint16 = (v22 << ((65535) & 63))
                _ = v28
            }
        case "hello":
            var v29 int16 = 42
            _ = v29
        }
        if ((1e10 / 1e10) > (0.0 * 3.14159)) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v30 uint = (v3 * (v3 * v3))
        _ = v30
    }
    v16 -= (-(v16))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
