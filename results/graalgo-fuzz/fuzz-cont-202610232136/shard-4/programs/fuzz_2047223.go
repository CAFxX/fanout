package main

import (
    "fmt"
)

func genFn0(p0 uint8) string {
    fmt.Println(p0)
    v5 := []int{0, 32767}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return "\n\t"
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 0
    var v3 float64 = 3.14159
    var v4 uint32 = 1
    v6 := []int{32768, 9223372036854775807, 0}
    switch "foo bar" {
    case "hello":
        v1 = (^(v1))
    case "\n\t":
        v2 -= v2
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v4 -= (^(v4))
        v2 += v2
        _ = fi7
    }
    switch ("" + "a") {
    case "":
        v4 = 2
    case "hello":
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v2 = (v2 << ((42) & 63))
                var v10 string = (("hello" + ("foo bar" + ((("a" + "foo bar") + ("hello" + "")) + ("hello" + ((("hello" + (("\n\t" + (("\n\t" + "\n\t") + "")) + "a")) + (("a" + ("\n\t" + "\n\t")) + (("hello" + "hello") + ((("hello" + "") + "hello") + "\n\t")))) + ("\n\t" + "\n\t")))))) + ("foo bar" + "a"))
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    fmt.Println(v2, v1)
    if true {
        v4 -= uint32(v2)
    }
    var v11 uint16 = v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint8 = 255
        _ = fi12
        _ = v13
    }
    v11 += 2
    fmt.Println(v1)
    if len(v6) > 0 {
        v6[0] = 42
    }
    if len(v6) > 0 {
        v6[0] = 127
    }
    v11 = (v1 - v1)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
