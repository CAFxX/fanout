package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }

func main() {
    v1 := []string{"a", "", (("" + ("a" + (("hello" + "\n\t") + (("\n\t" + ("a" + ("hello" + (("a" + ("foo bar" + "foo bar")) + "a")))) + "foo bar")))) + "\n\t"), "foo bar"}
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    for fi4 := 0; fi4 < 20; fi4++ {
        v3 -= ((float32(1e10) * -2.25) / (((v3 - v3) + (v3 - 3.14159)) + ((v3 / (-2.25 - (-2.25 - 3.14159))) * float32(((3.14159 / 1.5) / 3.14159)))))
        v3 = float32(3.14159)
        switch "\n\t" {
        case "\n\t":
            v3 += (3.14159 * 0.0)
        case "foo bar":
            var v5 int16 = 0
            _ = v5
        default:
            v1 = append(v1, "a")
        }
        _ = fi4
    }
    delete(v2, -282)
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: -1, F1: 130, F2: 2}
        _ = fi6
        _ = v7
    }
    v3 -= (0.1 * v3)
    switch 505 {
    case -69:
        var v8 int64 = -9223372036854775808
        _ = v8
    default:
        if true {
            v3 = v3
        }
    }
    v3 += 0.1
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := v2[fi9]
        _ = fi9
        _ = v10
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "a")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
