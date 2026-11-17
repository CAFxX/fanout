package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float64 = (0.1 * float64(float32(1.5)))
    var v2 int32 = -814
    v3 := make(map[string]uint32)
    v4 := []uint32{1, 534, 2147483647, 4294967295}
    v1 -= float64((0.1 - -2.25))
    delete(v3, "foo bar")
    v5 := GenS1{F0: "hello", F1: 2147483647, F2: 18446744073709551615}
    fmt.Println(v2, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 bool = (!("" > "hello"))
    if (1 != 442) {
        fmt.Println(v6, v1, v2)
        v1 = v1
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi8 := 0; fi8 < 17; fi8++ {
                if v6 {
                    var v9 uint8 = 2
                    var v10 string = "hello"
                    var v11 int16 = -686
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    v12 := GenS1{F0: "\n\t", F1: 2147483647, F2: 65535}
                    if v6 {
                        fmt.Println(v1, fi8)
                        switch 227 {
                        case 2147483647:
                            var v13 uint8 = 2
                            _ = v13
                        case 9223372036854775807:
                            var v14 int16 = 1
                            _ = v14
                        case 0:
                            v6 = v6
                        default:
                            v15 := GenS1{F0: "foo bar", F1: 42, F2: 2}
                            _ = v15
                        }
                        v1 -= 0.1
                    } else {
                        v16 := v3["\n\t"]
                        v6 = (true || v6)
                        _ = v16
                    }
                    _ = v12
                }
                _ = fi8
            }
            _ = fi7
        }
        if (!(0 != 0)) {
            if false {
                var v17 string = ("foo bar" + "")
                _ = v17
            }
        }
    }
    var v18 uint8 = 2
    v2 = (v2 << ((916) & 63))
    if false {
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS2{F0: 774}
            var v21 float64 = (v1 * float64(0.0))
            delete(v3, "hello")
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        if v6 {
            var v22 float64 = 0.0
            fmt.Println(v22, v2, v6)
            _ = v22
        } else {
            var v23 uint32 = 380
            v6 = false
            _ = v23
        }
        switch "foo bar" {
        case "":
            delete(v3, ("foo bar" + "hello"))
        case "hello":
            fmt.Println(v6)
        }
    }
    var v24 int16 = -282
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
