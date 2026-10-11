package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint32 }
func genFn0(p0 int8) int32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 -= p0
        _ = fi7
    }
    switch 96 {
    case -1:
        switch ("hello" + ("foo bar" + "\n\t")) {
        case "hello":
            p0 += (p0 ^ 127)
        case "":
            var v8 uint16 = 32767
            _ = v8
        default:
            switch "\n\t" {
            case "foo bar":
                switch ("\n\t" + (("" + ("hello" + "hello")) + "foo bar")) {
                case "hello":
                    var v9 int32 = 157
                    _ = v9
                default:
                    v10 := []int{9223372036854775807, 495, 4294967295}
                    _ = v10
                }
            case "hello":
                var v11 bool = true
                _ = v11
            case "a":
                if (!(-563 >= 127)) {
                    p0 = p0
                    if (1 != -1) {
                        if false {
                            v12 := GenS1{F0: 1, F1: 0}
                            if true {
                                p0 = (p0 - p0)
                            } else {
                                p0 += (p0 & -1)
                                p0 += (p0 + (p0 % ((120) | 1)))
                                p0 = p0
                            }
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v14 := GenS1{F0: 0, F1: 65535}
                            switch fi13 {
                            case 32768:
                                for fi15 := 0; fi15 < 13; fi15++ {
                                    v16 := GenS2{F0: p0}
                                    if true {
                                        if true {
                                            p0 += 1
                                        } else {
                                            p0 = (p0 ^ 100)
                                            var v17 uint64 = 0
                                            _ = v17
                                        }
                                    } else {
                                        var v18 int32 = 42
                                        _ = v18
                                    }
                                    var v19 int8 = 42
                                    _ = fi15
                                    _ = v16
                                    _ = v19
                                }
                            }
                            _ = fi13
                            _ = v14
                        }
                        if (128 > 212) {
                            p0 = (p0 % (((p0 * p0)) | 1))
                        }
                    }
                }
            }
        }
    case 9223372036854775807:
        p0 -= (-(p0))
    }
    p0 = (p0 / ((p0) | 1))
    v20 := []int{4294967295}
    _ = p0
    _ = v20
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 int = 357
    v4 := []int{2147483647, -70, (v3 ^ v3)}
    var v5 float32 = 0.0
    var v6 uint8 = 0
    _ = len(v4)
    var v21 uint = 155
    _ = len(v1)
    fmt.Println(v21)
    v22 := GenS3{F0: "a"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v23 uint16 = 42
    if v2 {
        v3 = (v3 + 128)
    } else {
        v24 := GenS4{F0: 60, F1: v6, F2: 2147483647}
        v23 = v23
        v2 = false
        _ = v24
    }
    var v25 int = v3
    var v26 int = v3
    var v27 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
