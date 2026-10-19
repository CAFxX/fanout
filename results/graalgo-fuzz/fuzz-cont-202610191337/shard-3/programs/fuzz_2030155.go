package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    var v7 uint64 = 1
    if (32767 != 991) {
        v8 := GenS1{F0: 86}
        var v9 int8 = 115
        _ = v8
        _ = v9
    } else {
        v10 := []int{32767}
        fmt.Println(v7)
        if len(v10) > 0 {
            v10[0] = 128
        }
        _ = v10
    }
    _ = v7
    return "foo bar"
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{636, 146, 2, 2147483648}
    var v3 int16 = 32767
    var v4 int8 = 1
    v5 := make(map[string]string)
    v6 := make(map[int]int)
    v11 := v5[""]
    if false {
        var v12 uint16 = 42
        if true {
            var v13 float64 = float64(((v1 / v1) + (v1 * 0.0)))
            _ = v13
        }
        var v14 bool = false
        _ = v12
        _ = v14
    } else {
        switch 0 {
        case -348:
            if true {
                v15 := GenS1{F0: -1}
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint32 = 663
                    _ = fi16
                    _ = v17
                }
                var v18 uint = 65535
                _ = v15
                _ = v18
            } else {
                if ((!(((0.1 - 0.0) - (1.5 - 1.5)) == (0.1 / 1.0))) && false) {
                    fmt.Println(v1, v3, v4)
                    var v19 int16 = v3
                    _ = v19
                } else {
                    fmt.Println(v11)
                }
                if ((0.0 * v1) < (0.1 * v1)) {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        for fi21 := 0; fi21 < 5; fi21++ {
                            var v22 uint16 = 732
                            v23 := GenS1{F0: -1}
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        _ = fi20
                    }
                }
            }
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            v4 -= v4
            _ = fi24
        }
        v4 += (-(v4))
    }
    var v25 float64 = 1e10
    v26 := GenS1{F0: 183}
    for fi27 := 0; fi27 < 7; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            v4 = v4
            _ = fi28
        }
        _ = fi27
    }
    var v29 uint8 = 108
    _ = len(v2)
    for fi30 := 0; fi30 < 20; fi30++ {
        v6[128] = -375
        v29 += v29
        _ = fi30
    }
    switch v11 {
    case "\n\t":
        fmt.Println(v29)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
