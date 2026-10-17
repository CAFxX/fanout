package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int32; F2 int16 }

func main() {
    var v1 uint = 779
    var v2 string = "foo bar"
    v3 := []int{32768, 74, -827, -9223372036854775808}
    switch -701 {
    case 0:
        fmt.Println(v1)
    }
    v4 := GenS1{F0: 426}
    v3 = append(v3, -2147483648)
    v2 = v2
    var v5 float64 = 1e10
    v6 := GenS1{F0: -1}
    _ = len(v3)
    if true {
        if len(v3) > 0 {
            v3[0] = 834
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            v3 = append(v3, fi7)
            switch (fi7 ^ (fi7 >> ((153) & 63))) {
            case 32767:
                v8 := GenS2{F0: 92}
                _ = v8
            case -736:
                v9 := GenS2{F0: 127}
                _ = v9
            }
            for fi10 := 0; fi10 < 17; fi10++ {
                fmt.Println(v5)
                v2 = "foo bar"
                v11 := GenS3{F0: 109, F1: -1, F2: 1000}
                _ = fi10
                _ = v11
            }
            _ = fi7
        }
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 string = "a"
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 string = ("\n\t" + "a")
        var v16 uint = (v1 / ((65535) | 1))
        var v17 uint = 42
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
