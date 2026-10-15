package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    var v3 string = ""
    v4 := []string{("" + v3)}
    if true {
        var v5 int8 = -34
        var v6 string = ((("foo bar" + v3) + "") + "")
        for fi7 := 0; fi7 < 20; fi7++ {
            v5 = v5
            switch "\n\t" {
            case "":
                v8 := GenS1{F0: (((v6 + v6) + ((v6 + v3) + v3)) + "foo bar"), F1: 0}
                _ = v8
            case "a":
                v3 = "foo bar"
            default:
                fmt.Println(fi7, v5)
            }
            var v9 uint16 = 0
            _ = fi7
            _ = v9
        }
        _ = v5
        _ = v6
    }
    v2 -= v2
    v1[v3] = v3
    var v10 float32 = 3.14159
    fmt.Println(v3, v2)
    v10 = v10
    var v11 int32 = v2
    var v12 int32 = (v2 >> ((856) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
