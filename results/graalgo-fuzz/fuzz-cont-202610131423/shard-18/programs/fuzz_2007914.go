package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64) uint16 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{fi6, (fi6 + fi6), fi6, fi6}
        _ = fi6
        _ = v7
    }
    _ = p0
    return 882
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]string)
    var v3 uint16 = 642
    var v4 uint32 = 718
    var v5 int16 = 1000
    delete(v2, v4)
    switch "a" {
    case "a":
        switch "" {
        case "\n\t":
            fmt.Println(v5, v3)
        default:
            v4 -= v4
        }
    case "hello":
        v8 := []int{128, -229, 32767}
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := v2[(v4 | v4)]
        var v11 uint = 2
        switch fi9 {
        case 1:
            fmt.Println(v4, v1, v3)
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 bool = false
    v12 = v12
    if (v1 == v1) {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: -1, F1: uint(v3), F2: 94}
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int64 = 1
                var v17 uint = 42
                fmt.Println(fi13, v4)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        v3 += v3
    } else {
        switch int(v5) {
        case 127:
            var v18 int16 = v5
            _ = v18
        case -329:
            var v19 bool = v12
            _ = v19
        default:
            v5 = (v5 >> ((65535) & 63))
        }
    }
    if v12 {
        var v20 int64 = -210
        var v21 uint8 = 167
        _ = v20
        _ = v21
    } else {
        v22 := GenS1{F0: 287, F1: 42, F2: 1}
        if (v4 < (^(v4))) {
            if v12 {
                if v12 {
                    var v23 int16 = v5
                    _ = v23
                }
                v24 := GenS1{F0: -537, F1: 0, F2: 1}
                _ = v24
            }
            v12 = true
        }
        _ = v22
    }
    v25 := GenS1{F0: 42, F1: 2, F2: uint8(v3)}
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 float64 = 0.0
        for fi28 := 0; fi28 < 16; fi28++ {
            if v12 {
                if true {
                    if (v1 < v1) {
                        v2[4294967295] = ""
                        v12 = v12
                        v29 := GenS1{F0: -9223372036854775808, F1: 2, F2: 0}
                        _ = v29
                    } else {
                        var v30 float64 = (v27 * (v27 + float64(-2.25)))
                        delete(v2, 881)
                        _ = v30
                    }
                    var v31 float32 = -2.25
                    v32 := GenS2{F0: 484}
                    var v33 int16 = (v5 * int16(v4))
                    var v34 int32 = -577
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            } else {
                v4 = (v4 % ((v4) | 1))
            }
            v35 := GenS2{F0: 916}
            _ = fi28
            _ = v35
        }
        _ = fi26
        _ = v27
    }
    delete(v2, v4)
    if ((v1 / ((v1) | 1)) == v1) {
        delete(v2, 2)
        v36 := GenS3{F0: 202}
        for fi37 := 0; fi37 < 13; fi37++ {
            v38 := GenS3{F0: 1}
            var v39 float32 = 1.5
            v5 += 1
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = v36
    }
    fmt.Println(v1, v3)
    for fi40 := 0; fi40 < 9; fi40++ {
        v3 = 1
        var v41 uint = 1
        _ = fi40
        _ = v41
    }
    if (4294967295 > 4294967295) {
        var v42 uint8 = 181
        switch "hello" {
        case "hello":
            fmt.Println(v4)
        case "a":
            var v43 bool = (("a" > "foo bar") || true)
            _ = v43
        case "\n\t":
            v44 := []int{42}
            _ = v44
        }
        _ = v42
    }
    v5 = int16(v4)
    v5 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
