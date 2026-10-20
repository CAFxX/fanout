package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int32 }
func genFn0(p0 int64) uint32 {
    p0 = p0
    fmt.Println(p0)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p0)
    case "a":
        v5 := GenS1{F0: 519}
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: -1}
        _ = v6
    default:
        p0 = (^(p0))
    }
    if false {
        var v7 string = ""
        v8 := []int{-372}
        _ = v7
        _ = v8
    } else {
        var v9 int8 = 100
        p0 += p0
        switch "foo bar" {
        case "a":
            switch "" {
            case "hello":
                for fi10 := 0; fi10 < 3; fi10++ {
                    if true {
                        var v11 uint8 = 127
                        _ = v11
                    }
                    fmt.Println(fi10, p0)
                    _ = fi10
                }
            case "foo bar":
                v12 := []int{188, -9223372036854775808, 42}
                _ = v12
            case "":
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS1{F0: 0}
                    _ = fi13
                    _ = v14
                }
            }
        case "foo bar":
            v15 := []int{9223372036854775807, 1}
            _ = v15
        }
        _ = v9
    }
    v16 := GenS2{F0: (^(p0))}
    _ = p0
    _ = v16
    return 2
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]string)
    var v3 float32 = v1
    var v4 int16 = 1000
    v17 := []int{931, -203, -376}
    v3 = float32(1e10)
    v18 := v2[848]
    v19 := GenS3{F0: 1, F1: genFn0(9223372036854775807), F2: -175}
    fmt.Println(v3)
    var v20 bool = false
    var v21 string = v18
    v20 = true
    if true {
        v3 += -2.25
    }
    v3 = float32(0.1)
    v4 += (v4 >> ((42) & 63))
    var v22 bool = v20
    for fi23 := 0; fi23 < 2; fi23++ {
        v18 = (("" + "foo bar") + v18)
        for fi24 := 0; fi24 < 3; fi24++ {
            v1 -= (v1 / 0.1)
            fmt.Println(v21)
            _ = fi24
        }
        if true {
            switch (fi23 % ((-9223372036854775808) | 1)) {
            case 4294967295:
                switch (((("hello" + "") + v21) + v18) + v21) {
                case "\n\t":
                    var v25 string = v21
                    _ = v25
                case "hello":
                    fmt.Println(fi23)
                default:
                    v3 += v1
                }
            case -1:
                var v26 float32 = float32(((3.14159 * 3.14159) + -2.25))
                _ = v26
            default:
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            v21 = v21
        } else {
            fmt.Println(v21, v21, v4)
            var v27 string = v18
            var v28 float32 = -2.25
            _ = v27
            _ = v28
        }
        _ = fi23
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        _ = len(v2)
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 float32 = (0.1 + ((v1 + v1) / (((-2.25 + v3) + (1.5 - 1.5)) - (0.1 / v3))))
            v17 = append(v17, fi29)
            _ = fi30
            _ = v31
        }
        v32 := GenS1{F0: (^(v4))}
        _ = fi29
        _ = v32
    }
    var v33 uint64 = 1
    if true {
        v34 := GenS3{F0: 12, F1: 203, F2: -627}
        _ = v34
    }
    if len(v17) > 0 {
        v17[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
