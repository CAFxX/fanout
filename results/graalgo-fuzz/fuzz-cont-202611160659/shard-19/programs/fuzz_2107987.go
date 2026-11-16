package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
func genFn0(p0 int8, p1 string, p2 string) int64 {
    fmt.Println(p0)
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 1
    var v3 int = -36
    v4 := make(map[string]bool)
    var v5 bool = false
    v6 := []string{("" + "a"), "hello"}
    var v7 uint8 = (v2 & v2)
    var v8 int = (-(v3))
    v2 -= v2
    switch ((("" + ("" + ("hello" + ("\n\t" + "\n\t")))) + "\n\t") + ("\n\t" + ("a" + (("foo bar" + "foo bar") + "hello")))) {
    case "foo bar":
        v3 = (v8 % ((v8) | 1))
    case "a":
        v6 = append(v6, "\n\t")
    case "\n\t":
        var v9 int = -9223372036854775808
        _ = v9
    default:
        var v10 string = ("foo bar" + "")
        _ = v10
    }
    var v11 int8 = -128
    var v12 uint16 = v1
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v3)
        _ = fi13
    }
    switch ("\n\t" + (("foo bar" + ("foo bar" + "")) + "a")) {
    case "foo bar":
        v14 := GenS1{F0: (("" + ("" + ("a" + "a"))) + ((("foo bar" + "") + ("a" + ("foo bar" + "\n\t"))) + "")), F1: ("\n\t" + "a"), F2: 1}
        _ = v14
    default:
        v7 = v2
    }
    fmt.Println(v2, v5)
    switch 32767 {
    case 1:
        v15 := GenS1{F0: "foo bar", F1: "", F2: (v12 << ((54) & 63))}
        _ = v15
    case 9223372036854775807:
        v16 := GenS1{F0: "", F1: "", F2: (v12 >> ((65535) & 63))}
        _ = v16
    default:
        v4["foo bar"] = (!(255 < 255))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
