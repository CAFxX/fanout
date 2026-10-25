package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint8; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int8; F1 uint8; F2 int8 }
type GenS7 struct { F0 int8; F1 bool }
func genFn0(p0 bool, p1 bool, p2 int64) uint16 {
    v4 := GenS1{F0: -77}
    p0 = ((p2 << ((1) & 63)) != -737)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1() bool {
    var v5 bool = true
    fmt.Println(v5)
    var v6 bool = v5
    v7 := []int{1, 42}
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v1 = 42
    fmt.Println(v1)
    var v8 int8 = 80
    v9 := GenS2{F0: -845, F1: 128, F2: true}
    var v10 int32 = 42
    v10 -= -759
    v8 = int8(v1)
    var v11 bool = false
    for fi12 := 0; fi12 < 6; fi12++ {
        if false {
            fmt.Println(v10, v10, fi12)
            if (0.0 >= (1.5 / (((3.14159 - (((0.0 - 3.14159) - 0.0) - -2.25)) + 0.0) / -2.25))) {
                var v13 float32 = (1e10 * 0.0)
                _ = v13
            } else {
                var v14 int32 = v10
                var v15 bool = v11
                v8 = v8
                _ = v14
                _ = v15
            }
        } else {
            switch fi12 {
            case 128:
                v11 = v11
            case 628:
                switch "hello" {
                case "":
                    v16 := GenS3{F0: -1}
                    _ = v16
                default:
                    switch fi12 {
                    case 9223372036854775807:
                        if (0 >= v1) {
                            v17 := GenS4{F0: 65535}
                            v18 := GenS5{F0: v11, F1: (fi12 | 0)}
                            v19 := []int{(fi12 ^ (fi12 / ((fi12) | 1)))}
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    case 2147483647:
                        v8 += (-(v8))
                    case 96:
                        if true {
                            _ = len(v2)
                        } else {
                            v20 := GenS2{F0: 1000, F1: 42, F2: v11}
                            var v21 float32 = ((0.0 / 0.1) * (1.5 - (0.1 / 1.0)))
                            _ = v20
                            _ = v21
                        }
                    default:
                        v22 := GenS6{F0: (v8 & v8), F1: 227, F2: v8}
                        _ = v22
                    }
                }
            }
            v10 = -1
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(v10, v11, v1)
            _ = fi23
        }
        switch (fi12 % ((fi12) | 1)) {
        case 319:
            if v11 {
                for fi24 := 0; fi24 < 2; fi24++ {
                    v25 := []int{fi12, fi24, 578, (fi12 / ((9223372036854775807) | 1))}
                    var v26 uint64 = 27
                    switch "" {
                    case "a":
                        fmt.Println(v26, v10)
                    }
                    var v27 int32 = 850
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                v1 = v1
            }
        default:
            var v28 uint16 = 1
            _ = v28
        }
        _ = fi12
    }
    var v29 int16 = -32768
    for fi30 := 0; fi30 < 12; fi30++ {
        switch (((("foo bar" + "foo bar") + "\n\t") + "foo bar") + "hello") {
        case "foo bar":
            for fi31 := 0; fi31 < 13; fi31++ {
                if (!v11) {
                    if ((true && v11) || true) {
                        v32 := GenS7{F0: -128, F1: (1 != -9223372036854775808)}
                        fmt.Println(fi30, v1, v10)
                        if v11 {
                            fmt.Println(fi30)
                            var v33 int64 = -1
                            _ = v33
                        }
                        _ = v32
                    } else {
                        _ = len(v2)
                        fmt.Println(fi31)
                        var v34 uint = 18446744073709551615
                        _ = v34
                    }
                    if true {
                        v8 += v8
                        v35 := GenS4{F0: uint(v1)}
                        _ = v35
                    }
                }
                _ = fi31
            }
        case "\n\t":
            var v36 int8 = (v8 - v8)
            _ = v36
        default:
            v29 = (v29 - -321)
        }
        _ = fi30
    }
    switch -927 {
    case 558:
        fmt.Println(v29, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v8
    _ = v9
}
