package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS5 struct { F0 bool; F1 int64 }
func genFn0(p0 uint, p1 uint64) int64 {
    var v7 int8 = 81
    if (18446744073709551615 > (p0 % (((p0 << (((p0 % ((p0) | 1))) & 63))) | 1))) {
        p1 = 1
    }
    var v8 float64 = 1.5
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int) int64 {
    var v10 int16 = 32767
    if true {
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := []int{2147483647}
            var v13 bool = false
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint64 = 0
        v16 := []int{2147483647, (fi14 % ((591) | 1)), fi14, p0}
        var v17 string = ("hello" + "hello")
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 float64 = float64((float32(float64(1e10)) / 1e10))
    v19 := GenS1{F0: 127, F1: -361, F2: 240}
    _ = p0
    _ = v10
    _ = v18
    _ = v19
    return -56
}

func main() {
    var v1 int = 42
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    v4 := []string{"hello", ("hello" + ""), ("a" + "hello"), "foo bar"}
    v5 := make(map[string]string)
    v6 := make(map[int]int)
    if (!(!(!(false || (!(1 != 65535)))))) {
        v1 = v1
        if (false && true) {
            v1 = v1
            var v20 bool = (42 >= 538)
            v1 = v1
            _ = v20
        } else {
            if (!true) {
                switch v1 {
                case -298:
                    fmt.Println(v1, v2)
                default:
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 uint8 = 26
                        if (((65535 >= 0) || true) && true) {
                            switch "" {
                            case "foo bar":
                                v4 = append(v4, "a")
                            case "\n\t":
                                var v23 int8 = 42
                                _ = v23
                            case "":
                                v24 := v6[fi21]
                                _ = v24
                            }
                            v4 = append(v4, "foo bar")
                        } else {
                            var v25 uint8 = (^(v22))
                            v4 = append(v4, ("hello" + "hello"))
                            _ = v25
                        }
                        v1 = v1
                        _ = fi21
                        _ = v22
                    }
                }
                v2 -= 32767
            }
        }
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        switch "hello" {
        case "a":
            var v27 uint8 = 128
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 uint16 = 753
                _ = len(v4)
                _ = fi28
                _ = v29
            }
        default:
            fmt.Println(fi26)
        }
        _ = fi26
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            if true {
                v32 := GenS2{F0: 42, F1: 363}
                _ = v32
            } else {
                v33 := GenS3{F0: 65535}
                var v34 int16 = v2
                _ = len(v4)
                v2 = v2
                v4 = append(v4, "\n\t")
                _ = v33
                _ = v34
            }
            switch "\n\t" {
            case "hello":
                var v35 uint8 = 2
                _ = v35
            case "a":
                var v36 string = ("hello" + "foo bar")
                _ = v36
            case "":
                switch v1 {
                case 4294967295:
                    var v37 uint16 = 2
                    _ = v37
                case 9223372036854775807:
                    _ = len(v4)
                case 32768:
                    var v38 uint64 = 470
                    _ = v38
                }
            default:
                var v39 int32 = 1
                _ = v39
            }
            _ = fi31
        }
        var v40 int16 = v2
        _ = fi30
        _ = v40
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        switch (v1 << ((2) & 63)) {
        case 855:
            v42 := GenS4{F0: 255, F1: 192, F2: 0}
            _ = v42
        default:
            fmt.Println(v1)
        }
        v4 = append(v4, "\n\t")
        for fi43 := 0; fi43 < 5; fi43++ {
            switch int(v2) {
            case 32768:
                v1 -= (v1 ^ (v1 % ((-9223372036854775808) | 1)))
            }
            v44 := GenS2{F0: 255, F1: 657}
            var v45 bool = (944 > 1)
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi41
    }
    v1 -= 4294967295
    v2 += (v2 & v2)
    switch "foo bar" {
    case "\n\t":
        v2 = v2
    default:
        fmt.Println(v1)
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        v47 := GenS2{F0: 1, F1: 1}
        v48 := GenS3{F0: 42}
        _ = fi46
        _ = v47
        _ = v48
    }
    if true {
        for fi49 := 0; fi49 < 19; fi49++ {
            if (!false) {
                var v50 uint32 = 4294967295
                v6[(v1 ^ fi49)] = int(v2)
                _ = v50
            } else {
                v51 := GenS5{F0: (114 > 1), F1: 377}
                _ = v51
            }
            for fi52 := 0; fi52 < 10; fi52++ {
                switch fi49 {
                case 42:
                    var v53 uint = 255
                    _ = v53
                }
                _ = fi52
            }
            v54 := GenS5{F0: (511 == -169), F1: -519}
            _ = fi49
            _ = v54
        }
    } else {
        v55 := GenS1{F0: 42, F1: 9223372036854775807, F2: (-(v2))}
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
