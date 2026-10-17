package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 int32 }
type GenS3 struct { F0 int64; F1 string; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int; F2 int }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 int32, p1 int16, p2 int32) uint64 {
    fmt.Println(p0, p1)
    v6 := []int{0, 128, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 bool = true
    var v2 float64 = -2.25
    var v3 int64 = 0
    var v4 uint8 = 1
    var v5 int32 = -2147483648
    var v7 uint = 182
    fmt.Println(v2, v1, v2)
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: -581, F1: (2147483647 >= 42), F2: 265}
        v10 := GenS2{F0: -181, F1: "hello", F2: -216}
        v1 = v1
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 38, F1: (1000 <= -1), F2: 1}
    v12 := GenS1{F0: (-(v3)), F1: v1, F2: uint16(v7)}
    var v13 float64 = 3.14159
    switch 987 {
    case 429:
        var v14 uint16 = 0
        _ = v14
    case 2147483647:
        fmt.Println(v5)
    default:
        var v15 uint8 = (v4 << (((^(v7))) & 63))
        _ = v15
    }
    v16 := GenS1{F0: v3, F1: (!v1), F2: 255}
    var v17 uint32 = 888
    switch 127 {
    case -69:
        switch ("hello" + "hello") {
        case "\n\t":
            v2 -= 1.5
        case "hello":
            fmt.Println(v7)
        case "a":
            if (v1 && false) {
                var v18 float64 = 3.14159
                v19 := GenS1{F0: v3, F1: (255 > (v4 * 168)), F2: 1}
                v5 = v5
                _ = v18
                _ = v19
            } else {
                fmt.Println(v13, v4)
            }
        default:
            if (!false) {
                for fi20 := 0; fi20 < 8; fi20++ {
                    fmt.Println(v7, v17, v7)
                    _ = fi20
                }
            } else {
                if v1 {
                    if v1 {
                        v21 := GenS1{F0: v3, F1: v1, F2: 255}
                        _ = v21
                    } else {
                        v22 := []int{2147483647, 4294967295}
                        _ = v22
                    }
                    v4 -= v4
                }
                if ((v7 + uint(v17)) > 1) {
                    var v23 int16 = 0
                    v24 := GenS2{F0: (^(v3)), F1: "\n\t", F2: v5}
                    _ = v23
                    _ = v24
                }
                var v25 bool = (-286 == 702)
                _ = v25
            }
        }
    case 0:
        v26 := GenS3{F0: v3, F1: "foo bar", F2: 0}
        _ = v26
    }
    switch "a" {
    case "a":
        v5 -= 0
    case "\n\t":
        if true {
            v17 += 2147483647
            switch "hello" {
            case "foo bar":
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 uint = (^(v7))
                    for fi29 := 0; fi29 < 16; fi29++ {
                        var v30 int32 = v5
                        v31 := []int{fi27, fi29, fi27}
                        if len(v31) > 0 {
                            v31[0] = (fi29 << ((v7) & 63))
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
            case "hello":
                v32 := []int{42, 2147483647}
                _ = v32
            default:
                v5 = 1
            }
            var v33 uint8 = 127
            _ = v33
        } else {
            if (65535 == 32767) {
                for fi34 := 0; fi34 < 17; fi34++ {
                    var v35 uint64 = 4294967295
                    _ = fi34
                    _ = v35
                }
                fmt.Println(v2, v17)
            }
            if v1 {
                for fi36 := 0; fi36 < 5; fi36++ {
                    if false {
                        if (!(-2.25 > float32(v13))) {
                            var v37 uint = 255
                            _ = v37
                        } else {
                            v17 = 2147483647
                            for fi38 := 0; fi38 < 9; fi38++ {
                                v17 = (^(v17))
                                var v39 bool = v1
                                _ = fi38
                                _ = v39
                            }
                        }
                    } else {
                        if false {
                            v40 := GenS4{F0: -560}
                            for fi41 := 0; fi41 < 5; fi41++ {
                                var v42 int8 = int8(fi36)
                                v4 -= 0
                                v43 := GenS5{F0: 225, F1: (fi36 % (((fi36 % ((fi41) | 1))) | 1)), F2: fi41}
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        } else {
                            v44 := []int{fi36, 2147483647, (fi36 << (((v7 << (((^(v7))) & 63))) & 63))}
                            var v45 int16 = 42
                            _ = v44
                            _ = v45
                        }
                        v7 += (v7 / (((^(v7))) | 1))
                        for fi46 := 0; fi46 < 6; fi46++ {
                            v47 := GenS6{F0: 1000, F1: v3}
                            v48 := GenS2{F0: (v3 * v3), F1: "hello", F2: (v5 << (((v7 ^ (v7 * v7))) & 63))}
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                    }
                    v49 := []int{fi36}
                    var v50 int = 9223372036854775807
                    _ = fi36
                    _ = v49
                    _ = v50
                }
            }
        }
    default:
        v51 := GenS3{F0: v3, F1: ((("\n\t" + "a") + "foo bar") + "foo bar"), F2: int16(v3)}
        _ = v51
    }
    if false {
        v52 := GenS6{F0: 1, F1: (^(v3))}
        v1 = v1
        v2 -= float64(3.14159)
        _ = v52
    } else {
        var v53 uint64 = 2
        switch 128 {
        case -2147483648:
            fmt.Println(v13, v2)
        }
        fmt.Println(v4, v7, v7)
        _ = v53
    }
    var v54 uint8 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v7
}
