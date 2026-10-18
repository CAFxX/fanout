package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483647
    var v3 int64 = -787
    v4 := []uint32{(v2 & 2147483648), (^(v2))}
    var v5 float32 = -2.25
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: -571}
    var v8 string = "foo bar"
    v6 = (v6 / ((910) | 1))
    fmt.Println(v3, v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = v6
    fmt.Println(v8)
    if (!true) {
        if (18446744073709551615 >= 1) {
            for fi9 := 0; fi9 < 9; fi9++ {
                for fi10 := 0; fi10 < 5; fi10++ {
                    if (!true) {
                        v2 += (v2 + v6)
                    } else {
                        var v11 int = fi10
                        v2 = (v6 & v2)
                        if true {
                            var v12 float32 = ((((v5 - 1e10) * (((float32(1.5) - v5) * 3.14159) - (v5 - 1e10))) - (1.5 + ((v5 * 1.5) + v5))) + (v5 / (v5 + v5)))
                            _ = v12
                        }
                        _ = v11
                    }
                    var v13 float32 = (float32(0.0) / (v5 - v5))
                    v3 -= 0
                    _ = fi10
                    _ = v13
                }
                var v14 int64 = 0
                _ = fi9
                _ = v14
            }
        }
        switch v8 {
        case "foo bar":
            v1["\n\t"] = -2147483648
        default:
            v5 = 1.5
        }
    } else {
        var v15 int8 = -28
        _ = v15
    }
    if false {
        switch v8 {
        case "hello":
            v5 += v5
        }
        v3 -= -1
        for fi16 := 0; fi16 < 19; fi16++ {
            v2 = v2
            switch "\n\t" {
            case "foo bar":
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 uint8 = 12
                    _ = fi17
                    _ = v18
                }
            case "hello":
                fmt.Println(v3, fi16)
            case "\n\t":
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        v19 := GenS1{F0: -1}
                        _ = v19
                    case "a":
                        var v20 float32 = v5
                        _ = v20
                    default:
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v2 -= v2
                            _ = fi21
                        }
                    }
                    if true {
                        v22 := GenS1{F0: 788}
                        var v23 float64 = 0.0
                        _ = v22
                        _ = v23
                    }
                    var v24 uint64 = 0
                    _ = v24
                }
            }
            _ = fi16
        }
    }
    var v25 int32 = 42
    for fi26 := 0; fi26 < 4; fi26++ {
        if len(v4) > 0 {
            v4[0] = (v6 + v2)
        }
        var v27 int32 = (v25 + 42)
        _ = fi26
        _ = v27
    }
    _ = len(v1)
    v25 -= v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
