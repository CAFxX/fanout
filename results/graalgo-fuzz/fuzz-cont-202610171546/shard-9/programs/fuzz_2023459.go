package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 float64 = (((float64(3.14159) + 3.14159) / float64(3.14159)) - (0.1 / 0.1))
    v4 := []string{"\n\t"}
    var v5 uint8 = 255
    var v6 uint32 = 65535
    v3 = v3
    fmt.Println(v5, v5)
    v5 = v5
    v3 -= 3.14159
    v5 += 255
    var v7 float32 = 1.5
    v5 = (v5 % ((255) | 1))
    var v8 bool = (!false)
    switch 32767 {
    case 0:
        if true {
            var v9 int64 = -489
            if true {
                for fi10 := 0; fi10 < 15; fi10++ {
                    if v8 {
                        v11 := GenS1{F0: 247, F1: v9}
                        var v12 uint = 42
                        _ = len(v4)
                        _ = v11
                        _ = v12
                    }
                    var v13 bool = (v9 == (v9 << ((1) & 63)))
                    _ = fi10
                    _ = v13
                }
            } else {
                v9 = v9
            }
            _ = v9
        }
    case 1:
        v8 = (true || v8)
    case 673:
        v7 = v7
    default:
        if true {
            for fi14 := 0; fi14 < 15; fi14++ {
                switch fi14 {
                case -2147483648:
                    var v15 int32 = int32(fi14)
                    _ = v15
                case 2147483647:
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 string = ("a" + ("hello" + "a"))
                        _ = fi16
                        _ = v17
                    }
                case 32767:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        fmt.Println(fi18)
                        v4 = append(v4, ("" + ("hello" + (("a" + ("" + "")) + "foo bar"))))
                        fmt.Println(fi18, fi18)
                        v19 := GenS1{F0: 4294967295, F1: 839}
                        _ = fi18
                        _ = v19
                    }
                default:
                    _ = len(v2)
                }
                var v20 int64 = 447
                delete(v2, "hello")
                _ = fi14
                _ = v20
            }
            switch 32767 {
            case 127:
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := GenS2{F0: 668, F1: (!v8)}
                    _ = fi21
                    _ = v22
                }
            case 32768:
                v7 = (float32(((v3 * (((v3 * v3) / 1e10) + ((((v3 + float64(-2.25)) + v3) + 0.1) - (0.0 * 1e10)))) + (float64(((((1.5 / 0.1) - v7) - v7) - (v7 + v7))) * v3))) - (v7 * ((0.0 - v7) / 1e10)))
            case 1:
                var v23 int16 = 32767
                _ = v23
            default:
                v3 = v3
            }
            v24 := GenS3{F0: (^(v6)), F1: 100}
            _ = v24
        } else {
            var v25 uint8 = 42
            v4 = append(v4, "hello")
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        switch "" {
        case "\n\t":
            _ = len(v4)
        }
        if false {
            fmt.Println(v6)
        } else {
            v8 = (v8 && v8)
        }
        _ = fi26
    }
    var v27 bool = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
