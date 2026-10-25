package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int8 = -128
    var v2 uint32 = 2
    var v3 int16 = -1
    var v4 int64 = 735
    v5 := []int{4294967295}
    var v6 string = ("" + (("\n\t" + "a") + "a"))
    v7 := GenS1{F0: 11}
    var v8 int32 = -2147483648
    var v9 int8 = 42
    v8 = -2147483648
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 float32 = float32((float64(((0.1 + 0.0) - 1e10)) - -2.25))
        v11 = v11
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi10
        _ = v11
    }
    v6 = v6
    var v12 uint8 = 2
    var v13 float32 = 1e10
    if (!false) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (!true) {
            v14 := GenS1{F0: v2}
            _ = v14
        } else {
            v8 -= (^(v8))
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            if true {
                for fi16 := 0; fi16 < 18; fi16++ {
                    fmt.Println(fi15, v9, v1)
                    v9 += (v9 >> ((18446744073709551615) & 63))
                    v5 = append(v5, int(v12))
                    _ = fi16
                }
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int64 = (v4 & v4)
                    var v19 int64 = v18
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 int32 = (^(v8))
                _ = v20
            }
            _ = fi15
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 bool = false
        _ = fi21
        _ = v22
    }
    var v23 float64 = (-2.25 - 1e10)
    switch 319 {
    case 681:
        if true {
            var v24 int = 2147483647
            var v25 uint16 = 391
            _ = v24
            _ = v25
        }
    case 128:
        v26 := GenS1{F0: (^(v2))}
        _ = v26
    }
    v2 = (v2 % (((v2 | (v2 - (v2 ^ 161)))) | 1))
    v27 := GenS2{F0: (v12 & (v12 % (((v12 / (((v12 & v12)) | 1))) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
