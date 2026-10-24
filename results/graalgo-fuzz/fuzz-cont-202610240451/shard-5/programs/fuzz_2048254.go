package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int32 }

func main() {
    var v1 string = "a"
    v2 := []string{(v1 + ("" + ("a" + (v1 + (v1 + v1)))))}
    v3 := []uint32{550, 434, 728}
    v4 := []string{"a", v1, "a", ""}
    if (0 <= -32768) {
        v1 = (v1 + v1)
    }
    var v5 int = -432
    var v6 float32 = (1e10 + 0.1)
    for fi7 := 0; fi7 < 1; fi7++ {
        v5 = (v5 - fi7)
        switch 1 {
        case -9223372036854775808:
            if true {
                v8 := GenS1{F0: 32767, F1: 127, F2: 2147483647}
                for fi9 := 0; fi9 < 5; fi9++ {
                    v10 := GenS1{F0: -32768, F1: 224, F2: -2147483648}
                    _ = fi9
                    _ = v10
                }
                if true {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v12 := GenS1{F0: -32768, F1: 0, F2: 1}
                        _ = fi11
                        _ = v12
                    }
                    v1 = (v1 + "a")
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 int32 = 1
                        _ = fi13
                        _ = v14
                    }
                }
                _ = v8
            }
        case 419:
            fmt.Println(fi7)
        case 4294967295:
            v15 := GenS1{F0: 32767, F1: 42, F2: -93}
            _ = v15
        default:
            if (235 >= 18446744073709551615) {
                fmt.Println(v1)
                var v16 int64 = 1
                _ = v16
            } else {
                switch fi7 {
                case -678:
                    fmt.Println(v5, v6)
                case -2147483648:
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 = append(v2, "hello")
                        _ = fi17
                    }
                case 450:
                    v1 = v1
                }
                v1 = v1
            }
        }
        _ = fi7
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        if (!true) {
            fmt.Println(v5, v1, fi18)
        } else {
            var v19 string = (v1 + v1)
            _ = v19
        }
        _ = fi18
    }
    if len(v3) > 0 {
        v3[0] = 540
    }
    v6 += v6
    fmt.Println(v1, v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
