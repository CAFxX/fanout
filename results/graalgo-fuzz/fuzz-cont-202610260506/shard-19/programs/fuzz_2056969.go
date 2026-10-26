package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint64) uint16 {
    var v7 uint64 = p0
    v8 := []int{}
    if false {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                if (!true) {
                    fmt.Println(fi10, fi10, p0)
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                } else {
                    p0 = (^(p0))
                }
                _ = fi10
            }
            _ = fi9
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: fi11, F1: 1, F2: fi11}
            var v13 int = 32768
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        var v14 int8 = -128
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS1{F0: fi15, F1: 1, F2: -9223372036854775808}
            var v17 uint32 = 1
            fmt.Println(v17, v7)
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = len(v8)
        _ = v14
    }
    v18 := GenS1{F0: 0, F1: 10, F2: 127}
    _ = p0
    _ = v18
    _ = v7
    _ = v8
    return 642
}
func genFn1(p0 int32, p1 uint8) uint32 {
    v19 := GenS1{F0: -9223372036854775808, F1: 343, F2: 4294967295}
    v20 := []int{32767, 32767}
    if false {
        var v21 uint8 = (p1 % ((p1) | 1))
        var v22 uint32 = 1
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(fi23, p0, p0)
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi23
        }
        v24 := GenS1{F0: -2147483648, F1: 0, F2: -9223372036854775808}
        v20 = append(v20, 2147483647)
        _ = v24
    }
    if true {
        p1 = (p1 / (((p1 / ((p1) | 1))) | 1))
        p0 += p0
        switch 32768 {
        case 1:
            var v25 string = "foo bar"
            _ = v25
        default:
            p1 = 42
        }
    } else {
        switch 128 {
        case 926:
            v26 := GenS2{F0: 2147483648, F1: -32768, F2: (^(p1))}
            _ = v26
        case 825:
            p0 += (-(p0))
        default:
            p0 -= (p0 | (p0 / ((p0) | 1)))
        }
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 1
}

func main() {
    var v1 uint8 = 26
    var v2 int16 = -661
    v3 := make(map[uint32]uint32)
    var v4 int64 = -1
    var v5 uint32 = 353
    var v6 uint64 = uint64(v5)
    if ((v1 | (v1 | (^(v1)))) != (v1 * (v1 ^ 1))) {
        for fi27 := 0; fi27 < 14; fi27++ {
            delete(v3, v5)
            var v28 uint32 = v5
            _ = fi27
            _ = v28
        }
    } else {
        var v29 int16 = v2
        _ = v29
    }
    v30 := GenS3{F0: 1, F1: v5}
    if ((false && (32767 > 484)) || false) {
        var v31 float32 = 0.0
        _ = v31
    } else {
        var v32 int16 = (v2 | v2)
        var v33 int8 = -38
        for fi34 := 0; fi34 < 1; fi34++ {
            if false {
                v3[(v5 & genFn1(664, (v1 & 255)))] = (v5 % ((v5) | 1))
                fmt.Println(v1, v6)
            } else {
                var v35 string = "hello"
                var v36 bool = true
                var v37 int8 = v33
                _ = v35
                _ = v36
                _ = v37
            }
            v38 := []int{}
            _ = fi34
            _ = v38
        }
        _ = v32
        _ = v33
    }
    v39 := GenS1{F0: 2147483647, F1: (v4 + v4), F2: 429}
    if (((true && false) && false) && (-1 != 0)) {
        if false {
            v40 := GenS4{F0: 395}
            v2 += (v2 / ((v2) | 1))
            _ = v40
        } else {
            v6 += 4294967295
            fmt.Println(v6)
        }
    } else {
        var v41 uint64 = v6
        var v42 int8 = 100
        var v43 uint = 18446744073709551615
        _ = v41
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        for fi45 := 0; fi45 < 2; fi45++ {
            v3[(v5 ^ v5)] = v5
            v6 -= (v6 + v6)
            fmt.Println(v1)
            if false {
                v46 := v3[(v5 | 4294967295)]
                fmt.Println(v2, v1, v6)
                _ = v46
            } else {
                var v47 uint = 1
                var v48 uint16 = 395
                _ = v47
                _ = v48
            }
            _ = len(v3)
            _ = fi45
        }
        _ = fi44
    }
    if true {
        for fi49 := 0; fi49 < 9; fi49++ {
            _ = len(v3)
            _ = fi49
        }
        if false {
            v50 := GenS3{F0: (v2 % ((304) | 1)), F1: uint32(v6)}
            if false {
                v5 = v5
                _ = len(v3)
            } else {
                fmt.Println(v4, v4, v1)
                for fi51 := 0; fi51 < 6; fi51++ {
                    var v52 uint32 = v5
                    v53 := GenS2{F0: genFn1(-991, v1), F1: -1, F2: 183}
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                v2 -= (v2 ^ (v2 * (v2 & int16(v6))))
            }
            for fi54 := 0; fi54 < 10; fi54++ {
                for fi55 := 0; fi55 < 15; fi55++ {
                    v56 := GenS4{F0: (v6 - (v6 % (((v6 >> ((255) & 63))) | 1)))}
                    _ = fi55
                    _ = v56
                }
                _ = fi54
            }
            _ = v50
        } else {
            v5 = genFn1(385, v1)
            for fi57 := 0; fi57 < 16; fi57++ {
                if ((^(v6)) > uint64(v5)) {
                    var v58 int64 = v4
                    _ = v58
                } else {
                    v59 := GenS1{F0: (fi57 >> ((260) & 63)), F1: (v4 << ((788) & 63)), F2: (fi57 * fi57)}
                    _ = v59
                }
                _ = fi57
            }
            var v60 uint16 = genFn0(v6)
            _ = v60
        }
    } else {
        var v61 int16 = -155
        _ = v61
    }
    v62 := GenS2{F0: (v5 << ((4294967295) & 63)), F1: (v2 ^ v2), F2: uint8(v5)}
    v5 = (v5 * 120)
    for fi63 := 0; fi63 < 11; fi63++ {
        v4 -= v4
        _ = fi63
    }
    v4 = (v4 ^ (v4 - 42))
    var v64 int = -2147483648
    v64 -= -718
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v64
}
