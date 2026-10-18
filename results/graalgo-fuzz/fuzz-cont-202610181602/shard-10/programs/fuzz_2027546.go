package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float64 = (0.0 / 1.5)
    var v2 int32 = -1
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 1
    var v5 int64 = (^(v3))
    v6 := []int{949, 32768}
    for fi7 := 0; fi7 < 13; fi7++ {
        _ = len(v6)
        _ = fi7
    }
    if (0 > (v4 + (v4 * v4))) {
        var v8 bool = false
        _ = v8
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            var v9 int32 = v2
            _ = v9
        case "a":
            v10 := GenS1{F0: 255}
            _ = v10
        }
    }
    var v11 int = 32768
    if (true || true) {
        var v12 string = ""
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v12
    } else {
        v13 := GenS1{F0: 255}
        _ = v13
    }
    v14 := GenS1{F0: 42}
    var v15 float64 = (v1 * 1e10)
    v5 = v5
    v4 = 127
    v2 += 42
    for fi16 := 0; fi16 < 1; fi16++ {
        v2 += v2
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS1{F0: (v4 << ((42) & 63))}
            var v19 int64 = (v3 * -1)
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := GenS1{F0: v4}
        _ = fi16
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        _ = len(v6)
        v6 = append(v6, 128)
        if ((v4 | (v4 / ((v4) | 1))) > 42) {
            v22 := GenS1{F0: 42}
            _ = v22
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
