package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int16, p2 int) uint32 {
    p2 = p2
    var v6 float32 = 0.1
    v7 := GenS1{F0: p0, F1: 127}
    v8 := GenS1{F0: p0, F1: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 843
}

func main() {
    var v1 string = "hello"
    var v2 float32 = (float32(3.14159) + -2.25)
    var v3 float32 = float32(1e10)
    v4 := make(map[string]bool)
    var v5 int = 32767
    v2 += ((v3 * 1e10) / float32((((float64(1.5) + 0.1) / 1e10) + (0.1 - 1e10))))
    var v9 int32 = 854
    v10 := GenS1{F0: 18446744073709551615, F1: (v5 % ((-2147483648) | 1))}
    if false {
        v1 = (v1 + ((("foo bar" + (("\n\t" + (("\n\t" + v1) + v1)) + v1)) + v1) + v1))
        v11 := GenS1{F0: 2, F1: v5}
        switch "\n\t" {
        case "a":
            v12 := []int{851, (v5 >> ((841) & 63)), (v5 << ((42) & 63))}
            _ = v12
        case "foo bar":
            v1 = (((v1 + v1) + ((v1 + "\n\t") + ("a" + (v1 + ("" + "\n\t"))))) + (("\n\t" + "a") + "\n\t"))
        default:
            v13 := GenS1{F0: 139, F1: (v5 / ((127) | 1))}
            _ = v13
        }
        _ = v11
    } else {
        var v14 int64 = 42
        v15 := GenS2{F0: 1}
        fmt.Println(v2, v1, v1)
        _ = v14
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 float32 = (1.5 - ((0.1 / v3) / v3))
            _ = fi16
            _ = v17
        }
        var v18 bool = (((true || (!(-1 == -1))) || (4 < 304)) || (2 != genFn0(954, 696, v5)))
        v19 := GenS1{F0: 155, F1: v5}
        _ = v18
        _ = v19
    }
    v5 += (v5 % (((v5 / ((v5) | 1))) | 1))
    if (426 < 32767) {
        v5 += 676
    } else {
        var v20 int16 = 0
        var v21 bool = (false || ((1e10 >= v2) || true))
        v22 := v4["hello"]
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int64 = 794
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 uint = 0
        var v26 uint64 = 1
        fmt.Println(v3, v9, v2)
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS3{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
