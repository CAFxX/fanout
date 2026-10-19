package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint16; F2 int32 }

func main() {
    var v1 int32 = 42
    var v2 string = "hello"
    var v3 string = ("" + v2)
    v3 = "foo bar"
    v4 := GenS1{F0: 1, F1: v2}
    var v5 bool = false
    var v6 int32 = -35
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS2{F0: v5}
        v9 := GenS2{F0: v5}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 100
    if v5 {
        var v11 int8 = 100
        v3 = ""
        var v12 float32 = (0.1 / 1e10)
        _ = v11
        _ = v12
    }
    v13 := []int{0}
    switch 678 {
    case 32767:
        v1 += (v6 * 42)
    case 127:
        var v14 float32 = 3.14159
        _ = v14
    case -1:
        var v15 int16 = -383
        _ = v15
    }
    fmt.Println(v2)
    var v16 uint = 2
    var v17 int32 = v1
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            switch -1 {
            case 4294967295:
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            case 988:
                v20 := GenS3{F0: (9223372036854775807 >= 9223372036854775807), F1: 880, F2: -838}
                _ = v20
            case 128:
                var v21 int64 = -452
                _ = v21
            }
            fmt.Println(v1, v6)
            _ = fi19
        }
        v13 = append(v13, fi18)
        var v22 string = ((("a" + "\n\t") + v3) + (v2 + (v2 + v3)))
        _ = fi18
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v3 = v3
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 float64 = (1e10 - 3.14159)
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    if len(v13) > 0 {
        v13[0] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
