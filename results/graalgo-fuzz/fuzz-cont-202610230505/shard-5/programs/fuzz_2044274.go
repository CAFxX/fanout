package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint, p1 int64) int32 {
    v7 := []int{9223372036854775807, -456}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}
func genFn1(p0 int8) uint {
    p0 = (p0 ^ p0)
    var v8 float64 = (1.5 - (3.14159 / 1e10))
    _ = p0
    _ = v8
    return 138
}

func main() {
    v1 := []int{2147483647, 32767, 0, 0}
    var v2 int16 = 1000
    var v3 float32 = 1.5
    var v4 float64 = float64((0.0 + ((0.1 / (v3 - 3.14159)) / (1e10 * v3))))
    var v5 int = -1
    v6 := make(map[int]uint32)
    if false {
        var v9 int = (v5 ^ (v5 / ((v5) | 1)))
        switch v9 {
        case 42:
            var v10 string = "hello"
            _ = v10
        default:
            v11 := v6[int(v2)]
            _ = v11
        }
        _ = v9
    } else {
        if false {
            var v12 uint32 = 980
            _ = v12
        }
        var v13 uint32 = 2147483648
        fmt.Println(v5, v2)
        _ = v13
    }
    switch (-(v5)) {
    case 127:
        var v14 uint = 1
        _ = v14
    case 4294967295:
        var v15 uint8 = 60
        _ = v15
    case 736:
        v16 := GenS1{F0: (("a" + "a") + "hello"), F1: 2147483647, F2: "foo bar"}
        _ = v16
    }
    var v17 bool = (42 >= -115)
    for fi18 := 0; fi18 < 17; fi18++ {
        v2 = (v2 * (v2 / ((-32768) | 1)))
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 uint16 = 6
            if (v4 != v4) {
                var v21 uint = 2
                var v22 bool = (-961 != 1)
                v23 := GenS2{F0: 50}
                _ = v21
                _ = v22
                _ = v23
            } else {
                var v24 string = (("foo bar" + "") + "\n\t")
                v2 = v2
                _ = v24
            }
            v25 := GenS3{F0: 0, F1: 4294967295}
            _ = fi19
            _ = v20
            _ = v25
        }
        v26 := GenS4{F0: 788}
        _ = fi18
        _ = v26
    }
    v27 := v6[(v5 / (((v5 | v5)) | 1))]
    v28 := GenS3{F0: v5, F1: 2}
    switch int(v27) {
    case 2147483647:
        v29 := GenS2{F0: v27}
        _ = v29
    case 1:
        for fi30 := 0; fi30 < 4; fi30++ {
            v17 = v17
            var v31 int8 = -128
            _ = fi30
            _ = v31
        }
    default:
        v6[v5] = (v27 | (v27 * (v27 << ((genFn1(-128)) & 63))))
    }
    v17 = v17
    if len(v1) > 0 {
        v1[0] = v5
    }
    if (86 < 42) {
        if v17 {
            var v32 int32 = -758
            var v33 int8 = 34
            v4 = (v4 + v4)
            _ = v32
            _ = v33
        } else {
            switch -788 {
            case 1:
                var v34 float32 = 0.0
                _ = v34
            }
        }
        fmt.Println(v4)
    } else {
        var v35 string = "foo bar"
        _ = v35
    }
    var v36 float64 = (((float64((1.5 - 1e10)) + (((1e10 + (v4 + 1e10)) - v4) - v4)) - (1e10 + v4)) / 1.0)
    var v37 uint16 = 327
    if ("\n\t" <= ("foo bar" + "foo bar")) {
        var v38 float64 = (float64(((0.1 - float32(v4)) - v3)) - ((v4 * ((v4 + (3.14159 - 3.14159)) - ((0.0 + 1e10) + float64((v3 / v3))))) / 0.1))
        for fi39 := 0; fi39 < 4; fi39++ {
            if v17 {
                v37 -= uint16(v27)
                v3 -= float32(v4)
            } else {
                v40 := GenS5{F0: 2}
                switch "" {
                case "a":
                    var v41 int64 = 42
                    _ = v41
                default:
                    var v42 int64 = 42
                    _ = v42
                }
                _ = v40
            }
            v36 = -2.25
            v37 = 506
            _ = fi39
        }
        v17 = v17
        _ = v38
    }
    var v43 string = "hello"
    v44 := GenS6{F0: 1}
    var v45 int32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
