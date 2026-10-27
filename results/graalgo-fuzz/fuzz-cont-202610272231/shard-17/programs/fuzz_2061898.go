package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
func genFn0(p0 uint8, p1 uint8, p2 uint32) bool {
    fmt.Println(p1, p1)
    p0 = 255
    _ = p0
    _ = p1
    _ = p2
    return (!(1 < 1))
}

func main() {
    var v1 uint16 = 312
    var v2 string = ("hello" + "")
    var v3 uint16 = 0
    v4 := []int{-299}
    v5 := []string{}
    var v6 uint64 = 698
    if false {
        v3 += (^(v1))
    }
    v5 = append(v5, (v2 + "hello"))
    for fi7 := 0; fi7 < 5; fi7++ {
        switch (fi7 / (((fi7 - fi7)) | 1)) {
        case 4294967295:
            v6 = v6
        }
        if ((!true) || false) {
            v8 := GenS1{F0: v6}
            for fi9 := 0; fi9 < 7; fi9++ {
                _ = len(v4)
                _ = fi9
            }
            _ = v8
        }
        _ = fi7
    }
    var v10 uint64 = v6
    fmt.Println(v1, v6, v3)
    v2 = v2
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int32 = -1
        _ = fi11
        _ = v12
    }
    switch 0 {
    case -1:
        v13 := GenS2{F0: 940, F1: 975, F2: -2147483648}
        _ = v13
    case 29:
        if false {
            v4 = append(v4, 42)
        }
    case -2147483648:
        var v14 int64 = -1
        _ = v14
    }
    v2 = (v2 + v2)
    if false {
        var v15 uint = 0
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int32 = 42
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        var v18 float32 = 0.1
        _ = v18
    }
    v10 -= v10
    _ = len(v5)
    fmt.Println(v1)
    var v19 int32 = -128
    if false {
        var v20 uint8 = 127
        _ = v20
    } else {
        if len(v4) > 0 {
            v4[0] = int(v10)
        }
        fmt.Println(v2, v19, v1)
    }
    var v21 float32 = 1e10
    var v22 int32 = -122
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
