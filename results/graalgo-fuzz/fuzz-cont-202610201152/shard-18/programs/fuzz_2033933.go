package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int, p1 uint8) uint {
    var v4 uint32 = 0
    fmt.Println(p1, v4, p1)
    var v5 int16 = 1000
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: p1, F1: 1, F2: -496}
        switch "a" {
        case "hello":
            switch -2147483648 {
            case 136:
                var v8 int64 = int64(v4)
                _ = v8
            case 127:
                p1 = p1
            default:
                fmt.Println(p0, v4)
            }
        case "foo bar":
            if (false && false) {
                p1 = p1
                p0 = (p0 % (((p0 << ((290) & 63))) | 1))
            } else {
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint64 = 0
                    var v11 int32 = 409
                    v11 = 42
                    p1 -= (p1 % (((p1 >> ((65535) & 63))) | 1))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                if (true && false) {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        if false {
                            if false {
                                v13 := GenS2{F0: -9223372036854775808}
                                var v14 uint64 = 4294967295
                                var v15 float64 = ((1.5 * 1e10) / -2.25)
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            var v16 string = ("\n\t" + "\n\t")
                            v17 := []int{(^(fi12)), -928, fi6}
                            _ = v16
                            _ = v17
                        }
                        fmt.Println(fi12)
                        var v18 uint8 = 1
                        _ = fi12
                        _ = v18
                    }
                    p1 -= p1
                }
            }
        case "":
            fmt.Println(v4, v5)
        default:
            p0 += p0
        }
        var v19 uint = 0
        _ = fi6
        _ = v19
        _ = v7
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1() int64 {
    var v20 int16 = 1
    if false {
        var v21 int8 = 42
        var v22 int32 = -1
        _ = v21
        _ = v22
    } else {
        var v23 int = -448
        for fi24 := 0; fi24 < 13; fi24++ {
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 int32 = 42
                if (4294967295 < 42) {
                    v27 := GenS3{F0: ((-(v23)) <= (v23 * fi24)), F1: false}
                    _ = v27
                } else {
                    fmt.Println(fi25)
                    v28 := []int{(fi25 * 537), fi25, 0}
                    fmt.Println(fi25)
                    _ = v28
                }
                fmt.Println(v20, v20, v23)
                var v29 int64 = -1
                _ = fi25
                _ = v26
                _ = v29
            }
            for fi30 := 0; fi30 < 13; fi30++ {
                v31 := GenS4{F0: 2}
                _ = fi30
                _ = v31
            }
            var v32 int32 = 552
            _ = fi24
            _ = v32
        }
        var v33 int16 = 32767
        _ = v23
        _ = v33
    }
    v20 = -1
    _ = v20
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -1
    var v3 uint8 = 0
    v3 -= (v3 * (v3 + v3))
    var v34 float32 = 3.14159
    v35 := GenS1{F0: (v3 / ((v3) | 1)), F1: 826, F2: 0}
    var v36 int64 = 0
    v36 += 1
    fmt.Println(v36)
    v36 = 42
    if (false && (128 == 32768)) {
        v2 = v2
        for fi37 := 0; fi37 < 1; fi37++ {
            v3 = 70
            var v38 int = int(v3)
            v36 += genFn1()
            _ = fi37
            _ = v38
        }
    }
    var v39 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
}
