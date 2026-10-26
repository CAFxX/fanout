package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32, p1 int32, p2 uint64) int16 {
    switch -174 {
    case -358:
        fmt.Println(p0, p0)
    case 32768:
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: (p2 << ((42) & 63))}
            _ = fi7
            _ = v8
        }
    case -9223372036854775808:
        v9 := GenS2{F0: 716}
        _ = v9
    }
    var v10 int64 = -669
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    v1 := []string{"foo bar", "\n\t"}
    var v2 uint8 = 255
    var v3 uint32 = 196
    var v4 uint32 = v3
    v5 := make(map[string]string)
    v6 := []string{("\n\t" + "foo bar"), (("\n\t" + "hello") + "a"), "a", "foo bar"}
    v2 = (^(v2))
    var v11 string = "foo bar"
    v2 = v2
    var v12 uint64 = 2
    fmt.Println(v3, v4, v2)
    v4 += v4
    v13 := GenS3{F0: v11, F1: v12, F2: 1}
    var v14 int32 = -657
    v2 -= (v2 * v2)
    fmt.Println(v3)
    v15 := GenS1{F0: v12}
    v4 -= 4294967295
    v16 := GenS4{F0: 0}
    v17 := GenS3{F0: v11, F1: v12, F2: 1}
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "a" {
        case "":
            if len(v6) > 0 {
                v6[0] = (v11 + ("foo bar" + v11))
            }
        default:
            var v19 int32 = v14
            _ = v19
        }
        fmt.Println(v12, fi18)
        _ = fi18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
