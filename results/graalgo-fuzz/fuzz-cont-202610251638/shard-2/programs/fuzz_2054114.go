package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0(p0 string) uint16 {
    v5 := GenS1{F0: -1, F1: -128, F2: -316}
    var v6 uint8 = 141
    var v7 float32 = 0.1
    v7 = ((v7 - v7) / (((float32(1e10) / (((v7 * (v7 + float32(3.14159))) / v7) / (v7 - 3.14159))) * v7) / float32(((0.1 + 0.0) / 3.14159))))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 uint32 = 0
    v2 := []int{127, -393, -1, 32767}
    var v3 uint8 = 127
    var v4 uint = 873
    if (-2.25 <= ((0.0 + -2.25) * 1e10)) {
        v1 = 2147483648
    }
    switch "" {
    case "":
        v3 += uint8(v1)
    default:
        switch ("foo bar" + "\n\t") {
        case "":
            var v8 int8 = -32
            _ = v8
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "foo bar":
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS2{F0: 0, F1: 65535, F2: (v3 & (v3 & (v3 | (^(v3)))))}
                _ = fi9
                _ = v10
            }
        }
    }
    var v11 string = "hello"
    v12 := GenS1{F0: 1, F1: 100, F2: -9223372036854775808}
    switch v11 {
    case "\n\t":
        fmt.Println(v3, v11)
    case "foo bar":
        if (v1 <= (v1 & v1)) {
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint16 = 0
                if (false && false) {
                    v11 = (v11 + v11)
                } else {
                    if false {
                        v2 = append(v2, fi13)
                    }
                    var v15 float32 = 3.14159
                    v16 := GenS2{F0: (v14 + 255), F1: v1, F2: 0}
                    _ = v15
                    _ = v16
                }
                for fi17 := 0; fi17 < 8; fi17++ {
                    v4 -= (v4 / ((uint(v1)) | 1))
                    v18 := GenS3{F0: 127, F1: 100}
                    _ = len(v2)
                    _ = fi17
                    _ = v18
                }
                _ = fi13
                _ = v14
            }
            fmt.Println(v1, v1, v1)
        }
    }
    var v19 int8 = 0
    var v20 uint8 = v3
    var v21 string = (v11 + v11)
    v20 = v20
    v2 = append(v2, -1)
    v22 := GenS2{F0: 65535, F1: 2, F2: 2}
    v11 = v11
    v23 := GenS2{F0: 65535, F1: (v1 ^ uint32(v20)), F2: v3}
    v24 := GenS3{F0: (v19 + 0), F1: v19}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
