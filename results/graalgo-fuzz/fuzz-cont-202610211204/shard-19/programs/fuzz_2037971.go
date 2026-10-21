package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    v1 := []string{""}
    var v2 string = ("hello" + "")
    var v3 uint = 255
    var v4 uint = (v3 - v3)
    var v5 uint32 = 4294967295
    var v6 float32 = 1e10
    var v7 bool = false
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v5, v3)
        v9 := GenS1{F0: 55}
        _ = fi8
        _ = v9
    }
    if true {
        v10 := GenS1{F0: 9223372036854775807}
        var v11 int = -9223372036854775808
        v2 = v2
        _ = v10
        _ = v11
    }
    switch -9223372036854775808 {
    case -2147483648:
        var v12 float64 = 1.5
        _ = v12
    case 42:
        var v13 int16 = 335
        _ = v13
    default:
        if (v2 > v2) {
            v4 -= 65535
            _ = len(v1)
        }
    }
    v4 -= v4
    v3 -= v3
    v2 = (v2 + "")
    fmt.Println(v6)
    if false {
        if (v5 != v5) {
            var v14 int64 = 815
            var v15 int = -9223372036854775808
            _ = v14
            _ = v15
        } else {
            v3 = v4
            v6 = ((3.14159 / 1.5) * v6)
            v1 = append(v1, "hello")
        }
        v6 = v6
        v1 = append(v1, v2)
    } else {
        var v16 uint64 = 18446744073709551615
        if false {
            if (v6 != v6) {
                v17 := GenS1{F0: 42}
                _ = v17
            }
            v2 = ""
        } else {
            v18 := GenS1{F0: -693}
            v19 := GenS2{F0: v4, F1: 42}
            fmt.Println(v6, v6, v6)
            _ = v18
            _ = v19
        }
        _ = v16
    }
    var v20 int32 = -1
    switch 32768 {
    case 127:
        for fi21 := 0; fi21 < 19; fi21++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            v7 = ((false || ((v5 + v5) >= v5)) && false)
            for fi22 := 0; fi22 < 3; fi22++ {
                v1 = append(v1, "foo bar")
                switch ("\n\t" + v2) {
                case "hello":
                    var v23 int8 = -126
                    _ = v23
                }
                var v24 bool = ((^(v4)) != 18446744073709551615)
                var v25 int64 = int64(fi22)
                if true {
                    _ = len(v1)
                    if (((v5 <= v5) || true) || false) {
                        var v26 int32 = v20
                        v2 = (((v2 + (("hello" + "a") + v2)) + v2) + (v2 + "foo bar"))
                        v24 = ((v20 | v20) < (v20 - (v20 & v26)))
                        var v27 int16 = -178
                        _ = v26
                        _ = v27
                    }
                    v4 = v3
                }
                v4 -= 1
                v7 = v7
                _ = fi22
                _ = v24
                _ = v25
            }
            _ = fi21
        }
    case -411:
        var v28 uint = 1
        _ = v28
    default:
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := GenS2{F0: (v3 % ((v4) | 1)), F1: -340}
            _ = fi29
            _ = v30
        }
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    var v31 uint = v4
    var v32 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
