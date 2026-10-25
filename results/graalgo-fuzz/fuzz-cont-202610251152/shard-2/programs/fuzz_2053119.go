package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 810
    v2 := make(map[string]uint32)
    var v3 float32 = (float32(0.0) + -2.25)
    var v4 float32 = (v3 - v3)
    v5 := []int{42, 32768, 0}
    var v6 int = 891
    if (2 == 4294967295) {
        for fi7 := 0; fi7 < 18; fi7++ {
            if true {
                var v8 int32 = -889
                _ = v8
            }
            if true {
                v3 -= float32((3.14159 * 3.14159))
                if false {
                    v1 += (^(v1))
                    fmt.Println(v6, v1, v6)
                }
            } else {
                var v9 uint = 300
                if (true && true) {
                    v10 := GenS1{F0: -2147483648, F1: 0, F2: v9}
                    var v11 int32 = -512
                    fmt.Println(v6)
                    _ = v10
                    _ = v11
                }
                v12 := GenS2{F0: -1}
                _ = v12
                _ = v9
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                v3 = 0.1
                var v14 uint64 = 4294967295
                _ = fi13
                _ = v14
            }
            _ = fi7
        }
        var v15 float32 = 1.5
        fmt.Println(v15)
        _ = v15
    }
    v16 := GenS2{F0: -9223372036854775808}
    if (4294967295 == 2) {
        var v17 uint64 = 2
        v18 := GenS3{F0: v17}
        _ = v17
        _ = v18
    }
    var v19 float32 = v4
    for fi20 := 0; fi20 < 5; fi20++ {
        v3 -= v3
        fmt.Println(v1, v6, v19)
        _ = fi20
    }
    delete(v2, "a")
    if false {
        v6 += (v6 % ((4294967295) | 1))
        v3 = ((v4 - -2.25) - v3)
        v21 := GenS2{F0: 9223372036854775807}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS1{F0: -719, F1: 1, F2: 42}
            _ = fi22
            _ = v23
        }
    }
    if len(v5) > 0 {
        v5[0] = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
