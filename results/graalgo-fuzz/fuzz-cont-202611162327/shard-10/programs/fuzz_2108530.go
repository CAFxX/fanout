package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string; F2 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 361
    var v3 int8 = 30
    var v4 uint32 = 2147483648
    v4 -= v4
    var v5 bool = true
    v6 := []int{}
    if len(v6) > 0 {
        v6[0] = -656
    }
    v7 := GenS1{F0: -1}
    switch 555 {
    case -952:
        var v8 uint8 = uint8(v2)
        _ = v8
    case -2147483648:
        fmt.Println(v3, v4)
    case 32768:
        switch "a" {
        case "a":
            var v9 float64 = -2.25
            _ = v9
        default:
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(v4)
                v11 := GenS1{F0: 1}
                _ = fi10
                _ = v11
            }
        }
    }
    v4 -= v4
    v4 = v4
    if (760 < 848) {
        var v12 bool = ((!(v5 && (18446744073709551615 <= (^(v2))))) || v5)
        _ = v12
    } else {
        switch "\n\t" {
        case "\n\t":
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 float32 = float32(float64((-2.25 * (1.5 - 1.5))))
                v15 := GenS1{F0: 1}
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 10; fi16++ {
                v6 = append(v6, fi16)
                _ = fi16
            }
        }
        if true {
            if v5 {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 uint64 = v2
                    _ = fi17
                    _ = v18
                }
                v3 = v3
                var v19 uint = 255
                _ = v19
            }
        } else {
            fmt.Println(v3)
            v3 = -128
            var v20 int16 = -1
            _ = v20
        }
    }
    var v21 int16 = 24
    if ("hello" > "a") {
        switch "hello" {
        case "":
            v22 := GenS1{F0: 42}
            _ = v22
        case "hello":
            var v23 int16 = v21
            _ = v23
        case "\n\t":
            var v24 int = -467
            _ = v24
        }
        v4 = v4
    } else {
        var v25 float64 = 3.14159
        switch 32768 {
        case 128:
            v5 = v5
        case 127:
            v3 -= (^(v3))
        case 12:
            var v26 float64 = v25
            _ = v26
        default:
            v27 := GenS1{F0: 42}
            _ = v27
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 int32 = 2147483647
            if (v3 <= v3) {
                for fi30 := 0; fi30 < 2; fi30++ {
                    switch (fi30 * 9223372036854775807) {
                    case -1:
                        v1[("foo bar" + "a")] = (v4 | (^(v4)))
                    case 128:
                        v25 = -2.25
                    case -867:
                        _ = len(v1)
                    }
                    v31 := v1["\n\t"]
                    _ = fi30
                    _ = v31
                }
            } else {
                var v32 bool = v5
                v33 := GenS1{F0: 42}
                _ = v32
                _ = v33
            }
            _ = fi28
            _ = v29
        }
        _ = v25
    }
    if v5 {
        v34 := GenS2{F0: v2, F1: "hello", F2: (v3 & (v3 ^ (^(v3))))}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 10; fi35++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                v37 := GenS1{F0: 569}
                _ = len(v6)
                fmt.Println(v21, v3, v4)
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        if (2 < 32767) {
            var v38 uint32 = v4
            _ = v38
        }
        var v39 uint8 = 1
        _ = v39
    }
    _ = len(v1)
    for fi40 := 0; fi40 < 14; fi40++ {
        v41 := GenS1{F0: 0}
        var v42 int8 = (v3 * 0)
        switch fi40 {
        case -36:
            var v43 int8 = v3
            _ = v43
        case 42:
            _ = len(v6)
        case 1:
            fmt.Println(v5, v3, v21)
        }
        _ = fi40
        _ = v41
        _ = v42
    }
    if (127 == (v3 & (v3 + 42))) {
        v21 = v21
        v3 = v3
        if (int64(v2) >= 42) {
            for fi44 := 0; fi44 < 13; fi44++ {
                delete(v1, ("foo bar" + "hello"))
                var v45 int8 = int8(v2)
                _ = fi44
                _ = v45
            }
        } else {
            if v5 {
                v46 := GenS2{F0: v2, F1: ("hello" + "hello"), F2: v3}
                v5 = (v5 && ((v2 >> ((255) & 63)) != (v2 / ((430) | 1))))
                _ = v46
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
