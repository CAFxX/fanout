package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint64) uint16 {
    v7 := GenS1{F0: 0, F1: 42, F2: (627 > 128)}
    if false {
        var v8 int = -2147483648
        v9 := []int{v8, v8, v8, (^(v8))}
        fmt.Println(p0)
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{(fi10 | fi10), 52, fi10}
            var v12 int16 = 1
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = v7
    return 65535
}

func main() {
    v1 := []string{}
    var v2 uint32 = 1
    var v3 float64 = float64(3.14159)
    var v4 uint16 = 1
    v5 := make(map[string]uint32)
    var v6 uint64 = 487
    if true {
        v13 := GenS1{F0: (v4 ^ v4), F1: -328, F2: (213 == 13)}
        var v14 uint8 = 42
        _ = v13
        _ = v14
    } else {
        fmt.Println(v2, v3, v3)
    }
    v6 -= 4294967295
    if true {
        v15 := v5["\n\t"]
        fmt.Println(v15, v3)
        _ = v15
    }
    var v16 string = "hello"
    v17 := GenS2{F0: 2, F1: 0, F2: ("hello" + v16)}
    var v18 int8 = 77
    if (-127 == -1) {
        if (int32(v4) == -2147483648) {
            var v19 float64 = (v3 * 0.0)
            _ = v19
        } else {
            if (!true) {
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 uint32 = 1
                    _ = fi20
                    _ = v21
                }
                v22 := GenS2{F0: v2, F1: 1, F2: "foo bar"}
                _ = v22
            } else {
                v6 -= uint64(v4)
            }
        }
        fmt.Println(v3, v6, v18)
    } else {
        if false {
            fmt.Println(v16, v3, v2)
        }
        v6 = v6
        v23 := GenS2{F0: 1, F1: 42, F2: (v16 + v16)}
        _ = v23
    }
    v1 = append(v1, v16)
    _ = len(v5)
    if (255 != 144) {
        var v24 bool = false
        for fi25 := 0; fi25 < 8; fi25++ {
            delete(v5, v16)
            v26 := GenS1{F0: uint16(v2), F1: -1, F2: v24}
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        if false {
            if (false || false) {
                switch 32768 {
                case 835:
                    var v27 uint16 = v4
                    _ = v27
                case -1:
                    v6 += 774
                default:
                    v6 += 0
                }
                v6 += 1
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                _ = len(v5)
                v29 := GenS1{F0: v4, F1: 42, F2: false}
                var v30 int = (fi28 + fi28)
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            v6 -= (v6 | v6)
            v31 := GenS1{F0: (v4 << ((4294967295) & 63)), F1: 623, F2: true}
            v3 = v3
            _ = v31
        }
        fmt.Println(v16, v18, v4)
        switch v16 {
        case "hello":
            var v32 bool = false
            _ = v32
        case "a":
            switch "a" {
            case "foo bar":
                for fi33 := 0; fi33 < 7; fi33++ {
                    var v34 float32 = 1e10
                    for fi35 := 0; fi35 < 13; fi35++ {
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        fmt.Println(v2, v4)
                        v4 += 1
                        _ = fi35
                    }
                    _ = fi33
                    _ = v34
                }
            case "hello":
                v36 := GenS1{F0: (v4 | (v4 % ((265) | 1))), F1: 0, F2: false}
                _ = v36
            default:
                switch 127 {
                case 127:
                    for fi37 := 0; fi37 < 17; fi37++ {
                        for fi38 := 0; fi38 < 15; fi38++ {
                            v39 := GenS2{F0: (v2 & v2), F1: -1, F2: v16}
                            _ = fi38
                            _ = v39
                        }
                        if (true || true) {
                            if (!(-1 != (fi37 + -2147483648))) {
                                if len(v1) > 0 {
                                    v1[0] = (v16 + v16)
                                }
                            } else {
                                v18 = -77
                                v3 = v3
                                fmt.Println(fi37, v6)
                            }
                            fmt.Println(v16, fi37)
                        }
                        _ = fi37
                    }
                case 32768:
                    v16 = "a"
                }
            }
        default:
            v4 = uint16(v6)
        }
    }
    v40 := v5[(("foo bar" + v16) + "a")]
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v41 uint = 472
    v42 := GenS1{F0: 32767, F1: 70, F2: true}
    var v43 uint16 = (v4 % ((v4) | 1))
    var v44 uint64 = 18446744073709551615
    switch 2147483647 {
    case 9223372036854775807:
        v1 = append(v1, v16)
    case 4294967295:
        v45 := GenS2{F0: (v2 % (((v2 * 2)) | 1)), F1: -9223372036854775808, F2: "foo bar"}
        _ = v45
    case 288:
        if (1.5 > float32(((v3 * float64(1.5)) / float64(3.14159)))) {
            fmt.Println(v16, v41, v18)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        for fi46 := 0; fi46 < 6; fi46++ {
            v47 := GenS3{F0: v4, F1: 157}
            for fi48 := 0; fi48 < 3; fi48++ {
                fmt.Println(v18, v43)
                _ = fi48
            }
            if false {
                for fi49 := 0; fi49 < 20; fi49++ {
                    v50 := GenS4{F0: 42}
                    _ = fi49
                    _ = v50
                }
                var v51 bool = (42 != v41)
                _ = len(v5)
                _ = v51
            }
            _ = fi46
            _ = v47
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
