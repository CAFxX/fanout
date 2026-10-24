package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0(p0 uint32, p1 uint32, p2 uint) string {
    var v4 int16 = 0
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 string = "\n\t"
        v4 = v4
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 100, F1: "a"}
    var v8 int64 = 9223372036854775807
    v4 = (v4 >> ((p2) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 bool = false
    var v2 bool = (!false)
    v3 := make(map[int]string)
    if true {
        v2 = (!(1.5 != (1e10 / 3.14159)))
        v1 = (4294967295 < 65535)
        fmt.Println(v1)
    } else {
        v9 := GenS1{F0: 127, F1: (("a" + ((("" + "\n\t") + "\n\t") + (((("\n\t" + "a") + "a") + "a") + "\n\t"))) + "foo bar")}
        var v10 uint8 = 255
        _ = v10
        _ = v9
    }
    v1 = (127 == 49)
    var v11 float64 = 0.0
    fmt.Println(v2, v11)
    var v12 int16 = -228
    v13 := GenS2{F0: 2, F1: v12, F2: 816}
    v1 = v1
    var v14 uint8 = 2
    if (!v1) {
        v15 := GenS2{F0: 4294967295, F1: v12, F2: 65535}
        v11 += 1e10
        if (0 != 255) {
            v11 = (v11 * 3.14159)
        }
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int64 = 0
        _ = fi16
        _ = v17
    }
    var v18 int = -9223372036854775808
    var v19 int8 = -127
    fmt.Println(v2, v12, v18)
    v2 = (!v1)
    if v1 {
        var v20 int8 = (v19 / (((v19 % (((^(v19))) | 1))) | 1))
        v21 := v3[v18]
        if v1 {
            var v22 int32 = 0
            _ = v22
        } else {
            v23 := GenS2{F0: 1, F1: -1, F2: 2}
            for fi24 := 0; fi24 < 3; fi24++ {
                v20 += (-(v20))
                var v25 uint32 = 2
                v26 := []int{v18, fi24, (v18 >> ((4294967295) & 63)), (fi24 + (v18 * v18))}
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(v1)
            _ = v23
        }
        _ = v20
        _ = v21
    } else {
        v27 := GenS2{F0: 211, F1: (-(v12)), F2: 578}
        var v28 uint64 = 18446744073709551615
        _ = v27
        _ = v28
    }
    v14 = 1
    v12 -= v12
    for fi29 := 0; fi29 < 9; fi29++ {
        v14 = (v14 >> ((184) & 63))
        v30 := GenS3{F0: (^(v12)), F1: -9223372036854775808}
        fmt.Println(v1, v18)
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
