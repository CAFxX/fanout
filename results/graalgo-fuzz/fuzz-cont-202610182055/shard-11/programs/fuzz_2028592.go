package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 uint8 }
func genFn0() uint64 {
    v4 := GenS1{F0: 0, F1: 31, F2: -173}
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 0, F1: 1, F2: 42}
        fmt.Println(fi5)
        switch -1 {
        case 73:
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 int = (fi7 / ((1) | 1))
                v9 := []int{}
                _ = fi7
                _ = v8
                _ = v9
            }
        case 2147483647:
            v10 := []int{fi5, (fi5 % (((fi5 * fi5)) | 1)), (fi5 + fi5), fi5}
            _ = v10
        case 9223372036854775807:
            var v11 int8 = 123
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    _ = v4
    return 956
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{-987, -1}
    var v3 float64 = ((-2.25 + float64((float32(float64(3.14159)) * (3.14159 + (-2.25 / ((0.1 * 0.0) - float32(-2.25))))))) + float64(0.1))
    var v12 float32 = (3.14159 - float32(1e10))
    _ = len(v1)
    var v13 bool = false
    switch ("" + "a") {
    case "foo bar":
        fmt.Println(v3, v13)
    case "":
        var v14 float64 = (((v3 * ((1.5 / 1.0) * float64(1.5))) - (1.5 - 1.5)) * -2.25)
        _ = v14
    case "hello":
        for fi15 := 0; fi15 < 1; fi15++ {
            v3 -= (3.14159 + (1.5 - 3.14159))
            _ = fi15
        }
    }
    var v16 int64 = -9223372036854775808
    v17 := GenS2{F0: 42, F1: 96, F2: v16}
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int64 = (v16 % ((-1) | 1))
        v3 = v3
        if ((!((v12 - v12) <= 1.5)) || v13) {
            v20 := GenS2{F0: 42, F1: 42, F2: (v16 << ((42) & 63))}
            v21 := GenS2{F0: 0, F1: 42, F2: v16}
            for fi22 := 0; fi22 < 3; fi22++ {
                v23 := GenS3{F0: 1}
                _ = fi22
                _ = v23
            }
            _ = v20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    var v24 bool = (4294967295 == -579)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v16 -= v16
    var v25 uint64 = 101
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
