package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 uint8; F2 bool }

func main() {
    v1 := []uint32{285, 65535}
    v2 := make(map[uint32]string)
    var v3 int8 = 42
    var v4 uint8 = 128
    v5 := make(map[string]bool)
    var v6 uint16 = 2
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint16 = (^(v6))
        v6 = v6
        _ = fi7
        _ = v8
    }
    v1 = append(v1, 0)
    v6 = v6
    if false {
        var v9 uint16 = v6
        _ = v9
    } else {
        fmt.Println(v3, v4)
    }
    v6 -= (v6 * (^(v6)))
    v10 := GenS1{F0: -75}
    _ = len(v2)
    for fi11 := 0; fi11 < 5; fi11++ {
        fmt.Println(fi11, v4, v6)
        _ = fi11
    }
    if (false && (v6 == 174)) {
        for fi12 := 0; fi12 < 16; fi12++ {
            switch -2147483648 {
            case 0:
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 string = "foo bar"
                    _ = fi13
                    _ = v14
                }
            case 32767:
                for fi15 := 0; fi15 < 3; fi15++ {
                    v3 = (v3 ^ (v3 * v3))
                    v3 += (v3 >> ((196) & 63))
                    v4 += v4
                    _ = fi15
                }
            case -750:
                v16 := GenS1{F0: (v3 / (((^(v3))) | 1))}
                _ = v16
            }
            v6 = (v6 + 65535)
            _ = fi12
        }
        fmt.Println(v6)
        v17 := GenS2{F0: 841, F1: (^(v6)), F2: -312}
        _ = v17
    } else {
        v3 -= (v3 ^ (v3 % ((v3) | 1)))
        var v18 int32 = -869
        v6 -= v6
        _ = v18
    }
    var v19 int8 = 100
    if len(v1) > 0 {
        v1[0] = 42
    }
    v6 = (v6 - v6)
    if (32767 == v6) {
        if (!((!(42 > 0)) && (-703 <= 1))) {
            v20 := GenS3{F0: 128, F1: 1, F2: false}
            v21 := v2[838]
            var v22 int = 577
            _ = v20
            _ = v21
            _ = v22
        }
        var v23 uint = 4294967295
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS1{F0: (v19 << (((v23 + (^(v23)))) & 63))}
            switch (("a" + ("\n\t" + "hello")) + "a") {
            case "hello":
                for fi26 := 0; fi26 < 1; fi26++ {
                    if false {
                        var v27 uint16 = 65535
                        v28 := GenS1{F0: v19}
                        switch (fi24 & -1) {
                        case -1:
                            fmt.Println(v6, v19)
                        }
                        v4 -= v4
                        v6 = v6
                        _ = v27
                        _ = v28
                    }
                    _ = fi26
                }
            case "foo bar":
                v29 := GenS1{F0: -62}
                _ = v29
            }
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 int8 = -1
                _ = fi30
                _ = v31
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 string = "foo bar"
        fmt.Println(fi32)
        _ = fi32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
