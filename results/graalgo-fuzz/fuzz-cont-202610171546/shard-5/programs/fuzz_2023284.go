package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0() int16 {
    var v5 uint8 = 255
    v6 := GenS1{F0: 2147483647, F1: 17}
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 uint64 = 0
    var v2 uint8 = 2
    v3 := []string{(("a" + "\n\t") + ((((((("a" + (("hello" + ("" + "foo bar")) + ("\n\t" + "a"))) + "") + "hello") + "\n\t") + "a") + ("foo bar" + "foo bar")) + ("hello" + ((("a" + ("hello" + ("" + ""))) + "a") + "a"))))}
    v4 := make(map[int]int)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS2{F0: 766, F1: "a", F2: 127}
        _ = fi7
        _ = v8
    }
    v1 = v1
    if true {
        fmt.Println(v2)
    }
    v9 := GenS2{F0: 216, F1: "hello", F2: -128}
    v3 = append(v3, "a")
    switch ("a" + ("\n\t" + "foo bar")) {
    case "hello":
        var v10 uint32 = 0
        _ = v10
    case "foo bar":
        if (("\n\t" + "hello") <= "") {
            var v11 uint32 = 65535
            if false {
                for fi12 := 0; fi12 < 17; fi12++ {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 uint16 = 116
                        var v15 int64 = 1
                        v16 := GenS2{F0: 82, F1: "\n\t", F2: -1}
                        switch (fi12 + fi12) {
                        case 4294967295:
                            v11 = v11
                        default:
                            v1 = v1
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    v3 = append(v3, "foo bar")
                    _ = fi12
                }
            }
            var v17 uint32 = v11
            _ = v11
            _ = v17
        }
    }
    switch "a" {
    case "":
        var v18 int8 = 60
        _ = v18
    case "\n\t":
        fmt.Println(v2, v2)
    case "foo bar":
        var v19 uint32 = 4294967295
        _ = v19
    }
    v20 := GenS2{F0: 65535, F1: "\n\t", F2: 42}
    if false {
        var v21 float64 = 1.5
        _ = v21
    }
    v2 += 0
    v22 := v4[0]
    v2 -= (^(v2))
    v23 := v4[v22]
    v22 = (v23 ^ (v23 - (v23 & 9223372036854775807)))
    v24 := GenS3{F0: (2 != 4294967295), F1: (v1 - 18446744073709551615)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
