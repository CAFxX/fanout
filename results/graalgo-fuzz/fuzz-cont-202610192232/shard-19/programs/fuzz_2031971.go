package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint = 65535
    v2 := []string{"hello", "\n\t", ""}
    v3 := []uint32{779}
    v1 -= (^(v1))
    if len(v3) > 0 {
        v3[0] = 72
    }
    for fi4 := 0; fi4 < 18; fi4++ {
        v1 -= (^(v1))
        if len(v3) > 0 {
            v3[0] = 2
        }
        v5 := GenS1{F0: fi4, F1: (fi4 | (fi4 % ((-872) | 1))), F2: 42}
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint64 = 356
        v9 := GenS1{F0: (fi7 ^ (fi7 << ((v1) & 63))), F1: fi7, F2: fi7}
        switch "hello" {
        case "foo bar":
            if (255 <= 128) {
                v8 = 4294967295
                if (127 <= 127) {
                    var v10 int = (fi7 / ((-9223372036854775808) | 1))
                    var v11 bool = false
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint32 = v6
                    var v13 int16 = 32767
                    _ = v12
                    _ = v13
                }
            } else {
                v6 = 761
            }
        default:
            for fi14 := 0; fi14 < 1; fi14++ {
                v1 += v1
                _ = fi14
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "" {
    case "foo bar":
        var v15 string = ("" + "")
        _ = v15
    case "\n\t":
        _ = len(v3)
    case "":
        if true {
            fmt.Println(v1, v6)
        } else {
            if (!((true && ((!(1 >= -9223372036854775808)) && ((9223372036854775807 <= 1) || (!false)))) && true)) {
                if len(v3) > 0 {
                    v3[0] = (v6 - v6)
                }
                v6 = v6
                v16 := GenS1{F0: -1, F1: -2147483648, F2: -161}
                _ = v16
            }
            var v17 int = 4294967295
            for fi18 := 0; fi18 < 13; fi18++ {
                if len(v3) > 0 {
                    v3[0] = v6
                }
                var v19 uint8 = 58
                if (!false) {
                    v20 := GenS2{F0: uint16(v19)}
                    _ = v20
                } else {
                    var v21 float32 = 0.0
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    }
    v22 := GenS2{F0: 255}
    fmt.Println(v6)
    if true {
        v23 := GenS1{F0: 42, F1: -9223372036854775808, F2: 2147483647}
        _ = v23
    } else {
        v1 -= (v1 ^ 737)
        var v24 int64 = -1
        v25 := GenS1{F0: 42, F1: 1, F2: 127}
        _ = v24
        _ = v25
    }
    var v26 uint64 = 42
    v27 := GenS1{F0: -1, F1: 74, F2: -158}
    v26 += 2
    v26 -= (v26 + 2)
    switch (((("\n\t" + "foo bar") + "\n\t") + "foo bar") + "") {
    case "\n\t":
        v28 := GenS1{F0: -1, F1: -1, F2: 32768}
        _ = v28
    case "a":
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "")
        }
    case "hello":
        var v29 int8 = -128
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v6
}
