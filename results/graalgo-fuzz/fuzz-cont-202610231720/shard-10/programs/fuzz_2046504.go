package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64; F1 int32; F2 uint }
func genFn0(p0 uint, p1 int, p2 uint16) uint32 {
    v7 := GenS1{F0: -806, F1: 42}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint16 = p2
        v10 := []int{835}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        v11 := []int{9223372036854775807}
        _ = v11
    case "\n\t":
        var v12 uint64 = 0
        _ = v12
    case "a":
        var v13 bool = true
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 float64 = -2.25
    var v4 int8 = -1
    v5 := make(map[string]bool)
    v6 := []uint32{2147483647}
    switch "foo bar" {
    case "":
        if true {
            switch "" {
            case "a":
                for fi14 := 0; fi14 < 4; fi14++ {
                    v6 = append(v6, 0)
                    fmt.Println(fi14, v3, v3)
                    var v15 int32 = 0
                    _ = fi14
                    _ = v15
                }
            case "":
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 uint64 = 18446744073709551615
                    var v18 int8 = -82
                    fmt.Println(v3, v3)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            case "hello":
                for fi19 := 0; fi19 < 19; fi19++ {
                    if len(v6) > 0 {
                        v6[0] = 42
                    }
                    _ = fi19
                }
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 float64 = ((v3 + v3) * 1e10)
                _ = fi20
                _ = v21
            }
            var v22 float32 = 1.5
            _ = v22
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 uint32 = 65535
                _ = fi23
                _ = v24
            }
            if (!(!true)) {
                v4 = (v4 << ((2) & 63))
                for fi25 := 0; fi25 < 18; fi25++ {
                    v3 -= -2.25
                    v26 := GenS2{F0: 2147483647, F1: 320}
                    if false {
                        if (true || (43 <= 0)) {
                            v4 = v4
                            fmt.Println(v4, fi25, v3)
                        }
                    } else {
                        delete(v2, 392)
                        if false {
                            v6 = append(v6, 977)
                        } else {
                            v3 += 0.0
                            v4 = v4
                            if (65535 < 853) {
                                var v27 float64 = v3
                                _ = v27
                            } else {
                                v4 += (v4 - 10)
                                v4 += (v4 / ((-124) | 1))
                                v3 = 1.5
                            }
                        }
                        var v28 int64 = 9223372036854775807
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
                v29 := v2[42]
                _ = v29
            } else {
                if (42 > 299) {
                    var v30 float32 = 0.0
                    for fi31 := 0; fi31 < 20; fi31++ {
                        v32 := GenS1{F0: -196, F1: -32768}
                        v3 = (float64(-2.25) / (v3 + (v3 + 1e10)))
                        fmt.Println(v3, v30)
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                } else {
                    _ = len(v6)
                }
                switch "foo bar" {
                case "":
                    v33 := v2[777]
                    _ = v33
                case "a":
                    for fi34 := 0; fi34 < 20; fi34++ {
                        v3 += -2.25
                        _ = fi34
                    }
                case "foo bar":
                    switch "hello" {
                    case "foo bar":
                        switch ("foo bar" + ("hello" + "")) {
                        case "hello":
                            fmt.Println(v3, v3)
                        case "a":
                            var v35 int32 = 1
                            _ = v35
                        }
                    case "\n\t":
                        v4 = v4
                    case "":
                        var v36 int64 = 0
                        _ = v36
                    default:
                        var v37 int64 = 9223372036854775807
                        _ = v37
                    }
                default:
                    v38 := GenS2{F0: 536, F1: 733}
                    _ = v38
                }
            }
            v3 -= v3
        }
    }
    _ = len(v6)
    switch "\n\t" {
    case "foo bar":
        v39 := GenS2{F0: 1, F1: 1}
        _ = v39
    case "":
        if true {
            v40 := GenS3{F0: 9223372036854775807, F1: -1, F2: 0}
            v4 -= v4
            _ = v40
        } else {
            fmt.Println(v3)
            v41 := GenS1{F0: 42, F1: 1}
            _ = v41
        }
    case "a":
        delete(v2, 2147483647)
    }
    if false {
        v3 += 3.14159
        fmt.Println(v4)
        for fi42 := 0; fi42 < 15; fi42++ {
            if true {
                v3 -= 1.5
                var v43 string = "\n\t"
                var v44 float64 = (v3 / -2.25)
                _ = v43
                _ = v44
            } else {
                var v45 int = (^(fi42))
                if ((!(!((v4 % (((v4 << ((42) & 63))) | 1)) >= (v4 % ((v4) | 1))))) || (true && true)) {
                    switch (-(fi42)) {
                    case 127:
                        var v46 uint8 = 0
                        _ = v46
                    case 1:
                        for fi47 := 0; fi47 < 4; fi47++ {
                            v3 += v3
                            fmt.Println(fi47, v3, fi47)
                            _ = fi47
                        }
                    }
                    var v48 uint16 = 2
                    _ = v48
                }
                _ = v45
            }
            var v49 uint64 = 42
            v49 = (v49 / (((^(v49))) | 1))
            _ = fi42
            _ = v49
        }
    }
    var v50 uint8 = 120
    var v51 bool = ((1e10 / ((((v3 * (float64(3.14159) * 1.5)) + v3) * 3.14159) * 0.0)) > (0.0 / ((float64((3.14159 - ((1.5 + 0.0) * 3.14159))) - v3) / 3.14159)))
    var v52 int8 = int8(v50)
    v53 := GenS3{F0: 763, F1: 0, F2: 510}
    var v54 uint64 = 42
    var v55 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
