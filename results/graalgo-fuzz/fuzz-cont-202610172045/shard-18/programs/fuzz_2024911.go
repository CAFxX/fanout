package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 uint64, p1 bool) string {
    if p1 {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch fi7 {
            case 128:
                fmt.Println(fi7, fi7, fi7)
            case -544:
                switch fi7 {
                case 42:
                    for fi8 := 0; fi8 < 19; fi8++ {
                        var v9 bool = ((p0 << ((65535) & 63)) == p0)
                        v9 = p1
                        v10 := []int{fi8, 351, 42, fi8}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
            case 1:
                p1 = p1
            }
            switch fi7 {
            case -728:
                fmt.Println(fi7, fi7, p0)
            case -2147483648:
                var v11 int8 = 0
                _ = v11
            }
            _ = fi7
        }
    }
    v12 := GenS1{F0: 960, F1: "hello"}
    if p1 {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint16 = 32767
            var v15 uint64 = 18446744073709551615
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 bool = p1
                _ = fi16
                _ = v17
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(p1, p0)
        v18 := []int{0, 128, -127}
        _ = v18
    } else {
        v19 := []int{32768, 9223372036854775807, -9223372036854775808}
        _ = v19
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v12
    return ("\n\t" + "\n\t")
}

func main() {
    var v1 int8 = 100
    var v2 int64 = -263
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    var v5 uint = 0
    var v6 int64 = -136
    if (!(889 > 42)) {
        v2 = (v2 | v2)
        var v20 int64 = -690
        v21 := GenS2{F0: 0, F1: v20}
        _ = v20
        _ = v21
    }
    if (false && (0 < 127)) {
        fmt.Println(v2, v5)
    } else {
        v22 := GenS3{F0: 692, F1: 4294967295}
        _ = v22
    }
    v23 := GenS4{F0: 42}
    v24 := GenS5{F0: 0, F1: v5, F2: 180}
    var v25 float64 = (float64(((1.5 / ((0.1 / 1.5) * 3.14159)) * float32(0.1))) + float64((1e10 * ((1e10 * (3.14159 - 0.0)) / 1.0))))
    for fi26 := 0; fi26 < 5; fi26++ {
        switch fi26 {
        case 32768:
            switch "a" {
            case "\n\t":
                v5 += v5
            case "foo bar":
                v27 := []int{fi26, -2147483648}
                _ = v27
            case "hello":
                switch ("\n\t" + "hello") {
                case "a":
                    switch fi26 {
                    case 533:
                        v25 -= 0.0
                    case -1:
                        var v28 int16 = -1
                        _ = v28
                    }
                }
            default:
                if (true && true) {
                    v2 = v2
                    delete(v3, 1)
                    fmt.Println(v6, fi26)
                }
            }
        default:
            fmt.Println(v1)
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int8 = (v1 % ((int8(v5)) | 1))
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    v31 := []int{127, 0, 832}
    var v32 int64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
