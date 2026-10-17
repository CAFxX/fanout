package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint; F2 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 uint = 42
    var v4 int8 = -128
    for fi5 := 0; fi5 < 10; fi5++ {
        v3 -= (v3 & 1)
        _ = fi5
    }
    if true {
        v3 += v3
        if (!false) {
            v6 := GenS1{F0: 416}
            _ = v6
        }
        fmt.Println(v4, v3)
    }
    delete(v2, "\n\t")
    var v7 int64 = -73
    v8 := GenS2{F0: false}
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            switch "\n\t" {
            case "\n\t":
                v11 := GenS2{F0: (757 < 1)}
                _ = v11
            case "":
                if false {
                    var v12 int64 = (v7 ^ (v7 % ((v7) | 1)))
                    _ = v12
                } else {
                    v13 := []int{fi10}
                    fmt.Println(v7, v4, fi10)
                    _ = len(v13)
                    _ = v13
                }
            }
            if ((255 > 1) && true) {
                v2["\n\t"] = fi10
                v14 := GenS3{F0: (fi10 % ((-2147483648) | 1)), F1: (v3 & (^(v3))), F2: v7}
                _ = v14
            } else {
                var v15 uint64 = 79
                var v16 int = fi10
                v17 := v2[("\n\t" + ("" + ("foo bar" + "")))]
                v15 = 2
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := []int{(fi10 + (fi9 >> ((v3) & 63))), (fi10 | -9223372036854775808), (-(fi10)), fi10}
            _ = fi10
            _ = v18
        }
        _ = fi9
    }
    var v19 int8 = v4
    delete(v2, "foo bar")
    v20 := []int{-124, -751, 1, 127}
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int32 = 42
        _ = fi21
        _ = v22
    }
    if len(v20) > 0 {
        v20[0] = 497
    }
    v23 := GenS1{F0: 674}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v3 = 477
    var v24 float64 = 1e10
    switch 2147483647 {
    case -425:
        if (3.14159 < 3.14159) {
            switch "foo bar" {
            case "\n\t":
                switch 1 {
                case 127:
                    v24 = float64((-2.25 * 0.0))
                default:
                    switch -951 {
                    case 42:
                        v1["hello"] = ("hello" + ("\n\t" + ("hello" + "foo bar")))
                    default:
                        v3 = (v3 * v3)
                    }
                }
            case "a":
                var v25 int8 = v19
                _ = v25
            case "hello":
                if false {
                    switch 4294967295 {
                    case 0:
                        var v26 uint8 = 81
                        _ = v26
                    case -1:
                        _ = len(v20)
                    }
                    v19 = v4
                    var v27 int16 = 0
                    _ = v27
                }
            }
            fmt.Println(v19, v4, v3)
        }
    case 32767:
        v19 += (v4 << ((4294967295) & 63))
    case 127:
        if true {
            _ = len(v20)
            if (-820 > 42) {
                v20 = append(v20, 4294967295)
                v7 -= v7
            }
        } else {
            v28 := GenS1{F0: 4294967295}
            _ = v28
        }
    }
    if true {
        if true {
            v29 := GenS3{F0: 42, F1: 255, F2: v7}
            switch (("foo bar" + "foo bar") + "") {
            case "hello":
                fmt.Println(v19, v24)
            case "":
                v19 = -1
            default:
                var v30 int32 = 692
                _ = v30
            }
            _ = v29
        } else {
            var v31 int8 = 0
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := GenS3{F0: -9223372036854775808, F1: v3, F2: -1}
            var v34 bool = (false && true)
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v7
    _ = v8
}
