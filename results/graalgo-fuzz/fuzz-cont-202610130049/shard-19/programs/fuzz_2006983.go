package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64; F1 bool }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    var v3 float64 = 0.1
    var v4 uint64 = 2
    var v5 float64 = v3
    delete(v2, 128)
    v3 -= (v5 * v5)
    var v6 int64 = 42
    var v7 float64 = v5
    for fi8 := 0; fi8 < 3; fi8++ {
        if (!true) {
            var v9 int64 = -351
            if (!false) {
                var v10 float64 = float64(-2.25)
                _ = v10
            } else {
                switch (fi8 ^ -52) {
                case -403:
                    v2[(fi8 - (fi8 << ((336) & 63)))] = 2147483648
                default:
                    var v11 uint8 = 155
                    _ = v11
                }
                var v12 float64 = float64((-2.25 / 1.5))
                _ = v12
            }
            _ = v9
        } else {
            if true {
                v13 := []int{(^(fi8)), (fi8 % (((fi8 >> ((65535) & 63))) | 1)), (-(fi8)), (fi8 >> ((105) & 63))}
                _ = len(v2)
                fmt.Println(v7, fi8)
                _ = v13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 bool = true
                _ = len(v2)
                _ = fi14
                _ = v15
            }
        }
        var v16 int = (fi8 & (fi8 >> ((456) & 63)))
        _ = fi8
        _ = v16
    }
    if (!false) {
        v17 := GenS1{F0: 42, F1: -1}
        switch v1 {
        case "foo bar":
            switch -347 {
            case 462:
                switch -1 {
                case 1:
                    if true {
                        var v18 float32 = (3.14159 - (float32(v5) / (1e10 + float32(v7))))
                        v19 := []int{}
                        _ = v18
                        _ = v19
                    }
                case 4294967295:
                    if (!false) {
                        v2[0] = 2147483647
                        for fi20 := 0; fi20 < 13; fi20++ {
                            switch "" {
                            case "a":
                                if false {
                                    v3 = 0.0
                                    var v21 float64 = 1e10
                                    if true {
                                        delete(v2, fi20)
                                        v6 -= (v6 << ((2) & 63))
                                    } else {
                                        v22 := v2[(-(fi20))]
                                        v23 := GenS2{F0: "", F1: 2147483648}
                                        _ = v22
                                        _ = v23
                                    }
                                    _ = v21
                                } else {
                                    fmt.Println(v6, v6)
                                    fmt.Println(v6, v6)
                                }
                            case "":
                                v4 -= 42
                            }
                            v24 := GenS2{F0: v1, F1: 2}
                            _ = fi20
                            _ = v24
                        }
                        v1 = v1
                    }
                case -9223372036854775808:
                    v25 := []int{}
                    _ = v25
                default:
                    var v26 uint8 = 127
                    _ = v26
                }
            case 567:
                switch v1 {
                case "":
                    v5 = (v3 * -2.25)
                case "foo bar":
                    v27 := GenS1{F0: 180, F1: -774}
                    _ = v27
                case "a":
                    switch 410 {
                    case 32768:
                        for fi28 := 0; fi28 < 7; fi28++ {
                            v29 := GenS2{F0: v1, F1: 345}
                            _ = fi28
                            _ = v29
                        }
                    }
                }
            }
        case "\n\t":
            var v30 int = 42
            _ = v30
        default:
            v31 := GenS1{F0: 65535, F1: 219}
            _ = v31
        }
        _ = v17
    } else {
        for fi32 := 0; fi32 < 5; fi32++ {
            var v33 float32 = -2.25
            if (42 >= 2147483648) {
                for fi34 := 0; fi34 < 9; fi34++ {
                    var v35 int = fi34
                    switch int(v4) {
                    case -267:
                        if ((true || (fi34 == int(v4))) && true) {
                            v4 += 2
                            fmt.Println(fi32)
                        }
                        v6 = v6
                        fmt.Println(v35, v7, v35)
                    case -901:
                        var v36 int8 = -102
                        _ = v36
                    case -9223372036854775808:
                        if (!false) {
                            switch (v1 + v1) {
                            case "foo bar":
                                if true {
                                    if true {
                                        fmt.Println(v3, v33, v7)
                                    }
                                    var v37 uint64 = (v4 << ((18446744073709551615) & 63))
                                    _ = v37
                                } else {
                                    fmt.Println(fi32, v4, v4)
                                    v3 -= (float64(float32(1.5)) / (v7 * (v3 / float64(v33))))
                                    v4 -= v4
                                }
                            }
                        } else {
                            var v38 uint8 = 128
                            if true {
                                if false {
                                    var v39 int32 = 42
                                    fmt.Println(v39, v38, v35)
                                    _ = v39
                                }
                            } else {
                                if (1 >= 42) {
                                    delete(v2, (fi34 & v35))
                                    v2[int(v6)] = 1
                                }
                                fmt.Println(fi34, v33)
                                var v40 int8 = 100
                                _ = v40
                            }
                            v6 += (v6 >> ((0) & 63))
                            _ = v38
                        }
                    default:
                        var v41 float32 = 3.14159
                        _ = v41
                    }
                    var v42 float64 = ((v5 * -2.25) / v7)
                    _ = fi34
                    _ = v35
                    _ = v42
                }
                v43 := GenS3{F0: v4, F1: true}
                _ = v43
            }
            var v44 uint8 = 2
            _ = fi32
            _ = v33
            _ = v44
        }
        if false {
            v3 -= v5
            v4 = v4
        } else {
            fmt.Println(v7, v4)
            var v45 uint16 = 1
            _ = v45
        }
        v7 = (3.14159 / 1.5)
    }
    v4 = v4
    var v46 float32 = ((3.14159 / -2.25) - 1.5)
    for fi47 := 0; fi47 < 2; fi47++ {
        for fi48 := 0; fi48 < 20; fi48++ {
            v7 += (float64(v46) + v3)
            var v49 int8 = -96
            _ = fi48
            _ = v49
        }
        var v50 float32 = v46
        if ((fi47 ^ (fi47 ^ (^(fi47)))) > fi47) {
            var v51 uint16 = 65535
            v1 = ("hello" + ("a" + v1))
            switch (v1 + ("foo bar" + "a")) {
            case "a":
                var v52 string = "foo bar"
                _ = v52
            case "foo bar":
                var v53 uint = 4294967295
                _ = v53
            case "":
                delete(v2, int(v51))
            default:
                v54 := v2[(fi47 & (fi47 - fi47))]
                _ = v54
            }
            _ = v51
        }
        _ = fi47
        _ = v50
    }
    var v55 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v55
    _ = v6
    _ = v7
}
