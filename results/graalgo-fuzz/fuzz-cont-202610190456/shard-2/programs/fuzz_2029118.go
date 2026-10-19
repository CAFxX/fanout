package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int64 = -957
    v2 := make(map[string]bool)
    v3 := []string{(("hello" + "a") + "a"), ("\n\t" + "hello")}
    v3 = append(v3, "foo bar")
    switch ((("hello" + "\n\t") + "a") + (("foo bar" + ("" + "hello")) + ((("hello" + ("a" + "hello")) + "a") + (("hello" + "foo bar") + "")))) {
    case "a":
        if true {
            v1 += (^(v1))
            var v4 int = 128
            _ = v4
        }
    case "hello":
        var v5 int = -2147483648
        _ = v5
    case "\n\t":
        v1 = (v1 & 1)
    }
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 1; fi7++ {
        v1 = v1
        _ = fi7
    }
    fmt.Println(v6, v1)
    var v8 uint64 = 931
    if ((float32((float64(0.1) - (3.14159 * (-2.25 * v6)))) >= 1.5) && (((3.14159 < 3.14159) && (65535 >= 551)) && (!(!false)))) {
        for fi9 := 0; fi9 < 20; fi9++ {
            v8 = v8
            _ = fi9
        }
        switch 4294967295 {
        case 440:
            v10 := GenS1{F0: v1}
            _ = v10
        }
    }
    switch 584 {
    case -319:
        v11 := GenS2{F0: 18446744073709551615}
        _ = v11
    }
    fmt.Println(v8)
    var v12 bool = (0 >= (^(v8)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v8
}
