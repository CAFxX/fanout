package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 633
    v3 := []string{"\n\t", "foo bar", "hello", "hello"}
    if (2147483648 == 105) {
        if len(v3) > 0 {
            v3[0] = ("" + "a")
        }
        for fi4 := 0; fi4 < 4; fi4++ {
            fmt.Println(v2)
            v3 = append(v3, ("" + "hello"))
            var v5 int8 = -88
            _ = fi4
            _ = v5
        }
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    }
    v2 = v2
    if len(v3) > 0 {
        v3[0] = ("hello" + ("foo bar" + "hello"))
    }
    var v6 uint16 = (v2 * (v2 << ((1) & 63)))
    _ = len(v3)
    v6 -= v2
    v7 := GenS1{F0: 976}
    v8 := GenS1{F0: -1}
    fmt.Println(v2, v2)
    switch ("\n\t" + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v9 uint32 = 0
        _ = v9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
