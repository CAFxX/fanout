package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    if true {
        v4 := GenS1{F0: 255, F1: -679}
        if (32767 >= -861) {
            switch "a" {
            case "\n\t":
                if (!(42 != 42)) {
                    v3 = (float64((1e10 + 3.14159)) * (1e10 - (v3 * -2.25)))
                } else {
                    v5 := GenS2{F0: 517}
                    for fi6 := 0; fi6 < 19; fi6++ {
                        v7 := GenS1{F0: 2, F1: 42}
                        var v8 uint = 180
                        var v9 uint8 = 2
                        _ = fi6
                        _ = v7
                        _ = v8
                        _ = v9
                    }
                    var v10 uint = 42
                    _ = v10
                    _ = v5
                }
            case "":
                v3 = (v3 * float64((1e10 / -2.25)))
            default:
                var v11 int = 127
                _ = v11
            }
            switch 4294967295 {
            case 1:
                v3 = v3
            case 387:
                fmt.Println(v3)
            }
        }
        var v12 uint32 = 4294967295
        _ = v12
        _ = v4
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint64 = 4294967295
            fmt.Println(fi13)
            _ = fi13
            _ = v14
        }
    }
    if false {
        v15 := []int{32768}
        _ = v15
    } else {
        var v16 float64 = v3
        var v17 uint16 = 347
        switch 128 {
        case -2147483648:
            v16 = v16
        }
        _ = v16
        _ = v17
    }
    var v18 int8 = 100
    var v19 uint64 = 678
    var v20 uint64 = v19
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS3{F0: 42, F1: -1}
            v24 := []int{fi21, (fi22 | 2147483647), fi21, -2147483648}
            var v25 uint = 1
            fmt.Println(fi22)
            var v26 bool = (true || true)
            v27 := GenS1{F0: (v25 & v25), F1: 1000}
            fmt.Println(v18)
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi21
    }
    if false {
        v28 := []int{4294967295}
        _ = v28
    } else {
        fmt.Println(v19)
        var v29 int64 = -824
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 uint8 = 255
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    var v32 int32 = 158
    v33 := GenS4{F0: 104, F1: 274}
    v34 := GenS5{F0: (v18 - v18), F1: (v20 % ((415) | 1)), F2: 33}
    fmt.Println(v3)
    v18 = v18
    v3 += (v3 - v3)
    fmt.Println(v19)
    _ = len(v2)
    var v35 string = ("a" + "")
    if true {
        var v36 uint64 = v19
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
