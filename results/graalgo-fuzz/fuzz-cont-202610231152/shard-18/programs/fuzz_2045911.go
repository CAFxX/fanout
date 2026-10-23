package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 int, p1 int8, p2 int8) uint64 {
    var v6 bool = (0 <= 32767)
    p2 = p2
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{fi8, fi8, 9223372036854775807, (^(p0))}
            if (4294967295 != 239) {
                var v10 uint64 = 603
                v11 := GenS1{F0: v10}
                if false {
                    v6 = v6
                    p0 += (fi7 * (p0 / ((fi8) | 1)))
                    var v12 uint64 = v10
                    _ = v12
                }
                _ = v10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        if v6 {
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 uint64 = 730
                _ = fi13
                _ = v14
            }
            p1 = (p1 % (((p2 ^ 57)) | 1))
            fmt.Println(p2, p0, p0)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1() uint8 {
    if true {
        for fi15 := 0; fi15 < 6; fi15++ {
            if (65535 <= 255) {
                var v16 float32 = 0.0
                v17 := GenS2{F0: 2}
                v18 := GenS2{F0: 127}
                _ = v16
                _ = v17
                _ = v18
            } else {
                fmt.Println(fi15)
            }
            fmt.Println(fi15)
            _ = fi15
        }
        v19 := []int{42, 4294967295}
        var v20 uint = 18446744073709551615
        _ = v19
        _ = v20
    }
    var v21 uint = 18446744073709551615
    var v22 string = "hello"
    if false {
        if (false || true) {
            var v23 bool = ((1 >= 134) && true)
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := GenS1{F0: 815}
                _ = fi24
                _ = v25
            }
            _ = v23
        } else {
            v22 = ("\n\t" + (v22 + v22))
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            switch (^(fi26)) {
            case -2147483648:
                var v27 int64 = 42
                _ = v27
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := []int{2147483647}
                var v30 int32 = 538
                fmt.Println(v21)
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi26
        }
        v31 := GenS3{F0: 0, F1: v21, F2: 2147483647}
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            if false {
                fmt.Println(fi33)
            }
            _ = fi33
        }
        v34 := []int{(fi32 | (fi32 - (fi32 + 9223372036854775807)))}
        _ = fi32
        _ = v34
    }
    _ = v21
    _ = v22
    return 150
}

func main() {
    var v1 int32 = 197
    v2 := make(map[int]uint32)
    var v3 uint8 = 127
    var v4 string = "hello"
    var v5 string = (("foo bar" + v4) + v4)
    v3 = (v3 / ((v3) | 1))
    v35 := GenS1{F0: 0}
    fmt.Println(v4, v5, v4)
    fmt.Println(v1, v5)
    v36 := []int{}
    for fi37 := 0; fi37 < 1; fi37++ {
        delete(v2, -2147483648)
        _ = fi37
    }
    if (!true) {
        var v38 uint8 = v3
        v38 -= v38
        v3 += 42
        _ = v38
    } else {
        for fi39 := 0; fi39 < 14; fi39++ {
            v4 = v4
            if len(v36) > 0 {
                v36[0] = fi39
            }
            v2[(fi39 * fi39)] = 2
            _ = fi39
        }
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
        if false {
            if (!false) {
                if (true || true) {
                    fmt.Println(v3)
                    var v40 uint64 = 4294967295
                    var v41 uint16 = 42
                    _ = v40
                    _ = v41
                }
                switch 752 {
                case 32768:
                    var v42 bool = (!false)
                    _ = v42
                case -230:
                    if (((^(v3)) > (v3 << ((uint(v3)) & 63))) && ((42 != 42) || false)) {
                        _ = len(v36)
                        var v43 uint32 = 2147483647
                        _ = v43
                    } else {
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                        var v44 int64 = 485
                        _ = v44
                    }
                default:
                    var v45 int64 = -907
                    _ = v45
                }
            } else {
                switch 32767 {
                case 127:
                    fmt.Println(v4, v3)
                case -9223372036854775808:
                    if len(v36) > 0 {
                        v36[0] = 781
                    }
                case -2147483648:
                    v4 = v4
                }
                v46 := GenS4{F0: -385, F1: 32767, F2: 42}
                _ = v46
            }
            fmt.Println(v1, v1)
        } else {
            if false {
                var v47 int8 = -89
                switch v4 {
                case "hello":
                    if (42 < 0) {
                        var v48 float64 = (1e10 * 3.14159)
                        _ = v48
                    } else {
                        v47 += (v47 | int8(v1))
                    }
                case "":
                    switch ("" + v4) {
                    case "foo bar":
                        v49 := GenS2{F0: 229}
                        _ = v49
                    case "":
                        if false {
                            v3 -= v3
                            if len(v36) > 0 {
                                v36[0] = 32768
                            }
                            fmt.Println(v47)
                        }
                    }
                default:
                    v50 := GenS4{F0: 0, F1: 380, F2: 2}
                    _ = v50
                }
                _ = v47
            }
            var v51 uint64 = 2
            _ = v51
        }
    }
    switch 2147483647 {
    case 1:
        var v52 uint64 = 4294967295
        _ = v52
    default:
        switch v4 {
        case "foo bar":
            v1 += (v1 * v1)
        case "":
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
        case "\n\t":
            v53 := GenS4{F0: -913, F1: 0, F2: 328}
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
