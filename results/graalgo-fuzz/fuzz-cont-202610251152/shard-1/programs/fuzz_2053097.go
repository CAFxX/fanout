package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }

func main() {
    var v1 float64 = 3.14159
    v2 := []string{}
    v3 := []int{1, 595, -35}
    v4 := GenS1{F0: -32768, F1: ("foo bar" + (("foo bar" + "\n\t") + (("\n\t" + "") + "a")))}
    v1 = v1
    var v5 uint16 = 32767
    var v6 bool = false
    if v6 {
        for fi7 := 0; fi7 < 3; fi7++ {
            v5 -= 2
            switch ("a" + ("hello" + ("\n\t" + "hello"))) {
            case "foo bar":
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 int32 = 474
                    var v10 int16 = 1000
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "a":
                v1 = v1
            }
            v1 = 0.1
            _ = fi7
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v11 := GenS1{F0: -32768, F1: ""}
    if v6 {
        for fi12 := 0; fi12 < 19; fi12++ {
            v6 = true
            var v13 uint64 = 4294967295
            _ = fi12
            _ = v13
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v3) > 0 {
        v3[0] = 39
    }
    if (128 != 162) {
        v1 -= v1
        v1 = v1
        var v14 float64 = (float64(0.0) - (float64((0.1 * float32(1e10))) * 1.5))
        _ = v14
    } else {
        var v15 int = 42
        _ = v15
    }
    var v16 uint = 0
    switch 1 {
    case 0:
        v5 += (v5 & (v5 | 2))
    case -931:
        var v17 int32 = -447
        _ = v17
    default:
        fmt.Println(v6, v1, v1)
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        switch ((("foo bar" + "foo bar") + "\n\t") + "") {
        case "\n\t":
            var v19 uint = uint(v5)
            _ = v19
        case "hello":
            var v20 int64 = 1
            _ = v20
        case "foo bar":
            fmt.Println(fi18, v1, v5)
        default:
            for fi21 := 0; fi21 < 16; fi21++ {
                _ = len(v3)
                var v22 int32 = -1
                _ = fi21
                _ = v22
            }
        }
        v1 += 0.1
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
