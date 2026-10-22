package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 int = 32768
    var v3 int = -9223372036854775808
    var v4 uint16 = 561
    v4 -= v4
    if (4294967295 >= 967) {
        var v5 int = v2
        for fi6 := 0; fi6 < 16; fi6++ {
            switch "foo bar" {
            case "":
                fmt.Println(fi6, v2)
            default:
                switch (fi6 ^ (^(fi6))) {
                case 42:
                    v3 = (fi6 ^ v2)
                }
            }
            if (-29 <= -128) {
                var v7 int16 = 42
                v7 = (v7 >> ((715) & 63))
                _ = v7
            }
            v3 = 127
            _ = fi6
        }
        _ = v5
    } else {
        fmt.Println(v2, v3)
        if (1e10 < float32((0.0 * 1.5))) {
            v4 -= v4
            if false {
                switch 127 {
                case 128:
                    v4 = v4
                case 42:
                    v3 = (v3 << ((172) & 63))
                case 32768:
                    var v8 int8 = -88
                    _ = v8
                }
                fmt.Println(v4, v3)
            }
            v9 := []int{(v2 % (((v3 - 32768)) | 1)), v3, 4294967295}
            _ = v9
        } else {
            v10 := GenS1{F0: 42}
            for fi11 := 0; fi11 < 6; fi11++ {
                if false {
                    v12 := GenS1{F0: 255}
                    v2 = v2
                    _ = v12
                } else {
                    var v13 uint32 = 77
                    v14 := []int{(v2 ^ (fi11 ^ -18))}
                    _ = v13
                    _ = v14
                }
                if true {
                    var v15 int = v2
                    if false {
                        v16 := []int{(fi11 * v15), (^(v15))}
                        _ = v16
                    } else {
                        fmt.Println(v3)
                    }
                    _ = v15
                } else {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        _ = len(v1)
                        fmt.Println(v4)
                        _ = fi17
                    }
                    for fi18 := 0; fi18 < 14; fi18++ {
                        var v19 bool = (!false)
                        _ = fi18
                        _ = v19
                    }
                    v4 -= (v4 + 42)
                }
                _ = fi11
            }
            _ = v10
        }
    }
    v20 := GenS2{F0: "foo bar", F1: v4}
    var v21 int8 = -128
    var v22 uint16 = v4
    fmt.Println(v2, v3, v22)
    v23 := []int{v3, v3}
    var v24 int8 = v21
    v25 := v1[(v3 - int(v21))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
