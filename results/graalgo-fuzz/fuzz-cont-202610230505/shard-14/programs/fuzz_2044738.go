package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }

func main() {
    var v1 int8 = 1
    v2 := make(map[string]string)
    var v3 uint16 = 222
    var v4 uint = 18446744073709551615
    fmt.Println(v4, v3)
    if (9223372036854775807 > -49) {
        fmt.Println(v4, v4)
    } else {
        if (!(859 == 9223372036854775807)) {
            if (true || (v4 != v4)) {
                if ((-1 != 0) && false) {
                    for fi5 := 0; fi5 < 20; fi5++ {
                        var v6 bool = (!true)
                        _ = fi5
                        _ = v6
                    }
                }
            }
        } else {
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: (("foo bar" + ((("a" + "foo bar") + "") + "foo bar")) + "hello"), F1: 0, F2: v1}
                var v9 string = ("hello" + ("\n\t" + "\n\t"))
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                v3 += v3
                var v11 bool = true
                _ = fi10
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint32 = 42
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        switch "" {
        case "hello":
            var v15 int8 = 127
            _ = v15
        case "a":
            v16 := []int{9223372036854775807, (fi14 * (fi14 ^ fi14)), fi14, 32768}
            _ = v16
        case "\n\t":
            v3 += v3
        default:
            v4 -= v4
        }
        fmt.Println(v4, fi14, v1)
        v4 = v4
        _ = fi14
    }
    fmt.Println(v4, v3, v4)
    for fi17 := 0; fi17 < 19; fi17++ {
        v1 -= v1
        _ = fi17
    }
    v3 = v3
    v4 -= v4
    if (!(((true && true) && (!(true || false))) || false)) {
        fmt.Println(v3, v3)
        v18 := v2["\n\t"]
        var v19 int8 = v1
        _ = v18
        _ = v19
    } else {
        var v20 int16 = 593
        v20 += v20
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
