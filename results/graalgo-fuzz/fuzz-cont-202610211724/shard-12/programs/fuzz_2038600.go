package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 65535
    v2 := []int{22}
    v3 := []uint32{(v1 * 643), 286, 2147483647, (v1 & (v1 ^ (v1 << ((1) & 63))))}
    var v4 uint32 = 0
    var v5 int16 = 32767
    v1 = v1
    switch "hello" {
    case "\n\t":
        var v6 float32 = 1.5
        _ = v6
    case "foo bar":
        v5 += 32767
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v4, v1, v4)
        _ = fi7
    }
    var v8 float32 = 0.1
    if ((255 <= 149) || ((v5 * 1000) != (-(v5)))) {
        v8 = v8
        v4 = 42
    }
    var v9 uint = 63
    var v10 int16 = -1
    v4 = v1
    v3 = append(v3, 2147483648)
    if false {
        if (("foo bar" < "") || true) {
            v10 += v10
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(v1, v4, fi11)
                var v12 string = "a"
                _ = fi11
                _ = v12
            }
        }
    } else {
        v3 = append(v3, 2147483647)
        var v13 uint = (v9 + v9)
        for fi14 := 0; fi14 < 3; fi14++ {
            v8 += (-2.25 + v8)
            switch ("" + (("hello" + ("\n\t" + "")) + (("a" + "") + ""))) {
            case "foo bar":
                v8 = v8
            }
            fmt.Println(v10, v13)
            _ = fi14
        }
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
