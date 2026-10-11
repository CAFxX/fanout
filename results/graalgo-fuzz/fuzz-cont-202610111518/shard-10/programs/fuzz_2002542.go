package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 uint = 1
    var v4 uint64 = 794
    var v5 int8 = 83
    v3 = v3
    var v6 int16 = -1
    if (-845 <= 2147483647) {
        var v7 uint64 = v4
        switch -1 {
        case 61:
            v2[127] = 441
        case 128:
            var v8 uint8 = 1
            _ = v8
        case 127:
            v5 = (-(v5))
        }
        _ = v7
    }
    switch "" {
    case "\n\t":
        v9 := []int{-797, 0, -404, 9223372036854775807}
        _ = v9
    case "hello":
        if (!(float64(0.0) >= 0.1)) {
            var v10 uint16 = 434
            for fi11 := 0; fi11 < 1; fi11++ {
                v6 -= v6
                _ = fi11
            }
            _ = v10
        }
    case "":
        switch "" {
        case "\n\t":
            _ = len(v1)
        case "foo bar":
            v4 += 2
        }
    }
    var v12 int8 = (v5 + -128)
    var v13 int64 = 18
    if (1e10 >= float64(3.14159)) {
        v5 -= (v5 & -128)
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: "hello"}
            _ = fi14
            _ = v15
        }
        var v16 float32 = 0.1
        _ = v16
    }
    v17 := GenS1{F0: "a"}
    switch -828 {
    case 32768:
        v1["\n\t"] = 1
    case 887:
        switch (("\n\t" + "") + "\n\t") {
        case "":
            v6 = (v6 ^ v6)
        case "hello":
            v18 := GenS1{F0: (("foo bar" + "hello") + "")}
            _ = v18
        }
    }
    var v19 bool = (!(v3 == (^(v3))))
    v3 -= (v3 | (v3 - (v3 * v3)))
    v20 := GenS2{F0: v12, F1: (v3 >> (((v3 - 0)) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
