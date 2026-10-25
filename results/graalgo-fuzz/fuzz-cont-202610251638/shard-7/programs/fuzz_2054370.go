package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64 }
func genFn0() string {
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 20; fi6++ {
        v5 = v5
        if true {
            fmt.Println(fi6)
            fmt.Println(v5, v5)
        }
        v7 := GenS1{F0: (fi6 | 836), F1: -2147483648, F2: 0}
        _ = fi6
        _ = v7
    }
    _ = v5
    return ("foo bar" + "foo bar")
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 128
    var v3 int16 = 1
    var v4 uint16 = 27
    var v8 int16 = 42
    v9 := GenS2{F0: (v2 - 255)}
    switch 32767 {
    case 4294967295:
        v8 = v8
    case -576:
        fmt.Println(v8, v2, v3)
    case 905:
        v10 := GenS2{F0: (v2 & (^(v2)))}
        _ = v10
    default:
        v11 := []int{-878}
        _ = v11
    }
    fmt.Println(v3, v2)
    v12 := []int{-9223372036854775808}
    v8 = (v3 * (v3 / (((v3 | v8)) | 1)))
    for fi13 := 0; fi13 < 8; fi13++ {
        switch ("hello" + ("hello" + "a")) {
        case "a":
            if false {
                var v14 uint = 633
                v3 = v8
                _ = v14
            } else {
                if false {
                    switch fi13 {
                    case 0:
                        v3 += (v8 + int16(v2))
                    case 128:
                        var v15 int64 = 0
                        _ = v15
                    case 42:
                        if false {
                            for fi16 := 0; fi16 < 8; fi16++ {
                                switch 4294967295 {
                                case -9223372036854775808:
                                    v3 = 0
                                case 32767:
                                    fmt.Println(v3, fi13)
                                }
                                var v17 int = fi16
                                v18 := GenS1{F0: (v17 / ((fi13) | 1)), F1: 42, F2: (v4 - v4)}
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS3{F0: 42}
                            _ = v19
                        } else {
                            v20 := GenS3{F0: 100}
                            v8 = v8
                            _ = v20
                        }
                    default:
                        var v21 int64 = -615
                        _ = v21
                    }
                }
            }
        case "\n\t":
            v12 = append(v12, fi13)
        case "hello":
            switch (^(fi13)) {
            case 32767:
                v22 := GenS4{F0: 42, F1: v2, F2: 28}
                _ = v22
            case 128:
                v4 = (v4 >> ((666) & 63))
            case 32768:
                v23 := GenS5{F0: 18446744073709551615}
                _ = v23
            default:
                fmt.Println(v3, v4)
            }
        }
        _ = fi13
    }
    var v24 bool = false
    var v25 uint16 = uint16(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
