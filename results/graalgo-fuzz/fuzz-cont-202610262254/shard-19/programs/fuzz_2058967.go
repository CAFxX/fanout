package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []int{4294967295, 42, 32768}
    v2 := make(map[string]int)
    var v3 int = 128
    v4 := GenS1{F0: 1}
    var v5 int8 = 127
    v3 -= (v3 / (((v3 >> ((18446744073709551615) & 63))) | 1))
    v6 := GenS1{F0: 178}
    switch "hello" {
    case "foo bar":
        v7 := GenS1{F0: 1}
        _ = v7
    case "a":
        v8 := GenS1{F0: 8}
        _ = v8
    case "hello":
        var v9 float32 = -2.25
        _ = v9
    default:
        switch (("hello" + ("foo bar" + "a")) + "hello") {
        case "":
            for fi10 := 0; fi10 < 17; fi10++ {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch ("\n\t" + ("hello" + ("hello" + "hello"))) {
                    case "foo bar":
                        v5 += (v5 << ((42) & 63))
                    default:
                        fmt.Println(fi11, fi11, v5)
                    }
                    switch "" {
                    case "\n\t":
                        var v12 uint64 = 1
                        _ = v12
                    default:
                        var v13 uint = 0
                        _ = v13
                    }
                    _ = fi11
                }
                _ = fi10
            }
        default:
            v14 := GenS2{F0: (v5 % (((v5 ^ -1)) | 1))}
            _ = v14
        }
    }
    delete(v2, "foo bar")
    v2["foo bar"] = v3
    var v15 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
