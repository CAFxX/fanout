package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 int8 }
func genFn0(p0 int8) int {
    fmt.Println(p0)
    p0 += (p0 - -128)
    var v4 uint16 = 2
    _ = p0
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 int64, p1 int32, p2 int32) int8 {
    v5 := GenS1{F0: 0, F1: p1}
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(p1)
        v7 := GenS2{F0: (p1 % ((635) | 1)), F1: -68, F2: 235}
        switch (("a" + "") + "") {
        case "":
            for fi8 := 0; fi8 < 2; fi8++ {
                p0 += p0
                _ = fi8
            }
        case "hello":
            if (4294967295 < 2147483647) {
                fmt.Println(p2, fi6)
                switch (fi6 % ((fi6) | 1)) {
                case 0:
                    for fi9 := 0; fi9 < 20; fi9++ {
                        if true {
                            v10 := []int{0, (^(fi9)), fi6}
                            _ = v10
                        } else {
                            v11 := []int{(fi6 ^ fi9)}
                            _ = v11
                        }
                        fmt.Println(p1, fi6, p2)
                        _ = fi9
                    }
                case 127:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 int16 = 42
                        v13 += v13
                        p0 += (p0 - p0)
                        v14 := []int{genFn0(-1)}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 174:
                    switch (fi6 - genFn0(115)) {
                    case 2147483647:
                        var v15 int8 = 15
                        _ = v15
                    case 127:
                        var v16 int8 = -128
                        _ = v16
                    }
                default:
                    v17 := []int{(fi6 - fi6)}
                    _ = v17
                }
            }
        case "foo bar":
            var v18 uint8 = 1
            _ = v18
        }
        _ = fi6
        _ = v7
    }
    var v19 int32 = -634
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v5
    return 66
}

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 42
    var v3 uint = 0
    if (false || (58 < 255)) {
        switch genFn0(1) {
        case 42:
            v20 := []int{}
            _ = v20
        default:
            v21 := []int{genFn0(1), 128}
            _ = v21
        }
        var v22 string = ""
        for fi23 := 0; fi23 < 1; fi23++ {
            switch (fi23 << ((v3) & 63)) {
            case 9223372036854775807:
                switch (v22 + v22) {
                case "foo bar":
                    switch 2147483647 {
                    case -9223372036854775808:
                        fmt.Println(v1, v2, v1)
                    case -1:
                        switch v22 {
                        case "hello":
                            v24 := []int{fi23, fi23, 128, (fi23 % ((genFn0(0)) | 1))}
                            _ = v24
                        case "a":
                            var v25 string = (v22 + v22)
                            _ = v25
                        case "":
                            v1 = v1
                        }
                    case -954:
                        v2 -= (v2 & v2)
                    default:
                        var v26 uint64 = 2
                        _ = v26
                    }
                case "hello":
                    var v27 float32 = 3.14159
                    _ = v27
                }
            default:
                var v28 bool = ((0.1 != (1.5 + 1e10)) && ((float64(-2.25) - 0.1) >= 1e10))
                _ = v28
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                v30 := GenS3{F0: "\n\t", F1: 42, F2: 1}
                v3 = (v3 - v3)
                var v31 uint8 = 1
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi23
        }
        _ = v22
    }
    v32 := GenS1{F0: 456, F1: 42}
    if false {
        switch -1 {
        case 32768:
            var v33 uint = v3
            _ = v33
        case 1:
            v3 = (v3 ^ (v3 + v3))
        }
        if ((v3 + 0) == 26) {
            switch 127 {
            case 1:
                if ((v2 >> ((v3) & 63)) <= 0) {
                    v34 := GenS2{F0: 42, F1: -1, F2: (v2 - v2)}
                    _ = v34
                } else {
                    v35 := []int{32768, 32767}
                    v36 := GenS4{F0: 127, F1: 784}
                    _ = v35
                    _ = v36
                }
            default:
                var v37 float32 = float32(0.0)
                _ = v37
            }
            for fi38 := 0; fi38 < 6; fi38++ {
                v39 := GenS5{F0: "a", F1: 32767}
                var v40 float32 = 1e10
                v41 := GenS3{F0: "a", F1: 42, F2: 0}
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 1; fi42++ {
                v43 := GenS3{F0: "", F1: 94, F2: 127}
                v44 := GenS6{F0: int(v3), F1: 406}
                _ = fi42
                _ = v43
                _ = v44
            }
        }
    }
    if (true && true) {
        if ((true || false) || true) {
            v2 -= v2
            v45 := GenS7{F0: (v1 - 0)}
            _ = v45
        } else {
            v46 := GenS2{F0: 42, F1: 486, F2: 18446744073709551615}
            _ = v46
        }
        v47 := GenS8{F0: 42}
        _ = v47
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        v1 = v1
        _ = fi48
    }
    var v49 int32 = -650
    switch "hello" {
    case "hello":
        if (895 >= v3) {
            v50 := GenS8{F0: -688}
            _ = v50
        } else {
            if (127 != 1) {
                v51 := GenS9{F0: (v49 << ((255) & 63))}
                _ = v51
            }
            v52 := GenS10{F0: (v49 * -2147483648), F1: v3}
            v53 := GenS11{F0: 42}
            _ = v52
            _ = v53
        }
    default:
        var v54 string = "hello"
        _ = v54
    }
    v2 -= (^(v2))
    var v55 uint64 = 18446744073709551615
    v55 = (v2 >> ((678) & 63))
    switch "hello" {
    case "":
        v1 = 726
    case "a":
        switch 0 {
        case 1:
            var v56 uint32 = 2147483648
            _ = v56
        case -9223372036854775808:
            v3 -= uint(v1)
        case -134:
            if false {
                var v57 int8 = -77
                v55 += (v2 | v2)
                if false {
                    for fi58 := 0; fi58 < 4; fi58++ {
                        v59 := GenS8{F0: 0}
                        v2 += uint64(v3)
                        v60 := []int{(^(fi58)), fi58, 128}
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                }
                _ = v57
            }
        }
    }
    v2 = (v2 >> ((v3) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v49
    _ = v55
}
