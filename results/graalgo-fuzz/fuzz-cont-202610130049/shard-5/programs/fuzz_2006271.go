package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
func genFn0(p0 int16, p1 uint32) int64 {
    if true {
        p0 = (p0 | (p0 * (p0 ^ -32768)))
    } else {
        v5 := []int{0, 776}
        if (1 != 2) {
            var v6 int8 = -37
            _ = v6
        }
        _ = v5
    }
    fmt.Println(p1)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 string = ("" + ("a" + "foo bar"))
    v2 := make(map[uint32]int)
    var v3 uint8 = 127
    var v4 int64 = 42
    var v7 int64 = (v4 % ((v4) | 1))
    var v8 uint64 = 42
    v9 := GenS1{F0: 0}
    v8 = v8
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int8 = 100
        var v12 bool = ((!true) || false)
        v13 := []int{352}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{fi14, 9223372036854775807, fi14, -666}
        for fi16 := 0; fi16 < 10; fi16++ {
            if false {
                fmt.Println(v1)
                v17 := GenS2{F0: 60}
                if (false && false) {
                    v18 := GenS3{F0: 1, F1: 42, F2: v7}
                    v8 = v8
                    fmt.Println(v8, v7)
                    _ = v18
                } else {
                    var v19 uint64 = v8
                    _ = len(v2)
                    _ = v19
                }
                if true {
                    if false {
                        v7 += v7
                        _ = len(v15)
                        v1 = v1
                    }
                    var v20 int16 = 480
                    v8 -= v8
                    _ = v20
                } else {
                    v4 += v7
                    if len(v15) > 0 {
                        v15[0] = fi14
                    }
                }
                v21 := GenS1{F0: -293}
                _ = v17
                _ = v21
            } else {
                v22 := GenS2{F0: 0}
                switch (v1 + v1) {
                case "foo bar":
                    v7 += (^(v7))
                }
                _ = v22
            }
            v3 = v3
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v23 int8 = 100
    var v24 int = 565
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
