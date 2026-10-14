package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []int{649}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 uint32 = 2147483647
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 445}
    switch 42 {
    case -734:
        v7 := GenS1{F0: 177}
        _ = v7
    }
    if false {
        v8 := GenS1{F0: 4294967295}
        v4 += (v4 & v4)
        fmt.Println(v4)
        _ = v8
    }
    var v9 uint8 = 1
    fmt.Println(v9, v9)
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS2{F0: 5}
        var v12 uint32 = 2
        var v13 bool = (269 >= 1)
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(v4, v9)
    if false {
        v14 := v2["a"]
        var v15 int64 = 0
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        fmt.Println(fi16, v4, v9)
        _ = fi16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        switch (fi17 % (((-(fi17))) | 1)) {
        case -9223372036854775808:
            var v18 uint8 = (^(v9))
            _ = v18
        case 127:
            v19 := GenS2{F0: 65535}
            _ = v19
        default:
            v20 := GenS1{F0: 811}
            _ = v20
        }
        var v21 string = ("foo bar" + ("foo bar" + ""))
        _ = fi17
        _ = v21
    }
    v22 := GenS3{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
