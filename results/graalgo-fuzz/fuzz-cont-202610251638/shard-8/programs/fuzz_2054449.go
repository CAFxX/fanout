package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 15
    var v3 float64 = 0.0
    v4 := make(map[uint32]bool)
    var v5 string = ("hello" + "\n\t")
    var v6 uint32 = 2147483647
    if (255 == 18446744073709551615) {
        var v7 uint = 1
        if false {
            switch 127 {
            case 0:
                v5 = (v5 + (v5 + v5))
            case -914:
                if true {
                    v8 := []int{105}
                    v9 := GenS1{F0: int32(v2), F1: (859 == 9223372036854775807)}
                    v2 = 255
                    _ = v8
                    _ = v9
                }
            }
            v10 := []int{32767, 42}
            _ = v10
        } else {
            if (!false) {
                var v11 string = ("\n\t" + v5)
                _ = v11
            }
        }
        _ = v7
    }
    fmt.Println(v5, v2)
    if false {
        if (2 < 0) {
            fmt.Println(v5)
        }
        v12 := GenS2{F0: 32767, F1: (v5 + "hello")}
        _ = v12
    } else {
        v2 = uint8(v6)
    }
    v6 += uint32(v2)
    v5 = ((v5 + v5) + v5)
    if ((true && true) && false) {
        if (!false) {
            delete(v4, v6)
            v13 := GenS1{F0: 42, F1: (!true)}
            _ = v13
        }
    }
    v3 -= float64((-2.25 - 0.0))
    v2 = (v2 << ((538) & 63))
    v14 := []int{}
    var v15 int64 = -964
    fmt.Println(v3, v15)
    switch 401 {
    case -9223372036854775808:
        if true {
            var v16 uint32 = (^(v6))
            for fi17 := 0; fi17 < 4; fi17++ {
                if false {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        var v19 int16 = 172
                        delete(v1, (fi18 & (-(fi18))))
                        _ = fi18
                        _ = v19
                    }
                    var v20 string = (((v5 + v5) + v5) + "\n\t")
                    v3 -= v3
                    _ = v20
                } else {
                    var v21 int = fi17
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v23 := GenS3{F0: uint8(v6), F1: true}
                        var v24 bool = false
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = v21
                }
                _ = fi17
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 uint64 = 328
                v27 := GenS1{F0: 628, F1: (985 < (v16 - (v6 / (((^(v16))) | 1))))}
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = v16
        }
    case 1:
        fmt.Println(v2, v5, v15)
    case 32767:
        v5 = (v5 + (v5 + "foo bar"))
    default:
        for fi28 := 0; fi28 < 7; fi28++ {
            for fi29 := 0; fi29 < 9; fi29++ {
                fmt.Println(v2, v15, v3)
                _ = fi29
            }
            v5 = v5
            _ = fi28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
