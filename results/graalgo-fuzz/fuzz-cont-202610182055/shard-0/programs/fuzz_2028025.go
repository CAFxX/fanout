package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int = 0
    var v2 uint8 = 128
    var v3 bool = false
    v4 := []int{32767, 2147483647, v1}
    v5 := []string{"hello", (("" + (("a" + "a") + "hello")) + ("a" + "hello")), ("foo bar" + (("hello" + "") + ("hello" + "foo bar")))}
    var v6 float32 = 3.14159
    fmt.Println(v6, v1)
    for fi7 := 0; fi7 < 13; fi7++ {
        switch "a" {
        case "a":
            v6 -= 3.14159
        case "hello":
            for fi8 := 0; fi8 < 13; fi8++ {
                if false {
                    v9 := GenS1{F0: 60}
                    _ = v9
                } else {
                    fmt.Println(fi8, v3)
                    var v10 int16 = 1
                    _ = v10
                }
                switch ("a" + "foo bar") {
                case "\n\t":
                    v11 := GenS1{F0: 32767}
                    _ = v11
                case "":
                    switch "foo bar" {
                    case "\n\t":
                        switch "" {
                        case "a":
                            if v3 {
                                v2 = (v2 >> ((523) & 63))
                            }
                        case "hello":
                            v12 := GenS1{F0: 1}
                            _ = v12
                        }
                    case "foo bar":
                        var v13 uint32 = 65535
                        _ = v13
                    default:
                        var v14 int8 = -74
                        _ = v14
                    }
                }
                _ = fi8
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi7
    }
    fmt.Println(v6, v6)
    v2 = v2
    for fi15 := 0; fi15 < 4; fi15++ {
        if len(v4) > 0 {
            v4[0] = fi15
        }
        _ = fi15
    }
    if ((2147483647 <= 65535) && true) {
        _ = len(v5)
        if (65535 >= 0) {
            var v16 float32 = v6
            v1 -= (v1 % ((42) | 1))
            var v17 uint32 = 65535
            _ = v16
            _ = v17
        }
        v4 = append(v4, 9223372036854775807)
    } else {
        v2 -= v2
        _ = len(v4)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if false {
        var v18 int64 = 290
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 string = "foo bar"
        _ = fi19
        _ = v20
    }
    if len(v5) > 0 {
        v5[0] = ("\n\t" + "")
    }
    var v21 int32 = 0
    for fi22 := 0; fi22 < 19; fi22++ {
        _ = len(v5)
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(v6, fi23)
            _ = fi23
        }
        _ = fi22
    }
    v3 = v3
    fmt.Println(v6, v21)
    var v24 bool = true
    switch "hello" {
    case "a":
        fmt.Println(v21)
    case "\n\t":
        if (255 <= v2) {
            var v25 uint32 = 1
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 int64 = 56
                if false {
                    var v28 int32 = v21
                    var v29 float32 = v6
                    v2 = (v2 ^ (v2 + (v2 ^ v2)))
                    _ = v28
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
