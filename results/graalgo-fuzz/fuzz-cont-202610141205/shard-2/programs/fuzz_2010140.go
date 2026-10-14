package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int16; F2 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8) bool {
    switch 1 {
    case 128:
        for fi7 := 0; fi7 < 13; fi7++ {
            switch 1 {
            case 1:
                p0 -= (p0 | p0)
            case 128:
                var v8 int16 = 387
                _ = v8
            default:
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: 42, F1: 2147483647}
                    var v11 string = "hello"
                    v12 := GenS1{F0: 595, F1: 2147483647}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
            _ = fi7
        }
    case 1:
        var v13 uint32 = 432
        _ = v13
    case -108:
        var v14 int = 32767
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int8 = -48
            _ = fi15
            _ = v16
        }
        v17 := []int{}
        _ = v17
    } else {
        if (-702 < -1) {
            if ((42 <= 1) && (false || (false || ((p0 - p0) > 42)))) {
                p0 = 157
            } else {
                var v18 int32 = 42
                var v19 uint8 = p0
                _ = v18
                _ = v19
            }
            p0 = (p0 << ((2) & 63))
        }
        var v20 int16 = int16(p0)
        _ = v20
    }
    var v21 int16 = 1000
    v22 := GenS2{F0: 42}
    _ = p0
    _ = v21
    _ = v22
    return false
}

func main() {
    var v1 float64 = ((float64((-2.25 / 1.5)) / float64(1e10)) * -2.25)
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483647
    v4 := []string{}
    v5 := make(map[int]bool)
    var v6 float64 = v1
    var v23 int64 = 1
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 uint8 = 255
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    switch -2147483648 {
    case 128:
        v4 = append(v4, "a")
    case 9223372036854775807:
        var v27 uint = 967
        _ = v27
    case 32767:
        var v28 int8 = 127
        _ = v28
    default:
        for fi29 := 0; fi29 < 8; fi29++ {
            v6 = 3.14159
            for fi30 := 0; fi30 < 1; fi30++ {
                switch int(v3) {
                case -483:
                    var v31 float64 = v1
                    _ = v31
                case 1:
                    fmt.Println(fi30, fi30)
                case 32767:
                    var v32 uint16 = 399
                    _ = v32
                }
                _ = fi30
            }
            _ = fi29
        }
    }
    v1 += v6
    v33 := GenS3{F0: 782, F1: -551, F2: (v3 ^ (v3 % ((v3) | 1)))}
    _ = len(v4)
    v23 += (v23 ^ (^(v23)))
    if false {
        v1 -= float64(1e10)
    } else {
        fmt.Println(v1, v1)
        v34 := GenS4{F0: 127}
        _ = v34
    }
    var v35 uint32 = v3
    var v36 uint64 = uint64(v3)
    v3 -= v3
    var v37 int = 32768
    v6 -= v6
    if (v37 > v37) {
        v36 += v36
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 uint64 = 4294967295
        _ = fi38
        _ = v39
    }
    v40 := GenS4{F0: 79}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
