package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 bool = (2147483647 == 695)
    v2 := make(map[uint32]int)
    var v3 uint16 = 2
    var v4 int8 = 118
    var v5 string = "\n\t"
    var v6 string = (v5 + "hello")
    switch ((v5 + ("\n\t" + "")) + "a") {
    case "foo bar":
        fmt.Println(v6, v5, v6)
    case "hello":
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{(fi7 * (fi7 ^ (fi7 / (((fi7 * 0)) | 1)))), 127}
            _ = fi7
            _ = v8
        }
    default:
        v9 := GenS1{F0: 0}
        _ = v9
    }
    var v10 uint16 = v3
    fmt.Println(v1, v5, v3)
    var v11 uint32 = 479
    if v1 {
        var v12 uint64 = 1
        switch v5 {
        case "":
            v1 = (!(v10 == (v10 % ((1) | 1))))
        case "\n\t":
            var v13 string = v6
            _ = v13
        }
        _ = v12
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint8 = 6
            _ = fi14
            _ = v15
        }
        var v16 int32 = -1
        _ = v16
    }
    switch 42 {
    case 0:
        var v17 uint64 = 671
        _ = v17
    case -9223372036854775808:
        fmt.Println(v1)
    case 128:
        var v18 uint64 = 546
        _ = v18
    default:
        v19 := GenS1{F0: 1}
        _ = v19
    }
    switch (v6 + ("hello" + "")) {
    case "":
        var v20 int32 = 2147483647
        _ = v20
    case "\n\t":
        fmt.Println(v6, v5, v6)
    }
    fmt.Println(v1)
    fmt.Println(v5)
    v10 = v10
    var v21 int8 = (v4 + 127)
    switch 128 {
    case 128:
        fmt.Println(v6)
    case 32767:
        v22 := GenS1{F0: 42}
        _ = v22
    case -208:
        var v23 uint64 = 318
        _ = v23
    default:
        v6 = "foo bar"
    }
    var v24 int8 = v4
    for fi25 := 0; fi25 < 14; fi25++ {
        if (!(((!v1) && ((3.14159 + 1.5) != (3.14159 - ((-2.25 / 1.5) * (3.14159 / (-2.25 - 1e10)))))) && v1)) {
            var v26 uint32 = uint32(v3)
            v26 = (v26 + (^(v11)))
            _ = v26
        }
        if v1 {
            var v27 int8 = -25
            v28 := GenS1{F0: 255}
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 int32 = 2147483647
                var v31 uint8 = 107
                v5 = v5
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        _ = fi25
    }
    var v32 uint16 = v10
    v33 := GenS1{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
