package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16; F1 int16; F2 int8 }
type GenS3 struct { F0 int64; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("foo bar" + ("" + "hello"))}
    var v3 float64 = -2.25
    fmt.Println(v3)
    v3 = v3
    v3 += v3
    switch 9223372036854775807 {
    case 1:
        _ = len(v2)
    case 128:
        var v4 int64 = -1
        _ = v4
    case 32768:
        v3 -= (float64(float32(v3)) + ((((v3 + (float64(3.14159) - (v3 - 0.1))) / v3) + float64(0.0)) + float64(0.1)))
    default:
        if true {
            v5 := GenS1{F0: -32768, F1: 39}
            switch 9223372036854775807 {
            case 32767:
                var v6 uint32 = 2147483648
                _ = v6
            case -2147483648:
                v7 := GenS1{F0: 9, F1: -1}
                _ = v7
            case -67:
                if len(v2) > 0 {
                    v2[0] = (("hello" + "a") + "\n\t")
                }
            }
            _ = v5
        } else {
            switch ("" + "\n\t") {
            case "a":
                for fi8 := 0; fi8 < 3; fi8++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v9 uint32 = 2147483647
                    _ = fi8
                    _ = v9
                }
            default:
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint16 = 0
                    v12 := GenS2{F0: 299, F1: 42, F2: -1}
                    delete(v1, (-(fi10)))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            switch ("foo bar" + "") {
            case "":
                _ = len(v1)
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint = 894
                v3 = (0.0 + v3)
                _ = fi13
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(fi15, v3)
        _ = fi15
    }
    v16 := GenS3{F0: -196, F1: 0, F2: 32767}
    v3 += v3
    if true {
        for fi17 := 0; fi17 < 3; fi17++ {
            v3 = (v3 - (1.5 + 3.14159))
            v18 := GenS3{F0: 212, F1: 65535, F2: -32768}
            _ = fi17
            _ = v18
        }
        if (true || true) {
            v3 -= v3
            fmt.Println(v3)
        }
    }
    var v19 string = ("hello" + "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
