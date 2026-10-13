package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint32 }
func genFn0() bool {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: 1, F2: -2147483648}
        var v8 string = (("hello" + "hello") + "\n\t")
        v9 := GenS1{F0: 65535, F1: 42, F2: 825}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 599
    var v11 string = ""
    for fi12 := 0; fi12 < 3; fi12++ {
        v11 = v11
        var v13 float64 = 3.14159
        _ = fi12
        _ = v13
    }
    _ = v10
    _ = v11
    return (((((1.5 - float32(((3.14159 - 1e10) * 0.0))) + -2.25) / 1e10) + (0.0 * (0.1 * 0.0))) == 1.5)
}

func main() {
    var v1 int64 = 42
    var v2 uint64 = 2
    var v3 uint64 = v2
    var v4 int = 32767
    v5 := []string{"hello"}
    if ((v4 << ((18446744073709551615) & 63)) >= (v4 >> ((1) & 63))) {
        v14 := GenS1{F0: 2, F1: 387, F2: -1}
        _ = v14
    } else {
        v1 = -1
        for fi15 := 0; fi15 < 3; fi15++ {
            v1 = v1
            v16 := GenS1{F0: 18446744073709551615, F1: -1, F2: -244}
            var v17 uint8 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v4 += 582
    if true {
        for fi18 := 0; fi18 < 18; fi18++ {
            v1 += 1
            v2 -= v2
            v19 := GenS2{F0: -1}
            _ = fi18
            _ = v19
        }
    }
    var v20 float64 = (3.14159 - (1.5 / 0.1))
    for fi21 := 0; fi21 < 18; fi21++ {
        v4 = v4
        var v22 uint8 = 0
        _ = fi21
        _ = v22
    }
    var v23 int16 = 1000
    var v24 uint32 = 389
    fmt.Println(v23, v23)
    v25 := GenS3{F0: 133, F1: (v24 << ((255) & 63))}
    var v26 bool = (0 == -768)
    var v27 int16 = int16(v4)
    var v28 uint64 = v2
    if v26 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        v3 -= (v2 / ((v3) | 1))
        if v26 {
            fmt.Println(v26, v3, v20)
            if false {
                _ = len(v5)
            } else {
                switch 2147483647 {
                case 4294967295:
                    _ = len(v5)
                default:
                    v5 = append(v5, (("\n\t" + "a") + ("hello" + "foo bar")))
                }
            }
        }
        var v29 string = ("" + "foo bar")
        _ = v29
    }
    fmt.Println(v3, v20)
    var v30 int8 = 42
    var v31 uint32 = (v24 * v24)
    v28 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
