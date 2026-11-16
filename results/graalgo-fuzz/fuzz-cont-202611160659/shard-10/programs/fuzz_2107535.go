package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 bool }

func main() {
    var v1 int = 0
    v2 := []uint32{1, 2147483648}
    v3 := []uint32{167, 1, 2147483648}
    v4 := make(map[uint32]uint32)
    var v5 int32 = -2147483648
    var v6 int64 = -9223372036854775808
    v1 = (^(v1))
    switch v1 {
    case 9223372036854775807:
        var v7 int8 = 122
        _ = v7
    case -2147483648:
        var v8 int64 = 0
        _ = v8
    default:
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 42
            v11 := GenS1{F0: 0}
            if (444 > 806) {
                v1 -= v1
                for fi12 := 0; fi12 < 3; fi12++ {
                    v10 = v10
                    _ = fi12
                }
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 float64 = 3.14159
                    v2 = append(v2, 1)
                    v14 = 0.1
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    fmt.Println(v6, v6, v5)
    var v15 float32 = 3.14159
    v5 += 656
    v5 = (-(v5))
    var v16 uint16 = 99
    var v17 uint32 = 917
    var v18 int8 = int8(v5)
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 uint8 = 127
        var v21 float64 = 0.1
        delete(v4, v17)
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if ((v18 >> ((42) & 63)) == (v18 >> ((255) & 63))) {
            for fi23 := 0; fi23 < 9; fi23++ {
                v6 += (v6 % ((v6) | 1))
                var v24 float64 = float64(-2.25)
                v6 = 9223372036854775807
                _ = fi23
                _ = v24
            }
        } else {
            switch "hello" {
            case "hello":
                if false {
                    _ = len(v4)
                    v25 := v4[v17]
                    _ = v25
                }
            }
            var v26 int8 = (v18 ^ 100)
            _ = v26
        }
        v27 := GenS2{F0: 935, F1: int16(v17), F2: 255}
        if true {
            switch (v1 >> ((809) & 63)) {
            case 4294967295:
                for fi28 := 0; fi28 < 4; fi28++ {
                    fmt.Println(v18, v15, fi22)
                    _ = fi28
                }
            case 0:
                if ((v17 % ((v17) | 1)) < v17) {
                    v2 = append(v2, v17)
                    var v29 int64 = v6
                    v30 := GenS3{F0: (^(v17)), F1: (!((v5 & v5) <= (v5 / (((v5 - v5)) | 1))))}
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(v15, fi22, v16)
                    fmt.Println(v1, v5)
                }
            case -2147483648:
                if false {
                    var v31 string = ("" + ("" + "\n\t"))
                    v15 = 0.1
                    _ = v31
                }
            }
        }
        _ = fi22
        _ = v27
    }
    fmt.Println(v5, v5, v5)
    v32 := GenS1{F0: 0}
    v17 = (v17 + (v17 + (^(v17))))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
