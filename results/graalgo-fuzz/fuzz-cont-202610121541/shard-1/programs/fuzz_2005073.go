package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint }
func genFn0() int32 {
    v4 := []int{366, 4294967295, 32767, 128}
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            v4 = append(v4, (^(fi5)))
            _ = fi6
        }
        _ = fi5
    }
    _ = v4
    return 42
}

func main() {
    var v1 int32 = 0
    var v2 int32 = v1
    v3 := []uint32{}
    v3 = append(v3, 14)
    fmt.Println(v1, v2)
    fmt.Println(v2, v2)
    switch -473 {
    case 1:
        var v7 bool = (!(127 >= 127))
        _ = v7
    case 9223372036854775807:
        var v8 int16 = 32767
        _ = v8
    case -9223372036854775808:
        v1 = genFn0()
    }
    var v9 uint = 2
    if ((1.5 / 1e10) < ((0.0 + -2.25) - (3.14159 * -2.25))) {
        if true {
            v1 += (v2 % ((genFn0()) | 1))
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            switch "a" {
            case "foo bar":
                for fi11 := 0; fi11 < 4; fi11++ {
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    _ = fi11
                }
            }
            v3 = append(v3, 4294967295)
            _ = fi10
        }
    } else {
        var v12 uint16 = 42
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v13 := GenS1{F0: 2, F1: 2, F2: (true || true)}
    var v14 bool = false
    v2 = v1
    switch "\n\t" {
    case "a":
        v15 := GenS2{F0: (v9 >> (((v9 - v9)) & 63))}
        _ = v15
    case "\n\t":
        var v16 uint64 = 806
        _ = v16
    default:
        if v14 {
            if len(v3) > 0 {
                v3[0] = 1
            }
            var v17 uint8 = 127
            v14 = v14
            _ = v17
        }
    }
    var v18 bool = v14
    var v19 int16 = -1
    if (128 >= -9223372036854775808) {
        if true {
            v2 = genFn0()
            for fi20 := 0; fi20 < 20; fi20++ {
                for fi21 := 0; fi21 < 15; fi21++ {
                    fmt.Println(v18)
                    _ = fi21
                }
                fmt.Println(v2, v2, v19)
                _ = fi20
            }
        } else {
            if true {
                fmt.Println(v14)
                v18 = (33 < -1)
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 16; fi23++ {
                    var v24 uint32 = uint32(v9)
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    v18 = true
                    var v25 uint = 65535
                    var v26 uint64 = 849
                    fmt.Println(fi23, fi23)
                    var v27 uint = v9
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = fi22
            }
        }
        v28 := GenS2{F0: (v9 / ((1) | 1))}
        _ = v28
    }
    v14 = true
    v19 = (v19 + (v19 - (v19 % (((v19 ^ 1)) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
