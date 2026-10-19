package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint; F1 uint32; F2 int32 }
type GenS3 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"hello", ("foo bar" + "a"), "\n\t", (("foo bar" + "a") + "foo bar")}
    var v3 bool = true
    var v4 string = "a"
    switch 32767 {
    case 128:
        v5 := GenS1{F0: 0, F1: (-1 > 0)}
        _ = v5
    case -157:
        var v6 int8 = 0
        _ = v6
    case -897:
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 bool = v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS2{F0: 1, F1: 4294967295, F2: 42}
    for fi10 := 0; fi10 < 15; fi10++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi10
    }
    v3 = (100 > 0)
    v3 = v3
    switch "" {
    case "a":
        for fi11 := 0; fi11 < 14; fi11++ {
            v3 = v3
            _ = fi11
        }
    default:
        fmt.Println(v3)
    }
    var v12 float64 = 1e10
    v4 = "a"
    var v13 uint16 = 42
    if v3 {
        switch (v4 + "hello") {
        case "a":
            v14 := GenS1{F0: uint64(v13), F1: (646 <= (^(v13)))}
            _ = v14
        case "hello":
            v4 = v4
        }
        switch ((v4 + "") + v4) {
        case "a":
            if (v3 && (!v3)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v15 := v1[4294967295]
                if (!(32767 < 199)) {
                    var v16 uint64 = 0
                    if (0.0 > -2.25) {
                        v3 = (!v15)
                    } else {
                        if (1 < 65535) {
                            v12 += v12
                            if v15 {
                                v17 := GenS2{F0: 2, F1: 610, F2: -2147483648}
                                _ = v17
                            } else {
                                var v18 uint = 4294967295
                                v19 := GenS2{F0: 255, F1: 2, F2: 42}
                                _ = v18
                                _ = v19
                            }
                            v20 := v1[32768]
                            _ = v20
                        }
                    }
                    _ = v16
                }
                _ = v15
            }
        case "":
            if v3 {
                var v21 int32 = -2147483648
                switch v4 {
                case "\n\t":
                    v22 := GenS1{F0: 42, F1: ((-99 > -128) && ((!true) || (255 > 576)))}
                    _ = v22
                default:
                    var v23 uint32 = 4294967295
                    _ = v23
                }
                v2 = append(v2, v4)
                _ = v21
            }
        default:
            v24 := GenS3{F0: 91, F1: 2147483648}
            _ = v24
        }
    } else {
        v12 += v12
        v12 += v12
        v3 = true
    }
    v3 = (v13 <= (v13 << ((216) & 63)))
    for fi25 := 0; fi25 < 17; fi25++ {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            v12 = ((v12 + (float64((1e10 / 1.0)) / ((-2.25 / (float64(3.14159) * (1e10 / v12))) * v12))) + (1e10 - float64(1e10)))
            _ = fi26
        }
        var v27 uint64 = 18446744073709551615
        _ = fi25
        _ = v27
    }
    v28 := v1[-903]
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
