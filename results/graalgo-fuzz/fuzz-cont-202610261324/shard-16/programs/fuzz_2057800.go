package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int64) int8 {
    p0 += p0
    v6 := GenS1{F0: "foo bar"}
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 uint = 0
    var v2 string = ("foo bar" + "hello")
    var v3 uint16 = 1
    var v4 bool = false
    v5 := make(map[int]bool)
    v4 = true
    delete(v5, 42)
    switch -2147483648 {
    case -1:
        if v4 {
            v7 := GenS2{F0: 42, F1: 2147483647, F2: 42}
            _ = v7
        }
    case 2147483647:
        if v4 {
            v1 = 1
        }
    case 688:
        _ = len(v5)
    }
    fmt.Println(v3)
    var v8 uint64 = 42
    v2 = "hello"
    v1 = 1
    var v9 int32 = 1
    v3 = (v3 / ((v3) | 1))
    v3 = 255
    if (v1 >= (v1 / (((v1 ^ 18446744073709551615)) | 1))) {
        if true {
            v10 := GenS2{F0: 0, F1: (v9 * (v9 >> (((v1 + (v1 | (v1 / (((^(v1))) | 1))))) & 63))), F2: 4294967295}
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 float64 = 0.1
                _ = fi11
                _ = v12
            }
            v8 += 247
            _ = v10
        }
        if (v4 && ((false || (!(203 <= 127))) && (!v4))) {
            switch ("\n\t" + v2) {
            case "hello":
                delete(v5, int(v3))
            case "foo bar":
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 int64 = -9223372036854775808
                    _ = fi13
                    _ = v14
                }
            }
            if ((v4 || false) && v4) {
                v15 := []int{374, -419, 32768}
                v1 += 18446744073709551615
                _ = v15
            } else {
                switch 4294967295 {
                case 32768:
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := []int{32768}
                        v18 := GenS3{F0: v3}
                        if v4 {
                            fmt.Println(v2)
                            v19 := GenS4{F0: -117}
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 20; fi20++ {
                                var v21 string = (v2 + (((v2 + ("\n\t" + "")) + "\n\t") + "a"))
                                var v22 float32 = float32(1e10)
                                v2 = v2
                                var v23 uint8 = 177
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            v1 += (v1 * (v1 << ((v1) & 63)))
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                case 347:
                    fmt.Println(v1, v4, v3)
                default:
                    var v24 int32 = 2147483647
                    _ = v24
                }
            }
        }
        var v25 float64 = (-2.25 / 1e10)
        _ = v25
    } else {
        v4 = v4
        fmt.Println(v4)
    }
    v26 := []int{4294967295, 642}
    switch -2147483648 {
    case 9223372036854775807:
        var v27 int64 = 1
        _ = v27
    case 32768:
        v28 := GenS2{F0: 675, F1: v9, F2: 1}
        _ = v28
    case -9223372036854775808:
        v9 = v9
    default:
        var v29 uint16 = (v3 * v3)
        _ = v29
    }
    v8 = (v8 - (v8 ^ 18446744073709551615))
    _ = len(v26)
    fmt.Println(v3, v4)
    var v30 string = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
