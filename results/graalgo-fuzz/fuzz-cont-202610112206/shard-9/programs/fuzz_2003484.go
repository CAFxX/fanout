package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() int16 {
    var v4 bool = true
    var v5 int16 = 645
    _ = v4
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 float64 = (3.14159 + 0.1)
    if (-546 != 445) {
        v3 = v3
        if (1 < 18446744073709551615) {
            fmt.Println(v3)
        }
    }
    v6 := GenS1{F0: 42}
    switch (("" + ("foo bar" + "")) + "a") {
    case "foo bar":
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint32 = 65535
            _ = fi7
            _ = v8
        }
    default:
        switch 32767 {
        case 4294967295:
            var v9 uint32 = 0
            _ = v9
        case 42:
            if true {
                v3 = 3.14159
                var v10 string = "hello"
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS1{F0: 188}
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
        default:
            v13 := GenS1{F0: 4294967295}
            _ = v13
        }
    }
    var v14 float64 = (v3 * 0.0)
    v3 = v3
    v15 := GenS1{F0: 0}
    v14 -= v14
    var v16 uint = 686
    v17 := GenS1{F0: 1}
    var v18 int16 = -958
    if (2 != 18446744073709551615) {
        v14 -= (float64((-2.25 * (0.0 - ((1e10 / 1.5) / float32(v14))))) + 3.14159)
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v18, v16)
            _ = fi19
        }
        v20 := GenS1{F0: uint32(v16)}
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        fmt.Println(v16, v18, v14)
        fmt.Println(v16, v3, v16)
        v14 = float64(1e10)
        _ = fi21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int = -9223372036854775808
        _ = fi22
        _ = v23
    }
    if false {
        var v24 int8 = -75
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v6
}
