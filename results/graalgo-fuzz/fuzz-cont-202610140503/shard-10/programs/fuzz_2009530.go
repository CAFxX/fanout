package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int16; F2 uint }

func main() {
    var v1 int8 = 127
    v2 := make(map[int]int)
    var v3 int16 = -1
    var v4 uint16 = 1
    v5 := make(map[uint32]string)
    if ((!((!((-669 < 127) || false)) && true)) || true) {
        var v6 int = 618
        _ = v6
    }
    var v7 uint = 1
    v7 = v7
    v8 := GenS1{F0: 890, F1: v4, F2: 231}
    v1 = (v1 * v1)
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v3 = (v3 % (((v3 * v3)) | 1))
            switch 2147483647 {
            case 127:
                v3 += 42
            }
            delete(v2, 4294967295)
            v7 -= (v7 & (v7 | v7))
            _ = fi10
        }
        v7 = (v7 * 2)
        _ = fi9
    }
    switch "\n\t" {
    case "a":
        v7 = v7
    default:
        v11 := GenS2{F0: 1, F1: -913}
        _ = v11
    }
    if true {
        if true {
            fmt.Println(v3)
        } else {
            var v12 int64 = -128
            var v13 uint16 = (^(v4))
            v5[65535] = "foo bar"
            _ = v12
            _ = v13
        }
        var v14 uint = 2
        v7 += v14
        _ = v14
    } else {
        var v15 int16 = v3
        _ = v15
    }
    var v16 float64 = 0.0
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v3)
        for fi18 := 0; fi18 < 20; fi18++ {
            switch fi18 {
            case 2147483647:
                var v19 int32 = 42
                _ = v19
            case 301:
                var v20 float32 = float32((v16 + v16))
                _ = v20
            }
            _ = fi18
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS1{F0: 0, F1: 2, F2: 1}
            if (true && (fi21 <= fi17)) {
                switch 42 {
                case 267:
                    var v23 uint16 = (v4 << ((255) & 63))
                    _ = v23
                case 42:
                    delete(v5, 42)
                case 127:
                    var v24 uint16 = 42
                    _ = v24
                default:
                    v25 := GenS1{F0: 2147483647, F1: v4, F2: 42}
                    _ = v25
                }
                v26 := []int{(^(fi21))}
                _ = v26
            }
            v16 += -2.25
            _ = fi21
            _ = v22
        }
        _ = fi17
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        _ = len(v2)
        switch fi27 {
        case 128:
            switch (fi27 ^ 0) {
            case -1:
                v7 += v7
            case 0:
                v4 += v4
            case 128:
                _ = len(v5)
            default:
                if false {
                    v28 := GenS3{F0: uint16(v7)}
                    var v29 uint8 = 127
                    v3 -= v3
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS4{F0: (v3 / ((int16(v7)) | 1)), F1: -1, F2: v7}
                    v2[fi27] = fi27
                    var v31 float64 = (v16 + (v16 + 0.0))
                    _ = v30
                    _ = v31
                }
            }
        }
        v16 += v16
        _ = fi27
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 uint16 = v4
        if false {
            fmt.Println(fi32, v3, v7)
            var v34 int = (fi32 & -971)
            _ = v34
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v7)
    var v35 uint16 = 626
    var v36 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
