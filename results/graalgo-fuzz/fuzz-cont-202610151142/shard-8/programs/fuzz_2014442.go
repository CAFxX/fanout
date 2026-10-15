package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }

func main() {
    var v1 int32 = 1
    v2 := []int{4294967295}
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint32 = 0
        _ = fi5
        _ = v6
    }
    v4[("hello" + ((("a" + ("foo bar" + "foo bar")) + "") + ((("a" + ("a" + "\n\t")) + ("" + ("" + ("a" + "foo bar")))) + "\n\t")))] = (1 <= 42)
    for fi7 := 0; fi7 < 13; fi7++ {
        v1 = v1
        v1 += -2147483648
        v1 -= -1
        _ = fi7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v1 -= (v1 * v1)
        v9 := v4[("" + "")]
        _ = fi8
        _ = v9
    }
    switch -2147483648 {
    case 32767:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -1:
        switch ("" + ("foo bar" + "foo bar")) {
        case "hello":
            v10 := GenS1{F0: false, F1: 649, F2: 1}
            _ = v10
        }
    default:
        delete(v4, "\n\t")
    }
    v4[""] = false
    fmt.Println(v1)
    v1 -= (v1 + 2147483647)
    v11 := GenS1{F0: (65535 <= 2147483648), F1: 402, F2: 32767}
    var v12 float64 = 0.1
    var v13 string = ("hello" + "a")
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(fi14)
        if true {
            v1 = (v1 >> ((184) & 63))
            v1 = v1
            v1 = 1
        }
        _ = fi14
    }
    delete(v4, "\n\t")
    fmt.Println(v12)
    for fi15 := 0; fi15 < 17; fi15++ {
        switch (fi15 - 32768) {
        case 644:
            _ = len(v2)
        case -2147483648:
            v2 = append(v2, (fi15 << ((717) & 63)))
        case -838:
            fmt.Println(v12, v13)
        default:
            var v16 uint32 = 2147483648
            _ = v16
        }
        v12 += v12
        _ = fi15
    }
    v1 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
