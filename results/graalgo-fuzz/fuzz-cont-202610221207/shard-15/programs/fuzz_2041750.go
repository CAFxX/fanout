package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 int8; F1 uint64; F2 string }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 int32, p1 int) uint16 {
    switch 2147483647 {
    case 0:
        p1 = 32767
    case 127:
        for fi5 := 0; fi5 < 15; fi5++ {
            if ((1 != 42) && (p0 == p0)) {
                fmt.Println(p1, fi5)
                var v6 float32 = (-2.25 * 0.0)
                v6 -= 1e10
                _ = v6
            } else {
                if (0 < 127) {
                    switch fi5 {
                    case 127:
                        for fi7 := 0; fi7 < 4; fi7++ {
                            fmt.Println(p1)
                            v8 := []int{}
                            var v9 float64 = 3.14159
                            var v10 int = fi5
                            _ = fi7
                            _ = v10
                            _ = v8
                            _ = v9
                        }
                    case 32767:
                        p1 = -102
                    default:
                        var v11 bool = false
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 13; fi12++ {
                        p0 -= 0
                        v13 := []int{-184, 429}
                        var v14 uint = 4294967295
                        var v15 int8 = -1
                        var v16 uint = v14
                        v17 := GenS1{F0: (v15 << ((v14) & 63)), F1: 1, F2: ("" + ("foo bar" + ("\n\t" + ("hello" + "\n\t"))))}
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
                p1 += 127
            }
            v18 := GenS1{F0: 0, F1: 0, F2: "hello"}
            if false {
                for fi19 := 0; fi19 < 20; fi19++ {
                    fmt.Println(p1, p1, fi19)
                    switch (^(fi5)) {
                    case 127:
                        v20 := GenS2{F0: 42, F1: 419, F2: "foo bar"}
                        _ = v20
                    case 0:
                        var v21 float32 = (3.14159 - 3.14159)
                        _ = v21
                    default:
                        switch (fi19 - fi19) {
                        case -9223372036854775808:
                            if (1 >= -32768) {
                                v22 := []int{p1, -2147483648, 127}
                                _ = v22
                            }
                        default:
                            v23 := []int{2147483647, 2147483647, 9223372036854775807}
                            _ = v23
                        }
                    }
                    var v24 string = ""
                    var v25 uint16 = 642
                    _ = fi19
                    _ = v24
                    _ = v25
                }
                p1 += 1
            } else {
                var v26 int8 = -63
                _ = v26
            }
            _ = fi5
            _ = v18
        }
    case 661:
        p0 = (p0 & -2147483648)
    default:
        p0 += p0
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 458
}
func genFn1(p0 int64, p1 uint16) int64 {
    p0 += p0
    var v27 bool = false
    _ = p0
    _ = p1
    _ = v27
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -293
    v3 := make(map[uint32]bool)
    v4 := []string{("" + ("foo bar" + "hello")), ((("" + "") + "a") + "a"), "foo bar"}
    var v28 uint32 = 855
    v28 = v28
    if true {
        v2 = v2
        var v29 int64 = 1
        if (true && true) {
            for fi30 := 0; fi30 < 2; fi30++ {
                switch fi30 {
                case 4294967295:
                    if (true || true) {
                        var v31 int = fi30
                        v28 = (v28 - (v28 << ((704) & 63)))
                        _ = v31
                    } else {
                        var v32 bool = (797 < 717)
                        _ = v32
                    }
                case 2147483647:
                    if false {
                        var v33 int16 = (v2 >> ((281) & 63))
                        _ = v33
                    } else {
                        var v34 uint32 = (v28 % ((v28) | 1))
                        _ = v34
                    }
                }
                _ = fi30
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v29
    }
    var v35 float64 = -2.25
    fmt.Println(v2, v28, v28)
    var v36 uint8 = 128
    v35 = (v35 + v35)
    v37 := GenS3{F0: 0, F1: v2}
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 uint8 = (v36 << ((684) & 63))
        for fi40 := 0; fi40 < 12; fi40++ {
            v4 = append(v4, ("" + "\n\t"))
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v41 int16 = v2
        _ = v41
    default:
        var v42 uint32 = v28
        _ = v42
    }
    v35 += v35
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
