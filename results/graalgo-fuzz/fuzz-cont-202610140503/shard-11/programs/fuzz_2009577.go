package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 128
    v3 := make(map[int]int)
    var v4 int8 = -128
    v5 := []uint32{896, 2}
    v6 := make(map[int]string)
    v7 := GenS1{F0: 2}
    v5 = append(v5, 0)
    if true {
        if true {
            fmt.Println(v4, v2)
        }
    }
    var v8 int16 = 42
    v9 := GenS1{F0: 116}
    v8 = v8
    var v10 float32 = (3.14159 - 1.5)
    for fi11 := 0; fi11 < 15; fi11++ {
        v4 = (-(v4))
        if false {
            if (0 != -1) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    fmt.Println(v8, v8)
                    var v13 uint64 = 486
                    var v14 string = "hello"
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 int16 = v8
                var v16 int = (fi11 ^ (-(fi11)))
                _ = v15
                _ = v16
            }
            var v17 uint8 = 2
            _ = v17
        } else {
            if (!((v8 & 1000) <= v8)) {
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := GenS1{F0: 0}
                    _ = fi18
                    _ = v19
                }
                v10 = (v10 * 0.1)
            } else {
                var v20 int64 = -871
                _ = v20
            }
        }
        var v21 int8 = -1
        _ = fi11
        _ = v21
    }
    var v22 int32 = int32(v2)
    v22 -= v22
    var v23 uint32 = 587
    v24 := v3[-1]
    if false {
        v24 = -1
        switch "foo bar" {
        case "":
            if ("" != "foo bar") {
                if false {
                    var v25 float64 = float64(v10)
                    _ = v25
                } else {
                    switch (("hello" + ("foo bar" + "hello")) + "hello") {
                    case "":
                        if true {
                            v8 -= v8
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if (1 <= v4) {
                                v8 = v8
                            }
                        }
                    }
                    for fi26 := 0; fi26 < 12; fi26++ {
                        fmt.Println(v22, v22)
                        _ = fi26
                    }
                }
            } else {
                v24 += -438
                var v27 float32 = v10
                _ = v27
            }
        }
    }
    var v28 int = (v24 - v24)
    for fi29 := 0; fi29 < 8; fi29++ {
        if (!(v10 == v10)) {
            v28 -= -427
            fmt.Println(v23, fi29)
        } else {
            for fi30 := 0; fi30 < 5; fi30++ {
                v2 = (v2 ^ v2)
                v31 := GenS1{F0: 42}
                if true {
                    if (int8(v23) < v4) {
                        v28 -= (fi29 << ((1) & 63))
                        var v32 int8 = (v4 ^ v4)
                        _ = v32
                    } else {
                        if len(v5) > 0 {
                            v5[0] = (v23 + 0)
                        }
                        v22 += v22
                        var v33 uint64 = 720
                        _ = v33
                    }
                    v22 += (v22 - v22)
                } else {
                    _ = len(v5)
                }
                _ = fi30
                _ = v31
            }
            v34 := GenS1{F0: 147}
            _ = len(v6)
            _ = v34
        }
        fmt.Println(v22, v28)
        _ = fi29
    }
    fmt.Println(v2, v22)
    v22 -= v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
