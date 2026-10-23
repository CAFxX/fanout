package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }
func genFn0() uint16 {
    var v5 string = ("" + "")
    fmt.Println(v5)
    fmt.Println(v5)
    _ = v5
    return 65535
}

func main() {
    var v1 float64 = ((1.5 / float64(((0.1 * (((float32((3.14159 - float64(1.5))) + 0.0) * 0.0) + (1e10 - 0.0))) - -2.25))) + 0.1)
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v4 int32 = -2147483648
    v6 := GenS1{F0: 9223372036854775807}
    v3 = v3
    v7 := GenS2{F0: false, F1: 42, F2: -675}
    fmt.Println(v1, v1, v1)
    fmt.Println(v4, v3, v4)
    for fi8 := 0; fi8 < 16; fi8++ {
        switch 1 {
        case 4294967295:
            for fi9 := 0; fi9 < 1; fi9++ {
                switch (^(fi9)) {
                case 957:
                    var v10 bool = (2 >= 18446744073709551615)
                    _ = v10
                }
                if (true || true) {
                    if false {
                        v11 := GenS3{F0: 0}
                        fmt.Println(v1)
                        if false {
                            var v12 uint32 = 4294967295
                            _ = v12
                        }
                        _ = v11
                    } else {
                        if true {
                            var v13 uint = 0
                            v4 = v4
                            if false {
                                var v14 uint8 = 127
                                switch fi8 {
                                case 1:
                                    v15 := []int{fi9, fi9}
                                    _ = v15
                                default:
                                    v16 := GenS1{F0: (fi8 >> (((^(v13))) & 63))}
                                    _ = v16
                                }
                                _ = v14
                            }
                            var v17 int32 = v4
                            _ = v13
                            _ = v17
                        }
                        fmt.Println(v3)
                    }
                }
                v18 := []int{(fi8 % (((-(fi8))) | 1))}
                _ = fi9
                _ = v18
            }
        case 32767:
            v19 := []int{fi8, fi8, 9223372036854775807}
            _ = v19
        default:
            for fi20 := 0; fi20 < 4; fi20++ {
                switch v3 {
                case "\n\t":
                    v21 := GenS4{F0: (!true)}
                    _ = v21
                case "":
                    if (9223372036854775807 >= 1) {
                        fmt.Println(v1, v4)
                    }
                }
                _ = fi20
            }
        }
        v22 := v2["hello"]
        v23 := []int{4294967295, (fi8 ^ fi8)}
        _ = fi8
        _ = v22
        _ = v23
    }
    if (true && (genFn0() > 42)) {
        if true {
            var v24 float32 = 3.14159
            var v25 uint16 = 2
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 uint = 65535
                v3 = "\n\t"
                _ = fi26
                _ = v27
            }
            _ = v24
            _ = v25
        } else {
            var v28 int = 9223372036854775807
            var v29 uint = 632
            v4 -= v4
            _ = v28
            _ = v29
        }
    }
    switch v3 {
    case "hello":
        v3 = "a"
    default:
        delete(v2, ("" + ("hello" + v3)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
