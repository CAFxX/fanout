package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint; F1 bool }

func main() {
    var v1 float32 = 1e10
    var v2 uint = 62
    var v3 uint = (v2 | (^(v2)))
    var v4 uint = v3
    var v5 int = 32767
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(v3)
        _ = fi6
    }
    var v7 float64 = 0.0
    if (503 >= 42) {
        if (-1 < -21) {
            if true {
                v8 := GenS1{F0: v3, F1: (^(v3)), F2: 42}
                fmt.Println(v4)
                v2 = v2
                _ = v8
            }
            v9 := []int{0, (v5 << (((^(v4))) & 63)), 32768, (^(v5))}
            _ = v9
        }
        v2 = 92
    }
    v4 = (v3 >> ((34) & 63))
    if (!(42 >= -1)) {
        v3 += v2
        switch (v5 >> (((v3 * 255)) & 63)) {
        case 32767:
            if true {
                v10 := GenS1{F0: v2, F1: 0, F2: 212}
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        if false {
                            var v13 uint = (^(v2))
                            v14 := GenS1{F0: v4, F1: (v13 & v4), F2: 95}
                            var v15 int32 = 287
                            var v16 int = v5
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v17 := []int{v5, (-(fi12)), int(v2), (v5 + 0)}
                            if false {
                                v2 += (v4 ^ v4)
                            } else {
                                if (42 < 1) {
                                    var v18 uint = (v3 << ((4294967295) & 63))
                                    v5 += (fi12 ^ (-(fi12)))
                                    fmt.Println(v7, v3)
                                    _ = v18
                                }
                            }
                            var v19 int32 = 1
                            _ = v17
                            _ = v19
                        }
                        v20 := []int{128, fi12, v5}
                        _ = fi12
                        _ = v20
                    }
                    v21 := GenS1{F0: (v2 * (^(v4))), F1: 375, F2: 2}
                    _ = fi11
                    _ = v21
                }
                if false {
                    switch (v5 | (-(v5))) {
                    case 128:
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v1 += (v1 * v1)
                            v23 := GenS1{F0: 18446744073709551615, F1: v4, F2: 154}
                            _ = fi22
                            _ = v23
                        }
                    case -243:
                        if false {
                            v24 := []int{v5}
                            var v25 int = (v5 ^ v5)
                            _ = v24
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 20; fi26++ {
                                v27 := GenS2{F0: 0, F1: (69 > 100), F2: ("hello" + "foo bar")}
                                _ = fi26
                                _ = v27
                            }
                            v28 := []int{v5, v5, (v5 - v5), (v5 / (((v5 % (((^(v5))) | 1))) | 1))}
                            var v29 int8 = 42
                            _ = v28
                            _ = v29
                        }
                    default:
                        var v30 uint64 = 42
                        _ = v30
                    }
                } else {
                    v31 := []int{}
                    v7 = float64(1.5)
                    v32 := GenS3{F0: (v3 << ((2) & 63)), F1: 1}
                    _ = v31
                    _ = v32
                }
                _ = v10
            } else {
                for fi33 := 0; fi33 < 14; fi33++ {
                    v1 += (((v1 + -2.25) - (0.1 + v1)) - (v1 - float32(v7)))
                    _ = fi33
                }
                v34 := []int{v5, v5, (v5 & v5), 331}
                _ = v34
            }
        case 127:
            fmt.Println(v2, v4)
        case -1:
            var v35 uint64 = 8
            _ = v35
        default:
            v4 += (v2 >> ((v3) & 63))
        }
    }
    if (127 != 0) {
        switch (v5 | v5) {
        case -2147483648:
            v4 = (v4 & v3)
        case 32768:
            for fi36 := 0; fi36 < 1; fi36++ {
                v37 := []int{}
                v37 = append(v37, 4294967295)
                if true {
                    switch (fi36 % (((fi36 + 0)) | 1)) {
                    case -2147483648:
                        switch ("" + ("" + "a")) {
                        case "hello":
                            v7 += v7
                        case "foo bar":
                            var v38 float64 = (((v7 - v7) - ((1.5 * v7) + v7)) * 3.14159)
                            _ = v38
                        case "":
                            v39 := GenS4{F0: "foo bar", F1: -1}
                            _ = v39
                        }
                    case 127:
                        switch (-(v5)) {
                        case 128:
                            _ = len(v37)
                        }
                    case 4294967295:
                        _ = len(v37)
                    }
                }
                _ = fi36
                _ = v37
            }
        default:
            var v40 int64 = 361
            _ = v40
        }
        v41 := GenS4{F0: "\n\t", F1: 1}
        if false {
            v4 = v2
        } else {
            v4 = (v3 / (((v4 & 42)) | 1))
        }
        _ = v41
    }
    if (65535 >= 189) {
        var v42 bool = (4294967295 >= 4294967295)
        switch "a" {
        case "":
            for fi43 := 0; fi43 < 10; fi43++ {
                var v44 int8 = -128
                _ = fi43
                _ = v44
            }
        case "foo bar":
            for fi45 := 0; fi45 < 12; fi45++ {
                var v46 float32 = v1
                _ = fi45
                _ = v46
            }
        default:
            v47 := []int{}
            _ = v47
        }
        for fi48 := 0; fi48 < 10; fi48++ {
            for fi49 := 0; fi49 < 6; fi49++ {
                fmt.Println(v5, v5, v42)
                v50 := GenS3{F0: v2, F1: 851}
                _ = fi49
                _ = v50
            }
            v1 += -2.25
            _ = fi48
        }
        _ = v42
    } else {
        v51 := GenS5{F0: v4, F1: true}
        var v52 uint8 = 128
        _ = v51
        _ = v52
    }
    v53 := GenS1{F0: (v2 & v4), F1: v3, F2: 2}
    fmt.Println(v5, v1)
    v54 := GenS2{F0: 255, F1: ((953 > 1) || (v1 == (v1 - v1))), F2: ("hello" + "a")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v7
}
