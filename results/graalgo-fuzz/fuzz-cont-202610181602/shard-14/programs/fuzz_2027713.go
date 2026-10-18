package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int, p1 uint8, p2 string) uint {
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 239, F1: 65535, F2: 1000}
        for fi7 := 0; fi7 < 6; fi7++ {
            p1 = p1
            _ = fi7
        }
        p0 -= (p0 - (^(fi5)))
        _ = fi5
        _ = v6
    }
    v8 := GenS2{F0: 1}
    var v9 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int8 = 100
    var v2 uint32 = 4294967295
    v3 := []int{9223372036854775807, -1, -28, -9223372036854775808}
    v4 := []int{42, -9223372036854775808}
    if false {
        v10 := GenS1{F0: 9223372036854775807, F1: v2, F2: 2147483647}
        var v11 uint16 = 42
        v11 = (v11 + v11)
        _ = v10
        _ = v11
    } else {
        var v12 float64 = -2.25
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= 127
    var v13 uint = 65535
    v2 = v2
    if (!(42 != 42)) {
        if false {
            switch "hello" {
            case "":
                v2 -= v2
            case "foo bar":
                v13 += (v13 << ((genFn0(0, 0, "a")) & 63))
            }
        } else {
            if (!true) {
                v3 = append(v3, 0)
                var v14 string = ("hello" + ("foo bar" + ((("" + "") + ("foo bar" + "foo bar")) + "")))
                _ = v14
            } else {
                fmt.Println(v1, v1)
            }
            var v15 uint64 = 1
            _ = v15
        }
        v2 -= (v2 | v2)
        for fi16 := 0; fi16 < 8; fi16++ {
            v13 += v13
            var v17 int16 = 413
            _ = fi16
            _ = v17
        }
    }
    v1 += (v1 ^ v1)
    var v18 int32 = -2147483648
    if len(v4) > 0 {
        v4[0] = 0
    }
    fmt.Println(v18, v1, v1)
    var v19 float32 = ((3.14159 * 3.14159) * float32((((3.14159 - 3.14159) + -2.25) + float64((-2.25 * 0.1)))))
    if (v2 >= (^(v2))) {
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v2, fi20)
            _ = fi20
        }
        v1 += (v1 / ((v1) | 1))
    }
    v21 := GenS3{F0: 1}
    var v22 uint16 = 373
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
