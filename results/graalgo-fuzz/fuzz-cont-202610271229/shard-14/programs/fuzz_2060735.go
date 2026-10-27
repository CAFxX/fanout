package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    v1 := []uint32{316}
    var v2 int32 = 42
    var v3 uint64 = 18446744073709551615
    v4 := []int{2147483647, 2147483647, -149}
    fmt.Println(v2, v2)
    v5 := GenS1{F0: 255, F1: 127}
    for fi6 := 0; fi6 < 16; fi6++ {
        if true {
            var v7 int8 = -128
            _ = v7
        } else {
            var v8 bool = true
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 uint8 = 1
                fmt.Println(fi6, fi6)
                var v11 int32 = v2
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(fi6, fi6)
            _ = v8
        }
        _ = fi6
    }
    v2 = (v2 | (v2 & v2))
    var v12 float32 = 1.5
    v3 = v3
    v3 = 42
    if false {
        if true {
            switch "" {
            case "\n\t":
                v13 := GenS2{F0: 611, F1: -1}
                _ = v13
            case "hello":
                v2 = v2
            case "foo bar":
                var v14 int32 = 905
                _ = v14
            default:
                var v15 int8 = 0
                _ = v15
            }
        }
        var v16 string = ("" + "\n\t")
        v17 := GenS2{F0: 1, F1: 604}
        _ = v16
        _ = v17
    } else {
        var v18 int8 = 42
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 float64 = 0.0
        for fi21 := 0; fi21 < 8; fi21++ {
            if len(v4) > 0 {
                v4[0] = fi21
            }
            if false {
                var v22 uint8 = 133
                _ = v22
            }
            _ = fi21
        }
        _ = len(v4)
        _ = fi19
        _ = v20
    }
    v3 = (v3 >> ((42) & 63))
    if (!(42 < uint(v3))) {
        for fi23 := 0; fi23 < 17; fi23++ {
            if (!(32767 != 338)) {
                for fi24 := 0; fi24 < 11; fi24++ {
                    v12 += float32(-2.25)
                    v12 += float32(1.5)
                    v2 = -2147483648
                    _ = fi24
                }
                v3 -= v3
            } else {
                var v25 int = (fi23 + 677)
                _ = v25
            }
            v3 = (v3 - v3)
            v12 -= 1e10
            _ = fi23
        }
        fmt.Println(v3)
        v12 += v12
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 float32 = (v12 + (v12 - v12))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            switch (fi26 % (((fi26 / ((fi26) | 1))) | 1)) {
            case 2147483647:
                var v29 uint8 = 53
                _ = v29
            }
            v30 := GenS3{F0: false, F1: (127 != 1)}
            _ = fi28
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    if (42 >= -1) {
        var v31 int16 = 1
        if false {
            fmt.Println(v3)
            if ((false || (true && (0 <= -905))) || true) {
                for fi32 := 0; fi32 < 3; fi32++ {
                    fmt.Println(v2, v3)
                    v1 = append(v1, 4294967295)
                    _ = fi32
                }
            }
        }
        switch 0 {
        case -2147483648:
            v3 = 18446744073709551615
        case -9223372036854775808:
            var v33 uint = 41
            _ = v33
        }
        _ = v31
    }
    var v34 uint16 = 255
    v12 = float32(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
