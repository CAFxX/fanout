package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }
func genFn0(p0 int32) string {
    v7 := []int{32768, 42}
    v8 := GenS1{F0: 32767, F1: 65535, F2: p0}
    var v9 float64 = (3.14159 / -2.25)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 uint16 = 362
    v2 := []int{0, -9223372036854775808}
    var v3 int = 32768
    v4 := []uint32{}
    v5 := make(map[uint32]bool)
    v6 := make(map[string]uint32)
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            _ = len(v5)
            var v12 bool = true
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v3 = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS2{F0: 100}
        var v15 string = "foo bar"
        _ = fi13
        _ = v14
        _ = v15
    }
    if (1 <= 42) {
        switch "\n\t" {
        case "\n\t":
            if true {
                switch (("" + "foo bar") + "hello") {
                case "hello":
                    var v16 bool = (42 == -159)
                    _ = v16
                }
                if false {
                    if (("" >= ("" + ("" + "a"))) || (214 >= 85)) {
                        v17 := GenS1{F0: (^(v1)), F1: 531, F2: 2147483647}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v3, v3)
                        _ = v17
                    } else {
                        if true {
                            fmt.Println(v3)
                        }
                    }
                } else {
                    var v18 int16 = -32768
                    var v19 bool = true
                    _ = v18
                    _ = v19
                }
            }
        case "hello":
            switch "" {
            case "\n\t":
                if (true || (false || false)) {
                    var v20 string = ""
                    switch v20 {
                    case "foo bar":
                        v3 = (v3 - 32767)
                    }
                    v21 := GenS3{F0: 1, F1: 965, F2: 0}
                    _ = v20
                    _ = v21
                } else {
                    var v22 int16 = 42
                    v22 = 541
                    _ = v22
                }
            case "a":
                v23 := GenS4{F0: 18446744073709551615, F1: 2, F2: 42}
                _ = v23
            default:
                var v24 int32 = -859
                _ = v24
            }
        default:
            var v25 bool = true
            _ = v25
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            _ = len(v2)
            var v27 int16 = 0
            for fi28 := 0; fi28 < 14; fi28++ {
                v27 = v27
                v29 := v5[2147483647]
                _ = fi28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
    } else {
        v3 -= v3
    }
    var v30 uint8 = 245
    v31 := GenS4{F0: 18446744073709551615, F1: v30, F2: 4294967295}
    var v32 bool = (-54 > -128)
    switch 9223372036854775807 {
    case -303:
        v32 = true
    default:
        switch (v3 << ((255) & 63)) {
        case 9223372036854775807:
            for fi33 := 0; fi33 < 3; fi33++ {
                v34 := v5[2]
                _ = fi33
                _ = v34
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
