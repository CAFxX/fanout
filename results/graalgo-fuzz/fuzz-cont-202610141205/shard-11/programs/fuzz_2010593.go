package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool; F1 string; F2 uint64 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []int{127, -2147483648, 2147483647}
    var v2 string = ("foo bar" + ("\n\t" + "\n\t"))
    var v3 int16 = 570
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    for fi6 := 0; fi6 < 17; fi6++ {
        switch 665 {
        case 42:
            var v7 int16 = v3
            _ = v7
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            v1 = append(v1, (fi6 >> ((2) & 63)))
            _ = len(v1)
            if false {
                var v9 string = ("a" + v2)
                _ = v9
            }
            _ = fi8
        }
        if len(v1) > 0 {
            v1[0] = 127
        }
        _ = fi6
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        switch ((((v2 + v2) + ((v2 + v2) + ("a" + (v2 + (("a" + (((v2 + v2) + v2) + "")) + v2))))) + v2) + (v2 + v2)) {
        case "a":
            v4[870] = "foo bar"
        case "hello":
            v11 := GenS1{F0: (1e10 >= 0.1), F1: 458}
            _ = v11
        default:
            var v12 float64 = 3.14159
            _ = v12
        }
        if (v2 >= v2) {
            fmt.Println(v2)
        } else {
            v2 = "a"
        }
        _ = fi10
    }
    if (((v2 + v2) + (v2 + (v2 + v2))) < (((v2 + v2) + "foo bar") + "a")) {
        v4[633] = "hello"
        v13 := GenS2{F0: 32768, F1: 32767}
        v14 := GenS1{F0: true, F1: 875}
        _ = v13
        _ = v14
    }
    if false {
        if false {
            for fi15 := 0; fi15 < 8; fi15++ {
                v3 += (v3 % ((v3) | 1))
                switch fi15 {
                case -2147483648:
                    fmt.Println(v3)
                case -9223372036854775808:
                    fmt.Println(v2)
                }
                _ = fi15
            }
            var v16 float64 = (0.0 / -2.25)
            _ = len(v5)
            _ = v16
        }
    } else {
        v17 := GenS3{F0: true, F1: ("\n\t" + ("foo bar" + "a")), F2: 4294967295}
        _ = v17
    }
    fmt.Println(v3, v3)
    v3 += (^(v3))
    if false {
        var v18 int32 = 91
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint16 = 1
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int32 = v18
            fmt.Println(v2, v22)
            _ = fi21
            _ = v22
        }
        _ = v18
    } else {
        v3 = v3
        v23 := GenS4{F0: 4294967295}
        _ = v23
    }
    v24 := GenS1{F0: true, F1: 744}
    _ = len(v5)
    fmt.Println(v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
