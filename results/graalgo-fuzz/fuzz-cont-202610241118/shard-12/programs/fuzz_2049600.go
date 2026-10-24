package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint16 {
    if (((1 < -128) && true) || true) {
        var v7 float32 = float32(1.5)
        var v8 int64 = 9223372036854775807
        var v9 int8 = 100
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{138, 32767, 4294967295}
    var v11 int32 = -1
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 int = 1
    var v2 int64 = -803
    var v3 float64 = 1e10
    v4 := make(map[int]bool)
    v5 := []uint32{2, 42}
    var v6 int64 = v2
    v2 = (v2 + 42)
    for fi12 := 0; fi12 < 2; fi12++ {
        v3 = (v3 + float64(0.1))
        var v13 int64 = (-(v2))
        _ = fi12
        _ = v13
    }
    fmt.Println(v2)
    v3 = (v3 / v3)
    v14 := GenS1{F0: 42}
    if ((^(v2)) <= 9223372036854775807) {
        var v15 int = -2147483648
        for fi16 := 0; fi16 < 20; fi16++ {
            if true {
                var v17 int = 32767
                _ = v17
            } else {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS1{F0: 51}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v6 += -144
                    var v20 int16 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 uint = 192
                    fmt.Println(v15)
                    fmt.Println(v22, v1, v1)
                    v2 += (^(v6))
                    delete(v4, v1)
                    v6 += v6
                    _ = fi21
                    _ = v22
                }
            }
            v23 := v4[2147483647]
            var v24 float64 = v3
            _ = fi16
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(v2, v6, v2)
            var v26 float32 = (((0.0 + 0.0) - 1.5) - -2.25)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi25
            _ = v26
        }
        _ = v15
    }
    var v27 int64 = v6
    switch "hello" {
    case "\n\t":
        var v28 uint8 = 250
        _ = v28
    }
    var v29 float64 = (v3 / v3)
    var v30 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
