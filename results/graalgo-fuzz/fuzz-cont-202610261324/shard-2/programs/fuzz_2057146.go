package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64; F1 string; F2 int }

func main() {
    var v1 int = -569
    var v2 int16 = 1000
    var v3 uint16 = 558
    v4 := make(map[string]uint32)
    switch "hello" {
    case "a":
        v1 -= v1
    case "hello":
        if (!false) {
            for fi5 := 0; fi5 < 2; fi5++ {
                if false {
                    for fi6 := 0; fi6 < 10; fi6++ {
                        var v7 float64 = 3.14159
                        if (((0.1 + ((0.1 / 0.1) * 3.14159)) != 0.0) || true) {
                            var v8 uint64 = 326
                            v1 = fi6
                            _ = v8
                        } else {
                            var v9 float64 = 3.14159
                            var v10 uint = 616
                            _ = v10
                            _ = v9
                        }
                        _ = fi6
                        _ = v7
                    }
                } else {
                    if true {
                        var v11 int16 = (v2 & (v2 | (v2 >> ((784) & 63))))
                        _ = v11
                    }
                    v1 += (-(v1))
                }
                _ = fi5
            }
            switch "foo bar" {
            case "foo bar":
                var v12 float64 = ((1.5 / 1.5) / 1.0)
                _ = v12
            case "":
                for fi13 := 0; fi13 < 8; fi13++ {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 uint = 2
                        v2 -= 0
                        var v16 int64 = 762
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS1{F0: v3, F1: -1, F2: uint8(v3)}
                    _ = fi13
                    _ = v17
                }
            case "a":
                v18 := []int{int(v3)}
                _ = v18
            default:
                v2 = v2
            }
        }
    default:
        switch (^(v1)) {
        case -933:
            var v19 int32 = 42
            _ = v19
        case 980:
            if true {
                var v20 uint32 = 42
                _ = v20
            } else {
                var v21 int16 = -32768
                _ = v21
            }
        case -1:
            var v22 int8 = 98
            _ = v22
        }
    }
    switch 42 {
    case 32767:
        fmt.Println(v1, v2)
    }
    fmt.Println(v3)
    var v23 string = "a"
    if ((!(!true)) && true) {
        var v24 bool = (v23 >= v23)
        var v25 int64 = 1
        for fi26 := 0; fi26 < 2; fi26++ {
            v27 := []int{}
            _ = fi26
            _ = v27
        }
        _ = v24
        _ = v25
    } else {
        switch ((v23 + "") + v23) {
        case "foo bar":
            var v28 uint16 = 158
            _ = v28
        default:
            v3 -= 2
        }
        v1 = 127
    }
    v29 := GenS1{F0: 32767, F1: v1, F2: 184}
    fmt.Println(v2)
    var v30 uint16 = (v3 - (v3 >> ((42) & 63)))
    v30 = v30
    if true {
        v30 += 42
    } else {
        fmt.Println(v30, v23, v30)
    }
    fmt.Println(v30, v1)
    v31 := []int{(v1 + v1)}
    if (v2 <= v2) {
        v1 -= (v1 * v1)
        var v32 uint64 = 42
        _ = v32
    } else {
        v33 := GenS2{F0: 0, F1: (v23 + v23), F2: v1}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
