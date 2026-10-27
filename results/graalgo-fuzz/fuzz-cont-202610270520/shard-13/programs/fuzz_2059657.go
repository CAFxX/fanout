package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 bool }

func main() {
    var v1 bool = true
    var v2 uint16 = 675
    var v3 uint = 493
    var v4 int = -294
    v5 := []int{(v4 | v4)}
    var v6 uint16 = (v2 >> (((v3 >> ((1) & 63))) & 63))
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: 104}
            _ = fi8
            _ = v9
        }
        if (fi7 > (fi7 ^ fi7)) {
            var v10 uint8 = 127
            if v1 {
                var v11 int16 = int16(fi7)
                switch "" {
                case "foo bar":
                    fmt.Println(v2, v1, fi7)
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v11
            } else {
                var v12 string = ("hello" + "hello")
                var v13 int32 = -2147483648
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            switch (fi7 >> (((^(v3))) & 63)) {
            case 636:
                if true {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        var v15 uint64 = 18446744073709551615
                        v16 := GenS1{F0: 2}
                        var v17 int16 = 1000
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    if true {
                        v6 -= (v2 >> (((v3 * 1)) & 63))
                        var v18 string = "hello"
                        fmt.Println(v2, v4)
                        _ = v18
                    } else {
                        v19 := GenS2{F0: 4294967295}
                        v20 := GenS2{F0: 42}
                        _ = v19
                        _ = v20
                    }
                }
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v21 := GenS1{F0: 128}
            _ = v21
        }
        _ = fi7
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        if v1 {
            var v23 bool = v1
            v6 -= v2
            for fi24 := 0; fi24 < 8; fi24++ {
                v5 = append(v5, (v4 * 32767))
                _ = fi24
            }
            _ = v23
        }
        v25 := GenS2{F0: 438}
        _ = fi22
        _ = v25
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 int32 = -2147483648
            v1 = (v4 >= -1)
            v4 = 32767
            _ = fi27
            _ = v28
        }
        v29 := GenS1{F0: 255}
        _ = fi26
        _ = v29
    }
    switch v4 {
    case 32768:
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 float32 = 1.5
            _ = fi30
            _ = v31
        }
    case 9223372036854775807:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 100:
        var v32 int16 = 0
        _ = v32
    }
    if v1 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = len(v5)
    }
    v1 = false
    if (2 > (v3 * v3)) {
        v33 := GenS2{F0: 2}
        _ = v33
    } else {
        if v1 {
            var v34 uint64 = 4294967295
            _ = v34
        }
        v5 = append(v5, (v4 * -2147483648))
    }
    if (v3 >= 924) {
        v35 := GenS3{F0: (v6 + 547), F1: (v2 + (v6 & v2)), F2: false}
        v4 += -696
        _ = v35
    } else {
        for fi36 := 0; fi36 < 7; fi36++ {
            fmt.Println(fi36)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi36
        }
        v4 = v4
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
