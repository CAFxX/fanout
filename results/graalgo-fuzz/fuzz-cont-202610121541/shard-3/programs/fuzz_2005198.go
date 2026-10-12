package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 float32 = float32(float64(3.14159))
    var v2 int16 = 32767
    var v3 uint32 = 0
    var v4 float64 = 0.0
    v5 := make(map[uint32]string)
    var v6 int8 = 127
    var v7 uint64 = 0
    for fi8 := 0; fi8 < 6; fi8++ {
        switch fi8 {
        case 32767:
            if (1 > -1) {
                v9 := GenS1{F0: -198}
                _ = v9
            } else {
                v10 := GenS1{F0: -1}
                _ = v10
            }
        case -1:
            v1 += v1
        case -2147483648:
            v11 := GenS1{F0: 2147483647}
            _ = v11
        default:
            fmt.Println(fi8)
        }
        v1 = v1
        var v12 float32 = v1
        _ = fi8
        _ = v12
    }
    if true {
        v13 := GenS1{F0: int32(v6)}
        for fi14 := 0; fi14 < 19; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int64 = 42
                if (!((3.14159 - -2.25) <= (((v4 - 3.14159) - v4) + v4))) {
                    var v17 uint16 = 65535
                    fmt.Println(v3)
                    _ = v17
                } else {
                    var v18 uint8 = 45
                    var v19 int8 = v6
                    _ = v18
                    _ = v19
                }
                var v20 uint8 = 127
                var v21 float64 = 1.5
                _ = fi15
                _ = v16
                _ = v20
                _ = v21
            }
            _ = fi14
        }
        var v22 int16 = int16(v7)
        _ = v13
        _ = v22
    }
    v23 := GenS2{F0: (-(v6)), F1: (v7 | 782), F2: (-1 <= 0)}
    v24 := GenS1{F0: 0}
    v25 := []int{0}
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(fi26, v2, v6)
        switch ("foo bar" + "a") {
        case "hello":
            v2 = (v2 + (v2 * v2))
        case "":
            switch -1 {
            case 1:
                var v27 int64 = -1
                _ = v27
            case -603:
                fmt.Println(v7, v6)
            case 42:
                var v28 int32 = 1
                _ = v28
            }
        case "foo bar":
            var v29 uint64 = v7
            _ = v29
        }
        v25 = append(v25, int(v2))
        _ = fi26
    }
    v4 -= float64(0.0)
    if true {
        v30 := v5[(v3 << ((1) & 63))]
        _ = v30
    } else {
        if false {
            if true {
                for fi31 := 0; fi31 < 5; fi31++ {
                    switch "foo bar" {
                    case "":
                        v5[(v3 + v3)] = "a"
                    }
                    v7 -= v7
                    v2 += (v2 * v2)
                    _ = fi31
                }
                fmt.Println(v2, v7)
            }
            if ((v3 / ((4294967295) | 1)) > v3) {
                v32 := GenS2{F0: v6, F1: 1, F2: false}
                if (42 > 127) {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        v4 -= v4
                        v5[v3] = ("hello" + "a")
                        _ = fi33
                    }
                    fmt.Println(v4, v6)
                } else {
                    switch ("hello" + "\n\t") {
                    case "a":
                        v2 = 76
                    default:
                        v2 = 498
                    }
                    for fi34 := 0; fi34 < 14; fi34++ {
                        for fi35 := 0; fi35 < 9; fi35++ {
                            fmt.Println(fi35, fi35, v4)
                            switch "\n\t" {
                            case "foo bar":
                                var v36 float64 = float64(v1)
                                _ = v36
                            case "a":
                                v37 := GenS2{F0: (v6 % (((v6 % (((v6 + int8(v2))) | 1))) | 1)), F1: 927, F2: true}
                                v6 = (-(v6))
                                v6 -= (v6 | v6)
                                if false {
                                    v38 := GenS1{F0: -977}
                                    v2 += v2
                                    var v39 uint8 = 127
                                    _ = v38
                                    _ = v39
                                } else {
                                    v2 = 1
                                    v4 = v4
                                    _ = len(v5)
                                    fmt.Println(v6)
                                    fmt.Println(v7)
                                }
                                _ = v37
                            }
                            _ = fi35
                        }
                        _ = fi34
                    }
                    v6 = 7
                }
                _ = v32
            }
            if (v1 != 1.5) {
                switch ("foo bar" + "") {
                case "hello":
                    var v40 int = -668
                    _ = v40
                case "a":
                    delete(v5, v3)
                }
                for fi41 := 0; fi41 < 12; fi41++ {
                    if (759 >= 0) {
                        v42 := GenS1{F0: -770}
                        _ = v42
                    }
                    var v43 uint = 438
                    _ = fi41
                    _ = v43
                }
            } else {
                var v44 int16 = (v2 & (v2 + v2))
                _ = v44
            }
        } else {
            delete(v5, v3)
            if (v4 <= v4) {
                var v45 bool = true
                if true {
                    var v46 uint = 2
                    v6 = v6
                    _ = v46
                }
                _ = len(v25)
                _ = v45
            }
            if (-169 > 2147483647) {
                v1 = (v1 + (v1 / v1))
            }
        }
    }
    switch 32767 {
    case 9223372036854775807:
        var v47 bool = (!(v4 == (v4 + 0.1)))
        _ = v47
    case -710:
        var v48 int = 128
        _ = v48
    case -2147483648:
        var v49 uint16 = 255
        _ = v49
    }
    var v50 float32 = float32(v4)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
