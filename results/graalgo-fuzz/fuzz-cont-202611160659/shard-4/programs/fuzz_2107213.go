package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16; F1 bool; F2 uint64 }
func genFn0(p0 int64, p1 bool, p2 uint64) uint64 {
    v6 := []int{9223372036854775807, 4294967295, 9223372036854775807}
    var v7 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 bool = ((127 > 127) && (((!(true || (true && (2147483647 > -1)))) || ("foo bar" >= ("foo bar" + "foo bar"))) && ((!(-128 < -70)) || false)))
    v2 := make(map[int]string)
    var v3 int = 548
    var v4 uint = 846
    v5 := []uint32{845, 0, 42}
    fmt.Println(v3, v4, v3)
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 9; fi9++ {
        delete(v2, fi9)
        var v10 float64 = 1e10
        v4 = 214
        _ = fi9
        _ = v10
    }
    fmt.Println(v3, v8)
    fmt.Println(v1)
    v11 := GenS1{F0: (^(v4)), F1: 255, F2: int8(v8)}
    fmt.Println(v1, v3, v8)
    switch "hello" {
    case "":
        if (v1 && (!(-2.25 != 3.14159))) {
            if v1 {
                var v12 uint64 = 824
                if ("a" < "") {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v3 = 318
                        _ = fi13
                    }
                    var v14 uint32 = 804
                    _ = v14
                }
                _ = v12
            }
            var v15 float64 = 3.14159
            _ = v15
        }
    case "\n\t":
        var v16 bool = (!(genFn0(9223372036854775807, (65535 <= 0), 42) > 4294967295))
        _ = v16
    }
    if v1 {
        var v17 uint = 202
        var v18 uint32 = 1
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 uint64 = 2
            _ = fi19
            _ = v20
        }
        var v21 int16 = 0
        v22 := GenS2{F0: v21, F1: true, F2: genFn0(0, ((v3 % ((v3) | 1)) < -9223372036854775808), genFn0(94, (false || (!(v21 < v21))), 389))}
        _ = v21
        _ = v22
    }
    var v23 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
