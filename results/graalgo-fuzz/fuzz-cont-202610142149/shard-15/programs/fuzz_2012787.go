package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int {
    var v7 uint32 = 0
    var v8 int8 = -128
    v9 := []int{}
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (0.0 - float32(1.5))
    var v3 float32 = 0.0
    var v4 string = "hello"
    var v5 uint = 65535
    v6 := []int{42, -635}
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v3)
        if (100 <= 104) {
            if (true || (0 == v5)) {
                v5 += (v5 ^ v5)
            } else {
                for fi11 := 0; fi11 < 9; fi11++ {
                    switch (fi10 - fi11) {
                    case 32767:
                        v12 := GenS1{F0: 320}
                        _ = v12
                    }
                    var v13 int = (fi11 - fi11)
                    v2 -= v2
                    _ = fi11
                    _ = v13
                }
                var v14 int8 = -128
                _ = v14
            }
            v5 = (^(v5))
            if len(v6) > 0 {
                v6[0] = (^(fi10))
            }
        } else {
            var v15 uint32 = 1
            v5 = 2
            var v16 int32 = int32(v5)
            _ = v15
            _ = v16
        }
        var v17 float32 = v3
        _ = fi10
        _ = v17
    }
    v3 -= v3
    var v18 float32 = v2
    for fi19 := 0; fi19 < 3; fi19++ {
        _ = len(v6)
        _ = fi19
    }
    if true {
        var v20 float32 = (v3 * v18)
        var v21 float64 = 1.5
        v20 = v20
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 42}
    v4 = v4
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 int32 = -173
            v2 = 0.0
            fmt.Println(v4)
            if len(v6) > 0 {
                v6[0] = -29
            }
            _ = fi24
            _ = v25
        }
        v3 += 0.1
        if (false && (float64(v18) > 0.0)) {
            switch genFn0() {
            case -123:
                switch fi23 {
                case -81:
                    v2 = (v18 - v3)
                case 4294967295:
                    v4 = (("a" + ("\n\t" + ((v4 + v4) + ("a" + v4)))) + "hello")
                case 9223372036854775807:
                    v2 -= (((v3 - v3) * (v2 / 1e10)) / ((1e10 / ((v18 / (v18 * v2)) - v3)) * v18))
                }
            default:
                _ = len(v1)
            }
            if (v18 >= 3.14159) {
                var v26 float64 = -2.25
                v27 := GenS1{F0: -430}
                _ = v26
                _ = v27
            } else {
                v4 = v4
                v4 = "\n\t"
                v28 := GenS1{F0: 256}
                _ = v28
            }
            var v29 uint64 = 1
            _ = v29
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
