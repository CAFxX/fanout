package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }

func main() {
    var v1 uint8 = 146
    var v2 uint16 = 2
    v3 := make(map[uint32]int)
    var v4 uint32 = 2
    v5 := []int{}
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint = 65535
        _ = fi7
        _ = v8
    }
    var v9 bool = ((0 != 2147483647) || (v2 == v2))
    var v10 int8 = -6
    fmt.Println(v1, v4, v4)
    fmt.Println(v2, v4)
    var v11 uint16 = v2
    var v12 uint16 = v2
    if true {
        if (v9 || false) {
            v11 += (v12 % ((v12) | 1))
            for fi13 := 0; fi13 < 3; fi13++ {
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(v2)
                    fmt.Println(v2, fi14, v12)
                    _ = fi14
                }
                v15 := GenS1{F0: 1}
                var v16 int32 = 2147483647
                _ = fi13
                _ = v15
                _ = v16
            }
        } else {
            v9 = (v10 >= (v10 & v10))
        }
        v4 -= 974
    }
    v9 = ("foo bar" < "hello")
    var v17 uint32 = 135
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS1{F0: 42}
            _ = fi18
            _ = v19
        }
        var v20 string = "\n\t"
        var v21 int8 = (v10 % (((^(v10))) | 1))
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 97}
    v23 := GenS2{F0: (v4 % (((v17 >> ((300) & 63))) | 1)), F1: "a", F2: "a"}
    switch 32767 {
    case -268:
        var v24 int8 = int8(v1)
        _ = v24
    case 0:
        if (2147483647 <= -651) {
            switch -2147483648 {
            case 179:
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 int64 = -1
                    for fi27 := 0; fi27 < 19; fi27++ {
                        fmt.Println(v12, fi27)
                        switch fi25 {
                        case -729:
                            var v28 float64 = 0.1
                            _ = v28
                        case 1:
                            v29 := GenS3{F0: 805, F1: 2, F2: (v26 >> ((930) & 63))}
                            _ = v29
                        default:
                            var v30 uint32 = 65535
                            _ = v30
                        }
                        _ = fi27
                    }
                    if (9223372036854775807 >= (fi25 + fi25)) {
                        if (v9 && v9) {
                            switch fi25 {
                            case 172:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            default:
                                var v31 int32 = 738
                                _ = v31
                            }
                        } else {
                            v32 := GenS4{F0: (v12 + (v12 * (v12 / (((v12 | v11)) | 1)))), F1: 1, F2: 2147483647}
                            _ = v32
                        }
                    }
                    _ = fi25
                    _ = v26
                }
            case -1:
                for fi33 := 0; fi33 < 13; fi33++ {
                    v2 += 255
                    fmt.Println(fi33)
                    _ = fi33
                }
            default:
                v10 = v10
            }
            v11 -= (v11 ^ v12)
            var v34 int32 = 2147483647
            _ = v34
        } else {
            v2 += (v12 - 0)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
