package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64) uint32 {
    if false {
        var v7 int = 654
        if true {
            v8 := []int{v7}
            _ = v8
        } else {
            v9 := GenS1{F0: -908, F1: 42, F2: 0}
            fmt.Println(v7, p0)
            _ = v9
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int32 = 2147483647
            _ = fi10
            _ = v11
        }
        _ = v7
    } else {
        p0 += (p0 - (p0 + p0))
        v12 := GenS1{F0: -587, F1: 0, F2: 128}
        v13 := GenS1{F0: 580, F1: 767, F2: 1}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := []int{128}
        _ = fi14
        _ = v15
    }
    p0 = (p0 | p0)
    _ = p0
    return 2147483648
}

func main() {
    var v1 int64 = 756
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 503
    v4 := make(map[int]int)
    var v5 int16 = 1000
    var v6 uint32 = 2147483647
    v16 := []int{-105, 32768, 423}
    var v17 uint64 = v3
    for fi18 := 0; fi18 < 20; fi18++ {
        v1 -= -631
        fmt.Println(v3, v3, v5)
        _ = fi18
    }
    _ = len(v16)
    if len(v16) > 0 {
        v16[0] = 1
    }
    if true {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int16 = int16(fi19)
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 uint32 = v6
            var v23 int8 = 1
            var v24 uint32 = (^(v6))
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 7; fi25++ {
            v26 := GenS2{F0: 2}
            _ = fi25
            _ = v26
        }
    }
    v5 -= -1
    v27 := GenS2{F0: v3}
    v5 = (^(v5))
    var v28 string = (("a" + (("" + "hello") + "hello")) + "a")
    var v29 int = 971
    var v30 uint32 = v6
    if ((v1 & v1) == (v1 << ((139) & 63))) {
        if (-585 <= -444) {
            fmt.Println(v1)
            switch (v29 << ((uint(v17)) & 63)) {
            case 32767:
                v31 := GenS1{F0: 672, F1: v5, F2: 110}
                _ = v31
            }
            var v32 int = 9223372036854775807
            _ = v32
        }
        var v33 uint16 = 2
        _ = v33
    } else {
        v5 += int16(v29)
        var v34 uint64 = 0
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 uint8 = 221
        _ = fi35
        _ = v36
    }
    v28 = v28
    var v37 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
