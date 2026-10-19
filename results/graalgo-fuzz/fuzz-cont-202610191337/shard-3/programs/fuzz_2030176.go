package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint32 = 142
    var v2 bool = false
    v3 := make(map[uint32]string)
    var v4 int64 = 9223372036854775807
    switch 518 {
    case 1:
        v1 = v1
    case 2147483647:
        if (v2 && v2) {
            v2 = v2
            v2 = v2
        } else {
            fmt.Println(v4, v4, v1)
        }
    case 128:
        v5 := GenS1{F0: ("foo bar" + "\n\t"), F1: 255, F2: 152}
        _ = v5
    default:
        var v6 uint32 = (v1 * v1)
        _ = v6
    }
    v2 = (-993 != -32768)
    var v7 int = 325
    v7 = (v7 ^ v7)
    if true {
        v8 := GenS2{F0: 890, F1: v4}
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := []int{v7}
        switch "hello" {
        case "foo bar":
            switch "" {
            case "a":
                if (-64 == 0) {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v2 = false
                        _ = fi11
                    }
                }
            default:
                var v12 bool = false
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    v4 -= 180
    v3[v1] = "a"
    var v13 int8 = 113
    v1 -= (v1 + v1)
    v13 = 60
    fmt.Println(v1, v7)
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint64 = 1
        v15 = v15
        if false {
            v16 := []int{fi14, 42}
            var v17 float64 = -2.25
            switch "hello" {
            case "hello":
                v1 += v1
            }
            _ = v16
            _ = v17
        } else {
            var v18 bool = v2
            var v19 int64 = (-(v4))
            switch 2147483647 {
            case -885:
                v2 = (1 < 42)
            case 491:
                fmt.Println(v19)
            }
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v20 := GenS3{F0: v2}
        _ = v20
    case "hello":
        v21 := []int{v7}
        _ = v21
    case "a":
        v22 := GenS3{F0: (false && v2)}
        _ = v22
    }
    if false {
        switch "hello" {
        case "a":
            fmt.Println(v1)
        case "foo bar":
            v4 = (v4 * v4)
        default:
            if (v4 > (v4 * (v4 ^ v4))) {
                v3[(v1 * v1)] = ("foo bar" + "\n\t")
                switch (v7 % (((v7 * v7)) | 1)) {
                case -2147483648:
                    v4 -= v4
                case 0:
                    v23 := GenS3{F0: (0.0 > ((float64((0.0 - 3.14159)) - 0.0) - 1.5))}
                    _ = v23
                }
                v13 = -77
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
