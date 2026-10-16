package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 int }
type GenS3 struct { F0 bool; F1 bool }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{653, -9223372036854775808, 0}
        var v7 int32 = 172
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        var v8 int64 = 1
        if (0 < 390) {
            v8 = 1
            fmt.Println(v8)
            v9 := []int{-499, -852}
            _ = v9
        } else {
            switch "a" {
            case "foo bar":
                v8 = 1
            case "a":
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := GenS1{F0: -128, F1: (fi10 >> ((188) & 63)), F2: 39}
                    v12 := GenS1{F0: 127, F1: (fi10 << ((698) & 63)), F2: 127}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "":
                if true {
                    if (112 >= 2) {
                        v13 := []int{}
                        v14 := GenS1{F0: -1, F1: 9223372036854775807, F2: 0}
                        for fi15 := 0; fi15 < 4; fi15++ {
                            _ = len(v13)
                            v16 := GenS1{F0: 127, F1: fi15, F2: 50}
                            _ = fi15
                            _ = v16
                        }
                        _ = v13
                        _ = v14
                    } else {
                        v17 := []int{127, 2147483647, 1}
                        fmt.Println(v8)
                        _ = v17
                    }
                    if true {
                        var v18 int = 97
                        _ = v18
                    }
                }
            default:
                var v19 int64 = (v8 % ((v8) | 1))
                _ = v19
            }
            var v20 uint16 = 622
            v21 := GenS2{F0: 1, F1: (^(v20)), F2: 1}
            _ = v20
            _ = v21
        }
        _ = v8
    }
    v22 := []int{32768}
    var v23 int64 = 797
    if (573 < -321) {
        var v24 uint64 = 2
        var v25 int16 = -1
        _ = len(v22)
        _ = v24
        _ = v25
    }
    _ = v22
    _ = v23
    return 517
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 bool = ((("hello" + ("foo bar" + "\n\t")) + "hello") <= (("foo bar" + "foo bar") + "foo bar"))
    v1 = v1
    if (!false) {
        v1 = 42
    }
    v4 = v4
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS3{F0: (v1 <= 255), F1: (!v4)}
        v28 := GenS2{F0: 911, F1: 42, F2: fi26}
        _ = fi26
        _ = v27
        _ = v28
    }
    v4 = (!(42 < 2))
    var v29 int = 9223372036854775807
    switch -1 {
    case 1:
        if (!v4) {
            var v30 uint64 = 1
            _ = v30
        } else {
            var v31 int8 = 1
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 10; fi32++ {
            for fi33 := 0; fi33 < 2; fi33++ {
                fmt.Println(v4)
                fmt.Println(fi33)
                _ = fi33
            }
            _ = fi32
        }
    }
    var v34 float64 = -2.25
    if v4 {
        v35 := GenS3{F0: ((3.14159 + (-2.25 - 3.14159)) == (-2.25 * 1.5)), F1: true}
        var v36 uint32 = 2
        fmt.Println(v34, v29)
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS3{F0: (!false), F1: v4}
        v39 := GenS3{F0: true, F1: v4}
        _ = fi37
        _ = v38
        _ = v39
    }
    v4 = true
    v40 := []int{(v29 % ((776) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v40
}
