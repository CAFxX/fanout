package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 uint = 65535
    var v2 float32 = ((-2.25 - -2.25) * (1.5 / ((-2.25 + 0.0) * 0.1)))
    var v3 string = ("foo bar" + "a")
    switch ("" + v3) {
    case "":
        var v4 int16 = 762
        _ = v4
    default:
        v5 := []int{32767, 127, 143, -9223372036854775808}
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v2, v3, fi6)
        if (42 < 1) {
            fmt.Println(fi6)
            v1 = (^(v1))
            switch (fi6 | (-(fi6))) {
            case 9223372036854775807:
                switch 1 {
                case 42:
                    var v7 string = v3
                    _ = v7
                }
            }
        } else {
            v3 = "foo bar"
            switch 42 {
            case 32768:
                fmt.Println(v3, v2, v2)
            }
        }
        var v8 int8 = -128
        _ = fi6
        _ = v8
    }
    v9 := []int{-9223372036854775808, 128, 2147483647, -249}
    fmt.Println(v3, v2, v2)
    switch v3 {
    case "foo bar":
        v2 = (v2 + v2)
    case "hello":
        v2 += v2
    case "":
        var v10 int8 = 100
        _ = v10
    default:
        switch 32768 {
        case 32768:
            v2 -= v2
        case 42:
            v3 = "a"
        }
    }
    v1 += 574
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint8 = 90
        _ = fi11
        _ = v12
    }
    if (!("a" < (v3 + "\n\t"))) {
        v13 := GenS1{F0: v3, F1: 1}
        var v14 int32 = -808
        if false {
            v9 = append(v9, 2147483647)
        } else {
            _ = len(v9)
            if (false || (42 < -128)) {
                if false {
                    v3 = v3
                } else {
                    v15 := GenS2{F0: -116, F1: (v1 / ((v1) | 1))}
                    v1 += v1
                    v16 := GenS3{F0: 947, F1: 65535}
                    _ = v15
                    _ = v16
                }
                var v17 int16 = 1000
                _ = v17
            }
            var v18 bool = ((v3 + ("a" + v3)) <= "")
            _ = v18
        }
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
