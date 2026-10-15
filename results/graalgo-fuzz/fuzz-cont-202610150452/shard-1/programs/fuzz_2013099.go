package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []string{"hello"}
    v2 := make(map[uint32]string)
    v3 := []string{}
    var v4 float32 = 3.14159
    var v5 uint32 = 1
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 float64 = 1.5
        v8 := GenS1{F0: fi6}
        var v9 uint32 = (v5 - 2147483648)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(v4, v4)
    _ = len(v2)
    switch "foo bar" {
    case "\n\t":
        delete(v2, (v5 * v5))
    case "a":
        v5 = 2147483648
    default:
        if (!(42 != 42)) {
            var v10 uint64 = 183
            _ = v10
        }
    }
    if len(v3) > 0 {
        v3[0] = (("foo bar" + "a") + ("hello" + ("hello" + ("foo bar" + "foo bar"))))
    }
    var v11 float64 = float64(float32(-2.25))
    if false {
        v12 := GenS1{F0: 127}
        v13 := GenS1{F0: 32767}
        _ = v12
        _ = v13
    } else {
        v11 += (v11 + v11)
        v14 := GenS1{F0: -9223372036854775808}
        var v15 uint64 = 1
        _ = v14
        _ = v15
    }
    var v16 int32 = -1
    v17 := GenS2{F0: 42, F1: 128, F2: 255}
    v5 += v5
    if (!(723 != 177)) {
        fmt.Println(v11)
        delete(v2, 2147483648)
    } else {
        if true {
            v18 := GenS1{F0: -253}
            _ = v18
        } else {
            v19 := GenS3{F0: (^(v5))}
            v20 := GenS1{F0: 462}
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            if (!false) {
                v16 = (v16 - v16)
                fmt.Println(v16)
            }
            v2[0] = "hello"
            _ = fi21
        }
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            delete(v2, v5)
            v4 -= -2.25
            var v24 uint32 = 2147483648
            _ = fi23
            _ = v24
        }
        var v25 bool = (true && (1 == -1))
        _ = fi22
        _ = v25
    }
    v11 = v11
    switch -452 {
    case 563:
        var v26 int = -2147483648
        _ = v26
    case -359:
        v27 := GenS4{F0: 65535}
        _ = v27
    default:
        v16 -= (-(v16))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
