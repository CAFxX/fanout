package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }

func main() {
    v1 := []uint32{369, 857, 2147483647, 4294967295}
    v2 := []uint32{2147483647, 2147483647, 2147483648, 236}
    v3 := make(map[uint32]int)
    var v4 int16 = 0
    v5 := []int{-402, 0}
    var v6 bool = false
    v7 := GenS1{F0: 65535}
    var v8 int32 = -171
    fmt.Println(v8, v4)
    v4 += 0
    switch "hello" {
    case "hello":
        var v9 int = 405
        _ = v9
    case "foo bar":
        var v10 uint64 = 600
        _ = v10
    default:
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS1{F0: 334}
            _ = fi11
            _ = v12
        }
    }
    switch -259 {
    case 42:
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: 1}
            _ = fi13
            _ = v14
        }
    case 32768:
        if (v6 && v6) {
            v4 += 42
            fmt.Println(v4, v6, v6)
        } else {
            v15 := GenS1{F0: 1}
            _ = v15
        }
    case -2147483648:
        if ((!(128 > 2)) && (1 < 301)) {
            v4 += v4
        } else {
            v6 = v6
            switch ("\n\t" + "\n\t") {
            case "a":
                var v16 uint = 76
                _ = v16
            case "\n\t":
                var v17 uint64 = 2
                _ = v17
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v18 uint = 18446744073709551615
                _ = v18
            }
        }
    default:
        if true {
            var v19 int32 = v8
            var v20 int64 = -51
            _ = v19
            _ = v20
        } else {
            if (v6 && (v6 || (!v6))) {
                v8 -= 2147483647
                var v21 string = ""
                fmt.Println(v4, v6)
                _ = v21
            }
        }
    }
    var v22 uint16 = 2
    fmt.Println(v6)
    var v23 int8 = 1
    var v24 float32 = (1.5 - 0.1)
    v24 = 0.1
    if (42 > int64(v8)) {
        fmt.Println(v4, v8)
    }
    if (!(v6 && v6)) {
        switch (((("" + ("" + "a")) + "") + "") + "a") {
        case "a":
            fmt.Println(v6)
        default:
            if true {
                v25 := GenS2{F0: -169}
                if v6 {
                    for fi26 := 0; fi26 < 14; fi26++ {
                        if (!(2 == 228)) {
                            for fi27 := 0; fi27 < 3; fi27++ {
                                v23 += v23
                                _ = fi27
                            }
                            fmt.Println(v23)
                        }
                        if (672 > 18446744073709551615) {
                            v28 := v3[2147483648]
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                            _ = v28
                        }
                        _ = fi26
                    }
                    v24 -= -2.25
                } else {
                    var v29 int8 = (v23 | 86)
                    _ = v29
                }
                var v30 uint64 = 0
                _ = v25
                _ = v30
            }
        }
    }
    var v31 int8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
