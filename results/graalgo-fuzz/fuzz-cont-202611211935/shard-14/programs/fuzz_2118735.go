package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 string) string {
    var v4 int8 = -104
    v4 += (v4 + (v4 & 100))
    fmt.Println(p0)
    var v5 bool = true
    _ = p0
    _ = v4
    _ = v5
    return p0
}

func main() {
    v1 := []int{}
    var v2 int = -2147483648
    var v3 float64 = float64(-2.25)
    fmt.Println(v3, v2)
    if false {
        var v6 int = (v2 * v2)
        var v7 uint32 = 2
        _ = v6
        _ = v7
    } else {
        _ = len(v1)
    }
    switch "hello" {
    case "a":
        if (true && (1000 <= 129)) {
            var v8 int16 = 1
            _ = v8
        }
    case "\n\t":
        var v9 uint32 = 891
        _ = v9
    default:
        v3 -= v3
    }
    var v10 uint8 = 0
    _ = len(v1)
    fmt.Println(v10, v10)
    if len(v1) > 0 {
        v1[0] = -482
    }
    switch ("foo bar" + "a") {
    case "\n\t":
        switch v2 {
        case 1:
            switch "foo bar" {
            case "foo bar":
                for fi11 := 0; fi11 < 6; fi11++ {
                    if (!(true && true)) {
                        if (("" + "a") > ("\n\t" + ("\n\t" + ""))) {
                            var v12 int8 = -36
                            var v13 uint32 = uint32(v10)
                            _ = v12
                            _ = v13
                        } else {
                            var v14 uint64 = 216
                            _ = v14
                        }
                    } else {
                        fmt.Println(v2, v2)
                        if true {
                            v15 := GenS1{F0: 545, F1: -63, F2: 1}
                            v2 = (^(fi11))
                            _ = v15
                        }
                    }
                    fmt.Println(v3)
                    v2 = (v2 * (v2 | (v2 - fi11)))
                    _ = fi11
                }
            case "a":
                var v16 int32 = 2147483647
                _ = v16
            case "":
                v3 += 1.5
            }
        }
    default:
        if false {
            var v17 string = "\n\t"
            _ = v17
        } else {
            _ = len(v1)
        }
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS2{F0: "foo bar", F1: 65535}
        _ = fi18
        _ = v19
    }
    var v20 int16 = -917
    var v21 uint32 = 2
    if false {
        v22 := GenS3{F0: 42}
        var v23 uint = uint(v21)
        _ = v22
        _ = v23
    }
    fmt.Println(v3)
    if (uint(v10) == 2) {
        if len(v1) > 0 {
            v1[0] = (v2 / (((v2 ^ 4294967295)) | 1))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
