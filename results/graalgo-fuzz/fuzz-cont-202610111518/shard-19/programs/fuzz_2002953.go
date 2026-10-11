package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int32, p1 int16) uint {
    if false {
        for fi5 := 0; fi5 < 4; fi5++ {
            if (-423 > fi5) {
                p1 += p1
                var v6 int32 = 42
                _ = v6
            }
            _ = fi5
        }
        var v7 int8 = -128
        var v8 int32 = 0
        _ = v7
        _ = v8
    }
    var v9 int16 = p1
    _ = p0
    _ = p1
    _ = v9
    return 164
}

func main() {
    var v1 float64 = 1.5
    v2 := []int{-2147483648}
    var v3 int32 = -2147483648
    var v4 uint16 = 773
    v4 = v4
    v4 = 32767
    var v10 int = 127
    v11 := GenS1{F0: 1}
    switch "a" {
    case "hello":
        fmt.Println(v4)
    case "foo bar":
        if false {
            v4 -= 70
            v1 = v1
        }
    case "\n\t":
        fmt.Println(v1, v3)
    }
    v12 := GenS1{F0: 42}
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        for fi13 := 0; fi13 < 12; fi13++ {
            v1 -= -2.25
            _ = fi13
        }
    case "a":
        _ = len(v2)
    }
    if true {
        var v14 uint8 = 1
        _ = v14
    } else {
        var v15 string = ""
        v3 += -725
        if (-964 <= 1) {
            var v16 int32 = -876
            var v17 uint16 = (v4 >> ((65535) & 63))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v16
            _ = v17
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
