package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint; F2 uint32 }
type GenS3 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0() bool {
    var v5 uint16 = 1
    for fi6 := 0; fi6 < 20; fi6++ {
        if (true && (128 > 0)) {
            switch (("\n\t" + (("" + "") + ("" + ("a" + "\n\t")))) + "") {
            case "hello":
                v7 := []int{fi6}
                _ = v7
            case "foo bar":
                if false {
                    var v8 uint = 42
                    var v9 int8 = -1
                    _ = v8
                    _ = v9
                }
            case "\n\t":
                fmt.Println(fi6)
            default:
                if (0 == -1) {
                    v5 = (v5 ^ (v5 - (v5 << ((2) & 63))))
                } else {
                    v10 := []int{(fi6 | -1), fi6, fi6, (fi6 / (((-(fi6))) | 1))}
                    _ = v10
                }
            }
            v11 := GenS1{F0: 42, F1: 2147483648}
            _ = v11
        } else {
            v12 := GenS1{F0: 988, F1: 2147483648}
            v13 := GenS1{F0: -224, F1: 222}
            _ = v12
            _ = v13
        }
        _ = fi6
    }
    v5 = (v5 >> ((846) & 63))
    var v14 int32 = 0
    _ = v14
    _ = v5
    return false
}

func main() {
    v1 := []string{"", "\n\t", ("a" + ("" + (("hello" + "foo bar") + ("a" + "\n\t")))), "a"}
    v2 := make(map[string]int)
    var v3 string = "hello"
    var v4 uint = 259
    if ((1.5 * (-2.25 / 1.0)) == 0.0) {
        for fi15 := 0; fi15 < 7; fi15++ {
            if false {
                if true {
                    var v16 int16 = -1
                    _ = v16
                }
            }
            if false {
                fmt.Println(fi15, v4)
                delete(v2, v3)
                if true {
                    var v17 string = v3
                    _ = v17
                }
            } else {
                v4 -= (v4 << ((65535) & 63))
            }
            v4 -= (v4 >> ((v4) & 63))
            _ = fi15
        }
        fmt.Println(v4, v4)
    }
    var v18 uint8 = 1
    v3 = v3
    var v19 int64 = 9223372036854775807
    var v20 float32 = 1.5
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int64 = 42
        if len(v1) > 0 {
            v1[0] = v3
        }
        _ = fi21
        _ = v22
    }
    var v23 uint8 = v18
    fmt.Println(v20, v20)
    v24 := GenS2{F0: v3, F1: (v4 | (v4 << (((v4 | (v4 ^ (v4 >> (((v4 % ((v4) | 1))) & 63))))) & 63))), F2: 13}
    var v25 int = 32768
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int16 = 1000
        v28 := GenS3{F0: 647, F1: 526, F2: v18}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint8 = 42
    v29 = 42
    var v30 int = v25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
