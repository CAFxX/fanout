package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64) int32 {
    var v7 uint = 152
    v8 := []int{2147483647, 9223372036854775807, 2147483647}
    _ = p0
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint8, p1 uint64, p2 uint16) uint32 {
    var v9 uint = 1
    v9 += v9
    if (!true) {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 float32 = -2.25
            _ = fi10
            _ = v11
        }
        switch -740 {
        case -344:
            fmt.Println(p1, p0, v9)
        }
    } else {
        var v12 int16 = 42
        _ = v12
    }
    switch ("hello" + "foo bar") {
    case "hello":
        for fi13 := 0; fi13 < 10; fi13++ {
            if (!(((v9 * v9) >= (v9 | (v9 & 65535))) && true)) {
                var v14 uint32 = 2147483648
                v15 := GenS1{F0: int32(v14)}
                _ = v14
                _ = v15
            } else {
                v16 := GenS1{F0: 2147483647}
                _ = v16
            }
            var v17 uint = v9
            _ = fi13
            _ = v17
        }
    case "a":
        fmt.Println(p2)
    case "foo bar":
        var v18 string = ("foo bar" + "\n\t")
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 609
    var v2 uint8 = 2
    var v3 float32 = float32(0.0)
    var v4 uint16 = 615
    var v5 int64 = v1
    var v6 uint64 = 18446744073709551615
    var v19 int32 = 0
    v20 := GenS2{F0: 57, F1: ("" + (("foo bar" + "foo bar") + ("\n\t" + ("hello" + "hello")))), F2: v5}
    fmt.Println(v4, v19)
    var v21 int16 = 412
    var v22 bool = ((!true) || ((v19 % (((v19 % (((v19 << ((4294967295) & 63))) | 1))) | 1)) == v19))
    fmt.Println(v22)
    var v23 int16 = v21
    v24 := []int{32768, int(v1), 419}
    fmt.Println(v23, v4, v21)
    if v22 {
        if (12 == -87) {
            v19 -= v19
            fmt.Println(v21)
            v25 := GenS3{F0: 581}
            _ = v25
        }
    }
    v26 := GenS1{F0: (v19 / ((-1) | 1))}
    v6 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
