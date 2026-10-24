package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint }
type GenS5 struct { F0 int64 }

func main() {
    var v1 float32 = 0.0
    v2 := []int{}
    var v3 int = 9223372036854775807
    var v4 int = (v3 << ((4294967295) & 63))
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 14; fi6++ {
        if ((!(4294967295 < 2)) || ((!((v3 - (fi6 ^ v4)) != fi6)) || ((0.0 - ((0.0 + (((3.14159 - v1) + (v1 / ((float32((0.0 / -2.25)) * v1) - 1.5))) * 0.1)) - 3.14159)) > v1))) {
            var v7 string = ("hello" + "hello")
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v7
        }
        v4 += fi6
        for fi8 := 0; fi8 < 12; fi8++ {
            if false {
                var v9 uint64 = 300
                if true {
                    v1 = v1
                    delete(v5, (^(v3)))
                    v10 := GenS1{F0: 255, F1: ("\n\t" + "")}
                    v11 := GenS2{F0: 1, F1: -347, F2: -9223372036854775808}
                    if (1 <= 1) {
                        var v12 uint16 = 370
                        v3 = 127
                        v12 = (v12 << ((255) & 63))
                        _ = v12
                    }
                    _ = v10
                    _ = v11
                }
                var v13 float32 = (((v1 - 0.0) / 0.1) + v1)
                var v14 string = ""
                var v15 int16 = 42
                if len(v2) > 0 {
                    v2[0] = (v3 >> ((682) & 63))
                }
                _ = v13
                _ = v14
                _ = v15
                _ = v9
            } else {
                v16 := GenS3{F0: 42}
                v17 := GenS4{F0: 116, F1: 1, F2: 2}
                var v18 int64 = -294
                _ = v16
                _ = v17
                _ = v18
            }
            switch v4 {
            case 528:
                v5[(v4 + 127)] = ""
            case 32767:
                v1 -= v1
            }
            _ = fi8
        }
        _ = fi6
    }
    fmt.Println(v1, v1)
    switch v4 {
    case 2147483647:
        v1 -= 1e10
    case 42:
        var v19 int16 = -134
        _ = v19
    case -394:
        var v20 int8 = 112
        _ = v20
    }
    var v21 uint16 = 32767
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int32 = 2147483647
        fmt.Println(v23, v21, v3)
        _ = fi22
        _ = v23
    }
    v24 := GenS5{F0: -9223372036854775808}
    var v25 int64 = -9223372036854775808
    v26 := GenS1{F0: 65535, F1: ""}
    for fi27 := 0; fi27 < 3; fi27++ {
        v3 = v3
        if (3.14159 <= 0.1) {
            v21 -= v21
        } else {
            v28 := v5[(fi27 * (v3 | 2147483647))]
            v28 = (("foo bar" + v28) + "hello")
            _ = v28
        }
        _ = fi27
    }
    v29 := GenS5{F0: int64(v21)}
    for fi30 := 0; fi30 < 6; fi30++ {
        v4 -= 42
        v5[fi30] = ("a" + "")
        _ = fi30
    }
    v4 += (v4 << ((0) & 63))
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
