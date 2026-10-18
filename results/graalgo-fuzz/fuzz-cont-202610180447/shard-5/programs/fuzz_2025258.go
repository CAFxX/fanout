package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 255
    v2 := make(map[int]string)
    v3 := []string{(("a" + ("a" + ((("a" + (("hello" + "a") + "")) + ("a" + "\n\t")) + "a"))) + (((("\n\t" + ("a" + "\n\t")) + "\n\t") + "foo bar") + "")), ""}
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 float32 = 3.14159
            v1 = 0
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v7 float32 = 1.5
    switch 128 {
    case -192:
        var v8 int8 = 127
        _ = v8
    case 32768:
        fmt.Println(v7, v1)
    case 727:
        v7 -= float32((1.5 - float64(v7)))
    default:
        v7 = float32(1e10)
    }
    delete(v2, -848)
    fmt.Println(v1)
    if ((0 > -1) || true) {
        var v9 float64 = float64(0.1)
        _ = v9
    }
    switch ("\n\t" + "") {
    case "\n\t":
        v1 = (^(v1))
    }
    v1 += 127
    for fi10 := 0; fi10 < 17; fi10++ {
        v7 = (v7 - float32((1.5 / 1e10)))
        v1 -= (v1 - 42)
        _ = fi10
    }
    var v11 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
}
