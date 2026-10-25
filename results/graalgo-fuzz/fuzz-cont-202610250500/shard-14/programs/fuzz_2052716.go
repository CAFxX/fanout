package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = (("hello" + "") + "hello")
    var v3 uint64 = 584
    v4 := []int{42, 0}
    var v5 uint32 = 117
    v6 := GenS1{F0: v5, F1: 42}
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: v5, F1: 615}
    if false {
        if false {
            if false {
                v9 := GenS2{F0: 0, F1: v3, F2: -32768}
                switch -497 {
                case 894:
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 uint32 = v5
                        var v12 bool = (128 > 1)
                        delete(v1, 884)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case -1:
                    var v13 int8 = 127
                    _ = v13
                default:
                    v3 -= (v3 + v3)
                }
                _ = v9
            } else {
                v14 := GenS3{F0: (v3 * (v3 ^ (v3 >> ((684) & 63)))), F1: v2}
                switch v2 {
                case "":
                    fmt.Println(v2, v7, v2)
                case "\n\t":
                    v3 = 59
                }
                _ = v14
            }
        } else {
            v15 := GenS4{F0: v5, F1: (v5 - v5)}
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 uint = 65535
                var v18 uint32 = v5
                _ = fi16
                _ = v17
                _ = v18
            }
            if false {
                var v19 int8 = 123
                var v20 int = 0
                _ = v19
                _ = v20
            } else {
                switch int(v3) {
                case 42:
                    if (0 >= v3) {
                        var v21 float64 = 1.5
                        var v22 int = 9223372036854775807
                        _ = v21
                        _ = v22
                    } else {
                        if ("\n\t" <= v2) {
                            v2 = v2
                        }
                        v23 := GenS2{F0: 114, F1: 1, F2: -190}
                        var v24 uint8 = 0
                        _ = v23
                        _ = v24
                    }
                default:
                    v25 := GenS5{F0: v5, F1: (v5 >> ((241) & 63))}
                    _ = v25
                }
            }
            _ = v15
        }
        v4 = append(v4, 1)
        var v26 bool = true
        _ = v26
    } else {
        fmt.Println(v2)
        for fi27 := 0; fi27 < 15; fi27++ {
            v3 = v3
            if true {
                v7 += v7
                var v28 int8 = 0
                _ = v28
            } else {
                v29 := GenS5{F0: v5, F1: 597}
                v3 += (v3 * (v3 * (v3 & v3)))
                var v30 bool = false
                _ = v29
                _ = v30
            }
            var v31 int8 = 125
            _ = fi27
            _ = v31
        }
    }
    v2 = v2
    v4 = append(v4, 0)
    switch -1 {
    case -2147483648:
        v7 -= v7
    default:
        var v32 uint64 = v3
        _ = v32
    }
    v3 = (v3 % ((42) | 1))
    v33 := GenS6{F0: -586}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
