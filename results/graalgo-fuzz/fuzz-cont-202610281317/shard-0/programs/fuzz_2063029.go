package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64) bool {
    fmt.Println(p0)
    p0 = 42
    p0 = p0
    if (false || true) {
        p0 = 0
        for fi7 := 0; fi7 < 14; fi7++ {
            if (!(92 != 100)) {
                var v8 uint8 = 128
                if true {
                    var v9 uint64 = uint64(v8)
                    v9 = (v9 & v9)
                    v10 := GenS1{F0: 42}
                    _ = v10
                    _ = v9
                }
                _ = v8
            } else {
                v11 := GenS1{F0: 255}
                p0 = 843
                var v12 float64 = -2.25
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                p0 -= (p0 & p0)
                v14 := []int{(fi13 * fi7)}
                _ = fi13
                _ = v14
            }
            var v15 uint16 = 255
            _ = fi7
            _ = v15
        }
    } else {
        p0 = (p0 * p0)
    }
    _ = p0
    return false
}
func genFn1(p0 int64, p1 uint) int16 {
    p1 = p1
    fmt.Println(p0, p1)
    var v16 int = 0
    v17 := GenS1{F0: p1}
    p1 = 255
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 42
    v3 := make(map[int]string)
    var v4 int8 = 107
    var v5 uint32 = 2147483647
    var v6 uint = 4294967295
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            if true {
                var v20 int16 = 32767
                _ = v20
            } else {
                fmt.Println(v4, v2)
                var v21 int32 = 2147483647
                switch fi18 {
                case -834:
                    v5 += (v5 | (v5 >> ((42) & 63)))
                default:
                    v2 = 42
                }
                v4 -= (v4 & v4)
                _ = v21
            }
            var v22 uint8 = v2
            _ = fi19
            _ = v22
        }
        fmt.Println(v2, v6, fi18)
        _ = fi18
    }
    if false {
        var v23 uint16 = 0
        var v24 float32 = 0.0
        v24 += v24
        _ = v23
        _ = v24
    } else {
        v25 := GenS1{F0: v6}
        v26 := []int{}
        _ = v25
        _ = v26
    }
    var v27 float64 = float64(1.5)
    v28 := GenS2{F0: "\n\t", F1: 1, F2: false}
    delete(v3, 2147483647)
    v29 := GenS3{F0: 58}
    var v30 string = "foo bar"
    for fi31 := 0; fi31 < 20; fi31++ {
        switch fi31 {
        case 4294967295:
            v5 = v5
        default:
            var v32 bool = (!false)
            _ = v32
        }
        switch ("foo bar" + v30) {
        case "a":
            switch "foo bar" {
            case "":
                fmt.Println(v27, v30, v6)
            }
        default:
            switch "\n\t" {
            case "hello":
                if (!((^(v4)) >= v4)) {
                    v2 += v2
                } else {
                    v27 = -2.25
                }
            case "":
                v1[-9223372036854775808] = (-(fi31))
            default:
                var v33 float32 = 3.14159
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(v2, v6, v5)
            v35 := GenS1{F0: v6}
            _ = fi34
            _ = v35
        }
        _ = fi31
    }
    v36 := GenS4{F0: 1}
    switch -980 {
    case 127:
        switch ((("" + v30) + v30) + "hello") {
        case "a":
            v6 = v6
        case "hello":
            v27 = float64((3.14159 * 3.14159))
        case "":
            v37 := v3[0]
            _ = v37
        }
    case 42:
        v27 = 1e10
    case -9223372036854775808:
        var v38 int64 = -488
        _ = v38
    }
    v39 := v1[-526]
    fmt.Println(v30, v39)
    var v40 uint64 = 18446744073709551615
    v40 = v40
    var v41 float32 = 0.1
    var v42 uint32 = (^(v5))
    var v43 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
