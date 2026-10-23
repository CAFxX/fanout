package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 bool }

func main() {
    var v1 uint8 = 1
    var v2 float64 = -2.25
    var v3 int16 = 1
    v4 := []uint32{0}
    switch 4294967295 {
    case 32767:
        var v5 int = -1
        _ = v5
    default:
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 float32 = (3.14159 * 0.1)
            _ = fi6
            _ = v7
        }
    }
    var v8 int8 = 100
    v1 = (v1 << ((582) & 63))
    var v9 string = "\n\t"
    v10 := GenS1{F0: 539, F1: true, F2: 532}
    switch "" {
    case "\n\t":
        fmt.Println(v2)
    case "":
        v4 = append(v4, 1)
    case "a":
        if (!false) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                switch fi11 {
                case 42:
                    var v12 int = fi11
                    _ = v12
                case 2147483647:
                    if (v8 > 127) {
                        var v13 int32 = 42
                        var v14 uint = 4294967295
                        for fi15 := 0; fi15 < 12; fi15++ {
                            if len(v4) > 0 {
                                v4[0] = 2147483648
                            }
                            var v16 string = ""
                            var v17 float32 = 1.5
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        _ = v13
                        _ = v14
                    }
                }
                if true {
                    fmt.Println(v3, v9, v9)
                }
                _ = fi11
            }
            var v18 int32 = 313
            _ = v18
        }
    }
    v19 := GenS2{F0: 42}
    fmt.Println(v8, v3)
    fmt.Println(v3, v9)
    v1 = (^(v1))
    switch 2147483647 {
    case -995:
        var v20 bool = true
        _ = v20
    case 2147483647:
        switch 128 {
        case 127:
            switch 529 {
            case -190:
                v3 += v3
            case -756:
                if false {
                    fmt.Println(v9, v8)
                } else {
                    var v21 int64 = 1
                    v22 := GenS3{F0: "", F1: true}
                    _ = v21
                    _ = v22
                }
            default:
                v23 := GenS1{F0: 2, F1: (false && (-1 != -9223372036854775808)), F2: 345}
                _ = v23
            }
        }
    case 983:
        v24 := GenS3{F0: (("hello" + ("a" + (v9 + ("foo bar" + (v9 + "foo bar"))))) + v9), F1: false}
        _ = v24
    default:
        var v25 int32 = 259
        _ = v25
    }
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
