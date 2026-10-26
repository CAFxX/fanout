package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
func genFn0(p0 bool) uint8 {
    p0 = (743 <= 1)
    v6 := GenS1{F0: 451, F1: 32767, F2: 42}
    _ = p0
    _ = v6
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 887
    var v3 uint64 = 4294967295
    var v4 string = "hello"
    var v5 uint = 18446744073709551615
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := v1[4294967295]
        fmt.Println(v2, v8)
        _ = fi7
        _ = v8
    }
    if (!((!(v2 <= 328)) && (648 == 2147483648))) {
        v9 := GenS1{F0: 838, F1: 42, F2: 497}
        _ = v9
    } else {
        fmt.Println(v3, v3, v5)
    }
    v3 += (v3 << ((1) & 63))
    if (((0.1 * ((1.5 + ((0.1 - 1e10) + 1e10)) / ((3.14159 / (float32(1e10) / 1.0)) / 1.0))) + -2.25) >= 1.5) {
        for fi10 := 0; fi10 < 7; fi10++ {
            if true {
                var v11 int = (fi10 * fi10)
                v12 := v1[2147483647]
                if (true && (!(620 != 9223372036854775807))) {
                    var v13 float64 = 3.14159
                    fmt.Println(v2, v5, v4)
                    _ = v13
                } else {
                    v14 := GenS1{F0: 65535, F1: -32768, F2: (v3 - 42)}
                    var v15 float64 = (((3.14159 / 1e10) / 1.0) * -2.25)
                    _ = v14
                    _ = v15
                }
                _ = v11
                _ = v12
            }
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint16 = 1
                var v18 string = "foo bar"
                v19 := []int{(-(fi10)), -1, (fi16 & fi10)}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi10
        }
        var v20 bool = true
        var v21 uint64 = 436
        _ = v20
        _ = v21
    }
    switch -635 {
    case -1:
        var v22 bool = true
        _ = v22
    case -258:
        v23 := []int{1, 9223372036854775807, 836, 32767}
        _ = v23
    case 923:
        if false {
            var v24 int32 = int32(v2)
            if true {
                v2 = v2
                if (true || (false && true)) {
                    fmt.Println(v4, v24)
                    var v25 uint32 = 65535
                    var v26 uint8 = 33
                    _ = v25
                    _ = v26
                }
            }
            v27 := GenS2{F0: true}
            _ = v24
            _ = v27
        }
    }
    var v28 uint = (^(v5))
    switch 949 {
    case -708:
        fmt.Println(v5)
    case 2147483647:
        v28 = v28
    case -641:
        v4 = v4
    default:
        v29 := GenS2{F0: true}
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS3{F0: true}
        _ = fi30
        _ = v31
    }
    v32 := []int{128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
