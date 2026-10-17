package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0() string {
    var v6 uint = 993
    v7 := GenS1{F0: -128}
    v6 = v6
    switch 127 {
    case -2147483648:
        v6 = 956
    case 9223372036854775807:
        var v8 float64 = 1e10
        _ = v8
    case 1:
        v6 = v6
    default:
        fmt.Println(v6)
    }
    fmt.Println(v6)
    _ = v6
    _ = v7
    return ""
}
func genFn1() uint {
    if false {
        v9 := []int{-1, 474}
        var v10 int64 = 876
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: -128}
        _ = v11
    }
    var v12 uint8 = 2
    v12 = v12
    _ = v12
    return 283
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 168
    v3 := []string{"foo bar", "a", ("a" + ("foo bar" + "a")), ""}
    var v4 string = (("hello" + "a") + ("a" + ("a" + "foo bar")))
    v5 := []int{}
    if false {
        v4 = "foo bar"
        v4 = v4
    } else {
        v13 := v1[v4]
        var v14 uint = 275
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: -128}
    v1[v4] = 143
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if true {
            v16 := GenS2{F0: 184, F1: 882}
            v17 := GenS1{F0: -1}
            _ = v16
            _ = v17
        }
    default:
        v18 := GenS2{F0: 255, F1: 32767}
        _ = v18
    }
    fmt.Println(v4, v2)
    if true {
        if true {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 float32 = (1.5 + (0.1 + (1e10 / 0.1)))
                v21 := GenS3{F0: (("foo bar" + ("hello" + v4)) + v4), F1: 701}
                var v22 string = (v4 + v4)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            if len(v3) > 0 {
                v3[0] = "hello"
            }
        } else {
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 int32 = -2147483648
                _ = fi23
                _ = v24
            }
            switch 127 {
            case 32768:
                var v25 int64 = -351
                _ = v25
            case 9223372036854775807:
                if (0 <= 1) {
                    fmt.Println(v4, v2)
                }
            }
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 uint8 = 2
                _ = fi26
                _ = v27
            }
        }
    } else {
        v4 = v4
        var v28 uint = 2
        if true {
            var v29 int16 = (v2 / ((v2) | 1))
            v30 := v1[v4]
            _ = v29
            _ = v30
        } else {
            v28 = (v28 - v28)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v31 bool = ((v2 == v2) && (1.5 != (0.1 / 1.0)))
            _ = v31
        }
        _ = v28
    }
    fmt.Println(v4, v4)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
