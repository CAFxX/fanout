package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = float64(0.1)
    var v3 int32 = -2147483648
    var v4 int16 = -167
    var v5 int8 = 100
    if (!(-2.25 != 3.14159)) {
        v6 := v1[""]
        if (!(-128 < v5)) {
            for fi7 := 0; fi7 < 18; fi7++ {
                for fi8 := 0; fi8 < 13; fi8++ {
                    switch 1 {
                    case 643:
                        var v9 uint = 838
                        _ = v9
                    }
                    _ = fi8
                }
                _ = len(v1)
                v4 = v4
                _ = fi7
            }
        } else {
            v2 += v2
        }
        var v10 string = "\n\t"
        _ = v10
        _ = v6
    }
    v4 = v4
    v11 := GenS1{F0: (("" + "a") + "")}
    fmt.Println(v3, v4, v2)
    v12 := GenS2{F0: -1, F1: (!(!true)), F2: 520}
    for fi13 := 0; fi13 < 1; fi13++ {
        switch 9223372036854775807 {
        case -107:
            if (4294967295 > 1) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint32 = 2
                    v16 := v1["foo bar"]
                    var v17 uint = 4294967295
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            } else {
                v18 := v1[("\n\t" + ((("\n\t" + (("hello" + "hello") + "foo bar")) + "a") + "\n\t"))]
                var v19 int32 = (v3 >> ((255) & 63))
                _ = v18
                _ = v19
            }
        case 834:
            v20 := v1["foo bar"]
            _ = v20
        }
        _ = fi13
    }
    v21 := []int{128, -160, int(v3), 127}
    v5 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
