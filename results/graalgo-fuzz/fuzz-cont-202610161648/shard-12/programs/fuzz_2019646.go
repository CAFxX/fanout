package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = v1
    var v3 uint64 = 4294967295
    var v4 int32 = 0
    v2 -= v1
    if true {
        v4 = v4
    } else {
        v2 = (v1 / (((v1 % ((v1) | 1))) | 1))
        switch (("hello" + "foo bar") + "\n\t") {
        case "a":
            var v5 bool = true
            _ = v5
        default:
            if true {
                for fi6 := 0; fi6 < 17; fi6++ {
                    if true {
                        var v7 float32 = (((float32(0.0) / 1.0) - (0.0 + float32(0.0))) + 1e10)
                        _ = v7
                    }
                    fmt.Println(v3, v3, v1)
                    fmt.Println(fi6, v4)
                    _ = fi6
                }
                for fi8 := 0; fi8 < 16; fi8++ {
                    fmt.Println(v4)
                    var v9 float32 = float32(float64((((3.14159 + float32((-2.25 * 1e10))) + (float32(1.5) + float32(1.5))) / 1e10)))
                    _ = fi8
                    _ = v9
                }
                v10 := []int{-1, -1, -973, 960}
                _ = v10
            } else {
                v11 := GenS1{F0: (v1 >> ((65535) & 63))}
                _ = v11
            }
        }
    }
    switch 2147483647 {
    case 746:
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 string = "a"
            _ = fi12
            _ = v13
        }
    }
    if (!(false && ((226 != 255) && false))) {
        var v14 uint8 = 255
        if false {
            var v15 bool = (34 != 1)
            v15 = ((v14 >> ((42) & 63)) < 170)
            _ = v15
        }
        var v16 int = 32768
        _ = v14
        _ = v16
    }
    fmt.Println(v4, v1)
    var v17 int64 = 0
    var v18 int = 42
    v19 := []int{v18, (v18 - (^(v18))), v18}
    fmt.Println(v3, v3, v2)
    for fi20 := 0; fi20 < 10; fi20++ {
        v1 -= 2147483647
        var v21 uint8 = 127
        if (4294967295 < 18446744073709551615) {
            if len(v19) > 0 {
                v19[0] = v18
            }
            v22 := GenS1{F0: -582}
            var v23 uint = 2
            _ = v22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if (127 <= -128) {
        var v24 int = v18
        var v25 float32 = (0.1 * 1.5)
        var v26 uint = 18446744073709551615
        _ = v24
        _ = v25
        _ = v26
    } else {
        fmt.Println(v17)
        var v27 uint32 = 0
        v28 := GenS1{F0: (v4 * v2)}
        _ = v27
        _ = v28
    }
    if ((^(v1)) < (v2 >> ((255) & 63))) {
        var v29 uint16 = 527
        v30 := GenS1{F0: v1}
        _ = v29
        _ = v30
    } else {
        var v31 uint64 = (v3 & (v3 + (v3 % ((v3) | 1))))
        for fi32 := 0; fi32 < 1; fi32++ {
            if (v17 >= 42) {
                v33 := GenS1{F0: v4}
                if (false && (2 == 4294967295)) {
                    for fi34 := 0; fi34 < 5; fi34++ {
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        fmt.Println(v2)
                        _ = fi34
                    }
                    fmt.Println(fi32)
                } else {
                    fmt.Println(v2)
                    for fi35 := 0; fi35 < 10; fi35++ {
                        switch 127 {
                        case 127:
                            var v36 float64 = 0.1
                            _ = v36
                        default:
                            v1 = (v4 / (((v1 >> ((931) & 63))) | 1))
                        }
                        _ = fi35
                    }
                    v37 := GenS1{F0: -103}
                    _ = v37
                }
                v19 = append(v19, fi32)
                _ = v33
            }
            v2 = v2
            v1 = (v1 ^ (-(v2)))
            _ = fi32
        }
        _ = v31
    }
    v38 := GenS1{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
