package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 bool }

func main() {
    var v1 int8 = -13
    var v2 bool = true
    var v3 float32 = (1e10 / (((1.5 + 1.5) + (0.0 * 1e10)) + 1e10))
    var v4 bool = false
    v5 := []int{323, -112, 711}
    v6 := GenS1{F0: 956}
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v4, fi7, fi7)
        fmt.Println(fi7, v1, fi7)
        _ = fi7
    }
    if (197 == 42) {
        if v4 {
            switch -472 {
            case -2147483648:
                for fi8 := 0; fi8 < 1; fi8++ {
                    _ = len(v5)
                    for fi9 := 0; fi9 < 19; fi9++ {
                        v1 += 127
                        _ = fi9
                    }
                    v5 = append(v5, (fi8 / ((fi8) | 1)))
                    _ = fi8
                }
            }
            v10 := GenS2{F0: "", F1: 0, F2: v4}
            _ = v10
        } else {
            var v11 int32 = 1
            _ = v11
        }
    }
    switch ("hello" + "\n\t") {
    case "":
        v2 = (32767 < 1)
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        if (-981 > 32767) {
            switch 9223372036854775807 {
            case -9223372036854775808:
                if v4 {
                    fmt.Println(v2, v3, v3)
                } else {
                    var v12 float64 = 0.1
                    v12 -= 0.0
                    v1 += 92
                    _ = v12
                }
            case 4294967295:
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS2{F0: "\n\t", F1: 588, F2: v2}
                    _ = fi13
                    _ = v14
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    if (127 <= 434) {
        switch (("hello" + "") + "hello") {
        case "":
            var v15 bool = (1 > -721)
            _ = v15
        }
    } else {
        fmt.Println(v4)
        v16 := GenS2{F0: "a", F1: 229, F2: ("\n\t" >= "\n\t")}
        _ = v16
    }
    var v17 uint16 = 794
    v1 += -63
    v18 := GenS1{F0: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
