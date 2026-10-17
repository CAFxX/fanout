package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }

func main() {
    var v1 int = 1
    var v2 string = "hello"
    v3 := []string{"\n\t", v2, v2}
    v4 := []int{}
    var v5 int8 = 42
    v6 := []string{"a", "hello", "hello", v2}
    v1 = (v1 ^ -521)
    for fi7 := 0; fi7 < 3; fi7++ {
        v2 = (v2 + "")
        _ = fi7
    }
    v8 := GenS1{F0: 42, F1: -2147483648, F2: -658}
    var v9 bool = (-556 >= 1000)
    var v10 string = ""
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: 65535, F1: 0, F2: 42}
            fmt.Println(v9, fi12)
            var v14 int = fi12
            _ = fi12
            _ = v13
            _ = v14
        }
        v2 = ((v10 + v2) + (v2 + v10))
        _ = fi11
    }
    v1 += 127
    v15 := GenS1{F0: 255, F1: -214, F2: 9223372036854775807}
    for fi16 := 0; fi16 < 14; fi16++ {
        if len(v6) > 0 {
            v6[0] = (v2 + "a")
        }
        v2 = v10
        v17 := GenS1{F0: 959, F1: -2147483648, F2: 31}
        _ = fi16
        _ = v17
    }
    switch "" {
    case "a":
        if (194 >= 197) {
            if false {
                var v18 string = "foo bar"
                _ = v18
            } else {
                if (!v9) {
                    v1 -= v1
                    for fi19 := 0; fi19 < 1; fi19++ {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v1 = v1
                            v5 = 1
                            v3 = append(v3, v10)
                            _ = fi20
                        }
                        var v21 int32 = 0
                        _ = fi19
                        _ = v21
                    }
                } else {
                    v5 += v5
                    v22 := GenS1{F0: 2, F1: -2147483648, F2: int64(v1)}
                    _ = v22
                }
                v2 = "\n\t"
                var v23 float64 = 0.1
                _ = v23
            }
        }
    default:
        v5 += (-(v5))
    }
    var v24 uint16 = 735
    v10 = v10
    var v25 float64 = ((1e10 - (3.14159 / (0.1 / 1e10))) / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
