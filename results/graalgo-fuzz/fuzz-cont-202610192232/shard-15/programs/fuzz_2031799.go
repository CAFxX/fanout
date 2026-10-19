package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int16) uint64 {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 894}
        p0 = (p0 + (p0 + (p0 & 0)))
        _ = fi7
        _ = v8
    }
    var v9 int64 = 1
    v9 = -409
    switch -592 {
    case 42:
        if (true && true) {
            var v10 int32 = 0
            _ = v10
        }
    case 897:
        var v11 bool = false
        _ = v11
    }
    v12 := GenS1{F0: 32767}
    _ = p0
    _ = v12
    _ = v9
    return 150
}
func genFn1() uint {
    var v13 uint8 = 244
    v13 = (v13 * 127)
    fmt.Println(v13)
    _ = v13
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 32768
    var v3 uint = 42
    var v4 string = ("\n\t" + ("" + "foo bar"))
    var v5 uint = (v3 - v3)
    var v6 int32 = 0
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint = v3
        v16 := GenS1{F0: 1}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 0}
    switch v4 {
    case "hello":
        v6 = (v6 * (v6 * v6))
    }
    var v18 int64 = -9223372036854775808
    v18 = 9223372036854775807
    v6 += v6
    var v19 uint32 = 42
    if false {
        if true {
            fmt.Println(v5)
        } else {
            _ = len(v1)
            v19 += v19
        }
    }
    if (210 == 2) {
        v19 = v19
        v20 := GenS1{F0: 42}
        var v21 int64 = v18
        _ = v20
        _ = v21
    } else {
        fmt.Println(v5)
        v3 = (v5 | (^(v5)))
    }
    var v22 uint32 = uint32(v3)
    v22 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
