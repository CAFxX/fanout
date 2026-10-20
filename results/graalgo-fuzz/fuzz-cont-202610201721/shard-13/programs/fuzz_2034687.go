package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int8; F1 int; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint }
func genFn0() int16 {
    var v4 uint16 = 2
    var v5 int32 = 2147483647
    if (true || false) {
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(v4, fi6)
            _ = fi6
        }
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int32 = -1
            v5 = (-(v8))
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint = 18446744073709551615
                v11 := []int{(^(fi7)), fi9}
                v8 = v8
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
            _ = v8
        }
        v12 := GenS1{F0: 115}
        var v13 int = 990
        _ = v12
        _ = v13
    }
    var v14 int64 = 172
    v15 := []int{42}
    _ = v14
    _ = v15
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint64) uint16 {
    var v16 string = ""
    if true {
        if false {
            v16 = (v16 + "a")
        }
    } else {
        v17 := []int{853, -1, 128}
        var v18 uint64 = 42
        _ = v17
        _ = v18
    }
    v16 = v16
    for fi19 := 0; fi19 < 9; fi19++ {
        switch -9223372036854775808 {
        case -960:
            for fi20 := 0; fi20 < 5; fi20++ {
                fmt.Println(v16, v16)
                var v21 string = "foo bar"
                var v22 int64 = 9223372036854775807
                _ = fi20
                _ = v21
                _ = v22
            }
        case -800:
            p0 = p0
        default:
            for fi23 := 0; fi23 < 16; fi23++ {
                fmt.Println(v16)
                v24 := []int{(fi23 / ((fi23) | 1))}
                switch "\n\t" {
                case "hello":
                    p0 -= 4294967295
                }
                _ = fi23
                _ = v24
            }
        }
        if (127 < 127) {
            v16 = (v16 + ((v16 + v16) + v16))
            var v25 int8 = 113
            switch ("a" + v16) {
            case "hello":
                v25 += int8(fi19)
            }
            _ = v25
        } else {
            for fi26 := 0; fi26 < 16; fi26++ {
                p0 = (p0 | p0)
                _ = fi26
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := GenS2{F0: 508}
                _ = fi27
                _ = v28
            }
            fmt.Println(p0, p0, p0)
        }
        p0 -= (p0 | (p0 * p0))
        _ = fi19
    }
    _ = p0
    _ = v16
    return 32767
}

func main() {
    var v1 uint8 = 166
    var v2 uint = 847
    var v3 float32 = 0.1
    if (false || false) {
        v2 = 286
    } else {
        v29 := GenS1{F0: (v1 % ((v1) | 1))}
        fmt.Println(v3)
        _ = v29
    }
    fmt.Println(v2)
    var v30 uint32 = 409
    v31 := GenS3{F0: 962, F1: false}
    v32 := GenS4{F0: 127, F1: 4294967295, F2: ""}
    v33 := []int{2147483647, 42}
    v1 -= v1
    var v34 int32 = -1
    for fi35 := 0; fi35 < 8; fi35++ {
        v34 = (v34 & v34)
        _ = fi35
    }
    if (true || (!false)) {
        var v36 string = "\n\t"
        if (!true) {
            switch -925 {
            case 9223372036854775807:
                _ = len(v33)
            default:
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 uint32 = v30
                    v34 = (v34 >> ((uint(v38)) & 63))
                    var v39 string = v36
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            if false {
                var v40 uint64 = 42
                if (2147483648 != (v30 << (((v2 | uint(v30))) & 63))) {
                    var v41 int = -669
                    _ = v41
                } else {
                    v33 = append(v33, 128)
                    for fi42 := 0; fi42 < 7; fi42++ {
                        fmt.Println(v3, v2, v1)
                        v34 = (^(v34))
                        v40 = (v40 >> ((42) & 63))
                        _ = fi42
                    }
                }
                if len(v33) > 0 {
                    v33[0] = 42
                }
                _ = v40
            }
        }
        _ = v36
    }
    v43 := GenS1{F0: 128}
    for fi44 := 0; fi44 < 18; fi44++ {
        var v45 float32 = v3
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        fmt.Println(v34, v34, v34)
        _ = fi46
    }
    if false {
        for fi47 := 0; fi47 < 6; fi47++ {
            var v48 int16 = 42
            v49 := GenS5{F0: -9223372036854775808}
            var v50 uint32 = (v30 << ((v2) & 63))
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
        }
    } else {
        v51 := GenS6{F0: v2}
        _ = v51
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        if false {
            var v53 uint16 = 255
            var v54 uint16 = v53
            _ = v53
            _ = v54
        }
        fmt.Println(v30, v30)
        var v55 int64 = -9223372036854775808
        _ = fi52
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v43
}
