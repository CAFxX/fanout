package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0() int {
    var v4 uint = 1
    v5 := GenS1{F0: -1}
    v6 := []int{}
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            v8 := GenS1{F0: 1}
            fmt.Println(fi7, v4)
            _ = v8
        }
        v4 = v4
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return 704
}
func genFn1(p0 int32) int64 {
    v9 := []int{-9223372036854775808}
    _ = len(v9)
    var v10 int8 = 1
    var v11 int64 = 0
    v10 = v10
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return v11
}

func main() {
    v1 := []uint32{158, 124}
    var v2 int = -668
    var v3 uint8 = 128
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v12 string = "hello"
    _ = len(v1)
    v13 := GenS2{F0: 0}
    for fi14 := 0; fi14 < 13; fi14++ {
        if (508 == 2147483647) {
            switch (fi14 / ((genFn0()) | 1)) {
            case 42:
                fmt.Println(v3)
            case 281:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v15 := GenS3{F0: 883, F1: (fi14 * v2)}
                _ = v15
            }
            switch fi14 {
            case -9223372036854775808:
                v3 += (^(v3))
            case 2147483647:
                v3 = (v3 % ((v3) | 1))
            }
            _ = len(v1)
        }
        v16 := GenS3{F0: 150, F1: 32767}
        v3 = 255
        _ = fi14
        _ = v16
    }
    switch 4294967295 {
    case 0:
        var v17 uint64 = 42
        _ = v17
    default:
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 uint32 = 100
            _ = fi18
            _ = v19
        }
    }
    if len(v1) > 0 {
        v1[0] = 975
    }
    var v20 uint64 = 753
    for fi21 := 0; fi21 < 13; fi21++ {
        switch (fi21 << ((65535) & 63)) {
        case 128:
            v3 += (v3 | v3)
        case -567:
            var v22 string = ("a" + (((("\n\t" + "hello") + v12) + "a") + v12))
            _ = v22
        default:
            fmt.Println(v2, fi21, v3)
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
