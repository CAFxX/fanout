package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint = 1
    var v2 int64 = 0
    var v3 float64 = 0.1
    var v4 string = "\n\t"
    var v5 float64 = 0.1
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 float32 = 0.0
        for fi8 := 0; fi8 < 3; fi8++ {
            v1 = (v1 >> (((^(v1))) & 63))
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{}
            var v11 int = (fi9 + (-(fi9)))
            _ = fi9
            _ = v10
            _ = v11
        }
        switch 482 {
        case -38:
            var v12 uint16 = 42
            _ = v12
        case 32768:
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS1{F0: 255, F1: (("" + v4) + v4)}
                v15 := GenS2{F0: (!false)}
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            v3 = float64(0.1)
        }
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint32 = 2147483648
        v17 += v17
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS1{F0: 35, F1: v4}
            v4 = ""
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        if (true || false) {
            v1 -= 242
            for fi21 := 0; fi21 < 17; fi21++ {
                fmt.Println(v5, v5)
                v22 := []int{}
                v3 = v3
                var v23 float32 = (3.14159 - 1.5)
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 19; fi24++ {
                if ((v2 % (((v2 ^ v2)) | 1)) <= (v2 << ((v1) & 63))) {
                    fmt.Println(v5)
                    v3 = v5
                }
                _ = fi24
            }
        }
        _ = fi20
    }
    var v25 uint = (v1 % ((v1) | 1))
    v26 := []int{4294967295}
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 uint8 = 255
        _ = fi27
        _ = v28
    }
    v4 = "foo bar"
    if ((0.0 == float32(v3)) && (true || (0 != 42))) {
        for fi29 := 0; fi29 < 20; fi29++ {
            fmt.Println(v5, v1, v4)
            if false {
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
            } else {
                if true {
                    var v30 float64 = 3.14159
                    for fi31 := 0; fi31 < 16; fi31++ {
                        fmt.Println(v25, fi29)
                        fmt.Println(v4)
                        _ = fi31
                    }
                    _ = v30
                }
                v1 = (^(v1))
                var v32 int64 = (v2 << (((v25 + (v1 << ((65535) & 63)))) & 63))
                _ = v32
            }
            _ = fi29
        }
    }
    var v33 int64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
