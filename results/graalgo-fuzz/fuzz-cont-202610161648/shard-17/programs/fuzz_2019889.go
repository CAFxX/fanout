package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 163
    var v3 bool = false
    var v4 int32 = -488
    switch 4294967295 {
    case 869:
        var v5 uint64 = 18446744073709551615
        _ = v5
    case -334:
        for fi6 := 0; fi6 < 6; fi6++ {
            v3 = v3
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
    }
    if v3 {
        var v8 uint64 = 602
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(v3)
            if ((1.5 + 1e10) <= 3.14159) {
                if false {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 string = (((("\n\t" + "\n\t") + ("a" + "a")) + "\n\t") + "hello")
                        v11 = ""
                        _ = fi10
                        _ = v11
                    }
                    _ = len(v1)
                    v12 := []int{(fi9 >> ((534) & 63)), (fi9 | fi9)}
                    _ = v12
                }
                for fi13 := 0; fi13 < 11; fi13++ {
                    v4 = (v4 + v4)
                    _ = fi13
                }
            } else {
                v14 := GenS1{F0: (-(v4)), F1: 42}
                v15 := []int{(fi9 * fi9), (-(fi9))}
                _ = v14
                _ = v15
            }
            var v16 int16 = 827
            _ = fi9
            _ = v16
        }
        _ = v8
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        switch (fi17 << ((42) & 63)) {
        case 32768:
            v18 := GenS1{F0: v4, F1: -28}
            _ = v18
        }
        fmt.Println(fi17, v3, v3)
        _ = fi17
    }
    var v19 string = "a"
    var v20 int = 367
    if false {
        v3 = true
        var v21 uint8 = (v2 * 241)
        _ = v21
    } else {
        v22 := GenS1{F0: (v4 - (v4 + (^(v4)))), F1: 42}
        var v23 int = (v20 * (v20 & (v20 % ((v20) | 1))))
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: v20}
    switch v20 {
    case -2147483648:
        for fi25 := 0; fi25 < 11; fi25++ {
            v19 = v19
            for fi26 := 0; fi26 < 17; fi26++ {
                v20 = 418
                _ = fi26
            }
            _ = fi25
        }
    case -9223372036854775808:
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 uint8 = (v2 - (v2 % ((v2) | 1)))
            _ = fi27
            _ = v28
        }
    case -844:
        v29 := GenS3{F0: (("\n\t" + "\n\t") + ("" + "a")), F1: 2}
        _ = v29
    default:
        var v30 int8 = -109
        _ = v30
    }
    fmt.Println(v4, v4)
    v31 := GenS4{F0: 1}
    for fi32 := 0; fi32 < 8; fi32++ {
        v20 += 42
        v4 = 100
        var v33 int8 = -128
        _ = fi32
        _ = v33
    }
    v19 = (v19 + "foo bar")
    v2 += v2
    v2 = (v2 % (((^(v2))) | 1))
    switch v20 {
    case 0:
        _ = len(v1)
    default:
        v34 := GenS3{F0: ((v19 + v19) + v19), F1: 72}
        _ = v34
    }
    var v35 uint16 = 683
    var v36 float32 = 0.1
    v36 = (3.14159 / 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
}
