package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
func genFn0(p0 uint16) uint16 {
    fmt.Println(p0)
    var v7 string = "\n\t"
    _ = p0
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (-2.25 / 1.5)
    var v3 uint64 = 2
    var v4 int8 = 127
    var v5 float64 = 3.14159
    var v6 string = ("\n\t" + ("\n\t" + "a"))
    fmt.Println(v2, v4)
    v5 += (0.0 * 1e10)
    var v8 int64 = 822
    v9 := GenS1{F0: 42}
    var v10 int8 = (v4 - v4)
    v11 := GenS1{F0: 139}
    if (!false) {
        v5 = v2
        v5 = v2
        var v12 string = v6
        _ = v12
    }
    var v13 float32 = 3.14159
    for fi14 := 0; fi14 < 14; fi14++ {
        switch (fi14 * fi14) {
        case 4294967295:
            var v15 int16 = 1000
            _ = v15
        case 1:
            if false {
                for fi16 := 0; fi16 < 16; fi16++ {
                    if (2 == (v3 << ((65535) & 63))) {
                        var v17 uint = 32
                        _ = v17
                    } else {
                        var v18 bool = (float32(0.0) != (v13 - ((v13 / float32(v5)) * v13)))
                        _ = v18
                    }
                    fmt.Println(v13, fi14, v2)
                    var v19 int64 = (v8 + v8)
                    _ = fi16
                    _ = v19
                }
            } else {
                var v20 int = (-(fi14))
                _ = v20
            }
        case -482:
            v5 += 1.5
        }
        if (true || (451 <= 65535)) {
            v1[(fi14 ^ 42)] = true
            v21 := GenS2{F0: 18446744073709551615, F1: v8, F2: (v10 / (((v4 >> ((871) & 63))) | 1))}
            _ = v21
        }
        _ = fi14
    }
    var v22 float32 = v13
    var v23 uint = 422
    switch "\n\t" {
    case "hello":
        v24 := GenS1{F0: (v23 / (((^(v23))) | 1))}
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
