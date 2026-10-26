package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 bool = ((!true) || (-372 == -32768))
    var v5 uint32 = 42
    v4 = v4
    var v6 string = ("a" + (("" + (("\n\t" + ("hello" + "foo bar")) + "")) + "foo bar"))
    for fi7 := 0; fi7 < 12; fi7++ {
        v5 += (v5 ^ v5)
        var v8 uint32 = 205
        v8 -= (v5 & (v8 * v5))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        if (408 > 4294967295) {
            v10 := GenS1{F0: 42, F1: 480, F2: (v5 / ((v5) | 1))}
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS2{F0: v5, F1: (fi9 * (^(fi11))), F2: "hello"}
                v13 := GenS1{F0: -2147483648, F1: 255, F2: 543}
                v14 := GenS1{F0: -2147483648, F1: 1, F2: (^(v5))}
                var v15 bool = v4
                v6 = v6
                v16 := GenS3{F0: 255, F1: -1, F2: 42}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := GenS4{F0: -945, F1: 742, F2: (v5 | 838)}
            _ = v10
            _ = v17
        } else {
            v4 = false
            v18 := v1[-9223372036854775808]
            var v19 uint32 = 4294967295
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            v6 = ""
            var v21 uint = 58
            v5 -= 2147483648
            _ = fi20
            _ = v21
        }
        var v22 string = v6
        _ = fi9
        _ = v22
    }
    switch (("\n\t" + v6) + v6) {
    case "foo bar":
        for fi23 := 0; fi23 < 19; fi23++ {
            if (0 < 4294967295) {
                var v24 int16 = -32768
                _ = v24
            }
            v25 := GenS1{F0: -612, F1: 1, F2: v5}
            _ = fi23
            _ = v25
        }
    default:
        v2[-579] = (!v4)
    }
    if (!v4) {
        v5 = (v5 * (v5 + v5))
        v5 -= (v5 | v5)
        var v26 uint16 = 2
        _ = v26
    }
    if (0.0 != ((-2.25 * (0.1 / 3.14159)) + (1e10 - 0.0))) {
        var v27 int = 853
        _ = v27
    } else {
        v6 = (("foo bar" + (v6 + "\n\t")) + v6)
        v5 -= (^(v5))
    }
    switch "foo bar" {
    case "":
        v28 := v3[v6]
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
