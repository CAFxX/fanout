package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []uint32{2147483647, 2147483648, 427, 541}
    var v2 string = (("hello" + ("hello" + "\n\t")) + "\n\t")
    v3 := make(map[int]int)
    var v4 uint64 = 18446744073709551615
    var v5 uint = 0
    v6 := make(map[uint32]uint32)
    v2 = ((v2 + "foo bar") + v2)
    v5 -= 2
    v1 = append(v1, 164)
    v7 := v6[2]
    if true {
        v5 = v5
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 float32 = 3.14159
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 234, F1: 706}
        _ = v10
    } else {
        v2 = v2
        if (65535 != (v5 + (v5 + 0))) {
            var v11 int32 = 1
            fmt.Println(v5, v5)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                for fi13 := 0; fi13 < 5; fi13++ {
                    if ((fi13 | fi12) >= fi13) {
                        delete(v3, fi12)
                        if true {
                            v2 = ((((v2 + v2) + "a") + v2) + v2)
                            v4 = 4294967295
                        }
                        v1 = append(v1, (v7 >> (((v5 | v5)) & 63)))
                        v5 = (v5 ^ v5)
                        if true {
                            v14 := GenS1{F0: 42, F1: 1}
                            var v15 int32 = 859
                            _ = v14
                            _ = v15
                        } else {
                            if true {
                                switch 127 {
                                case 128:
                                    switch 42 {
                                    case 1:
                                        var v16 int8 = -16
                                        _ = v16
                                    default:
                                        if false {
                                            v17 := GenS2{F0: (true || (594 < -32768))}
                                            _ = len(v6)
                                            if len(v1) > 0 {
                                                v1[0] = (v7 << ((v5) & 63))
                                            }
                                            _ = v17
                                        }
                                    }
                                case 42:
                                    v2 = v2
                                default:
                                    if false {
                                        fmt.Println(fi12)
                                        var v18 float32 = 0.1
                                        _ = v18
                                    }
                                    var v19 uint32 = uint32(v5)
                                    _ = v19
                                }
                                var v20 int16 = 1
                                _ = v20
                            } else {
                                v7 += (v7 & v7)
                                var v21 bool = false
                                var v22 uint32 = 111
                                _ = v21
                                _ = v22
                            }
                            if true {
                                _ = len(v1)
                            } else {
                                fmt.Println(v5, fi13, fi12)
                                var v23 float32 = 1.5
                                _ = v23
                            }
                            v24 := GenS2{F0: (!false)}
                            if (9223372036854775807 == 42) {
                                v4 -= (v4 << ((v5) & 63))
                            } else {
                                var v25 int = -2147483648
                                v7 += 2147483647
                                _ = v25
                            }
                            v26 := GenS2{F0: (false || true)}
                            _ = v24
                            _ = v26
                        }
                    } else {
                        v4 = uint64(v5)
                        fmt.Println(v2)
                        if (0 <= -95) {
                            v5 += v5
                            v4 = 1
                            switch (^(fi12)) {
                            case 128:
                                v4 = 537
                            default:
                                switch fi12 {
                                case 32768:
                                    v27 := GenS3{F0: 42, F1: 32767}
                                    _ = v27
                                case 520:
                                    var v28 uint64 = 2
                                    _ = v28
                                case -996:
                                    v3[fi13] = fi12
                                }
                            }
                        } else {
                            _ = len(v1)
                            if len(v1) > 0 {
                                v1[0] = (v7 | (^(v7)))
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    v29 := GenS3{F0: -98, F1: 376}
                    _ = fi13
                    _ = v29
                }
                var v30 int32 = 42
                _ = fi12
                _ = v30
            }
            fmt.Println(v2, v2)
        }
        v31 := GenS2{F0: (!(true || true))}
        _ = v31
    }
    if true {
        v5 = v5
    }
    v32 := GenS3{F0: 42, F1: 1}
    for fi33 := 0; fi33 < 11; fi33++ {
        if true {
            v34 := GenS4{F0: -2147483648}
            v5 += (v5 + uint(v7))
            _ = v34
        } else {
            v35 := GenS5{F0: -1}
            var v36 int8 = 0
            _ = v35
            _ = v36
        }
        v37 := v3[fi33]
        _ = fi33
        _ = v37
    }
    v7 += (^(v7))
    var v38 string = (((v2 + "") + "a") + "\n\t")
    var v39 uint = v5
    _ = len(v1)
    if true {
        var v40 int64 = -1
        var v41 float64 = 1e10
        _ = v40
        _ = v41
    }
    v2 = (v2 + (v38 + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
