package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 int; F2 int32 }

func main() {
    var v1 uint32 = 65535
    var v2 int = 32767
    var v3 uint8 = 88
    v4 := []int{}
    v5 := make(map[string]string)
    var v6 string = "hello"
    v5[v6] = v6
    var v7 int32 = -307
    var v8 int64 = 571
    delete(v5, v6)
    if true {
        fmt.Println(v8, v2)
        var v9 float32 = ((float32(float64(0.0)) * -2.25) * 3.14159)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            fmt.Println(v2, v8)
            if false {
                if (false && (true && ((v3 + v3) <= v3))) {
                    v11 := GenS1{F0: 42, F1: v2}
                    _ = v11
                }
                switch -2147483648 {
                case 4294967295:
                    switch v6 {
                    case "a":
                        switch (v6 + v6) {
                        case "\n\t":
                            var v12 string = v6
                            _ = v12
                        default:
                            fmt.Println(v1, v6)
                        }
                    case "foo bar":
                        v6 = ("hello" + v6)
                    default:
                        v13 := GenS2{F0: -795, F1: (^(fi10)), F2: int32(v8)}
                        _ = v13
                    }
                case 128:
                    var v14 string = v6
                    _ = v14
                case 42:
                    v8 = (v8 | 9223372036854775807)
                default:
                    var v15 int8 = -128
                    _ = v15
                }
                var v16 uint16 = 542
                _ = v16
            }
            _ = fi10
        }
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 string = (v6 + (("" + (v6 + (v6 + "hello"))) + v6))
        if (!false) {
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 int32 = v7
                if ((3.14159 - 1e10) != 1e10) {
                    v8 = (-(v8))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v4 = append(v4, 32768)
                _ = fi19
                _ = v20
            }
        }
        _ = fi17
        _ = v18
    }
    if true {
        switch "hello" {
        case "hello":
            v3 += 255
        case "a":
            v21 := GenS2{F0: 1, F1: 411, F2: int32(v3)}
            _ = v21
        case "\n\t":
            v8 -= v8
        }
    } else {
        fmt.Println(v3)
        v8 = (v8 + v8)
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS1{F0: 515, F1: fi22}
            fmt.Println(v8, v1)
            _ = fi22
            _ = v23
        }
    }
    switch v2 {
    case 9223372036854775807:
        var v24 int = int(v8)
        _ = v24
    default:
        var v25 bool = (v2 >= v2)
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
