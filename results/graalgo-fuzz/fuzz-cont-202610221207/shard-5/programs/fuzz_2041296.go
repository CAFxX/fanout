package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int16 }

func main() {
    v1 := []string{"foo bar", "", "hello"}
    var v2 bool = ((true || ((!false) || true)) && true)
    v3 := []uint32{1, 561, 4294967295, 101}
    v3 = append(v3, 2147483647)
    var v4 string = "a"
    switch 32768 {
    case 32768:
        for fi5 := 0; fi5 < 2; fi5++ {
            switch 0 {
            case 736:
                if v2 {
                    v3 = append(v3, 1)
                }
            case 80:
                if v2 {
                    v6 := GenS1{F0: 389}
                    if v2 {
                        var v7 int32 = -28
                        var v8 uint64 = 109
                        _ = v7
                        _ = v8
                    } else {
                        v2 = ((42 != 1) && true)
                        for fi9 := 0; fi9 < 14; fi9++ {
                            v1 = append(v1, v4)
                            _ = fi9
                        }
                    }
                    _ = v6
                } else {
                    v2 = (1.5 == (1.5 - -2.25))
                    _ = len(v3)
                }
            default:
                for fi10 := 0; fi10 < 9; fi10++ {
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    var v11 string = v4
                    _ = fi10
                    _ = v11
                }
            }
            var v12 bool = ((!true) || v2)
            _ = fi5
            _ = v12
        }
    }
    var v13 int64 = 42
    var v14 int = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 float32 = float32(0.0)
        v17 := GenS2{F0: 745, F1: 32767, F2: -800}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint16 = 236
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
