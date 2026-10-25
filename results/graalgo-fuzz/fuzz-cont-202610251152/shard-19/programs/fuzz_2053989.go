package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    v1 := []int{-466, 127, -2147483648, 42}
    var v2 int16 = 1
    var v3 int64 = 271
    var v4 uint32 = 500
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 float32 = 1.5
        for fi7 := 0; fi7 < 18; fi7++ {
            v1 = append(v1, (fi7 - fi7))
            fmt.Println(fi5)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    v4 += 677
    v8 := GenS1{F0: 2147483647}
    v1 = append(v1, 181)
    if true {
        if (1 != 65535) {
            var v9 uint64 = 42
            switch int(v3) {
            case -2147483648:
                if true {
                    fmt.Println(v3, v4, v2)
                    for fi10 := 0; fi10 < 17; fi10++ {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            var v12 bool = (3.14159 == 3.14159)
                            _ = fi11
                            _ = v12
                        }
                        _ = fi10
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            var v13 int16 = (-(v2))
            _ = v13
            _ = v9
        } else {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
        var v14 string = "\n\t"
        var v15 int64 = (v3 ^ -1)
        _ = v14
        _ = v15
    } else {
        v4 += (v4 & (v4 * 105))
        switch -2147483648 {
        case 128:
            v4 = (v4 ^ v4)
        }
        var v16 int32 = 2147483647
        _ = v16
    }
    v3 = v3
    v2 = -32768
    v1 = append(v1, 127)
    v4 = 65535
    fmt.Println(v4, v2, v3)
    if (!(2147483647 > 903)) {
        v2 -= (^(v2))
        switch "" {
        case "\n\t":
            var v17 int8 = 100
            _ = v17
        case "":
            v2 += v2
        }
        var v18 int = 32768
        _ = v18
    } else {
        if (976 >= 813) {
            v1 = append(v1, 0)
            if (false || true) {
                v19 := GenS2{F0: false, F1: 100}
                var v20 uint = 181
                _ = v19
                _ = v20
            }
        }
    }
    fmt.Println(v4, v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
