package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint8, p1 int16, p2 uint) int64 {
    p1 -= 32767
    fmt.Println(p2, p2)
    var v6 int16 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    var v3 float32 = 3.14159
    v4 := make(map[int]string)
    var v5 bool = (3.14159 != v3)
    var v7 int = -130
    v8 := v1[2147483648]
    for fi9 := 0; fi9 < 9; fi9++ {
        switch (fi9 - (fi9 / ((v7) | 1))) {
        case 128:
            v2 = append(v2, "a")
        case -1:
            v8 = "\n\t"
        }
        v5 = v5
        _ = fi9
    }
    switch v8 {
    case "":
        var v10 int8 = -2
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: (^(fi11)), F1: -2147483648}
        if v5 {
            if (!v5) {
                if v5 {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        fmt.Println(v8, fi11)
                        v3 = v3
                        _ = fi13
                    }
                    fmt.Println(v5)
                }
                var v14 uint = 65535
                var v15 uint = 4294967295
                _ = v14
                _ = v15
            }
            var v16 uint = 2
            var v17 int8 = 51
            _ = v16
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    v8 = (v8 + v8)
    if true {
        v8 = "hello"
        v5 = ((!v5) && ("a" <= v8))
        v5 = ((v5 && (!true)) && ((v3 + (v3 / float32(0.1))) == ((float32(((0.0 + 1.5) - 0.0)) / v3) * 1.5)))
    } else {
        var v18 uint16 = 1
        _ = v18
    }
    v1[1] = v8
    if ((v7 | (v7 | v7)) < 401) {
        v3 -= (3.14159 - (-2.25 - 3.14159))
        v19 := GenS1{F0: -524, F1: 0}
        var v20 bool = (-2.25 < 0.1)
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: 2147483647, F1: 0}
    v22 := GenS2{F0: 461, F1: -9223372036854775808}
    var v23 int32 = 0
    if v5 {
        switch (v7 ^ v7) {
        case -461:
            v7 -= (v7 + (^(v7)))
        }
    } else {
        if (v5 || v5) {
            v5 = v5
        } else {
            var v24 string = "foo bar"
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 uint = 18446744073709551615
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        switch v7 {
        case 0:
            if true {
                v3 = 0.0
            }
        case 1:
            var v27 int16 = 962
            _ = v27
        case -9223372036854775808:
            var v28 bool = (v5 || v5)
            _ = v28
        default:
            _ = len(v1)
        }
        fmt.Println(v5, v7, v7)
    }
    if v5 {
        fmt.Println(v8, v7, v3)
        _ = len(v1)
        v7 += (v7 | (v7 % (((v7 * 0)) | 1)))
    }
    v29 := v4[313]
    v30 := GenS1{F0: (v7 * (v7 * v7)), F1: (v23 ^ 2147483647)}
    var v31 uint16 = 274
    v8 = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
