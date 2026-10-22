package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
func genFn0(p0 int64) string {
    v6 := []int{}
    var v7 string = "hello"
    v7 = v7
    _ = p0
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 65535
    var v3 bool = true
    v4 := make(map[string]int)
    var v5 bool = (v2 != (^(v2)))
    v2 = v2
    v8 := []int{128, 4294967295, 128}
    v2 = v2
    if (32767 > 932) {
        v5 = false
        var v9 uint = v2
        v10 := GenS1{F0: (!(18446744073709551615 != 1)), F1: -9223372036854775808}
        _ = v10
        _ = v9
    } else {
        v8 = append(v8, 0)
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS1{F0: v3, F1: 398}
        var v13 uint = 18446744073709551615
        _ = fi11
        _ = v12
        _ = v13
    }
    if v5 {
        if len(v8) > 0 {
            v8[0] = 128
        }
        v2 = (v2 + (v2 - v2))
    } else {
        var v14 uint16 = 531
        var v15 int = 801
        _ = v14
        _ = v15
    }
    v2 = (v2 ^ 1)
    fmt.Println(v3, v3, v5)
    v8 = append(v8, 128)
    var v16 string = ("" + "a")
    _ = len(v4)
    _ = len(v8)
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 float32 = (((0.1 * (0.0 + 0.1)) + 1e10) * 0.1)
        v8 = append(v8, fi17)
        v5 = (2147483647 == 2)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 int16 = 0
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS1{F0: v5, F1: fi21}
            var v24 int16 = -1
            v2 = (v2 << ((436) & 63))
            _ = fi22
            _ = v23
            _ = v24
        }
        if (((0.1 * (0.0 / float32(3.14159))) != 0.1) || true) {
            for fi25 := 0; fi25 < 14; fi25++ {
                var v26 int32 = 386
                var v27 int = (^(fi25))
                if ((!(v2 == 2)) || v5) {
                    switch v16 {
                    case "\n\t":
                        _ = len(v4)
                    case "":
                        var v28 uint16 = 65535
                        _ = v28
                    case "a":
                        v3 = (v3 || (false || (v3 || (!false))))
                    default:
                        var v29 uint = (v2 >> ((v2) & 63))
                        _ = v29
                    }
                    var v30 uint64 = 42
                    _ = v30
                } else {
                    v8 = append(v8, fi25)
                    v8 = append(v8, fi21)
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            v2 += (v2 - (v2 << (((^(v2))) & 63)))
        }
        v5 = (v2 > (v2 * 699))
        _ = fi21
    }
    v8 = append(v8, 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
