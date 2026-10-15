package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint64 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint = 786
        switch (fi7 << ((2) & 63)) {
        case 4294967295:
            var v9 int64 = 1
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if false {
        if false {
            var v10 int64 = 0
            for fi11 := 0; fi11 < 18; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 float64 = (((3.14159 - (1.5 * 1e10)) * (0.0 * ((3.14159 / -2.25) / 0.1))) / ((-2.25 * 1.5) * -2.25))
                    switch fi12 {
                    case -430:
                        v14 := []int{(fi12 & 2147483647), (fi12 & (fi12 << ((4294967295) & 63))), fi12, fi12}
                        _ = v14
                    case 42:
                        v15 := []int{}
                        _ = v15
                    default:
                        v10 = (v10 | (v10 - (v10 & 0)))
                    }
                    v10 += (v10 << ((18446744073709551615) & 63))
                    _ = fi12
                    _ = v13
                }
                v10 = v10
                for fi16 := 0; fi16 < 6; fi16++ {
                    v10 = v10
                    fmt.Println(v10)
                    v17 := []int{(fi16 | fi16), (fi16 * fi16), 32768, -9223372036854775808}
                    _ = fi16
                    _ = v17
                }
                _ = fi11
            }
            _ = v10
        } else {
            var v18 float64 = -2.25
            v19 := GenS1{F0: 255, F1: 255, F2: 524}
            _ = v18
            _ = v19
        }
        var v20 int32 = 1
        v20 += v20
        _ = v20
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 bool = false
            v23 := GenS1{F0: 2, F1: 65535, F2: 65535}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    if false {
        if (true && true) {
            if ((-876 > -435) && ((65535 > 42) && true)) {
                switch "hello" {
                case "":
                    switch 42 {
                    case -1:
                        for fi24 := 0; fi24 < 3; fi24++ {
                            v25 := []int{32768, (^(fi24)), 323}
                            _ = len(v25)
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            _ = fi24
                            _ = v25
                        }
                    case -9223372036854775808:
                        var v26 uint8 = 42
                        _ = v26
                    default:
                        v27 := GenS2{F0: "foo bar", F1: 127, F2: 9223372036854775807}
                        _ = v27
                    }
                case "foo bar":
                    var v28 float32 = 0.0
                    _ = v28
                case "a":
                    v29 := GenS3{F0: 894, F1: -47}
                    _ = v29
                }
            }
            switch "foo bar" {
            case "foo bar":
                v30 := GenS4{F0: 42}
                _ = v30
            case "":
                v31 := GenS4{F0: 291}
                _ = v31
            default:
                var v32 uint64 = 2
                _ = v32
            }
            var v33 uint16 = 0
            _ = v33
        }
    }
    v34 := []int{-2147483648, -1, 9223372036854775807, 3}
    _ = v34
    return 248
}
func genFn1() int {
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 uint = 255
        fmt.Println(v36, fi35)
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 bool = true
            v39 := []int{(fi37 + 42)}
            v36 = v36
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    var v40 int64 = 388
    _ = v40
    return 954
}

func main() {
    v1 := []uint32{250, 42, 79, 4294967295}
    var v2 float64 = 3.14159
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    var v5 uint8 = 210
    var v6 uint32 = 4294967295
    if (true || true) {
        v5 += v5
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 int = genFn1()
        _ = fi41
        _ = v42
    }
    fmt.Println(v2, v5)
    switch "foo bar" {
    case "\n\t":
        for fi43 := 0; fi43 < 10; fi43++ {
            fmt.Println(v5, v2)
            _ = len(v4)
            _ = fi43
        }
    case "hello":
        for fi44 := 0; fi44 < 11; fi44++ {
            _ = len(v3)
            _ = fi44
        }
    }
    v6 -= 2
    v45 := GenS5{F0: 2, F1: 42}
    if ((v6 / (((v6 / ((v6) | 1))) | 1)) != 42) {
        fmt.Println(v2, v6, v6)
        for fi46 := 0; fi46 < 12; fi46++ {
            v6 = (v6 | v6)
            _ = fi46
        }
        v5 = (v5 >> ((18446744073709551615) & 63))
    } else {
        var v47 uint8 = v5
        v47 = (^(v47))
        for fi48 := 0; fi48 < 8; fi48++ {
            v6 = (v6 << ((944) & 63))
            fmt.Println(v5, fi48)
            if true {
                var v49 int64 = 295
                var v50 uint32 = (v6 % (((v6 ^ v6)) | 1))
                _ = v49
                _ = v50
            }
            _ = fi48
        }
        _ = v47
    }
    v51 := GenS6{F0: (^(v5)), F1: 955, F2: 1}
    var v52 uint64 = 1
    v53 := GenS6{F0: v5, F1: 553, F2: (^(v52))}
    fmt.Println(v52, v52, v52)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
