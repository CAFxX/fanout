package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 string = ""
    v2 := []uint32{}
    var v3 int8 = 74
    v4 := make(map[int]string)
    var v5 uint8 = 2
    var v6 float64 = (1.5 - ((float64(0.0) + 0.1) * float64((3.14159 - (float32(((1e10 * float64(0.1)) - (1.5 / (1e10 / 1e10)))) - 1.5)))))
    v3 = (v3 + (v3 % ((v3) | 1)))
    for fi7 := 0; fi7 < 7; fi7++ {
        if (true && (!(-32768 != -32768))) {
            if ((v5 * (v5 ^ v5)) < (v5 * v5)) {
                if true {
                    for fi8 := 0; fi8 < 2; fi8++ {
                        var v9 int16 = 830
                        _ = fi8
                        _ = v9
                    }
                    v1 = ""
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v10 uint64 = 752
                        var v11 uint16 = 42
                        var v12 uint64 = (v10 | (v10 ^ (v10 + v10)))
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                }
                v1 = "hello"
            }
            v4[(-(fi7))] = (v1 + "a")
        }
        _ = fi7
    }
    switch "foo bar" {
    case "foo bar":
        var v13 string = v1
        _ = v13
    case "hello":
        var v14 uint16 = 255
        _ = v14
    case "\n\t":
        if (true || false) {
            for fi15 := 0; fi15 < 14; fi15++ {
                v5 -= v5
                v5 = v5
                _ = fi15
            }
            v5 += 127
            if true {
                delete(v4, 128)
                fmt.Println(v5, v3)
            }
        }
    }
    var v16 uint = 65535
    var v17 float32 = 1e10
    var v18 int16 = 32767
    for fi19 := 0; fi19 < 12; fi19++ {
        v20 := GenS1{F0: 240}
        var v21 uint16 = 255
        v22 := GenS1{F0: 42}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    fmt.Println(v3, v6, v18)
    var v23 uint16 = 687
    delete(v4, 9223372036854775807)
    v18 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
