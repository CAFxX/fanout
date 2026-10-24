package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint32 = 65535
    v2 := []uint32{(v1 / (((v1 / (((v1 * (v1 - 2147483647))) | 1))) | 1)), 65535, (v1 | (v1 | (v1 >> ((2) & 63))))}
    var v3 bool = true
    v4 := make(map[uint32]string)
    v5 := GenS1{F0: 0}
    var v6 uint32 = 1
    v7 := GenS2{F0: 0}
    if v3 {
        switch (("\n\t" + "\n\t") + ((((("foo bar" + "foo bar") + "a") + "hello") + "") + "a")) {
        case "hello":
            for fi8 := 0; fi8 < 1; fi8++ {
                v3 = v3
                fmt.Println(fi8, v6, v1)
                _ = fi8
            }
        case "a":
            fmt.Println(v3, v3)
        case "foo bar":
            var v9 int = 4294967295
            _ = v9
        }
        fmt.Println(v6)
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        switch 32768 {
        case 2147483647:
            for fi11 := 0; fi11 < 2; fi11++ {
                v1 -= 2147483648
                switch ((("a" + "\n\t") + "hello") + "a") {
                case "\n\t":
                    v6 = v6
                case "":
                    switch (fi10 / ((525) | 1)) {
                    case -155:
                        v3 = v3
                    case 902:
                        var v12 int32 = -183
                        _ = v12
                    }
                case "hello":
                    v1 += v1
                default:
                    v13 := GenS2{F0: fi11}
                    _ = v13
                }
                _ = fi11
            }
        }
        _ = fi10
    }
    v3 = v3
    v14 := GenS1{F0: 1}
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(v1, v1, fi15)
        _ = fi15
    }
    var v16 uint = 4294967295
    for fi17 := 0; fi17 < 3; fi17++ {
        v2 = append(v2, v1)
        _ = fi17
    }
    var v18 uint64 = 683
    fmt.Println(v3, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
