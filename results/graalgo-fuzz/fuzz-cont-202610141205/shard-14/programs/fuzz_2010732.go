package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int32 = -288
    var v2 uint = 2
    v3 := []uint32{2}
    var v4 int16 = -1
    var v5 int32 = -1
    v6 := make(map[string]int)
    var v7 int64 = -1
    var v8 uint16 = 1
    v2 += (v2 / ((v2) | 1))
    var v9 float64 = 0.1
    var v10 uint8 = 78
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: 7}
        var v13 float32 = 1e10
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int32 = 771
    _ = len(v3)
    v4 += -392
    if false {
        delete(v6, "foo bar")
        v4 = (v4 ^ (v4 | v4))
        v15 := GenS1{F0: 216}
        _ = v15
    }
    v3 = append(v3, 803)
    if (2 >= uint16(v2)) {
        if false {
            fmt.Println(v9)
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 uint32 = uint32(v2)
                for fi18 := 0; fi18 < 13; fi18++ {
                    var v19 float64 = ((v9 / v9) - (float64((0.1 * (1e10 / 1.0))) - v9))
                    v20 := GenS2{F0: (v17 & 4294967295)}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 float32 = 1.5
                _ = fi16
                _ = v17
                _ = v21
            }
            v10 -= v10
        } else {
            v22 := GenS3{F0: v7}
            var v23 int64 = int64(v8)
            _ = v22
            _ = v23
        }
    } else {
        var v24 uint16 = v8
        _ = v24
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v14, v2)
    case "\n\t":
        fmt.Println(v2, v4, v8)
    case "hello":
        var v25 uint16 = v8
        _ = v25
    default:
        delete(v6, (("hello" + (("foo bar" + ("hello" + ((("\n\t" + "\n\t") + "") + "foo bar"))) + "a")) + ("" + ("" + ""))))
    }
    switch 32768 {
    case 9223372036854775807:
        if ("\n\t" < "") {
            var v26 uint8 = v10
            v14 += (v14 * v5)
            fmt.Println(v7, v1, v7)
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
