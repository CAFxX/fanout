package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 int8, p1 int64) uint16 {
    if true {
        var v6 uint64 = 42
        p0 -= p0
        if ((2147483647 != 485) && true) {
            var v7 string = ""
            _ = v7
        }
        _ = v6
    } else {
        var v8 int = 585
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        if (715 < 888) {
            for fi10 := 0; fi10 < 3; fi10++ {
                switch "foo bar" {
                case "hello":
                    p1 -= p1
                case "":
                    switch fi9 {
                    case 32768:
                        var v11 int64 = p1
                        _ = v11
                    case -169:
                        fmt.Println(p0)
                    case 42:
                        v12 := GenS1{F0: -1}
                        _ = v12
                    default:
                        p1 += (p1 | p1)
                    }
                case "\n\t":
                    var v13 uint16 = 0
                    _ = v13
                default:
                    v14 := []int{}
                    _ = v14
                }
                _ = fi10
            }
            p1 = 1
            switch ("foo bar" + (("\n\t" + "hello") + "foo bar")) {
            case "foo bar":
                v15 := []int{523}
                _ = v15
            case "a":
                var v16 int32 = 2147483647
                _ = v16
            default:
                v17 := GenS2{F0: "hello"}
                _ = v17
            }
        }
        _ = fi9
    }
    p0 -= 84
    p1 = (p1 | -197)
    _ = p0
    _ = p1
    return 909
}
func genFn1(p0 int64, p1 uint16) uint64 {
    p0 = (p0 & (p0 | p0))
    var v18 int16 = -886
    if false {
        v18 = 32767
        switch "hello" {
        case "foo bar":
            if true {
                switch "foo bar" {
                case "hello":
                    fmt.Println(p1)
                }
            } else {
                var v19 bool = (!(42 == 1))
                p1 = (p1 ^ 42)
                v20 := GenS3{F0: 34, F1: 1}
                _ = v19
                _ = v20
            }
        case "hello":
            var v21 uint64 = 4294967295
            _ = v21
        default:
            v22 := GenS3{F0: 0, F1: 901}
            _ = v22
        }
        switch 32767 {
        case 127:
            var v23 uint32 = 2
            _ = v23
        }
    } else {
        p0 = p0
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 int = (-(fi24))
            v26 := GenS4{F0: "a", F1: (p0 % (((-(p0))) | 1))}
            v27 := []int{fi24}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        v18 = 0
    }
    _ = p0
    _ = p1
    _ = v18
    return 18446744073709551615
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 0
    var v3 int8 = -128
    var v4 string = "\n\t"
    v5 := []int{-685}
    v28 := GenS1{F0: 68}
    for fi29 := 0; fi29 < 9; fi29++ {
        if true {
            fmt.Println(v4)
            _ = len(v1)
        }
        v4 = (v4 + v4)
        var v30 int8 = (v3 << ((0) & 63))
        _ = fi29
        _ = v30
    }
    v4 = "foo bar"
    fmt.Println(v2)
    v2 += v2
    if (v2 >= 1) {
        v5 = append(v5, 55)
        var v31 uint = 0
        if true {
            var v32 float32 = 3.14159
            v2 = (^(v2))
            for fi33 := 0; fi33 < 14; fi33++ {
                v4 = v4
                _ = fi33
            }
            _ = v32
        } else {
            var v34 uint8 = 1
            v35 := GenS5{F0: (v34 ^ (v34 + v34))}
            v4 = (v4 + v4)
            _ = v34
            _ = v35
        }
        _ = v31
    } else {
        fmt.Println(v4, v2)
        var v36 float32 = 0.0
        _ = v36
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi37
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        if true {
            v39 := GenS6{F0: 4294967295, F1: 255, F2: v2}
            _ = v39
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi38
    }
    _ = len(v5)
    var v40 uint8 = 2
    v41 := GenS3{F0: v3, F1: 349}
    if (!false) {
        for fi42 := 0; fi42 < 16; fi42++ {
            v3 += v3
            _ = fi42
        }
        switch 957 {
        case 937:
            var v43 int = 298
            _ = v43
        case 44:
            var v44 string = (v4 + ("hello" + ((v4 + ((v4 + v4) + v4)) + v4)))
            _ = v44
        }
        fmt.Println(v40, v40, v4)
    } else {
        var v45 uint64 = 2
        var v46 uint32 = v2
        _ = v45
        _ = v46
    }
    v2 = v2
    v4 = ((v4 + ((v4 + ("a" + "foo bar")) + (v4 + v4))) + v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
