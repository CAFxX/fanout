package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int }

func main() {
    var v1 string = ("foo bar" + (("a" + ("\n\t" + "foo bar")) + ("foo bar" + ("\n\t" + "foo bar"))))
    var v2 uint16 = 42
    var v3 uint32 = 488
    var v4 uint = 641
    v5 := []uint32{458, v3, 839}
    var v6 uint16 = v2
    v6 = v2
    fmt.Println(v1, v3, v1)
    switch v1 {
    case "foo bar":
        if true {
            var v7 float32 = 3.14159
            _ = v7
        }
    case "hello":
        if ((0 > -797) && false) {
            if len(v5) > 0 {
                v5[0] = v3
            }
            for fi8 := 0; fi8 < 3; fi8++ {
                for fi9 := 0; fi9 < 10; fi9++ {
                    v5 = append(v5, (v3 >> ((v4) & 63)))
                    fmt.Println(v4, v3, v3)
                    var v10 uint16 = 2
                    v1 = v1
                    var v11 bool = false
                    switch 128 {
                    case 42:
                        var v12 bool = false
                        _ = v12
                    case -9223372036854775808:
                        v13 := GenS1{F0: v11, F1: v1}
                        _ = v13
                    default:
                        v11 = true
                    }
                    fmt.Println(v4)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v14 int8 = -128
                v5 = append(v5, v3)
                _ = fi8
                _ = v14
            }
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi15 := 0; fi15 < 10; fi15++ {
                    v4 -= 4294967295
                    var v16 int8 = 100
                    _ = fi15
                    _ = v16
                }
            }
        }
    case "a":
        switch int(v3) {
        case 4294967295:
            fmt.Println(v4, v3, v1)
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint32 = 42
                for fi19 := 0; fi19 < 15; fi19++ {
                    fmt.Println(fi19)
                    switch (-(fi19)) {
                    case 9223372036854775807:
                        v1 = "a"
                    case 128:
                        v20 := GenS2{F0: (fi19 % ((fi19) | 1))}
                        _ = v20
                    }
                    _ = fi19
                }
                v18 -= 161
                _ = fi17
                _ = v18
            }
        }
    default:
        v4 -= 1
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 int32 = 42
        _ = fi21
        _ = v22
    }
    v5 = append(v5, 4294967295)
    v4 = 543
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v23 uint32 = v3
    var v24 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
