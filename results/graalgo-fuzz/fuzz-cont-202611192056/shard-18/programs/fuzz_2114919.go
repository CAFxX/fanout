package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint8 = 83
    v2 := []int{}
    var v3 uint32 = 2
    var v4 float64 = 0.0
    fmt.Println(v1, v1, v3)
    var v5 int8 = 127
    v1 = (v1 * (v1 & (v1 - 128)))
    v4 = v4
    v5 = -72
    v3 -= 2147483648
    fmt.Println(v1)
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "hello" {
        case "foo bar":
            v4 += 1.5
        case "":
            if (("a" + "") <= (((("foo bar" + "") + "") + (("\n\t" + "") + ("foo bar" + ("\n\t" + "a")))) + ("hello" + "a"))) {
                if (42 > -537) {
                    v5 = (v5 / (((v5 % (((v5 << ((1) & 63))) | 1))) | 1))
                }
                v7 := GenS1{F0: "a"}
                _ = v7
            }
        case "a":
            fmt.Println(v5, v5)
        }
        _ = fi6
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 = uint8(v3)
        if false {
            var v9 uint8 = v1
            _ = v9
        }
        _ = fi8
    }
    fmt.Println(v4, v3)
    var v10 uint16 = 924
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
