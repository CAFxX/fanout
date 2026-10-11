package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 string; F2 int32 }

func main() {
    var v1 string = ("" + ("" + ("hello" + "\n\t")))
    var v2 int64 = 0
    v3 := make(map[string]string)
    var v4 uint32 = 969
    if true {
        var v5 bool = (false || true)
        _ = v5
    } else {
        var v6 int8 = 100
        v7 := GenS1{F0: "a"}
        if (-2147483648 <= 101) {
            v8 := GenS1{F0: ("\n\t" + (v1 + "foo bar"))}
            _ = v8
        }
        _ = v6
        _ = v7
    }
    _ = len(v3)
    switch v1 {
    case "\n\t":
        v9 := GenS1{F0: v1}
        _ = v9
    default:
        v10 := GenS1{F0: v1}
        _ = v10
    }
    v2 = -1
    var v11 float64 = 1.5
    var v12 uint32 = (v4 * (v4 & 997))
    var v13 uint16 = 418
    var v14 uint8 = 1
    v4 = 1
    if (false && true) {
        switch 32767 {
        case 813:
            fmt.Println(v13, v13)
        case -2147483648:
            if (-32768 >= 887) {
                delete(v3, (v1 + v1))
                v13 += 2
            } else {
                if false {
                    v15 := []int{32767, -2147483648, 4294967295, -211}
                    _ = v15
                }
                switch (v1 + v1) {
                case "a":
                    v16 := v3[v1]
                    _ = v16
                }
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS2{F0: 0, F1: v1, F2: int32(fi17)}
                    var v19 float64 = v11
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        }
        switch v1 {
        case "foo bar":
            fmt.Println(v2, v12)
        case "a":
            var v20 uint64 = 972
            _ = v20
        case "\n\t":
            if (!(true || (float64(3.14159) != (v11 - (1.5 / v11))))) {
                var v21 float64 = v11
                switch 42 {
                case 128:
                    v22 := GenS2{F0: 998, F1: "foo bar", F2: -1}
                    _ = v22
                case 32768:
                    v12 = (v4 + (v12 * (v12 << ((18446744073709551615) & 63))))
                case 0:
                    var v23 int32 = -2147483648
                    _ = v23
                default:
                    fmt.Println(v11)
                }
                _ = v21
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
