package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int16, p1 uint, p2 bool) int32 {
    var v7 int8 = -33
    var v8 uint = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int8, p1 int64) int16 {
    v9 := []int{-395, 128, 0, -254}
    var v10 float64 = (3.14159 - -2.25)
    var v11 int64 = 0
    v12 := GenS1{F0: 2, F1: 1, F2: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    var v5 int = 2147483647
    var v6 int16 = 1
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 float64 = 3.14159
        _ = fi13
        _ = v14
    }
    var v15 bool = true
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: 1, F1: -1, F2: 127}
        v15 = (10 > 127)
        if (913 <= 783) {
            var v18 uint64 = 2
            _ = v18
        } else {
            var v19 uint64 = 219
            v20 := v1[v5]
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v5 = v5
    v15 = (("foo bar" < "foo bar") || (v15 && v15))
    var v21 float64 = 1.5
    v5 += v5
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 string = (("\n\t" + ("foo bar" + "foo bar")) + ("hello" + ("hello" + "foo bar")))
        v5 += v5
        _ = fi22
        _ = v23
    }
    switch v5 {
    case -2147483648:
        var v24 uint32 = 2147483647
        _ = v24
    case 222:
        var v25 uint64 = 42
        _ = v25
    case -9223372036854775808:
        fmt.Println(v5, v5)
    default:
        fmt.Println(v21)
    }
    switch v5 {
    case 32768:
        if (v15 && (v15 && (100 != 1))) {
            for fi26 := 0; fi26 < 19; fi26++ {
                v5 += fi26
                _ = fi26
            }
        } else {
            switch 128 {
            case 128:
                v15 = v15
            case -1:
                v5 += v5
            case 525:
                v27 := GenS1{F0: 65535, F1: -53, F2: 227}
                _ = v27
            default:
                v28 := GenS1{F0: 32767, F1: 32, F2: 168}
                _ = v28
            }
            var v29 uint16 = 2
            for fi30 := 0; fi30 < 15; fi30++ {
                v31 := []int{fi30, (v5 * 128), (v5 + v5), fi30}
                fmt.Println(v5, v5, v15)
                _ = fi30
                _ = v31
            }
            _ = v29
        }
    case 128:
        if ("" != "a") {
            fmt.Println(v21, v21)
            switch (v5 - v5) {
            case -2147483648:
                v32 := GenS2{F0: 42, F1: -669}
                _ = v32
            case -835:
                v33 := GenS1{F0: 42, F1: 0, F2: 255}
                _ = v33
            case 0:
                v34 := []int{}
                _ = v34
            }
            v5 = 4294967295
        }
    default:
        var v35 int64 = 1
        _ = v35
    }
    switch (v5 % (((v5 / (((v5 >> ((42) & 63))) | 1))) | 1)) {
    case -894:
        switch (^(v5)) {
        case 517:
            if (true || ((!(1 <= 42)) || v15)) {
                var v36 uint64 = 2
                _ = v36
            } else {
                for fi37 := 0; fi37 < 13; fi37++ {
                    fmt.Println(v15, v21)
                    switch (v5 + (v5 - -2147483648)) {
                    case 9223372036854775807:
                        _ = len(v4)
                    case 128:
                        v38 := []int{(v5 + -9223372036854775808), (v5 ^ fi37), (fi37 / ((v5) | 1))}
                        _ = v38
                    case 2147483647:
                        for fi39 := 0; fi39 < 2; fi39++ {
                            v6 += v6
                            _ = fi39
                        }
                    default:
                        v40 := GenS2{F0: 65535, F1: -9223372036854775808}
                        _ = v40
                    }
                    _ = fi37
                }
                var v41 int16 = 32767
                _ = v41
            }
        case 1:
            v42 := GenS3{F0: v5}
            _ = v42
        }
    case -519:
        var v43 int64 = 696
        _ = v43
    default:
        v44 := GenS4{F0: "a", F1: 361}
        _ = v44
    }
    var v45 int16 = (v6 - (v6 % ((v6) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
