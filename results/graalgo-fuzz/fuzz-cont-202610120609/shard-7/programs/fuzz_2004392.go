package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 int32 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v1 = v1
    for fi4 := 0; fi4 < 1; fi4++ {
        fmt.Println(v1, v1)
        _ = fi4
    }
    var v5 uint = 1
    var v6 int = 947
    var v7 uint = v5
    switch v6 {
    case -242:
        fmt.Println(v7, v5)
    }
    if (!(true && (v6 > v6))) {
        v5 = v5
        v8 := GenS1{F0: 535, F1: 0}
        v9 := GenS1{F0: 608, F1: 4294967295}
        _ = v8
        _ = v9
    } else {
        v7 += 0
        v1 = "\n\t"
        v10 := GenS2{F0: 6, F1: 42, F2: -2147483648}
        _ = v10
    }
    v6 -= (v6 * v6)
    for fi11 := 0; fi11 < 10; fi11++ {
        v1 = "foo bar"
        fmt.Println(v1)
        v2[4294967295] = v1
        _ = fi11
    }
    if false {
        v6 += v6
        if false {
            var v12 bool = ((4294967295 >= (^(v7))) && false)
            v13 := v2[0]
            _ = v12
            _ = v13
        } else {
            var v14 int = 194
            var v15 int8 = 127
            _ = v14
            _ = v15
        }
    }
    var v16 float32 = -2.25
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float32 = v16
        v16 = 1.5
        v19 := GenS2{F0: 99, F1: 1, F2: 0}
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v5)
    var v20 uint16 = 255
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 int32 = -1
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
