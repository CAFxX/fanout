package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0(p0 int8, p1 uint8) uint16 {
    var v4 float64 = 0.0
    v5 := GenS1{F0: 42, F1: true}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int64 = 357
    var v2 bool = false
    var v3 uint8 = 42
    var v6 bool = v2
    switch 9223372036854775807 {
    case 32768:
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{}
            var v9 int64 = int64(v3)
            if (!v6) {
                var v10 float64 = (1.5 / -2.25)
                v11 := GenS2{F0: 0, F1: 0}
                v12 := GenS1{F0: (^(fi7)), F1: (v9 == (v9 ^ (v1 >> ((129) & 63))))}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    fmt.Println(v1, v3)
    var v13 int16 = 1000
    v2 = (!(!v2))
    v14 := []int{}
    var v15 uint8 = 255
    if v6 {
        var v16 string = "a"
        _ = v16
    } else {
        if v6 {
            fmt.Println(v15, v15, v1)
        }
        var v17 string = (("" + ((("\n\t" + "a") + "a") + "")) + ("hello" + ""))
        var v18 int16 = 1
        _ = v17
        _ = v18
    }
    var v19 uint32 = 2147483647
    var v20 int32 = 42
    v3 = v15
    v14 = append(v14, 457)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
