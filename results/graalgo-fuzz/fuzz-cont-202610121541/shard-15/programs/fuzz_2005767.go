package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint64; F1 int32 }

func main() {
    v1 := []int{-783}
    v2 := make(map[string]string)
    var v3 string = "\n\t"
    v4 := []int{}
    v5 := make(map[int]string)
    if true {
        v3 = (v3 + v3)
        for fi6 := 0; fi6 < 6; fi6++ {
            v3 = ""
            var v7 int32 = -838
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: 4294967295}
                var v10 uint32 = 65535
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v3)
    switch "" {
    case "":
        var v11 string = (("" + ((v3 + v3) + "hello")) + (v3 + v3))
        _ = v11
    case "foo bar":
        switch -513 {
        case 9223372036854775807:
            fmt.Println(v3)
        case 764:
            if false {
                v3 = v3
                var v12 int16 = -32768
                if true {
                    switch 1 {
                    case -9223372036854775808:
                        if true {
                            fmt.Println(v3)
                        } else {
                            v13 := GenS1{F0: 632}
                            v14 := GenS2{F0: 190, F1: v12}
                            fmt.Println(v12, v3)
                            _ = v13
                            _ = v14
                        }
                    case 242:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    if (1 < 2147483647) {
                        switch -2147483648 {
                        case -2147483648:
                            v12 = 464
                        case 4294967295:
                            _ = len(v4)
                        case -307:
                            var v15 uint32 = 4294967295
                            _ = v15
                        }
                        v12 = v12
                        if len(v4) > 0 {
                            v4[0] = 0
                        }
                    }
                }
                _ = v12
            }
        case -189:
            if true {
                v1 = append(v1, 32767)
                v3 = v3
                switch -665 {
                case 32767:
                    if (-456 == 667) {
                        v3 = "foo bar"
                        var v16 uint8 = 127
                        v16 += (v16 / (((v16 - 255)) | 1))
                        _ = v16
                    } else {
                        var v17 int = -1
                        _ = v17
                    }
                case 2147483647:
                    var v18 float32 = 0.1
                    _ = v18
                case 4294967295:
                    var v19 int32 = 1
                    _ = v19
                default:
                    var v20 int8 = -1
                    _ = v20
                }
            } else {
                var v21 int32 = 42
                _ = v21
            }
        }
    case "a":
        var v22 int16 = -32768
        _ = v22
    default:
        if (0 > 42) {
            for fi23 := 0; fi23 < 13; fi23++ {
                v3 = v3
                _ = fi23
            }
        }
    }
    v24 := GenS1{F0: 613}
    var v25 string = (("hello" + (v3 + v3)) + v3)
    if true {
        if (2147483647 == 42) {
            var v26 uint16 = 2
            v3 = ("\n\t" + v25)
            var v27 uint = 340
            _ = v26
            _ = v27
        }
        v28 := GenS3{F0: 42, F1: 2147483647, F2: 65535}
        _ = v28
    }
    _ = len(v5)
    if (1 > 0) {
        v25 = v25
        var v29 float64 = -2.25
        var v30 int8 = 100
        _ = v29
        _ = v30
    } else {
        v31 := GenS4{F0: 717, F1: -654}
        var v32 uint16 = 0
        var v33 int = 128
        _ = v31
        _ = v32
        _ = v33
    }
    v34 := GenS5{F0: 42, F1: -333}
    fmt.Println(v3)
    v25 = (v3 + v25)
    if (42 <= 2) {
        v25 = v25
    }
    var v35 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
