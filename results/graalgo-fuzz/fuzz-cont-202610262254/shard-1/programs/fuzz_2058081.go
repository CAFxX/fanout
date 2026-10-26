package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int = -2147483648
    var v2 float64 = float64(float32(3.14159))
    var v3 uint32 = 786
    v4 := []int{}
    v5 := GenS1{F0: "hello"}
    _ = len(v4)
    v2 -= -2.25
    switch ("foo bar" + (("a" + "hello") + ("foo bar" + ("\n\t" + "")))) {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v6 int = v1
    for fi7 := 0; fi7 < 3; fi7++ {
        v2 -= 3.14159
        _ = fi7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v8 uint32 = v3
    if len(v4) > 0 {
        v4[0] = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
