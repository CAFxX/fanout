package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint32; F1 string; F2 int16 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int8; F1 bool; F2 int64 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int32 }
func genFn0(p0 int32) uint16 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint = 1
        var v7 float64 = 0.1
        if true {
            for fi8 := 0; fi8 < 11; fi8++ {
                p0 += p0
                _ = fi8
            }
            v9 := GenS1{F0: p0, F1: v6}
            v10 := GenS2{F0: 813, F1: 2, F2: fi5}
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
        var v11 int64 = 9223372036854775807
        _ = v11
    }
    _ = p0
    return 2
}
func genFn1(p0 int8, p1 uint64) int32 {
    if (!false) {
        switch 32768 {
        case 0:
            p0 = 1
        case -9223372036854775808:
            var v12 int64 = 177
            _ = v12
        }
    } else {
        if true {
            fmt.Println(p0, p1)
            if (!true) {
                var v13 string = "foo bar"
                v14 := GenS3{F0: 42, F1: 42}
                if (!false) {
                    switch int(p0) {
                    case 127:
                        for fi15 := 0; fi15 < 11; fi15++ {
                            var v16 float32 = -2.25
                            _ = fi15
                            _ = v16
                        }
                    case 2147483647:
                        v17 := []int{43, 137, 127}
                        _ = v17
                    }
                    switch 128 {
                    case 877:
                        v18 := []int{}
                        _ = v18
                    case -9223372036854775808:
                        for fi19 := 0; fi19 < 4; fi19++ {
                            v20 := GenS1{F0: -1, F1: 1}
                            p0 += 99
                            p1 = (p1 ^ (p1 ^ (p1 & (p1 & p1))))
                            _ = fi19
                            _ = v20
                        }
                    case -441:
                        fmt.Println(p1)
                    }
                }
                _ = v13
                _ = v14
            } else {
                switch 128 {
                case 30:
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 int = fi21
                        p1 = p1
                        _ = fi21
                        _ = v22
                    }
                case -1:
                    var v23 bool = (false || false)
                    _ = v23
                case 0:
                    if true {
                        v24 := GenS3{F0: 255, F1: 1000}
                        p0 = p0
                        _ = v24
                    } else {
                        p1 = p1
                        v25 := GenS4{F0: 183, F1: "\n\t", F2: 331}
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 int8 = (p0 + 100)
                    v28 := GenS5{F0: 42, F1: -1}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                fmt.Println(p0, p1)
            }
        } else {
            p0 += 127
            if false {
                switch 2147483647 {
                case -196:
                    var v29 uint16 = 2
                    _ = v29
                case -399:
                    var v30 uint8 = 4
                    _ = v30
                }
                if false {
                    v31 := GenS6{F0: p0, F1: (!true), F2: int64(p0)}
                    p1 -= p1
                    _ = v31
                }
                v32 := []int{-2147483648, 0, -9223372036854775808, 32767}
                _ = v32
            } else {
                switch 128 {
                case -965:
                    var v33 uint = 1
                    _ = v33
                case -9223372036854775808:
                    for fi34 := 0; fi34 < 16; fi34++ {
                        for fi35 := 0; fi35 < 10; fi35++ {
                            var v36 uint16 = 2
                            v37 := GenS3{F0: 255, F1: 42}
                            if true {
                                v38 := []int{fi35, 76, 9223372036854775807}
                                switch (fi35 % ((fi34) | 1)) {
                                case 32768:
                                    _ = len(v38)
                                }
                                _ = v38
                            } else {
                                v39 := []int{(^(fi34)), -933}
                                _ = len(v39)
                                var v40 float32 = 0.1
                                _ = v39
                                _ = v40
                            }
                            v41 := []int{-9223372036854775808, (fi35 | 1)}
                            if len(v41) >= 2 {
                                v41 = v41[1:len(v41)]
                            }
                            var v42 uint64 = (p1 >> ((1) & 63))
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v41
                            _ = v42
                        }
                        p0 = (p0 + (^(p0)))
                        _ = fi34
                    }
                case 1:
                    var v43 uint = 721
                    _ = v43
                default:
                    v44 := []int{-20, 139}
                    _ = v44
                }
            }
        }
        var v45 uint16 = 255
        fmt.Println(p0, v45)
        _ = v45
    }
    var v46 int64 = 1
    p0 = (p0 >> ((4294967295) & 63))
    fmt.Println(v46, p1)
    _ = p0
    _ = p1
    _ = v46
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 42
    v3 := []string{("hello" + ((("\n\t" + (("hello" + (("foo bar" + "") + (("foo bar" + ("" + "a")) + "\n\t"))) + "hello")) + "hello") + "")), "\n\t", "a", ""}
    var v4 uint8 = v2
    if true {
        var v47 uint = 0
        var v48 uint64 = 328
        v1["hello"] = true
        _ = v47
        _ = v48
    } else {
        v49 := GenS3{F0: (v2 | 127), F1: 32767}
        fmt.Println(v2, v2)
        _ = v49
    }
    switch 32767 {
    case 42:
        for fi50 := 0; fi50 < 14; fi50++ {
            for fi51 := 0; fi51 < 6; fi51++ {
                var v52 bool = (0 == genFn0(1))
                var v53 uint16 = 2
                var v54 float32 = -2.25
                switch (fi50 << ((18446744073709551615) & 63)) {
                case 4294967295:
                    var v55 uint64 = 4294967295
                    _ = v55
                case 32767:
                    var v56 int = fi51
                    _ = v56
                default:
                    v52 = ((-128 != 75) && true)
                }
                var v57 float32 = (float32(float64(0.0)) - (1.5 - -2.25))
                _ = fi51
                _ = v52
                _ = v53
                _ = v54
                _ = v57
            }
            _ = fi50
        }
    case 790:
        for fi58 := 0; fi58 < 15; fi58++ {
            switch -99 {
            case 4294967295:
                v3 = append(v3, "\n\t")
            case 9223372036854775807:
                if true {
                    v59 := GenS1{F0: 1, F1: 18446744073709551615}
                    _ = v59
                } else {
                    switch fi58 {
                    case 852:
                        v1["a"] = (((1 >= -656) && false) && true)
                    }
                    var v60 int32 = -729
                    _ = v60
                }
            case 776:
                v61 := GenS7{F0: 1, F1: 42}
                _ = v61
            }
            var v62 uint16 = 388
            fmt.Println(v4, v4, v4)
            _ = fi58
            _ = v62
        }
    case 128:
        var v63 uint16 = 32767
        _ = v63
    }
    switch -875 {
    case 32767:
        fmt.Println(v4, v2)
    case 127:
        v64 := GenS8{F0: 1}
        _ = v64
    case -1:
        v4 = (v4 | 218)
    default:
        v4 = v4
    }
    for fi65 := 0; fi65 < 9; fi65++ {
        v2 -= v4
        var v66 float32 = 1e10
        fmt.Println(fi65, v66, v2)
        _ = fi65
        _ = v66
    }
    v67 := GenS7{F0: 2, F1: 1}
    v3 = append(v3, ((("foo bar" + "hello") + "") + "a"))
    fmt.Println(v2, v4)
    v68 := v1["a"]
    v4 = v2
    v68 = (!v68)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v68
}
