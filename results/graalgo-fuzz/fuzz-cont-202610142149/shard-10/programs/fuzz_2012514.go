package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }
func genFn0(p0 int8, p1 uint, p2 uint32) uint32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{(fi7 | (fi7 | (^(fi7)))), 21}
            v8 = append(v8, fi6)
            var v9 uint32 = 2147483647
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS1{F0: 1, F1: 524, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2147483647
}

func main() {
    v1 := []uint32{65535, 1, 2147483647, 567}
    var v2 float32 = 1.5
    v3 := []uint32{1, 2147483648, 813, 1}
    var v4 uint = 42
    var v5 float64 = -2.25
    v5 -= (0.0 * (0.0 - (v5 / -2.25)))
    if (4294967295 > 642) {
        var v11 uint32 = 2147483647
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(v4, v11)
            _ = fi12
        }
        switch ("foo bar" + ("hello" + (("hello" + "a") + "\n\t"))) {
        case "":
            switch "a" {
            case "\n\t":
                fmt.Println(v11)
            case "a":
                fmt.Println(v4)
            }
        case "hello":
            var v13 uint = v4
            _ = v13
        case "a":
            v14 := GenS1{F0: 1, F1: 42, F2: 100}
            _ = v14
        }
        _ = v11
    }
    if true {
        fmt.Println(v4)
    } else {
        v5 = (v5 - v5)
        v3 = append(v3, 42)
        fmt.Println(v2)
    }
    switch 4294967295 {
    case -230:
        var v15 uint32 = genFn0(42, (v4 - (^(v4))), 4294967295)
        _ = v15
    case 398:
        var v16 uint8 = 2
        _ = v16
    case 1:
        var v17 uint64 = 2
        _ = v17
    default:
        var v18 int16 = 1000
        _ = v18
    }
    switch 128 {
    case -2147483648:
        if (false || ((42 > 32768) || (!true))) {
            v1 = append(v1, 2147483648)
        } else {
            if false {
                v2 = -2.25
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 int8 = 42
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 3; fi21++ {
                    _ = len(v3)
                    _ = fi21
                }
            }
        }
    case -1:
        for fi22 := 0; fi22 < 11; fi22++ {
            if len(v1) > 0 {
                v1[0] = 42
            }
            v2 = float32(v5)
            _ = fi22
        }
    case 137:
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                v4 += 2
                fmt.Println(v2, v4, v2)
                _ = fi24
            }
            if (0 <= -55) {
                var v25 int32 = -1
                _ = v25
            }
            _ = fi23
        }
    default:
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 uint = (v4 << (((v4 << (((v4 << ((v4) & 63))) & 63))) & 63))
            _ = fi26
            _ = v27
        }
    }
    v2 = v2
    var v28 uint64 = 264
    switch "foo bar" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = genFn0(127, v4, 42)
        }
    case "\n\t":
        for fi29 := 0; fi29 < 18; fi29++ {
            switch "hello" {
            case "":
                var v30 uint16 = uint16(v4)
                _ = v30
            case "hello":
                v31 := GenS1{F0: -1, F1: v28, F2: 73}
                _ = v31
            default:
                var v32 int64 = -974
                _ = v32
            }
            var v33 int8 = 1
            for fi34 := 0; fi34 < 18; fi34++ {
                v35 := GenS1{F0: 146, F1: v28, F2: (v33 >> ((372) & 63))}
                if false {
                    switch 4294967295 {
                    case 808:
                        var v36 uint = (v4 * (^(v4)))
                        _ = v36
                    case 581:
                        var v37 uint64 = 728
                        _ = v37
                    }
                    v38 := GenS1{F0: 0, F1: 18446744073709551615, F2: (v33 + (v33 - (v33 >> ((1) & 63))))}
                    v4 -= (v4 >> ((v4) & 63))
                    _ = v38
                }
                _ = fi34
                _ = v35
            }
            _ = fi29
            _ = v33
        }
    case "a":
        switch -2147483648 {
        case 4294967295:
            for fi39 := 0; fi39 < 6; fi39++ {
                fmt.Println(fi39, v28)
                _ = fi39
            }
        case 32767:
            fmt.Println(v5)
        case 2147483647:
            var v40 float32 = v2
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 11; fi41++ {
        v42 := GenS1{F0: 42, F1: uint64(v4), F2: -1}
        _ = fi41
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
