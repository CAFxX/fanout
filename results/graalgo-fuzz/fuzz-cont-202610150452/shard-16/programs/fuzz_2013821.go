package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []int{-660, 1, 128}
    var v2 int8 = 100
    var v3 string = "a"
    var v4 float64 = 0.1
    v5 := GenS1{F0: v3}
    if false {
        fmt.Println(v2, v2)
        if (1 >= 127) {
            var v6 float64 = ((-2.25 / v4) - ((0.1 * v4) - (float64((0.1 + -2.25)) * v4)))
            v1 = append(v1, -2147483648)
            if false {
                v2 = (v2 - (v2 + v2))
            }
            _ = v6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            v1 = append(v1, fi7)
            var v8 uint32 = 0
            v4 -= ((3.14159 * 0.0) / v4)
            _ = fi7
            _ = v8
        }
    } else {
        v2 = -105
    }
    var v9 uint32 = 2147483648
    if len(v1) > 0 {
        v1[0] = 1
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v4 += float64(float32(((v4 * 1e10) * ((v4 + v4) * ((v4 - (float64(3.14159) - ((v4 - v4) * (v4 + (v4 * v4))))) * (((float64(1.5) - 0.0) * v4) - 0.1))))))
        switch fi10 {
        case 4294967295:
            fmt.Println(v4)
        case 304:
            var v11 int64 = 42
            _ = v11
        }
        v3 = v3
        _ = fi10
    }
    v2 = v2
    var v12 int = 515
    var v13 string = (("foo bar" + v3) + v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 uint32 = v9
    var v15 string = v3
    var v16 string = ((("\n\t" + (v13 + v15)) + v13) + v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
