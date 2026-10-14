package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint = 916
    v2 := []int{-9223372036854775808, 32767, 0, -9223372036854775808}
    v3 := make(map[string]string)
    var v4 uint32 = 2147483648
    for fi5 := 0; fi5 < 7; fi5++ {
        v4 += v4
        _ = fi5
    }
    v6 := GenS1{F0: 2147483647, F1: 1}
    if ("foo bar" == ("hello" + ("\n\t" + ("hello" + "hello")))) {
        v1 += v1
        fmt.Println(v4, v1)
        if (!true) {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 int32 = 42
                var v9 uint16 = 1
                var v10 string = "hello"
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
    }
    v1 += (v1 % ((v1) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int16 = -32768
    var v12 int8 = 127
    if false {
        delete(v3, (((((("hello" + ("foo bar" + "\n\t")) + "\n\t") + "") + "foo bar") + "foo bar") + ("\n\t" + "a")))
        fmt.Println(v4, v12, v12)
        var v13 uint = v1
        _ = v13
    } else {
        fmt.Println(v4)
        if (v4 > (^(v4))) {
            for fi14 := 0; fi14 < 9; fi14++ {
                if ((v12 != v12) || (!((((!((!((58 > 128) || true)) || true)) && true) || (-32768 == 42)) && false))) {
                    _ = len(v3)
                }
                _ = fi14
            }
            v2 = append(v2, 127)
        }
    }
    switch 2147483647 {
    case -2147483648:
        var v15 int64 = 0
        _ = v15
    case 9223372036854775807:
        v1 += 4294967295
    default:
        v16 := GenS1{F0: 926, F1: 1}
        _ = v16
    }
    v11 = (v11 + v11)
    var v17 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
