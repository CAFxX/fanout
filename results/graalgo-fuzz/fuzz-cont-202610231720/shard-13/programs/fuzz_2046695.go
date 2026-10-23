package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
func genFn0(p0 int16) int16 {
    p0 = -32768
    var v5 bool = false
    fmt.Println(v5)
    _ = p0
    _ = v5
    return (p0 + (-(p0)))
}

func main() {
    var v1 uint = 629
    v2 := []string{"a", "hello"}
    var v3 float32 = (-2.25 / float32(0.1))
    var v4 int = -2147483648
    switch ("hello" + "\n\t") {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: (("foo bar" + "foo bar") + (((("hello" + "hello") + "hello") + "hello") + "\n\t")), F1: 989, F2: 921}
            _ = fi6
            _ = v7
        }
    case "foo bar":
        if (3.14159 >= 3.14159) {
            if (100 != -4) {
                var v8 uint = (^(v1))
                var v9 int64 = -545
                fmt.Println(v8, v8, v3)
                _ = v8
                _ = v9
            } else {
                var v10 int16 = 1000
                _ = v10
            }
            var v11 uint16 = 1
            _ = v11
        }
    }
    v4 = (v4 * (-(v4)))
    if (!false) {
        v3 = ((v3 + 0.1) - float32(3.14159))
    } else {
        if (23 == 42) {
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := GenS1{F0: ("foo bar" + "\n\t"), F1: 0, F2: 18446744073709551615}
                v1 = 18446744073709551615
                _ = len(v2)
                _ = fi12
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int = v4
                _ = fi14
                _ = v15
            }
        }
        fmt.Println(v1, v3, v3)
        fmt.Println(v4)
    }
    v1 -= (v1 >> ((2) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += 1.5
    var v16 int64 = 171
    v17 := GenS1{F0: "foo bar", F1: 2, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
