package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int; F1 bool; F2 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 string = ""
    var v3 int32 = 0
    var v4 int = 0
    var v5 string = v2
    delete(v1, (v2 + (v2 + v5)))
    if false {
        v6 := GenS1{F0: (v3 / ((1) | 1)), F1: 42}
        v2 = ""
        _ = v6
    }
    v7 := GenS1{F0: (v3 << ((18446744073709551615) & 63)), F1: 800}
    switch ((v5 + v5) + v2) {
    case "a":
        v1[("foo bar" + ("" + v2))] = v4
    }
    switch (v4 % (((v4 * (v4 % ((v4) | 1)))) | 1)) {
    case 390:
        delete(v1, v5)
    }
    v8 := v1[v2]
    v2 = ("a" + v5)
    if false {
        var v9 float64 = 3.14159
        _ = v9
    }
    var v10 uint = 794
    fmt.Println(v3)
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v2, v2, v10)
        if (fi11 == (v4 * (fi11 & (fi11 ^ (^(fi11)))))) {
            delete(v1, v2)
        } else {
            var v12 string = "\n\t"
            switch fi11 {
            case -130:
                switch v8 {
                case -2147483648:
                    var v13 float32 = -2.25
                    _ = v13
                case 32768:
                    var v14 bool = false
                    _ = v14
                default:
                    v10 = 1
                }
            case -1:
                if false {
                    if (2 >= 1) {
                        var v15 uint = 18446744073709551615
                        v1["foo bar"] = (v8 / ((-9223372036854775808) | 1))
                        v1[("\n\t" + (("" + v2) + v12))] = v8
                        _ = v15
                    } else {
                        var v16 uint = v10
                        v17 := []int{1}
                        v18 := GenS1{F0: v3, F1: 0}
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v4 += 2147483647
                }
            default:
                var v19 string = "\n\t"
                _ = v19
            }
            v20 := GenS2{F0: (^(v8)), F1: false, F2: v3}
            _ = v12
            _ = v20
        }
        _ = fi11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
