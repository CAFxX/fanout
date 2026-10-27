package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint8) uint16 {
    var v5 float64 = 1.5
    v5 += v5
    _ = p0
    _ = p1
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 931
    v3 := make(map[int]string)
    var v4 bool = true
    v4 = (!v4)
    switch 2147483647 {
    case 128:
        switch -9223372036854775808 {
        case 32768:
            var v6 string = ""
            _ = v6
        default:
            v2 -= (v2 * v2)
        }
    case 855:
        v7 := GenS1{F0: 127}
        _ = v7
    case -2147483648:
        switch "" {
        case "a":
            v2 -= (v2 + (v2 << ((42) & 63)))
        case "\n\t":
            v8 := GenS1{F0: 87}
            _ = v8
        default:
            var v9 int16 = 1
            _ = v9
        }
    }
    var v10 uint16 = 32767
    v2 = 1
    var v11 uint8 = 127
    fmt.Println(v11, v11, v10)
    v12 := v1[""]
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint32 = 232
            var v16 uint16 = (v10 & 2)
            switch ("" + ("" + "a")) {
            case "foo bar":
                v17 := GenS1{F0: (v11 ^ v11)}
                _ = v17
            case "":
                var v18 uint = 764
                _ = v18
            case "a":
                var v19 uint64 = 4294967295
                _ = v19
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        switch 0 {
        case -685:
            fmt.Println(v2, v4)
        case 127:
            var v20 uint64 = 0
            _ = v20
        }
        v21 := GenS1{F0: v11}
        _ = fi13
        _ = v21
    }
    var v22 uint8 = (v11 | (v11 >> ((255) & 63)))
    v4 = ("foo bar" < "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
