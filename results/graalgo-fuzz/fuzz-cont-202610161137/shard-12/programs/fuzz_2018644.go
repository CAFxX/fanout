package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32, p1 uint16) int {
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            if (!false) {
                p0 -= p0
            } else {
                var v6 string = ""
                var v7 uint8 = 42
                var v8 int8 = -54
                var v9 int64 = int64(v8)
                _ = v6
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := []int{}
            p1 = p1
            _ = fi5
            _ = v10
        }
        p0 = p0
        _ = fi4
    }
    fmt.Println(p0)
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint8 = 127
            _ = fi11
            _ = v12
        }
        fmt.Println(p0, p0)
    } else {
        v13 := []int{1, -1, 959, -577}
        var v14 int32 = -1
        if (42 > 42) {
            v15 := GenS1{F0: ((42 > 4294967295) || true), F1: 277}
            _ = v15
        }
        _ = v13
        _ = v14
    }
    var v16 uint64 = 1
    v17 := []int{4294967295}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 int8 = -53
    var v2 int = 42
    v3 := make(map[string]int)
    _ = len(v3)
    for fi18 := 0; fi18 < 4; fi18++ {
        v2 -= genFn0(0, 32767)
        _ = fi18
    }
    if (!false) {
        v1 = v1
        switch (("" + "a") + "\n\t") {
        case "hello":
            v2 += v2
        default:
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 int16 = 0
                switch fi19 {
                case -767:
                    v21 := []int{}
                    _ = v21
                case -1:
                    v22 := GenS2{F0: 65535, F1: v2}
                    _ = v22
                default:
                    var v23 float64 = float64((-2.25 * (-2.25 - (float32(0.0) / (3.14159 + (0.0 + (0.0 + 0.0)))))))
                    _ = v23
                }
                var v24 int32 = 756
                _ = fi19
                _ = v20
                _ = v24
            }
        }
        if false {
            for fi25 := 0; fi25 < 18; fi25++ {
                if (492 <= 18446744073709551615) {
                    v26 := GenS3{F0: 0}
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 int64 = 665
                        v1 = v1
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                }
                _ = fi25
            }
        }
    }
    v29 := []int{int(v1), v2, (v2 * (-(v2))), genFn0(65535, 1)}
    v30 := GenS4{F0: 18446744073709551615}
    v2 = v2
    v31 := GenS1{F0: (0 == 0), F1: 42}
    v2 -= v2
    var v32 uint32 = 2147483648
    v33 := GenS2{F0: 312, F1: v2}
    var v34 int16 = -1
    if len(v29) > 0 {
        v29[0] = -2147483648
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 int32 = 1
        _ = fi35
        _ = v36
    }
    v29 = append(v29, genFn0(v32, 65535))
    var v37 int64 = -397
    var v38 bool = (0.1 < 1e10)
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS2{F0: 65535, F1: 32768}
        if v38 {
            v41 := v3["a"]
            for fi42 := 0; fi42 < 15; fi42++ {
                v43 := GenS3{F0: v32}
                _ = fi42
                _ = v43
            }
            _ = v41
        }
        fmt.Println(v2, v34, fi39)
        _ = fi39
        _ = v40
    }
    v37 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
}
