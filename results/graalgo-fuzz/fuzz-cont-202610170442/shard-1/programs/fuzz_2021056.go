package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int; F2 uint8 }
type GenS6 struct { F0 uint; F1 uint }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]uint32)
    var v3 float64 = (3.14159 / 1.0)
    var v4 bool = (float64(-2.25) <= -2.25)
    var v5 int = 1
    v3 += v3
    for fi6 := 0; fi6 < 4; fi6++ {
        v4 = (255 > 255)
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 896}
            var v9 bool = false
            if v9 {
                fmt.Println(v3, v1)
            } else {
                var v10 int16 = 0
                var v11 uint64 = 687
                v12 := GenS2{F0: 2147483647}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := []int{fi6, (fi13 >> ((0) & 63))}
            var v15 int64 = v1
            _ = len(v2)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi6
    }
    if true {
        v16 := GenS3{F0: -604, F1: 65535}
        var v17 float64 = v3
        if (879 > -1) {
            for fi18 := 0; fi18 < 20; fi18++ {
                for fi19 := 0; fi19 < 16; fi19++ {
                    v5 = 1
                    fmt.Println(v17, fi18)
                    v20 := GenS2{F0: 2147483647}
                    var v21 int64 = v1
                    v22 := []int{(-(fi18)), fi18}
                    v23 := GenS1{F0: 176}
                    v21 -= 925
                    v21 += (v1 & v1)
                    switch (v5 * v5) {
                    case 264:
                        var v24 int64 = v1
                        _ = v24
                    default:
                        v25 := GenS4{F0: 42, F1: 2}
                        _ = v25
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                if true {
                    var v26 int32 = -1
                    for fi27 := 0; fi27 < 5; fi27++ {
                        switch ("hello" + ("a" + "")) {
                        case "a":
                            v5 -= (fi18 ^ 9223372036854775807)
                        default:
                            v28 := []int{(v5 | fi18), 0, fi18, -379}
                            _ = v28
                        }
                        _ = fi27
                    }
                    _ = v26
                } else {
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        v29 := []int{fi18, (v5 - fi18)}
                        _ = v29
                    case "a":
                        v5 = (fi18 | (-(fi18)))
                    default:
                        if (4294967295 < 65535) {
                            v30 := []int{(fi18 + 9223372036854775807)}
                            fmt.Println(v3, v1, v3)
                            _ = v30
                        }
                    }
                    if (2 >= 1) {
                        switch "" {
                        case "a":
                            v31 := GenS4{F0: 0, F1: 0}
                            _ = v31
                        case "foo bar":
                            v17 = (float64(0.0) / (v17 * v3))
                        default:
                            v4 = false
                        }
                        v5 -= (v5 >> ((1) & 63))
                        if v4 {
                            var v32 uint = 65535
                            _ = v32
                        } else {
                            v5 = (fi18 >> ((65535) & 63))
                            var v33 bool = v4
                            var v34 float64 = v3
                            _ = v33
                            _ = v34
                        }
                    } else {
                        for fi35 := 0; fi35 < 12; fi35++ {
                            if v4 {
                                switch 826 {
                                case 2147483647:
                                    switch (-(v5)) {
                                    case -164:
                                        var v36 int = 128
                                        _ = v36
                                    }
                                case -994:
                                    switch 1 {
                                    case -9223372036854775808:
                                        _ = len(v2)
                                    case 193:
                                        var v37 int16 = -32768
                                        _ = v37
                                    }
                                    fmt.Println(v3, v4)
                                    var v38 string = ("a" + (("" + ("hello" + "hello")) + "\n\t"))
                                    v38 = v38
                                    _ = v38
                                }
                            }
                            _ = fi35
                        }
                    }
                    if false {
                        for fi39 := 0; fi39 < 15; fi39++ {
                            fmt.Println(v5)
                            _ = len(v2)
                            var v40 float32 = float32(v3)
                            _ = fi39
                            _ = v40
                        }
                    } else {
                        v41 := []int{(fi18 ^ 32767), -805, 32767}
                        _ = v41
                    }
                }
                v3 -= ((v17 + v17) - v17)
                _ = fi18
            }
            v42 := []int{-9223372036854775808, v5}
            var v43 uint8 = 0
            _ = v42
            _ = v43
        }
        _ = v16
        _ = v17
    } else {
        if (v1 >= (-(v1))) {
            var v44 int8 = -21
            v45 := GenS4{F0: 40, F1: 786}
            _ = v44
            _ = v45
        }
        switch "\n\t" {
        case "foo bar":
            for fi46 := 0; fi46 < 8; fi46++ {
                delete(v2, fi46)
                _ = fi46
            }
        case "hello":
            var v47 int = (v5 * v5)
            _ = v47
        default:
            v5 = v5
        }
    }
    if (!(4294967295 > 0)) {
        v48 := []int{-2147483648, 1, 32767}
        _ = v48
    } else {
        var v49 bool = v4
        _ = v49
    }
    if (35 < 98) {
        var v50 float32 = (1.5 / (0.1 * 1.5))
        _ = v50
    }
    fmt.Println(v5)
    for fi51 := 0; fi51 < 11; fi51++ {
        if (false || v4) {
            fmt.Println(v3, v3, v4)
        } else {
            v52 := GenS5{F0: 4294967295, F1: (v5 % ((0) | 1)), F2: 127}
            _ = v52
        }
        v5 -= (v5 + (v5 * v5))
        for fi53 := 0; fi53 < 19; fi53++ {
            _ = len(v2)
            v54 := GenS6{F0: 2, F1: 768}
            _ = fi53
            _ = v54
        }
        _ = fi51
    }
    v4 = true
    if (255 > 65535) {
        fmt.Println(v4, v4, v4)
    } else {
        v2[v5] = 0
        if v4 {
            var v55 uint64 = 4294967295
            switch (v5 % ((-1) | 1)) {
            case -9223372036854775808:
                fmt.Println(v3)
            }
            _ = v55
        }
        fmt.Println(v5, v4)
    }
    v3 = (1e10 / float64((0.1 * (0.0 + float32(1e10)))))
    var v56 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
