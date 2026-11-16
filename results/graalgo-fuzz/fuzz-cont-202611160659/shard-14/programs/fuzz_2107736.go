package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []int{32768, -2147483648, 9223372036854775807}
    var v2 bool = (!true)
    v3 := make(map[string]string)
    switch "hello" {
    case "foo bar":
        var v4 float32 = float32(((0.1 / 1.5) + (3.14159 + float64(1e10))))
        _ = v4
    }
    var v5 int16 = -32768
    var v6 uint16 = 65535
    fmt.Println(v5, v2)
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint8 = 1
            _ = fi7
            _ = v8
        }
        if ((1.5 * 1.5) <= 0.0) {
            v6 += v6
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 float32 = 1e10
                _ = fi9
                _ = v10
            }
            var v11 uint8 = 0
            _ = v11
        } else {
            var v12 int8 = 100
            var v13 int = -9223372036854775808
            _ = v12
            _ = v13
        }
    } else {
        v6 -= (v6 << ((0) & 63))
        switch ("\n\t" + "\n\t") {
        case "hello":
            v2 = (!false)
        }
        v5 = (^(v5))
    }
    v5 += 32767
    switch ("" + "\n\t") {
    case "foo bar":
        v3["\n\t"] = ((("hello" + ("\n\t" + ("hello" + ""))) + "a") + "\n\t")
    }
    v5 -= (v5 * 1)
    switch 2147483647 {
    case 42:
        for fi14 := 0; fi14 < 8; fi14++ {
            v5 += 32767
            fmt.Println(fi14)
            _ = fi14
        }
    case 127:
        v2 = true
    }
    var v15 uint = 725
    var v16 int32 = 42
    v17 := GenS1{F0: 2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
