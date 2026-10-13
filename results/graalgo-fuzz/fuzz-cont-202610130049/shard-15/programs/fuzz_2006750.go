package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint) int32 {
    p0 += p0
    for fi6 := 0; fi6 < 8; fi6++ {
        p0 = (^(p0))
        _ = fi6
    }
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        v7 += (v7 * (^(v7)))
        var v9 uint = (p0 | p0)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return -2147483648
}

func main() {
    var v1 string = (((((((("" + "foo bar") + ("hello" + "a")) + ("" + "foo bar")) + "a") + "hello") + "a") + "foo bar") + (("" + "\n\t") + "\n\t"))
    v2 := []int{721, 2147483647, -446}
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    var v5 float32 = 1.5
    if false {
        v5 -= v5
        if true {
            v1 = v1
            v10 := GenS1{F0: 267, F1: 255, F2: 255}
            for fi11 := 0; fi11 < 17; fi11++ {
                switch v1 {
                case "\n\t":
                    v5 = (v5 - v5)
                }
                _ = fi11
            }
            _ = v10
        }
        switch (v1 + "hello") {
        case "foo bar":
            v5 = (v5 * 0.1)
        }
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int32 = 0
        _ = fi12
        _ = v13
    }
    var v14 int8 = 0
    v15 := GenS1{F0: 1, F1: 2, F2: 65535}
    v16 := v3[0]
    var v17 bool = false
    if (v14 > v14) {
        if v17 {
            v18 := GenS2{F0: 0}
            v5 -= 0.0
            _ = v18
        }
        _ = len(v3)
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            v5 -= 1e10
            v20 := GenS1{F0: 14, F1: 42, F2: 403}
            v16 = v16
            _ = fi19
            _ = v20
        }
        fmt.Println(v16, v14, v16)
        var v21 uint8 = 128
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 int16 = int16(v14)
        var v24 uint16 = 2
        for fi25 := 0; fi25 < 2; fi25++ {
            if (v17 && ((v17 && ((fi25 | (fi25 << ((42) & 63))) != 0)) || v17)) {
                if v17 {
                    v4[0] = 2147483648
                    var v26 uint8 = 63
                    v24 = 0
                    _ = v26
                }
            }
            _ = fi25
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
