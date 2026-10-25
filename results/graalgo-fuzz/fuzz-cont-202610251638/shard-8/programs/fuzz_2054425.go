package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int }
func genFn0() uint64 {
    var v7 int64 = 1
    var v8 uint = 4294967295
    var v9 uint = 304
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint32, p1 uint16) int {
    p1 -= 456
    var v10 int64 = 625
    var v11 uint32 = p0
    var v12 uint32 = (p0 - p0)
    fmt.Println(v11)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 107
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 uint32 = 42
    var v5 float64 = -2.25
    var v6 string = "foo bar"
    fmt.Println(v1, v5, v4)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int = (v1 >> ((0) & 63))
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 int16 = 32767
            switch v6 {
            case "":
                fmt.Println(fi15, v1)
            default:
                var v17 int64 = 9223372036854775807
                _ = v17
            }
            v18 := GenS1{F0: genFn0(), F1: 1}
            var v19 int8 = 0
            v20 := GenS1{F0: 1, F1: -9223372036854775808}
            if (!true) {
                fmt.Println(v4)
                v21 := GenS2{F0: (v4 & v4), F1: 0, F2: 2}
                if (0 < genFn0()) {
                    var v22 int16 = (^(v16))
                    switch v6 {
                    case "\n\t":
                        if true {
                            v23 := GenS1{F0: 958, F1: 9223372036854775807}
                            _ = v23
                        } else {
                            var v24 uint32 = (v4 / (((^(v4))) | 1))
                            _ = v24
                        }
                    case "foo bar":
                        if (0 >= 128) {
                            var v25 uint16 = 180
                            v26 := []int{genFn1((v4 / (((v4 - v4)) | 1)), (^(v25))), fi13, (v1 | (fi15 >> ((284) & 63))), (fi13 * 246)}
                            if (true && false) {
                                var v27 uint64 = 4294967295
                                if false {
                                    var v28 uint64 = 42
                                    var v29 uint16 = 540
                                    _ = v28
                                    _ = v29
                                }
                                _ = v27
                            }
                            var v30 int16 = (v16 & v22)
                            delete(v3, (fi15 * (v14 & (v14 + v14))))
                            _ = v25
                            _ = v26
                            _ = v30
                        }
                    }
                    v4 += (v4 - v4)
                    _ = v22
                }
                var v31 int32 = -462
                if true {
                    v32 := GenS3{F0: -128}
                    _ = v32
                }
                _ = v21
                _ = v31
            }
            _ = fi15
            _ = v16
            _ = v18
            _ = v19
            _ = v20
        }
        var v33 bool = true
        _ = fi13
        _ = v14
        _ = v33
    }
    v34 := GenS4{F0: -1, F1: -1}
    if (1000 <= 0) {
        v1 -= -470
        fmt.Println(v6, v5)
        for fi35 := 0; fi35 < 19; fi35++ {
            switch (v1 % ((-9223372036854775808) | 1)) {
            case 4294967295:
                v36 := []int{32767}
                _ = v36
            case -1:
                delete(v2, (v4 / ((v4) | 1)))
            }
            v37 := GenS4{F0: -1, F1: (fi35 / ((v1) | 1))}
            _ = fi35
            _ = v37
        }
    }
    var v38 uint32 = v4
    v5 -= ((((v5 * v5) - 3.14159) * (v5 + v5)) + float64(0.0))
    var v39 uint16 = 2
    v40 := GenS2{F0: 42, F1: v39, F2: genFn0()}
    v5 += -2.25
    fmt.Println(v5)
    fmt.Println(v39)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
