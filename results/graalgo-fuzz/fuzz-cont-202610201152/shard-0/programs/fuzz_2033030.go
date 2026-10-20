package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 100
    v4 -= (v4 % ((v4) | 1))
    var v5 string = v1
    v5 = v1
    var v6 int8 = v4
    var v7 bool = (42 != -706)
    for fi8 := 0; fi8 < 12; fi8++ {
        if v7 {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 int8 = 100
                v11 := GenS1{F0: -1, F1: 42, F2: 0}
                v12 := v3[0]
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        v13 := []int{(-(fi8)), (fi8 - (fi8 / (((^(fi8))) | 1))), fi8, fi8}
        _ = fi8
        _ = v13
    }
    fmt.Println(v6)
    if (957 != 1) {
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS1{F0: -32768, F1: 42, F2: 2147483647}
            _ = fi14
            _ = v15
        }
        v5 = v1
        var v16 bool = (!true)
        _ = v16
    } else {
        if (!(808 > 816)) {
            var v17 int16 = -158
            _ = v17
        } else {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS1{F0: -421, F1: 0, F2: 772}
                _ = fi18
                _ = v19
            }
            var v20 int8 = (v4 % ((v6) | 1))
            v7 = (!true)
            _ = v20
        }
    }
    v7 = v7
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int16 = 42
        delete(v3, 2147483648)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v24 := []int{(fi23 >> ((65535) & 63)), (fi23 & (fi23 - 128))}
        switch (fi23 | fi23) {
        case 128:
            var v25 int = (fi23 << ((4294967295) & 63))
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 int = (fi26 << ((65535) & 63))
        _ = fi26
        _ = v27
    }
    v6 += v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
