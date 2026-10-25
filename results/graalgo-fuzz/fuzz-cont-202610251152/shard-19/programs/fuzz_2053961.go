package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{65535, 0}
    var v3 uint64 = 4294967295
    var v4 float32 = (float32(3.14159) - 0.1)
    v5 := make(map[string]string)
    var v6 int16 = 42
    var v7 uint16 = 915
    if false {
        v7 += uint16(v3)
    }
    var v8 uint32 = uint32(v7)
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: -810}
            switch ("\n\t" + ("hello" + (("" + (("hello" + "a") + "\n\t")) + ""))) {
            case "\n\t":
                if (!true) {
                    v8 = v8
                    if false {
                        v5["\n\t"] = (("\n\t" + "foo bar") + "")
                        fmt.Println(fi10, v3)
                        switch "\n\t" {
                        case "a":
                            v7 += v7
                        case "hello":
                            switch (fi10 % ((fi10) | 1)) {
                            case -2147483648:
                                v3 -= 18446744073709551615
                            case 0:
                                fmt.Println(fi9, v3)
                            }
                        default:
                            if false {
                                if false {
                                    v4 = float32(0.1)
                                }
                                switch ("" + "hello") {
                                case "":
                                    v4 = (0.1 - 3.14159)
                                case "hello":
                                    var v12 uint32 = (v8 % (((v8 ^ (v8 | v8))) | 1))
                                    _ = v12
                                default:
                                    v13 := GenS1{F0: 91}
                                    _ = v13
                                }
                                fmt.Println(v3, v7)
                            }
                        }
                    }
                }
            }
            _ = fi10
            _ = v11
        }
        v7 = 1
        _ = fi9
    }
    if false {
        v14 := GenS2{F0: v8, F1: 61, F2: -2147483648}
        delete(v1, -181)
        fmt.Println(v4, v7)
        _ = v14
    }
    var v15 uint16 = 1
    v8 += v8
    v3 = v3
    var v16 uint = 2
    var v17 string = "foo bar"
    var v18 int16 = v6
    fmt.Println(v6, v8, v6)
    var v19 uint32 = v8
    if true {
        v7 += v15
        switch int(v16) {
        case -711:
            var v20 int64 = -9223372036854775808
            _ = v20
        }
        fmt.Println(v6, v6)
    } else {
        if (true || true) {
            for fi21 := 0; fi21 < 20; fi21++ {
                fmt.Println(v15, v3)
                _ = fi21
            }
            if ((v17 + ("" + "foo bar")) == v17) {
                v15 = (v15 | v7)
                v8 = (v8 ^ 4294967295)
                var v22 int8 = 100
                _ = v22
            }
            v2 = append(v2, v8)
        }
        var v23 uint = 0
        var v24 bool = false
        _ = v23
        _ = v24
    }
    fmt.Println(v7, v8, v19)
    switch 0 {
    case 32768:
        v17 = v17
    case 529:
        var v25 uint64 = v3
        _ = v25
    default:
        for fi26 := 0; fi26 < 5; fi26++ {
            v19 += v19
            _ = fi26
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := v1[1]
        v19 -= v8
        v3 = (v3 * (v3 / ((v3) | 1)))
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
