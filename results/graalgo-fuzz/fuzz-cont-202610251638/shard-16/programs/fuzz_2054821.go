package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float64 = (0.1 + 1e10)
    var v2 string = "a"
    v3 := make(map[string]bool)
    switch -834 {
    case -1:
        var v4 int8 = 127
        _ = v4
    }
    v5 := v3[(v2 + (v2 + v2))]
    switch 2147483647 {
    case -11:
        v1 = v1
    case 127:
        v2 = v2
    }
    var v6 uint16 = 17
    if v5 {
        v5 = (!v5)
        fmt.Println(v5, v5, v1)
    }
    v7 := GenS1{F0: -9223372036854775808}
    switch v2 {
    case "foo bar":
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint64 = 4294967295
            _ = fi8
            _ = v9
        }
    default:
        v2 = v2
    }
    fmt.Println(v2)
    var v10 int = 1
    var v11 uint8 = 1
    if v5 {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: (v10 & (^(fi12)))}
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        switch 1 {
        case 42:
            v3["hello"] = v5
        case 127:
            if true {
                if v5 {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v11 = v11
                        v16 := []int{(fi15 >> ((149) & 63)), (fi15 / (((v10 + fi14)) | 1)), (v10 / ((fi14) | 1)), fi15}
                        _ = fi15
                        _ = v16
                    }
                    var v17 bool = v5
                    var v18 uint = 213
                    _ = v17
                    _ = v18
                }
            }
        }
        var v19 int16 = -79
        v2 = "\n\t"
        _ = fi14
        _ = v19
    }
    switch 2147483647 {
    case 127:
        v5 = (!v5)
    case -497:
        v20 := []int{(v10 ^ v10)}
        _ = v20
    case 2147483647:
        _ = len(v3)
    default:
        v11 = v11
    }
    if (((v6 - v6) < v6) && (!true)) {
        v2 = v2
        fmt.Println(v6, v10, v6)
        v1 = v1
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint16 = 23
            v23 := GenS2{F0: 718}
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 uint64 = 0
        _ = v24
    }
    fmt.Println(v5, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
