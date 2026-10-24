package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint; F2 int16 }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0(p0 int8) uint32 {
    v4 := []int{}
    if (255 < 0) {
        switch "" {
        case "a":
            var v5 bool = (-2147483648 == 2147483647)
            _ = v5
        }
        var v6 int32 = 1
        var v7 uint32 = 763
        _ = v6
        _ = v7
    } else {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint32 = 637
        _ = fi8
        _ = v9
    }
    if (1e10 > 1e10) {
        v4 = append(v4, -593)
        v10 := GenS1{F0: -333, F1: 65, F2: -591}
        var v11 int8 = 42
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 float32 = ((float32(3.14159) * 3.14159) / 1.0)
    var v3 uint8 = 255
    var v12 int8 = 0
    v13 := GenS1{F0: -1, F1: v3, F2: -660}
    if (1e10 <= float64(v2)) {
        v14 := GenS2{F0: v12}
        var v15 uint = 4294967295
        _ = v14
        _ = v15
    } else {
        v1 = v1
    }
    if true {
        fmt.Println(v3, v12, v1)
    }
    v16 := []int{165, 0}
    for fi17 := 0; fi17 < 8; fi17++ {
        switch fi17 {
        case 55:
            fmt.Println(v1, fi17, v12)
        }
        _ = fi17
    }
    fmt.Println(v1, v2, v12)
    if (genFn0((v12 % ((v12) | 1))) > 1) {
        var v18 string = "a"
        for fi19 := 0; fi19 < 18; fi19++ {
            switch (v18 + v18) {
            case "hello":
                if (v12 > (v12 + (v12 % (((-(v12))) | 1)))) {
                    var v20 uint64 = uint64(v3)
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        if false {
                            switch fi19 {
                            case -757:
                                if false {
                                    fmt.Println(v18, fi19)
                                }
                            case 42:
                                var v22 int16 = -126
                                _ = v22
                            case 127:
                                if (18446744073709551615 <= 1) {
                                    var v23 uint64 = 42
                                    var v24 int8 = int8(v1)
                                    v25 := GenS3{F0: (v18 + (v18 + v18)), F1: 1, F2: 797}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            }
                        }
                        v26 := GenS4{F0: 65535, F1: v12}
                        _ = fi21
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v28 := GenS4{F0: 4294967295, F1: 68}
                        _ = fi27
                        _ = v28
                    }
                }
            default:
                var v29 int64 = 216
                _ = v29
            }
            _ = fi19
        }
        _ = v18
    }
    var v30 uint16 = 715
    _ = len(v16)
    if (false && false) {
        fmt.Println(v1)
    }
    var v31 uint32 = 2147483647
    fmt.Println(v30)
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := GenS3{F0: "", F1: 4294967295, F2: -1}
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 int64 = 1
            var v36 int64 = v35
            fmt.Println(v31)
            _ = len(v16)
            var v37 string = "hello"
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        fmt.Println(v12, v30)
        _ = fi32
        _ = v33
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 int32 = 841
        _ = fi38
        _ = v39
    }
    var v40 uint16 = (v30 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v40
}
