package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    var v1 int8 = 99
    v2 := []int{-2147483648, 0, 218}
    var v3 uint64 = 42
    var v4 int16 = 1
    var v5 int = -936
    v6 := GenS1{F0: 5}
    v4 -= v4
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if true {
            var v7 string = "hello"
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 bool = (v4 < v4)
                    v3 = v3
                    _ = fi9
                    _ = v10
                }
                if (v3 >= v3) {
                    v4 -= (v4 / ((v4) | 1))
                    var v11 uint32 = 42
                    var v12 int32 = -2147483648
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint32 = 2
                    _ = v13
                }
                var v14 uint64 = v3
                _ = fi8
                _ = v14
            }
            v4 -= (^(v4))
            _ = v7
        } else {
            if (v3 >= v3) {
                if len(v2) > 0 {
                    v2[0] = (v5 - v5)
                }
            }
            var v15 uint16 = 65535
            v1 = v1
            _ = v15
        }
        fmt.Println(v1, v4, v3)
    }
    v16 := GenS2{F0: -1}
    var v17 bool = true
    var v18 uint32 = 0
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS2{F0: 32767}
        _ = fi19
        _ = v20
    }
    var v21 int8 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        v21 = 127
        if (60 >= 42) {
            fmt.Println(v1)
            v5 -= (^(fi22))
            var v23 uint8 = 44
            _ = v23
        }
        _ = fi22
    }
    v18 = (^(v18))
    v3 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v24 := GenS3{F0: 42, F1: v21}
    v21 += 43
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
