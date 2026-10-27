package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint32 }
func genFn0() uint16 {
    var v4 int16 = 1
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{(fi6 & fi5), -9223372036854775808}
            if false {
                var v8 int = (fi6 - 9223372036854775807)
                v8 = v8
                fmt.Println(v8, fi5)
                _ = v8
            }
            _ = fi6
            _ = v7
        }
        if true {
            v4 = 32767
            if (9223372036854775807 == -1) {
                for fi9 := 0; fi9 < 20; fi9++ {
                    v4 = (v4 + v4)
                    v10 := GenS1{F0: (!false)}
                    _ = fi9
                    _ = v10
                }
                if false {
                    fmt.Println(fi5, v4)
                }
                var v11 int32 = 383
                _ = v11
            } else {
                var v12 int64 = -9223372036854775808
                if false {
                    v13 := []int{(-(fi5)), (fi5 + (fi5 / (((fi5 / ((fi5) | 1))) | 1))), int(v12), fi5}
                    _ = v13
                }
                _ = v12
            }
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(fi14)
            if (-333 <= v4) {
                var v15 bool = false
                _ = v15
            } else {
                var v16 uint64 = 18446744073709551615
                _ = v16
            }
            v17 := []int{}
            v4 = (v4 * v4)
            fmt.Println(v4, v4)
            v17 = append(v17, fi5)
            _ = fi14
            _ = v17
        }
        _ = fi5
    }
    v18 := GenS1{F0: true}
    if true {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 float32 = (0.1 - ((0.0 + 1.5) + 0.1))
            v4 += v4
            v21 := GenS2{F0: ("hello" < ("foo bar" + "foo bar")), F1: 1}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    _ = v18
    _ = v4
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    v22 := GenS3{F0: -2147483648, F1: 18446744073709551615, F2: -9223372036854775808}
    v23 := []int{2147483647}
    var v24 uint64 = 42
    v25 := GenS4{F0: 42}
    var v26 uint16 = uint16(v24)
    var v27 float32 = 1.5
    v2 = false
    v24 += v24
    if (58 <= 42) {
        v26 += (v26 - 0)
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 float32 = -2.25
            _ = fi28
            _ = v29
        }
    } else {
        fmt.Println(v2, v26)
        v2 = v2
    }
    v30 := GenS5{F0: 960, F1: uint16(v24), F2: -1}
    var v31 int = -332
    _ = len(v1)
    var v32 bool = true
    v31 = v31
    v33 := GenS6{F0: 38}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
