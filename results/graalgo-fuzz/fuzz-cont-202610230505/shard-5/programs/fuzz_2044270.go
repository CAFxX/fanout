package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int16; F2 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint, p1 int16, p2 int8) int32 {
    switch -717 {
    case -182:
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int32 = -2147483648
            v9 := GenS1{F0: fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        v10 := []int{291, -229, 1, -337}
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: int(p1)}
        if (!false) {
            fmt.Println(p2, p0, fi11)
            if false {
                v13 := GenS2{F0: int32(fi11), F1: (p1 >> ((p0) & 63)), F2: 4294967295}
                _ = v13
            } else {
                var v14 uint = p0
                p0 += v14
                _ = v14
            }
            if ((42 < 2147483647) || ((p0 != (p0 % (((p0 ^ p0)) | 1))) || ((-1 <= -1) || false))) {
                fmt.Println(p2, p0)
                v15 := []int{127}
                var v16 uint8 = 255
                _ = v15
                _ = v16
            } else {
                v17 := GenS2{F0: -2147483648, F1: (p1 / (((p1 << ((p0) & 63))) | 1)), F2: 0}
                for fi18 := 0; fi18 < 6; fi18++ {
                    if (false && true) {
                        if (!(!(p1 <= p1))) {
                            p1 -= p1
                        } else {
                            p1 -= p1
                        }
                    }
                    v19 := []int{4294967295, fi11, fi11, 9223372036854775807}
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
        } else {
            var v20 uint = p0
            if false {
                switch fi11 {
                case 2147483647:
                    for fi21 := 0; fi21 < 3; fi21++ {
                        if false {
                            v22 := []int{int(p0)}
                            _ = v22
                        }
                        switch fi21 {
                        case -748:
                            v23 := []int{(fi11 - fi21), fi11}
                            _ = v23
                        }
                        _ = fi21
                    }
                case 820:
                    v24 := []int{}
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 11; fi25++ {
                    if true {
                        var v26 uint8 = 103
                        _ = v26
                    }
                    _ = fi25
                }
            }
            _ = v20
        }
        var v27 int = fi11
        _ = fi11
        _ = v12
        _ = v27
    }
    var v28 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 0
}

func main() {
    var v1 int64 = 42
    var v2 uint = 292
    var v3 float64 = (0.0 / 1.0)
    var v4 int64 = v1
    var v5 uint32 = 966
    var v6 int32 = 42
    v3 = 1e10
    if (!true) {
        var v29 int32 = (v6 ^ v6)
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        switch fi30 {
        case 1:
            v1 = 156
        case 221:
            var v31 float32 = 0.0
            _ = v31
        }
        _ = fi30
    }
    v32 := GenS2{F0: v6, F1: 990, F2: 4294967295}
    v1 += -921
    if (-32768 >= -32768) {
        var v33 uint8 = 70
        _ = v33
    }
    if true {
        v4 = 42
        v34 := GenS3{F0: genFn0((^(v2)), 32767, 42), F1: -9223372036854775808, F2: 65535}
        v35 := GenS2{F0: (v6 >> ((v2) & 63)), F1: -32768, F2: 0}
        _ = v34
        _ = v35
    }
    var v36 string = "hello"
    v37 := GenS4{F0: v4}
    v5 = v5
    v38 := []int{32768, 83, -862, 9223372036854775807}
    var v39 string = v36
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
