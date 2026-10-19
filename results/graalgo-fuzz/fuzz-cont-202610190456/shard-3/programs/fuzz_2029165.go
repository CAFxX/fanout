package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 string; F1 string; F2 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint8, p1 string, p2 int64) uint {
    fmt.Println(p2, p1)
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 bool = (1 >= 680)
    var v5 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 20; fi7++ {
        v5 = (v5 | v5)
        v1 = v1
        _ = fi7
    }
    var v8 int8 = 100
    switch v1 {
    case "\n\t":
        if (v5 <= (-(v5))) {
            for fi9 := 0; fi9 < 1; fi9++ {
                if (false || v4) {
                    v10 := []int{(fi9 << ((738) & 63)), 128}
                    _ = v10
                }
                _ = fi9
            }
        }
    default:
        var v11 float64 = (3.14159 + 3.14159)
        _ = v11
    }
    v12 := GenS1{F0: 65535, F1: 359}
    for fi13 := 0; fi13 < 6; fi13++ {
        v4 = (v4 && true)
        fmt.Println(v5, v8)
        switch (-(fi13)) {
        case 127:
            for fi14 := 0; fi14 < 6; fi14++ {
                if ((v1 == "\n\t") && v4) {
                    if (true && v4) {
                        var v15 float64 = -2.25
                        _ = v15
                    }
                    fmt.Println(v1, fi13, fi14)
                } else {
                    v16 := GenS1{F0: 65535, F1: 2147483648}
                    switch (fi13 | fi13) {
                    case -440:
                        v5 = v5
                    case 810:
                        if v4 {
                            var v17 int = (fi14 >> ((genFn0(2, v1, 0)) & 63))
                            var v18 float32 = (-2.25 + 0.1)
                            _ = v17
                            _ = v18
                        }
                    case -9223372036854775808:
                        var v19 uint = 0
                        _ = v19
                    default:
                        v20 := GenS2{F0: (v1 + v1), F1: v1, F2: 32767}
                        _ = v20
                    }
                    _ = v16
                }
                var v21 int32 = 878
                delete(v2, fi14)
                _ = fi14
                _ = v21
            }
        case -2147483648:
            if false {
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 uint8 = 42
                    var v24 uint16 = 0
                    v25 := GenS3{F0: v23, F1: "", F2: 1}
                    var v26 int16 = int16(v24)
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        case -923:
            v8 += 0
        default:
            var v27 int = fi13
            _ = v27
        }
        _ = fi13
    }
    if (true || false) {
        v28 := []int{}
        _ = v28
    }
    v29 := []int{906}
    if (v4 && (v4 && (!(v4 && v4)))) {
        var v30 uint = 4294967295
        if v4 {
            var v31 uint = 4294967295
            v32 := GenS1{F0: v30, F1: 1}
            v33 := GenS4{F0: 0, F1: v31, F2: 498}
            _ = v31
            _ = v32
            _ = v33
        } else {
            switch -9223372036854775808 {
            case -2147483648:
                var v34 int = 1
                _ = v34
            default:
                v4 = true
            }
        }
        _ = v30
    }
    v8 = 42
    fmt.Println(v1)
    delete(v3, v1)
    if v4 {
        for fi35 := 0; fi35 < 8; fi35++ {
            switch (fi35 - 4) {
            case -9223372036854775808:
                v8 = v8
            case 365:
                v36 := GenS5{F0: 850}
                _ = v36
            default:
                v37 := GenS5{F0: 0}
                _ = v37
            }
            var v38 int16 = -223
            fmt.Println(fi35)
            _ = fi35
            _ = v38
        }
    }
    fmt.Println(v5, v4, v8)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
