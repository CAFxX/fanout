package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 int8) int64 {
    v4 := []int{42, -948}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = v4
    return 500
}
func genFn1() uint64 {
    var v5 int = 396
    v5 = (v5 * 32767)
    var v6 string = "a"
    v7 := []int{(^(v5))}
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint8 = 228
    v2 := []uint32{4294967295, 872}
    var v3 int8 = 78
    v3 = v3
    if (1 <= 1) {
        switch "\n\t" {
        case "\n\t":
            var v8 float32 = 1e10
            _ = v8
        case "foo bar":
            var v9 uint32 = 65535
            _ = v9
        case "":
            v1 = v1
        default:
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 int8 = v3
                v2 = append(v2, 411)
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 uint = 4294967295
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        }
        v3 += v3
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 := GenS1{F0: (222 == (v1 ^ (v1 ^ v1)))}
    if ("foo bar" >= "\n\t") {
        var v15 string = "\n\t"
        var v16 bool = true
        _ = v15
        _ = v16
    } else {
        v17 := GenS2{F0: 42}
        v3 = 70
        _ = v17
    }
    var v18 int64 = 42
    for fi19 := 0; fi19 < 17; fi19++ {
        fmt.Println(v18, v1, fi19)
        switch (fi19 - fi19) {
        case -926:
            var v20 int64 = v18
            _ = v20
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v21 := GenS2{F0: 1000}
            _ = v21
        }
        _ = fi19
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v18 = (v18 | genFn0(940, (v3 ^ (v3 * v3))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
