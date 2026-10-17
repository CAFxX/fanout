package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 uint32) int16 {
    fmt.Println(p1, p1)
    if false {
        if (!false) {
            var v4 int32 = 2147483647
            if ((1.5 / -2.25) != 1e10) {
                p0 = (p0 / (((p0 >> ((18446744073709551615) & 63))) | 1))
                for fi5 := 0; fi5 < 2; fi5++ {
                    p1 = 505
                    var v6 int16 = 0
                    _ = fi5
                    _ = v6
                }
            }
            var v7 float32 = ((-2.25 - ((1.5 - -2.25) / (1.5 - -2.25))) - 3.14159)
            _ = v4
            _ = v7
        }
    } else {
        p0 = 127
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        p0 += p0
        v9 := GenS1{F0: (p1 - p1), F1: 486, F2: "a"}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (p1 - (^(p1))), F1: 914, F2: ((("" + "hello") + "a") + "a")}
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    v1 := []uint32{1}
    var v2 int16 = -997
    v3 := []uint32{42, 65535}
    v11 := GenS1{F0: 336, F1: 0, F2: (("\n\t" + "\n\t") + "\n\t")}
    v12 := GenS1{F0: 0, F1: -1, F2: "hello"}
    for fi13 := 0; fi13 < 13; fi13++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(fi14, v2)
            switch "foo bar" {
            case "":
                var v16 string = (("\n\t" + "\n\t") + "")
                _ = v16
            default:
                switch fi14 {
                case 32767:
                    var v17 float64 = 1.5
                    _ = v17
                case 32768:
                    v18 := GenS1{F0: 942, F1: 0, F2: "a"}
                    _ = v18
                case -588:
                    fmt.Println(fi14, v2)
                }
            }
            if (true || (611 != 2)) {
                var v19 int16 = genFn0((fi14 ^ fi14), 42)
                _ = v19
            } else {
                _ = len(v3)
                var v20 int64 = int64(fi15)
                _ = v20
            }
            _ = fi15
        }
        _ = fi14
    }
    v21 := GenS2{F0: 255}
    var v22 float32 = -2.25
    v22 = v22
    v2 += 1
    v23 := GenS2{F0: 210}
    v1 = append(v1, 12)
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 float64 = 3.14159
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
