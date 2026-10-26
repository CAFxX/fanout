package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 bool, p2 int16) uint64 {
    p2 -= (p2 % (((p2 >> ((0) & 63))) | 1))
    switch -614 {
    case -2147483648:
        fmt.Println(p1)
    case 42:
        v7 := []int{-924, 1, -667}
        _ = v7
    }
    if p1 {
        fmt.Println(p0, p2)
        if ((((((0.0 * 1.5) * float32(0.0)) + 1e10) + 0.0) < 3.14159) || false) {
            p0 = -957
            if ((!p1) || true) {
                switch -2147483648 {
                case 929:
                    fmt.Println(p1, p2, p0)
                case -257:
                    var v8 int8 = 1
                    _ = v8
                case -817:
                    for fi9 := 0; fi9 < 9; fi9++ {
                        v10 := []int{-773, fi9, 128}
                        for fi11 := 0; fi11 < 19; fi11++ {
                            v12 := GenS1{F0: fi9, F1: "foo bar", F2: -129}
                            _ = fi11
                            _ = v12
                        }
                        _ = fi9
                        _ = v10
                    }
                }
            }
            var v13 uint8 = 255
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 784
}

func main() {
    var v1 uint16 = 2
    var v2 string = ("" + "")
    v3 := make(map[string]string)
    var v4 uint = 42
    var v5 int16 = -297
    var v6 uint16 = (v1 + (v1 << ((779) & 63)))
    v5 = v5
    var v14 int8 = 1
    v1 = v1
    fmt.Println(v2, v2)
    if (!(true && (!true))) {
        var v15 int64 = 42
        _ = v15
    } else {
        v16 := []int{}
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int32 = -1
        _ = fi17
        _ = v18
    }
    switch v2 {
    case "\n\t":
        var v19 int = -223
        _ = v19
    default:
        v20 := GenS1{F0: 842, F1: v2, F2: 1}
        _ = v20
    }
    var v21 uint8 = 1
    v22 := GenS2{F0: 0}
    if (!false) {
        for fi23 := 0; fi23 < 20; fi23++ {
            for fi24 := 0; fi24 < 20; fi24++ {
                if false {
                    switch v2 {
                    case "a":
                        v25 := GenS2{F0: 574}
                        _ = v25
                    default:
                        v26 := GenS2{F0: -334}
                        _ = v26
                    }
                    fmt.Println(v21, v21)
                }
                v27 := []int{(fi24 * (-(fi24))), (fi23 | fi24)}
                _ = fi24
                _ = v27
            }
            v2 = ((("" + v2) + v2) + v2)
            _ = fi23
        }
    } else {
        var v28 int16 = 32767
        var v29 uint32 = 39
        _ = v28
        _ = v29
    }
    fmt.Println(v6, v1)
    var v30 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
