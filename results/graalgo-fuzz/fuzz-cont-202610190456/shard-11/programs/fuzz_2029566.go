package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int32 = 64
    var v2 uint = 791
    v3 := []int{821, 32767, -1, -304}
    v4 := make(map[int]bool)
    var v5 int = 290
    var v6 float32 = 1e10
    v6 = (v6 + 0.0)
    var v7 int32 = (^(v1))
    v1 -= v7
    switch "hello" {
    case "\n\t":
        if ((828 <= 993) || (-1 >= 1000)) {
            if false {
                fmt.Println(v6)
            }
        } else {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: -1, F1: 236, F2: 1}
                v10 := GenS1{F0: 25, F1: 48, F2: 880}
                for fi11 := 0; fi11 < 2; fi11++ {
                    _ = len(v3)
                    _ = fi11
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(v2, v5)
            var v12 int16 = -1
            _ = v12
        }
    case "hello":
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS2{F0: 698}
            fmt.Println(fi13)
            for fi15 := 0; fi15 < 1; fi15++ {
                v6 -= v6
                if true {
                    switch "foo bar" {
                    case "foo bar":
                        if (true || true) {
                            if (!false) {
                                if (20 == v2) {
                                    var v16 int64 = 0
                                    var v17 float32 = (((v6 / v6) * 3.14159) * float32((3.14159 / 1.5)))
                                    _ = v16
                                    _ = v17
                                } else {
                                    switch "" {
                                    case "hello":
                                        if (!true) {
                                            v18 := GenS2{F0: -624}
                                            _ = v18
                                        }
                                    case "":
                                        fmt.Println(v5, v6)
                                    default:
                                        var v19 bool = false
                                        _ = v19
                                    }
                                    var v20 float64 = 0.0
                                    _ = v20
                                }
                                v21 := GenS3{F0: -317, F1: 1, F2: 351}
                                _ = v21
                            }
                        } else {
                            var v22 uint64 = 4294967295
                            if true {
                                var v23 int32 = (v7 / ((0) | 1))
                                var v24 int = v5
                                v22 -= 42
                                _ = v23
                                _ = v24
                            }
                            var v25 float64 = -2.25
                            fmt.Println(v25)
                            _ = v22
                            _ = v25
                        }
                    }
                    var v26 uint = v2
                    _ = len(v3)
                    _ = v26
                } else {
                    fmt.Println(v2)
                    if (255 <= 2) {
                        if (!true) {
                            _ = len(v3)
                        }
                    } else {
                        v6 += 3.14159
                        var v27 int16 = 0
                        _ = v27
                    }
                }
                if (-1 == 42) {
                    var v28 uint32 = 2
                    var v29 uint64 = 1
                    if (fi13 <= fi13) {
                        if false {
                            var v30 uint16 = 255
                            var v31 int32 = v1
                            _ = v30
                            _ = v31
                        } else {
                            fmt.Println(fi13)
                            var v32 int8 = -24
                            v7 -= -104
                            _ = v32
                        }
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    default:
        var v33 int32 = (v7 * v7)
        _ = v33
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v4)
    fmt.Println(v5)
    var v34 float64 = (1.5 * (3.14159 * -2.25))
    if (-21 >= -1) {
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 uint = v2
            v37 := GenS4{F0: 0}
            v2 += v2
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = len(v4)
        v7 = (v1 << (((v2 | 158)) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
