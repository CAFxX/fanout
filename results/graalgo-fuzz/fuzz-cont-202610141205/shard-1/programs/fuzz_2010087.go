package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 string; F1 string; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint64; F2 int16 }

func main() {
    var v1 int8 = 127
    v2 := make(map[string]uint32)
    var v3 float64 = 1e10
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    var v6 uint16 = 1
    var v7 int8 = (v1 * (-(v1)))
    var v8 uint64 = 42
    _ = len(v2)
    v9 := []int{956}
    v10 := GenS1{F0: 32767, F1: (v6 % (((v6 ^ (v6 << ((18446744073709551615) & 63)))) | 1)), F2: -32768}
    var v11 bool = true
    v12 := v4["\n\t"]
    if v11 {
        v13 := GenS1{F0: v6, F1: v6, F2: 32767}
        _ = v13
    } else {
        if (56 == 1) {
            var v14 string = v12
            v15 := GenS1{F0: v6, F1: uint16(v8), F2: 1}
            var v16 uint = 613
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 int = 81
        _ = v17
    }
    var v18 uint16 = (v6 >> ((4294967295) & 63))
    var v19 uint = 849
    v20 := GenS2{F0: ("\n\t" + v12), F1: ((("a" + v12) + "a") + (v12 + v12)), F2: 32767}
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS3{F0: fi21}
        _ = fi21
        _ = v22
    }
    if (!(345 >= 65535)) {
        v4["a"] = ((v12 + "foo bar") + v12)
    }
    v3 += (v3 * v3)
    v23 := GenS2{F0: v12, F1: v12, F2: 918}
    if v11 {
        if v11 {
            var v24 uint = v19
            for fi25 := 0; fi25 < 10; fi25++ {
                v7 -= 100
                var v26 float32 = float32((v3 - (v3 * (float64(3.14159) / (float64(1.5) + (1.5 - v3))))))
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        v11 = true
        v27 := GenS3{F0: -661}
        _ = v27
    } else {
        if (127 > 128) {
            for fi28 := 0; fi28 < 11; fi28++ {
                if v11 {
                    var v29 int32 = 400
                    switch v12 {
                    case "\n\t":
                        fmt.Println(v11)
                    default:
                        var v30 string = (("\n\t" + (v12 + (v12 + "hello"))) + (v12 + v12))
                        _ = v30
                    }
                    if (2 < 2) {
                        _ = len(v9)
                    }
                    _ = v29
                }
                if v11 {
                    if ((2147483647 != 2) || (413 >= 65535)) {
                        v31 := v5[1]
                        var v32 int32 = 2147483647
                        switch fi28 {
                        case 42:
                            _ = len(v4)
                        case -1:
                            switch "foo bar" {
                            case "\n\t":
                                v33 := GenS4{F0: 786, F1: 0, F2: 4294967295}
                                _ = v33
                            default:
                                var v34 uint8 = uint8(v19)
                                _ = v34
                            }
                        default:
                            v35 := GenS1{F0: (v6 >> ((v19) & 63)), F1: 2, F2: 1}
                            _ = v35
                        }
                        _ = v31
                        _ = v32
                    } else {
                        fmt.Println(v1)
                        _ = len(v9)
                        v8 = v8
                    }
                    v36 := GenS2{F0: (v12 + (v12 + v12)), F1: "", F2: (v18 | (v6 + v6))}
                    _ = v36
                } else {
                    fmt.Println(v3, v6)
                    var v37 int = -336
                    v38 := GenS5{F0: -671, F1: (v8 | (v8 ^ 506)), F2: 19}
                    _ = v37
                    _ = v38
                }
                _ = fi28
            }
            var v39 uint32 = 4294967295
            _ = v39
        } else {
            var v40 bool = (v11 || false)
            switch (v12 + (("a" + "a") + v12)) {
            case "hello":
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            case "\n\t":
                v7 -= (v7 - v7)
            case "":
                v6 = 161
            }
            _ = v40
        }
        if v11 {
            var v41 int64 = 42
            _ = v41
        } else {
            v42 := GenS5{F0: -1, F1: (v8 << ((uint(v6)) & 63)), F2: -1}
            _ = v42
        }
    }
    v43 := v5[444]
    if (v6 != v18) {
        for fi44 := 0; fi44 < 20; fi44++ {
            v45 := GenS4{F0: 147, F1: 367, F2: 0}
            v12 = "hello"
            var v46 uint64 = v8
            _ = fi44
            _ = v45
            _ = v46
        }
        v7 += v1
        fmt.Println(v43, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
