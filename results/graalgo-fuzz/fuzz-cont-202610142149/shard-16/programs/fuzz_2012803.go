package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }

func main() {
    v1 := []int{42, 875, 32767}
    var v2 int8 = 11
    v3 := []uint32{42, 2147483647}
    v4 := []string{}
    var v5 int16 = -32768
    v6 := make(map[int]bool)
    if (2147483647 <= -217) {
        switch "a" {
        case "a":
            v7 := v6[-1]
            _ = v7
        default:
            v5 = v5
        }
    } else {
        fmt.Println(v2, v2)
    }
    v3 = append(v3, 65535)
    fmt.Println(v2)
    switch ("hello" + "\n\t") {
    case "hello":
        v4 = append(v4, "\n\t")
    default:
        v8 := GenS1{F0: 42, F1: 2, F2: 42}
        _ = v8
    }
    var v9 int = -2147483648
    v9 -= -9223372036854775808
    switch "foo bar" {
    case "\n\t":
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS1{F0: -24, F1: 0, F2: 42}
                v5 = (v5 * (v5 % ((v5) | 1)))
                _ = fi11
                _ = v12
            }
            v5 += v5
            _ = fi10
        }
    default:
        var v13 int8 = v2
        _ = v13
    }
    var v14 int16 = (v5 >> ((0) & 63))
    fmt.Println(v2, v5)
    fmt.Println(v14)
    for fi15 := 0; fi15 < 18; fi15++ {
        switch (v9 & (fi15 - 32767)) {
        case 9223372036854775807:
            v16 := GenS1{F0: (v9 - (v9 % ((v9) | 1))), F1: 255, F2: 32767}
            _ = v16
        }
        fmt.Println(v2)
        _ = fi15
    }
    var v17 uint64 = 587
    fmt.Println(v17)
    switch -9223372036854775808 {
    case 128:
        fmt.Println(v14, v14)
    case 1:
        fmt.Println(v2, v17)
    case -77:
        for fi18 := 0; fi18 < 19; fi18++ {
            v9 = (v9 % ((fi18) | 1))
            for fi19 := 0; fi19 < 2; fi19++ {
                v9 = 1
                _ = fi19
            }
            _ = fi18
        }
    default:
        v20 := v6[(v9 >> ((4294967295) & 63))]
        _ = v20
    }
    v17 += v17
    for fi21 := 0; fi21 < 17; fi21++ {
        if (2147483648 < 160) {
            var v22 uint32 = 0
            if false {
                for fi23 := 0; fi23 < 8; fi23++ {
                    fmt.Println(v17, v2, v17)
                    switch fi23 {
                    case -1:
                        var v24 int8 = 1
                        _ = v24
                    case 512:
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                    }
                    _ = fi23
                }
                switch (fi21 & -2147483648) {
                case -782:
                    if false {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v5 = 1000
                            _ = fi25
                        }
                        for fi26 := 0; fi26 < 5; fi26++ {
                            _ = len(v6)
                            v27 := GenS1{F0: -117, F1: 630, F2: 2}
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        var v28 float32 = 3.14159
                        _ = v28
                    }
                case 737:
                    if ((false && (v14 != (v14 | int16(v22)))) && false) {
                        if false {
                            var v29 string = "hello"
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                            _ = v29
                        } else {
                            v2 = (v2 >> ((1) & 63))
                            v17 = (v17 + (v17 + (v17 + v17)))
                        }
                    }
                case -9223372036854775808:
                    var v30 string = (((("" + ("a" + ("a" + ""))) + "") + "") + ("\n\t" + (("hello" + "foo bar") + "")))
                    _ = v30
                }
                v5 += (v14 >> ((39) & 63))
            }
            fmt.Println(v14)
            _ = v22
        }
        var v31 uint32 = 2147483648
        v32 := GenS1{F0: 2147483647, F1: 269, F2: 0}
        _ = fi21
        _ = v31
        _ = v32
    }
    v14 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
