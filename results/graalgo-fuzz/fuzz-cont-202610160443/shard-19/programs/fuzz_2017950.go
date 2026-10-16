package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 0.1
    var v3 uint32 = 1
    var v4 string = "\n\t"
    var v5 uint = 65535
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 2147483647, F1: -36}
    v5 = 0
    if true {
        switch 42 {
        case 938:
            v8 := v6[(v3 & 65535)]
            _ = v8
        case 2147483647:
            var v9 bool = (1 <= 1)
            _ = v9
        default:
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 float32 = 3.14159
                v12 := GenS1{F0: 0, F1: 100}
                v2 -= (v2 - v11)
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        switch "hello" {
        case "\n\t":
            v4 = v4
        default:
            for fi13 := 0; fi13 < 6; fi13++ {
                if true {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        if true {
                            v15 := GenS1{F0: int(v3), F1: 100}
                            _ = v15
                        }
                        var v16 uint8 = 128
                        _ = fi14
                        _ = v16
                    }
                    var v17 int64 = 0
                    var v18 float64 = (1e10 - float64((float32(0.0) + 0.0)))
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS1{F0: (fi13 % ((-405) | 1)), F1: 123}
                    v4 = (v4 + v4)
                    _ = v19
                }
                var v20 uint64 = 18446744073709551615
                v3 -= v3
                _ = fi13
                _ = v20
            }
        }
        v5 -= 71
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        if false {
            v4 = (v4 + (v4 + (v4 + "hello")))
        }
        switch "" {
        case "\n\t":
            delete(v6, (^(v3)))
        default:
            v22 := v6[v3]
            _ = v22
        }
        var v23 int32 = 682
        _ = fi21
        _ = v23
    }
    switch (v4 + v4) {
    case "hello":
        for fi24 := 0; fi24 < 6; fi24++ {
            v3 = (v3 & v3)
            var v25 float32 = v2
            v1[v4] = (true || true)
            _ = fi24
            _ = v25
        }
    case "foo bar":
        if true {
            switch 463 {
            case 325:
                v26 := []int{0, -182, -141}
                _ = v26
            case -586:
                switch 127 {
                case -2147483648:
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v28 := v6[2]
                        var v29 uint = 951
                        switch -1 {
                        case -1:
                            fmt.Println(v3, v5, v28)
                        case -2147483648:
                            for fi30 := 0; fi30 < 17; fi30++ {
                                fmt.Println(fi27, v28)
                                v5 = (v29 & 48)
                                _ = fi30
                            }
                        case -728:
                            fmt.Println(v3, v3)
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
            case -833:
                var v31 uint64 = 1
                _ = v31
            default:
                for fi32 := 0; fi32 < 13; fi32++ {
                    v2 -= float32(((3.14159 * 1.5) - 0.0))
                    _ = len(v6)
                    _ = fi32
                }
            }
            fmt.Println(v3, v3, v2)
        } else {
            var v33 uint8 = 53
            _ = v33
        }
    case "a":
        var v34 int16 = 608
        _ = v34
    default:
        var v35 uint32 = v3
        _ = v35
    }
    if (32767 != 65535) {
        v36 := []int{-287, 32768, 2147483647}
        v5 += 85
        _ = v36
    }
    v4 = v4
    v37 := GenS1{F0: 42, F1: 94}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
