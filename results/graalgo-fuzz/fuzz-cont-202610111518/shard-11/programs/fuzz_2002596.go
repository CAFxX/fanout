package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float32 = 0.1
    var v2 string = "foo bar"
    v3 := make(map[int]string)
    var v4 float32 = 1e10
    var v5 uint16 = 634
    var v6 int16 = 623
    v6 = v6
    v7 := v3[128]
    if (float64(((1.5 * 1.5) - -2.25)) < 0.0) {
        fmt.Println(v1)
    }
    fmt.Println(v4, v2)
    var v8 int8 = 100
    var v9 float64 = 0.1
    for fi10 := 0; fi10 < 9; fi10++ {
        v4 += v1
        var v11 int8 = v8
        for fi12 := 0; fi12 < 7; fi12++ {
            v4 = float32(v9)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(v9, v2, v7)
        switch (fi13 << ((2) & 63)) {
        case 32768:
            delete(v3, (fi13 - -9223372036854775808))
        case -977:
            for fi14 := 0; fi14 < 18; fi14++ {
                v9 = ((-2.25 / v9) - (-2.25 - (((((((1e10 * 1.5) * v9) / -2.25) / (0.0 - ((v9 / float64(v4)) / v9))) / (v9 / v9)) * v9) * (v9 - (((v9 - v9) / v9) - (v9 / 1e10))))))
                v6 = (v6 | (v6 >> ((0) & 63)))
                fmt.Println(v2, v2, v9)
                _ = fi14
            }
        }
        if ((!((!(!(127 == 255))) && (53 >= 83))) || (127 != 105)) {
            v7 = ""
            v15 := GenS1{F0: v6}
            var v16 float32 = ((v1 * 3.14159) - 0.0)
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            v9 = (v9 * (v9 + v9))
            v19 := []int{(-(fi18))}
            v5 -= v5
            _ = fi18
            _ = v19
        }
        var v20 int64 = 0
        v21 := []int{(fi17 << ((530) & 63)), 420, 1}
        _ = fi17
        _ = v20
        _ = v21
    }
    v22 := []int{}
    v23 := GenS1{F0: (v6 << ((347) & 63))}
    v7 = (v2 + v2)
    v4 = float32(float64(((v1 / ((((v1 + v1) / 1.0) + 1e10) / 1e10)) / (((v4 - 0.1) * 0.0) * (v4 + ((((v1 + 3.14159) * 0.1) + v1) * 0.1))))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
