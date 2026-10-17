package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0() int64 {
    var v7 uint64 = 42
    v8 := GenS1{F0: 1, F1: -51, F2: v7}
    v7 += v7
    v7 += v7
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = (0.0 / (0.1 - (float32(1e10) / (float32(1e10) + float32(1.5)))))
    var v2 uint32 = 42
    v3 := []int{}
    v4 := make(map[int]string)
    v5 := []uint32{(v2 & (v2 ^ v2)), (v2 >> ((459) & 63)), v2, 2}
    var v6 int32 = -1
    var v9 int8 = 100
    switch "a" {
    case "hello":
        fmt.Println(v1, v6)
    }
    if (0.1 != 0.1) {
        switch 635 {
        case 127:
            if ((!false) && true) {
                switch -553 {
                case 1:
                    v6 += v6
                case 128:
                    switch -9223372036854775808 {
                    case 1:
                        v9 = (v9 % (((v9 % ((127) | 1))) | 1))
                    }
                default:
                    v2 = (v2 ^ v2)
                }
                v10 := GenS1{F0: (v9 * v9), F1: (-(v9)), F2: 4294967295}
                fmt.Println(v9)
                _ = v10
            } else {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v2 = 705
                    if false {
                        if false {
                            v12 := v4[int(v2)]
                            v9 += (v9 << ((853) & 63))
                            v13 := GenS1{F0: (v9 ^ -56), F1: v9, F2: 4294967295}
                            _ = v12
                            _ = v13
                        } else {
                            fmt.Println(fi11)
                        }
                    }
                    _ = fi11
                }
                fmt.Println(v9)
                v14 := GenS1{F0: v9, F1: (v9 % ((v9) | 1)), F2: 18446744073709551615}
                _ = v14
            }
        case 1:
            switch ("hello" + "hello") {
            case "foo bar":
                if ((^(v2)) >= 2147483647) {
                    _ = len(v5)
                }
            }
        case -1:
            var v15 uint32 = v2
            _ = v15
        default:
            v16 := GenS2{F0: (^(v2)), F1: 435}
            _ = v16
        }
    }
    if (true || false) {
        _ = len(v4)
        if (!false) {
            v2 = v2
        }
    }
    switch 29 {
    case -368:
        v17 := GenS2{F0: (v2 / (((v2 & v2)) | 1)), F1: 2}
        _ = v17
    case -2147483648:
        if false {
            fmt.Println(v6, v1, v6)
            var v18 float64 = 0.0
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := v4[fi19]
                for fi21 := 0; fi21 < 15; fi21++ {
                    delete(v4, int(v2))
                    delete(v4, 32767)
                    _ = fi21
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            var v22 uint32 = v2
            _ = v22
        }
    case 517:
        for fi23 := 0; fi23 < 19; fi23++ {
            switch (-(fi23)) {
            case 32768:
                var v24 int = -492
                _ = v24
            case 127:
                switch (fi23 & (fi23 * (fi23 >> ((255) & 63)))) {
                case 9223372036854775807:
                    var v25 float32 = (float32((0.0 / -2.25)) + v1)
                    _ = v25
                default:
                    _ = len(v5)
                }
            case 32767:
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 int32 = v6
                    var v28 int32 = (v27 + v6)
                    v2 += (v2 - 42)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            default:
                if (!true) {
                    for fi29 := 0; fi29 < 4; fi29++ {
                        v30 := GenS1{F0: (v9 % ((42) | 1)), F1: v9, F2: 1}
                        _ = fi29
                        _ = v30
                    }
                    var v31 int = fi23
                    _ = v31
                }
            }
            v9 = v9
            _ = fi23
        }
    default:
        switch -9223372036854775808 {
        case -9223372036854775808:
            _ = len(v4)
        }
    }
    v32 := GenS1{F0: (v9 / ((v9) | 1)), F1: v9, F2: 0}
    fmt.Println(v9, v2)
    switch 945 {
    case -887:
        v6 += v6
    case 42:
        var v33 int64 = 667
        _ = v33
    }
    v1 -= 1.5
    var v34 uint32 = v2
    var v35 uint64 = 0
    var v36 int = 32768
    var v37 int64 = 1
    v38 := GenS1{F0: v9, F1: 100, F2: v35}
    v36 = (v36 / ((v36) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
