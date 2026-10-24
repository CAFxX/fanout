package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int = 1
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int16 = -236
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            if ((-349 == 32767) || false) {
                v1 += (fi7 % ((fi8) | 1))
                if (!false) {
                    var v9 uint32 = 357
                    _ = v9
                }
                var v10 uint = 255
                _ = v10
            } else {
                fmt.Println(v1, fi8, fi7)
                v11 := GenS1{F0: 180}
                v1 = 42
                v12 := GenS1{F0: 1}
                _ = v11
                _ = v12
            }
            _ = fi8
        }
        v13 := GenS1{F0: 839}
        fmt.Println(fi7)
        _ = fi7
        _ = v13
    }
    switch ((("\n\t" + (("\n\t" + "foo bar") + "")) + "\n\t") + (("\n\t" + "\n\t") + "hello")) {
    case "foo bar":
        fmt.Println(v1)
    case "":
        v14 := []int{128}
        _ = v14
    }
    switch (-(v1)) {
    case 42:
        var v15 string = "\n\t"
        _ = v15
    case 2147483647:
        v1 = v1
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := v3["foo bar"]
            v18 := v2["a"]
            if (v17 && v17) {
                if (true || false) {
                    fmt.Println(v17, v18)
                    v1 = (v1 - -390)
                }
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if (100 <= 100) {
        v19 := GenS1{F0: 32767}
        _ = v19
    }
    v20 := GenS1{F0: 255}
    v1 = 127
    fmt.Println(v1)
    var v21 float32 = (3.14159 * (3.14159 * 1e10))
    v1 += (v1 * v1)
    v1 -= v1
    var v22 float32 = v21
    if (390 == -2147483648) {
        if false {
            var v23 uint16 = 42
            v24 := GenS1{F0: 255}
            _ = v23
            _ = v24
        }
        v1 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
