package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    v3 := []string{"foo bar"}
    v4 := make(map[string]int)
    v5 := []string{(("a" + "foo bar") + ("\n\t" + "\n\t")), "\n\t", ("a" + ""), ("foo bar" + ("\n\t" + "foo bar"))}
    v6 := make(map[string]bool)
    var v7 uint32 = 42
    for fi8 := 0; fi8 < 13; fi8++ {
        v7 += (v7 + 708)
        if len(v3) > 0 {
            v3[0] = ""
        }
        v7 -= v7
        _ = fi8
    }
    var v9 uint32 = v7
    switch -2147483648 {
    case -1:
        v10 := GenS1{F0: 9223372036854775807, F1: 128, F2: "hello"}
        _ = v10
    case 1:
        for fi11 := 0; fi11 < 13; fi11++ {
            v5 = append(v5, "foo bar")
            if (0 < int32(v7)) {
                switch ("a" + "") {
                case "":
                    _ = len(v2)
                default:
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                }
            }
            var v12 uint16 = 255
            _ = fi11
            _ = v12
        }
    }
    var v13 uint32 = (v7 % (((v9 * v9)) | 1))
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v7 = (v13 % (((v7 - v7)) | 1))
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    var v14 float64 = ((float64(-2.25) * 3.14159) / 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
