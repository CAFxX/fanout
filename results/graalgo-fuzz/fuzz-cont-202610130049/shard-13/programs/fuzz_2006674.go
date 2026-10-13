package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int64; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 bool; F2 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 string; F2 uint16 }
func genFn0(p0 int64) uint64 {
    v5 := []int{}
    fmt.Println(p0)
    var v6 int64 = p0
    p0 -= 42
    _ = p0
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 uint8, p1 uint8) int64 {
    var v7 int32 = 1
    var v8 bool = false
    v8 = true
    var v9 int64 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (v9 & (-(v9)))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 957
    var v3 uint16 = 65535
    var v4 int16 = -886
    switch "foo bar" {
    case "":
        if (!true) {
            var v10 int16 = (-(v4))
            _ = v10
        } else {
            v3 += v3
            v2 += -902
            v4 -= v4
        }
    case "foo bar":
        if (702 <= 65535) {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int64 = 386
                fmt.Println(v3, fi11, v3)
                v4 = 9
                _ = fi11
                _ = v12
            }
            v4 += 1000
            fmt.Println(v3, v2, v2)
        }
    case "\n\t":
        if (-746 == -9223372036854775808) {
            v13 := GenS1{F0: 100, F1: 0, F2: 128}
            _ = v13
        } else {
            v14 := GenS1{F0: -84, F1: 0, F2: 2}
            _ = v14
        }
    default:
        v2 -= (v2 >> ((1) & 63))
    }
    switch "foo bar" {
    case "hello":
        for fi15 := 0; fi15 < 11; fi15++ {
            if (86 >= 255) {
                switch "foo bar" {
                case "foo bar":
                    var v16 bool = (0 <= 263)
                    _ = v16
                case "\n\t":
                    var v17 int32 = -220
                    _ = v17
                case "a":
                    v18 := []int{}
                    _ = v18
                }
                var v19 int16 = v4
                fmt.Println(v3, v2, v3)
                _ = v19
            } else {
                v4 = v4
                for fi20 := 0; fi20 < 1; fi20++ {
                    v4 += (v4 >> ((255) & 63))
                    v4 = v4
                    _ = fi20
                }
                var v21 uint64 = genFn0(1)
                _ = v21
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                switch -1 {
                case 42:
                    v3 += (^(v3))
                case 128:
                    v4 = v4
                    var v23 uint64 = 4294967295
                    var v24 uint16 = v3
                    _ = v23
                    _ = v24
                case 32768:
                    switch "a" {
                    case "hello":
                        v2 = fi22
                    }
                }
                _ = fi22
            }
            if (-2.25 > 1e10) {
                delete(v1, "hello")
            }
            _ = fi15
        }
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v1[("\n\t" + ((("\n\t" + "\n\t") + ("\n\t" + "a")) + "hello"))] = (1 <= 127)
        v4 += 32767
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 bool = false
            v28 := GenS1{F0: 100, F1: 107, F2: 128}
            if v27 {
                v29 := GenS2{F0: 336}
                _ = v29
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    v30 := GenS3{F0: -128, F1: 1, F2: ((-2.25 + float32(((float64((1e10 - 3.14159)) + -2.25) / (1.5 + -2.25)))) < float32(((0.0 + 1e10) * -2.25)))}
    v4 += -1
    v31 := []int{v2, v2, 170, (v2 / ((154) | 1))}
    var v32 float32 = (3.14159 - -2.25)
    v32 += v32
    if (-2147483648 > 2147483647) {
        v33 := GenS3{F0: 42, F1: -556, F2: ((!false) || true)}
        _ = v33
    } else {
        v2 -= (-(v2))
        for fi34 := 0; fi34 < 15; fi34++ {
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 uint64 = 42
                _ = fi35
                _ = v36
            }
            var v37 uint32 = 2
            v3 -= v3
            _ = fi34
            _ = v37
        }
        switch (v2 + (v2 & (v2 >> ((0) & 63)))) {
        case 128:
            if true {
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 int8 = -1
                    if (-9223372036854775808 <= genFn1(146, 79)) {
                        v40 := GenS1{F0: 0, F1: (v39 << ((255) & 63)), F2: 210}
                        _ = v40
                    } else {
                        switch "a" {
                        case "":
                            v41 := GenS3{F0: 42, F1: 42, F2: true}
                            _ = v41
                        default:
                            v42 := GenS4{F0: 18446744073709551615}
                            _ = v42
                        }
                    }
                    _ = fi38
                    _ = v39
                }
                switch (v2 & v2) {
                case -855:
                    var v43 uint8 = 42
                    _ = v43
                case -1:
                    v4 += -32768
                case -304:
                    switch "" {
                    case "":
                        var v44 int64 = -1
                        _ = v44
                    case "hello":
                        v45 := GenS5{F0: 2, F1: true, F2: -9223372036854775808}
                        _ = v45
                    }
                default:
                    if len(v31) >= 2 {
                        v31 = v31[1:len(v31)]
                    }
                }
                v4 -= v4
            } else {
                var v46 int8 = -90
                if true {
                    v47 := GenS6{F0: v2}
                    if true {
                        var v48 bool = (false && (false || (!true)))
                        _ = v48
                    } else {
                        if false {
                            var v49 float32 = (v32 + 1e10)
                            v50 := GenS7{F0: 460, F1: (("" + (((("foo bar" + "") + (("\n\t" + "hello") + "")) + "\n\t") + ("foo bar" + ("" + ("foo bar" + "\n\t"))))) + "foo bar"), F2: 352}
                            _ = v49
                            _ = v50
                        } else {
                            for fi51 := 0; fi51 < 16; fi51++ {
                                v32 -= (1e10 * ((0.0 - -2.25) + ((float32(3.14159) * v32) + ((float32(1e10) / v32) + (((v32 * v32) * v32) - v32)))))
                                _ = fi51
                            }
                            var v52 int = (v2 << ((255) & 63))
                            _ = v52
                        }
                    }
                    v32 += v32
                    _ = v47
                }
                _ = v46
            }
        case -1:
            for fi53 := 0; fi53 < 18; fi53++ {
                var v54 uint64 = 42
                if true {
                    for fi55 := 0; fi55 < 13; fi55++ {
                        var v56 bool = false
                        fmt.Println(v2)
                        _ = fi55
                        _ = v56
                    }
                    v1["hello"] = (float64(((v32 + 0.1) * (float32(1e10) * ((v32 / v32) * 0.1)))) != 3.14159)
                }
                _ = fi53
                _ = v54
            }
        default:
            fmt.Println(v2)
        }
    }
    v2 -= v2
    v57 := GenS1{F0: -1, F1: -1, F2: 165}
    v31 = append(v31, v2)
    v1[""] = true
    var v58 uint32 = 579
    var v59 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v57
    _ = v58
    _ = v59
}
