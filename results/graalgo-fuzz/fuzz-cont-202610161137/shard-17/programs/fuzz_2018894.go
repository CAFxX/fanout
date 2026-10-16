package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 bool; F2 int8 }
func genFn0(p0 int, p1 int32) uint16 {
    p1 += (-(p1))
    if ((-2.25 / 1e10) == float64((float32(float64(-2.25)) + 1.5))) {
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(p0, p0, p0)
            _ = fi7
        }
    } else {
        var v8 uint32 = 1
        _ = v8
    }
    p0 -= p0
    v9 := []int{4294967295, p0}
    _ = p0
    _ = p1
    _ = v9
    return 404
}

func main() {
    var v1 int = 32768
    var v2 uint64 = 0
    var v3 bool = false
    v4 := make(map[int]uint32)
    var v5 uint64 = v2
    v6 := []int{4294967295, v1, v1}
    var v10 uint8 = uint8(v2)
    v3 = ((v3 && false) || v3)
    v10 += v10
    fmt.Println(v10, v10)
    if true {
        var v11 string = "hello"
        if true {
            switch v1 {
            case 1:
                if ((1e10 - 1.5) > float32(float64(1.5))) {
                    v5 += v5
                    if (4294967295 != 42) {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            v2 = uint64(v10)
                            _ = fi12
                        }
                    }
                    v4[v1] = 0
                }
            case 877:
                var v13 int64 = 1
                _ = v13
            default:
                if v3 {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v15 := GenS1{F0: v3, F1: true, F2: 242}
                        var v16 int64 = 1
                        v17 := GenS1{F0: true, F1: v3, F2: 2147483648}
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
            }
            if len(v6) > 0 {
                v6[0] = v1
            }
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                v3 = v3
                v19 := GenS1{F0: true, F1: v3, F2: 0}
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 7; fi20++ {
                v10 = v10
                var v21 int8 = 100
                _ = fi20
                _ = v21
            }
        }
        if (((4294967295 < 933) || (!false)) && v3) {
            for fi22 := 0; fi22 < 10; fi22++ {
                v23 := GenS1{F0: v3, F1: v3, F2: 65535}
                v24 := GenS2{F0: 65535}
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS2{F0: 0}
            var v26 bool = true
            _ = v25
            _ = v26
        } else {
            var v27 uint16 = 1
            fmt.Println(v3, v11)
            _ = v27
        }
        _ = v11
    } else {
        v28 := v4[v1]
        var v29 int = int(v2)
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        switch (-(fi30)) {
        case 42:
            fmt.Println(v5)
        case -1:
            if (false && (1 != 42)) {
                var v31 uint8 = 255
                _ = v31
            }
        }
        v1 += -9223372036854775808
        _ = fi30
    }
    _ = len(v4)
    v32 := GenS2{F0: 42}
    _ = len(v4)
    v33 := GenS3{F0: v3, F1: true, F2: -1}
    var v34 float64 = 1.5
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := GenS3{F0: (42 > -9223372036854775808), F1: v3, F2: 0}
        var v37 bool = v3
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
