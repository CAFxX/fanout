package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int64 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]uint32)
    var v3 float64 = (0.1 + (float64(-2.25) - 3.14159))
    var v4 float32 = v1
    var v5 int8 = 100
    var v6 uint16 = 65535
    fmt.Println(v3, v5, v3)
    for fi7 := 0; fi7 < 15; fi7++ {
        switch 0 {
        case 666:
            var v8 int = (-(fi7))
            _ = v8
        default:
            if false {
                for fi9 := 0; fi9 < 12; fi9++ {
                    fmt.Println(fi7)
                    v4 = 0.0
                    _ = fi9
                }
            } else {
                v4 += 0.1
                v6 -= 0
                delete(v2, "\n\t")
            }
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            v3 -= -2.25
            _ = fi10
        }
        _ = fi7
    }
    fmt.Println(v3, v3, v6)
    v4 += (v1 + (3.14159 + v1))
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(v1, v5, fi11)
        for fi12 := 0; fi12 < 16; fi12++ {
            v6 = (v6 | v6)
            var v13 int = fi12
            switch "" {
            case "a":
                if false {
                    var v14 uint32 = 65535
                    v15 := v2[("hello" + "hello")]
                    if true {
                        var v16 float32 = v4
                        v2[("" + "")] = 2
                        _ = v16
                    } else {
                        var v17 float64 = (v3 + 1.5)
                        if (42 != int32(v6)) {
                            v18 := GenS1{F0: 7}
                            fmt.Println(v6, v13)
                            v19 := GenS2{F0: -128}
                            _ = v18
                            _ = v19
                        }
                        v20 := []int{(-(v13)), fi12}
                        _ = v17
                        _ = v20
                    }
                    _ = v14
                    _ = v15
                } else {
                    v5 += v5
                    if false {
                        v6 -= v6
                        v21 := []int{}
                        _ = v21
                    }
                }
            case "\n\t":
                v22 := GenS3{F0: "", F1: -751}
                _ = v22
            case "hello":
                v23 := GenS2{F0: (v5 & (^(v5)))}
                _ = v23
            default:
                v24 := []int{(fi12 & v13)}
                _ = v24
            }
            var v25 bool = false
            v4 -= 3.14159
            _ = fi12
            _ = v13
            _ = v25
        }
        _ = fi11
    }
    var v26 uint = 341
    v27 := v2[("" + "")]
    for fi28 := 0; fi28 < 17; fi28++ {
        _ = len(v2)
        v29 := []int{fi28, (fi28 >> (((v26 + (v26 & 18446744073709551615))) & 63)), fi28, (fi28 ^ -2147483648)}
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
