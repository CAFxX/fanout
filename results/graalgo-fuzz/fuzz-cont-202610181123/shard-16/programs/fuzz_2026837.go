package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 int8 }
type GenS3 struct { F0 int8; F1 int }
func genFn0() int8 {
    var v5 string = ("\n\t" + "\n\t")
    if (814 < 32767) {
        v6 := GenS1{F0: false, F1: 4294967295}
        _ = v6
    }
    var v7 uint64 = 1
    _ = v5
    _ = v7
    return 1
}

func main() {
    var v1 int64 = 0
    var v2 int8 = 127
    v3 := make(map[int]int)
    var v4 uint16 = 255
    var v8 uint16 = (^(v4))
    if false {
        var v9 int16 = -682
        _ = v9
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint = 612
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: (false && (255 >= 164)), F1: 733}
    v13 := GenS1{F0: true, F1: 519}
    v14 := GenS2{F0: 541, F1: 627, F2: v2}
    if false {
        _ = len(v3)
        fmt.Println(v1)
    } else {
        v4 = v4
        if (true && (true && true)) {
            v15 := GenS1{F0: false, F1: 255}
            v16 := GenS3{F0: v2, F1: -2147483648}
            v1 += -369
            _ = v15
            _ = v16
        } else {
            fmt.Println(v1, v2)
            var v17 int16 = 42
            _ = v17
        }
    }
    var v18 bool = false
    if false {
        for fi19 := 0; fi19 < 20; fi19++ {
            if v18 {
                switch "" {
                case "\n\t":
                    fmt.Println(v8, v1)
                case "hello":
                    if v18 {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            var v21 float64 = float64(((-2.25 + (1.5 - 0.1)) * 3.14159))
                            var v22 int16 = 32767
                            v8 = (^(v8))
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    } else {
                        v18 = v18
                        v4 = (v8 >> ((65535) & 63))
                        if true {
                            v4 = v4
                        }
                    }
                }
                v18 = v18
            }
            _ = fi19
        }
        var v23 uint = 1
        var v24 int = 4294967295
        _ = v23
        _ = v24
    }
    v4 = (v4 & (v8 << ((2) & 63)))
    v25 := GenS3{F0: v2, F1: 813}
    for fi26 := 0; fi26 < 7; fi26++ {
        fmt.Println(v8, v1, v4)
        v18 = false
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v8
}
