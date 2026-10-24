package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int32 {
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int16 = -1
            var v9 int32 = 2147483647
            var v10 uint = 2
            switch "a" {
            case "hello":
                switch fi7 {
                case -2147483648:
                    var v11 string = ("foo bar" + "hello")
                    _ = v11
                }
                var v12 int64 = -1
                _ = v12
            case "foo bar":
                var v13 uint64 = 18446744073709551615
                _ = v13
            default:
                v14 := []int{(fi6 ^ fi6), (fi6 | (fi7 << ((4294967295) & 63))), 136}
                _ = v14
            }
            v15 := []int{fi6, fi6, fi7}
            _ = fi7
            _ = v10
            _ = v15
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v16 uint8 = 42
    var v17 int = 81
    _ = v16
    _ = v17
    return 1
}
func genFn1(p0 uint, p1 uint, p2 string) int32 {
    fmt.Println(p1, p2, p1)
    var v18 float32 = 0.1
    v19 := GenS1{F0: -9223372036854775808}
    v20 := []int{652, -2147483648}
    fmt.Println(p1, v18)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return 42
}

func main() {
    var v1 int64 = 1
    var v2 uint = 474
    v3 := []string{"hello"}
    var v4 int16 = 1000
    var v5 int16 = (-(v4))
    var v21 uint64 = 738
    if ((float32(0.1) < 3.14159) || (("hello" + (("foo bar" + "a") + "foo bar")) > ((("a" + "hello") + "") + "a"))) {
        var v22 uint32 = 65535
        _ = v22
    }
    var v23 float64 = (1.5 + 3.14159)
    v1 += v1
    var v24 float32 = 0.0
    var v25 float64 = (v23 * float64((1e10 * 1e10)))
    fmt.Println(v5, v24, v21)
    v5 -= 32767
    v26 := GenS1{F0: 0}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
