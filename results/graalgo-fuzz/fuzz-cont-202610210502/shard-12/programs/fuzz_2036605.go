package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 uint32) uint {
    var v7 int16 = 279
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 string, p1 int, p2 int32) uint32 {
    p0 = "foo bar"
    if false {
        p1 -= -828
        p0 = "foo bar"
    } else {
        v8 := GenS1{F0: -179, F1: 1, F2: (!(-128 >= 13))}
        fmt.Println(p0, p0, p1)
        v9 := GenS1{F0: (p1 | (p1 + p1)), F1: 127, F2: (!false)}
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 309
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = 1.5
    var v3 int16 = 235
    var v4 int32 = -958
    var v5 int8 = 127
    v6 := make(map[uint32]bool)
    if true {
        var v10 int8 = (v5 * 100)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 float64 = -2.25
            for fi13 := 0; fi13 < 16; fi13++ {
                if false {
                    var v14 float64 = v2
                    _ = v14
                }
                var v15 uint16 = 65535
                if ((19 > -9223372036854775808) || false) {
                    fmt.Println(v5, v3)
                    delete(v6, 42)
                    var v16 uint16 = v1
                    v17 := GenS1{F0: fi11, F1: 42, F2: (938 <= 2)}
                    _ = v16
                    _ = v17
                }
                if (v5 <= -1) {
                    v15 = v1
                } else {
                    v2 -= (((v12 * -2.25) - v12) / v2)
                    fmt.Println(fi11, v5, v2)
                }
                v4 = v4
                _ = fi13
                _ = v15
            }
            fmt.Println(v1, v5)
            _ = fi11
            _ = v12
        }
    }
    var v18 bool = ((v4 * v4) != -189)
    v4 = v4
    if (("\n\t" <= "foo bar") || (32768 <= 487)) {
        v19 := GenS1{F0: 1, F1: (v5 * v5), F2: v18}
        switch 777 {
        case -1:
            if (genFn0(227, 709) == 213) {
                var v20 int16 = v3
                _ = v20
            } else {
                v21 := GenS2{F0: -9223372036854775808}
                var v22 int = -374
                var v23 uint64 = 18446744073709551615
                _ = v21
                _ = v22
                _ = v23
            }
        }
        _ = v19
    }
    v24 := []int{128, 9223372036854775807}
    v25 := GenS3{F0: 0}
    delete(v6, 67)
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
