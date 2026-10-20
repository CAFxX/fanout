package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 string; F1 uint; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 int16; F1 int64; F2 uint8 }
func genFn0() int {
    v5 := []int{9223372036854775807, 983, 127, 972}
    var v6 uint64 = 4294967295
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 113
    var v3 uint16 = v2
    v4 := make(map[int]int)
    var v7 int8 = -1
    v2 = (v3 % (((v2 % ((1) | 1))) | 1))
    var v8 int = 127
    v9 := GenS1{F0: 656}
    if false {
        if false {
            v7 += v7
        } else {
            var v10 int64 = -101
            switch ("foo bar" + ("hello" + "a")) {
            case "foo bar":
                v7 = (v7 >> ((1) & 63))
            case "\n\t":
                fmt.Println(v2, v3, v2)
            default:
                fmt.Println(v8, v3, v10)
            }
            fmt.Println(v2)
            _ = v10
        }
        v2 = v2
    } else {
        switch (^(v8)) {
        case 9223372036854775807:
            var v11 uint32 = 0
            _ = v11
        case 1:
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint16 = (^(v3))
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS2{F0: (fi12 * (v8 | 127)), F1: (-(v7)), F2: v13}
                    v16 := GenS3{F0: 2, F1: 1}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := GenS3{F0: 128, F1: (v2 % (((v13 - 32767)) | 1))}
                    fmt.Println(v2, fi12, v8)
                    switch (fi17 ^ fi12) {
                    case 32768:
                        fmt.Println(v3, fi12)
                    default:
                        v19 := GenS2{F0: fi12, F1: -1, F2: (v2 * (v13 & v2))}
                        _ = v19
                    }
                    v20 := []int{(v8 >> ((2) & 63)), (v8 % ((fi17) | 1))}
                    _ = fi17
                    _ = v18
                    _ = v20
                }
                _ = fi12
                _ = v13
            }
        case 32768:
            var v21 int = v8
            _ = v21
        default:
            fmt.Println(v8, v7, v2)
        }
        v22 := GenS4{F0: "", F1: 0, F2: 1}
        v23 := v4[(^(v8))]
        _ = v22
        _ = v23
    }
    if false {
        for fi24 := 0; fi24 < 16; fi24++ {
            v7 += 100
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int32 = int32(v2)
                v8 = 2147483647
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        _ = len(v4)
    } else {
        v27 := []int{}
        var v28 uint64 = 144
        if (359 >= 564) {
            v29 := GenS5{F0: 0}
            fmt.Println(v7, v7, v7)
            v28 += 4294967295
            _ = v29
        } else {
            switch "" {
            case "":
                v3 = (v2 & (v3 << ((18446744073709551615) & 63)))
            case "hello":
                if true {
                    delete(v1, genFn0())
                    for fi30 := 0; fi30 < 15; fi30++ {
                        if len(v27) > 0 {
                            v27[0] = genFn0()
                        }
                        _ = fi30
                    }
                    fmt.Println(v7, v28, v28)
                }
            case "\n\t":
                v2 -= (v2 << ((255) & 63))
            default:
                for fi31 := 0; fi31 < 20; fi31++ {
                    v32 := GenS6{F0: 128}
                    _ = fi31
                    _ = v32
                }
            }
        }
        _ = v27
        _ = v28
    }
    var v33 uint64 = 1000
    if true {
        if (true || (-937 >= 0)) {
            for fi34 := 0; fi34 < 18; fi34++ {
                var v35 float64 = 1e10
                v3 += (v2 % ((2) | 1))
                _ = fi34
                _ = v35
            }
        } else {
            switch "hello" {
            case "a":
                if (42 <= 4) {
                    switch "" {
                    case "foo bar":
                        var v36 string = "\n\t"
                        _ = v36
                    }
                } else {
                    var v37 int32 = 0
                    switch genFn0() {
                    case 1:
                        fmt.Println(v3)
                    case 840:
                        var v38 int32 = 0
                        _ = v38
                    case 42:
                        var v39 int64 = int64(v37)
                        _ = v39
                    default:
                        for fi40 := 0; fi40 < 17; fi40++ {
                            v41 := GenS4{F0: "foo bar", F1: 676, F2: 42}
                            if true {
                                v4[int(v7)] = int(v37)
                            }
                            v42 := GenS7{F0: 42, F1: v2}
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    }
                    if false {
                        v33 = (v33 << ((42) & 63))
                        for fi43 := 0; fi43 < 7; fi43++ {
                            v33 = (v33 >> ((11) & 63))
                            _ = fi43
                        }
                    } else {
                        var v44 uint64 = 4294967295
                        _ = v44
                    }
                    _ = v37
                }
            case "":
                switch v8 {
                case -9223372036854775808:
                    var v45 int8 = 68
                    _ = v45
                case -2147483648:
                    for fi46 := 0; fi46 < 13; fi46++ {
                        v33 -= (v33 ^ (v33 / ((v33) | 1)))
                        _ = fi46
                    }
                }
            case "hello":
                if (true && false) {
                    switch v8 {
                    case 32768:
                        if false {
                            v7 += -33
                            v47 := GenS2{F0: 1, F1: v7, F2: 255}
                            v48 := GenS8{F0: 0, F1: int64(v33), F2: 127}
                            _ = v47
                            _ = v48
                        }
                    case -171:
                        fmt.Println(v3, v7)
                    case 1:
                        var v49 uint64 = (v33 & 286)
                        _ = v49
                    default:
                        var v50 uint32 = 0
                        _ = v50
                    }
                }
            }
        }
    }
    if (v7 != v7) {
        v51 := v1[(v8 + (v8 ^ (v8 - (v8 >> ((4294967295) & 63)))))]
        _ = v51
    } else {
        var v52 uint64 = 18446744073709551615
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
