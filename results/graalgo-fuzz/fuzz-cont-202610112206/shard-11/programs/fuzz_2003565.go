package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint16 = 204
    var v2 string = "hello"
    var v3 int64 = 9223372036854775807
    v4 := []string{(v2 + "hello"), v2}
    var v5 uint8 = 127
    switch 42 {
    case -255:
        v3 += v3
    default:
        for fi6 := 0; fi6 < 20; fi6++ {
            v3 = v3
            if len(v4) > 0 {
                v4[0] = v2
            }
            _ = fi6
        }
    }
    v4 = append(v4, (v2 + v2))
    fmt.Println(v5)
    if len(v4) > 0 {
        v4[0] = (v2 + v2)
    }
    if true {
        var v7 uint8 = 42
        v2 = v2
        v8 := GenS1{F0: 255}
        _ = v7
        _ = v8
    } else {
        if false {
            var v9 float32 = (1e10 - (0.1 + (float32(1.5) / -2.25)))
            _ = v9
        } else {
            v1 -= (v1 << ((65535) & 63))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v10 := GenS1{F0: v1}
            _ = v10
        }
        fmt.Println(v1, v1, v3)
        if true {
            var v11 int64 = (v3 | 1)
            _ = v11
        }
    }
    fmt.Println(v1, v3, v1)
    var v12 bool = (("hello" < "foo bar") && false)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v4 = append(v4, v2)
    switch ("foo bar" + ("foo bar" + ((((((((v2 + (v2 + v2)) + v2) + v2) + v2) + v2) + v2) + v2) + ("a" + (v2 + (v2 + v2)))))) {
    case "":
        switch -2147483648 {
        case -422:
            switch "hello" {
            case "foo bar":
                if v12 {
                    var v13 int64 = v3
                    v14 := GenS1{F0: v1}
                    var v15 bool = (((1.5 * 0.0) * 0.1) <= 1e10)
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    switch v2 {
                    case "":
                        v4 = append(v4, v2)
                    case "hello":
                        fmt.Println(v5, v2, v3)
                    case "a":
                        v16 := GenS1{F0: (v1 + (v1 / ((0) | 1)))}
                        _ = v16
                    }
                    v2 = (v2 + "")
                    for fi17 := 0; fi17 < 11; fi17++ {
                        switch fi17 {
                        case 1:
                            fmt.Println(fi17)
                        case -1:
                            var v18 uint8 = v5
                            _ = v18
                        case -532:
                            fmt.Println(v12, v5, fi17)
                        }
                        v3 -= (v3 | (v3 & v3))
                        _ = fi17
                    }
                }
            case "":
                var v19 int32 = 152
                _ = v19
            default:
                for fi20 := 0; fi20 < 14; fi20++ {
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v1 = (^(v1))
                        var v22 uint16 = 2
                        fmt.Println(v1)
                        _ = fi21
                        _ = v22
                    }
                    v4 = append(v4, v2)
                    v2 = "foo bar"
                    _ = fi20
                }
            }
        case 127:
            v3 -= 190
        default:
            if false {
                for fi23 := 0; fi23 < 5; fi23++ {
                    for fi24 := 0; fi24 < 1; fi24++ {
                        var v25 uint64 = 1
                        _ = fi24
                        _ = v25
                    }
                    _ = fi23
                }
            }
        }
    }
    switch (v2 + v2) {
    case "hello":
        for fi26 := 0; fi26 < 20; fi26++ {
            v5 -= v5
            _ = fi26
        }
    case "":
        if (!true) {
            var v27 int32 = -913
            _ = v27
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "foo bar":
        var v28 int8 = 0
        _ = v28
    }
    v12 = (((!v12) && true) || (!v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
