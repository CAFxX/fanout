package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 int }
func genFn0() int32 {
    var v4 int32 = -1
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := []int{(fi6 >> ((1) & 63)), 1}
            if false {
                if true {
                    v4 += 724
                    v4 = 42
                    var v8 uint = 65535
                    _ = v8
                } else {
                    var v9 uint8 = 126
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = v9
                }
                switch "\n\t" {
                case "a":
                    var v10 uint32 = 349
                    var v11 int32 = 2147483647
                    var v12 int64 = -356
                    _ = v10
                    _ = v11
                    _ = v12
                case "\n\t":
                    var v13 int32 = (^(v4))
                    _ = v13
                default:
                    if (!(false || false)) {
                        fmt.Println(fi5)
                    }
                }
            } else {
                v4 -= v4
                var v14 uint = 65535
                var v15 uint64 = 1
                fmt.Println(fi6)
                _ = len(v7)
                _ = v14
                _ = v15
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    _ = v4
    return v4
}

func main() {
    var v1 uint32 = 947
    var v2 int32 = 0
    var v3 int = 1000
    fmt.Println(v3, v2)
    fmt.Println(v2)
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS1{F0: false, F1: 314}
        _ = fi16
        _ = v17
    }
    switch 1 {
    case 42:
        if false {
            if true {
                v18 := GenS2{F0: (^(v2)), F1: (-(v3))}
                var v19 uint64 = 2
                _ = v18
                _ = v19
            }
            var v20 int64 = 766
            _ = v20
        }
    default:
        v21 := GenS3{F0: -1, F1: v3}
        _ = v21
    }
    fmt.Println(v1, v1, v1)
    var v22 uint8 = 128
    if true {
        v23 := []int{}
        if len(v23) > 0 {
            v23[0] = int(v1)
        }
        _ = v23
    }
    var v24 int32 = genFn0()
    v2 -= (v24 + v24)
    var v25 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
}
