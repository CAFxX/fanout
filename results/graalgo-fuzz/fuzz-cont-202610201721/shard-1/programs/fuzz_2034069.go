package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    v1 := []string{}
    var v2 uint8 = 127
    v3 := []int{0, 32768}
    var v4 int64 = 9223372036854775807
    var v5 uint8 = v2
    fmt.Println(v4)
    v4 -= 0
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: int(v4), F1: 65535}
        _ = fi6
        _ = v7
    }
    v4 += 751
    v8 := GenS1{F0: 127, F1: 303}
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: fi9, F1: 2}
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS2{F0: 2147483647, F1: (v4 & (v4 & v4)), F2: 100}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    v2 -= v2
    var v13 string = "hello"
    var v14 uint8 = 21
    var v15 uint = 4294967295
    for fi16 := 0; fi16 < 6; fi16++ {
        v5 -= v14
        _ = fi16
    }
    v17 := GenS3{F0: false, F1: 911}
    v4 -= (v4 + v4)
    fmt.Println(v14)
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 bool = (1000 != -1)
        switch 1 {
        case -63:
            v20 := GenS3{F0: v19, F1: 678}
            _ = v20
        case 32767:
            fmt.Println(v5, v2, v4)
        }
        var v21 uint = (v15 | (v15 & v15))
        _ = fi18
        _ = v19
        _ = v21
    }
    if len(v1) > 0 {
        v1[0] = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
