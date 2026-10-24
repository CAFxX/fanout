package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 9223372036854775807
    var v3 int32 = -961
    var v4 float32 = 3.14159
    var v5 int = 586
    v6 := make(map[int]uint32)
    var v7 bool = false
    var v8 uint16 = 1
    if ((-57 <= 9) && v7) {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS1{F0: "", F1: 127}
            v11 := []int{}
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        var v12 uint = 65535
        v13 := GenS1{F0: "", F1: 96}
        var v14 uint32 = 2
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: ("" + ("a" + ("foo bar" + ("foo bar" + ("hello" + (("" + ("a" + ("a" + "a"))) + "\n\t")))))), F1: -99}
    v2 -= (v2 << ((18446744073709551615) & 63))
    switch (("\n\t" + "\n\t") + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v16 string = (("foo bar" + "\n\t") + "")
        _ = v16
    case "":
        v17 := GenS2{F0: 0, F1: 4294967295}
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := []int{v2, (v2 | (v5 | fi18)), (fi18 * (-(v5)))}
        v20 := GenS2{F0: 0, F1: 2147483648}
        v5 += v5
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint8 = 127
    switch (("\n\t" + ("\n\t" + "foo bar")) + "a") {
    case "a":
        v7 = ((v3 * v3) > v3)
    default:
        v7 = (0 <= 2)
    }
    var v22 string = ("foo bar" + "")
    if false {
        v7 = true
        var v23 uint8 = 159
        _ = v23
    } else {
        delete(v1, "a")
        if (v4 <= 0.0) {
            fmt.Println(v5, v3, v2)
            for fi24 := 0; fi24 < 19; fi24++ {
                v7 = v7
                _ = fi24
            }
            var v25 int = (v5 % ((v5) | 1))
            _ = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
