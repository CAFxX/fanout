package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint, p1 uint) uint8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        p1 = p1
        _ = fi6
    }
    var v7 int16 = 32767
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 bool = (1 < 42)
    var v2 float64 = -2.25
    var v3 int16 = 1000
    var v4 string = (((("a" + "a") + "\n\t") + ("" + "\n\t")) + (("" + "") + "foo bar"))
    v5 := []uint32{531, 42, 287}
    if (9223372036854775807 >= 0) {
        var v8 float32 = 0.0
        var v9 uint32 = 2147483648
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if v1 {
                v3 = (v3 % ((677) | 1))
            }
            _ = fi11
        }
        if false {
            v3 = v3
            if true {
                fmt.Println(v3)
            } else {
                var v12 bool = (true && v1)
                v13 := GenS1{F0: false, F1: 0, F2: -128}
                if len(v5) > 0 {
                    v5[0] = 2
                }
                _ = v12
                _ = v13
            }
            v1 = (v1 && (!v1))
        } else {
            if true {
                switch fi10 {
                case 176:
                    v4 = ""
                case 834:
                    for fi14 := 0; fi14 < 13; fi14++ {
                        fmt.Println(v1)
                        var v15 string = v4
                        fmt.Println(fi14, v15)
                        _ = fi14
                        _ = v15
                    }
                default:
                    v2 = v2
                }
                fmt.Println(v4, v2, v3)
            }
            if (0 < 737) {
                var v16 int32 = -2147483648
                _ = v16
            } else {
                fmt.Println(v3, v4)
            }
        }
        _ = fi10
    }
    var v17 uint16 = 0
    fmt.Println(v4)
    v4 = (("\n\t" + "\n\t") + v4)
    switch 2147483647 {
    case -2147483648:
        v3 = (v3 % ((v3) | 1))
    case 128:
        fmt.Println(v2, v2)
    case 32767:
        fmt.Println(v2, v17)
    }
    v3 = 32767
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS2{F0: 128, F1: 4294967295}
        fmt.Println(v4)
        _ = fi18
        _ = v19
    }
    if len(v5) > 0 {
        v5[0] = 29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
