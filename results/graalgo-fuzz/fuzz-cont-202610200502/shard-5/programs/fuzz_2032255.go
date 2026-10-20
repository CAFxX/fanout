package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint64; F1 string; F2 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 uint8, p2 uint8) bool {
    var v5 int32 = 2147483647
    v5 += (v5 / ((int32(p1)) | 1))
    if ((32767 >= 2) && true) {
        var v6 int8 = -38
        p1 = p1
        v7 := []int{}
        _ = v6
        _ = v7
    }
    fmt.Println(v5, v5, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (!true)
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 uint32 = 2147483647
    v8 := GenS1{F0: -848, F1: (v4 >= v4), F2: 255}
    switch 128 {
    case 2147483647:
        if (-831 > -192) {
            v9 := GenS1{F0: 0, F1: (19 > 62), F2: 31}
            fmt.Println(v4, v1)
            if (844 == 782) {
                var v10 string = ""
                _ = v10
            } else {
                v1 -= uint16(v4)
                if (1.5 != float64(1.5)) {
                    v4 += (v4 - (v4 << ((42) & 63)))
                    for fi11 := 0; fi11 < 3; fi11++ {
                        switch "foo bar" {
                        case "\n\t":
                            delete(v2, (v4 & v4))
                        }
                        _ = fi11
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v12 := GenS1{F0: 312, F1: (65535 <= v1), F2: 42}
                        _ = v12
                    default:
                        v1 -= 42
                    }
                    var v13 uint16 = v1
                    fmt.Println(v4, v13, v4)
                    _ = v13
                }
                v4 += (v4 - 42)
            }
            _ = v9
        } else {
            var v14 int32 = 1
            for fi15 := 0; fi15 < 4; fi15++ {
                v1 = v1
                fmt.Println(v1)
                _ = fi15
            }
            var v16 uint32 = v4
            _ = v14
            _ = v16
        }
    }
    var v17 uint = 501
    v18 := GenS1{F0: int64(v1), F1: false, F2: 127}
    switch 128 {
    case 0:
        v19 := GenS1{F0: 588, F1: false, F2: 255}
        _ = v19
    case 42:
        v4 = 395
    case -779:
        if ((false || (709 != 0)) || (true && true)) {
            var v20 uint64 = 410
            v17 += v17
            _ = v20
        } else {
            v21 := GenS1{F0: 0, F1: (-128 == -103), F2: uint8(v4)}
            _ = v21
        }
    default:
        fmt.Println(v17, v17)
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := GenS2{F0: true, F1: 753}
        v24 := GenS1{F0: 837, F1: (!true), F2: uint8(v4)}
        _ = fi22
        _ = v23
        _ = v24
    }
    if false {
        switch "hello" {
        case "":
            v25 := GenS3{F0: 988, F1: (("" + (("hello" + (("foo bar" + ("a" + "a")) + (("foo bar" + "\n\t") + ""))) + "hello")) + "\n\t"), F2: 2}
            _ = v25
        default:
            v26 := GenS3{F0: 724, F1: "foo bar", F2: 0}
            _ = v26
        }
        if (true || (!false)) {
            v1 -= v1
            var v27 int32 = 42
            _ = v27
        }
    } else {
        v28 := GenS4{F0: -1}
        v17 += (v17 / ((v17) | 1))
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        switch fi29 {
        case -9223372036854775808:
            fmt.Println(v4, v4)
        case 2147483647:
            fmt.Println(v1, v1)
        case -2147483648:
            v30 := v3[""]
            _ = v30
        }
        if false {
            if true {
                switch ("\n\t" + "hello") {
                case "foo bar":
                    v31 := []int{(fi29 - fi29), fi29, (fi29 * fi29)}
                    _ = v31
                default:
                    var v32 int = (fi29 - (^(fi29)))
                    _ = v32
                }
            } else {
                v17 = 255
            }
            var v33 float32 = 1.5
            v34 := []int{fi29, fi29}
            _ = v33
            _ = v34
        }
        _ = fi29
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        fmt.Println(fi35, v17, v17)
        _ = fi35
    }
    var v36 uint16 = v1
    delete(v2, (^(v4)))
    fmt.Println(v36, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v8
}
