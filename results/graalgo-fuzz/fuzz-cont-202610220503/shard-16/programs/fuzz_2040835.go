package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16, p1 uint64) uint32 {
    var v7 int32 = -1
    var v8 uint16 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 57
}
func genFn1() int16 {
    v9 := GenS1{F0: 738}
    if (((-921 < -281) || (!true)) || true) {
        var v10 string = (("foo bar" + ("a" + "a")) + ("" + "hello"))
        for fi11 := 0; fi11 < 1; fi11++ {
            switch "\n\t" {
            case "\n\t":
                v12 := GenS1{F0: 949}
                _ = v12
            case "hello":
                v10 = v10
            default:
                v10 = v10
            }
            _ = fi11
        }
        _ = v10
    }
    var v13 uint32 = 2147483647
    var v14 uint = 567
    _ = v13
    _ = v14
    _ = v9
    return 343
}

func main() {
    var v1 uint64 = 1
    var v2 uint64 = v1
    v3 := make(map[int]string)
    var v4 uint32 = 589
    var v5 uint32 = v4
    var v6 uint32 = (^(v5))
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(fi15, v2, v2)
        v2 -= 1
        v6 += (v6 ^ v5)
        _ = fi15
    }
    var v16 uint = 18446744073709551615
    var v17 float64 = 0.1
    if (true && false) {
        var v18 int = 2147483647
        var v19 int32 = -2147483648
        v20 := GenS1{F0: 0}
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 bool = false
    v2 = v2
    var v22 bool = ((("\n\t" + "foo bar") <= (("\n\t" + (("foo bar" + "foo bar") + ("" + "foo bar"))) + "a")) || v21)
    var v23 uint32 = v4
    delete(v3, 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
