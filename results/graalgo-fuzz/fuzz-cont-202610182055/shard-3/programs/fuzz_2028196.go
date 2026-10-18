package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int16; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint64) int64 {
    var v6 uint = 509
    p0 += (p0 / (((^(p0))) | 1))
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 uint8 = 40
    var v2 uint = 4294967295
    var v3 uint16 = 32767
    var v4 int = 4294967295
    v5 := make(map[uint32]uint32)
    fmt.Println(v2, v3, v2)
    for fi7 := 0; fi7 < 3; fi7++ {
        if ((v1 | 98) > v1) {
            var v8 uint32 = 1
            _ = v8
        } else {
            switch fi7 {
            case 32767:
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS1{F0: -386, F1: "foo bar", F2: (fi7 ^ -1)}
                    var v11 uint64 = 2
                    v1 = 255
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                var v12 uint16 = (v3 & 32767)
                _ = v12
            }
        }
        delete(v5, 2147483648)
        _ = fi7
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        delete(v5, 65535)
        v3 += 2
        v4 -= (v4 % (((v4 & 32767)) | 1))
        _ = fi13
    }
    v3 = v3
    var v14 string = "hello"
    if (false && (9223372036854775807 < -816)) {
        v15 := GenS2{F0: v14}
        v2 -= v2
        var v16 float32 = float32((0.1 * 1e10))
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: 32767, F1: "foo bar", F2: (v4 % ((-915) | 1))}
        for fi18 := 0; fi18 < 5; fi18++ {
            if (!(uint(v3) <= (v2 >> ((v2) & 63)))) {
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS3{F0: 1, F1: 0, F2: (v3 % (((v3 | v3)) | 1))}
                    fmt.Println(fi19, fi18)
                    switch (v14 + "") {
                    case "\n\t":
                        switch "foo bar" {
                        case "":
                            var v21 int = fi18
                            _ = v21
                        }
                    case "":
                        if false {
                            v22 := []int{(fi19 | 127), 32767, (fi18 << ((v2) & 63))}
                            _ = v22
                        }
                    case "a":
                        v23 := GenS4{F0: 100}
                        _ = v23
                    default:
                        var v24 int64 = -9223372036854775808
                        var v25 uint = (^(v2))
                        _ = v24
                        _ = v25
                    }
                    _ = fi19
                    _ = v20
                }
                var v26 float32 = -2.25
                _ = v26
            }
            _ = fi18
        }
        _ = v17
    }
    v5[567] = 2
    if false {
        var v27 uint = (v2 * (v2 + v2))
        if true {
            var v28 uint16 = v3
            v29 := []int{(v4 - 2147483647), (-(v4))}
            v28 -= (v28 | (v28 << (((v2 & v27)) & 63)))
            _ = v28
            _ = v29
        }
        _ = v27
    }
    if (9223372036854775807 < genFn0(2)) {
        switch (v4 ^ v4) {
        case 127:
            var v30 bool = (((("foo bar" + "a") + v14) + v14) > v14)
            _ = v30
        case -480:
            v31 := GenS5{F0: 950, F1: -561}
            _ = v31
        case -1:
            switch "" {
            case "":
                v32 := GenS4{F0: -40}
                _ = v32
            case "hello":
                for fi33 := 0; fi33 < 3; fi33++ {
                    v14 = (v14 + "hello")
                    v5[4294967295] = 0
                    _ = fi33
                }
            case "\n\t":
                var v34 int32 = 410
                _ = v34
            }
        default:
            v2 = uint(v1)
        }
        fmt.Println(v4)
        _ = len(v5)
    }
    if true {
        v2 -= (v2 - (v2 | v2))
        var v35 int8 = -1
        v5[608] = 640
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v4 -= v4
        var v37 uint64 = 0
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
