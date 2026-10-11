package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64; F1 uint; F2 int32 }
type GenS4 struct { F0 uint64 }
func genFn0() uint8 {
    v5 := []int{-2147483648, 21}
    v6 := GenS1{F0: 2, F1: "", F2: -212}
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 string = ""
    var v2 uint = 0
    var v3 int16 = 1
    var v4 float32 = 0.1
    var v7 uint8 = 0
    v1 = v1
    for fi8 := 0; fi8 < 12; fi8++ {
        if (v1 >= "hello") {
            v3 += (v3 / ((v3) | 1))
        } else {
            var v9 uint = v2
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            v1 = v1
            v11 := []int{(fi10 / ((797) | 1)), (fi10 / (((-(fi10))) | 1)), -9223372036854775808, 32768}
            if true {
                v12 := GenS2{F0: 1, F1: 4294967295}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
    }
    v13 := GenS1{F0: (v7 % ((genFn0()) | 1)), F1: (("\n\t" + v1) + ((v1 + "a") + v1)), F2: v3}
    v14 := GenS3{F0: 415, F1: 0, F2: 0}
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v3)
        fmt.Println(v3, v1)
        if false {
            v16 := []int{}
            var v17 int32 = int32(fi15)
            var v18 uint64 = 1
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    var v19 uint8 = (v7 * v7)
    if (0 > 0) {
        v19 = v7
        v20 := []int{42}
        _ = v20
    }
    fmt.Println(v4)
    v21 := GenS4{F0: 871}
    for fi22 := 0; fi22 < 10; fi22++ {
        v7 = v19
        _ = fi22
    }
    v7 += v19
    var v23 int = -1
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 uint = (v2 << (((v2 | v2)) & 63))
        _ = fi24
        _ = v25
    }
    v7 = (v7 ^ v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
