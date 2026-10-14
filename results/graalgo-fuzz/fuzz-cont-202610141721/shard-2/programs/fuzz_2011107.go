package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 string = (v1 + "a")
    var v3 int32 = -2147483648
    var v4 int64 = -9223372036854775808
    var v5 string = ""
    var v6 float32 = -2.25
    var v7 float64 = (1.5 / 1.0)
    var v8 bool = (0 <= 230)
    if v8 {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 434, F1: 65535}
            fmt.Println(v1, v8)
            for fi11 := 0; fi11 < 6; fi11++ {
                if ((32767 > 742) && ((!(-956 == int16(fi9))) || true)) {
                    var v12 uint64 = 18446744073709551615
                    var v13 int = (fi9 | (fi11 * (-(fi9))))
                    v5 = v5
                    _ = v12
                    _ = v13
                } else {
                    var v14 float64 = (v7 * (float64(v6) + float64((1.5 / -2.25))))
                    _ = v14
                }
                var v15 int16 = -1
                _ = fi11
                _ = v15
            }
            _ = fi9
            _ = v10
        }
        if (42 < 592) {
            var v16 uint32 = 65535
            _ = v16
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            v2 = ("hello" + v2)
            _ = fi17
        }
    }
    v7 = ((v7 / ((v7 + float64(float32(float64(float32((float64(v6) + ((0.1 + float64(v6)) - (float64((1e10 / v6)) - (v7 * v7))))))))) / v7)) / 1e10)
    v18 := []int{-1, -44}
    switch 4294967295 {
    case 1:
        v3 = -2147483648
    case -940:
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 int = (fi19 >> ((0) & 63))
            _ = fi19
            _ = v20
        }
    }
    if ((v1 != v2) || v8) {
        v21 := GenS2{F0: v8, F1: 682, F2: 284}
        var v22 string = v2
        _ = v21
        _ = v22
    }
    v8 = v8
    fmt.Println(v5, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
