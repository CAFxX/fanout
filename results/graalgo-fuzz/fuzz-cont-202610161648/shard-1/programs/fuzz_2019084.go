package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 bool, p1 int16, p2 bool) uint32 {
    var v7 uint32 = 65535
    var v8 uint32 = 578
    for fi9 := 0; fi9 < 15; fi9++ {
        switch ("" + ((("foo bar" + "foo bar") + "") + "hello")) {
        case "\n\t":
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{1, 9223372036854775807}
                var v12 uint = 18446744073709551615
                _ = fi10
                _ = v11
                _ = v12
            }
        case "foo bar":
            v13 := GenS1{F0: -1, F1: (^(v7)), F2: 0}
            _ = v13
        }
        var v14 int16 = (p1 * (p1 ^ -1))
        v15 := GenS2{F0: 255}
        _ = fi9
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: 127, F1: 74, F2: 361}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    _ = v8
    return (v7 / ((v7) | 1))
}
func genFn1(p0 uint32, p1 uint64, p2 uint64) int {
    p0 += (^(p0))
    fmt.Println(p1, p0, p1)
    if true {
        var v17 uint8 = 127
        v17 = v17
        if false {
            fmt.Println(v17, v17, p2)
            for fi18 := 0; fi18 < 8; fi18++ {
                p2 = uint64(v17)
                var v19 int64 = 451
                v20 := GenS3{F0: -1, F1: 0, F2: uint32(v17)}
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            var v21 int64 = 1
            v22 := []int{}
            if (!false) {
                var v23 int = -9223372036854775808
                _ = v23
            } else {
                if true {
                    v21 = (v21 ^ (v21 << ((255) & 63)))
                } else {
                    v24 := GenS2{F0: v17}
                    var v25 uint32 = p0
                    _ = v24
                    _ = v25
                }
                p2 += 589
                var v26 uint16 = 42
                _ = v26
            }
            _ = v21
            _ = v22
        }
        _ = v17
    } else {
        var v27 int32 = -2147483648
        v28 := GenS1{F0: 100, F1: 0, F2: 223}
        _ = v27
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    v1 := []int{127, 127, -1}
    v2 := []uint32{23, 1, 42}
    var v3 string = "a"
    v4 := make(map[int]string)
    var v5 uint8 = 116
    var v6 uint32 = 2147483648
    for fi29 := 0; fi29 < 6; fi29++ {
        if (65535 == 2) {
            v6 = 27
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch v3 {
            case "hello":
                _ = len(v4)
            case "":
                v6 = v6
            default:
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS1{F0: 127, F1: v6, F2: 42}
                    _ = fi30
                    _ = v31
                }
            }
        }
        var v32 int16 = -912
        _ = fi29
        _ = v32
    }
    switch "foo bar" {
    case "":
        v3 = v3
    case "a":
        if true {
            v4[4294967295] = v3
            _ = len(v4)
        }
    }
    if false {
        v33 := GenS1{F0: 1, F1: 2147483648, F2: 526}
        var v34 float32 = 1e10
        _ = v33
        _ = v34
    } else {
        var v35 bool = false
        v5 -= 255
        _ = v35
    }
    var v36 uint64 = 1
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 int32 = -1
        fmt.Println(v3, v36)
        if true {
            for fi39 := 0; fi39 < 3; fi39++ {
                switch v3 {
                case "\n\t":
                    v38 -= (^(v38))
                default:
                    v38 = (v38 * -745)
                }
                if true {
                    if false {
                        if true {
                            var v40 uint8 = v5
                            if true {
                                v41 := GenS1{F0: 42, F1: v6, F2: 4294967295}
                                v4[genFn1((^(v6)), (v36 % ((v36) | 1)), 18446744073709551615)] = v3
                                _ = v41
                            }
                            var v42 int = 634
                            var v43 uint16 = 345
                            var v44 float64 = 1.5
                            _ = v40
                            _ = v42
                            _ = v43
                            _ = v44
                        } else {
                            v38 = (v38 - v38)
                        }
                        v45 := GenS4{F0: 0}
                        v46 := GenS3{F0: -128, F1: 32767, F2: uint32(v36)}
                        _ = v45
                        _ = v46
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch v3 {
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = (-(fi37))
                            }
                        case "foo bar":
                            v3 = v3
                        case "\n\t":
                            v36 = (v36 * 944)
                        default:
                            var v47 bool = true
                            _ = v47
                        }
                    }
                    if true {
                        v48 := GenS1{F0: 42, F1: v6, F2: 0}
                        _ = v48
                    } else {
                        var v49 string = v3
                        _ = v49
                    }
                    v3 = v3
                }
                _ = fi39
            }
        }
        _ = fi37
        _ = v38
    }
    for fi50 := 0; fi50 < 4; fi50++ {
        var v51 float64 = float64((3.14159 * 0.0))
        _ = fi50
        _ = v51
    }
    _ = len(v4)
    _ = len(v4)
    switch -9223372036854775808 {
    case 42:
        for fi52 := 0; fi52 < 19; fi52++ {
            delete(v4, int(v5))
            _ = fi52
        }
    case 127:
        var v53 int64 = -1
        _ = v53
    case 128:
        var v54 float32 = -2.25
        _ = v54
    default:
        delete(v4, 4294967295)
    }
    if false {
        v55 := GenS4{F0: 1}
        v36 += (^(v36))
        delete(v4, 32767)
        _ = v55
    }
    fmt.Println(v36, v5, v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
