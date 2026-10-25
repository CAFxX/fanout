package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }

func main() {
    var v1 float32 = (0.1 - 3.14159)
    var v2 string = ("hello" + ("hello" + "\n\t"))
    v3 := make(map[uint32]uint32)
    var v4 uint = 2
    v5 := []uint32{2, 2, 1, 65535}
    v4 = v4
    var v6 uint = v4
    v2 = v2
    v4 = 521
    v7 := GenS1{F0: 121}
    v8 := GenS1{F0: 127}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint = (v6 * (v6 & 65535))
        var v11 uint16 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    switch (("hello" + ("hello" + v2)) + ((v2 + v2) + ("\n\t" + v2))) {
    case "":
        v6 = 595
    default:
        if true {
            for fi12 := 0; fi12 < 15; fi12++ {
                v4 = (v4 << ((v4) & 63))
                fmt.Println(v2, v4, v6)
                _ = fi12
            }
        } else {
            v13 := GenS1{F0: 0}
            var v14 float64 = 1e10
            v15 := v3[42]
            _ = v13
            _ = v14
            _ = v15
        }
    }
    var v16 float64 = 1.5
    v17 := GenS1{F0: 0}
    v3[978] = 0
    if (v1 < v1) {
        fmt.Println(v6, v4, v2)
    }
    switch -2147483648 {
    case 127:
        var v18 uint = 65535
        _ = v18
    case -2147483648:
        v16 -= (0.1 + v16)
    case -9223372036854775808:
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int = 9223372036854775807
            _ = fi19
            _ = v20
        }
    default:
        var v21 uint64 = 2
        _ = v21
    }
    if (0 <= 4) {
        fmt.Println(v1, v4, v6)
        if (!true) {
            if (true || (!((18446744073709551615 < 18446744073709551615) || (!true)))) {
                var v22 float32 = ((3.14159 - v1) + v1)
                var v23 uint16 = 0
                var v24 string = v2
                _ = v22
                _ = v23
                _ = v24
            } else {
                var v25 string = v2
                for fi26 := 0; fi26 < 1; fi26++ {
                    if true {
                        var v27 uint8 = 0
                        for fi28 := 0; fi28 < 1; fi28++ {
                            v2 = v25
                            var v29 float64 = ((v16 + float64(1.5)) - v16)
                            v4 = v4
                            _ = fi28
                            _ = v29
                        }
                        v30 := GenS2{F0: 65535, F1: 65535, F2: 0}
                        _ = v27
                        _ = v30
                    }
                    _ = fi26
                }
                _ = v25
            }
        } else {
            v31 := GenS2{F0: 754, F1: 1, F2: -1}
            var v32 int16 = 32767
            _ = v31
            _ = v32
        }
    } else {
        var v33 uint32 = 65535
        v34 := GenS2{F0: 255, F1: v33, F2: 42}
        for fi35 := 0; fi35 < 2; fi35++ {
            delete(v3, v33)
            _ = fi35
        }
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
