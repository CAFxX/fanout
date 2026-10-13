package main

import (
    "fmt"
)

func genFn0(p0 int8) int {
    fmt.Println(p0)
    if (2 != 0) {
        p0 += p0
        p0 = 0
        var v6 int16 = 490
        _ = v6
    }
    _ = p0
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    var v4 string = "hello"
    var v5 float64 = 1e10
    switch (v4 + "a") {
    case "\n\t":
        v5 = (((1.5 + (float64(0.1) - 0.1)) * (0.1 + v5)) + (v5 * (((v5 * float64(float32((float64(1e10) / v5)))) / 1e10) - (v5 - ((float64(0.0) + (float64((3.14159 - (0.1 + 1e10))) + 0.0)) / (0.0 - 3.14159))))))
    case "":
        v3 = append(v3, 2147483647)
    case "a":
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                delete(v2, 1)
                var v9 uint16 = 65535
                _ = fi8
                _ = v9
            }
            switch "" {
            case "hello":
                v10 := v1[genFn0(-1)]
                _ = v10
            case "foo bar":
                v5 = v5
            case "\n\t":
                fmt.Println(v5)
            default:
                fmt.Println(fi7)
            }
            _ = fi7
        }
    }
    v5 -= (3.14159 * (float64(1.5) - v5))
    if true {
        for fi11 := 0; fi11 < 13; fi11++ {
            v4 = ""
            for fi12 := 0; fi12 < 12; fi12++ {
                if (((!true) && false) || false) {
                    v4 = (v4 + "a")
                }
                _ = fi12
            }
            _ = fi11
        }
    } else {
        var v13 int8 = -128
        _ = v13
    }
    v5 = (((v5 + (3.14159 / 1e10)) + v5) + 0.1)
    var v14 float32 = (1.5 / 1.0)
    switch -1 {
    case -1:
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 uint64 = 511
            _ = fi15
            _ = v16
        }
    default:
        v17 := v1[491]
        _ = v17
    }
    switch ("" + ("\n\t" + v4)) {
    case "a":
        if true {
            v18 := v2[710]
            _ = v18
        }
    case "foo bar":
        switch genFn0(127) {
        case -635:
            var v19 int32 = 0
            _ = v19
        default:
            v14 = (-2.25 / (v14 - (1e10 - 0.0)))
        }
    default:
        fmt.Println(v4, v14)
    }
    var v20 int32 = 263
    var v21 float32 = (v14 * 1.5)
    fmt.Println(v14)
    var v22 int32 = (^(v20))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
