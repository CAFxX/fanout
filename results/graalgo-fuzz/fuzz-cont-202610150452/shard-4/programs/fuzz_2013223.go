package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int64; F1 int64; F2 string }
func genFn0() bool {
    var v6 int64 = 965
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: 20, F1: 255, F2: 2}
            _ = fi7
            _ = v8
        }
        v6 = v6
        v9 := GenS1{F0: 535, F1: 128, F2: -1}
        _ = v9
    }
    var v10 int8 = 85
    _ = v10
    _ = v6
    return true
}
func genFn1(p0 uint16, p1 uint64, p2 uint64) uint {
    p0 += p0
    v11 := []int{9223372036854775807, -2147483648, -893, 128}
    var v12 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 868
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 float64 = 1e10
    var v3 string = "foo bar"
    v4 := []uint32{65535}
    v5 := make(map[int]int)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 uint16 = 0
        v15 := GenS1{F0: 832, F1: 0, F2: 71}
        _ = fi13
        _ = v14
        _ = v15
    }
    v1 += (v1 + v1)
    v3 = (v3 + ((v3 + v3) + v3))
    for fi16 := 0; fi16 < 6; fi16++ {
        v1 -= 700
        switch "a" {
        case "\n\t":
            var v17 uint16 = 32767
            _ = v17
        case "foo bar":
            fmt.Println(v2, v1)
        default:
            _ = len(v5)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi16
    }
    var v18 float64 = ((1.5 - (v2 + ((v2 / -2.25) - float64((-2.25 + 1e10))))) * 0.1)
    v19 := GenS2{F0: 2, F1: 2}
    v20 := GenS1{F0: 1000, F1: 255, F2: 127}
    var v21 uint8 = 23
    var v22 uint64 = 711
    v23 := GenS3{F0: (v3 + "\n\t"), F1: 42}
    switch 127 {
    case 32768:
        switch v3 {
        case "foo bar":
            var v24 int8 = 42
            _ = v24
        case "hello":
            var v25 uint32 = 2147483647
            _ = v25
        case "":
            if (v3 < "foo bar") {
                if true {
                    var v26 int8 = 100
                    _ = v26
                } else {
                    v18 += v2
                }
                v3 = v3
                var v27 string = ("hello" + v3)
                _ = v27
            } else {
                if ((((-539 <= -2147483648) && (2147483648 > uint32(v22))) && true) || (!(691 > uint16(v21)))) {
                    var v28 int8 = -84
                    _ = v28
                } else {
                    v22 += 307
                    _ = len(v5)
                }
                var v29 uint16 = 65535
                if false {
                    var v30 int32 = 42
                    switch 4294967295 {
                    case 127:
                        v31 := GenS2{F0: 0, F1: (v29 ^ (v29 & 25))}
                        _ = v31
                    case 32767:
                        for fi32 := 0; fi32 < 13; fi32++ {
                            var v33 uint64 = 42
                            _ = fi32
                            _ = v33
                        }
                    default:
                        var v34 uint8 = 0
                        _ = v34
                    }
                    _ = v30
                }
                _ = v29
            }
        default:
            switch v3 {
            case "foo bar":
                var v35 int = -1
                _ = v35
            case "\n\t":
                v36 := GenS4{F0: 1, F1: 42, F2: (("hello" + "") + v3)}
                _ = v36
            }
        }
    case 602:
        if ((!false) && true) {
            var v37 int8 = 77
            v4 = append(v4, 488)
            if (!false) {
                v38 := v5[-9223372036854775808]
                v39 := v5[1]
                _ = v38
                _ = v39
            }
            _ = v37
        }
    case 611:
        var v40 int64 = -9223372036854775808
        _ = v40
    }
    switch -1 {
    case 42:
        for fi41 := 0; fi41 < 18; fi41++ {
            var v42 uint16 = 145
            var v43 int64 = -104
            _ = fi41
            _ = v42
            _ = v43
        }
    case 127:
        v4 = append(v4, 355)
    case -2147483648:
        delete(v5, 9223372036854775807)
    default:
        v1 = (v1 >> (((v1 ^ v1)) & 63))
    }
    var v44 int8 = 100
    fmt.Println(v21, v2, v22)
    var v45 uint64 = v22
    var v46 float32 = float32(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
