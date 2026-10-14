package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = uint32(v1)
    v3 := []string{"hello", "a"}
    var v4 string = (("foo bar" + "a") + "")
    v5 := make(map[uint32]string)
    var v6 string = ""
    var v7 int64 = 889
    var v8 int32 = -1
    fmt.Println(v6)
    v9 := GenS1{F0: (v7 / (((v7 ^ -1)) | 1)), F1: true, F2: 4294967295}
    fmt.Println(v2, v4)
    var v10 float32 = 1e10
    var v11 uint64 = (v1 << ((65535) & 63))
    v7 = (v7 + 0)
    for fi12 := 0; fi12 < 6; fi12++ {
        v3 = append(v3, v6)
        _ = fi12
    }
    var v13 int64 = (^(v7))
    v14 := GenS1{F0: v7, F1: false, F2: 9223372036854775807}
    switch 0 {
    case 0:
        if (1 == 4294967295) {
            fmt.Println(v8, v1)
            for fi15 := 0; fi15 < 11; fi15++ {
                switch int(v8) {
                case 914:
                    if true {
                        if true {
                            switch "foo bar" {
                            case "":
                                v16 := v5[2]
                                _ = v16
                            case "\n\t":
                                v17 := v5[v2]
                                _ = v17
                            case "foo bar":
                                v18 := GenS2{F0: (v11 + v1), F1: v1}
                                _ = v18
                            }
                            v8 = (v8 >> ((65535) & 63))
                        }
                        if false {
                            var v19 int8 = 100
                            _ = v19
                        }
                    }
                case 279:
                    var v20 float32 = 1.5
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 2; fi21++ {
                        var v22 int16 = 42
                        var v23 float32 = (v10 + 1e10)
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
                _ = fi15
            }
            var v24 uint16 = 42
            _ = v24
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v8 = (v8 >> ((821) & 63))
    if (("" + v6) != (v4 + ((("hello" + "hello") + (v4 + ("" + v6))) + (((v6 + v6) + "a") + "hello")))) {
        v2 = (v2 - v2)
        v1 = v1
        v25 := v5[v2]
        _ = v25
    } else {
        fmt.Println(v13)
        v26 := GenS3{F0: 0}
        v11 -= v1
        _ = v26
    }
    var v27 uint16 = 761
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
