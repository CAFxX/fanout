package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []string{"hello"}
    var v2 bool = true
    var v3 int64 = 9223372036854775807
    var v4 int64 = v3
    var v5 uint8 = 42
    v6 := make(map[string]bool)
    v7 := GenS1{F0: 42}
    var v8 uint16 = 830
    if true {
        var v9 uint = 18446744073709551615
        _ = len(v1)
        fmt.Println(v2, v4)
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: -1}
        v6["a"] = ("" >= "")
        _ = len(v1)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v5, fi12)
        _ = fi12
    }
    v8 += (v8 & v8)
    if (!v2) {
        var v13 uint64 = 862
        v8 = (v8 >> ((4294967295) & 63))
        v14 := GenS1{F0: 1}
        _ = v13
        _ = v14
    } else {
        v1 = append(v1, "")
    }
    v6[""] = v2
    v2 = (0.0 >= 0.1)
    switch 1 {
    case -2147483648:
        var v15 int64 = (v4 << ((1) & 63))
        _ = v15
    case 9223372036854775807:
        v16 := GenS1{F0: 1}
        _ = v16
    case 569:
        v4 += v4
    default:
        fmt.Println(v2)
    }
    var v17 uint16 = 65535
    switch 32767 {
    case 32767:
        v2 = (4294967295 != 42)
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            v1 = append(v1, ("a" + ("\n\t" + (("foo bar" + "a") + ("\n\t" + "\n\t")))))
            var v20 bool = false
            v21 := GenS1{F0: 0}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            delete(v6, "hello")
            var v24 uint16 = v8
            v8 = v17
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v25 int64 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
