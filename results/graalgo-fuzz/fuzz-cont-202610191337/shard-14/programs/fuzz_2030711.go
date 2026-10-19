package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint8 = 2
    var v2 uint32 = 4294967295
    v3 := []string{(("" + ("hello" + "a")) + "\n\t")}
    for fi4 := 0; fi4 < 9; fi4++ {
        _ = len(v3)
        _ = fi4
    }
    v2 += 2
    var v5 uint16 = 466
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v6 uint64 = 18446744073709551615
    fmt.Println(v2, v2, v5)
    var v7 uint16 = v5
    var v8 int16 = -1
    v9 := GenS1{F0: (v8 - (v8 % (((v8 >> ((4294967295) & 63))) | 1)))}
    for fi10 := 0; fi10 < 12; fi10++ {
        switch (fi10 >> ((343) & 63)) {
        case 128:
            v8 = (^(v8))
        case -1:
            fmt.Println(v7)
        }
        _ = fi10
    }
    var v11 uint8 = (v1 / (((^(v1))) | 1))
    for fi12 := 0; fi12 < 7; fi12++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        fmt.Println(v5, v6)
        _ = len(v3)
        _ = fi12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int = fi13
        var v15 string = (("" + "foo bar") + "hello")
        v3 = append(v3, "")
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v1, v8, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
