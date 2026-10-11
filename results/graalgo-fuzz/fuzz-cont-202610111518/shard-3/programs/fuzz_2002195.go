package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 int, p2 string) bool {
    switch p1 {
    case -339:
        for fi4 := 0; fi4 < 17; fi4++ {
            v5 := []int{}
            _ = fi4
            _ = v5
        }
    case 0:
        switch (p1 / ((p1) | 1)) {
        case 128:
            v6 := GenS1{F0: -234, F1: -2147483648, F2: (p1 & (p1 - (p1 / ((p1) | 1))))}
            _ = v6
        case -474:
            var v7 float32 = 0.1
            _ = v7
        }
    default:
        fmt.Println(p0, p0)
    }
    switch -9223372036854775808 {
    case -479:
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{fi8, fi8}
            _ = fi8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            p1 = (fi10 * (p1 & fi10))
            for fi11 := 0; fi11 < 14; fi11++ {
                if (1 < 39) {
                    var v12 string = "hello"
                    if (p0 < p0) {
                        fmt.Println(fi11, v12, p1)
                        if false {
                            v13 := GenS1{F0: 42, F1: -1, F2: fi11}
                            var v14 bool = (((!(!(!false))) && (true || true)) && false)
                            _ = v13
                            _ = v14
                        }
                    }
                    v15 := []int{fi10}
                    _ = v12
                    _ = v15
                }
                var v16 string = p2
                var v17 float32 = (0.1 - 3.14159)
                _ = fi11
                _ = v16
                _ = v17
            }
            v18 := []int{(p1 - (^(p1)))}
            _ = fi10
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (0 <= 182)
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := make(map[int]uint32)
    var v3 float64 = 1e10
    var v19 uint = 0
    v2[127] = 65535
    v19 += v19
    v20 := GenS1{F0: 1, F1: 1, F2: 127}
    switch "\n\t" {
    case "":
        v3 = v3
    default:
        switch 32767 {
        case -1:
            if (-1 == -1) {
                if false {
                    _ = len(v2)
                    var v21 uint8 = 1
                    for fi22 := 0; fi22 < 5; fi22++ {
                        for fi23 := 0; fi23 < 2; fi23++ {
                            switch (fi23 / (((-(fi22))) | 1)) {
                            case 128:
                                var v24 uint64 = 42
                                _ = v24
                            }
                            var v25 int16 = 1000
                            var v26 int32 = 42
                            _ = fi23
                            _ = v25
                            _ = v26
                        }
                        _ = fi22
                    }
                    _ = v21
                }
            }
        case 127:
            var v27 bool = false
            _ = v27
        default:
            v28 := GenS2{F0: 1}
            _ = v28
        }
    }
    var v29 int8 = 0
    fmt.Println(v3, v29, v29)
    v29 += 127
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := GenS1{F0: 32767, F1: -2147483648, F2: (fi30 | 9223372036854775807)}
        var v32 float64 = ((float64(3.14159) / (v3 * ((v3 - (float64(0.1) * float64(1e10))) * v3))) + v3)
        _ = fi30
        _ = v31
        _ = v32
    }
    switch "" {
    case "foo bar":
        delete(v2, 687)
    case "hello":
        for fi33 := 0; fi33 < 5; fi33++ {
            v3 -= v3
            _ = fi33
        }
    default:
        v34 := GenS2{F0: -32768}
        _ = v34
    }
    v19 = 65535
    fmt.Println(v19, v29, v19)
    for fi35 := 0; fi35 < 12; fi35++ {
        v29 = v29
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
}
