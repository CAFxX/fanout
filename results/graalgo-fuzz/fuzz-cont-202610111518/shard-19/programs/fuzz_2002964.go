package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 int8, p1 int64) uint8 {
    var v6 uint8 = 43
    switch -1 {
    case 32767:
        var v7 string = ("" + "foo bar")
        _ = v7
    }
    v8 := GenS1{F0: 42}
    v9 := GenS1{F0: 935}
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    _ = v9
    return v6
}
func genFn1() bool {
    var v10 uint16 = 42
    switch "" {
    case "foo bar":
        if true {
            v11 := []int{127, 32767, 0, 32768}
            fmt.Println(v10)
            if false {
                v10 += v10
            }
            _ = v11
        }
    case "hello":
        switch -2147483648 {
        case 1:
            fmt.Println(v10)
        case 128:
            v10 = v10
        case -582:
            v12 := GenS2{F0: 42}
            _ = v12
        default:
            fmt.Println(v10)
        }
    case "a":
        switch 367 {
        case 149:
            v10 += v10
        case 4294967295:
            var v13 int = 9223372036854775807
            _ = v13
        }
    }
    fmt.Println(v10)
    v10 -= (v10 - (v10 ^ 65535))
    _ = v10
    return true
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 65535
    var v3 float32 = float32(1e10)
    var v4 uint = 2
    v5 := []string{"", ""}
    var v14 string = ("a" + "foo bar")
    for fi15 := 0; fi15 < 2; fi15++ {
        switch 0 {
        case 2147483647:
            switch fi15 {
            case -238:
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (true || ((219 >= 42) && false)) {
                        v17 := GenS3{F0: (v14 + "a"), F1: v4}
                        _ = v17
                    } else {
                        v14 = ""
                    }
                    fmt.Println(v14)
                    _ = fi16
                }
            default:
                fmt.Println(v2, fi15)
            }
        case 0:
            if false {
                var v18 bool = ((1.5 != -2.25) || false)
                _ = v18
            }
        case 861:
            if len(v5) > 0 {
                v5[0] = v14
            }
        default:
            v14 = "\n\t"
        }
        _ = fi15
    }
    var v19 int32 = -1
    delete(v1, 417)
    var v20 int64 = 908
    if true {
        var v21 uint32 = 731
        fmt.Println(v14, v19, v2)
        _ = v21
    } else {
        var v22 uint64 = uint64(v4)
        var v23 int8 = 0
        v24 := GenS2{F0: 128}
        _ = v22
        _ = v23
        _ = v24
    }
    _ = len(v1)
    var v25 bool = true
    v2 = (v2 / ((uint16(v4)) | 1))
    v25 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
