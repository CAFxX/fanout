package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
func genFn0(p0 int64, p1 int32, p2 uint8) int {
    var v5 int32 = p1
    p0 += (p0 / ((p0) | 1))
    var v6 int64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1() string {
    switch 4294967295 {
    case 32768:
        v7 := []int{-1, 0, 32767, 0}
        _ = v7
    default:
        var v8 uint8 = 42
        _ = v8
    }
    var v9 int = 2147483647
    _ = v9
    return ""
}

func main() {
    var v1 int = 79
    v2 := make(map[int]uint32)
    var v3 int16 = 32767
    v4 := []uint32{42, 139}
    if true {
        if (42 > 1) {
            if true {
                fmt.Println(v3, v1)
                v3 += v3
                v3 = v3
            } else {
                if ((v3 <= v3) && true) {
                    v1 -= 0
                    for fi10 := 0; fi10 < 20; fi10++ {
                        v11 := GenS1{F0: 42, F1: 2147483647, F2: 65535}
                        v2[genFn0(-156, 2147483647, 226)] = 2147483647
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS1{F0: 184, F1: 2147483647, F2: 726}
                    _ = v12
                }
            }
            v2[(-(v1))] = 1
        } else {
            if false {
                v13 := GenS2{F0: 350}
                v1 = 127
                _ = v13
            } else {
                for fi14 := 0; fi14 < 16; fi14++ {
                    fmt.Println(v3, v3)
                    if (!(!true)) {
                        v15 := GenS1{F0: 65535, F1: 42, F2: 1}
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 7; fi16++ {
                            if true {
                                var v17 bool = ((true || true) || false)
                                v18 := GenS3{F0: (v1 + (fi14 & genFn0(11, 1, 20))), F1: (fi16 ^ (v1 / (((-(fi14))) | 1))), F2: -394}
                                fmt.Println(fi16, v17)
                                _ = len(v2)
                                _ = v17
                                _ = v18
                            }
                            if false {
                                v1 = 9223372036854775807
                            } else {
                                var v19 uint32 = 2147483648
                                v20 := GenS2{F0: v19}
                                fmt.Println(v19)
                                _ = v19
                                _ = v20
                            }
                            var v21 float32 = float32((1.5 / ((1e10 * 0.0) + (1e10 + float64(-2.25)))))
                            _ = fi16
                            _ = v21
                        }
                        switch "\n\t" {
                        case "hello":
                            v3 -= (v3 | (^(v3)))
                        case "foo bar":
                            for fi22 := 0; fi22 < 9; fi22++ {
                                var v23 float64 = ((0.0 / -2.25) / 1e10)
                                var v24 uint16 = 65535
                                if (500 < 349) {
                                    var v25 float32 = 1.5
                                    v23 = (v23 / v23)
                                    var v26 int16 = v3
                                    _ = v25
                                    _ = v26
                                } else {
                                    fmt.Println(v24, fi22)
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        }
                        switch fi14 {
                        case -2147483648:
                            var v27 int64 = -9223372036854775808
                            _ = v27
                        case 127:
                            v28 := v2[fi14]
                            _ = v28
                        case 9223372036854775807:
                            _ = len(v4)
                        }
                    }
                    _ = fi14
                }
                var v29 float64 = 3.14159
                fmt.Println(v3, v29)
                _ = v29
            }
        }
        if false {
            v3 = v3
        }
        v4 = append(v4, 65535)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v30 int64 = -9223372036854775808
    _ = len(v2)
    if (!(float64(1.5) == 0.1)) {
        for fi31 := 0; fi31 < 16; fi31++ {
            fmt.Println(v3, v3, fi31)
            fmt.Println(v3, fi31, v3)
            _ = fi31
        }
        v32 := GenS4{F0: 2}
        if (false || false) {
            v33 := GenS5{F0: 1, F1: (v3 | 1)}
            switch "foo bar" {
            case "hello":
                v34 := GenS4{F0: 42}
                _ = v34
            case "a":
                if (0.1 == (1e10 + 0.0)) {
                    fmt.Println(v1)
                    fmt.Println(v1, v30, v1)
                    v1 = (-(v1))
                }
            case "foo bar":
                fmt.Println(v30, v1, v30)
            default:
                v1 += (v1 ^ (v1 << ((375) & 63)))
            }
            _ = v33
        }
        _ = v32
    }
    var v35 uint64 = 0
    fmt.Println(v35, v1, v35)
    var v36 float64 = ((1.5 + ((-2.25 * 1e10) / -2.25)) + 1.5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2[genFn0((v30 + (v30 >> ((42) & 63))), -2147483648, 84)] = 65535
    switch "hello" {
    case "a":
        var v37 uint64 = v35
        _ = v37
    }
    v38 := GenS4{F0: 32767}
    var v39 int16 = v3
    if (v3 != v39) {
        if ((88 != 100) || (false && false)) {
            if true {
                fmt.Println(v36)
            } else {
                v36 += v36
                var v40 string = ""
                _ = v40
            }
            for fi41 := 0; fi41 < 8; fi41++ {
                v4 = append(v4, 42)
                for fi42 := 0; fi42 < 17; fi42++ {
                    if true {
                        var v43 int = -2147483648
                        if false {
                            var v44 int64 = (-(v30))
                            if (false || (!false)) {
                                fmt.Println(v30)
                            }
                            if false {
                                if ((v3 ^ v39) == v3) {
                                    if false {
                                        if ((v36 / v36) >= 1e10) {
                                            switch (fi41 * fi42) {
                                            case -691:
                                                var v45 int64 = -589
                                                _ = v45
                                            case 614:
                                                v2[(v43 + (fi42 + fi41))] = uint32(v35)
                                            case 128:
                                                if false {
                                                    v36 = 1.5
                                                    v44 += 327
                                                    fmt.Println(fi42, fi41, v35)
                                                }
                                                if (false || false) {
                                                    var v46 uint32 = uint32(v35)
                                                    _ = v46
                                                } else {
                                                    var v47 uint8 = 1
                                                    var v48 uint32 = 193
                                                    v49 := GenS3{F0: v1, F1: -906, F2: 1}
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                            }
                                        }
                                    } else {
                                        v39 = (v39 << ((432) & 63))
                                        var v50 uint = 405
                                        v51 := GenS5{F0: 2147483648, F1: v3}
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = len(v4)
                                }
                                v3 = int16(v35)
                                if true {
                                    v52 := GenS4{F0: 913}
                                    var v53 float64 = ((0.0 - (v36 - 1.5)) - 3.14159)
                                    _ = v52
                                    _ = v53
                                } else {
                                    v1 = 1
                                }
                                if (-923 > 1) {
                                    v54 := GenS6{F0: 2}
                                    _ = v54
                                }
                                v4 = append(v4, 4294967295)
                            }
                            _ = v44
                        } else {
                            v30 -= (v30 >> ((42) & 63))
                        }
                        var v55 int = (-(v43))
                        _ = v43
                        _ = v55
                    } else {
                        if ((!((v1 & v1) > genFn0(v30, -2147483648, 255))) && (true || (!true))) {
                            switch v1 {
                            case 32767:
                                v56 := GenS7{F0: 0}
                                _ = v56
                            default:
                                _ = len(v4)
                            }
                            fmt.Println(v3, v3, fi42)
                            var v57 uint8 = 121
                            _ = v57
                        }
                    }
                    fmt.Println(fi42, fi42)
                    _ = fi42
                }
                _ = fi41
            }
        } else {
            v58 := GenS5{F0: 42, F1: 42}
            var v59 string = ""
            v2[(v1 + (v1 & v1))] = 200
            _ = v58
            _ = v59
        }
    }
    v60 := GenS8{F0: true}
    v61 := GenS1{F0: 65535, F1: -850, F2: 65535}
    v62 := GenS6{F0: 654}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v60
    _ = v61
    _ = v62
}
