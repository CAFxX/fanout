package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 1.5
    v3 := []string{(("a" + "foo bar") + (("\n\t" + (("foo bar" + "") + ("\n\t" + (("\n\t" + ("foo bar" + "a")) + "\n\t")))) + "\n\t")), "\n\t", ("a" + (("\n\t" + "foo bar") + "hello"))}
    var v4 uint16 = 2
    var v5 float32 = 1.5
    var v6 uint8 = 122
    v2 -= v2
    v6 -= 1
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 4294967295}
        if (!false) {
            v9 := GenS1{F0: 65535}
            v2 = (v2 / v2)
            _ = v9
        }
        var v10 int8 = 0
        _ = fi7
        _ = v10
        _ = v8
    }
    v2 = (v2 * (v2 + 3.14159))
    if false {
        var v11 uint64 = 385
        v12 := GenS1{F0: 0}
        switch "" {
        case "\n\t":
            var v13 int32 = int32(v6)
            _ = v13
        case "foo bar":
            var v14 int64 = 9223372036854775807
            _ = v14
        }
        _ = v11
        _ = v12
    } else {
        for fi15 := 0; fi15 < 10; fi15++ {
            v4 -= (v4 >> ((42) & 63))
            _ = fi15
        }
        v1[2147483647] = (true || false)
    }
    var v16 uint16 = (v4 % (((v4 >> ((767) & 63))) | 1))
    for fi17 := 0; fi17 < 6; fi17++ {
        v16 = v4
        _ = fi17
    }
    var v18 bool = true
    if (!v18) {
        var v19 int32 = -2147483648
        var v20 bool = false
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 int32 = v19
            fmt.Println(v16, v4)
            var v23 int8 = 127
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v19
        _ = v20
    } else {
        fmt.Println(v2, v18, v2)
        for fi24 := 0; fi24 < 1; fi24++ {
            fmt.Println(v18, v16)
            v25 := GenS1{F0: 2147483648}
            var v26 uint16 = (v16 | v4)
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
