package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 uint8; F1 uint; F2 uint32 }
func genFn0(p0 int) int64 {
    v5 := []int{32768, p0, (p0 + (p0 * (p0 % ((p0) | 1))))}
    var v6 int16 = 1
    _ = p0
    _ = v5
    _ = v6
    return int64(p0)
}

func main() {
    v1 := []string{"foo bar"}
    var v2 float64 = -2.25
    v3 := []int{9223372036854775807, -9223372036854775808}
    var v4 int = 32767
    v7 := GenS1{F0: "\n\t", F1: 2, F2: (v4 / (((v4 | v4)) | 1))}
    if false {
        switch "\n\t" {
        case "":
            v4 = v4
        }
        v8 := GenS1{F0: ("hello" + "foo bar"), F1: 32767, F2: (v4 ^ (^(v4)))}
        fmt.Println(v2)
        _ = v8
    }
    var v9 int16 = 1000
    v2 += (v2 / v2)
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v9, v2, v9)
        var v11 bool = (42 >= 231)
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "hello":
        v12 := GenS1{F0: "\n\t", F1: 161, F2: (-(v4))}
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v13 := GenS1{F0: "foo bar", F1: 425, F2: v4}
    v14 := GenS1{F0: "", F1: 65535, F2: v4}
    switch (v4 % (((v4 | v4)) | 1)) {
    case 985:
        switch (^(v4)) {
        case -1:
            v2 = (float64(1e10) * ((v2 * v2) + ((v2 - float64(1.5)) + (-2.25 / (0.0 - (v2 + (0.0 + 3.14159)))))))
        }
    case -898:
        if (1 < -854) {
            if true {
                fmt.Println(v4, v9, v9)
                v15 := GenS2{F0: 125, F1: 42, F2: 559}
                var v16 uint = 42
                _ = v15
                _ = v16
            } else {
                _ = len(v1)
                var v17 int16 = v9
                v1 = append(v1, "\n\t")
                _ = v17
            }
            var v18 uint = 4294967295
            v4 -= -1
            _ = v18
        }
    case 128:
        v4 += (v4 | (v4 + v4))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
