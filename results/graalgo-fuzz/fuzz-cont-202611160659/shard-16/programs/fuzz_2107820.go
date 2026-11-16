package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 uint64) uint32 {
    var v6 float64 = 1e10
    p0 = 0
    if true {
        v7 := GenS1{F0: 0, F1: 18446744073709551615, F2: -876}
        _ = v7
    } else {
        switch p0 {
        case 127:
            p0 += 127
        case 52:
            p1 = p1
        default:
            switch ("a" + "foo bar") {
            case "":
                p1 = (p1 / ((10) | 1))
            default:
                if false {
                    var v8 int = 2147483647
                    if (((((v6 + v6) - (-2.25 / v6)) - (0.0 + 1.5)) >= v6) && true) {
                        var v9 int64 = 9223372036854775807
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 4; fi10++ {
                        var v11 uint64 = 18446744073709551615
                        _ = fi10
                        _ = v11
                    }
                    _ = v8
                }
            }
        }
    }
    var v12 uint16 = 0
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            if false {
                v15 := []int{(fi14 >> ((1) & 63)), (fi14 % (((^(fi13))) | 1)), 127, -114}
                _ = v15
            }
            fmt.Println(v6)
            _ = fi14
        }
        if false {
            p0 = (^(fi13))
        }
        fmt.Println(p0, v12)
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := []uint32{493, 488, 65535, 2147483648}
    var v4 bool = ((!(false || false)) && false)
    var v5 int32 = -550
    if true {
        var v16 uint64 = 0
        for fi17 := 0; fi17 < 11; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                delete(v1, 465)
                v19 := GenS2{F0: "\n\t", F1: v5}
                var v20 bool = true
                _ = fi18
                _ = v19
                _ = v20
            }
            if false {
                var v21 int = -1
                _ = v21
            }
            _ = fi17
        }
        v4 = v4
        _ = v16
    }
    if true {
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 uint = 4294967295
            _ = fi22
            _ = v23
        }
    } else {
        v24 := GenS2{F0: "hello", F1: (v5 % ((-1) | 1))}
        v4 = (1 > -682)
        _ = v24
    }
    delete(v1, 42)
    if (-1 <= v5) {
        for fi25 := 0; fi25 < 16; fi25++ {
            if (42 < v5) {
                if (false && (!(1 > 9))) {
                    for fi26 := 0; fi26 < 14; fi26++ {
                        v3 = append(v3, genFn0(32767, 744))
                        _ = fi26
                    }
                }
            } else {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 float32 = 1e10
                    v5 -= v5
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi27
                    _ = v28
                }
                v4 = ((v5 | (v5 * (v5 | 1))) != v5)
            }
            v29 := GenS3{F0: 2, F1: 1}
            var v30 uint32 = 953
            _ = fi25
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v32 := GenS4{F0: -1, F1: fi31}
        for fi33 := 0; fi33 < 10; fi33++ {
            fmt.Println(v4, v5, fi31)
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    var v34 int16 = int16(v5)
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 uint32 = 2147483648
        v5 = int32(v36)
        fmt.Println(v5, fi35, v34)
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 uint8 = 185
        if v4 {
            var v39 int = fi37
            fmt.Println(v38, v39)
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 298)
    if v4 {
        delete(v1, 0)
    } else {
        var v40 float32 = ((0.1 - 1e10) / -2.25)
        for fi41 := 0; fi41 < 6; fi41++ {
            var v42 uint8 = 3
            _ = len(v3)
            var v43 uint64 = 0
            _ = fi41
            _ = v42
            _ = v43
        }
        v40 -= (((((v40 * ((((1e10 + (v40 * 1.5)) + 1e10) / v40) * v40)) / v40) - 1.5) * float32(1.5)) - (v40 * float32(0.0)))
        _ = v40
    }
    if (false && ((v5 ^ (v5 >> ((255) & 63))) == v5)) {
        var v44 uint64 = 557
        v5 = 2147483647
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
