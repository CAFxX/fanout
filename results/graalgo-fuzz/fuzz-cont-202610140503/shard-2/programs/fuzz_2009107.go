package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }

func main() {
    var v1 int = 1
    var v2 uint = 175
    var v3 int32 = -76
    var v4 uint32 = 1
    for fi5 := 0; fi5 < 17; fi5++ {
        if true {
            var v6 int = 32767
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 uint16 = 73
                v2 += v2
                fmt.Println(v2)
                _ = fi7
                _ = v8
            }
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    fmt.Println(v4)
                    var v10 uint8 = uint8(v4)
                    _ = fi9
                    _ = v10
                }
                if false {
                    v11 := GenS1{F0: 2}
                    _ = v11
                }
                fmt.Println(v2, v6, v6)
            }
            _ = v6
        } else {
            v12 := GenS2{F0: 167, F1: uint16(v4), F2: v1}
            _ = v12
        }
        switch int(v4) {
        case 1:
            v4 += (v4 & v4)
        case 988:
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: 0, F1: 1, F2: 42}
                var v15 uint64 = 4294967295
                switch -1 {
                case -304:
                    if (("" >= "\n\t") || (true || (32767 < 999))) {
                        fmt.Println(v1, v2)
                    }
                case 32768:
                    v16 := []int{4294967295}
                    _ = v16
                case 42:
                    fmt.Println(fi5)
                }
                var v17 uint = v2
                _ = fi13
                _ = v14
                _ = v15
                _ = v17
            }
        default:
            var v18 float64 = ((((0.0 * (float64(float32(1.5)) * 1e10)) + 0.1) - 0.1) + (1e10 + (0.0 - -2.25)))
            _ = v18
        }
        _ = fi5
    }
    var v19 int64 = 15
    switch v1 {
    case 42:
        var v20 uint8 = 1
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: uint16(v2)}
        _ = fi21
        _ = v22
    }
    var v23 int32 = (v3 % ((42) | 1))
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 int8 = -37
        for fi26 := 0; fi26 < 5; fi26++ {
            fmt.Println(v19, v23)
            v27 := []int{}
            switch fi24 {
            case -2147483648:
                v4 = 0
            default:
                if (true || (((("hello" + "foo bar") + "foo bar") != "a") || false)) {
                    var v28 string = ("a" + ("a" + "\n\t"))
                    v1 = -2147483648
                    _ = v28
                }
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
                var v29 uint64 = 4294967295
                v29 += v29
                _ = v29
            }
            _ = fi26
            _ = v27
        }
        v3 = 1
        _ = fi24
        _ = v25
    }
    v30 := GenS2{F0: 255, F1: 42, F2: 42}
    for fi31 := 0; fi31 < 3; fi31++ {
        switch (fi31 << ((v2) & 63)) {
        case 9223372036854775807:
            v4 -= 94
        default:
            var v32 bool = (!((((-2.25 * 3.14159) == 0.1) || false) || true))
            _ = v32
        }
        var v33 int8 = 23
        _ = fi31
        _ = v33
    }
    var v34 int64 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v34
    _ = v4
}
