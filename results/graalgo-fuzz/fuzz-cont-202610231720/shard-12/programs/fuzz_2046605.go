package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    var v3 int16 = 0
    v4 := GenS1{F0: 18446744073709551615, F1: 46}
    var v5 string = (("\n\t" + (("" + "foo bar") + "a")) + ("a" + "a"))
    v6 := GenS1{F0: 18446744073709551615, F1: 135}
    var v7 string = ("" + v5)
    v8 := GenS1{F0: 0, F1: 1}
    var v9 int64 = 9223372036854775807
    fmt.Println(v2, v5, v2)
    switch 0 {
    case -147:
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := []int{2147483647}
            _ = fi10
            _ = v11
        }
    case 32768:
        v12 := GenS1{F0: 501, F1: 42}
        _ = v12
    case -2147483648:
        if v2 {
            if (-1 > 555) {
                var v13 float32 = (float32(3.14159) * (0.0 + 1.5))
                var v14 string = v7
                _ = v13
                _ = v14
            } else {
                var v15 uint16 = 255
                v16 := []int{128, 1}
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(v2)
                fmt.Println(v9, v2)
                v18 := v1[(fi17 + -411)]
                _ = fi17
                _ = v18
            }
        } else {
            v19 := []int{680}
            _ = v19
        }
    }
    var v20 int64 = (^(v9))
    v5 = v7
    var v21 uint16 = 65535
    if true {
        v3 = v3
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := []int{(fi22 * (^(fi22)))}
            _ = fi22
            _ = v23
        }
        v5 = ("" + ("a" + (("hello" + v7) + "")))
    }
    var v24 int64 = 9223372036854775807
    v25 := GenS2{F0: (v7 + "foo bar"), F1: 65535}
    switch v5 {
    case "\n\t":
        var v26 int16 = (v3 - v3)
        _ = v26
    }
    v24 = 9223372036854775807
    v27 := GenS3{F0: (v24 | v20), F1: 255}
    if (v20 >= (v24 | v24)) {
        if false {
            var v28 int = 1
            v9 -= -231
            v29 := []int{(^(v28))}
            _ = v28
            _ = v29
        }
        fmt.Println(v24)
        switch (v7 + (v5 + v5)) {
        case "\n\t":
            switch (v7 + (v5 + "hello")) {
            case "foo bar":
                fmt.Println(v20)
            case "hello":
                var v30 string = v7
                _ = v30
            }
        default:
            v9 += v20
        }
    } else {
        if (!(v2 || (!true))) {
            v3 = (v3 + (v3 << ((0) & 63)))
            if false {
                for fi31 := 0; fi31 < 3; fi31++ {
                    v9 -= 9223372036854775807
                    if (!(42 <= uint8(v21))) {
                        v32 := GenS2{F0: v7, F1: 2147483648}
                        _ = v32
                    }
                    var v33 uint16 = v21
                    _ = fi31
                    _ = v33
                }
                fmt.Println(v2, v24)
            }
        } else {
            switch "hello" {
            case "":
                switch 42 {
                case 4294967295:
                    var v34 int8 = 127
                    _ = v34
                default:
                    delete(v1, 86)
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
