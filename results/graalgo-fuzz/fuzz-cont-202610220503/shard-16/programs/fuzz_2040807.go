package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint; F1 string; F2 uint }
func genFn0(p0 int64, p1 uint8, p2 int32) uint16 {
    v5 := GenS1{F0: 0}
    fmt.Println(p1, p0, p0)
    var v6 uint = 2
    var v7 int32 = (p2 % (((p2 & (p2 - p2))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return uint16(p1)
}

func main() {
    var v1 bool = true
    var v2 float32 = float32(1.5)
    v3 := []int{1, 4294967295, 620, 42}
    v4 := make(map[uint32]uint32)
    switch ("hello" + "foo bar") {
    case "hello":
        if false {
            v8 := GenS1{F0: 2}
            _ = v8
        } else {
            if true {
                _ = len(v3)
            }
            _ = len(v3)
        }
    }
    v9 := GenS1{F0: 641}
    var v10 int32 = 0
    _ = len(v4)
    var v11 uint32 = 172
    if (1000 <= 32767) {
        var v12 int = -1
        _ = v12
    }
    v11 = 42
    if v1 {
        fmt.Println(v1, v1)
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(v10)
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        if true {
            switch (-(fi14)) {
            case -313:
                switch "\n\t" {
                case "hello":
                    fmt.Println(v11, v1)
                case "":
                    v15 := GenS2{F0: 372}
                    _ = v15
                case "foo bar":
                    var v16 int32 = v10
                    _ = v16
                default:
                    var v17 int64 = 9223372036854775807
                    _ = v17
                }
            case 4294967295:
                var v18 uint64 = 38
                _ = v18
            }
        }
        if (135 <= 9223372036854775807) {
            v2 -= (v2 / v2)
            var v19 int8 = -1
            var v20 int8 = (v19 + (^(v19)))
            _ = v19
            _ = v20
        } else {
            if len(v3) > 0 {
                v3[0] = fi14
            }
            v10 -= v10
        }
        switch fi14 {
        case -1:
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(fi14, v2)
                _ = fi21
            }
        default:
            switch (fi14 << ((4294967295) & 63)) {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 uint16 = 317
                    _ = fi22
                    _ = v23
                }
            case 856:
                v24 := GenS1{F0: 454}
                _ = v24
            case 42:
                v25 := GenS3{F0: (v11 & v11), F1: 1}
                _ = v25
            }
        }
        _ = fi14
    }
    var v26 uint64 = 0
    fmt.Println(v11, v11, v10)
    v27 := GenS4{F0: 286, F1: "hello", F2: 590}
    v1 = (1 <= 255)
    delete(v4, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
