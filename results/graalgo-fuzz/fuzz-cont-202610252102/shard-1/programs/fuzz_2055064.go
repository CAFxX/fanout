package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 9223372036854775807
    v3 := []int{(^(v2))}
    v4 := make(map[string]int)
    var v5 bool = (!(1 >= 578))
    var v6 uint16 = 0
    for fi7 := 0; fi7 < 1; fi7++ {
        if v5 {
            v8 := v4[("\n\t" + ("\n\t" + ""))]
            _ = v8
        } else {
            var v9 int = fi7
            switch ("foo bar" + ("hello" + ((("a" + ("\n\t" + "hello")) + ("foo bar" + ("hello" + "\n\t"))) + (("foo bar" + "hello") + "hello")))) {
            case "":
                v10 := GenS1{F0: 18446744073709551615}
                _ = v10
            case "foo bar":
                var v11 uint8 = 128
                _ = v11
            default:
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 int16 = 335
                    _ = fi12
                    _ = v13
                }
            }
            _ = v9
        }
        if v5 {
            v2 += fi7
            fmt.Println(v2, v6, v5)
        }
        v5 = (v5 && v5)
        _ = fi7
    }
    var v14 int8 = 0
    switch "hello" {
    case "\n\t":
        v15 := GenS1{F0: 453}
        _ = v15
    case "":
        v6 = 0
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v16 uint = 520
        _ = v16
    }
    switch v2 {
    case -2147483648:
        v14 -= v14
    case 2147483647:
        switch ((("foo bar" + "\n\t") + "") + "hello") {
        case "\n\t":
            v6 = (v6 / ((v6) | 1))
        case "":
            var v17 uint = 42
            _ = v17
        }
    case 127:
        v14 -= v14
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    _ = len(v3)
    v2 -= 0
    switch v2 {
    case 127:
        if v5 {
            if ((((("foo bar" + "a") + "\n\t") + "\n\t") + (("hello" + "foo bar") + (("foo bar" + "foo bar") + ("\n\t" + ("hello" + "hello"))))) > ("hello" + ("a" + "hello"))) {
                delete(v4, ("hello" + "a"))
            } else {
                fmt.Println(v5)
                var v18 uint = 724
                _ = v18
            }
        } else {
            var v19 int8 = v14
            _ = v19
        }
    }
    v14 = 100
    v6 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
