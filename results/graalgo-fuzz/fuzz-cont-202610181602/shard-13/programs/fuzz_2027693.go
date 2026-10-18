package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string) string {
    p0 = p0
    switch p0 {
    case "":
        p0 = p0
    case "hello":
        if true {
            p0 = p0
        }
    case "foo bar":
        var v7 int64 = -705
        _ = v7
    default:
        p0 = p0
    }
    p0 = (p0 + (("a" + p0) + ("foo bar" + "")))
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = v8
    return (p0 + ("foo bar" + "hello"))
}

func main() {
    v1 := []string{"", ((((((("" + ((("a" + "hello") + "") + "")) + "a") + "hello") + "foo bar") + "a") + "\n\t") + ("foo bar" + (("a" + "foo bar") + ("foo bar" + "a"))))}
    v2 := make(map[int]bool)
    v3 := []string{(("\n\t" + ("" + "")) + "hello"), "a", ("foo bar" + "")}
    v4 := []string{"hello", ("foo bar" + ("hello" + "hello"))}
    var v5 uint = 463
    var v6 int64 = 9223372036854775807
    v9 := GenS1{F0: true, F1: 124, F2: 234}
    v5 = v5
    v2[32768] = (true || true)
    v5 = v5
    if false {
        var v10 int16 = 0
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        if false {
            var v11 bool = ((("foo bar" + ((("" + ("foo bar" + ("foo bar" + "\n\t"))) + "a") + "foo bar")) <= "a") && (254 >= 128))
            _ = v11
        }
        _ = v10
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: (165 > 128), F1: 58, F2: 65535}
            _ = fi12
            _ = v13
        }
    }
    var v14 int16 = 32767
    fmt.Println(v14, v6)
    v15 := GenS2{F0: 963}
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS1{F0: true, F1: 23, F2: 65535}
        v6 = v6
        if false {
            fmt.Println(v5)
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v5, v5, v5)
    var v18 float32 = (3.14159 - 3.14159)
    _ = len(v2)
    v14 += -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
