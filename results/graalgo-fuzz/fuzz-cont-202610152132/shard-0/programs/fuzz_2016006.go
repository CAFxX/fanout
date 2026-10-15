package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{}
    var v2 uint32 = 0
    v3 := make(map[string]uint32)
    var v4 float64 = 3.14159
    var v5 uint16 = 65535
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = 1000
        _ = fi7
        _ = v8
    }
    var v9 int16 = -32768
    v6 = (^(v6))
    for fi10 := 0; fi10 < 2; fi10++ {
        v4 -= 1.5
        if (!true) {
            if true {
                var v11 int16 = (v9 >> ((65535) & 63))
                _ = v11
            }
            var v12 uint32 = v2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v12
        } else {
            for fi13 := 0; fi13 < 6; fi13++ {
                if true {
                    var v14 float32 = 3.14159
                    var v15 uint64 = 4294967295
                    if (v5 != uint16(v15)) {
                        v5 -= (v5 << ((uint(v2)) & 63))
                        v2 = 245
                    } else {
                        var v16 int32 = -2147483648
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v16
                    }
                    _ = v14
                    _ = v15
                }
                _ = fi13
            }
            _ = len(v3)
        }
        _ = fi10
    }
    if (uint32(v5) == (v2 | v2)) {
        switch "hello" {
        case "":
            fmt.Println(v6, v5, v6)
        default:
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v5 = v5
            }
        }
        v5 -= (v5 * v5)
    }
    var v17 int8 = 42
    switch "hello" {
    case "foo bar":
        delete(v3, "foo bar")
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        switch fi18 {
        case 2147483647:
            for fi19 := 0; fi19 < 16; fi19++ {
                if (v6 > (v6 & (v6 & (v6 - (v6 >> ((0) & 63)))))) {
                    v17 = v17
                }
                _ = fi19
            }
        default:
            var v20 bool = true
            _ = v20
        }
        switch (fi18 & 127) {
        case 4294967295:
            v21 := GenS1{F0: (v9 ^ v9)}
            _ = v21
        case 0:
            v22 := v3[("hello" + "foo bar")]
            _ = v22
        }
        _ = fi18
    }
    var v23 uint = uint(v6)
    for fi24 := 0; fi24 < 2; fi24++ {
        v9 = (-(v9))
        _ = fi24
    }
    switch 2147483647 {
    case 127:
        v25 := GenS1{F0: -1}
        _ = v25
    case 631:
        fmt.Println(v5, v23)
    }
    v3["foo bar"] = uint32(v6)
    v9 = (v9 * v9)
    v6 = uint8(v2)
    fmt.Println(v9, v2, v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
