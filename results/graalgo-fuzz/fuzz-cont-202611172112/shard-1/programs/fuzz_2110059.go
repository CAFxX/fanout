package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int, p1 int32, p2 int8) bool {
    fmt.Println(p1, p0)
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return ("foo bar" >= "")
}
func genFn1(p0 bool) uint8 {
    switch 9223372036854775807 {
    case 4294967295:
        p0 = p0
    }
    p0 = p0
    var v5 int64 = 0
    v5 += 911
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    var v6 int = 1
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 = v6
        _ = fi7
    }
    var v8 int16 = -1
    v1 = (^(v1))
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    fmt.Println(v8, v8, v6)
    var v11 uint = 2
    if (!false) {
        v6 = v6
    } else {
        var v12 int = v6
        if true {
            _ = len(v3)
            if false {
                switch (v6 / ((2147483647) | 1)) {
                case 363:
                    var v13 int8 = -128
                    _ = v13
                case 4294967295:
                    v8 -= (^(v8))
                case -999:
                    v14 := GenS1{F0: 598}
                    _ = v14
                default:
                    v6 = v6
                }
                fmt.Println(v11, v8)
                var v15 float64 = v4
                _ = v15
            } else {
                v16 := GenS1{F0: 0}
                for fi17 := 0; fi17 < 10; fi17++ {
                    if (false || false) {
                        v18 := GenS1{F0: 42}
                        _ = v18
                    }
                    _ = fi17
                }
                _ = v16
            }
        }
        _ = v12
    }
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
