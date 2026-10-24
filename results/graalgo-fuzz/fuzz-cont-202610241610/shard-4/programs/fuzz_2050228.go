package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint64; F2 int32 }
type GenS4 struct { F0 int32 }
func genFn0() int8 {
    var v7 uint64 = 815
    if (160 <= 16) {
        var v8 int8 = 0
        _ = v8
    } else {
        v9 := []int{}
        _ = v9
    }
    fmt.Println(v7)
    var v10 uint64 = v7
    _ = v10
    _ = v7
    return 0
}
func genFn1(p0 uint64, p1 int64) int8 {
    var v11 float32 = 3.14159
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: -50}
            v15 := []int{127, 42, (fi12 >> ((825) & 63)), fi12}
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: (p1 - p1), F1: 961, F2: 2147483647}
        _ = fi12
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v11
    return 127
}

func main() {
    var v1 int8 = 82
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 int32 = 1
    v5 := make(map[uint32]bool)
    var v6 uint8 = 2
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int32 = v4
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint64 = 386
            fmt.Println(v4, fi20)
            if ((v4 + v4) != (v4 & v4)) {
                var v22 int64 = -723
                fmt.Println(fi20, v1)
                _ = v22
            } else {
                var v23 int64 = -1
                v1 = (-(v1))
                _ = v23
            }
            _ = fi20
            _ = v21
        }
        if true {
            if true {
                var v24 uint32 = 28
                _ = v24
            } else {
                for fi25 := 0; fi25 < 6; fi25++ {
                    if (0.0 > -2.25) {
                        fmt.Println(v4, fi25, fi25)
                    } else {
                        switch "hello" {
                        case "a":
                            fmt.Println(v6)
                        case "hello":
                            v4 -= v4
                        case "foo bar":
                            switch ("hello" + "foo bar") {
                            case "a":
                                var v26 float32 = ((0.1 * 0.0) + 1e10)
                                _ = v26
                            default:
                                var v27 int8 = v1
                                _ = v27
                            }
                        default:
                            v28 := GenS2{F0: -1, F1: 65535, F2: 578}
                            _ = v28
                        }
                        var v29 int32 = (v4 << ((18446744073709551615) & 63))
                        v1 -= v1
                        v30 := GenS3{F0: -916, F1: 286, F2: -1}
                        if false {
                            v31 := GenS4{F0: (v29 & int32(v1))}
                            _ = v31
                        } else {
                            fmt.Println(v1)
                            v4 -= v4
                            var v32 uint16 = 65535
                            _ = v32
                        }
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                }
                v4 = (v4 | v4)
            }
            var v33 float32 = 1.5
            _ = v33
        } else {
            v4 += (v4 ^ v4)
        }
        _ = fi19
    }
    fmt.Println(v4, v6)
    var v34 float64 = ((0.0 + 1.5) + (0.0 * float64(float32(0.1))))
    v35 := GenS3{F0: -293, F1: 0, F2: v4}
    v1 -= (v1 + v1)
    v4 -= (v4 + (^(v4)))
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 int16 = -32768
        _ = fi36
        _ = v37
    }
    fmt.Println(v34, v34)
    fmt.Println(v4)
    v1 = 13
    if (0.1 < (0.0 + 3.14159)) {
        fmt.Println(v34)
        var v38 uint32 = 2
        _ = v38
    } else {
        v4 += (v4 + (^(v4)))
    }
    var v39 int8 = v1
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v40 uint = 774
        _ = v40
    case "a":
        var v41 float32 = 0.0
        _ = v41
    case "hello":
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
