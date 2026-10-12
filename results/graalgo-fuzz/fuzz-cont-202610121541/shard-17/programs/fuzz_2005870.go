package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint32 = 74
    v2 := []uint32{2147483648, (v1 * (v1 & v1))}
    var v3 int = 4294967295
    v4 := make(map[string]string)
    var v5 float64 = (-2.25 * 0.1)
    if (!true) {
        var v6 uint16 = 1
        delete(v4, ("foo bar" + "a"))
        _ = v6
    }
    fmt.Println(v5, v1, v5)
    if len(v2) > 0 {
        v2[0] = 877
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v8 float64 = (-2.25 + 0.1)
        fmt.Println(v1)
        _ = fi7
        _ = v8
    }
    var v9 string = "foo bar"
    var v10 float32 = 3.14159
    for fi11 := 0; fi11 < 6; fi11++ {
        if false {
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v5, v3, v9)
                _ = fi12
            }
            var v13 uint16 = 785
            _ = v13
        }
        fmt.Println(v5, v5, v9)
        fmt.Println(v9, fi11, fi11)
        _ = fi11
    }
    delete(v4, v9)
    v14 := GenS1{F0: 42}
    var v15 int8 = -128
    for fi16 := 0; fi16 < 8; fi16++ {
        v10 = v10
        if (!false) {
            var v17 uint = 65535
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: 149}
                _ = fi18
                _ = v19
            }
            _ = v17
        } else {
            v5 += float64(v10)
            v3 += 4294967295
        }
        _ = fi16
    }
    var v20 uint32 = (v1 | 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
