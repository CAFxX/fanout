package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint16 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint64, p1 int32) int8 {
    v6 := []int{-952}
    switch ("" + "\n\t") {
    case "foo bar":
        v7 := GenS1{F0: -44, F1: -1, F2: 32767}
        _ = v7
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(fi8, p1, p1)
            v9 := GenS1{F0: 1, F1: 824, F2: 32767}
            _ = fi8
            _ = v9
        }
    }
    p0 += (p0 & (p0 << ((4294967295) & 63)))
    var v10 uint8 = 255
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return -103
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 757
    var v3 uint = uint(v2)
    var v4 uint = v3
    var v5 bool = (uint(v2) >= (v4 >> ((v4) & 63)))
    fmt.Println(v4, v4, v4)
    switch "hello" {
    case "":
        v11 := []int{630, 42, 32767}
        _ = v11
    case "foo bar":
        fmt.Println(v5, v5, v2)
    }
    switch -2147483648 {
    case 32767:
        v3 += (v4 * 18446744073709551615)
    case 32768:
        var v12 uint32 = 786
        _ = v12
    case 407:
        if v5 {
            v4 = 4294967295
            fmt.Println(v4, v2)
        } else {
            v13 := GenS2{F0: 42, F1: -1, F2: v2}
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 float32 = (1e10 / 0.1)
            _ = fi14
            _ = v15
        }
    }
    var v16 int32 = 84
    switch -107 {
    case -9223372036854775808:
        if (int16(v2) > 0) {
            v17 := []int{2147483647, 32768, 0}
            var v18 int32 = 981
            _ = v17
            _ = v18
        }
    case 128:
        var v19 uint64 = 42
        _ = v19
    default:
        v20 := GenS3{F0: v16, F1: 0, F2: (^(v2))}
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            v3 = v4
            if (!false) {
                if v5 {
                    var v23 string = "hello"
                    v24 := GenS2{F0: -714, F1: 4294967295, F2: (v2 - (v2 * v2))}
                    _ = len(v1)
                    _ = v23
                    _ = v24
                } else {
                    v25 := []int{(-(fi21))}
                    _ = v25
                }
                v3 += (v3 ^ (v3 & (v4 & 42)))
            }
            v26 := []int{fi21, 595}
            _ = fi22
            _ = v26
        }
        if (!(2 >= 2147483647)) {
            v27 := []int{(^(fi21)), 42}
            for fi28 := 0; fi28 < 20; fi28++ {
                if len(v27) > 0 {
                    v27[0] = 647
                }
                _ = fi28
            }
            _ = v27
        }
        _ = fi21
    }
    v16 = (v16 >> ((v4) & 63))
    v29 := GenS4{F0: 2147483648, F1: 42}
    if v5 {
        delete(v1, 967)
        if (v5 || v5) {
            var v30 int = 9223372036854775807
            if (v5 && v5) {
                var v31 uint = (^(v4))
                v16 = v16
                _ = v31
            }
            v5 = v5
            _ = v30
        }
    }
    v2 -= (v2 & uint16(v3))
    var v32 int64 = -9223372036854775808
    v33 := GenS5{F0: (v16 ^ (v16 * 2147483647))}
    v32 += 517
    switch 32767 {
    case 429:
        v1[423] = "foo bar"
    case -10:
        var v34 int32 = (v16 - 0)
        _ = v34
    default:
        var v35 string = "a"
        _ = v35
    }
    var v36 int = -963
    var v37 string = "\n\t"
    v37 = v37
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
