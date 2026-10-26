package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{-2147483648, -2147483648, 2147483647}
    var v2 int8 = -1
    var v3 string = "hello"
    if (42 > 1) {
        v2 += v2
    } else {
        switch 32767 {
        case 127:
            if false {
                if false {
                    if (true || (false || (1.5 >= 3.14159))) {
                        var v4 uint32 = 65535
                        _ = v4
                    } else {
                        v2 += v2
                    }
                    v5 := GenS1{F0: v3}
                    var v6 int16 = -1
                    _ = v5
                    _ = v6
                }
            }
        default:
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 int64 = 1
                _ = fi7
                _ = v8
            }
        }
    }
    var v9 uint16 = 32767
    if (((-2.25 * (0.1 / 1.5)) / 1.0) != 0.0) {
        var v10 int = 127
        if len(v1) > 0 {
            v1[0] = v10
        }
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: 77}
        var v13 uint32 = 4294967295
        _ = fi11
        _ = v12
        _ = v13
    }
    v2 -= v2
    v14 := GenS2{F0: 2}
    switch v3 {
    case "hello":
        var v15 int32 = 2147483647
        _ = v15
    case "foo bar":
        var v16 uint = 18446744073709551615
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 bool = (!(!(2 <= 0)))
            var v19 string = v3
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        switch v3 {
        case "\n\t":
            if (32768 >= 128) {
                v20 := GenS1{F0: v3}
                _ = len(v1)
                var v21 uint64 = 4294967295
                _ = v20
                _ = v21
            }
        }
    }
    v3 = "a"
    fmt.Println(v3, v3)
    switch (v3 + v3) {
    case "hello":
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS1{F0: v3}
            fmt.Println(v3)
            _ = fi22
            _ = v23
        }
    case "":
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 uint16 = (v9 & v9)
            for fi26 := 0; fi26 < 1; fi26++ {
                switch -126 {
                case 971:
                    var v27 uint16 = (v25 + 65535)
                    _ = v27
                case 32767:
                    var v28 bool = (303 < 42)
                    _ = v28
                case 0:
                    var v29 uint = 84
                    _ = v29
                }
                if false {
                    switch ((("a" + ("foo bar" + v3)) + v3) + ((v3 + ("hello" + (v3 + "foo bar"))) + "\n\t")) {
                    case "":
                        fmt.Println(v3, fi24)
                    case "a":
                        v30 := GenS1{F0: (v3 + ((("a" + v3) + v3) + (("foo bar" + "a") + v3)))}
                        _ = v30
                    default:
                        switch "hello" {
                        case "hello":
                            v2 += v2
                        }
                    }
                    v9 = 32767
                    var v31 float64 = (1.5 / float64(((((float32(0.0) - 1e10) + (-2.25 * 1.5)) * (0.0 + -2.25)) - 0.1)))
                    fmt.Println(v3, fi24)
                    _ = v31
                } else {
                    fmt.Println(v3, v9)
                }
                v3 = ""
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
    default:
        var v32 float32 = 1e10
        _ = v32
    }
    v1 = append(v1, 128)
    var v33 int16 = 352
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v9
}
