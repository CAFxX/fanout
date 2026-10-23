package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []string{"a", ("a" + ("" + ("a" + "a"))), ("foo bar" + "\n\t")}
    v2 := make(map[uint32]uint32)
    var v3 int8 = 127
    var v4 float32 = 1e10
    fmt.Println(v3)
    for fi5 := 0; fi5 < 5; fi5++ {
        if (true || true) {
            var v6 int16 = -32768
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int = (fi5 ^ -152)
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        _ = fi5
    }
    var v9 uint32 = 65535
    switch ("foo bar" + "hello") {
    case "hello":
        for fi10 := 0; fi10 < 14; fi10++ {
            v3 += -128
            v11 := GenS1{F0: "", F1: 2, F2: 66}
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS1{F0: "", F1: 32767, F2: 128}
        _ = v12
    }
    fmt.Println(v3)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint64 = 18446744073709551615
        for fi15 := 0; fi15 < 6; fi15++ {
            v3 = 37
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v9, v4)
    v1 = append(v1, "a")
    for fi16 := 0; fi16 < 4; fi16++ {
        v4 += v4
        fmt.Println(v3, v4)
        if false {
            if false {
                fmt.Println(v4, v4)
            } else {
                fmt.Println(v9, fi16)
                v17 := GenS1{F0: "\n\t", F1: 505, F2: 42}
                var v18 int = int(v3)
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: "hello", F1: 42, F2: 127}
            for fi20 := 0; fi20 < 3; fi20++ {
                if (true && (true && false)) {
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    v9 += v9
                } else {
                    fmt.Println(v4, v3)
                    v21 := GenS2{F0: 65535}
                    _ = v21
                }
                v22 := v2[464]
                _ = len(v2)
                fmt.Println(v9, v4)
                _ = fi20
                _ = v22
            }
            _ = v19
        } else {
            if true {
                fmt.Println(v9, fi16)
            }
            var v23 string = "a"
            _ = v23
        }
        _ = fi16
    }
    v9 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
