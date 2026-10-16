package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 bool; F2 int }

func main() {
    var v1 int8 = 61
    var v2 int8 = v1
    v3 := []int{467, 442, -714}
    v4 := []uint32{4294967295}
    fmt.Println(v1, v2)
    v1 = v2
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v1 = (v2 ^ (v1 - (v2 ^ 0)))
    }
    fmt.Println(v1)
    v2 += 100
    if false {
        var v5 float32 = 0.0
        _ = v5
    } else {
        if true {
            v3 = append(v3, 32767)
            _ = len(v4)
        } else {
            v6 := GenS1{F0: 2147483647, F1: (true || false), F2: -9223372036854775808}
            v7 := GenS2{F0: 622, F1: true, F2: 32768}
            fmt.Println(v2, v2)
            _ = v6
            _ = v7
        }
        var v8 int8 = 0
        v3 = append(v3, 32767)
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            if true {
                fmt.Println(v2, fi10)
                switch (fi9 & (fi10 & fi10)) {
                case 349:
                    _ = len(v4)
                case -186:
                    var v11 uint64 = 856
                    _ = v11
                }
            }
            _ = fi10
        }
        _ = fi9
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
