package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 int = -681
    v3 := make(map[uint32]int)
    var v4 int16 = v1
    v5 := GenS1{F0: 2147483647, F1: 73}
    v1 = (v4 >> ((2) & 63))
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v2)
        _ = fi6
    }
    var v7 float32 = 1e10
    v7 -= (1e10 + v7)
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 0, F1: 18446744073709551615}
        var v10 uint16 = 32767
        switch "foo bar" {
        case "":
            v2 = 32767
        default:
            for fi11 := 0; fi11 < 5; fi11++ {
                v1 = v4
                if (v7 > 3.14159) {
                    switch "\n\t" {
                    case "foo bar":
                        var v12 uint64 = 88
                        _ = v12
                    case "":
                        switch fi11 {
                        case 9223372036854775807:
                            v13 := GenS1{F0: 2, F1: 1}
                            v14 := []int{(fi8 << ((1) & 63))}
                            if false {
                                var v15 uint32 = 317
                                _ = v15
                            } else {
                                fmt.Println(v4, v7)
                                v16 := GenS2{F0: v10}
                                v4 = 32767
                                if (42 > 100) {
                                    v3[0] = fi11
                                } else {
                                    if (64 >= 84) {
                                        v17 := GenS1{F0: 670, F1: 42}
                                        v2 += v2
                                        delete(v3, 533)
                                        _ = v17
                                    }
                                    var v18 uint32 = 65535
                                    var v19 uint8 = 0
                                    _ = v18
                                    _ = v19
                                }
                                _ = v16
                            }
                            _ = v13
                            _ = v14
                        case -682:
                            var v20 uint = 0
                            _ = v20
                        case 4294967295:
                            if true {
                                if false {
                                    var v21 uint64 = 2
                                    var v22 uint = 40
                                    _ = v21
                                    _ = v22
                                } else {
                                    v1 -= (v1 << ((1) & 63))
                                }
                                delete(v3, 446)
                                var v23 uint32 = 1
                                var v24 uint8 = 42
                                var v25 uint = 798
                                v26 := GenS1{F0: (v23 << ((v25) & 63)), F1: v25}
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                switch (v2 + 931) {
                                case 0:
                                    fmt.Println(v2, v4, v2)
                                case 32768:
                                    if false {
                                        v27 := GenS2{F0: 65535}
                                        v28 := v3[1]
                                        _ = v27
                                        _ = v28
                                    }
                                case 672:
                                    v29 := GenS1{F0: 65535, F1: 41}
                                    v2 += 2147483647
                                    _ = v29
                                }
                                fmt.Println(fi8, fi8, v4)
                                fmt.Println(v2, v2)
                            }
                        default:
                            var v30 int8 = int8(v2)
                            v4 = (v1 + v1)
                            _ = v30
                        }
                    default:
                        v31 := GenS2{F0: 345}
                        _ = v31
                    }
                    var v32 float32 = 1.5
                    _ = v32
                } else {
                    v4 = 32767
                }
                _ = fi11
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v33 int64 = -9223372036854775808
    var v34 uint = 42
    v4 += (v4 * v1)
    switch (v2 * -711) {
    case -32:
        v7 = 1e10
    case 0:
        for fi35 := 0; fi35 < 15; fi35++ {
            if ((-344 > 0) || true) {
                for fi36 := 0; fi36 < 4; fi36++ {
                    switch fi36 {
                    case 128:
                        v1 -= 42
                    case 288:
                        var v37 int64 = v33
                        _ = v37
                    case -9223372036854775808:
                        var v38 bool = (5 >= 127)
                        _ = v38
                    }
                    _ = fi36
                }
                v7 = v7
            }
            _ = fi35
        }
    }
    v4 = v1
    v39 := []int{(v2 & v2), (v2 % ((128) | 1)), v2}
    var v40 uint16 = 42
    delete(v3, 2147483648)
    delete(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v7
}
