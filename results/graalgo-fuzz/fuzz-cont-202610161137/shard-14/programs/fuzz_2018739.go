package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int, p1 uint32, p2 uint16) uint16 {
    var v7 uint16 = (^(p2))
    fmt.Println(p0, p1, p1)
    var v8 int = 32767
    var v9 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := []string{("\n\t" + "a"), "hello", "foo bar", ""}
    v2 := []int{}
    var v3 float32 = float32(3.14159)
    v4 := []string{("a" + "a"), ("\n\t" + ("" + "foo bar")), "a", "\n\t"}
    var v5 int8 = -117
    v6 := make(map[int]bool)
    switch 4294967295 {
    case 128:
        var v10 int16 = 0
        _ = v10
    case 4294967295:
        if (-572 == 127) {
            var v11 int32 = -1
            var v12 int8 = -94
            for fi13 := 0; fi13 < 13; fi13++ {
                switch fi13 {
                case 1:
                    var v14 int16 = 32767
                    _ = v14
                default:
                    _ = len(v2)
                }
                v5 = (v5 << ((18446744073709551615) & 63))
                _ = fi13
            }
            _ = v11
            _ = v12
        }
    }
    v15 := GenS1{F0: 2}
    var v16 float64 = float64((v3 + v3))
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 float32 = v3
        v3 = float32(3.14159)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v3 = (((v3 + 3.14159) - v3) * 1e10)
        _ = fi19
    }
    var v20 uint = 1
    var v21 bool = true
    var v22 int64 = -9223372036854775808
    v21 = v21
    var v23 float64 = 3.14159
    v23 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
