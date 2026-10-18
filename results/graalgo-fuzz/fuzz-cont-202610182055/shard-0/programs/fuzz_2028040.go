package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() uint8 {
    var v7 int16 = -32768
    v7 = 42
    var v8 int16 = v7
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 bool = (255 > 255)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if ((42 >= 32767) && (2147483647 != 297)) {
            var v12 uint64 = 0
            v12 = v12
            _ = v12
        }
        v7 = (v7 * -32768)
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint8 = 128
            _ = fi13
            _ = v14
        }
        _ = fi11
    }
    _ = v7
    _ = v8
    return 175
}

func main() {
    var v1 uint64 = 796
    v2 := []int{-2147483648, 0, -146, 1}
    var v3 int32 = 1
    var v4 int64 = 0
    v5 := make(map[int]bool)
    v6 := make(map[uint32]bool)
    v4 = v4
    switch 9223372036854775807 {
    case -54:
        switch "hello" {
        case "\n\t":
            v15 := GenS1{F0: (-2.25 < ((1e10 * ((0.0 - -2.25) + 1e10)) - 3.14159)), F1: 1}
            _ = v15
        default:
            v16 := GenS1{F0: true, F1: 28}
            _ = v16
        }
    case 32767:
        if len(v2) > 0 {
            v2[0] = 214
        }
    case 127:
        if (-2.25 < 0.0) {
            fmt.Println(v4, v3, v1)
        }
    }
    v3 -= 0
    v4 += v4
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint32 = 758
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            switch -2147483648 {
            case -2147483648:
                var v21 int8 = -1
                _ = v21
            case 9223372036854775807:
                var v22 string = "a"
                _ = v22
            }
            if false {
                switch fi20 {
                case 128:
                    if (!false) {
                        v4 += v4
                    }
                case 991:
                    var v23 uint8 = 2
                    _ = v23
                }
                if false {
                    var v24 float32 = 0.1
                    if false {
                        v25 := GenS1{F0: (-32768 == 42), F1: 617}
                        var v26 int8 = 42
                        var v27 float32 = (v24 / ((((v24 * ((v24 * 1.5) / (v24 - -2.25))) + ((v24 * (1.5 / v24)) * (v24 * (0.0 + 1e10)))) + (((v24 + ((0.0 / (((v24 * float32(((-2.25 * (1e10 - (((1.5 + 0.0) - 1.5) * float64(1e10)))) + (1e10 + ((1e10 + -2.25) / (((0.1 * 1.5) + 0.1) + float64(float32(0.0)))))))) - v24) / v24)) * v24)) * (v24 / v24)) * v24)) / (v24 - ((v24 + v24) - v24))))
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    var v28 int64 = v4
                    _ = v24
                    _ = v28
                }
                if (!true) {
                    v3 += (v3 + (^(v3)))
                    if ((v4 >> ((1) & 63)) > v4) {
                        switch "a" {
                        case "hello":
                            var v29 float64 = 0.1
                            _ = v29
                        }
                        var v30 int8 = 0
                        fmt.Println(fi19, v3, fi20)
                        _ = v30
                    }
                    switch "a" {
                    case "\n\t":
                        var v31 uint8 = 28
                        _ = v31
                    }
                } else {
                    var v32 float32 = ((1.5 - 0.1) * 1e10)
                    v33 := GenS2{F0: 32767, F1: 677}
                    _ = v32
                    _ = v33
                }
            }
            var v34 int32 = v3
            _ = fi20
            _ = v34
        }
        fmt.Println(fi19, v3, v4)
        var v35 int32 = v3
        _ = fi19
        _ = v35
    }
    v36 := GenS1{F0: (-117 < 535), F1: 1}
    v37 := GenS1{F0: false, F1: 873}
    v38 := GenS1{F0: (!(65535 == 2)), F1: 66}
    v39 := GenS1{F0: false, F1: 2}
    if (1000 <= 1) {
        if true {
            v3 += v3
        }
        fmt.Println(v4, v3, v3)
        for fi40 := 0; fi40 < 12; fi40++ {
            v41 := GenS1{F0: false, F1: 32767}
            var v42 bool = true
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
