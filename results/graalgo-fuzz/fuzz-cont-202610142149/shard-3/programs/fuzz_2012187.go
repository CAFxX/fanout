package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }

func main() {
    var v1 uint64 = 402
    v2 := make(map[int]int)
    var v3 float64 = 1e10
    var v4 int64 = -36
    v5 := make(map[uint32]bool)
    if false {
        v6 := GenS1{F0: 42, F1: (v1 << ((42) & 63))}
        _ = v6
    }
    var v7 int32 = 1
    v7 = v7
    var v8 uint16 = 42
    var v9 float64 = v3
    v10 := GenS2{F0: (v7 - (v7 ^ v7))}
    switch "hello" {
    case "hello":
        switch -782 {
        case 128:
            fmt.Println(v7, v4, v4)
        case 32767:
            for fi11 := 0; fi11 < 2; fi11++ {
                v3 -= -2.25
                _ = fi11
            }
        default:
            switch ("a" + "a") {
            case "a":
                v12 := GenS2{F0: 524}
                _ = v12
            case "foo bar":
                if ((!(v1 <= v1)) && false) {
                    var v13 uint = 840
                    v4 = (v4 * (v4 ^ v4))
                    _ = v13
                } else {
                    v1 = (v1 ^ (^(v1)))
                    v14 := GenS1{F0: 40, F1: v1}
                    var v15 uint32 = 522
                    _ = v14
                    _ = v15
                }
            }
        }
    case "":
        if (!true) {
            fmt.Println(v4, v9)
            switch 1 {
            case -2147483648:
                switch "hello" {
                case "hello":
                    if (32768 == 8) {
                        v16 := GenS3{F0: ("foo bar" + ("\n\t" + ("\n\t" + "a"))), F1: -128}
                        _ = v16
                    } else {
                        v3 = v9
                        v17 := []int{int(v1), 9223372036854775807, -9223372036854775808}
                        var v18 uint = 42
                        _ = v17
                        _ = v18
                    }
                case "a":
                    v9 = -2.25
                case "\n\t":
                    v4 += v4
                default:
                    if false {
                        v1 = (v1 ^ 4294967295)
                        var v19 uint = 899
                        _ = v19
                    } else {
                        v8 = v8
                    }
                }
            case -1:
                v5[uint32(v8)] = true
            }
        } else {
            var v20 uint64 = (v1 * v1)
            for fi21 := 0; fi21 < 14; fi21++ {
                switch (fi21 << ((65535) & 63)) {
                case -658:
                    for fi22 := 0; fi22 < 20; fi22++ {
                        if (199 >= -1) {
                            switch ("" + "") {
                            case "a":
                                v23 := []int{(fi22 + (fi21 * 42)), (^(fi21)), 9223372036854775807}
                                _ = v23
                            case "\n\t":
                                v1 = v20
                            case "hello":
                                v24 := GenS1{F0: 2, F1: v1}
                                _ = v24
                            default:
                                var v25 bool = (int16(v7) < 0)
                                _ = v25
                            }
                        } else {
                            v7 -= 1
                        }
                        if false {
                            if false {
                                if (false && (true || (v4 == (v4 >> ((18446744073709551615) & 63))))) {
                                    fmt.Println(v8, v7)
                                    if (false || (0 != int8(v7))) {
                                        v5[464] = ("hello" <= ("hello" + ""))
                                    }
                                }
                                var v26 int16 = 18
                                v27 := GenS2{F0: -89}
                                _ = v26
                                _ = v27
                            }
                        }
                        v28 := GenS4{F0: -326}
                        _ = fi22
                        _ = v28
                    }
                case 42:
                    fmt.Println(v9, v9, v4)
                }
                _ = fi21
            }
            _ = v20
        }
    default:
        if (1.5 > 1e10) {
            var v29 uint16 = (v8 | v8)
            _ = v29
        } else {
            var v30 uint64 = v1
            var v31 uint64 = v30
            _ = v30
            _ = v31
        }
    }
    switch (("" + "") + "a") {
    case "\n\t":
        var v32 uint16 = (v8 / (((v8 * 42)) | 1))
        _ = v32
    default:
        if false {
            if (false || true) {
                v33 := []int{}
                _ = v33
            }
            var v34 uint64 = (v1 - (v1 % ((v1) | 1)))
            _ = v34
        } else {
            v35 := []int{-2147483648}
            v36 := GenS5{F0: 0}
            var v37 int32 = v7
            _ = v35
            _ = v36
            _ = v37
        }
    }
    v1 = 0
    var v38 int64 = v4
    v7 = -485
    for fi39 := 0; fi39 < 19; fi39++ {
        v40 := GenS6{F0: 2}
        var v41 int16 = 806
        _ = fi39
        _ = v40
        _ = v41
    }
    if (!((1.5 <= 1e10) || (!false))) {
        delete(v5, 308)
        v1 = 23
    }
    v8 -= v8
    var v42 int8 = -6
    if (42 == 74) {
        fmt.Println(v1)
        if false {
            var v43 bool = true
            fmt.Println(v7, v43, v8)
            _ = v43
        }
        var v44 int16 = -1
        _ = v44
    }
    v45 := GenS5{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v45
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
