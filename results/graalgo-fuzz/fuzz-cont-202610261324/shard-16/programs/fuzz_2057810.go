package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 bool; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    v4 := []int{4294967295, -2147483648, -24}
    v5 := GenS1{F0: 71, F1: 2, F2: 106}
    v2 -= -9223372036854775808
    fmt.Println(v2)
    v6 := v1["\n\t"]
    if false {
        v6 = v6
    } else {
        var v7 uint8 = 42
        v2 = v2
        _ = v7
    }
    var v8 bool = (!(((v6 >> ((65535) & 63)) == (v6 * (v6 % ((42) | 1)))) && true))
    v1[""] = 699
    var v9 int8 = 100
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint32 = (v6 & v6)
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 32767}
    v2 += -312
    if v8 {
        if false {
            v4 = append(v4, int(v6))
        }
    }
    var v13 int64 = 0
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            v2 += (v2 & v2)
            v16 := GenS3{F0: "foo bar", F1: (-(v2))}
            switch 32768 {
            case 42:
                delete(v3, ("foo bar" + "a"))
            case 127:
                if (((v9 << ((65535) & 63)) < -106) || v8) {
                    var v17 int64 = v13
                    _ = len(v3)
                    _ = v17
                }
            default:
                v13 += v2
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    switch (("a" + ("\n\t" + ("a" + "foo bar"))) + "a") {
    case "foo bar":
        if v8 {
            var v18 uint = 18446744073709551615
            _ = v18
        } else {
            fmt.Println(v6, v9)
            v8 = v8
            for fi19 := 0; fi19 < 11; fi19++ {
                v13 = v13
                var v20 uint64 = 42
                if v8 {
                    v21 := GenS1{F0: 65535, F1: 255, F2: 127}
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := GenS4{F0: (4294967295 >= v20), F1: uint16(v20), F2: (v13 | 0)}
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                } else {
                    var v24 uint64 = 815
                    var v25 int = -2147483648
                    if v8 {
                        var v26 int16 = -571
                        v2 = -1
                        _ = v26
                    } else {
                        var v27 uint8 = 169
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                }
                _ = fi19
                _ = v20
            }
        }
    case "\n\t":
        if (false || v8) {
            if (!v8) {
                fmt.Println(v13)
                fmt.Println(v8)
                v9 = -128
            }
        } else {
            delete(v3, "hello")
        }
    case "a":
        fmt.Println(v9, v8)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
