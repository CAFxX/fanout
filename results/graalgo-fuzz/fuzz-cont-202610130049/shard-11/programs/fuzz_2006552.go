package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int8 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0(p0 uint64) bool {
    switch "foo bar" {
    case "a":
        v6 := GenS1{F0: 1, F1: -28}
        _ = v6
    case "":
        p0 = (^(p0))
    case "\n\t":
        v7 := GenS2{F0: p0, F1: 42, F2: 80}
        _ = v7
    }
    v8 := []int{32767}
    _ = p0
    _ = v8
    return (1 >= 65535)
}

func main() {
    var v1 string = "hello"
    v2 := []int{415, 127, -2147483648}
    var v3 uint8 = 0
    var v4 uint32 = 4294967295
    var v5 int32 = -2147483648
    v9 := GenS1{F0: 9223372036854775807, F1: -1}
    v10 := GenS1{F0: 47, F1: 1000}
    var v11 int32 = v5
    var v12 string = ""
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            v12 = (v1 + "")
            switch (v1 + ((v1 + (v12 + ((("" + v12) + (v12 + (("foo bar" + (v12 + (("\n\t" + ("\n\t" + ("\n\t" + ""))) + (("\n\t" + "a") + (v1 + "hello"))))) + ((v1 + ((v12 + v12) + ((v12 + v12) + "\n\t"))) + v12)))) + (v1 + "hello")))) + (v12 + v12))) {
            case "foo bar":
                fmt.Println(v3)
            case "\n\t":
                switch 9223372036854775807 {
                case 1:
                    fmt.Println(v12)
                case 128:
                    if (false && false) {
                        fmt.Println(fi13, v4)
                        fmt.Println(fi13, v11, v4)
                    } else {
                        fmt.Println(v11)
                        v1 = v12
                        var v15 bool = (-152 == 9223372036854775807)
                        _ = v15
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if true {
                        switch v12 {
                        case "a":
                            v11 -= (v11 << ((65535) & 63))
                        case "hello":
                            fmt.Println(v1)
                        case "":
                            if len(v2) > 0 {
                                v2[0] = (fi13 ^ (fi13 >> ((4294967295) & 63)))
                            }
                        default:
                            fmt.Println(fi13)
                        }
                        switch v12 {
                        case "foo bar":
                            var v16 uint = 42
                            _ = v16
                        }
                    }
                case 32767:
                    if false {
                        var v17 float32 = 0.1
                        _ = v17
                    }
                }
            default:
                switch (fi13 >> ((943) & 63)) {
                case 0:
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch fi13 {
                        case 4294967295:
                            _ = len(v2)
                        default:
                            fmt.Println(v3)
                        }
                    } else {
                        v1 = v1
                        fmt.Println(v5, v5, v11)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                default:
                    var v18 string = (v12 + v1)
                    _ = v18
                }
            }
            var v19 int32 = (v11 | (v5 % ((334) | 1)))
            _ = fi14
            _ = v19
        }
        _ = fi13
    }
    switch 2147483647 {
    case -1:
        fmt.Println(v11, v1)
    default:
        _ = len(v2)
    }
    v20 := GenS3{F0: (^(v5)), F1: -883}
    v21 := GenS2{F0: 4294967295, F1: 113, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
