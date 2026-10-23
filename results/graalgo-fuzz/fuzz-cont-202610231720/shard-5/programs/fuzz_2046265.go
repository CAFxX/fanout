package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int16, p1 int8) int8 {
    fmt.Println(p1)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return (p1 << ((705) & 63))
}
func genFn1(p0 string, p1 uint64, p2 int) uint64 {
    v8 := []int{4294967295, p2, (p2 & (p2 & p2))}
    for fi9 := 0; fi9 < 12; fi9++ {
        if false {
            if (((false || (14 > 744)) && (true || (false || (138 < 65535)))) && false) {
                fmt.Println(p1, p2, p2)
                var v10 float32 = (0.1 - 0.1)
                _ = v10
            }
            var v11 float64 = -2.25
            fmt.Println(p0)
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS1{F0: -2147483648, F1: 462, F2: 42}
    v13 := GenS1{F0: 125, F1: -168, F2: 494}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v8
    return (^(p1))
}

func main() {
    var v1 int64 = -1
    var v2 float64 = 0.0
    var v3 uint16 = 32767
    v4 := []int{42, -9223372036854775808, -1}
    var v5 uint = 18446744073709551615
    v6 := make(map[uint32]bool)
    var v14 float64 = (0.1 * 1.5)
    v4 = append(v4, -2147483648)
    v2 = v2
    if false {
        if ((-2.25 - ((1.5 - 0.0) - -2.25)) < 1e10) {
            v3 += (v3 / ((255) | 1))
            v15 := GenS1{F0: 2147483647, F1: -728, F2: (v3 / ((v3) | 1))}
            v16 := v6[42]
            _ = v15
            _ = v16
        }
        _ = len(v4)
    } else {
        v3 = (v3 % ((65535) | 1))
        v3 -= v3
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS2{F0: 588}
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS2{F0: 497}
                v1 += 656
                var v21 uint = v5
                if (-630 <= -748) {
                    if true {
                        var v22 int32 = -750
                        if true {
                            switch ("hello" + "") {
                            case "a":
                                v2 = v14
                            case "\n\t":
                                var v23 uint32 = 0
                                v21 = 871
                                _ = v23
                            case "foo bar":
                                var v24 int32 = (^(v22))
                                _ = v24
                            }
                            v4 = append(v4, int(v1))
                        }
                        _ = v22
                    }
                    switch (fi17 >> ((v21) & 63)) {
                    case -1:
                        v25 := GenS3{F0: 1, F1: 1}
                        _ = v25
                    }
                    v26 := GenS1{F0: -2147483648, F1: 1, F2: (v3 ^ v3)}
                    _ = v26
                } else {
                    switch -2147483648 {
                    case 32768:
                        var v27 uint32 = 4294967295
                        _ = v27
                    case 9223372036854775807:
                        delete(v6, 65535)
                    }
                    v4 = append(v4, (fi19 % (((fi19 * -1)) | 1)))
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi17
            _ = v18
        }
    }
    v2 = ((0.0 * v2) / v2)
    switch 128 {
    case -677:
        v28 := v6[4294967295]
        _ = v28
    case 128:
        v1 = -95
    case -9223372036854775808:
        var v29 uint64 = 4294967295
        _ = v29
    }
    v30 := GenS2{F0: uint32(v5)}
    v1 += v1
    if (!((v3 << (((v5 >> (((v5 & (v5 ^ 770))) & 63))) & 63)) < 65535)) {
        var v31 uint16 = v3
        var v32 uint32 = 2147483647
        _ = v31
        _ = v32
    }
    _ = len(v6)
    var v33 float64 = (0.0 / ((v14 - v14) + v2))
    for fi34 := 0; fi34 < 19; fi34++ {
        if len(v4) > 0 {
            v4[0] = 119
        }
        fmt.Println(v1, v14)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi34
    }
    var v35 int = int(v1)
    var v36 uint = 763
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
