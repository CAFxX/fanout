package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0() int32 {
    v5 := GenS1{F0: 111, F1: 18, F2: 0}
    var v6 uint = 790
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int = 755
    var v2 bool = ("\n\t" < "a")
    var v3 uint64 = 213
    var v4 float64 = (1e10 / 0.1)
    if v2 {
        v2 = (!true)
        v1 += (v1 % ((v1) | 1))
        v7 := GenS2{F0: 0}
        _ = v7
    } else {
        var v8 uint8 = 72
        _ = v8
    }
    v9 := []int{42, v1, (-(v1))}
    fmt.Println(v2)
    if (v2 && (!(false || v2))) {
        if v2 {
            var v10 uint32 = 0
            fmt.Println(v3, v2)
            v11 := GenS3{F0: v3}
            _ = v10
            _ = v11
        } else {
            var v12 uint8 = 1
            if len(v9) > 0 {
                v9[0] = (v1 ^ 127)
            }
            if ((-2.25 >= float32(v4)) || false) {
                if ("a" == "a") {
                    var v13 int16 = 986
                    fmt.Println(v3, v13)
                    var v14 int8 = -128
                    _ = v13
                    _ = v14
                }
                v9 = append(v9, v1)
            } else {
                if v2 {
                    var v15 uint = 4294967295
                    v2 = ((v12 & v12) > 2)
                    var v16 string = "a"
                    _ = v15
                    _ = v16
                }
            }
            _ = v12
        }
        _ = len(v9)
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS3{F0: (v3 ^ v3)}
        v2 = false
        v2 = (1 >= 2147483647)
        _ = fi17
        _ = v18
    }
    v4 = 0.1
    if false {
        var v19 uint = 880
        _ = v19
    }
    if len(v9) > 0 {
        v9[0] = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
