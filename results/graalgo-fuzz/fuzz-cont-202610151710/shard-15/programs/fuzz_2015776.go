package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 string = "\n\t"
    v5 := make(map[int]uint32)
    v6 := []int{32768, -664, 32768, 2147483647}
    var v7 int8 = 7
    delete(v2, v4)
    v7 += (v7 - v7)
    var v8 bool = (v7 >= 36)
    fmt.Println(v4)
    var v9 int8 = v7
    switch 32767 {
    case -9223372036854775808:
        v10 := GenS1{F0: v7, F1: 112, F2: 127}
        _ = v10
    case 0:
        v2[(v4 + "hello")] = 65535
    default:
        delete(v5, 771)
    }
    v11 := GenS1{F0: v7, F1: 128, F2: (v7 % (((v7 - (-(v7)))) | 1))}
    var v12 int16 = -32768
    v4 = (((v4 + v4) + v4) + "foo bar")
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            v4 = "\n\t"
            switch (fi14 + (fi14 << ((849) & 63))) {
            case -496:
                if v8 {
                    v15 := GenS1{F0: 127, F1: 255, F2: v9}
                    _ = v15
                }
            case 128:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                var v16 bool = v8
                _ = v16
            }
            _ = fi14
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            switch "foo bar" {
            case "foo bar":
                var v18 int64 = -455
                _ = v18
            }
            switch fi17 {
            case 32768:
                if v8 {
                    var v19 int16 = (v12 % (((v12 & 0)) | 1))
                    var v20 int = fi17
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint16 = 1
                    var v22 int = fi17
                    _ = v21
                    _ = v22
                }
                fmt.Println(v8, v12)
                if ((true && (0 == 18446744073709551615)) && (!(2 != 4294967295))) {
                    fmt.Println(v9, fi13, fi13)
                    fmt.Println(fi13, v7, v4)
                }
            case 9223372036854775807:
                v23 := v2[v4]
                _ = v23
            default:
                var v24 uint16 = 2
                _ = v24
            }
            _ = fi17
        }
        _ = fi13
    }
    var v25 uint16 = 1
    if v8 {
        var v26 int64 = 822
        if v8 {
            if (180 <= 1) {
                _ = len(v5)
            } else {
                v25 += (^(v25))
            }
            for fi27 := 0; fi27 < 3; fi27++ {
                if ((v26 ^ 9223372036854775807) != v26) {
                    var v28 int32 = 0
                    v29 := GenS2{F0: (fi27 << ((2) & 63))}
                    _ = v28
                    _ = v29
                }
                v4 = v4
                _ = fi27
            }
            v30 := GenS1{F0: -1, F1: 42, F2: v9}
            _ = v30
        } else {
            delete(v1, 4294967295)
            v31 := GenS3{F0: 0}
            fmt.Println(v26, v25)
            _ = v31
        }
        _ = v26
    }
    var v32 uint8 = 127
    var v33 float64 = 0.1
    v34 := GenS1{F0: v7, F1: v32, F2: v9}
    _ = len(v5)
    var v35 int16 = (v12 - -32768)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
