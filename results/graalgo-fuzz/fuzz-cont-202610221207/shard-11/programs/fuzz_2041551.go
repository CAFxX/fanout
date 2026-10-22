package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint32, p1 int32) int64 {
    switch 301 {
    case -2147483648:
        for fi5 := 0; fi5 < 14; fi5++ {
            p0 -= p0
            if false {
                for fi6 := 0; fi6 < 2; fi6++ {
                    v7 := GenS1{F0: 687, F1: p1}
                    p0 = p0
                    p0 = 1
                    switch "a" {
                    case "\n\t":
                        var v8 int = (^(fi5))
                        _ = v8
                    case "a":
                        var v9 int64 = 455
                        _ = v9
                    default:
                        var v10 int32 = (^(p1))
                        _ = v10
                    }
                    _ = fi6
                    _ = v7
                }
                v11 := GenS1{F0: 9223372036854775807, F1: -2147483648}
                _ = v11
            } else {
                v12 := GenS2{F0: 65535, F1: 255}
                switch -2147483648 {
                case 4294967295:
                    fmt.Println(fi5, p1)
                default:
                    switch int(p0) {
                    case 32768:
                        var v13 uint = 0
                        _ = v13
                    case -840:
                        var v14 uint64 = 2
                        _ = v14
                    case 9223372036854775807:
                        v15 := GenS2{F0: 934, F1: 0}
                        _ = v15
                    default:
                        v16 := []int{(fi5 / ((128) | 1)), (^(fi5))}
                        _ = v16
                    }
                }
                _ = v12
            }
            _ = fi5
        }
    case 641:
        switch 4294967295 {
        case 9223372036854775807:
            v17 := GenS3{F0: p1, F1: (p0 / (((p0 - p0)) | 1)), F2: -9223372036854775808}
            _ = v17
        case -1:
            v18 := GenS4{F0: 386, F1: 38}
            _ = v18
        case 2147483647:
            v19 := []int{}
            _ = v19
        default:
            v20 := []int{495}
            _ = v20
        }
    }
    v21 := []int{-447}
    _ = p0
    _ = p1
    _ = v21
    return 1
}

func main() {
    var v1 int16 = 1
    var v2 int = 0
    var v3 float64 = 1.5
    v4 := make(map[int]string)
    switch (-(v2)) {
    case -2147483648:
        switch 0 {
        case -2147483648:
            delete(v4, (v2 & v2))
        case -127:
            v22 := GenS4{F0: 957, F1: 2}
            _ = v22
        case 127:
            v3 += v3
        }
    default:
        fmt.Println(v2, v2, v1)
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v2 += fi23
        var v24 uint16 = 2
        _ = fi23
        _ = v24
    }
    var v25 float64 = v3
    v1 -= (^(v1))
    var v26 int64 = 9223372036854775807
    var v27 bool = false
    switch 950 {
    case 650:
        var v28 int64 = v26
        _ = v28
    default:
        fmt.Println(v27, v26)
    }
    v29 := GenS2{F0: 0, F1: 127}
    var v30 bool = v27
    if (!(3.14159 > (v25 * float64((3.14159 - 0.0))))) {
        v31 := GenS3{F0: 0, F1: 2, F2: genFn0(2147483648, -2147483648)}
        v32 := GenS1{F0: -877, F1: 342}
        _ = v31
        _ = v32
    }
    v26 = v26
    for fi33 := 0; fi33 < 4; fi33++ {
        v34 := v4[fi33]
        for fi35 := 0; fi35 < 19; fi35++ {
            var v36 int64 = v26
            v3 = 1e10
            var v37 string = v34
            _ = fi35
            _ = v36
            _ = v37
        }
        v38 := GenS1{F0: 127, F1: -2147483648}
        _ = fi33
        _ = v34
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
