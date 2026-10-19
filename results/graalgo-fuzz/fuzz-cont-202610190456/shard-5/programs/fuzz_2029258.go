package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 string }

func main() {
    v1 := []uint32{1}
    var v2 int8 = -128
    var v3 uint32 = 663
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 2, F1: false}
    if false {
        var v6 bool = (!false)
        var v7 bool = false
        v6 = v7
        _ = v6
        _ = v7
    }
    if (!(546 > 1)) {
        v2 -= (v2 - (-(v2)))
        if false {
            fmt.Println(v2, v3)
            for fi8 := 0; fi8 < 2; fi8++ {
                v3 -= 2147483648
                _ = fi8
            }
            _ = len(v1)
        }
    }
    if (!false) {
        fmt.Println(v3, v2)
        if true {
            switch 9223372036854775807 {
            case -1:
                switch ("a" + ("a" + "foo bar")) {
                case "a":
                    if (-32768 <= 42) {
                        fmt.Println(v2, v2)
                    }
                case "foo bar":
                    for fi9 := 0; fi9 < 19; fi9++ {
                        v10 := v4[4294967295]
                        v2 -= v2
                        _ = fi9
                        _ = v10
                    }
                case "\n\t":
                    var v11 bool = false
                    _ = v11
                default:
                    if true {
                        var v12 uint8 = 128
                        _ = v12
                    } else {
                        var v13 uint8 = 128
                        _ = v13
                    }
                }
            default:
                var v14 uint32 = 844
                _ = v14
            }
            if true {
                if (-32768 != -413) {
                    var v15 float32 = 1.5
                    v2 += 100
                    var v16 uint64 = 4294967295
                    _ = v15
                    _ = v16
                }
                for fi17 := 0; fi17 < 4; fi17++ {
                    fmt.Println(v3)
                    switch (fi17 - (^(fi17))) {
                    case -430:
                        fmt.Println(v2, fi17, v2)
                    case -9223372036854775808:
                        fmt.Println(fi17, v3)
                    case 128:
                        delete(v4, fi17)
                    default:
                        delete(v4, (fi17 ^ fi17))
                    }
                    _ = fi17
                }
            }
            var v18 uint64 = 2
            _ = v18
        } else {
            var v19 int16 = 55
            fmt.Println(v3, v2, v3)
            v3 = v3
            _ = v19
        }
    }
    v3 -= (v3 & (v3 - v3))
    var v20 int8 = v2
    v20 -= v2
    if true {
        v2 = (v20 ^ 100)
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            if false {
                v20 = v2
                delete(v4, fi21)
                if (9223372036854775807 >= 42) {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        var v23 int16 = 777
                        _ = fi22
                        _ = v23
                    }
                } else {
                    v24 := GenS1{F0: 1, F1: (!(-21 > 2147483647))}
                    if true {
                        v25 := GenS2{F0: 2147483647, F1: 42, F2: 0}
                        v26 := GenS1{F0: 2, F1: ((56 > v2) || ((v3 + 78) < v3))}
                        _ = v25
                        _ = v26
                    } else {
                        v2 -= (v2 & (v2 + -65))
                    }
                    var v27 float32 = 1.5
                    _ = v24
                    _ = v27
                }
            } else {
                fmt.Println(v20)
                v20 = (^(v2))
                _ = len(v4)
            }
            var v28 int8 = v2
            _ = fi21
            _ = v28
        }
        v29 := GenS3{F0: 65535, F1: "hello", F2: "a"}
        v2 -= 42
        _ = v29
    }
    v3 = (v3 & v3)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
