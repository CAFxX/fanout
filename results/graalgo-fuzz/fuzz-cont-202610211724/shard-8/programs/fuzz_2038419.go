package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint; F1 bool; F2 bool }
func genFn0(p0 uint32) int8 {
    p0 -= p0
    var v4 float32 = 1.5
    v5 := GenS1{F0: 978}
    fmt.Println(p0, v4)
    _ = p0
    _ = v4
    _ = v5
    return -114
}
func genFn1(p0 string, p1 int64, p2 uint) uint32 {
    var v6 uint = 2
    var v7 uint64 = 0
    v8 := []int{0, 2147483647, -417}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint8 = 165
    var v2 bool = false
    var v3 string = ("" + (((("foo bar" + "") + "hello") + ("hello" + "")) + ""))
    if v2 {
        var v9 uint = 2
        _ = v9
    }
    v2 = v2
    fmt.Println(v1, v2)
    var v10 float32 = 3.14159
    v11 := GenS2{F0: "hello", F1: 768}
    var v12 uint16 = 255
    fmt.Println(v2, v3)
    fmt.Println(v10, v10, v10)
    v13 := GenS3{F0: -1, F1: 1, F2: -1}
    fmt.Println(v3, v10)
    if (v1 < (v1 + v1)) {
        for fi14 := 0; fi14 < 1; fi14++ {
            switch fi14 {
            case 32768:
                var v15 int8 = -49
                _ = v15
            case -1:
                v16 := []int{fi14, (fi14 ^ fi14), (fi14 * fi14)}
                _ = v16
            }
            _ = fi14
        }
        v17 := GenS4{F0: 2, F1: v2, F2: v2}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            v3 = (v3 + v3)
            v19 := []int{(fi18 >> ((1) & 63)), (fi18 & -183), 127}
            _ = fi18
            _ = v19
        }
        if v2 {
            var v20 int8 = -1
            var v21 uint8 = 45
            for fi22 := 0; fi22 < 11; fi22++ {
                if v2 {
                    var v23 uint64 = 42
                    _ = v23
                } else {
                    fmt.Println(v12, v3)
                    var v24 uint64 = 632
                    if v2 {
                        var v25 uint8 = (v21 - v1)
                        v26 := GenS4{F0: 65535, F1: (1 >= 810), F2: (0 == 42)}
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS3{F0: 55, F1: uint64(v12), F2: 993}
                        var v28 uint16 = v12
                        _ = v27
                        _ = v28
                    }
                    _ = v24
                }
                for fi29 := 0; fi29 < 18; fi29++ {
                    if ((-(v20)) == (v20 / ((v20) | 1))) {
                        v12 += (^(v12))
                        var v30 int = (fi29 >> ((638) & 63))
                        var v31 int32 = 2147483647
                        _ = v30
                        _ = v31
                    } else {
                        v1 = (v21 | 18)
                    }
                    v32 := []int{(fi22 - (-(fi22)))}
                    _ = fi29
                    _ = v32
                }
                if (-32768 != 42) {
                    var v33 float64 = float64(v10)
                    for fi34 := 0; fi34 < 9; fi34++ {
                        v35 := []int{(fi22 & fi22), (fi34 & (fi22 - (^(fi34))))}
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
                _ = fi22
            }
            _ = v20
            _ = v21
        } else {
            fmt.Println(v2, v12)
            v36 := []int{}
            _ = v36
        }
        var v37 int32 = 0
        _ = v37
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 uint8 = v1
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
