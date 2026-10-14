package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint8 = 42
    var v2 int32 = 0
    var v3 uint8 = (v1 << ((65535) & 63))
    var v4 uint64 = 168
    var v5 float32 = ((-2.25 / ((-2.25 / float32(float64(1e10))) * 1e10)) + 1.5)
    v6 := []string{"a", ("" + ""), ("hello" + "foo bar"), ""}
    fmt.Println(v2, v3)
    switch 867 {
    case 32767:
        for fi7 := 0; fi7 < 3; fi7++ {
            _ = len(v6)
            _ = fi7
        }
    default:
        fmt.Println(v1, v5)
    }
    v6 = append(v6, "a")
    var v8 int64 = 1
    if false {
        v8 -= v8
        switch (("" + ("a" + "\n\t")) + ("foo bar" + (("a" + ("foo bar" + "\n\t")) + ("foo bar" + "hello")))) {
        case "hello":
            if (((false || (255 >= 0)) && true) || ("" <= "foo bar")) {
                var v9 int32 = (-(v2))
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 18; fi10++ {
                if len(v6) > 0 {
                    v6[0] = "foo bar"
                }
                fmt.Println(v4, v8, v1)
                _ = fi10
            }
        }
        v11 := GenS1{F0: 239, F1: 293}
        _ = v11
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v4 -= 1
        _ = fi12
    }
    var v13 float64 = float64(v5)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
