package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool) string {
    v5 := GenS1{F0: 65535, F1: ("" + "foo bar")}
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v6, p0)
        var v8 int64 = -754
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return ("\n\t" + ("foo bar" + "hello"))
}
func genFn1() int16 {
    v9 := []int{458, -2147483648, -201, 32767}
    var v10 int8 = -128
    v10 = v10
    var v11 uint32 = 1
    _ = v10
    _ = v11
    _ = v9
    return 598
}

func main() {
    var v1 uint8 = 94
    var v2 uint64 = 0
    var v3 bool = false
    var v4 string = "hello"
    var v12 uint = 18446744073709551615
    v4 = v4
    v13 := []int{128, -2147483648, 980, -2147483648}
    v3 = (!((-32768 == 342) || true))
    v14 := GenS1{F0: 495, F1: "\n\t"}
    var v15 int8 = 127
    v2 += (^(v2))
    var v16 uint32 = 2
    var v17 string = v4
    if false {
        v4 = v17
        v17 = ((v4 + "") + (v4 + v4))
        var v18 int64 = 516
        _ = v18
    }
    if v3 {
        v19 := GenS2{F0: int32(v16)}
        v3 = false
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        _ = len(v13)
        switch ((v4 + "foo bar") + v4) {
        case "\n\t":
            if (!(-1 <= 2147483647)) {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS1{F0: (^(v16)), F1: ""}
                    _ = fi21
                    _ = v22
                }
            }
        }
        if len(v13) > 0 {
            v13[0] = int(v1)
        }
        _ = fi20
    }
    fmt.Println(v4, v15)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v15 = (v15 / ((v15) | 1))
    v16 = (v16 >> ((uint(v1)) & 63))
    var v23 uint8 = (^(v1))
    if len(v13) > 0 {
        v13[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
