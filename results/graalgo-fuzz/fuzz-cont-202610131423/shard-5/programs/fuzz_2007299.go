package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = -241
    var v2 string = (("foo bar" + ("foo bar" + "foo bar")) + (("" + "foo bar") + (((("\n\t" + "\n\t") + ("a" + "hello")) + "hello") + "foo bar")))
    var v3 uint64 = 332
    v4 := make(map[uint32]uint32)
    var v5 int = int(v3)
    if false {
        v1 = (v1 + v1)
        var v6 int32 = -509
        fmt.Println(v2)
        _ = v6
    }
    switch v2 {
    case "\n\t":
        fmt.Println(v5, v1, v3)
    case "":
        var v7 uint16 = 763
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(v3)
                var v10 uint8 = 128
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 1; fi11++ {
                if (3.14159 >= (1.5 * 0.1)) {
                    switch v2 {
                    case "hello":
                        v5 += v5
                    case "foo bar":
                        fmt.Println(v1, v2)
                    case "\n\t":
                        v2 = v2
                    }
                    v12 := []int{(v5 ^ v5), fi11, (v5 % (((v5 << ((18446744073709551615) & 63))) | 1))}
                    v13 := GenS1{F0: 2}
                    _ = v12
                    _ = v13
                }
                _ = fi11
            }
            var v14 bool = true
            _ = fi8
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint64 = (v3 & (v3 % (((v3 * 1)) | 1)))
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v3, v3)
    v17 := GenS1{F0: 0}
    var v18 int = v5
    var v19 int16 = -1
    fmt.Println(v19)
    var v20 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
