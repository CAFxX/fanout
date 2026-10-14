package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 197
    v3 := []int{32768, -16}
    if false {
        v2 -= v2
        if true {
            var v4 int8 = -1
            var v5 int = 42
            if false {
                v4 = (v4 % (((v4 | 1)) | 1))
                for fi6 := 0; fi6 < 16; fi6++ {
                    v1["hello"] = 65535
                    switch -223 {
                    case 2147483647:
                        var v7 uint64 = 0
                        _ = v7
                    case 436:
                        fmt.Println(v5)
                    case 0:
                        fmt.Println(v2)
                    }
                    _ = fi6
                }
                if true {
                    var v8 int = (-(v5))
                    switch v5 {
                    case 32767:
                        v2 = (v2 << ((0) & 63))
                    default:
                        var v9 uint32 = 370
                        _ = v9
                    }
                    var v10 uint = 388
                    _ = v10
                    _ = v8
                } else {
                    _ = len(v3)
                    fmt.Println(v4, v5)
                    fmt.Println(v5, v2)
                }
            }
            _ = v4
            _ = v5
        } else {
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(fi11, v2)
                fmt.Println(v2, v2)
                _ = fi11
            }
        }
        _ = len(v3)
    }
    var v12 string = ""
    if false {
        v13 := GenS1{F0: -1, F1: true}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v13
    } else {
        var v14 uint8 = 0
        _ = v14
    }
    var v15 float64 = ((0.0 - -2.25) / ((0.1 + ((((0.1 + 1.5) * 3.14159) + 3.14159) * 1e10)) / float64(1.5)))
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int32 = 307
        fmt.Println(v17)
        if ((^(v2)) < (v2 % ((128) | 1))) {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int64 = 1
                var v20 string = (v12 + (((v12 + v12) + v12) + v12))
                var v21 int16 = 831
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v12, v12, v2)
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS2{F0: 1}
        for fi24 := 0; fi24 < 20; fi24++ {
            v1["a"] = 42
            v25 := GenS2{F0: 1}
            v26 := GenS3{F0: 1, F1: (fi24 + fi22)}
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v27 uint64 = 42
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := GenS1{F0: -1, F1: (4294967295 <= 914)}
            var v30 uint64 = (v27 | 4294967295)
            v2 = 2
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = v27
    } else {
        var v31 int8 = 62
        var v32 int = 2147483647
        _ = v31
        _ = v32
    }
    v33 := v1[v12]
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v34 := v1[(v12 + v12)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
