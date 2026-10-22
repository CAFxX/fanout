package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
func genFn0(p0 uint32, p1 uint, p2 uint) uint32 {
    var v5 int32 = 42
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = (p1 >> (((p2 | p2)) & 63))
        v8 := []int{(fi6 * (fi6 / (((^(fi6))) | 1)))}
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 -= 255
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 & p0)
}

func main() {
    var v1 float64 = 0.0
    var v2 int = 32767
    var v3 uint8 = 0
    var v4 uint64 = 18446744073709551615
    v9 := []int{(v2 >> ((18446744073709551615) & 63)), v2, (v2 & 127), int(v4)}
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            if (!((250 >= 153) && true)) {
                v2 = (fi11 % ((0) | 1))
                switch (fi11 - -2147483648) {
                case 32768:
                    switch -9223372036854775808 {
                    case 42:
                        var v12 bool = (!(true && (2147483647 != 2147483648)))
                        v13 := GenS1{F0: 65535, F1: fi10, F2: 42}
                        var v14 float32 = ((0.1 - 0.1) * (float32((v1 * (-2.25 * float64((0.1 + 3.14159))))) / (3.14159 / 1.5)))
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                var v15 int16 = 32767
                _ = v15
            }
            var v16 int = v2
            _ = fi11
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int16 = 935
            _ = fi17
            _ = v18
        }
        if ((v2 % (((^(v2))) | 1)) > v2) {
            fmt.Println(v4, fi10, v2)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                if (!("hello" <= "foo bar")) {
                    var v20 int64 = -594
                    _ = v20
                }
                var v21 int8 = 127
                _ = fi19
                _ = v21
            }
        }
        _ = fi10
    }
    v22 := GenS1{F0: 592, F1: v2, F2: 255}
    var v23 int32 = 1
    var v24 bool = true
    var v25 int32 = 285
    v24 = v24
    fmt.Println(v3)
    var v26 int16 = -32768
    var v27 uint = 1
    var v28 uint32 = 2147483648
    var v29 int64 = 327
    v3 += 1
    v4 += v4
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
