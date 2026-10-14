package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 42
    var v3 uint16 = 316
    var v4 int64 = 9223372036854775807
    v5 := []int{}
    var v6 string = ""
    for fi7 := 0; fi7 < 2; fi7++ {
        v4 = (v4 % ((v4) | 1))
        _ = fi7
    }
    fmt.Println(v3, v6, v2)
    v8 := GenS1{F0: -1, F1: 156, F2: 2}
    v2 = v2
    v2 = (v2 * (v2 & v2))
    if true {
        _ = len(v5)
        switch (v6 + "") {
        case "a":
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        case "\n\t":
            if false {
                v4 = 231
            }
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            v3 += (v3 / ((530) | 1))
            v4 -= v4
            var v10 int16 = 32767
            _ = fi9
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            v6 = v6
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint64 = 1
                var v14 uint16 = (v3 >> ((58) & 63))
                var v15 uint8 = 127
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi11
        }
        v16 := GenS1{F0: -1, F1: 76, F2: 0}
        if true {
            switch "" {
            case "\n\t":
                if false {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        v5 = append(v5, (fi17 | (^(fi17))))
                        v2 = v2
                        _ = fi17
                    }
                    var v18 int64 = v4
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := GenS2{F0: 42, F1: -32768, F2: 0}
                            v6 = ("hello" + ("" + (("\n\t" + (v6 + v6)) + "a")))
                            _ = fi20
                            _ = v21
                        }
                        switch (v6 + (((("\n\t" + v6) + v6) + (((v6 + "a") + "\n\t") + v6)) + "foo bar")) {
                        case "":
                            var v22 int16 = -1
                            _ = v22
                        case "hello":
                            fmt.Println(v3, v6)
                        default:
                            var v23 uint16 = v3
                            _ = v23
                        }
                        _ = fi19
                    }
                    fmt.Println(v6, v4, v6)
                }
            case "":
                var v24 int8 = -1
                _ = v24
            }
            v2 = (v2 * 190)
        }
        _ = v16
    }
    var v25 int = 592
    fmt.Println(v25, v6, v25)
    for fi26 := 0; fi26 < 18; fi26++ {
        switch v6 {
        case "foo bar":
            var v27 uint8 = 127
            _ = v27
        case "hello":
            fmt.Println(v6, v3, fi26)
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(v25)
            _ = fi28
        }
        var v29 int = 32767
        _ = fi26
        _ = v29
    }
    fmt.Println(v4, v3)
    fmt.Println(v6)
    var v30 int = v25
    var v31 uint64 = 18446744073709551615
    v3 = (v3 % ((v3) | 1))
    var v32 float64 = 0.0
    v32 -= (1.5 * v32)
    var v33 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
