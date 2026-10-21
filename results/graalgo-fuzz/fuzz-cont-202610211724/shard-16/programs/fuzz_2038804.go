package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 2147483647
    v3 := []int{197, -1}
    var v4 uint32 = 2147483648
    var v5 bool = false
    v5 = ((0.0 + 1e10) >= (0.1 * 3.14159))
    var v6 uint16 = 0
    if len(v3) > 0 {
        v3[0] = 42
    }
    if (v5 || (0 >= 480)) {
        _ = len(v1)
    } else {
        v5 = (-81 == 0)
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: (v2 & v2)}
            v6 += (^(v6))
            fmt.Println(v2, v6)
            _ = fi7
            _ = v8
        }
    }
    v5 = v5
    if v5 {
        v9 := GenS2{F0: -485, F1: 408}
        _ = v9
    } else {
        var v10 int32 = (v2 + (v2 + v2))
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := v1[(("hello" + ("" + "a")) + "\n\t")]
        var v13 int = (fi11 & fi11)
        _ = fi11
        _ = v12
        _ = v13
    }
    if (-3 == 0) {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := v1["\n\t"]
            fmt.Println(v5)
            _ = fi14
            _ = v15
        }
    }
    var v16 bool = (v5 || (255 <= 18446744073709551615))
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            if (0 != -1) {
                switch ("" + "a") {
                case "\n\t":
                    switch "" {
                    case "a":
                        fmt.Println(fi17, fi18)
                    case "":
                        var v19 float64 = 1.5
                        _ = v19
                    }
                default:
                    v20 := GenS2{F0: 159, F1: v6}
                    fmt.Println(v16)
                    _ = v20
                }
                v21 := GenS2{F0: (^(fi17)), F1: v6}
                if v16 {
                    if len(v3) > 0 {
                        v3[0] = fi17
                    }
                    v2 = v2
                }
                v3 = append(v3, 42)
                _ = v21
            }
            var v22 int8 = 0
            _ = fi18
            _ = v22
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
