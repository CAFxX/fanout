package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 1
    var v3 int32 = -2147483648
    var v4 float32 = (0.1 - -2.25)
    v5 := []int{32768}
    var v6 uint64 = 944
    v7 := GenS1{F0: true, F1: v2, F2: true}
    v8 := GenS1{F0: ((("\n\t" + ("" + "foo bar")) != "foo bar") && false), F1: (v2 & v2), F2: ((-2.25 - (0.1 / 1.0)) < v4)}
    var v9 string = ("\n\t" + "hello")
    var v10 int32 = 1
    v11 := GenS2{F0: 18, F1: v3}
    v9 = ""
    var v12 uint16 = 583
    v4 = v4
    var v13 int64 = -714
    fmt.Println(v10, v4, v12)
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            switch fi15 {
            case 128:
                switch fi15 {
                case -1:
                    var v16 int = (fi14 << ((18446744073709551615) & 63))
                    _ = v16
                default:
                    fmt.Println(v2)
                }
            case 4294967295:
                if (((true || true) && (!(false || (int32(v6) >= v3)))) || (((820 > 32767) || false) || (1.5 <= (float32(-2.25) + v4)))) {
                    if true {
                        var v17 float32 = v4
                        v18 := GenS2{F0: 225, F1: -883}
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(v13, v12, fi15)
                } else {
                    switch int(v12) {
                    case 2147483647:
                        v19 := v1[255]
                        fmt.Println(v19)
                        _ = v19
                    case 414:
                        if true {
                            v10 = v3
                        } else {
                            v13 = (v13 & v13)
                        }
                    case 767:
                        v20 := GenS1{F0: true, F1: 127, F2: ((float64(float32(1e10)) - 0.1) < -2.25)}
                        _ = v20
                    }
                    if true {
                        v6 += (v6 | v6)
                        switch v9 {
                        case "\n\t":
                            if false {
                                v21 := GenS1{F0: ((0.1 - 0.1) == 1.5), F1: (v2 >> ((255) & 63)), F2: (true || true)}
                                _ = v21
                            }
                        default:
                            switch (fi15 << ((613) & 63)) {
                            case -9223372036854775808:
                                var v22 uint = 65535
                                _ = v22
                            default:
                                var v23 uint = 18446744073709551615
                                _ = v23
                            }
                        }
                        v24 := GenS3{F0: 1, F1: 0, F2: "\n\t"}
                        _ = v24
                    }
                }
            case -2147483648:
                if (false && false) {
                    v5 = append(v5, (fi14 << ((uint(v2)) & 63)))
                }
            default:
                var v25 uint8 = (v2 ^ v2)
                _ = v25
            }
            var v26 bool = true
            _ = len(v5)
            _ = fi15
            _ = v26
        }
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 uint = 18446744073709551615
            _ = fi27
            _ = v28
        }
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
