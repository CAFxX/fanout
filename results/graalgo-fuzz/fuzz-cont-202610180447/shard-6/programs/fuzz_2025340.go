package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint; F2 int8 }
type GenS4 struct { F0 int8 }
func genFn0() uint8 {
    switch 4294967295 {
    case -479:
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                switch fi6 {
                case -9223372036854775808:
                    v7 := []int{(-(fi6)), (fi6 ^ 0), (fi6 | fi6), 1}
                    _ = v7
                case 128:
                    switch (fi6 - fi6) {
                    case 128:
                        v8 := []int{fi6, fi6, 32767, (fi6 & (^(fi6)))}
                        _ = v8
                    default:
                        fmt.Println(fi6)
                    }
                }
                v9 := []int{}
                if true {
                    var v10 uint8 = 255
                    _ = v10
                }
                _ = fi6
                _ = v9
            }
            if false {
                var v11 int = 336
                switch v11 {
                case 995:
                    v11 = (v11 ^ 128)
                case 0:
                    v12 := []int{v11, v11}
                    _ = v12
                case 673:
                    switch v11 {
                    case -1:
                        v11 = v11
                    case 9223372036854775807:
                        if true {
                            var v13 int32 = 42
                            v14 := []int{32768}
                            v14 = append(v14, (v11 >> ((42) & 63)))
                            _ = v13
                            _ = v14
                        }
                    default:
                        for fi15 := 0; fi15 < 9; fi15++ {
                            v11 -= v11
                            for fi16 := 0; fi16 < 12; fi16++ {
                                var v17 uint8 = 211
                                v17 = v17
                                var v18 float32 = ((0.0 + 1e10) / 1.0)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            var v19 float32 = 0.0
                            _ = fi15
                            _ = v19
                        }
                    }
                }
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := GenS1{F0: 1}
                    _ = fi20
                    _ = v21
                }
                _ = v11
            } else {
                if true {
                    v22 := GenS2{F0: 420, F1: -128}
                    _ = v22
                } else {
                    var v23 uint8 = 165
                    _ = v23
                }
                var v24 uint8 = 255
                _ = v24
            }
        } else {
            v25 := []int{-91}
            _ = v25
        }
    default:
        v26 := GenS1{F0: -1}
        _ = v26
    }
    var v27 int32 = 1
    _ = v27
    return 255
}
func genFn1() uint {
    var v28 uint16 = 615
    if true {
        v28 = (v28 & v28)
        switch 42 {
        case 2147483647:
            v29 := GenS3{F0: 2, F1: 18446744073709551615, F2: 118}
            _ = v29
        case 4294967295:
            v28 -= v28
        case 1:
            v28 = 156
        }
        v30 := GenS3{F0: 1, F1: 42, F2: 100}
        _ = v30
    } else {
        if (!false) {
            fmt.Println(v28)
            v31 := []int{57, 2147483647, 364, -230}
            _ = v31
        }
    }
    var v32 bool = true
    _ = v28
    _ = v32
    return 372
}

func main() {
    var v1 uint64 = 426
    v2 := make(map[int]uint32)
    var v3 int32 = 413
    v4 := []int{}
    var v5 uint8 = 127
    v1 = v1
    v2[495] = 63
    var v33 int64 = -1
    _ = len(v2)
    var v34 string = ((("hello" + ("foo bar" + "a")) + "") + "")
    var v35 int = 127
    v36 := GenS2{F0: 65535, F1: 108}
    v37 := GenS4{F0: -67}
    v33 = (v33 | -385)
    fmt.Println(v33, v34, v33)
    if (!false) {
        var v38 float32 = 0.1
        if (v5 < v5) {
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 uint8 = v5
                switch (((v34 + "a") + ((((v34 + (v34 + ("hello" + v34))) + ((v34 + ("hello" + (v34 + v34))) + "")) + "") + (v34 + v34))) + "\n\t") {
                case "\n\t":
                    fmt.Println(v34, v35)
                }
                _ = fi39
                _ = v40
            }
            v33 = (v33 / ((v33) | 1))
        }
        _ = v38
    } else {
        if true {
            switch v35 {
            case 553:
                v3 = v3
            default:
                v35 = (v35 >> ((725) & 63))
            }
            var v41 uint = genFn1()
            _ = v41
        } else {
            v5 += genFn0()
            var v42 uint16 = 492
            _ = v42
        }
        var v43 uint16 = 32767
        _ = v43
    }
    v2[v35] = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
