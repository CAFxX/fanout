package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint16 = 862
    v2 := make(map[string]bool)
    var v3 uint = 4294967295
    v4 := []string{"\n\t"}
    var v5 int = -2147483648
    var v6 uint = (v3 % ((v3) | 1))
    _ = len(v2)
    _ = len(v4)
    v1 += 885
    v1 -= v1
    var v7 int = v5
    var v8 uint64 = 0
    var v9 float32 = (0.1 / 1.5)
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 bool = (-2147483648 <= -975)
        switch v7 {
        case 873:
            v12 := GenS1{F0: -1}
            _ = v12
        case -9223372036854775808:
            var v13 int16 = 224
            _ = v13
        default:
            for fi14 := 0; fi14 < 6; fi14++ {
                switch "hello" {
                case "foo bar":
                    v6 -= uint(v1)
                }
                switch ("foo bar" + "a") {
                case "hello":
                    switch "\n\t" {
                    case "a":
                        switch (fi10 * v5) {
                        case -9223372036854775808:
                            var v15 uint32 = 0
                            _ = v15
                        case 2147483647:
                            var v16 int = (v7 ^ (v5 & (-(fi14))))
                            _ = v16
                        }
                    }
                case "a":
                    v6 = v6
                default:
                    var v17 int64 = -864
                    _ = v17
                }
                _ = fi14
            }
        }
        var v18 bool = (v11 || ("foo bar" >= ""))
        _ = fi10
        _ = v11
        _ = v18
    }
    switch (v7 ^ (v7 % ((32767) | 1))) {
    case 74:
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v8, v9)
            delete(v2, "a")
            v20 := GenS2{F0: (("a" + "\n\t") <= "a")}
            _ = fi19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
