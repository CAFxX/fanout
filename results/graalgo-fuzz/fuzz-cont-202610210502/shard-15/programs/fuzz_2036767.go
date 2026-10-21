package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint8; F2 bool }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint8, p1 bool, p2 uint16) int16 {
    p0 = (p0 << ((2) & 63))
    v6 := []int{-9223372036854775808, 127}
    var v7 uint8 = (p0 & 127)
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            p1 = p1
            v10 := GenS1{F0: 1, F1: (-(fi8))}
            if p1 {
                var v11 int64 = 42
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -652
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 65535
    v3 := make(map[uint32]uint32)
    var v4 uint = 0
    var v5 int8 = -128
    switch ("\n\t" + "hello") {
    case "":
        var v12 bool = false
        _ = v12
    default:
        fmt.Println(v2)
    }
    var v13 float32 = v1
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 int32 = 2147483647
            switch (^(fi14)) {
            case 887:
                v17 := []int{}
                _ = v17
            case 127:
                v18 := GenS2{F0: (^(fi15))}
                _ = v18
            case -845:
                v19 := GenS3{F0: (-(fi14)), F1: 54, F2: true}
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        var v20 int16 = -1
        v21 := GenS4{F0: -1, F1: int8(v2)}
        _ = fi14
        _ = v20
        _ = v21
    }
    v22 := []int{127, 32767}
    if false {
        v4 -= 472
        var v23 uint32 = 2147483648
        _ = v23
    } else {
        var v24 int32 = -957
        for fi25 := 0; fi25 < 6; fi25++ {
            switch "foo bar" {
            case "\n\t":
                v26 := GenS4{F0: v24, F1: v5}
                _ = v26
            default:
                v5 = v5
            }
            var v27 bool = (-201 <= 1)
            switch fi25 {
            case 128:
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 uint16 = v2
                    if (uint32(v2) != 0) {
                        var v30 bool = ((v5 * (v5 & v5)) < v5)
                        fmt.Println(v24, v30)
                        _ = v30
                    } else {
                        var v31 uint64 = 0
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                }
            case -1:
                fmt.Println(v13)
            default:
                for fi32 := 0; fi32 < 10; fi32++ {
                    if (v27 && (v2 >= 0)) {
                        delete(v3, 4294967295)
                    } else {
                        v33 := v3[2147483647]
                        var v34 uint64 = 53
                        var v35 int32 = 1
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    _ = fi32
                }
            }
            _ = fi25
            _ = v27
        }
        _ = v24
    }
    var v36 bool = false
    switch 270 {
    case -2147483648:
        if (-323 <= 4294967295) {
            v4 = v4
        }
    case -397:
        if v36 {
            var v37 uint32 = 2147483648
            v38 := GenS5{F0: v2}
            var v39 uint = (^(v4))
            _ = v37
            _ = v38
            _ = v39
        } else {
            var v40 uint8 = 59
            var v41 int64 = -9223372036854775808
            for fi42 := 0; fi42 < 4; fi42++ {
                var v43 uint16 = 1
                v22 = append(v22, fi42)
                for fi44 := 0; fi44 < 20; fi44++ {
                    v45 := GenS2{F0: 4294967295}
                    _ = fi44
                    _ = v45
                }
                _ = fi42
                _ = v43
            }
            _ = v40
            _ = v41
        }
    default:
        v1 += (v13 * v1)
    }
    switch "" {
    case "":
        var v46 uint8 = 1
        _ = v46
    }
    v36 = false
    switch "" {
    case "\n\t":
        for fi47 := 0; fi47 < 18; fi47++ {
            fmt.Println(fi47, v36, v1)
            _ = fi47
        }
    case "":
        fmt.Println(v2, v2)
    case "hello":
        delete(v3, 0)
    default:
        v2 = v2
    }
    for fi48 := 0; fi48 < 12; fi48++ {
        switch "hello" {
        case "a":
            v13 += v13
        case "":
            fmt.Println(v2, fi48, fi48)
        case "\n\t":
            for fi49 := 0; fi49 < 7; fi49++ {
                fmt.Println(v2, fi49)
                _ = fi49
            }
        default:
            if (9223372036854775807 < -9223372036854775808) {
                var v50 bool = (2147483647 != 15)
                _ = v50
            }
        }
        v2 -= (v2 * 2)
        _ = fi48
    }
    var v51 uint8 = 255
    var v52 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
