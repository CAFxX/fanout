package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 int64) string {
    v6 := GenS1{F0: "\n\t", F1: 4294967295, F2: 42}
    if false {
        v7 := GenS1{F0: ("hello" + (("\n\t" + "hello") + ("" + (((((("a" + "foo bar") + "a") + ("\n\t" + ((((("a" + "hello") + "\n\t") + "hello") + "foo bar") + "foo bar"))) + "a") + "hello") + "a")))), F1: -736, F2: 65535}
        _ = v7
    } else {
        v8 := []int{}
        v9 := GenS1{F0: (("a" + ("\n\t" + ("foo bar" + "\n\t"))) + "a"), F1: -9223372036854775808, F2: 42}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    return "a"
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]string)
    var v3 string = "\n\t"
    v4 := []string{"a"}
    var v5 uint = 2
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v5, v5, v1)
        v11 := GenS1{F0: ("a" + (v3 + "foo bar")), F1: (^(fi10)), F2: 391}
        fmt.Println(v5, v1, v3)
        _ = fi10
        _ = v11
    }
    v3 = "foo bar"
    var v12 uint32 = v1
    v13 := GenS2{F0: 127, F1: v3, F2: 978}
    v14 := GenS1{F0: v3, F1: 42, F2: v12}
    for fi15 := 0; fi15 < 8; fi15++ {
        if ((("foo bar" + "hello") + "") <= (v3 + v3)) {
            fmt.Println(v3)
            fmt.Println(v5, v12)
            v12 += v12
        } else {
            v16 := GenS3{F0: 1, F1: 94}
            switch v3 {
            case "\n\t":
                switch (fi15 << ((255) & 63)) {
                case 806:
                    for fi17 := 0; fi17 < 11; fi17++ {
                        if true {
                            v18 := GenS1{F0: ("\n\t" + (v3 + v3)), F1: -2147483648, F2: (v12 + (v1 / ((v1) | 1)))}
                            _ = len(v4)
                            var v19 int32 = 552
                            fmt.Println(fi17, v3)
                            _ = v18
                            _ = v19
                        } else {
                            var v20 uint64 = uint64(v1)
                            v20 = (^(v20))
                            _ = v20
                        }
                        var v21 float64 = 0.0
                        v22 := GenS2{F0: 0, F1: v3, F2: 65535}
                        _ = fi17
                        _ = v21
                        _ = v22
                    }
                case 761:
                    switch v3 {
                    case "hello":
                        _ = len(v2)
                    case "foo bar":
                        delete(v2, (fi15 << ((v5) & 63)))
                    }
                case 833:
                    var v23 uint16 = 2
                    _ = v23
                }
            case "hello":
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 int32 = 0
                    fmt.Println(v1, v25, v5)
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                v5 -= v5
            }
            _ = v16
        }
        _ = fi15
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 string = v3
        _ = fi26
        _ = v27
    }
    if ("\n\t" >= v3) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        _ = len(v4)
    }
    if len(v4) > 0 {
        v4[0] = (v3 + v3)
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v1 = (^(v12))
        var v29 uint32 = v1
        _ = fi28
        _ = v29
    }
    if true {
        var v30 int64 = -1
        var v31 bool = true
        _ = v30
        _ = v31
    }
    var v32 int32 = -535
    v1 -= v1
    var v33 uint32 = v12
    fmt.Println(v32, v32, v1)
    for fi34 := 0; fi34 < 4; fi34++ {
        _ = len(v2)
        _ = fi34
    }
    _ = len(v2)
    delete(v2, 128)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
