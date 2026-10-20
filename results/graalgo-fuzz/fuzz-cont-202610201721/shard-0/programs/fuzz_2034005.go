package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int8 }
func genFn0(p0 int32, p1 uint32, p2 int8) int8 {
    v7 := GenS1{F0: 2}
    if true {
        if true {
            for fi8 := 0; fi8 < 18; fi8++ {
                switch fi8 {
                case -746:
                    var v9 int16 = -11
                    _ = v9
                }
                v10 := []int{fi8}
                var v11 float32 = 1e10
                _ = fi8
                _ = v10
                _ = v11
            }
        } else {
            if true {
                var v12 uint64 = 293
                switch "hello" {
                case "a":
                    v13 := GenS2{F0: p0, F1: 42, F2: 208}
                    _ = v13
                }
                p0 -= (p0 << ((2) & 63))
                _ = v12
            } else {
                p0 = -565
                var v14 int32 = (p0 + p0)
                if (0 <= 63) {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        fmt.Println(p0, fi15)
                        _ = fi15
                    }
                    if false {
                        p1 = (p1 - (p1 << ((2) & 63)))
                        switch -1 {
                        case 127:
                            var v16 uint64 = 614
                            _ = v16
                        case -642:
                            var v17 uint8 = 237
                            _ = v17
                        case 4294967295:
                            v18 := []int{}
                            _ = v18
                        default:
                            p2 = (p2 | p2)
                        }
                    }
                    var v19 uint16 = 2
                    _ = v19
                }
                _ = v14
            }
        }
    }
    switch -638 {
    case 128:
        switch "a" {
        case "hello":
            v20 := GenS3{F0: 32767, F1: (!false)}
            _ = v20
        }
    case 2147483647:
        for fi21 := 0; fi21 < 6; fi21++ {
            p0 -= -102
            _ = fi21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 ^ (^(p2)))
}
func genFn1(p0 bool, p1 int8, p2 uint8) uint8 {
    p1 = -116
    p1 += genFn0(573, 4294967295, (^(p1)))
    var v22 float32 = ((1.5 / -2.25) / 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return p2
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[int]bool)
    var v3 uint64 = 1
    var v4 uint8 = 39
    var v5 int = 127
    var v6 string = "\n\t"
    v23 := v2[v1]
    v24 := []int{32767, (v1 - v1), (v5 - -9223372036854775808)}
    if len(v24) > 0 {
        v24[0] = -1
    }
    delete(v2, (^(v1)))
    var v25 uint = 26
    v25 += 255
    var v26 int16 = 1000
    if (0 >= 1) {
        var v27 uint = (v25 + 65535)
        for fi28 := 0; fi28 < 13; fi28++ {
            switch -9223372036854775808 {
            case 1:
                v29 := GenS4{F0: 42}
                _ = v29
            case 2147483647:
                v2[v5] = ("hello" != v6)
            }
            _ = fi28
        }
        for fi30 := 0; fi30 < 15; fi30++ {
            if (0 >= 2) {
                var v31 int64 = 42
                _ = v31
            }
            _ = fi30
        }
        _ = v27
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        switch -931 {
        case -723:
            var v33 string = "hello"
            _ = v33
        case 128:
            v34 := GenS5{F0: (v26 % ((v26) | 1)), F1: 22}
            _ = v34
        }
        v35 := GenS1{F0: 42}
        var v36 bool = false
        _ = fi32
        _ = v35
        _ = v36
    }
    _ = len(v2)
    var v37 int8 = -1
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
