package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 string; F1 int32; F2 int64 }

func main() {
    var v1 uint64 = 799
    v2 := []int{-1, -744, -2147483648, -1}
    var v3 int16 = -1
    var v4 int16 = v3
    var v5 uint8 = 127
    var v6 string = "hello"
    var v7 int32 = 2147483647
    var v8 int32 = (v7 * -702)
    switch ("foo bar" + "foo bar") {
    case "a":
        if true {
            fmt.Println(v6, v1, v3)
            var v9 uint8 = v5
            _ = v9
        } else {
            var v10 uint64 = 2
            _ = v10
        }
    case "foo bar":
        var v11 int32 = 42
        _ = v11
    }
    v5 = v5
    v4 = int16(v5)
    var v12 int8 = -1
    for fi13 := 0; fi13 < 11; fi13++ {
        if false {
            v7 = v8
            var v14 uint32 = 2147483648
            _ = v14
        } else {
            var v15 int = (fi13 % ((fi13) | 1))
            _ = v15
        }
        _ = fi13
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            if (true && (82 == 147)) {
                var v16 uint64 = (^(v1))
                if len(v2) > 0 {
                    v2[0] = -772
                }
                v2 = append(v2, 0)
                _ = v16
            } else {
                if (4294967295 <= 1) {
                    var v17 uint32 = 282
                    for fi18 := 0; fi18 < 8; fi18++ {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            if (-448 != v3) {
                                switch "foo bar" {
                                case "":
                                    v8 -= v7
                                case "hello":
                                    if (0 > v4) {
                                        v20 := GenS1{F0: 1, F1: int16(v7), F2: (v3 >> ((1) & 63))}
                                        var v21 int = (fi19 & fi19)
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if (!false) {
                                            v22 := GenS1{F0: 18446744073709551615, F1: 1000, F2: 44}
                                            _ = v22
                                        } else {
                                            var v23 int32 = v7
                                            if true {
                                                var v24 uint = 846
                                                v25 := GenS2{F0: v23, F1: (^(v17)), F2: 2}
                                                _ = v24
                                                _ = v25
                                            } else {
                                                var v26 float32 = -2.25
                                                var v27 bool = true
                                                _ = v26
                                                _ = v27
                                            }
                                            var v28 int64 = -388
                                            _ = v23
                                            _ = v28
                                        }
                                    }
                                }
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            switch 32768 {
                            case -1:
                                if false {
                                    fmt.Println(fi18, fi19)
                                    v17 += (^(v17))
                                    if len(v2) > 0 {
                                        v2[0] = int(v1)
                                    }
                                } else {
                                    var v29 int64 = 42
                                    _ = v29
                                }
                            case 42:
                                v30 := GenS2{F0: -929, F1: v17, F2: 450}
                                _ = v30
                            }
                            v3 -= (v3 - (v3 / ((v3) | 1)))
                            v4 = (^(v4))
                            _ = len(v2)
                            _ = fi19
                        }
                        var v31 int16 = (v3 & v4)
                        switch "a" {
                        case "\n\t":
                            switch int(v12) {
                            case 128:
                                v12 = v12
                            case -1:
                                switch "a" {
                                case "\n\t":
                                    v7 += (^(v7))
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = fi18
                                    }
                                case "a":
                                    var v32 uint8 = v5
                                    _ = v32
                                default:
                                    var v33 int64 = -1
                                    _ = v33
                                }
                            case 0:
                                if ((v17 - (v17 * v17)) > (v17 ^ (v17 & 0))) {
                                    v34 := GenS3{F0: 4294967295}
                                    if len(v2) > 0 {
                                        v2[0] = fi18
                                    }
                                    _ = v34
                                }
                            }
                        case "a":
                            v35 := GenS2{F0: v7, F1: v17, F2: v1}
                            _ = v35
                        }
                        _ = fi18
                        _ = v31
                    }
                    var v36 uint16 = uint16(v17)
                    _ = v17
                    _ = v36
                }
                fmt.Println(v3, v3)
                v37 := GenS3{F0: 0}
                _ = v37
            }
            v38 := GenS4{F0: 255, F1: 4294967295}
            v39 := GenS2{F0: (v7 << ((2) & 63)), F1: 0, F2: 18446744073709551615}
            _ = v38
            _ = v39
        } else {
            var v40 uint64 = (v1 << ((18446744073709551615) & 63))
            _ = v40
        }
    case "a":
        var v41 int8 = (v12 ^ -128)
        _ = v41
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 int32 = (v7 & v7)
        if false {
            for fi44 := 0; fi44 < 3; fi44++ {
                var v45 int8 = v12
                if true {
                    var v46 string = ""
                    v3 -= int16(v12)
                    fmt.Println(v43, v6, v7)
                    _ = v46
                } else {
                    v47 := GenS5{F0: (v1 << ((1) & 63)), F1: v1, F2: (v1 * v1)}
                    _ = v47
                }
                v48 := GenS6{F0: v6, F1: -2147483648, F2: 9223372036854775807}
                _ = fi44
                _ = v45
                _ = v48
            }
            v2 = append(v2, fi42)
            if (v6 == ((v6 + "a") + "")) {
                for fi49 := 0; fi49 < 17; fi49++ {
                    var v50 uint8 = v5
                    _ = fi49
                    _ = v50
                }
            } else {
                var v51 int64 = -9223372036854775808
                _ = v51
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi42
        _ = v43
    }
    v4 = 366
    for fi52 := 0; fi52 < 9; fi52++ {
        var v53 uint64 = v1
        _ = fi52
        _ = v53
    }
    var v54 float64 = 1.5
    v7 += v8
    if len(v2) > 0 {
        v2[0] = 568
    }
    v5 = (^(v5))
    if len(v2) > 0 {
        v2[0] = int(v12)
    }
    for fi55 := 0; fi55 < 4; fi55++ {
        v6 = ("a" + v6)
        _ = len(v2)
        fmt.Println(v5)
        _ = fi55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v7
    _ = v8
}
