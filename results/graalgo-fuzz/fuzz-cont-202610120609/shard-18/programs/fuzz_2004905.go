package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
func genFn0() int8 {
    var v7 int64 = -1
    var v8 uint = 1
    for fi9 := 0; fi9 < 15; fi9++ {
        if false {
            var v10 int16 = 1000
            _ = v10
        }
        _ = fi9
    }
    switch ((("hello" + "foo bar") + "") + "foo bar") {
    case "a":
        v11 := GenS1{F0: -26, F1: ("hello" + "foo bar"), F2: 289}
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: 42, F1: "\n\t", F2: (v8 * v8)}
            var v14 bool = (!false)
            v14 = v14
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := []uint32{4294967295}
    var v2 bool = (127 < 146)
    v3 := make(map[string]uint32)
    v4 := []int{24}
    var v5 bool = (121 <= 2)
    v6 := make(map[int]bool)
    var v15 int32 = -1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v2 = (!v2)
    v16 := GenS2{F0: (v15 | -483)}
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS1{F0: -2147483648, F1: (("" + ("a" + "")) + "foo bar"), F2: 255}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v19 int32 = (v15 << ((208) & 63))
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: 1, F1: 534, F2: 2147483648}
    var v21 uint16 = 65535
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            v21 = v21
            _ = fi24
        }
        _ = fi22
    }
    var v25 uint16 = v21
    v26 := GenS2{F0: v15}
    var v27 uint8 = 199
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
