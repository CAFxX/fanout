package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 string; F1 int8; F2 uint64 }
func genFn0(p0 bool, p1 bool, p2 string) int8 {
    fmt.Println(p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return -9
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 160
    var v3 uint32 = 2147483648
    v4 := make(map[string]string)
    v5 := []uint32{1}
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 string = "\n\t"
        _ = fi6
        _ = v7
    }
    v1["\n\t"] = 1
    var v8 int64 = 860
    if false {
        var v9 uint32 = (v3 * 0)
        if len(v5) > 0 {
            v5[0] = v3
        }
        _ = v9
    }
    v4[(("foo bar" + "a") + "\n\t")] = ("\n\t" + "hello")
    v10 := GenS1{F0: (v3 * v3)}
    v8 = (v8 ^ (v8 << ((v2) & 63)))
    fmt.Println(v2, v2)
    if true {
        v8 = (v8 & v8)
    } else {
        v2 += v2
        var v11 int8 = 0
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            switch "\n\t" {
            case "a":
                _ = len(v5)
            case "\n\t":
                v14 := GenS2{F0: -1, F1: (^(v3)), F2: 18446744073709551615}
                _ = v14
            case "foo bar":
                if true {
                    var v15 uint64 = 628
                    fmt.Println(v2)
                    _ = v15
                } else {
                    v5 = append(v5, (v3 ^ v3))
                }
                v16 := GenS1{F0: (v3 ^ 2147483648)}
                _ = v16
            }
            var v17 float32 = 3.14159
            _ = len(v1)
            if (true || (("a" + "") > "a")) {
                var v18 uint8 = 94
                if (true || true) {
                    var v19 int = (fi13 % (((-(fi13))) | 1))
                    v18 += (v18 | (v18 % ((127) | 1)))
                    var v20 bool = false
                    _ = v19
                    _ = v20
                }
                _ = v18
            } else {
                var v21 float32 = v17
                if false {
                    fmt.Println(v3)
                    var v22 uint64 = 288
                    _ = len(v5)
                    _ = v22
                } else {
                    if true {
                        if len(v5) > 0 {
                            v5[0] = v3
                        }
                        v5 = append(v5, uint32(v2))
                        var v23 float64 = 0.0
                        _ = v23
                    } else {
                        v24 := GenS3{F0: v2, F1: int32(fi13), F2: 2147483648}
                        _ = v24
                    }
                    var v25 uint8 = 2
                    _ = v25
                }
                _ = v21
            }
            v26 := GenS4{F0: "foo bar", F1: 0, F2: 725}
            _ = fi13
            _ = v17
            _ = v26
        }
        v8 = (-(v8))
        _ = fi12
    }
    if len(v5) > 0 {
        v5[0] = v3
    }
    if true {
        fmt.Println(v3, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
