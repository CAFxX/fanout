package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int8; F2 bool }
func genFn0() uint64 {
    var v4 int8 = 1
    v4 = v4
    if (!false) {
        if false {
            var v5 int32 = 0
            _ = v5
        }
        if true {
            if (1 == 42) {
                v6 := []int{}
                switch "foo bar" {
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "a":
                    v7 := GenS1{F0: 0}
                    _ = v7
                case "\n\t":
                    v4 += 122
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = v6
            } else {
                var v8 string = "\n\t"
                _ = v8
            }
        } else {
            var v9 uint = 42
            _ = v9
        }
        if false {
            if (false || true) {
                fmt.Println(v4)
                for fi10 := 0; fi10 < 16; fi10++ {
                    if true {
                        v4 = v4
                    }
                    v11 := []int{-1, fi10, (fi10 & (fi10 << ((42) & 63))), 127}
                    _ = fi10
                    _ = v11
                }
                v12 := []int{int(v4), 32768}
                _ = v12
            }
            v13 := GenS2{F0: 128, F1: 127}
            if false {
                v14 := []int{127, 278}
                _ = v14
            }
            _ = v13
        }
    }
    _ = v4
    return 42
}

func main() {
    var v1 string = ""
    var v2 uint8 = 1
    v3 := make(map[string]uint32)
    switch 127 {
    case 128:
        v15 := GenS2{F0: 0, F1: 0}
        _ = v15
    case 107:
        v16 := GenS3{F0: -57}
        _ = v16
    case -2147483648:
        var v17 int8 = 1
        _ = v17
    default:
        if true {
            v2 -= (v2 << ((2) & 63))
        } else {
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS3{F0: 1}
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := GenS1{F0: 1}
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
            switch 32768 {
            case 2147483647:
                v22 := []int{-707, 1, 4294967295}
                _ = v22
            }
            for fi23 := 0; fi23 < 6; fi23++ {
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 float32 = 0.1
                    v26 := GenS3{F0: -1}
                    v27 := GenS1{F0: 917}
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                var v28 uint32 = 200
                _ = fi23
                _ = v28
            }
        }
    }
    v29 := []int{214, 980, 640, 311}
    if (true || (v1 <= v1)) {
        fmt.Println(v2)
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := GenS4{F0: 0, F1: 127, F2: (0 != -13)}
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 bool = (1 < 1)
            var v34 int8 = -1
            var v35 int8 = (v34 & v34)
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            v37 := GenS4{F0: 68, F1: 100, F2: (v2 == 42)}
            _ = fi36
            _ = v37
        }
        _ = fi30
        _ = v31
    }
    var v38 uint = 977
    _ = len(v29)
    var v39 int64 = -944
    var v40 float64 = float64((0.0 / 1e10))
    for fi41 := 0; fi41 < 11; fi41++ {
        v42 := GenS4{F0: -67, F1: 42, F2: false}
        v39 = -9223372036854775808
        _ = fi41
        _ = v42
    }
    v40 = (0.1 / v40)
    var v43 float32 = (1e10 - 0.0)
    switch -2147483648 {
    case 42:
        v2 = 189
    default:
        delete(v3, (v1 + v1))
    }
    fmt.Println(v1, v38, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v43
}
