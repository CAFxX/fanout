package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint32 = 65535
    v2 := []string{"a", "hello"}
    v3 := []string{}
    if false {
        switch "\n\t" {
        case "foo bar":
            v2 = append(v2, ("hello" + ""))
        default:
            if true {
                var v4 int8 = 0
                _ = v4
            }
        }
        var v5 int8 = -63
        var v6 int64 = -78
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 984}
        v1 = v1
        _ = v7
    }
    var v8 float64 = 3.14159
    v8 = v8
    _ = len(v2)
    v9 := GenS2{F0: -131}
    var v10 int16 = -32768
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS2{F0: 9223372036854775807}
        switch (fi11 + (fi11 << ((1) & 63))) {
        case 32767:
            fmt.Println(v1)
        case 678:
            var v13 uint64 = 18446744073709551615
            _ = v13
        default:
            fmt.Println(fi11, v10)
        }
        _ = fi11
        _ = v12
    }
    var v14 uint16 = uint16(v1)
    v3 = append(v3, ("hello" + ""))
    v1 = (^(v1))
    fmt.Println(v8, v10, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
