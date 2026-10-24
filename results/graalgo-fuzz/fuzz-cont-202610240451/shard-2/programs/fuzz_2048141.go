package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
type GenS4 struct { F0 int32; F1 uint64; F2 int32 }
func genFn0(p0 uint64, p1 uint16, p2 bool) uint32 {
    var v7 int32 = 704
    p1 = (p1 % ((p1) | 1))
    v8 := GenS1{F0: 1, F1: 1, F2: -118}
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 -= 42
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int32, p1 string, p2 bool) uint {
    if (p2 && true) {
        var v10 uint8 = 127
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(fi11, v10)
            _ = fi11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            p1 = p1
            if p2 {
                var v13 uint16 = 32767
                _ = v13
            } else {
                p1 = "a"
                p0 -= 409
                fmt.Println(v10, v10, fi12)
            }
            _ = fi12
        }
        _ = v10
    } else {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 float32 = (0.1 / (3.14159 - 0.1))
            _ = fi14
            _ = v15
        }
        if p2 {
            if ((!p2) || false) {
                v16 := []int{}
                v17 := GenS1{F0: 2147483648, F1: 1, F2: -1}
                if (255 != 2) {
                    if (-46 != -436) {
                        p1 = p1
                        p1 = p1
                        p2 = ((42 != 128) || (128 >= -2147483648))
                    } else {
                        var v18 uint8 = 42
                        for fi19 := 0; fi19 < 8; fi19++ {
                            if p2 {
                                p0 += p0
                                p2 = p2
                            } else {
                                var v20 int64 = -1
                                var v21 uint8 = 1
                                _ = v20
                                _ = v21
                            }
                            var v22 int64 = 0
                            fmt.Println(v18, v22)
                            _ = fi19
                            _ = v22
                        }
                        fmt.Println(v18, p2)
                        _ = v18
                    }
                    v23 := GenS2{F0: p1}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 18; fi24++ {
                        for fi25 := 0; fi25 < 13; fi25++ {
                            var v26 uint32 = 2
                            var v27 bool = true
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        fmt.Println(p1, p2)
                        _ = fi24
                    }
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 int8 = 0
                        v30 := GenS3{F0: 1, F1: p2, F2: 91}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                _ = v16
                _ = v17
            }
        }
    }
    var v31 bool = ((p0 <= p0) && (((1e10 / -2.25) > 0.0) || p2))
    var v32 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v32
    return 991
}

func main() {
    var v1 float32 = 0.1
    var v2 int = 32768
    var v3 uint8 = 221
    var v4 float32 = -2.25
    v5 := []uint32{65535, 520}
    var v6 int64 = 0
    v2 = v2
    fmt.Println(v3, v3, v4)
    if true {
        var v33 uint8 = v3
        if true {
            v3 = v33
        }
        _ = v33
    }
    v3 = v3
    v34 := GenS2{F0: ""}
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := GenS4{F0: 539, F1: 557, F2: 42}
        if len(v5) > 0 {
            v5[0] = 460
        }
        _ = fi35
        _ = v36
    }
    fmt.Println(v1, v3, v1)
    if (731 >= 65535) {
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 int8 = 0
            _ = fi37
            _ = v38
        }
        var v39 uint32 = 443
        switch 429 {
        case 127:
            fmt.Println(v4, v1)
        case -981:
            fmt.Println(v6, v6, v39)
        case 128:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            switch (v2 + (-(v2))) {
            case 9223372036854775807:
                var v40 uint32 = v39
                _ = v40
            case 32768:
                var v41 int = 2147483647
                _ = v41
            case 212:
                v39 = v39
            default:
                _ = len(v5)
            }
        }
        _ = v39
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        if true {
            switch "a" {
            case "hello":
                v1 = v4
            default:
                var v43 uint8 = v3
                _ = v43
            }
        }
        for fi44 := 0; fi44 < 10; fi44++ {
            var v45 uint32 = 4294967295
            v4 -= (v4 * v4)
            var v46 float32 = v1
            switch -9223372036854775808 {
            case 102:
                v47 := GenS3{F0: 32767, F1: ((fi44 >> ((2) & 63)) > -2147483648), F2: -1}
                _ = v47
            case 9223372036854775807:
                v2 = (-(v2))
            }
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi42
    }
    if len(v5) > 0 {
        v5[0] = genFn0(42, 1, (!((448 > genFn1(1, (("" + "") + "hello"), (true || true))) || true)))
    }
    v6 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
