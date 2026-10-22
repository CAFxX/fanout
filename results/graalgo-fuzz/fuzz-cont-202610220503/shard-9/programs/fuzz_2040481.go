package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }

func main() {
    v1 := []string{"foo bar", (("" + "foo bar") + "foo bar")}
    v2 := []int{-512, 282, -9223372036854775808}
    v3 := make(map[uint32]int)
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: -32768, F1: -779}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        if true {
            if (2 <= 1) {
                delete(v3, 927)
                if (772 != -216) {
                    switch fi6 {
                    case 128:
                        var v7 int = (^(fi6))
                        _ = v7
                    }
                    for fi8 := 0; fi8 < 13; fi8++ {
                        var v9 int32 = -25
                        _ = fi8
                        _ = v9
                    }
                }
                for fi10 := 0; fi10 < 3; fi10++ {
                    _ = len(v3)
                    fmt.Println(fi6, fi10)
                    var v11 float64 = (0.0 + 0.1)
                    _ = fi10
                    _ = v11
                }
            } else {
                fmt.Println(fi6)
                if (!false) {
                    fmt.Println(fi6)
                } else {
                    var v12 float32 = (3.14159 - 3.14159)
                    v12 += 0.0
                    _ = v12
                }
            }
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            switch fi13 {
            case 32768:
                delete(v3, 65535)
            case 0:
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            default:
                if (!true) {
                    var v14 float32 = (float32(1.5) / 3.14159)
                    delete(v3, 480)
                    v15 := GenS1{F0: 32767, F1: 42}
                    _ = v14
                    _ = v15
                }
            }
            _ = fi13
        }
        var v16 int = fi6
        _ = fi6
        _ = v16
    }
    var v17 int = 128
    for fi18 := 0; fi18 < 9; fi18++ {
        switch "" {
        case "":
            fmt.Println(fi18, v17)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v17 += 9223372036854775807
            _ = fi19
        }
        _ = fi18
    }
    var v20 uint16 = 343
    if false {
        if true {
            v21 := GenS1{F0: -32768, F1: 0}
            _ = len(v3)
            if (("hello" + ("a" + "foo bar")) <= ("\n\t" + (("\n\t" + ("\n\t" + "\n\t")) + "foo bar"))) {
                v3[901] = (v17 / ((v17) | 1))
                var v22 uint32 = 2
                fmt.Println(v22, v22)
                _ = v22
            } else {
                v20 -= v20
                fmt.Println(v20, v17)
            }
            _ = v21
        }
    }
    v17 -= v17
    v17 += (v17 % (((v17 - (^(v17)))) | 1))
    v23 := GenS1{F0: 42, F1: 0}
    var v24 uint8 = 255
    var v25 uint64 = 18446744073709551615
    for fi26 := 0; fi26 < 8; fi26++ {
        switch ("a" + "") {
        case "":
            var v27 uint64 = (^(v25))
            _ = v27
        case "\n\t":
            var v28 uint16 = (^(v20))
            _ = v28
        default:
            var v29 float64 = 1.5
            _ = v29
        }
        _ = fi26
    }
    var v30 int32 = 28
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
}
