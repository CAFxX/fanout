package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }

func main() {
    v1 := []int{896, -2147483648, 32767, 1}
    var v2 int64 = 9223372036854775807
    v3 := []string{"hello"}
    var v4 float32 = 3.14159
    var v5 bool = true
    v4 += v4
    var v6 bool = ((-683 > -1) || (v4 == 1e10))
    for fi7 := 0; fi7 < 18; fi7++ {
        v2 = 1
        v8 := GenS1{F0: 32767, F1: -272, F2: int32(v2)}
        _ = fi7
        _ = v8
    }
    var v9 int = 2147483647
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 = (v9 ^ (v9 - v9))
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            if (!(((-2.25 * -2.25) * 0.0) >= 0.1)) {
                v12 := GenS1{F0: -32768, F1: -808, F2: -2147483648}
                v13 := GenS1{F0: -32768, F1: 42, F2: 842}
                var v14 int32 = 180
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        switch 554 {
        case -1:
            fmt.Println(v9, v5, fi10)
        case 208:
            if false {
                v5 = (1 >= 0)
            } else {
                v4 += v4
            }
        default:
            var v15 int16 = 0
            _ = v15
        }
        _ = len(v1)
        _ = fi10
    }
    if (true || (127 >= 128)) {
        fmt.Println(v9)
    } else {
        if (v4 == (0.1 * v4)) {
            v6 = (!(float32(0.1) <= v4))
            for fi16 := 0; fi16 < 14; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    fmt.Println(fi17, v5)
                    v18 := GenS1{F0: -459, F1: -363, F2: 1}
                    _ = fi17
                    _ = v18
                }
                _ = len(v1)
                if ((v6 && ((true || (560 <= 65535)) && v6)) || true) {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        var v20 float64 = 1.5
                        var v21 uint8 = 0
                        if v5 {
                            if (!(1 < 255)) {
                                switch "" {
                                case "hello":
                                    v9 = 42
                                case "\n\t":
                                    v20 = v20
                                    v1 = append(v1, v9)
                                case "foo bar":
                                    v22 := GenS1{F0: -32768, F1: 42, F2: 1}
                                    _ = v22
                                }
                                v2 += (-(v2))
                            }
                            var v23 float64 = -2.25
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v24 := GenS2{F0: 0, F1: 42, F2: -276}
                    _ = v24
                }
                _ = fi16
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
