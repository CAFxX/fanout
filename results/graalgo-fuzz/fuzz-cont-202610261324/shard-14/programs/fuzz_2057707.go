package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 bool, p1 uint32, p2 uint8) uint64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint = 42
        _ = fi7
        _ = v8
    }
    p0 = p0
    p2 += (^(p2))
    var v9 int64 = 9223372036854775807
    v10 := GenS1{F0: ((!("" < "a")) && p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1
}
func genFn1(p0 int64, p1 int16) int64 {
    p0 -= p0
    var v11 uint8 = 127
    var v12 int64 = 42
    var v13 uint8 = v11
    if true {
        v12 += p0
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return v12
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = -35
    var v3 int16 = -1
    v4 := make(map[string]bool)
    v5 := []string{"a", ""}
    var v6 uint = 1
    var v14 uint8 = 1
    var v15 uint8 = (v14 - 175)
    switch (v1 + ("foo bar" + v1)) {
    case "a":
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS2{F0: 4294967295, F1: 100}
            v18 := GenS1{F0: true}
            _ = fi16
            _ = v17
            _ = v18
        }
    case "\n\t":
        _ = len(v5)
    case "hello":
        v19 := GenS1{F0: (0.1 != 3.14159)}
        _ = v19
    }
    v4[(v1 + ("a" + v1))] = (65535 >= 0)
    for fi20 := 0; fi20 < 9; fi20++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v21 := GenS1{F0: false}
        _ = fi20
        _ = v21
    }
    v15 += (v15 & v15)
    if len(v5) > 0 {
        v5[0] = v1
    }
    var v22 bool = true
    switch v1 {
    case "hello":
        v15 = (^(v15))
    case "":
        v6 -= v6
    case "a":
        for fi23 := 0; fi23 < 6; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                if len(v5) > 0 {
                    v5[0] = (v1 + v1)
                }
                var v25 int16 = (v3 ^ int16(v6))
                v26 := GenS1{F0: ((v22 || true) && (65535 < 65535))}
                _ = fi24
                _ = v25
                _ = v26
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi23
        }
    }
    switch -233 {
    case 32768:
        v6 += (^(v6))
    case 128:
        if v22 {
            v27 := v4[v1]
            for fi28 := 0; fi28 < 9; fi28++ {
                v2 = (v2 | v2)
                v29 := GenS2{F0: genFn0(true, 42, v15), F1: v2}
                if true {
                    var v30 uint = (v6 << (((v6 / (((^(v6))) | 1))) & 63))
                    var v31 uint64 = 0
                    if false {
                        v32 := v4[("" + "")]
                        _ = v32
                    }
                    _ = v30
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
            v3 += v3
            _ = v27
        }
    default:
        v33 := v4[(v1 + (v1 + v1))]
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
