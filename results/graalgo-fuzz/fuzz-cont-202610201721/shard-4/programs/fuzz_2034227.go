package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int; F2 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int8) int {
    var v4 uint8 = 128
    switch "hello" {
    case "\n\t":
        v4 -= 127
    case "a":
        p0 = 0
    default:
        var v5 uint32 = 2147483647
        _ = v5
    }
    _ = p0
    _ = v4
    return 128
}

func main() {
    var v1 uint = 673
    var v2 int = -9223372036854775808
    v3 := []uint32{1, 839, 865, 65535}
    for fi6 := 0; fi6 < 2; fi6++ {
        if (!false) {
            v7 := GenS1{F0: (v1 >> ((v1) & 63)), F1: "a"}
            _ = v7
        }
        var v8 uint = v1
        _ = fi6
        _ = v8
    }
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 2; fi10++ {
        if false {
            var v11 int64 = -316
            v2 = genFn0(-128)
            _ = v11
        } else {
            v12 := GenS2{F0: 42}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v12
        }
        v13 := GenS2{F0: 100}
        _ = fi10
        _ = v13
    }
    v9 -= ((v9 * (v9 + (v9 / v9))) * (v9 + v9))
    v14 := GenS3{F0: 0, F1: v2, F2: 253}
    fmt.Println(v2)
    switch (v2 + genFn0(-128)) {
    case 1:
        switch v2 {
        case -1:
            if false {
                if false {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        if true {
                            var v16 uint32 = uint32(v1)
                            v3 = append(v3, uint32(v1))
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v16
                        } else {
                            fmt.Println(v2)
                        }
                        _ = fi15
                    }
                } else {
                    v17 := GenS3{F0: -2147483648, F1: v2, F2: -1}
                    _ = v17
                }
            }
        case -390:
            v9 += v9
        }
    case 127:
        var v18 int8 = 42
        _ = v18
    default:
        if ((v2 >> (((^(v1))) & 63)) != v2) {
            var v19 uint16 = 32767
            if true {
                v20 := GenS1{F0: v1, F1: "hello"}
                fmt.Println(v19, v2)
                fmt.Println(v9)
                _ = v20
            }
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            _ = v19
        }
    }
    v21 := GenS4{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
