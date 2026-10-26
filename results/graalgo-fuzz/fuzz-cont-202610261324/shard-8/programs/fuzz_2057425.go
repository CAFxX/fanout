package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }

func main() {
    var v1 uint = 0
    var v2 int16 = -1
    var v3 float64 = (0.1 * 0.0)
    var v4 uint16 = 923
    var v5 float32 = float32(-2.25)
    v6 := make(map[uint32]bool)
    v4 = (v4 * v4)
    var v7 int8 = -56
    var v8 uint = v1
    var v9 int16 = v2
    fmt.Println(v7)
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{(fi10 * (fi10 / ((32767) | 1))), fi10, fi10}
        delete(v6, 369)
        v2 = (v9 / (((v9 % ((1) | 1))) | 1))
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    fmt.Println(v2, v5)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int8 = 100
        _ = fi13
        _ = v14
    }
    if (((^(v8)) >= (v1 << ((v8) & 63))) && (v7 != int8(v4))) {
        for fi15 := 0; fi15 < 11; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                if (v5 < (((3.14159 - v5) + v5) - (1.5 * (1.5 * v5)))) {
                    var v17 int8 = -1
                    _ = v17
                } else {
                    var v18 uint64 = 18446744073709551615
                    fmt.Println(v4, v1)
                    switch fi15 {
                    case 2147483647:
                        v19 := GenS1{F0: (v9 ^ v2)}
                        _ = v19
                    case -9223372036854775808:
                        var v20 uint32 = 65535
                        _ = v20
                    default:
                        if (true && false) {
                            var v21 int16 = (v2 & 42)
                            v22 := GenS2{F0: (v2 % (((-(v21))) | 1)), F1: 42, F2: -2147483648}
                            _ = v21
                            _ = v22
                        }
                    }
                    fmt.Println(v2)
                    _ = v18
                }
                var v23 bool = true
                _ = fi16
                _ = v23
            }
            switch ((("\n\t" + "foo bar") + ("\n\t" + "a")) + "a") {
            case "":
                for fi24 := 0; fi24 < 16; fi24++ {
                    v25 := GenS1{F0: 934}
                    _ = fi24
                    _ = v25
                }
            }
            _ = fi15
        }
        if true {
            var v26 float32 = 3.14159
            _ = v26
        } else {
            v8 -= (v8 << ((v1) & 63))
            delete(v6, 439)
        }
        var v27 int16 = (v2 % ((42) | 1))
        _ = v27
    }
    var v28 uint64 = 18446744073709551615
    switch (((("\n\t" + "hello") + ("\n\t" + "hello")) + "\n\t") + "") {
    case "":
        v1 -= 1
    case "\n\t":
        switch 843 {
        case 128:
            v2 = (v9 * v2)
        case -153:
            delete(v6, 2)
        case -2147483648:
            switch "a" {
            case "":
                var v29 int = 4294967295
                _ = v29
            case "hello":
                var v30 int8 = v7
                _ = v30
            default:
                v31 := v6[uint32(v28)]
                _ = v31
            }
        default:
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 int32 = 1
                var v34 int16 = v2
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    case "foo bar":
        for fi35 := 0; fi35 < 4; fi35++ {
            _ = len(v12)
            switch fi35 {
            case 1:
                fmt.Println(v1, v7, v28)
            case 42:
                v3 = (v3 / v3)
            default:
                for fi36 := 0; fi36 < 20; fi36++ {
                    v1 += v1
                    v9 -= int16(v7)
                    _ = fi36
                }
            }
            _ = fi35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
