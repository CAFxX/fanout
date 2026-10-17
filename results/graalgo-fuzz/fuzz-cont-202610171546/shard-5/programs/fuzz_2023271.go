package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int64 = 0
    var v2 uint16 = 538
    var v3 int64 = v1
    v2 = (v2 / (((v2 & (^(v2)))) | 1))
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 string = ""
        _ = fi4
        _ = v5
    }
    if (109 < 69) {
        fmt.Println(v2, v2)
        v3 += v1
        v2 = v2
    } else {
        v6 := GenS1{F0: v2}
        _ = v6
    }
    v7 := GenS1{F0: (^(v2))}
    v8 := []int{4294967295, 0}
    switch "a" {
    case "hello":
        if (((("a" + (((("hello" + ("foo bar" + "hello")) + "") + "\n\t") + "\n\t")) + "\n\t") > ("a" + "hello")) || true) {
            v8 = append(v8, -9223372036854775808)
            var v9 uint32 = 55
            if false {
                var v10 int32 = 200
                _ = v10
            }
            _ = v9
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v11 int8 = 100
            for fi12 := 0; fi12 < 20; fi12++ {
                v11 = (v11 << ((4294967295) & 63))
                v11 = (v11 >> ((1) & 63))
                switch fi12 {
                case 32767:
                    v1 = (v1 ^ v3)
                case -371:
                    var v13 int32 = 883
                    _ = v13
                case 1:
                    for fi14 := 0; fi14 < 10; fi14++ {
                        if (("foo bar" == "\n\t") || true) {
                            v8 = append(v8, (fi12 + (fi12 ^ fi14)))
                            v15 := GenS1{F0: v2}
                            v16 := GenS2{F0: (v1 ^ (v1 * v1))}
                            _ = v15
                            _ = v16
                        } else {
                            switch ("a" + "") {
                            case "a":
                                v2 += 65535
                            case "\n\t":
                                v11 -= (v11 * (v11 >> ((671) & 63)))
                            default:
                                var v17 string = "foo bar"
                                _ = v17
                            }
                        }
                        _ = fi14
                    }
                default:
                    v1 = v3
                }
                _ = fi12
            }
            _ = v11
        }
    case "":
        var v18 uint8 = 255
        _ = v18
    case "\n\t":
        var v19 uint64 = 1
        _ = v19
    }
    if true {
        var v20 int16 = -379
        _ = v20
    }
    var v21 uint64 = 18446744073709551615
    v1 += v1
    v22 := GenS2{F0: 1}
    v1 = v1
    fmt.Println(v1, v1, v3)
    for fi23 := 0; fi23 < 17; fi23++ {
        v1 += (v3 % (((v1 ^ v1)) | 1))
        _ = fi23
    }
    v24 := GenS2{F0: 42}
    var v25 bool = (((0.1 - 3.14159) < 1.5) || true)
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint64 = (v21 - (v21 / ((v21) | 1)))
        var v28 int64 = -32
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v7
    _ = v8
}
