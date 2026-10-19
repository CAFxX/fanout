package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int64; F2 int64 }
func genFn0() bool {
    switch 127 {
    case 42:
        var v4 uint16 = 1
        _ = v4
    case 235:
        switch ("foo bar" + (("foo bar" + "") + "hello")) {
        case "\n\t":
            for fi5 := 0; fi5 < 13; fi5++ {
                if true {
                    if (255 >= 0) {
                        v6 := GenS1{F0: 1, F1: "foo bar"}
                        v7 := []int{}
                        _ = v6
                        _ = v7
                    } else {
                        fmt.Println(fi5)
                    }
                    var v8 uint32 = 197
                    _ = v8
                }
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    v9 := GenS2{F0: 32767}
                    _ = v9
                }
                _ = fi5
            }
        default:
            var v10 string = "a"
            _ = v10
        }
    case 128:
        switch 912 {
        case -2147483648:
            var v11 bool = (!false)
            _ = v11
        default:
            var v12 int64 = 42
            _ = v12
        }
    }
    var v13 int8 = 42
    v13 -= 0
    _ = v13
    return true
}
func genFn1(p0 int) uint64 {
    var v14 float32 = 0.0
    if true {
        p0 += p0
        for fi15 := 0; fi15 < 15; fi15++ {
            v14 = (v14 + (((1.5 * (1.5 / float32((1e10 * 1.5)))) + ((0.1 / 1.0) + v14)) - (3.14159 / (float32(float64((((float32(0.0) * v14) - 0.0) + v14))) / (v14 - (v14 / v14))))))
            _ = fi15
        }
    }
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v16 := GenS3{F0: ((1e10 - 1e10) <= 1.5), F1: 2, F2: 0}
    v1 -= v1
    switch 9223372036854775807 {
    case 533:
        if false {
            var v17 int16 = 1
            _ = v17
        } else {
            var v18 float32 = 3.14159
            v19 := GenS4{F0: (v1 / (((v1 + v1)) | 1)), F1: 127}
            _ = v18
            _ = v19
        }
    case 39:
        for fi20 := 0; fi20 < 13; fi20++ {
            v1 -= -1
            var v21 int32 = int32(v1)
            v22 := GenS2{F0: 255}
            _ = fi20
            _ = v21
            _ = v22
        }
    case -1:
        if (127 > (v1 & v1)) {
            fmt.Println(v1)
            if true {
                v23 := []int{4294967295, 4294967295, 2147483647}
                _ = v23
            } else {
                var v24 int64 = 203
                var v25 uint8 = 127
                var v26 float64 = 3.14159
                _ = v24
                _ = v25
                _ = v26
            }
            v27 := []int{32767, -908, 2147483647, -1}
            _ = v27
        } else {
            var v28 float32 = -2.25
            var v29 uint = 0
            var v30 int32 = -514
            _ = v28
            _ = v29
            _ = v30
        }
    }
    v31 := GenS5{F0: 4294967295, F1: 1, F2: 1}
    var v32 float64 = (1e10 * (((3.14159 - -2.25) + float64(1e10)) - 0.0))
    v1 -= (v1 << ((65535) & 63))
    var v33 int32 = -2147483648
    if (false || (!((3.14159 - -2.25) != float32((((float64(0.0) + 1.5) * (v32 * v32)) - v32))))) {
        v34 := GenS3{F0: true, F1: 255, F2: 255}
        fmt.Println(v1, v32)
        v2[729] = 128
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
