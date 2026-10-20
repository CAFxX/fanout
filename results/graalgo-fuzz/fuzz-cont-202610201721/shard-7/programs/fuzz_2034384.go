package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 int) int16 {
    v6 := []int{(^(p0))}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint = 1
        var v9 string = "foo bar"
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return -533
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 32767
    v3 := make(map[uint32]bool)
    v4 := []string{"hello"}
    var v5 int8 = int8(v2)
    var v10 uint64 = 481
    if false {
        _ = len(v1)
        switch v2 {
        case 127:
            v3[4294967295] = false
        case -2147483648:
            v5 = (v5 | v5)
        case -830:
            var v11 int8 = 100
            _ = v11
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 int32 = -87
                var v14 int = (v2 ^ v2)
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    } else {
        var v15 float64 = 1.5
        _ = v15
    }
    if true {
        var v16 int64 = -168
        var v17 uint = 4294967295
        _ = v16
        _ = v17
    }
    delete(v3, 45)
    fmt.Println(v10)
    if (!(((false && true) || true) && false)) {
        v2 = v2
    } else {
        v18 := GenS1{F0: 9223372036854775807, F1: 32767}
        var v19 uint = 42
        var v20 uint8 = 2
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 9223372036854775807, F1: int16(v5)}
    var v22 string = ("a" + ("" + "\n\t"))
    v5 = (v5 / (((v5 ^ v5)) | 1))
    v4 = append(v4, v22)
    v2 -= (^(v2))
    _ = len(v4)
    if (true && (true || true)) {
        if false {
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS2{F0: (v10 % ((v10) | 1)), F1: v22}
                v5 -= v5
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: 0, F1: 844}
            v22 = v22
            _ = v25
        } else {
            if true {
                if (-290 < 9223372036854775807) {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        fmt.Println(v2)
                        var v27 uint8 = 184
                        fmt.Println(v5, v2, v10)
                        _ = fi26
                        _ = v27
                    }
                } else {
                    fmt.Println(v5, v10)
                    fmt.Println(v10)
                    switch 127 {
                    case -1:
                        v10 = v10
                    case 1:
                        v2 += v2
                    case 32767:
                        fmt.Println(v10, v5, v5)
                    }
                }
            } else {
                var v28 uint64 = (v10 + 2)
                v5 += v5
                _ = v28
            }
        }
    }
    var v29 uint16 = 0
    var v30 uint8 = 42
    v5 += v5
    fmt.Println(v22, v5)
    switch (((v22 + v22) + "hello") + ("hello" + "\n\t")) {
    case "\n\t":
        if ((v2 << ((0) & 63)) >= 2147483647) {
            var v31 int32 = -1000
            _ = v31
        }
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
