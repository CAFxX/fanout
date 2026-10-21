package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{2147483647, 517}
    var v2 int32 = -2147483648
    var v3 int16 = 1
    v4 := []uint32{2147483647}
    var v5 int32 = (^(v2))
    v4 = append(v4, 65535)
    if (!true) {
        var v6 float32 = 1.5
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int32 = 1
            _ = fi7
            _ = v8
        }
        var v9 int8 = 127
        _ = v6
        _ = v9
    }
    v2 += (v2 - v5)
    if (!(!false)) {
        var v10 int64 = 678
        _ = v10
    } else {
        v3 = 860
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: -1}
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: -128}
        _ = v13
    }
    v5 = -2147483648
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint16 = 65535
        var v16 int64 = -9223372036854775808
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v2, v3)
    v4 = append(v4, 1)
    fmt.Println(v2, v2, v5)
    if (100 > 100) {
        v5 = -2147483648
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        switch (("" + "hello") + ("a" + "hello")) {
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            v5 = v2
        case "a":
            switch 32768 {
            case -2147483648:
                v18 := GenS1{F0: 100}
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 3; fi19++ {
                switch (fi17 | -92) {
                case 127:
                    var v20 uint8 = 107
                    _ = v20
                case 2147483647:
                    v1 = append(v1, 4294967295)
                    v21 := GenS1{F0: -1}
                    v2 = v5
                    _ = v21
                case 32768:
                    if ((fi19 >> ((4294967295) & 63)) > fi19) {
                        v5 = v2
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    }
                default:
                    v23 := GenS1{F0: 0}
                    v5 = (v5 % ((v5) | 1))
                    fmt.Println(fi19)
                    _ = len(v4)
                    if len(v1) > 0 {
                        v1[0] = 710
                    }
                    v2 += 0
                    _ = v23
                }
                _ = fi19
            }
        }
        _ = fi17
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS1{F0: 100}
        v3 = (v3 * v3)
        _ = fi24
        _ = v25
    }
    var v26 uint32 = 42
    switch 1 {
    case 128:
        fmt.Println(v26, v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
