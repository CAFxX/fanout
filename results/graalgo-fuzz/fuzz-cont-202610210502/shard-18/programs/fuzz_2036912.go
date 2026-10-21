package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 string }

func main() {
    var v1 int = 780
    var v2 int64 = 390
    v3 := []int{v1, v1, (v1 << ((1) & 63))}
    v2 -= -1
    if (32767 < 1000) {
        if true {
            for fi4 := 0; fi4 < 11; fi4++ {
                v3 = append(v3, (fi4 >> ((255) & 63)))
                var v5 int64 = (v2 | (v2 - (v2 - 1)))
                _ = fi4
                _ = v5
            }
        } else {
            _ = len(v3)
        }
        v2 = (v2 % (((v2 >> ((4294967295) & 63))) | 1))
    }
    v6 := GenS1{F0: "a", F1: 359, F2: (("a" + "a") + "foo bar")}
    if ((v1 / ((v1) | 1)) <= (v1 << ((194) & 63))) {
        if len(v3) > 0 {
            v3[0] = (v1 ^ 32768)
        }
        v7 := GenS2{F0: "a"}
        v8 := GenS1{F0: "\n\t", F1: 2147483647, F2: ""}
        _ = v7
        _ = v8
    }
    switch (v1 - (v1 ^ v1)) {
    case 128:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -675:
        _ = len(v3)
    default:
        _ = len(v3)
    }
    fmt.Println(v1)
    var v9 bool = true
    switch ("hello" + "foo bar") {
    case "a":
        var v10 int32 = 42
        _ = v10
    case "foo bar":
        var v11 uint32 = 2147483648
        _ = v11
    case "\n\t":
        if (!v9) {
            v9 = v9
        } else {
            v12 := GenS1{F0: "foo bar", F1: 2147483647, F2: "foo bar"}
            var v13 int8 = -1
            _ = v12
            _ = v13
        }
    default:
        var v14 uint8 = 46
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
