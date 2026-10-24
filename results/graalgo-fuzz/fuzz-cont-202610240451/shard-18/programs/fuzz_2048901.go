package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 int16, p1 int8, p2 uint32) int8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        if ((1 <= 42) && (0 > 4294967295)) {
            switch "foo bar" {
            case "foo bar":
                v5 := []int{(fi4 / (((fi4 % ((1) | 1))) | 1)), fi4}
                _ = v5
            case "\n\t":
                fmt.Println(p1, p0)
            }
        }
        _ = fi4
    }
    switch -9223372036854775808 {
    case 128:
        p0 += p0
    case 127:
        v6 := []int{88, 4294967295}
        _ = v6
    case -918:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int8 = 0
            v9 := []int{(fi7 ^ fi7), fi7, 333}
            p2 -= 65535
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 + (p1 >> ((1) & 63)))
}
func genFn1() int16 {
    var v10 bool = true
    v10 = v10
    _ = v10
    return -656
}

func main() {
    v1 := []string{(("foo bar" + "hello") + ("foo bar" + ("\n\t" + "hello"))), "", "foo bar", ("hello" + "")}
    var v2 int = 230
    var v3 uint = 286
    v11 := GenS1{F0: -32768, F1: 179}
    switch "foo bar" {
    case "a":
        if (18446744073709551615 <= 18446744073709551615) {
            v2 = (v2 << ((18446744073709551615) & 63))
            switch v2 {
            case -658:
                v1 = append(v1, "\n\t")
            }
            v2 -= (v2 << (((^(v3))) & 63))
        }
    case "foo bar":
        if (((0.0 + (1e10 - 1.5)) - float64((float32((0.0 - 0.0)) - (3.14159 * 3.14159)))) == (float64((0.1 * 3.14159)) / (1e10 / (1.5 - 0.1)))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            fmt.Println(v3)
            v12 := GenS1{F0: 0, F1: 96}
            var v13 uint32 = 1
            _ = v12
            _ = v13
        }
    case "":
        v3 = (v3 & (^(v3)))
    }
    v3 += (v3 & (v3 * 4294967295))
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v14 int8 = 42
        v1 = append(v1, "a")
        _ = v14
    }
    var v15 int = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = 2
    fmt.Println(v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
}
