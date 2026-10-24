package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 bool, p1 uint8, p2 int64) string {
    switch "foo bar" {
    case "a":
        if false {
            v6 := GenS1{F0: 1, F1: 110, F2: 173}
            _ = v6
        } else {
            v7 := []int{}
            p2 = p2
            fmt.Println(p0, p2)
            _ = v7
        }
    case "hello":
        if p0 {
            v8 := []int{-2147483648}
            _ = v8
        }
    case "\n\t":
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    v10 := GenS1{F0: 65535, F1: -32768, F2: p1}
    v11 := []int{1, 128, 42}
    v12 := GenS2{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (("\n\t" + "\n\t") + "\n\t")
}

func main() {
    v1 := []int{0, 4294967295}
    var v2 float64 = (-2.25 * -2.25)
    var v3 int16 = 149
    v4 := []uint32{4294967295, 300, 65535}
    v5 := make(map[uint32]int)
    v13 := GenS1{F0: 2, F1: 1, F2: 1}
    var v14 string = "foo bar"
    var v15 uint16 = 763
    switch "foo bar" {
    case "hello":
        if true {
            switch v14 {
            case "\n\t":
                var v16 int8 = 1
                _ = v16
            case "a":
                v1 = append(v1, 1)
            case "foo bar":
                var v17 float64 = v2
                _ = v17
            default:
                v14 = v14
            }
            if len(v1) > 0 {
                v1[0] = -881
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS1{F0: 2147483648, F1: v3, F2: 128}
                var v20 string = v14
                fmt.Println(v2)
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    case "\n\t":
        if (895 < 129) {
            if false {
                var v21 float64 = (1e10 / 3.14159)
                _ = v21
            }
            v5[4294967295] = -167
        }
    case "a":
        var v22 string = (v14 + v14)
        _ = v22
    default:
        var v23 int64 = -370
        _ = v23
    }
    var v24 int8 = -1
    if false {
        v2 = 0.0
    } else {
        v25 := GenS3{F0: v24, F1: "a"}
        for fi26 := 0; fi26 < 9; fi26++ {
            fmt.Println(v15, v24, v3)
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 float32 = -2.25
                if true {
                    var v29 int64 = -1
                    _ = v29
                }
                delete(v5, 0)
                fmt.Println(v14, v3, v24)
                v3 += v3
                _ = fi27
                _ = v28
            }
            v14 = "\n\t"
            _ = fi26
        }
        _ = v25
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            if len(v1) > 0 {
                v1[0] = -51
            }
            v14 = "hello"
            v32 := GenS2{F0: 255}
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    var v33 int32 = -2147483648
    var v34 int16 = (v3 * (v3 ^ v3))
    if true {
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 uint16 = v15
            _ = fi35
            _ = v36
        }
        if (false && false) {
            v3 = v3
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 uint16 = v15
                _ = fi37
                _ = v38
            }
        }
    } else {
        if true {
            for fi39 := 0; fi39 < 20; fi39++ {
                v40 := GenS4{F0: (v15 + 65535), F1: (v33 + v33), F2: 481}
                _ = fi39
                _ = v40
            }
            v41 := v5[876]
            _ = v41
        }
    }
    fmt.Println(v15, v15, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
