package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []int{}
    var v2 uint16 = 1
    var v3 int16 = -32768
    var v4 int64 = 1
    fmt.Println(v2, v4, v3)
    for fi5 := 0; fi5 < 16; fi5++ {
        if (uint64(v2) < 2) {
            if len(v1) > 0 {
                v1[0] = fi5
            }
            v4 -= -1
            v1 = append(v1, fi5)
        }
        for fi6 := 0; fi6 < 18; fi6++ {
            if (false && false) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (true && (!(-128 > 0))) {
                    var v7 int = (fi6 & -220)
                    var v8 string = ("" + "")
                    var v9 int = v7
                    _ = v7
                    _ = v8
                    _ = v9
                }
                if len(v1) > 0 {
                    v1[0] = -168
                }
            }
            v10 := GenS1{F0: 60}
            var v11 uint = 42
            fmt.Println(v11, v2)
            v12 := GenS1{F0: 255}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v13 float32 = -2.25
                var v14 uint64 = 0
                _ = v13
                _ = v14
            } else {
                var v15 int32 = 0
                var v16 int = -719
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v17 float64 = (1e10 + -2.25)
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi6
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi5
    }
    if len(v1) > 0 {
        v1[0] = -552
    }
    if (!((true || true) && false)) {
        v18 := GenS1{F0: 185}
        _ = v18
    }
    if true {
        v19 := GenS1{F0: 2}
        _ = len(v1)
        _ = v19
    } else {
        switch -9223372036854775808 {
        case 127:
            v20 := GenS1{F0: 2}
            _ = v20
        case 1:
            fmt.Println(v2)
        }
        switch 0 {
        case 359:
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS2{F0: v3, F1: 255}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi21
                _ = v22
            }
        }
    }
    v4 = (v4 | (v4 * v4))
    v23 := GenS3{F0: -1}
    v4 -= v4
    fmt.Println(v2)
    var v24 int16 = (v3 / ((v3) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
