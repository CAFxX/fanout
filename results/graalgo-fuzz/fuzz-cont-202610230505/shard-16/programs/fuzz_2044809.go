package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint16, p2 uint16) uint {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p0, p1, p0)
        _ = fi7
    }
    v8 := []int{2147483647, 128, -178, 9223372036854775807}
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}
func genFn1(p0 int8, p1 uint8) uint {
    v9 := []int{}
    fmt.Println(p0, p1)
    if true {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 476
}

func main() {
    v1 := []int{421, 42, -523, -9223372036854775808}
    var v2 bool = true
    v3 := make(map[uint32]bool)
    var v4 int = 42
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]string)
    for fi10 := 0; fi10 < 18; fi10++ {
        if len(v1) > 0 {
            v1[0] = v4
        }
        _ = len(v5)
        _ = fi10
    }
    v2 = ((449 >= 18446744073709551615) && (642 == 2147483648))
    if (-1 < -589) {
        v4 += v4
    } else {
        var v11 uint32 = 1
        v5[v11] = "a"
        _ = len(v3)
        _ = v11
    }
    var v12 float32 = 0.1
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int64 = 0
        _ = fi13
        _ = v14
    }
    var v15 string = ("a" + "foo bar")
    fmt.Println(v2, v12)
    v2 = v2
    if v2 {
        v16 := GenS1{F0: 1}
        switch v4 {
        case 1:
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 uint32 = 42
                    v5[0] = "a"
                    v20 := GenS2{F0: 2147483647}
                    v12 = 1e10
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = fi17
            }
        case 9223372036854775807:
            switch v15 {
            case "foo bar":
                v2 = false
            case "\n\t":
                v12 += -2.25
            case "hello":
                v12 -= 3.14159
            default:
                for fi21 := 0; fi21 < 9; fi21++ {
                    if true {
                        fmt.Println(v15, v2, v12)
                        switch (fi21 >> ((4294967295) & 63)) {
                        case -2147483648:
                            var v22 uint8 = 192
                            _ = v22
                        case 2147483647:
                            var v23 uint64 = 4294967295
                            _ = v23
                        case 32768:
                            var v24 uint32 = 2147483647
                            _ = v24
                        default:
                            fmt.Println(fi21, v4)
                        }
                    } else {
                        var v25 uint64 = 4294967295
                        _ = v25
                    }
                    _ = len(v1)
                    if (!(694 > 42)) {
                        var v26 uint8 = 98
                        v15 = "hello"
                        _ = v26
                    }
                    _ = fi21
                }
            }
        case 32767:
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 int64 = 9223372036854775807
                _ = fi27
                _ = v28
            }
        default:
            var v29 int = (v4 * (v4 / ((127) | 1)))
            _ = v29
        }
        _ = v16
    }
    v12 -= 1e10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        delete(v5, 2147483647)
        _ = fi30
    }
    v4 -= (v4 ^ v4)
    v12 = 1.5
    if v2 {
        v31 := GenS1{F0: 100}
        v6[2] = ((("" + (("a" + v15) + ("" + "hello"))) + ((v15 + v15) + (v15 + ""))) + v15)
        _ = v31
    }
    var v32 int = (v4 - 128)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
