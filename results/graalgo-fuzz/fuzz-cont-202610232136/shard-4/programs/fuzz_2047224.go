package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 181
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 int16 = 373
    var v6 int64 = 42
    delete(v4, -347)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: true, F1: true}
        v5 += v5
        _ = fi7
        _ = v8
    }
    v5 -= (v5 & 1)
    var v9 uint8 = 1
    if ((v6 ^ (^(v6))) != 1) {
        v10 := GenS2{F0: ((("hello" + "hello") + "") + ""), F1: 127}
        switch "foo bar" {
        case "\n\t":
            v9 -= 1
        default:
            var v11 uint16 = 32767
            _ = v11
        }
        _ = v10
    }
    var v12 int8 = 0
    if true {
        v5 = int16(v6)
        v13 := []int{int(v1)}
        if (!(2147483647 != -1)) {
            var v14 int = 268
            v15 := GenS1{F0: false, F1: true}
            var v16 int = (v14 % (((v14 | -1)) | 1))
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        var v17 float64 = 1.5
        _ = v17
    }
    v18 := []int{42, -481}
    fmt.Println(v1, v12)
    if true {
        for fi19 := 0; fi19 < 15; fi19++ {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            switch 128 {
            case -830:
                if (!(255 != 945)) {
                    var v20 uint8 = (v9 << ((42) & 63))
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v22 := GenS1{F0: true, F1: true}
                        _ = len(v4)
                        _ = fi21
                        _ = v22
                    }
                    v12 += v12
                    _ = v20
                } else {
                    var v23 uint16 = 0
                    if (false || true) {
                        v24 := v3[1]
                        v5 += (v5 / ((v5) | 1))
                        v25 := GenS3{F0: fi19}
                        _ = v24
                        _ = v25
                    }
                    _ = v23
                }
            case 32768:
                var v26 int32 = -1
                _ = v26
            case 42:
                var v27 int64 = (^(v6))
                _ = v27
            default:
                for fi28 := 0; fi28 < 20; fi28++ {
                    v2 = (^(v9))
                    v2 = 128
                    v1 += (v1 | v1)
                    _ = fi28
                }
            }
            if false {
                var v29 int16 = 0
                v9 -= (v2 & uint8(v1))
                var v30 string = ("\n\t" + "hello")
                _ = v29
                _ = v30
            } else {
                if true {
                    v9 -= v9
                    var v31 int64 = (-(v6))
                    var v32 int = (fi19 % ((fi19) | 1))
                    _ = v31
                    _ = v32
                } else {
                    v33 := GenS4{F0: 2147483647, F1: -2147483648}
                    _ = v33
                }
                if (!true) {
                    _ = len(v18)
                }
                var v34 int32 = int32(v5)
                _ = v34
            }
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
