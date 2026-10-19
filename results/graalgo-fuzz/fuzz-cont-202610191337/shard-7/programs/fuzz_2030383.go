package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint16; F1 int16; F2 int16 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 int32, p1 uint16) int64 {
    p0 = p0
    var v7 uint64 = 4294967295
    v7 = v7
    if (!false) {
        v8 := []int{0}
        _ = v8
    } else {
        v9 := []int{4294967295, -1}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -122
    var v3 int8 = (v2 >> ((18446744073709551615) & 63))
    var v4 bool = true
    v5 := make(map[int]bool)
    var v6 int16 = -32768
    var v10 int16 = (v6 + (v6 ^ v6))
    if (!(-592 <= 205)) {
        if (0.0 > 1.5) {
            v3 += (v2 >> ((2) & 63))
            var v11 string = (("a" + "") + "\n\t")
            _ = v11
        } else {
            v3 -= (^(v3))
            switch "a" {
            case "hello":
                v2 = (^(v2))
            default:
                v4 = (65535 > 42)
            }
            fmt.Println(v10, v4)
        }
    } else {
        v6 = (v10 ^ v10)
    }
    var v12 uint64 = 0
    v13 := []int{-711, -1, -437}
    if ((true && (30 > 42)) && (42 == int32(v3))) {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int64 = 1
            _ = fi14
            _ = v15
        }
    } else {
        v10 = (v10 & (-(v6)))
    }
    switch 0 {
    case 719:
        if (v3 < v3) {
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 int64 = genFn0(1, 562)
                _ = fi16
                _ = v17
            }
            _ = len(v1)
            for fi18 := 0; fi18 < 16; fi18++ {
                for fi19 := 0; fi19 < 16; fi19++ {
                    if (false || (145 < 32767)) {
                        if (42 < -1) {
                            v20 := GenS1{F0: 104}
                            v10 = (v10 % ((v6) | 1))
                            _ = v20
                        }
                        v21 := GenS1{F0: (v10 & (-(v10)))}
                        var v22 int = (fi18 / (((fi18 - 919)) | 1))
                        v23 := GenS1{F0: -219}
                        v13 = append(v13, (v22 / (((v22 | fi19)) | 1)))
                        var v24 int8 = (v2 >> ((608) & 63))
                        if (uint8(v12) < 251) {
                            v2 -= (v24 << ((42) & 63))
                            var v25 int = (v22 % ((fi19) | 1))
                            _ = v25
                        } else {
                            if (383 >= 902) {
                                var v26 uint16 = 1
                                _ = v26
                            }
                        }
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    v4 = (!(2147483647 <= 2147483647))
                    _ = fi19
                }
                switch (fi18 + fi18) {
                case 591:
                    if v4 {
                        for fi27 := 0; fi27 < 1; fi27++ {
                            var v28 uint8 = 42
                            var v29 uint8 = v28
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 20; fi30++ {
                            v10 += v10
                            var v31 int32 = 859
                            _ = fi30
                            _ = v31
                        }
                    }
                case 41:
                    v4 = ((((15 == 0) || (v4 || v4)) || true) && (-9223372036854775808 <= genFn0(0, 2)))
                case 714:
                    v32 := v5[fi18]
                    _ = v32
                default:
                    var v33 uint32 = 42
                    _ = v33
                }
                _ = fi18
            }
        } else {
            delete(v1, 1)
            for fi34 := 0; fi34 < 11; fi34++ {
                v10 = (v6 & (v6 >> ((42) & 63)))
                if (-6 > v3) {
                    switch fi34 {
                    case 9223372036854775807:
                        switch "" {
                        case "foo bar":
                            fmt.Println(v3)
                        case "hello":
                            _ = len(v13)
                        default:
                            if v4 {
                                if (!v4) {
                                    v6 -= (v10 + (v10 + (^(v10))))
                                    var v35 float32 = float32(-2.25)
                                    _ = v35
                                } else {
                                    v6 -= v10
                                    var v36 uint64 = v12
                                    _ = v36
                                }
                            }
                        }
                    }
                }
                _ = fi34
            }
        }
    case 42:
        fmt.Println(v2)
    }
    var v37 int = -133
    v38 := GenS1{F0: v6}
    v37 = (^(v37))
    v10 = (v10 | int16(v3))
    v12 -= 18446744073709551615
    _ = len(v1)
    if (2 <= v12) {
        if ((!v4) || v4) {
            v10 = v10
            fmt.Println(v10, v2)
        } else {
            v37 += (v37 * int(v12))
        }
        switch -810 {
        case 9223372036854775807:
            var v39 int32 = -23
            _ = v39
        case 127:
            v40 := GenS2{F0: 807}
            _ = v40
        case 42:
            v41 := GenS3{F0: (v37 % ((-9223372036854775808) | 1)), F1: 90}
            _ = v41
        }
        if ((v10 + v10) != v6) {
            for fi42 := 0; fi42 < 16; fi42++ {
                fmt.Println(fi42, v6, v6)
                _ = fi42
            }
            for fi43 := 0; fi43 < 13; fi43++ {
                v37 += -425
                _ = fi43
            }
        } else {
            v3 = (v2 << ((857) & 63))
            v37 -= v37
            v44 := GenS3{F0: 127, F1: 1}
            _ = v44
        }
    } else {
        var v45 float32 = 0.1
        _ = v45
    }
    v2 -= (v2 / (((v2 | 0)) | 1))
    v46 := GenS4{F0: 32767, F1: (v10 >> ((255) & 63)), F2: 512}
    var v47 int8 = v2
    if v4 {
        for fi48 := 0; fi48 < 15; fi48++ {
            var v49 int32 = -1
            if (("foo bar" < "foo bar") && (!(0.1 == 3.14159))) {
                v50 := GenS5{F0: 389, F1: 422}
                _ = v50
            }
            _ = fi48
            _ = v49
        }
        v47 = v2
        var v51 uint8 = 42
        _ = v51
    }
    for fi52 := 0; fi52 < 18; fi52++ {
        for fi53 := 0; fi53 < 8; fi53++ {
            v12 = (v12 % (((^(v12))) | 1))
            var v54 uint32 = 65535
            _ = fi53
            _ = v54
        }
        var v55 float32 = 0.1
        _ = fi52
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
