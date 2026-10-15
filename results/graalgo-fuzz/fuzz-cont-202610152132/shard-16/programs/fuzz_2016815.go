package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 int32) uint {
    p0 = (p0 - (p0 + p0))
    var v4 uint32 = 2
    _ = p0
    _ = v4
    return 365
}
func genFn1() int32 {
    var v5 int = 127
    v5 += (v5 + (v5 ^ v5))
    var v6 float64 = 0.0
    switch "" {
    case "":
        v5 = -116
    default:
        v5 -= v5
    }
    v7 := GenS1{F0: 0, F1: 1, F2: 0}
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 bool = false
    var v2 bool = (!(2147483647 >= 473))
    var v3 int16 = -266
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -1, F1: -128, F2: 255}
    var v11 string = ""
    fmt.Println(v11)
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := []int{(^(fi12)), (^(fi12)), (fi12 * 801)}
        v2 = (v2 || true)
        v14 := GenS1{F0: -121, F1: 127, F2: 2}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint = 255
    var v16 int64 = 9223372036854775807
    var v17 int16 = (v3 << ((337) & 63))
    var v18 uint16 = 1
    fmt.Println(v15)
    var v19 uint32 = 0
    for fi20 := 0; fi20 < 1; fi20++ {
        switch ("foo bar" + "a") {
        case "":
            v1 = ((v11 + ("hello" + v11)) == v11)
        case "\n\t":
            if (!false) {
                if v2 {
                    fmt.Println(v17, v18, v3)
                    switch v11 {
                    case "":
                        var v21 int32 = 772
                        _ = v21
                    default:
                        v22 := []int{fi20, (^(fi20)), -9223372036854775808}
                        _ = v22
                    }
                } else {
                    fmt.Println(fi20)
                    v23 := GenS1{F0: 1, F1: -1, F2: v18}
                    _ = v23
                }
            } else {
                fmt.Println(v15)
                for fi24 := 0; fi24 < 3; fi24++ {
                    v25 := GenS1{F0: 101, F1: -128, F2: (v18 + (v18 % (((v18 + 42)) | 1)))}
                    v2 = (v16 <= (v16 >> (((v15 * (v15 + v15))) & 63)))
                    _ = fi24
                    _ = v25
                }
            }
        default:
            var v26 int = (fi20 + 103)
            _ = v26
        }
        v19 = v19
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 bool = (v19 >= 1)
            _ = fi27
            _ = v28
        }
        _ = fi20
    }
    var v29 uint16 = (v18 % ((v18) | 1))
    v30 := []int{32768, 888, 32768}
    var v31 int16 = int16(v29)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
