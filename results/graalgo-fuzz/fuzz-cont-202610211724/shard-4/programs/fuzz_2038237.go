package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 887
    var v2 int16 = v1
    v3 := []string{("" + "hello"), ("\n\t" + ("hello" + ("hello" + ("a" + (("\n\t" + ("a" + ("\n\t" + ("\n\t" + "\n\t")))) + "a")))))}
    var v4 float32 = float32(((float64(1e10) * (-2.25 / 1.0)) / 1.0))
    var v5 uint8 = 255
    v3 = append(v3, "a")
    for fi6 := 0; fi6 < 5; fi6++ {
        v2 = (v1 & (v1 + v1))
        fmt.Println(v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi6
    }
    v3 = append(v3, "")
    switch "hello" {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case "\n\t":
        v2 = v2
    case "a":
        for fi7 := 0; fi7 < 17; fi7++ {
            _ = len(v3)
            _ = fi7
        }
    }
    var v8 uint = 255
    v5 = 127
    v1 = -32768
    var v9 int = -539
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
