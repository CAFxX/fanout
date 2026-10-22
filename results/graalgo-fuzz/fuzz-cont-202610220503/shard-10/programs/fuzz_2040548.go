package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 string = v1
    v3 := []string{"a", v1}
    v4 := make(map[uint32]bool)
    if (3.14159 != (-2.25 + float32(((1e10 + 1.5) / 1.5)))) {
        v2 = v1
        var v5 uint16 = 1
        var v6 uint8 = 25
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 901, F1: 0}
    switch v2 {
    case "foo bar":
        if (((((((1e10 + (-2.25 + 3.14159)) * ((0.0 - ((((-2.25 + (1e10 * ((float64(float32(3.14159)) + 3.14159) + 3.14159))) - float64(float32(1e10))) * (3.14159 - float64(1.5))) * 0.1)) * float64(3.14159))) + 0.1) * float64(1e10)) - (3.14159 * (1e10 - 1.5))) < 0.0) && true) {
            switch "a" {
            case "a":
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 uint16 = 32767
                    _ = fi8
                    _ = v9
                }
            case "hello":
                if false {
                    var v10 float32 = ((((0.1 + (1e10 / 1.0)) / 1.0) - 0.1) + 0.0)
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 float64 = 3.14159
                        v4[701] = true
                        _ = fi11
                        _ = v12
                    }
                    if false {
                        var v13 int = 2147483647
                        v14 := GenS2{F0: 0, F1: 100, F2: 411}
                        v15 := v4[2147483647]
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v16 := GenS2{F0: 383, F1: 100, F2: 18446744073709551615}
                        if len(v3) > 0 {
                            v3[0] = (("" + ((("hello" + ("a" + "hello")) + v1) + v2)) + v2)
                        }
                        _ = v16
                    }
                    _ = v10
                }
            case "foo bar":
                var v17 uint16 = 42
                _ = v17
            default:
                switch ("\n\t" + v1) {
                case "\n\t":
                    v4[65535] = ((42 == 0) && false)
                }
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                switch fi18 {
                case 127:
                    var v19 uint8 = 128
                    _ = v19
                }
                var v20 float64 = -2.25
                _ = fi18
                _ = v20
            }
        }
    case "a":
        v2 = ("hello" + ("a" + "\n\t"))
    case "hello":
        v2 = ("a" + (("" + (v2 + v1)) + v2))
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    v2 = v2
    var v21 int = 0
    if (100 <= 100) {
        var v22 uint16 = 482
        _ = v22
    }
    if (1 < 0) {
        if true {
            delete(v4, 2147483648)
            switch v21 {
            case 2147483647:
                v21 += (v21 % ((v21) | 1))
            case 395:
                var v23 uint = 18446744073709551615
                _ = v23
            case 42:
                v24 := GenS1{F0: 42, F1: 1}
                _ = v24
            }
            v25 := GenS3{F0: 65535, F1: 2}
            _ = v25
        }
    } else {
        if false {
            var v26 uint16 = 1
            _ = v26
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            v3 = append(v3, ((v2 + v1) + v1))
            _ = fi27
        }
        v21 -= -2147483648
    }
    var v28 uint32 = 65535
    var v29 int16 = -1
    var v30 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}
