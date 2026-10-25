package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int64 {
    var v5 int16 = 605
    var v6 int64 = 58
    _ = v5
    _ = v6
    return (v6 ^ -508)
}

func main() {
    var v1 float64 = 3.14159
    var v2 int8 = 1
    var v3 uint16 = 255
    v4 := []string{"hello", ""}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    if ((224 != 128) && ((0.1 / ((-2.25 - (v1 * v1)) + (0.0 / (1.5 + v1)))) != v1)) {
        fmt.Println(v2, v1, v3)
        fmt.Println(v1, v3)
        v7 := GenS1{F0: -1}
        _ = v7
    } else {
        v8 := GenS1{F0: 1}
        if true {
            v1 = 0.0
            fmt.Println(v3)
        }
        v2 = (v2 - (-(v2)))
        _ = v8
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v9 float32 = (1e10 / -2.25)
    switch "\n\t" {
    case "a":
        switch 9223372036854775807 {
        case 42:
            v10 := GenS1{F0: 503}
            _ = v10
        default:
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        _ = len(v4)
                        _ = fi12
                    }
                    _ = fi11
                }
            }
        }
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = "hello"
        }
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int16 = 42
        _ = fi13
        _ = v14
    }
    if true {
        v9 += 3.14159
    } else {
        v3 = (v3 ^ 65535)
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if ((((float64((v9 * v9)) - (v1 / float64(3.14159))) + v1) / -2.25) >= (v1 - (((1.5 + v1) / v1) / ((0.1 * v1) + (v1 * v1))))) {
            var v16 uint = 781
            _ = v16
        }
        var v17 uint16 = v3
        _ = fi15
        _ = v17
    }
    v4 = append(v4, ((("" + "") + "\n\t") + "\n\t"))
    fmt.Println(v3, v1, v3)
    v2 += (-(v2))
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS1{F0: 555}
        v20 := GenS1{F0: 32767}
        v2 = v2
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float64 = v1
    v4 = append(v4, (((("foo bar" + "") + ("a" + (("\n\t" + "\n\t") + "hello"))) + ("a" + "")) + (("a" + ("\n\t" + ("foo bar" + "a"))) + ("a" + ""))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
