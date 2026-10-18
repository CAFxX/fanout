package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }

func main() {
    var v1 float32 = (0.0 + 0.1)
    var v2 float64 = float64(v1)
    var v3 uint32 = 637
    v4 := make(map[string]bool)
    v5 := []int{32767}
    v1 = v1
    switch (("a" + "\n\t") + "") {
    case "a":
        fmt.Println(v1, v3)
    }
    if (2147483647 >= 42) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v6 := v4["foo bar"]
        for fi7 := 0; fi7 < 19; fi7++ {
            _ = len(v5)
            _ = fi7
        }
        _ = v6
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = 1
        _ = fi8
        _ = v9
    }
    v4[("" + "hello")] = false
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint32 = v3
            if (!true) {
                if true {
                    var v13 uint64 = 37
                    _ = v13
                } else {
                    var v14 float32 = 0.0
                    fmt.Println(fi10, v2, v3)
                    if true {
                        fmt.Println(fi10, v12)
                    }
                    _ = v14
                }
                var v15 uint32 = (^(v3))
                _ = v15
            }
            _ = fi11
            _ = v12
        }
        var v16 float64 = (0.1 - v2)
        v16 = ((v16 * v16) * (v16 + 1.5))
        _ = fi10
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 bool = (false && (-426 >= 9223372036854775807))
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 int8 = 42
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v21 := GenS1{F0: 0, F1: 90, F2: -55}
    var v22 int16 = 537
    for fi23 := 0; fi23 < 11; fi23++ {
        v22 += (v22 >> ((65535) & 63))
        if len(v5) > 0 {
            v5[0] = fi23
        }
        _ = fi23
    }
    v3 += v3
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
