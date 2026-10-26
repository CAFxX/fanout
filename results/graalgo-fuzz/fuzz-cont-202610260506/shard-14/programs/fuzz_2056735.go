package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int32; F2 int16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 int16 = 42
    v4 := []uint32{}
    switch 2147483647 {
    case -2147483648:
        var v5 int64 = -9223372036854775808
        _ = v5
    case -9223372036854775808:
        fmt.Println(v3)
    default:
        var v6 int = 127
        _ = v6
    }
    if false {
        var v7 float64 = 3.14159
        v7 += (v7 / v7)
        _ = v7
    }
    v3 -= (v3 << ((484) & 63))
    for fi8 := 0; fi8 < 8; fi8++ {
        v3 = 0
        fmt.Println(v3, v3)
        switch "foo bar" {
        case "":
            var v9 string = "foo bar"
            _ = v9
        case "hello":
            v10 := GenS1{F0: 158, F1: 2}
            _ = v10
        case "\n\t":
            for fi11 := 0; fi11 < 10; fi11++ {
                switch -1 {
                case 32767:
                    var v12 int16 = v3
                    _ = v12
                case -595:
                    v3 += v3
                case 2147483647:
                    if (int8(fi11) != 0) {
                        v13 := v1[42]
                        _ = v13
                    } else {
                        fmt.Println(fi8, v3, fi8)
                        var v14 uint = 1
                        _ = v14
                    }
                default:
                    if true {
                        if (false && false) {
                            v3 -= (-(v3))
                        } else {
                            v3 -= (v3 - v3)
                            switch -9223372036854775808 {
                            case -637:
                                v15 := GenS2{F0: 1, F1: 0, F2: (v3 ^ (v3 & v3))}
                                _ = v15
                            case 2147483647:
                                v16 := GenS2{F0: 1, F1: -351, F2: v3}
                                _ = v16
                            }
                        }
                        v3 += 32767
                        v17 := GenS3{F0: 65535, F1: -308}
                        _ = v17
                    }
                }
                v3 -= (v3 & -32768)
                var v18 uint8 = 127
                _ = fi11
                _ = v18
            }
        }
        _ = fi8
    }
    if true {
        var v19 uint32 = 0
        v20 := GenS4{F0: 148}
        _ = len(v2)
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: 65535, F1: -1}
    switch 4294967295 {
    case 587:
        _ = len(v4)
    default:
        if (654 <= 18446744073709551615) {
            delete(v2, -452)
        }
    }
    var v22 int32 = 42
    delete(v2, 32768)
    if true {
        var v23 int = 0
        var v24 uint64 = 2
        _ = v23
        _ = v24
    } else {
        delete(v2, 4294967295)
    }
    if false {
        var v25 int32 = v22
        if (true || (!true)) {
            switch -587 {
            case 4294967295:
                v26 := GenS3{F0: 255, F1: 572}
                _ = v26
            }
            for fi27 := 0; fi27 < 1; fi27++ {
                for fi28 := 0; fi28 < 5; fi28++ {
                    var v29 uint16 = 2
                    if (796 >= -9223372036854775808) {
                        var v30 uint64 = 1
                        switch (fi27 | 142) {
                        case -9223372036854775808:
                            if (("hello" + "foo bar") <= "\n\t") {
                                var v31 int32 = v22
                                v31 += 874
                                _ = v31
                            }
                            var v32 int64 = -1
                            if false {
                                var v33 int = int(v32)
                                _ = v33
                            } else {
                                switch fi28 {
                                case -2147483648:
                                    v22 += (v25 & (v22 << ((42) & 63)))
                                default:
                                    v30 += (v30 ^ (v30 ^ 18446744073709551615))
                                }
                                if true {
                                    var v34 float32 = float32(-2.25)
                                    v34 += v34
                                    _ = v34
                                }
                            }
                            v30 -= (v30 * (v30 | (v30 & (v30 | (v30 / ((v30) | 1))))))
                            _ = v32
                        case 2147483647:
                            v25 += 1
                        default:
                            v35 := GenS4{F0: 9223372036854775807}
                            _ = v35
                        }
                        v36 := GenS2{F0: 127, F1: -77, F2: v3}
                        var v37 uint = 0
                        _ = v30
                        _ = v36
                        _ = v37
                    }
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
        } else {
            var v38 int64 = -9223372036854775808
            _ = len(v4)
            _ = v38
        }
        _ = v25
    } else {
        v39 := GenS2{F0: 128, F1: v22, F2: (v3 << ((12) & 63))}
        delete(v2, 127)
        _ = v39
    }
    var v40 int64 = 1
    var v41 int32 = v22
    fmt.Println(v22, v40, v40)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
