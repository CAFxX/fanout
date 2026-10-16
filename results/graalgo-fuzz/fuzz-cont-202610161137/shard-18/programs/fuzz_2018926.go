package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 1
    var v2 float32 = 1e10
    var v3 int8 = 0
    var v4 uint64 = 553
    var v5 uint = 0
    fmt.Println(v3, v5)
    var v6 bool = (!false)
    v6 = v6
    v5 -= v5
    switch v1 {
    case -2147483648:
        var v7 int = v1
        _ = v7
    case 0:
        var v8 float64 = -2.25
        _ = v8
    case 32768:
        var v9 uint = 255
        _ = v9
    }
    fmt.Println(v2, v1)
    var v10 float32 = v2
    switch v1 {
    case 9223372036854775807:
        switch "foo bar" {
        case "\n\t":
            switch v1 {
            case 0:
                v11 := GenS1{F0: 0}
                _ = v11
            case 219:
                var v12 uint32 = 782
                _ = v12
            default:
                var v13 uint16 = 1
                _ = v13
            }
        case "foo bar":
            if v6 {
                var v14 uint16 = 65535
                _ = v14
            } else {
                if (!(("\n\t" == "foo bar") && (!v6))) {
                    var v15 bool = ((v5 | v5) > v5)
                    _ = v15
                } else {
                    switch v1 {
                    case -441:
                        if v6 {
                            var v16 uint16 = 2
                            _ = v16
                        } else {
                            fmt.Println(v1)
                        }
                    default:
                        fmt.Println(v2)
                    }
                    fmt.Println(v2)
                    var v17 uint8 = 42
                    _ = v17
                }
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS1{F0: 42}
                    var v20 int32 = 42
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        case "":
            var v21 int32 = 2147483647
            _ = v21
        }
    case 1:
        v22 := GenS1{F0: 1}
        _ = v22
    case -1:
        var v23 int16 = 42
        _ = v23
    default:
        v4 -= (^(v4))
    }
    var v24 int = -36
    var v25 uint8 = 1
    v3 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
