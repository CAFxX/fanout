package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int16, p1 int16) uint {
    p1 += (p0 * 32767)
    if (-9223372036854775808 < 472) {
        var v7 float32 = 3.14159
        var v8 int = 589
        if (v7 >= (((float32(((float64((float32(float64((v7 - v7))) - (1.5 * (v7 * (0.1 / float32(3.14159)))))) / 3.14159) / 1.5)) - v7) * float32(1.5)) / v7)) {
            var v9 float64 = 0.0
            for fi10 := 0; fi10 < 10; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    p0 = (p0 / (((^(p0))) | 1))
                    if true {
                        if true {
                            v7 -= v7
                        }
                        var v12 uint32 = 2
                        if true {
                            v9 += v9
                            v13 := []int{fi10, 32768}
                            var v14 int8 = -1
                            var v15 int8 = v14
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            v16 := []int{(v8 & fi11)}
                            var v17 int32 = -1
                            var v18 uint = 0
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        var v19 uint8 = 255
                        _ = v12
                        _ = v19
                    } else {
                        v20 := []int{42}
                        _ = v20
                    }
                    if true {
                        if true {
                            v21 := []int{v8, fi10, (fi11 & (fi11 % ((fi11) | 1))), 2147483647}
                            v21 = append(v21, (fi10 | v8))
                            _ = v21
                        }
                    } else {
                        v22 := []int{}
                        switch (fi11 | fi10) {
                        case 4294967295:
                            var v23 int8 = -7
                            _ = v23
                        case 42:
                            p0 = 32767
                        case 9223372036854775807:
                            v24 := GenS1{F0: -9223372036854775808}
                            _ = v24
                        default:
                            var v25 uint16 = 0
                            _ = v25
                        }
                        v26 := GenS1{F0: -103}
                        _ = v22
                        _ = v26
                    }
                    var v27 bool = true
                    p0 += p0
                    _ = fi11
                    _ = v27
                }
                _ = fi10
            }
            if false {
                fmt.Println(p0)
                for fi28 := 0; fi28 < 16; fi28++ {
                    if false {
                        switch v8 {
                        case -899:
                            var v29 uint16 = 36
                            _ = v29
                        case 128:
                            for fi30 := 0; fi30 < 17; fi30++ {
                                v31 := GenS1{F0: -9223372036854775808}
                                _ = fi30
                                _ = v31
                            }
                        case 979:
                            var v32 int64 = -279
                            _ = v32
                        }
                        var v33 bool = (true || (false && false))
                        _ = v33
                    } else {
                        v34 := GenS1{F0: -267}
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 12; fi35++ {
                        var v36 float32 = ((v7 / v7) / 1.5)
                        _ = fi35
                        _ = v36
                    }
                    var v37 int8 = 107
                    _ = fi28
                    _ = v37
                }
                var v38 int = -9223372036854775808
                _ = v38
            }
            _ = v9
        } else {
            v8 += v8
            v39 := GenS1{F0: -785}
            _ = v39
        }
        _ = v7
        _ = v8
    } else {
        fmt.Println(p1)
        var v40 int8 = 37
        for fi41 := 0; fi41 < 20; fi41++ {
            if (0 <= 255) {
                var v42 int16 = p1
                if true {
                    p1 = (-(v42))
                    p1 = p0
                }
                v43 := GenS1{F0: -9223372036854775808}
                _ = v42
                _ = v43
            }
            _ = fi41
        }
        _ = v40
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    var v4 uint16 = 485
    v5 := []int{32768, -9223372036854775808, 2147483647, 2147483647}
    v6 := make(map[uint32]string)
    if false {
        v4 += v4
        v3 -= v3
    } else {
        var v44 float64 = (float64(3.14159) * 1.5)
        _ = v44
    }
    var v45 int64 = -1
    if (false && false) {
        v45 += 0
    } else {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            v1 = v1
        case "hello":
            fmt.Println(v3, v4, v45)
        case "a":
            v46 := GenS1{F0: 1}
            _ = v46
        default:
            fmt.Println(v1)
        }
        for fi47 := 0; fi47 < 2; fi47++ {
            for fi48 := 0; fi48 < 15; fi48++ {
                v45 += (v45 << ((42) & 63))
                _ = len(v2)
                v3 -= 1
                _ = fi48
            }
            _ = fi47
        }
        switch 32768 {
        case 2147483647:
            switch 612 {
            case 42:
                var v49 uint16 = (v4 & (^(v4)))
                _ = v49
            case -9223372036854775808:
                v50 := GenS1{F0: int64(v4)}
                _ = v50
            }
        case 42:
            var v51 bool = false
            _ = v51
        default:
            v1 = 2147483647
        }
    }
    v45 -= (v45 | v45)
    v52 := GenS1{F0: v45}
    _ = len(v5)
    fmt.Println(v3)
    if false {
        v1 -= (v1 - (v1 << ((2) & 63)))
    }
    if true {
        var v53 uint = 1
        _ = v53
    } else {
        if true {
            for fi54 := 0; fi54 < 12; fi54++ {
                v55 := GenS1{F0: 9223372036854775807}
                _ = len(v2)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi54
                _ = v55
            }
            v45 += v45
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v56 := GenS2{F0: 42}
        _ = v56
    }
    var v57 int16 = 42
    if true {
        v5 = append(v5, 2147483647)
        if (1.5 >= ((float32(0.1) + 1e10) + 3.14159)) {
            for fi58 := 0; fi58 < 17; fi58++ {
                if ((v45 % ((-1) | 1)) < v45) {
                    for fi59 := 0; fi59 < 7; fi59++ {
                        if (42 >= 4294967295) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v45 = (v45 | v45)
                            v45 += (v45 << ((18446744073709551615) & 63))
                        }
                        v60 := GenS3{F0: 1}
                        var v61 uint8 = 148
                        if (v61 != (v61 | uint8(v4))) {
                            switch (fi58 >> ((42) & 63)) {
                            case 424:
                                v57 += v57
                            case 32768:
                                var v62 uint8 = (v61 & v61)
                                _ = v62
                            default:
                                if (0.0 == (3.14159 - float32(0.1))) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(v4, fi59)
                                }
                            }
                            v63 := GenS1{F0: v45}
                            var v64 float32 = 1e10
                            _ = v63
                            _ = v64
                        }
                        var v65 int16 = (-(v57))
                        _ = fi59
                        _ = v60
                        _ = v61
                        _ = v65
                    }
                } else {
                    for fi66 := 0; fi66 < 8; fi66++ {
                        var v67 int16 = 32767
                        _ = fi66
                        _ = v67
                    }
                    if ((v57 <= (v57 << ((2) & 63))) && true) {
                        var v68 uint64 = 18446744073709551615
                        _ = v68
                    }
                    for fi69 := 0; fi69 < 8; fi69++ {
                        _ = len(v5)
                        fmt.Println(v4)
                        _ = fi69
                    }
                }
                v3 -= v3
                _ = fi58
            }
            if false {
                for fi70 := 0; fi70 < 16; fi70++ {
                    for fi71 := 0; fi71 < 10; fi71++ {
                        v45 = 9223372036854775807
                        _ = fi71
                    }
                    _ = len(v6)
                    v72 := GenS2{F0: 1}
                    _ = fi70
                    _ = v72
                }
                var v73 int8 = -66
                v4 = v4
                _ = v73
            } else {
                if true {
                    var v74 int32 = v3
                    var v75 float32 = (0.1 + 1.5)
                    fmt.Println(v75, v74, v45)
                    _ = v74
                    _ = v75
                } else {
                    var v76 bool = ((v1 ^ v1) <= (v1 ^ (^(v1))))
                    _ = v76
                }
                v77 := GenS4{F0: 2}
                _ = v77
            }
        }
    }
    var v78 string = ("" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v52
    _ = v57
    _ = v6
    _ = v78
}
