package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }

func main() {
    v1 := []int{128, -1, -688, -9223372036854775808}
    var v2 uint = 469
    v3 := make(map[int]string)
    var v4 uint8 = 115
    var v5 int64 = 0
    v6 := GenS1{F0: 452}
    fmt.Println(v5, v2, v5)
    switch 42 {
    case 2147483647:
        for fi7 := 0; fi7 < 9; fi7++ {
            v3[int(v5)] = "foo bar"
            v8 := GenS1{F0: 353}
            v1 = append(v1, fi7)
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            if false {
                if true {
                    fmt.Println(v2, v5, fi9)
                    v11 := GenS2{F0: 32767, F1: 143, F2: 42}
                    _ = v11
                }
                v12 := GenS2{F0: 1, F1: 1, F2: 968}
                _ = v12
            }
            if true {
                v4 += v4
                fmt.Println(v2, fi9)
                if len(v1) > 0 {
                    v1[0] = 9223372036854775807
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = fi10
                }
                v13 := v3[(fi9 ^ (fi9 ^ fi9))]
                var v14 uint8 = (^(v4))
                _ = v13
                _ = v14
            }
            var v15 int8 = -1
            _ = fi10
            _ = v15
        }
        v2 = 42
        var v16 uint = (^(v2))
        _ = fi9
        _ = v16
    }
    if false {
        var v17 uint8 = (v4 + 7)
        var v18 int64 = v5
        if len(v1) > 0 {
            v1[0] = -1
        }
        _ = v17
        _ = v18
    }
    fmt.Println(v2, v2, v4)
    v19 := GenS2{F0: 1, F1: 0, F2: 975}
    var v20 uint16 = 42
    _ = len(v1)
    v21 := GenS2{F0: v20, F1: 1, F2: v20}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
