package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint = 255
    var v2 int8 = 0
    var v3 float64 = (3.14159 + 1e10)
    var v4 uint8 = 227
    v5 := []int{}
    switch -439 {
    case 32767:
        if (!((v3 + 1e10) > 1e10)) {
            var v6 string = (((((("" + "a") + (((("a" + "hello") + "hello") + "hello") + "\n\t")) + (("a" + "a") + "a")) + "a") + "foo bar") + "foo bar")
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 52}
        _ = fi7
        _ = v8
    }
    var v9 int8 = (v2 + (^(v2)))
    for fi10 := 0; fi10 < 9; fi10++ {
        _ = len(v5)
        _ = fi10
    }
    v11 := GenS2{F0: 127}
    v3 = v3
    var v12 int = int(v4)
    if false {
        if true {
            v1 = v1
        }
    } else {
        v13 := GenS3{F0: 255, F1: (true || true)}
        var v14 float32 = (((((float32(3.14159) / 1.0) * 0.1) / 1e10) * (1.5 + 0.1)) - float32(v3))
        if false {
            for fi15 := 0; fi15 < 16; fi15++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v2 += 9
                _ = fi15
            }
            switch "a" {
            case "\n\t":
                fmt.Println(v2, v12)
            case "":
                var v16 bool = true
                _ = v16
            case "a":
                switch 32768 {
                case -9223372036854775808:
                    if false {
                        var v17 int8 = v9
                        v17 = 1
                        fmt.Println(v4)
                        _ = v17
                    } else {
                        var v18 int = (v12 & (v12 | (v12 % (((v12 >> ((v1) & 63))) | 1))))
                        v19 := GenS2{F0: v4}
                        _ = v18
                        _ = v19
                    }
                }
            }
            v12 = (v12 + 0)
        } else {
            v20 := GenS1{F0: 1}
            _ = v20
        }
        _ = v13
        _ = v14
    }
    if (!(-2.25 >= v3)) {
        v3 = 1e10
    } else {
        v2 = 1
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint32 = 4294967295
            _ = fi22
            _ = v23
        }
        var v24 float32 = ((((-2.25 + (0.0 * 1e10)) / 3.14159) * -2.25) * (1.5 * ((0.1 + -2.25) - float32(1.5))))
        _ = fi21
        _ = v24
    }
    v9 = (v2 % ((v2) | 1))
    var v25 string = "foo bar"
    if len(v5) > 0 {
        v5[0] = (-(v12))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
