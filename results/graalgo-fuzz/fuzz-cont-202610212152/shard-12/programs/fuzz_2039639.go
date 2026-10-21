package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 string = (("a" + "foo bar") + ("a" + ""))
    v2 := []uint32{}
    var v3 uint = 42
    var v4 uint = v3
    var v5 int = 146
    var v6 int8 = -5
    switch "hello" {
    case "a":
        v7 := GenS1{F0: 32767}
        _ = v7
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 1
        }
    case "\n\t":
        v4 = v4
    }
    if (0 >= -1) {
        var v8 uint64 = 42
        for fi9 := 0; fi9 < 17; fi9++ {
            v8 -= v8
            if true {
                if true {
                    switch (v5 << (((v3 - v4)) & 63)) {
                    case 128:
                        v4 = v4
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if true {
                        var v10 int16 = int16(v6)
                        _ = v10
                    } else {
                        var v11 bool = true
                        switch int(v3) {
                        case -1:
                            var v12 float64 = (0.1 + 1.5)
                            _ = v12
                        default:
                            switch v5 {
                            case 4294967295:
                                v13 := GenS1{F0: v5}
                                _ = v13
                            case 42:
                                fmt.Println(v4, v8, v4)
                            default:
                                var v14 uint16 = 33
                                _ = v14
                            }
                        }
                        _ = v11
                    }
                    var v15 uint64 = v8
                    _ = v15
                }
            }
            _ = fi9
        }
        fmt.Println(v4, v5, v8)
        _ = v8
    }
    switch v1 {
    case "\n\t":
        v16 := GenS1{F0: v5}
        _ = v16
    case "hello":
        if ((v5 - 127) >= -2147483648) {
            var v17 int32 = -2147483648
            fmt.Println(v17)
            v18 := GenS1{F0: v5}
            _ = v17
            _ = v18
        }
    case "":
        var v19 int32 = 607
        _ = v19
    default:
        v6 = (v6 | (v6 * v6))
    }
    v3 = (v3 >> ((65535) & 63))
    var v20 float32 = ((1e10 + ((0.1 * (float32((0.1 + 0.0)) + ((3.14159 * 1e10) * ((3.14159 - 1e10) * -2.25)))) + -2.25)) * 3.14159)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 bool = (2147483647 >= 2147483647)
    for fi22 := 0; fi22 < 5; fi22++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi22
    }
    switch ("foo bar" + "foo bar") {
    case "":
        v20 = (((1.5 + 1.5) * v20) + float32(-2.25))
    case "foo bar":
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 bool = (!((v6 << ((65535) & 63)) > 0))
            _ = fi23
            _ = v24
        }
    }
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
