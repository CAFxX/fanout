package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 string; F2 uint8 }

func main() {
    v1 := []string{("hello" + ((("\n\t" + "foo bar") + "\n\t") + ("foo bar" + "hello"))), "\n\t"}
    var v2 bool = false
    v3 := make(map[uint32]bool)
    var v4 uint8 = 7
    var v5 int32 = 65
    v2 = v2
    var v6 int64 = int64(v5)
    var v7 float64 = -2.25
    var v8 uint64 = 18446744073709551615
    var v9 bool = true
    if (0.0 == float32(1.5)) {
        v1 = append(v1, "foo bar")
        if v2 {
            if ((v8 != (^(v8))) || v9) {
                v10 := v3[2147483648]
                var v11 bool = (v2 && v9)
                _ = v10
                _ = v11
            }
            fmt.Println(v8, v9, v8)
        }
        switch "a" {
        case "\n\t":
            for fi12 := 0; fi12 < 5; fi12++ {
                delete(v3, 0)
                if (((v4 ^ (v4 ^ 255)) < (v4 - (v4 % (((^(v4))) | 1)))) && v9) {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v14 := GenS1{F0: 1}
                        fmt.Println(fi13, fi12)
                        switch (fi13 / ((-392) | 1)) {
                        case -1:
                            v15 := GenS2{F0: -1, F1: "\n\t", F2: 1}
                            _ = v15
                        }
                        _ = fi13
                        _ = v14
                    }
                } else {
                    if v2 {
                        var v16 uint32 = 1
                        _ = v16
                    } else {
                        v6 += v6
                        v8 -= (^(v8))
                    }
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v3[2147483648] = (!(2147483648 < 42))
                        switch "\n\t" {
                        case "foo bar":
                            var v18 int32 = (v5 & 810)
                            _ = v18
                        case "\n\t":
                            var v19 int8 = -29
                            _ = v19
                        case "a":
                            var v20 int32 = v5
                            _ = v20
                        default:
                            fmt.Println(v8)
                        }
                        _ = fi17
                    }
                }
                _ = fi12
            }
        case "a":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        default:
            switch ((((("hello" + ("a" + ((("hello" + ("a" + "\n\t")) + "a") + ("\n\t" + "\n\t")))) + "foo bar") + "") + "\n\t") + "") {
            case "hello":
                v8 += (v8 & v8)
            case "a":
                if ((!(uint32(v8) >= 42)) && true) {
                    var v21 string = "a"
                    v2 = true
                    v22 := GenS2{F0: -1, F1: "\n\t", F2: uint8(v8)}
                    _ = v21
                    _ = v22
                } else {
                    v9 = v2
                }
            }
        }
    } else {
        v23 := GenS1{F0: 32767}
        var v24 int32 = v5
        _ = v23
        _ = v24
    }
    var v25 int32 = (v5 - (v5 / ((v5) | 1)))
    var v26 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
