package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint32; F1 uint8; F2 string }
func genFn0(p0 uint, p1 uint64) int8 {
    v7 := GenS1{F0: "hello", F1: false}
    p1 = p1
    p1 = p1
    v8 := GenS2{F0: -898, F1: 42}
    switch 127 {
    case -2147483648:
        p0 = p0
    case 42:
        p0 -= 65535
    case 702:
        var v9 bool = true
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 10
}

func main() {
    var v1 uint32 = 272
    var v2 float32 = (1.5 / 1.5)
    v3 := make(map[uint32]int)
    var v4 int = 32768
    var v5 uint32 = v1
    var v6 uint8 = 1
    v1 = v1
    v2 -= ((((float32((float64((1.5 / 1.5)) + -2.25)) + (((v2 / float32(1e10)) + 0.0) / v2)) / float32(0.1)) + (((v2 / v2) * float32(((0.1 * (0.1 + (((((((float64((v2 - (((v2 * (v2 * float32(1.5))) + (v2 * v2)) / (0.1 / -2.25)))) - 1e10) - float64(1.5)) * -2.25) + 1.5) - (0.1 - 1e10)) / 1e10) - ((0.0 * 0.1) / 0.1)))) * (1.5 - (0.0 * 3.14159))))) * v2)) + float32((3.14159 + (float64(v2) + (-2.25 / float64(((v2 / 0.1) + v2)))))))
    switch 904 {
    case 0:
        for fi10 := 0; fi10 < 2; fi10++ {
            delete(v3, 2)
            _ = fi10
        }
    }
    v11 := []int{(v4 / ((-1) | 1)), (-(v4)), v4, v4}
    if (!((v4 ^ (v4 ^ 433)) > (^(v4)))) {
        for fi12 := 0; fi12 < 2; fi12++ {
            _ = len(v3)
            _ = fi12
        }
        switch 9223372036854775807 {
        case 32768:
            if ((-1 > -1) || true) {
                v13 := GenS1{F0: "\n\t", F1: true}
                _ = v13
            } else {
                if true {
                    v14 := GenS3{F0: 42, F1: (v4 / (((v4 * v4)) | 1))}
                    _ = v14
                } else {
                    v2 += (v2 + float32(float64(3.14159)))
                    v15 := GenS4{F0: 2, F1: (v6 + (v6 / ((v6) | 1))), F2: ("\n\t" + "foo bar")}
                    _ = v15
                }
                switch "\n\t" {
                case "foo bar":
                    var v16 uint8 = 255
                    _ = v16
                default:
                    v17 := GenS1{F0: "foo bar", F1: true}
                    _ = v17
                }
            }
        case -1:
            var v18 int64 = -667
            _ = v18
        case 32767:
            switch ("\n\t" + "") {
            case "a":
                var v19 int32 = 2147483647
                _ = v19
            default:
                if true {
                    var v20 int32 = 42
                    for fi21 := 0; fi21 < 4; fi21++ {
                        var v22 uint = 42
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
            }
        }
    }
    v1 -= v5
    switch "hello" {
    case "":
        var v23 int16 = -1
        _ = v23
    case "foo bar":
        v1 += v1
    case "a":
        var v24 string = "hello"
        _ = v24
    default:
        var v25 int = 42
        _ = v25
    }
    fmt.Println(v1, v4, v1)
    v5 -= (v1 % (((v1 ^ v1)) | 1))
    v6 = 0
    var v26 uint8 = 127
    var v27 uint16 = 65535
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
