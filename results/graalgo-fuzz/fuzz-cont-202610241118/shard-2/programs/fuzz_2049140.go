package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int32; F2 int8 }
func genFn0() int32 {
    v4 := []int{}
    v5 := GenS1{F0: 0, F1: ""}
    v6 := GenS1{F0: 127, F1: "a"}
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int16 = 362
    v2 := []uint32{}
    var v3 float32 = 0.0
    if true {
        v7 := GenS1{F0: 1, F1: "\n\t"}
        v3 += ((0.1 + float32(0.0)) - v3)
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS2{F0: 28}
            v3 = (3.14159 * float32(3.14159))
            if true {
                for fi10 := 0; fi10 < 16; fi10++ {
                    fmt.Println(fi8)
                    fmt.Println(v3)
                    var v11 uint8 = 2
                    v1 = (v1 >> ((65535) & 63))
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v12 := GenS2{F0: 127}
    var v13 string = ""
    var v14 uint32 = 979
    switch (v13 + ((("a" + v13) + v13) + (v13 + "\n\t"))) {
    case "foo bar":
        v15 := GenS2{F0: 5}
        _ = v15
    default:
        fmt.Println(v3)
    }
    v13 = v13
    if false {
        if (!(1.5 >= v3)) {
            if true {
                v2 = append(v2, (^(v14)))
                v1 = v1
            } else {
                v13 = ("\n\t" + v13)
                if len(v2) > 0 {
                    v2[0] = v14
                }
                v2 = append(v2, (v14 | v14))
            }
        } else {
            v16 := GenS3{F0: -1}
            switch 4294967295 {
            case 128:
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 uint64 = 0
                    v2 = append(v2, v14)
                    _ = fi17
                    _ = v18
                }
            case -9223372036854775808:
                if true {
                    v2 = append(v2, 1)
                    if (!true) {
                        var v19 uint8 = 2
                        fmt.Println(v14)
                        if len(v2) > 0 {
                            v2[0] = (v14 / (((v14 >> ((2) & 63))) | 1))
                        }
                        _ = v19
                    }
                }
            case 2147483647:
                fmt.Println(v14)
            }
            _ = v16
        }
        if true {
            if ((!false) || false) {
                var v20 int = -398
                if true {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if false {
                            v13 = v13
                        } else {
                            _ = len(v2)
                            v22 := GenS4{F0: 1, F1: 42, F2: 0}
                            _ = v22
                        }
                        if false {
                            var v23 bool = (352 >= 18446744073709551615)
                            _ = v23
                        }
                        _ = fi21
                    }
                    fmt.Println(v3, v1, v14)
                }
                v14 += (v14 + (v14 * v14))
                _ = v20
            }
            v14 = (v14 / ((v14) | 1))
            var v24 string = v13
            _ = v24
        }
        var v25 int32 = -1
        _ = v25
    }
    var v26 int8 = -1
    v26 += v26
    v27 := GenS4{F0: 0, F1: -27, F2: (v26 >> ((65535) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
