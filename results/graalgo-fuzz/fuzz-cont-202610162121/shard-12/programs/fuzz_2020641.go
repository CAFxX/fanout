package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 881
    var v2 int8 = 1
    v3 := []uint32{87}
    if false {
        switch ("\n\t" + "\n\t") {
        case "hello":
            if len(v3) > 0 {
                v3[0] = 730
            }
        default:
            v4 := GenS1{F0: 1}
            _ = v4
        }
        fmt.Println(v2, v1)
    }
    v5 := GenS1{F0: -1}
    v6 := GenS1{F0: -9223372036854775808}
    if (192 == 255) {
        if true {
            fmt.Println(v2, v2)
            v1 -= v1
            for fi7 := 0; fi7 < 16; fi7++ {
                fmt.Println(v2)
                fmt.Println(fi7, v1)
                if false {
                    v8 := GenS1{F0: 280}
                    switch "a" {
                    case "\n\t":
                        var v9 int = fi7
                        _ = v9
                    case "foo bar":
                        var v10 int16 = -1
                        _ = v10
                    }
                    v2 += v2
                    _ = v8
                }
                _ = fi7
            }
        }
        if false {
            var v11 int16 = 1000
            var v12 bool = false
            _ = v11
            _ = v12
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v13 := GenS1{F0: -9223372036854775808}
        _ = v13
    }
    v2 += v2
    v3 = append(v3, 898)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v1 & v1) {
    case -9223372036854775808:
        v2 = -44
    }
    var v14 uint = 42
    v3 = append(v3, 81)
    var v15 int16 = 1000
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 += (v15 >> ((v14) & 63))
    v2 = 127
    var v16 float64 = 3.14159
    var v17 int16 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
