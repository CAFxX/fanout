package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint8; F1 int8 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = false
    v3 := make(map[int]string)
    var v4 string = (("hello" + "\n\t") + ("foo bar" + "hello"))
    if v2 {
        v2 = (v2 || v2)
        var v5 int = 32767
        v6 := GenS1{F0: 179, F1: 475}
        _ = v5
        _ = v6
    }
    var v7 int32 = 30
    var v8 uint16 = 32767
    var v9 int16 = 1
    v7 += v7
    fmt.Println(v8)
    var v10 uint8 = 1
    if v2 {
        v11 := v3[395]
        v10 = v10
        var v12 int16 = -32768
        _ = v11
        _ = v12
    } else {
        v13 := v3[1]
        if v2 {
            if (v10 == uint8(v8)) {
                var v14 uint16 = (v8 + v8)
                for fi15 := 0; fi15 < 7; fi15++ {
                    v10 -= (v10 % ((1) | 1))
                    for fi16 := 0; fi16 < 3; fi16++ {
                        fmt.Println(v14)
                        switch fi16 {
                        case 127:
                            v4 = v13
                        case 546:
                            if (int(v1) != (-(fi15))) {
                                v17 := GenS2{F0: v7, F1: -1}
                                _ = v17
                            } else {
                                v18 := GenS1{F0: 1, F1: -1}
                                _ = v18
                            }
                            v4 = ("hello" + v13)
                            var v19 uint16 = (v8 + 510)
                            _ = v19
                        case 0:
                            switch "a" {
                            case "a":
                                v20 := GenS2{F0: (v7 & int32(fi15)), F1: -112}
                                _ = v20
                            default:
                                delete(v3, (fi15 * (-(fi16))))
                            }
                        default:
                            var v21 int64 = v1
                            _ = v21
                        }
                        v9 += (v9 ^ (v9 + (^(v9))))
                        fmt.Println(v1, v14, fi16)
                        _ = fi16
                    }
                    _ = fi15
                }
                _ = v14
            } else {
                v22 := []int{0}
                var v23 int16 = (^(v9))
                _ = v22
                _ = v23
            }
            switch v13 {
            case "a":
                v24 := v3[int(v10)]
                _ = v24
            case "":
                if v2 {
                    if (v13 <= (v4 + v4)) {
                        v25 := []int{0, 2147483647, 9223372036854775807, -649}
                        if true {
                            v4 = "hello"
                            for fi26 := 0; fi26 < 4; fi26++ {
                                var v27 uint32 = 42
                                _ = fi26
                                _ = v27
                            }
                        }
                        _ = v25
                    } else {
                        var v28 string = v4
                        _ = v28
                    }
                    _ = len(v3)
                } else {
                    _ = len(v3)
                }
            }
        } else {
            v4 = (v13 + ((v13 + v13) + "hello"))
        }
        switch 42 {
        case 1:
            for fi29 := 0; fi29 < 8; fi29++ {
                switch 1 {
                case 42:
                    fmt.Println(v7, v13, v4)
                case 32768:
                    v30 := GenS3{F0: v10, F1: 0}
                    _ = v30
                case -520:
                    v1 = (v1 + 0)
                }
                _ = fi29
            }
        case 194:
            fmt.Println(v1, v10, v13)
        }
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
