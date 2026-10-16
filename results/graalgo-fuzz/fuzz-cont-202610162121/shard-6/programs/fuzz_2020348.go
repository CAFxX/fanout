package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64, p1 bool) int32 {
    v4 := []int{2147483647, -1, 0}
    fmt.Println(p0)
    p0 += (p0 << ((184) & 63))
    if p1 {
        fmt.Println(p0, p1)
    } else {
        var v5 int = 42
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return -2147483648
}

func main() {
    v1 := []uint32{4294967295, 42, 596, 206}
    var v2 string = ""
    var v3 uint32 = 2147483647
    if (-1 > 42) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v6 := GenS1{F0: 194, F1: false, F2: 65535}
        var v7 uint16 = 32767
        _ = v6
        _ = v7
    }
    fmt.Println(v3, v3)
    var v8 int32 = 42
    fmt.Println(v2, v2)
    v2 = v2
    var v9 int64 = -9223372036854775808
    var v10 int16 = 1000
    if false {
        var v11 int16 = (v10 & v10)
        var v12 int32 = (v8 % ((2147483647) | 1))
        var v13 int = 9223372036854775807
        _ = v11
        _ = v12
        _ = v13
    } else {
        v2 = "hello"
        v2 = (v2 + (v2 + ("a" + ((("hello" + v2) + v2) + v2))))
        v8 += (v8 / ((-2147483648) | 1))
    }
    v9 += v9
    if true {
        v9 = 42
    } else {
        if ("\n\t" < "") {
            v14 := GenS2{F0: int8(v8), F1: 957}
            for fi15 := 0; fi15 < 5; fi15++ {
                _ = len(v1)
                _ = fi15
            }
            _ = v14
        } else {
            switch v2 {
            case "a":
                var v16 uint16 = 2
                _ = v16
            case "\n\t":
                v17 := GenS2{F0: -128, F1: 58}
                _ = v17
            }
        }
    }
    v18 := GenS3{F0: 127}
    v8 = (^(v8))
    for fi19 := 0; fi19 < 14; fi19++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v20 uint8 = 2
        _ = fi19
        _ = v20
    }
    v9 = (v9 - v9)
    if (!false) {
        v21 := GenS2{F0: -114, F1: 4294967295}
        _ = v21
    }
    v9 += 866
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
