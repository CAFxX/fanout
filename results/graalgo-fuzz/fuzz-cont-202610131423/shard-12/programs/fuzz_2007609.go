package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 int32 = 166
    var v4 uint = 65535
    var v5 uint16 = 32767
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int16 = 0
            var v10 uint16 = v5
            _ = fi8
            _ = v10
            _ = v9
        }
        if (((0.0 - float64(((v6 * v6) / (float32(0.1) * v6)))) + 0.0) <= 3.14159) {
            switch (fi7 / (((-(fi7))) | 1)) {
            case 128:
                var v11 string = (("" + ("\n\t" + "a")) + "foo bar")
                _ = v11
            case -711:
                v12 := v1[2]
                _ = v12
            case 32767:
                for fi13 := 0; fi13 < 14; fi13++ {
                    v3 = (v3 * (v3 % (((v3 - (v3 / ((v3) | 1)))) | 1)))
                    v1[318] = 2147483648
                    switch (fi13 << (((v4 & 707)) & 63)) {
                    case 42:
                        if (v5 != uint16(v4)) {
                            if (1 <= 594) {
                                switch fi13 {
                                case 0:
                                    delete(v2, "foo bar")
                                case -9223372036854775808:
                                    v14 := []int{(fi7 - (fi7 / ((fi7) | 1))), fi7}
                                    _ = v14
                                case 113:
                                    var v15 int16 = -945
                                    _ = v15
                                default:
                                    v16 := GenS1{F0: false}
                                    _ = v16
                                }
                                v17 := v1[42]
                                if ((("\n\t" + ("hello" + ("hello" + "hello"))) <= "hello") || true) {
                                    var v18 float32 = (v6 * float32(float64(float32(0.1))))
                                    if (!true) {
                                        fmt.Println(v4)
                                        switch "a" {
                                        case "":
                                            v19 := v2[("hello" + "hello")]
                                            _ = v19
                                        default:
                                            var v20 int64 = 776
                                            _ = v20
                                        }
                                        var v21 uint8 = 255
                                        _ = v21
                                    }
                                    _ = v18
                                }
                                _ = v17
                            }
                            var v22 uint64 = 4294967295
                            _ = v22
                        }
                    default:
                        v23 := GenS1{F0: true}
                        var v24 int64 = -9223372036854775808
                        v3 -= v3
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v5, fi13)
                    _ = fi13
                }
            default:
                delete(v2, "hello")
            }
        } else {
            v4 += 587
        }
        v3 = -2147483648
        _ = fi7
    }
    delete(v2, "hello")
    switch "a" {
    case "hello":
        v5 = (v5 - uint16(v4))
    case "\n\t":
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := GenS1{F0: true}
            var v27 int = (^(fi25))
            _ = fi25
            _ = v26
            _ = v27
        }
    case "foo bar":
        var v28 uint8 = 223
        _ = v28
    }
    v29 := v2["a"]
    var v30 int = 32767
    var v31 int = v30
    var v32 float64 = (0.1 / ((((3.14159 + 0.1) - 0.1) + float64(((1.5 + (v6 + (v6 / ((1e10 + v6) / v6)))) + 0.0))) * 3.14159))
    _ = len(v1)
    switch "\n\t" {
    case "hello":
        v33 := []int{(v31 << (((v4 | v4)) & 63))}
        _ = v33
    case "foo bar":
        v34 := v1[2147483648]
        _ = v34
    case "\n\t":
        fmt.Println(v5)
    }
    switch ("\n\t" + "hello") {
    case "\n\t":
        v35 := v2[((("hello" + ("hello" + "\n\t")) + ("foo bar" + "a")) + "\n\t")]
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
