package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
func genFn0(p0 uint16, p1 string, p2 int) int32 {
    v6 := GenS1{F0: "hello", F1: 2147483647, F2: p1}
    fmt.Println(p0, p2, p0)
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{837, 1, 0}
    v3 := []uint32{162, 381}
    var v4 uint64 = 400
    v5 := make(map[string]uint32)
    _ = len(v5)
    var v8 uint = 2
    if true {
        v8 = 18446744073709551615
        v8 = (v8 << ((v8) & 63))
        var v9 uint = (v8 >> ((287) & 63))
        _ = v9
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            if (525 == 2) {
                v11 := GenS1{F0: "hello", F1: -82, F2: ("" + ("foo bar" + "hello"))}
                _ = v11
            } else {
                if false {
                    v5["\n\t"] = 4294967295
                    var v12 int32 = -1
                    _ = v12
                } else {
                    var v13 uint8 = 2
                    _ = v13
                }
            }
            fmt.Println(v8, fi10)
            if false {
                v4 = (v4 * v4)
            }
            _ = fi10
        }
        v14 := v5["foo bar"]
        if (-9223372036854775808 == 44) {
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint16 = 259
                v14 += (v14 % ((v14) | 1))
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    v14 += 908
                case "a":
                    var v17 int = 32767
                    _ = v17
                default:
                    switch ("foo bar" + (("hello" + "hello") + ("foo bar" + ("\n\t" + "hello")))) {
                    case "\n\t":
                        var v18 int8 = 42
                        _ = v18
                    case "hello":
                        v8 += v8
                    case "":
                        var v19 uint8 = 0
                        _ = v19
                    default:
                        v14 = (^(v14))
                    }
                }
                _ = fi15
                _ = v16
            }
            v20 := GenS1{F0: "hello", F1: 738, F2: "hello"}
            var v21 float64 = (1e10 + 1e10)
            _ = v20
            _ = v21
        }
        _ = v14
    }
    fmt.Println(v4)
    fmt.Println(v8, v4)
    v8 += (v8 - (v8 ^ (v8 + uint(v4))))
    _ = len(v3)
    fmt.Println(v8, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
