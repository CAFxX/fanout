package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    var v3 float32 = (0.0 * 0.1)
    var v4 int = 9223372036854775807
    v5 := []int{(v4 * v4), (v4 & (v4 ^ (v4 & v4)))}
    switch (v4 - v4) {
    case -578:
        v1 = (v4 > v4)
    case 1:
        v1 = (!v1)
    }
    if (v1 || v1) {
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(v3)
        case 32768:
            var v6 int8 = int8(v4)
            _ = v6
        }
        var v7 bool = (!(867 >= 42))
        _ = v7
    } else {
        fmt.Println(v4, v3, v3)
        _ = len(v5)
        if (v4 >= v4) {
            var v8 int16 = -835
            _ = v8
        }
    }
    switch (("foo bar" + "") + "foo bar") {
    case "\n\t":
        if false {
            v3 += (-2.25 - v3)
        }
    case "foo bar":
        v1 = v1
    }
    v2[220] = 2147483648
    v1 = (!(854 == 0))
    switch ("hello" + "foo bar") {
    case "a":
        v9 := GenS1{F0: 966}
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 13; fi10++ {
            delete(v2, 0)
            v4 += fi10
            var v11 uint8 = 128
            _ = fi10
            _ = v11
        }
    default:
        switch v4 {
        case 9223372036854775807:
            switch "" {
            case "foo bar":
                v12 := GenS1{F0: 2}
                _ = v12
            }
        case 32768:
            var v13 int16 = 1
            _ = v13
        }
    }
    switch "foo bar" {
    case "hello":
        var v14 uint32 = 0
        _ = v14
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v3, v1, v4)
    var v15 string = (("\n\t" + "") + "a")
    var v16 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
