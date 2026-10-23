package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint16 = 0
    v2 := []string{}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    v5 := []uint32{54}
    v1 = 32767
    var v6 int32 = 42
    fmt.Println(v6)
    v6 = v6
    v1 += v1
    for fi7 := 0; fi7 < 17; fi7++ {
        if true {
            v1 += (v1 ^ (v1 ^ v1))
        } else {
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(fi7, v1)
                fmt.Println(v6, fi8)
                v4[(fi8 ^ (fi7 + 42))] = uint32(v1)
                _ = fi8
            }
        }
        v1 += 577
        if true {
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 uint = 475
                if false {
                    fmt.Println(fi9)
                    var v11 string = ("\n\t" + ((("hello" + "\n\t") + "\n\t") + "hello"))
                    _ = v11
                }
                v10 -= (v10 >> (((v10 ^ (v10 % (((v10 % (((v10 ^ 42)) | 1))) | 1)))) & 63))
                _ = fi9
                _ = v10
            }
            var v12 float32 = (0.0 * (0.1 * 1e10))
            v12 -= 1.5
            _ = v12
        } else {
            v13 := GenS1{F0: (v6 | v6), F1: 127}
            fmt.Println(v1, fi7, v1)
            var v14 uint64 = 106
            _ = v13
            _ = v14
        }
        _ = fi7
    }
    switch 32768 {
    case 127:
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int16 = -360
            fmt.Println(v6)
            _ = fi15
            _ = v16
        }
    case 453:
        v2 = append(v2, ("" + "hello"))
    }
    switch -1 {
    case 42:
        var v17 uint8 = 127
        _ = v17
    case 2147483647:
        var v18 int16 = 763
        _ = v18
    default:
        v19 := GenS2{F0: 1, F1: 91, F2: -9223372036854775808}
        _ = v19
    }
    v20 := GenS1{F0: v6, F1: 1}
    v1 -= v1
    var v21 int = 127
    var v22 string = "a"
    if ((float64(0.0) / (0.0 - 1.5)) != 0.0) {
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS1{F0: (v6 + 2147483647), F1: 1}
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 int = (fi25 % (((fi23 / (((fi23 + (fi23 + -2147483648))) | 1))) | 1))
                var v27 int64 = 9223372036854775807
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        v28 := GenS1{F0: -452, F1: 2}
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := GenS3{F0: -410}
            var v31 int = (v21 - (v21 << ((42) & 63)))
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    } else {
        switch v22 {
        case "foo bar":
            switch (v22 + (("a" + "\n\t") + v22)) {
            case "":
                v22 = v22
            case "hello":
                v32 := GenS3{F0: v21}
                _ = v32
            case "a":
                _ = len(v2)
            }
        case "":
            if (!false) {
                switch (v21 >> ((0) & 63)) {
                case 127:
                    _ = len(v4)
                }
            }
        case "\n\t":
            var v33 int32 = v6
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
