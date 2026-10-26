package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 int32, p1 uint64) uint16 {
    var v4 int8 = 42
    var v5 int32 = p0
    fmt.Println(p0, p1, p1)
    v6 := GenS1{F0: "a"}
    v7 := GenS1{F0: ("a" + "")}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int, p1 int) string {
    p0 += -2147483648
    var v8 uint8 = 42
    _ = p0
    _ = p1
    _ = v8
    return (("a" + "foo bar") + "foo bar")
}

func main() {
    var v1 uint16 = 510
    var v2 float32 = 0.0
    v3 := make(map[uint32]string)
    v3[2] = "\n\t"
    v9 := v3[2147483648]
    v9 = ((v9 + v9) + ((v9 + v9) + "a"))
    if true {
        switch v9 {
        case "\n\t":
            for fi10 := 0; fi10 < 20; fi10++ {
                switch "\n\t" {
                case "\n\t":
                    var v11 int64 = 239
                    _ = v11
                case "a":
                    for fi12 := 0; fi12 < 20; fi12++ {
                        switch (fi10 % (((fi10 - fi12)) | 1)) {
                        case -9223372036854775808:
                            v13 := GenS1{F0: ("a" + ("hello" + ((v9 + v9) + ("foo bar" + v9))))}
                            _ = v13
                        case 372:
                            var v14 int32 = 42
                            _ = v14
                        case 1:
                            fmt.Println(fi10, v2)
                        }
                        _ = fi12
                    }
                case "hello":
                    fmt.Println(v2, v9, fi10)
                default:
                    fmt.Println(v1, v9)
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 uint = 2
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 12; fi17++ {
                    switch ("hello" + ((v9 + v9) + v9)) {
                    case "a":
                        v2 = 1e10
                    case "\n\t":
                        v18 := GenS1{F0: ("a" + v9)}
                        _ = v18
                    case "":
                        fmt.Println(v1, v9)
                    default:
                        var v19 float64 = ((1e10 / 1.0) * (1.5 * -2.25))
                        _ = v19
                    }
                    fmt.Println(v1, v1, v1)
                    v20 := GenS2{F0: 948, F1: -2147483648, F2: 42}
                    _ = fi17
                    _ = v20
                }
                _ = fi10
            }
        case "":
            v1 = (v1 + 470)
        default:
            if true {
                var v21 uint64 = 18446744073709551615
                _ = v21
            } else {
                if true {
                    switch 0 {
                    case -1:
                        var v22 int16 = 1000
                        _ = v22
                    case 0:
                        fmt.Println(v1)
                    case -777:
                        fmt.Println(v9)
                    default:
                        v23 := []int{707, 218, -567, 128}
                        _ = v23
                    }
                }
            }
        }
    }
    v24 := []int{525, 4294967295}
    if true {
        for fi25 := 0; fi25 < 16; fi25++ {
            fmt.Println(v2, v9, fi25)
            _ = fi25
        }
        v26 := GenS3{F0: 4294967295, F1: 716, F2: 42}
        v27 := GenS1{F0: (((v9 + (v9 + v9)) + v9) + v9)}
        _ = v26
        _ = v27
    }
    switch 453 {
    case -1:
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    case 9223372036854775807:
        v28 := GenS2{F0: 916, F1: 760, F2: 100}
        _ = v28
    }
    if (255 < 255) {
        fmt.Println(v9)
    }
    v29 := GenS2{F0: 18446744073709551615, F1: -2147483648, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v9
}
