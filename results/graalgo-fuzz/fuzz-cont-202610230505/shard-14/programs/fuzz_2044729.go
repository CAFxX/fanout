package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    var v3 uint8 = (v1 | (v1 % ((127) | 1)))
    var v4 uint16 = 1
    var v5 uint64 = 2
    v6 := v2["foo bar"]
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int16 = -954
        fmt.Println(v4, v5)
        _ = fi7
        _ = v8
    }
    if false {
        v4 = v4
        if false {
            switch 4294967295 {
            case 4294967295:
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := GenS1{F0: 2147483647}
                    _ = fi9
                    _ = v10
                }
            default:
                if true {
                    if true {
                        v2[("" + ("hello" + "hello"))] = 4294967295
                        delete(v2, (((("a" + "") + "foo bar") + "hello") + ("" + "foo bar")))
                        var v11 uint = 252
                        _ = v11
                    } else {
                        v6 = v6
                        var v12 uint64 = v5
                        var v13 string = "hello"
                        _ = v12
                        _ = v13
                    }
                    v3 -= (v1 ^ v1)
                    fmt.Println(v6, v3)
                }
            }
            v14 := []int{4294967295, 228}
            _ = v14
        }
        v4 += v4
    } else {
        v15 := GenS1{F0: -1}
        for fi16 := 0; fi16 < 18; fi16++ {
            if false {
                switch int(v4) {
                case 2147483647:
                    if (v6 <= 612) {
                        if (66 > 72) {
                            v17 := GenS1{F0: 2147483647}
                            v3 = (v1 % ((v3) | 1))
                            var v18 uint16 = v4
                            _ = v17
                            _ = v18
                        }
                    }
                default:
                    fmt.Println(v6, fi16, fi16)
                }
            } else {
                var v19 uint64 = (v5 ^ v5)
                v4 = v4
                _ = v19
            }
            var v20 uint = 2
            _ = fi16
            _ = v20
        }
        _ = v15
    }
    var v21 int64 = 0
    var v22 uint16 = (v4 * v4)
    var v23 int16 = 0
    var v24 int = -9223372036854775808
    var v25 float32 = 3.14159
    v26 := GenS1{F0: 2147483647}
    fmt.Println(v24, v3)
    fmt.Println(v25, v5)
    var v27 int8 = 1
    if (v27 <= (v27 & v27)) {
        var v28 uint8 = (v3 & (^(v1)))
        _ = len(v2)
        _ = v28
    }
    v29 := GenS1{F0: -273}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
