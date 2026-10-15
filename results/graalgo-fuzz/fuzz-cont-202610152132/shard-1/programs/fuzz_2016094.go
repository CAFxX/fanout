package main

import (
    "fmt"
)


func main() {
    var v1 float64 = 1e10
    var v2 string = "\n\t"
    var v3 string = ("\n\t" + (v2 + (("" + "\n\t") + (v2 + ((v2 + v2) + "a")))))
    v4 := []int{0, 32767}
    v5 := make(map[int]string)
    v3 = ("foo bar" + "foo bar")
    var v6 string = v2
    if (32767 == 1000) {
        v7 := v5[-1]
        _ = v7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v6 {
    case "\n\t":
        v6 = "foo bar"
    case "a":
        if len(v4) > 0 {
            v4[0] = -1
        }
    case "":
        var v8 uint32 = 0
        _ = v8
    }
    _ = len(v5)
    v3 = v6
    v9 := v5[-593]
    var v10 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
