package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint; F2 string }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int = -9223372036854775808
    var v2 int8 = 127
    var v3 string = "\n\t"
    var v4 uint32 = 65535
    var v5 uint64 = 18446744073709551615
    v6 := []uint32{(v4 >> ((255) & 63)), (v4 ^ v4), v4, v4}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint32 = v4
        fmt.Println(fi7, v5, v8)
        _ = fi7
        _ = v8
    }
    var v9 int32 = -210
    if (!true) {
        var v10 uint64 = v5
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 float32 = -2.25
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        v13 := GenS1{F0: 1, F1: (v2 >> ((0) & 63)), F2: 255}
        _ = v13
    }
    v14 := GenS1{F0: 255, F1: (-(v2)), F2: 201}
    v15 := GenS2{F0: 2147483648}
    v1 = int(v2)
    var v16 float32 = 0.1
    v17 := GenS2{F0: v4}
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            v1 = fi19
            _ = fi19
        }
        if true {
            fmt.Println(v16, v4)
            v4 -= (^(v4))
            v20 := GenS3{F0: 65535, F1: 255, F2: (v3 + "foo bar")}
            _ = v20
        } else {
            fmt.Println(v1, v9)
        }
        v21 := GenS4{F0: (-(v2))}
        _ = fi18
        _ = v21
    }
    v1 = v1
    var v22 uint64 = v5
    var v23 uint32 = 827
    var v24 int8 = (^(v2))
    switch v3 {
    case "":
        v6 = append(v6, v4)
    case "hello":
        if len(v6) > 0 {
            v6[0] = v4
        }
    }
    v6 = append(v6, (v23 - (v4 - 997)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
