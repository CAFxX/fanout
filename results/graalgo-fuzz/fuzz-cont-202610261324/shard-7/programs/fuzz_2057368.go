package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 uint16 }
type GenS4 struct { F0 int; F1 int; F2 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    var v3 int16 = 42
    v4 := make(map[int]int)
    v3 -= (^(v3))
    switch -2147483648 {
    case 272:
        fmt.Println(v2, v2)
    case -9223372036854775808:
        v3 = (v3 - (v3 * 446))
    case 128:
        var v5 float64 = 3.14159
        _ = v5
    }
    if false {
        switch ("\n\t" + (("\n\t" + "\n\t") + (("\n\t" + "a") + "a"))) {
        case "":
            v3 = v3
        case "a":
            v6 := GenS1{F0: 127}
            _ = v6
        case "foo bar":
            var v7 float32 = 1.5
            _ = v7
        }
    } else {
        fmt.Println(v2, v3)
        delete(v4, -704)
        var v8 float32 = (float32(-2.25) / float32((0.0 + 0.1)))
        _ = v8
    }
    v3 = 32767
    switch 32768 {
    case 32767:
        delete(v4, 9223372036854775807)
    case -1:
        var v9 float64 = 0.1
        _ = v9
    case 399:
        if false {
            var v10 float64 = float64((1.5 + (((0.0 * 0.0) / float32((3.14159 / (0.0 + ((-2.25 + (float64((0.0 - 0.0)) + 0.1)) - (1.5 + 0.1)))))) + -2.25)))
            var v11 uint32 = 2147483647
            _ = v10
            _ = v11
        }
    default:
        v12 := GenS2{F0: (^(v3)), F1: ("hello" + ("" + ("foo bar" + "foo bar"))), F2: -9223372036854775808}
        _ = v12
    }
    var v13 int8 = 0
    switch 0 {
    case -9223372036854775808:
        v1[719] = 127
    case 90:
        v3 += (v3 % ((v3) | 1))
    case 925:
        if false {
            v3 = v3
            fmt.Println(v3)
            for fi14 := 0; fi14 < 7; fi14++ {
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 int16 = v3
                    var v17 uint64 = 0
                    if false {
                        if true {
                            switch "hello" {
                            case "foo bar":
                                if (4294967295 < 2) {
                                    if (false || false) {
                                        var v18 int64 = int64(v16)
                                        switch "hello" {
                                        case "hello":
                                            var v19 uint32 = 788
                                            _ = v19
                                        }
                                        _ = v18
                                    } else {
                                        v20 := GenS2{F0: v16, F1: "hello", F2: 9223372036854775807}
                                        _ = v20
                                    }
                                    v21 := []int{}
                                    var v22 uint32 = 65535
                                    _ = v21
                                    _ = v22
                                } else {
                                    delete(v4, fi15)
                                }
                            }
                            fmt.Println(v13)
                        }
                        var v23 int8 = v13
                        _ = v23
                    } else {
                        v24 := v1[733]
                        var v25 float64 = (1e10 / (1.5 + (0.0 + (float64(0.1) / 1e10))))
                        _ = v24
                        _ = v25
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v13 = -104
                var v26 uint16 = (v2 ^ v2)
                _ = fi14
                _ = v26
            }
        }
    }
    var v27 uint32 = 4294967295
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := []int{}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 float64 = -2.25
        if false {
            _ = len(v1)
        } else {
            v3 -= 1
        }
        _ = fi30
        _ = v31
    }
    v32 := GenS3{F0: 2, F1: 91, F2: uint16(v27)}
    if ((-2.25 != (float64(-2.25) - 0.0)) || true) {
        var v33 int16 = (-(v3))
        v34 := GenS4{F0: -9223372036854775808, F1: 568, F2: (^(v13))}
        v35 := GenS2{F0: v33, F1: "hello", F2: 0}
        _ = v33
        _ = v34
        _ = v35
    } else {
        if false {
            v4[int(v3)] = 389
            var v36 string = "\n\t"
            v37 := []int{}
            _ = v36
            _ = v37
        } else {
            var v38 float32 = 1e10
            _ = v38
        }
    }
    v39 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v39
    _ = v4
}
