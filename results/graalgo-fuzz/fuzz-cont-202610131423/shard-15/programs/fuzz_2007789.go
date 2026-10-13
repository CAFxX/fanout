package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16) bool {
    fmt.Println(p0)
    p0 += p0
    if true {
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := []int{(-(fi5)), (fi5 / ((fi5) | 1)), fi5, 32767}
            v7 := GenS1{F0: 9223372036854775807, F1: 666, F2: (p0 ^ p0)}
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 int8 = -1
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return true
}

func main() {
    var v1 int16 = 1
    var v2 float64 = 1e10
    var v3 uint = 2
    var v4 int64 = int64(v3)
    v9 := []int{-9223372036854775808}
    switch ("foo bar" + "hello") {
    case "a":
        switch 2147483647 {
        case 32768:
            var v10 uint32 = 2
            _ = v10
        default:
            fmt.Println(v3)
        }
    default:
        var v11 bool = true
        _ = v11
    }
    switch "foo bar" {
    case "a":
        var v12 int16 = v1
        _ = v12
    case "hello":
        var v13 string = "\n\t"
        _ = v13
    case "foo bar":
        if len(v9) > 0 {
            v9[0] = 128
        }
    default:
        var v14 string = "foo bar"
        _ = v14
    }
    var v15 uint = (v3 + (v3 + 42))
    fmt.Println(v1)
    var v16 uint16 = 883
    var v17 bool = false
    if v17 {
        var v18 float32 = ((-2.25 - -2.25) + float32(v2))
        v9 = append(v9, 32768)
        switch -9223372036854775808 {
        case 32767:
            v4 += (v4 % ((v4) | 1))
        case 2147483647:
            switch 4294967295 {
            case 32767:
                for fi19 := 0; fi19 < 18; fi19++ {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        fmt.Println(v16)
                        _ = fi20
                    }
                    v9 = append(v9, (fi19 << (((^(v15))) & 63)))
                    _ = len(v9)
                    _ = fi19
                }
            case 599:
                var v21 int32 = 0
                _ = v21
            default:
                v9 = append(v9, 362)
            }
        default:
            var v22 int32 = 1
            _ = v22
        }
        _ = v18
    }
    switch 32768 {
    case -57:
        if (!false) {
            switch 127 {
            case 127:
                if v17 {
                    fmt.Println(v3, v4)
                    var v23 uint16 = 255
                    _ = v23
                } else {
                    if true {
                        _ = len(v9)
                        v24 := GenS2{F0: 2, F1: -1}
                        _ = v24
                    }
                }
            case -233:
                if false {
                    v25 := GenS3{F0: 2147483647, F1: 42}
                    _ = v25
                }
            case -1:
                v17 = ((v16 > v16) || v17)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
