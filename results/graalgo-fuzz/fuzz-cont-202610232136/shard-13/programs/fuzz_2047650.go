package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS3 struct { F0 int32; F1 int }
func genFn0() int32 {
    var v6 int8 = 100
    if (-1 == -279) {
        if ((1 == 9223372036854775807) || (2147483647 != 833)) {
            v7 := GenS1{F0: 4294967295, F1: 42}
            if (1 == 2) {
                v8 := []int{32767}
                fmt.Println(v6)
                _ = v8
            }
            _ = v7
        }
        v6 = (v6 >> ((546) & 63))
    } else {
        var v9 int = -115
        switch "a" {
        case "foo bar":
            v6 = v6
        case "hello":
            for fi10 := 0; fi10 < 8; fi10++ {
                for fi11 := 0; fi11 < 1; fi11++ {
                    v12 := GenS2{F0: 0, F1: true, F2: 65535}
                    var v13 float32 = float32((3.14159 + (1e10 * float64(0.0))))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                var v14 int64 = 9223372036854775807
                _ = fi10
                _ = v14
            }
        case "\n\t":
            if (255 >= 0) {
                var v15 int64 = -871
                if (true || false) {
                    v16 := GenS1{F0: 916, F1: 36}
                    fmt.Println(v15, v9)
                    _ = v16
                } else {
                    var v17 int = (^(v9))
                    if false {
                        var v18 uint64 = 18446744073709551615
                        _ = v18
                    }
                    _ = v17
                }
                _ = v15
            } else {
                fmt.Println(v6, v6)
                var v19 bool = (3.14159 == (((-2.25 - -2.25) - (1.5 + 0.0)) * ((0.1 - -2.25) - (1.5 - 0.1))))
                _ = v19
            }
        }
        _ = v9
    }
    var v20 int64 = -9223372036854775808
    _ = v20
    _ = v6
    return 1
}

func main() {
    var v1 float32 = (1e10 + 0.1)
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 int = 0
    var v5 int16 = 1000
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: 0, F1: 0}
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 bool = true
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    v25 := GenS3{F0: genFn0(), F1: 118}
    v26 := v2[int(v5)]
    v26 = (!((!(-9223372036854775808 >= 42)) && v26))
    var v27 int16 = v5
    var v28 uint = 411
    _ = len(v2)
    var v29 uint8 = 1
    v27 += v5
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
