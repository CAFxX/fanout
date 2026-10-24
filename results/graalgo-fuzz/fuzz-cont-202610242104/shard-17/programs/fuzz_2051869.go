package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int8, p1 uint64, p2 int16) bool {
    switch ("\n\t" + "") {
    case "\n\t":
        if (42 >= 42) {
            switch -2147483648 {
            case -25:
                v7 := GenS1{F0: 149}
                _ = v7
            case 42:
                v8 := []int{32768, int(p2), -951}
                _ = v8
            default:
                var v9 uint8 = 2
                _ = v9
            }
        } else {
            p1 -= (p1 + p1)
        }
    case "hello":
        fmt.Println(p1, p2, p0)
    default:
        fmt.Println(p2)
    }
    v10 := []int{42, 848, -539, 32767}
    p2 = -450
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return true
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    var v4 int32 = -247
    v5 := []uint32{}
    var v6 int8 = 1
    v4 = -228
    var v11 float64 = 3.14159
    if (!false) {
        for fi12 := 0; fi12 < 9; fi12++ {
            v1 = v1
            _ = fi12
        }
        if (v6 == (v6 / (((v6 / ((v6) | 1))) | 1))) {
            var v13 uint8 = 1
            var v14 bool = true
            _ = v13
            _ = v14
        }
        var v15 uint32 = 2
        _ = v15
    }
    switch 42 {
    case 32767:
        if false {
            var v16 bool = true
            if (898 < (v1 ^ (v1 + v1))) {
                var v17 float64 = 1e10
                if len(v5) > 0 {
                    v5[0] = 360
                }
                var v18 float32 = 0.1
                _ = v17
                _ = v18
            }
            if v16 {
                switch 1 {
                case -2147483648:
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v16 = false
                        switch "" {
                        case "hello":
                            _ = len(v2)
                        case "":
                            if true {
                                v4 += (v4 << ((622) & 63))
                            } else {
                                var v20 float32 = 0.1
                                var v21 uint64 = 4294967295
                                _ = v20
                                _ = v21
                            }
                        }
                        _ = fi19
                    }
                case 127:
                    var v22 uint = 652
                    _ = v22
                default:
                    v23 := GenS1{F0: (-(v4))}
                    _ = v23
                }
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 uint16 = 42
                    _ = fi24
                    _ = v25
                }
                v4 += v4
            } else {
                var v26 uint16 = 181
                var v27 bool = v16
                _ = v26
                _ = v27
            }
            _ = v16
        }
    case -1:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 32768:
        var v28 uint64 = (v1 & 1)
        _ = v28
    default:
        var v29 int64 = -816
        _ = v29
    }
    var v30 int8 = (^(v6))
    v31 := GenS1{F0: (v4 << ((423) & 63))}
    fmt.Println(v30)
    v11 = -2.25
    if ((18446744073709551615 == 255) && (-2.25 > -2.25)) {
        var v32 int64 = -9223372036854775808
        for fi33 := 0; fi33 < 6; fi33++ {
            if (42 <= 128) {
                fmt.Println(v11)
                for fi34 := 0; fi34 < 18; fi34++ {
                    var v35 uint8 = 42
                    fmt.Println(fi33)
                    v30 += v30
                    _ = fi34
                    _ = v35
                }
            }
            _ = fi33
        }
        _ = v32
    } else {
        if (750 > 127) {
            if false {
                v6 -= (^(v30))
                v4 = (-(v4))
                v36 := v3[2147483647]
                _ = v36
            }
            v37 := GenS2{F0: 244}
            fmt.Println(v4, v1)
            _ = v37
        }
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            var v38 int = -428
            _ = v38
        case "hello":
            switch -893 {
            case 128:
                var v39 int = -140
                _ = v39
            case -392:
                switch "foo bar" {
                case "foo bar":
                    var v40 uint = 4294967295
                    _ = v40
                case "":
                    v41 := GenS3{F0: (("\n\t" + "a") + "hello")}
                    _ = v41
                case "\n\t":
                    var v42 float64 = v11
                    _ = v42
                }
            default:
                v30 -= v30
            }
        }
        v43 := GenS4{F0: (674 == 131), F1: 4294967295}
        _ = v43
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        v30 = int8(fi44)
        v45 := GenS5{F0: 65, F1: 2}
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        for fi47 := 0; fi47 < 16; fi47++ {
            v1 += v1
            var v48 int32 = (v4 * v4)
            v3[2] = 65535
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    fmt.Println(v6)
    v30 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
