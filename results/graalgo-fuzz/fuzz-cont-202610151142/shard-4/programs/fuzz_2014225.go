package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 uint8) uint16 {
    switch "hello" {
    case "":
        var v6 int32 = 42
        _ = v6
    case "foo bar":
        var v7 float64 = float64(0.1)
        _ = v7
    case "hello":
        if (false && true) {
            p0 += 2
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := []int{fi8, fi8, fi8}
                var v10 int32 = 727
                _ = fi8
                _ = v10
                _ = v9
            }
            if false {
                p0 = p0
                var v11 uint8 = p0
                _ = v11
            }
        }
    default:
        if ((p0 ^ p0) < p0) {
            var v12 int32 = 892
            _ = v12
        } else {
            if (uint16(p0) > 255) {
                var v13 int16 = 0
                _ = v13
            }
        }
    }
    switch -9223372036854775808 {
    case 42:
        p0 += p0
    case 32768:
        var v14 float64 = (-2.25 - 3.14159)
        _ = v14
    case 127:
        p0 = (p0 ^ (^(p0)))
    default:
        var v15 uint = 611
        _ = v15
    }
    p0 = 1
    if false {
        v16 := []int{}
        p0 = (p0 ^ p0)
        for fi17 := 0; fi17 < 8; fi17++ {
            if (false || (100 > -90)) {
                v18 := GenS1{F0: true, F1: 0, F2: 2147483647}
                p0 = p0
                _ = v18
            }
            _ = fi17
        }
        _ = v16
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 float32 = ((1e10 - 1e10) / ((-2.25 + 1e10) + 0.0))
            v20 -= 1e10
            if false {
                if true {
                    var v21 float64 = (1e10 - (3.14159 / 1.5))
                    for fi22 := 0; fi22 < 3; fi22++ {
                        switch fi19 {
                        case -376:
                            if ((v21 - v21) < float64((v20 + (-2.25 - 1e10)))) {
                                v20 += (v20 - v20)
                            }
                        default:
                            v23 := []int{-1, (fi19 >> ((2) & 63))}
                            _ = v23
                        }
                        v21 = v21
                        var v24 int8 = -87
                        _ = fi22
                        _ = v24
                    }
                    v20 = v20
                    _ = v21
                }
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 int8 = 127
                    v27 := GenS1{F0: false, F1: -9223372036854775808, F2: 42}
                    v28 := GenS2{F0: -1}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            }
            _ = fi19
            _ = v20
        }
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 0
    var v5 string = ""
    v4 -= (v4 << ((908) & 63))
    var v29 uint = 12
    for fi30 := 0; fi30 < 12; fi30++ {
        v5 = ("\n\t" + v5)
        fmt.Println(v29, v29)
        _ = fi30
    }
    v31 := GenS3{F0: 4294967295, F1: (true || (true || (false || (!true))))}
    switch 127 {
    case 128:
        var v32 uint64 = 18446744073709551615
        _ = v32
    case 55:
        if (!(true || true)) {
            _ = len(v2)
            if false {
                for fi33 := 0; fi33 < 13; fi33++ {
                    fmt.Println(v29, v5, fi33)
                    _ = fi33
                }
            } else {
                v34 := v3[(^(v4))]
                v29 = v29
                _ = v34
            }
            v35 := GenS2{F0: -9223372036854775808}
            _ = v35
        }
    }
    v5 = v5
    v36 := GenS2{F0: -792}
    switch v5 {
    case "a":
        v5 = v5
    case "":
        fmt.Println(v29, v29, v5)
    }
    var v37 int8 = -74
    v38 := GenS1{F0: false, F1: 0, F2: 1}
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 9; fi40++ {
            v4 += v4
            fmt.Println(fi40, v5)
            if ((false || true) && true) {
                var v41 uint = 42
                _ = v41
            }
            _ = fi40
        }
        switch v5 {
        case "a":
            switch "foo bar" {
            case "":
                var v42 uint8 = 127
                _ = v42
            default:
                for fi43 := 0; fi43 < 10; fi43++ {
                    var v44 int8 = v37
                    _ = fi43
                    _ = v44
                }
            }
        case "":
            v29 = v29
        }
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
