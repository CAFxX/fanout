package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 uint; F1 int64; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 bool }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 bool; F1 int; F2 int32 }
type GenS8 struct { F0 uint32; F1 uint8; F2 int8 }

func main() {
    var v1 int32 = -2147483648
    var v2 int = 262
    var v3 string = "foo bar"
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: 2147483648, F1: -9223372036854775808, F2: 0}
    if ((0.1 * 0.1) >= 0.0) {
        switch v3 {
        case "\n\t":
            v6 := []int{v2, (v2 - v2)}
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: 1, F1: -729, F2: 929}
                _ = fi7
                _ = v8
            }
        case "":
            v1 += (v1 * (v1 + -1))
        default:
            var v9 uint8 = 2
            _ = v9
        }
        var v10 string = v3
        _ = v10
    }
    if (1 <= 339) {
        fmt.Println(v2)
        v11 := GenS2{F0: 255, F1: 0, F2: (v2 & v2)}
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        switch (-(fi12)) {
        case 954:
            v13 := GenS3{F0: 255}
            _ = v13
        case 42:
            v2 -= fi12
        case 0:
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := []int{fi14, (-(v2))}
                var v16 uint8 = 127
                v2 = fi14
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            var v17 int16 = 1
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            if (v3 < "\n\t") {
                if true {
                    fmt.Println(v2)
                    var v19 string = "\n\t"
                    v1 -= v1
                    var v20 float64 = 3.14159
                    var v21 string = v3
                    var v22 int = 941
                    v1 = -943
                    v1 = v1
                    v20 += 1e10
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    var v23 uint = 42
                    if (fi12 >= (v2 ^ (-(fi18)))) {
                        v24 := GenS4{F0: false, F1: 1}
                        if false {
                            var v25 int32 = v1
                            _ = v25
                        }
                        _ = v24
                    }
                    _ = v23
                }
            }
            _ = fi18
        }
        _ = fi12
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v2 += fi26
        var v27 float32 = (((0.1 - 3.14159) - (1e10 / 3.14159)) - 0.0)
        if (true && false) {
            if (4294967295 > 2147483647) {
                for fi28 := 0; fi28 < 10; fi28++ {
                    fmt.Println(v3)
                    v29 := []int{}
                    if false {
                        if ((-(v1)) >= (v1 ^ (v1 ^ (v1 & -673)))) {
                            fmt.Println(fi28)
                        }
                        _ = len(v29)
                        var v30 uint8 = 255
                        _ = v30
                    }
                    _ = fi28
                    _ = v29
                }
            } else {
                fmt.Println(v27)
            }
        }
        _ = fi26
        _ = v27
    }
    var v31 uint8 = 255
    switch 42 {
    case 42:
        var v32 uint8 = (v31 % ((2) | 1))
        _ = v32
    case 127:
        v33 := GenS5{F0: (v31 & v31), F1: 2, F2: (((false || (958 <= 0)) && (2 > 1)) || true)}
        _ = v33
    case -23:
        var v34 uint = 0
        _ = v34
    }
    if false {
        v1 = (-(v1))
        switch v3 {
        case "foo bar":
            v35 := []int{0, 2147483647, (^(v2))}
            _ = v35
        default:
            v36 := GenS5{F0: (v31 + 128), F1: 255, F2: false}
            _ = v36
        }
    } else {
        switch (v3 + v3) {
        case "\n\t":
            var v37 uint = 65535
            _ = v37
        case "foo bar":
            v38 := GenS4{F0: (v2 != (v2 | v2)), F1: 0}
            _ = v38
        case "a":
            for fi39 := 0; fi39 < 16; fi39++ {
                fmt.Println(v2, v31, fi39)
                _ = fi39
            }
        }
    }
    if true {
        v31 -= (^(v31))
        for fi40 := 0; fi40 < 11; fi40++ {
            if true {
                if true {
                    v1 += (v1 << ((0) & 63))
                }
                if false {
                    var v41 int8 = 36
                    _ = v41
                } else {
                    fmt.Println(v31, v1, v31)
                }
                var v42 int16 = 1
                _ = v42
            } else {
                var v43 int16 = 0
                _ = v43
            }
            var v44 float64 = 0.1
            v45 := GenS6{F0: 2147483648, F1: 0}
            _ = fi40
            _ = v44
            _ = v45
        }
    }
    switch v3 {
    case "hello":
        v46 := []int{(^(v2)), -296, v2}
        _ = v46
    case "foo bar":
        var v47 uint16 = 65535
        _ = v47
    case "\n\t":
        v2 += 708
    }
    if false {
        var v48 uint64 = 42
        _ = v48
    }
    switch -134 {
    case 4294967295:
        for fi49 := 0; fi49 < 12; fi49++ {
            v50 := []int{(v2 << ((2) & 63)), 127, -394}
            switch v3 {
            case "foo bar":
                var v51 float64 = (3.14159 / 1.5)
                _ = v51
            default:
                for fi52 := 0; fi52 < 9; fi52++ {
                    if (!(float64(0.0) >= 1e10)) {
                        if len(v50) >= 2 {
                            v50 = v50[1:len(v50)]
                        }
                        _ = len(v4)
                        var v53 int = v2
                        _ = v53
                    }
                    if true {
                        var v54 uint8 = v31
                        _ = v54
                    } else {
                        v3 = v3
                        switch v3 {
                        case "a":
                            if (2 == 2147483647) {
                                if (!((v31 % ((165) | 1)) == (v31 ^ 1))) {
                                    v31 = v31
                                    if false {
                                        var v55 uint16 = 1
                                        v1 = v1
                                        _ = v55
                                    }
                                }
                                v56 := GenS7{F0: (false && (42 <= -838)), F1: (-(fi49)), F2: 324}
                                v50 = append(v50, (v2 / (((v2 & fi49)) | 1)))
                                _ = v56
                            }
                        case "foo bar":
                            _ = len(v4)
                        }
                    }
                    _ = fi52
                }
            }
            v1 += (^(v1))
            _ = fi49
            _ = v50
        }
    case -9223372036854775808:
        fmt.Println(v2, v1)
    case 128:
        v57 := []int{-1, -51}
        _ = v57
    default:
        if (2147483647 >= -2147483648) {
            fmt.Println(v1, v1, v1)
        }
    }
    var v58 uint8 = v31
    v59 := GenS8{F0: 729, F1: v58, F2: 3}
    v60 := []int{9223372036854775807, v2}
    v31 -= (v31 + v58)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
}
