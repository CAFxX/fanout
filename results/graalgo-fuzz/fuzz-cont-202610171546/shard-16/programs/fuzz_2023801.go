package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint32; F2 bool }

func main() {
    var v1 int64 = -162
    var v2 bool = false
    var v3 uint8 = 2
    var v4 uint = 2
    var v5 uint32 = 806
    switch 985 {
    case 0:
        v6 := []int{-1, 127, 0}
        _ = v6
    }
    fmt.Println(v3, v2)
    v7 := []int{}
    v3 += v3
    var v8 float32 = 3.14159
    var v9 int = -974
    for fi10 := 0; fi10 < 18; fi10++ {
        if ((true || true) && false) {
            if true {
                if (0 == -55) {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        v1 = (^(v1))
                        _ = fi11
                    }
                    v12 := GenS1{F0: 2}
                    v13 := GenS2{F0: v1, F1: 127}
                    _ = v12
                    _ = v13
                }
                if v2 {
                    fmt.Println(v4)
                    v2 = false
                    v3 += 131
                } else {
                    if len(v7) > 0 {
                        v7[0] = (v9 - (v9 << ((v4) & 63)))
                    }
                }
                var v14 float64 = (1.5 / 0.1)
                _ = v14
            } else {
                var v15 int16 = 1
                var v16 int64 = (v1 + v1)
                _ = v15
                _ = v16
            }
        } else {
            v17 := GenS1{F0: v4}
            _ = v17
        }
        _ = fi10
    }
    if v2 {
        for fi18 := 0; fi18 < 8; fi18++ {
            v5 -= (^(v5))
            var v19 float64 = 1e10
            _ = fi18
            _ = v19
        }
        var v20 uint = 2
        v4 -= v4
        _ = v20
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            if len(v7) > 0 {
                v7[0] = (v9 | -1)
            }
            v5 = 1
            v22 := GenS3{F0: 2, F1: (v5 - v5), F2: (v1 < int64(fi21))}
            _ = fi21
            _ = v22
        }
    }
    v2 = (0.1 <= 0.0)
    v8 = ((float32(3.14159) + (((v8 + v8) / 3.14159) - float32(1e10))) - float32(1.5))
    var v23 int16 = 1
    for fi24 := 0; fi24 < 3; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 uint = 576
            switch v9 {
            case 571:
                fmt.Println(v2, fi25)
            default:
                var v27 string = "foo bar"
                _ = v27
            }
            _ = fi25
            _ = v26
        }
        var v28 uint64 = 4294967295
        _ = fi24
        _ = v28
    }
    var v29 int = v9
    if true {
        v23 += -1
        v7 = append(v7, v29)
    } else {
        var v30 bool = ((!v2) || ((4294967295 <= 692) || (1 < 22)))
        if (42 != 860) {
            fmt.Println(v5, v8)
            var v31 float32 = (float32(0.1) / (v8 / (v8 - -2.25)))
            _ = v31
        } else {
            v4 += v4
            if false {
                v32 := GenS2{F0: int64(v4), F1: 0}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 bool = false
                    _ = fi33
                    _ = v34
                }
            }
            for fi35 := 0; fi35 < 16; fi35++ {
                var v36 int8 = -19
                _ = fi35
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 15; fi37++ {
            if len(v7) > 0 {
                v7[0] = (fi37 / (((v29 * 315)) | 1))
            }
            _ = fi37
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
