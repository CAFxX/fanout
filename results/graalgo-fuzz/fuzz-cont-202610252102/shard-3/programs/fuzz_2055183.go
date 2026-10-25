package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int; F2 int64 }

func main() {
    v1 := []int{0}
    var v2 uint32 = 688
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 42, F1: 2147483648}
    v2 = 621
    v2 = v2
    if false {
        if (0 < -2147483648) {
            switch 17 {
            case 1:
                for fi5 := 0; fi5 < 14; fi5++ {
                    v2 = v2
                    if (-32768 != 1000) {
                        v2 = v2
                    }
                    for fi6 := 0; fi6 < 1; fi6++ {
                        v2 = v2
                        v2 -= v2
                        _ = len(v1)
                        v7 := GenS2{F0: 0, F1: uint(v2), F2: 486}
                        v8 := GenS3{F0: 903}
                        _ = fi6
                        _ = v7
                        _ = v8
                    }
                    _ = fi5
                }
            case 2147483647:
                if false {
                    v2 -= 820
                } else {
                    fmt.Println(v2)
                    v9 := GenS1{F0: -715, F1: v2}
                    _ = v9
                }
            case 240:
                fmt.Println(v2)
            default:
                switch ("a" + "") {
                case "foo bar":
                    fmt.Println(v2)
                case "hello":
                    v10 := GenS4{F0: (0 <= 32767), F1: 128, F2: -1}
                    _ = v10
                default:
                    if false {
                        if ("" > "hello") {
                            v1 = append(v1, 128)
                            var v11 uint16 = 65535
                            var v12 int8 = -90
                            _ = v11
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 14; fi13++ {
                                var v14 int8 = 0
                                var v15 int16 = -118
                                fmt.Println(v14, fi13)
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        }
                    }
                }
            }
        }
        var v16 float32 = (3.14159 - (1.5 + 0.0))
        _ = v16
    }
    var v17 uint = 353
    v2 += (v2 / (((v2 / (((v2 & 4294967295)) | 1))) | 1))
    var v18 uint8 = 127
    switch 128 {
    case -771:
        if false {
            switch "" {
            case "":
                v19 := GenS2{F0: 4294967295, F1: 152, F2: (^(v17))}
                _ = v19
            default:
                fmt.Println(v18)
            }
        } else {
            v2 = v2
        }
    default:
        for fi20 := 0; fi20 < 10; fi20++ {
            if (false && (true || ("a" != "\n\t"))) {
                delete(v3, (v2 + (v2 - (^(v2)))))
                v2 = v2
                var v21 uint32 = 2147483647
                _ = v21
            } else {
                var v22 float32 = 1.5
                _ = len(v3)
                _ = v22
            }
            _ = fi20
        }
    }
    var v23 int32 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v24 uint8 = (v18 + uint8(v17))
    fmt.Println(v2, v18)
    if true {
        v18 = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
