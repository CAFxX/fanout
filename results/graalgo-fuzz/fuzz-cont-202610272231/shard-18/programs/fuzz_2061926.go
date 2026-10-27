package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 bool }
func genFn0() int64 {
    v7 := []int{590, 2147483647}
    v8 := GenS1{F0: 1, F1: 0, F2: (("foo bar" + "") <= ("\n\t" + ((("foo bar" + "\n\t") + "a") + ("hello" + "hello"))))}
    switch "hello" {
    case "hello":
        if (65535 == 0) {
            var v9 int64 = -1
            for fi10 := 0; fi10 < 8; fi10++ {
                v9 = v9
                var v11 bool = (!true)
                _ = fi10
                _ = v11
            }
            v7 = append(v7, 185)
            _ = v9
        }
    case "foo bar":
        var v12 int64 = -1
        _ = v12
    case "":
        var v13 uint32 = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 int = fi14
        switch "\n\t" {
        case "foo bar":
            v16 := GenS1{F0: 65535, F1: 42, F2: (true || (!false))}
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 32767
    var v3 uint64 = 18446744073709551615
    var v4 int16 = -1
    var v5 uint = 0
    v6 := make(map[int]uint32)
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int16 = (v4 & -32768)
        _ = fi17
        _ = v18
    }
    v2 = v2
    fmt.Println(v1, v5)
    var v19 int8 = 42
    var v20 uint32 = 2147483648
    var v21 uint64 = v3
    v22 := GenS1{F0: v20, F1: int64(v21), F2: (false && true)}
    fmt.Println(v5)
    fmt.Println(v19)
    delete(v6, int(v2))
    var v23 uint8 = 0
    fmt.Println(v23, v19, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
