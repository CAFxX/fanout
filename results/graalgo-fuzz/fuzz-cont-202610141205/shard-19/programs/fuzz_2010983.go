package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8, p1 int8, p2 uint16) string {
    var v5 bool = (-962 != -2147483648)
    p1 = p0
    p0 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("foo bar" + "hello")
}

func main() {
    v1 := []int{}
    v2 := []string{((("a" + "") + "foo bar") + (("a" + "\n\t") + "\n\t")), ("foo bar" + "hello"), "foo bar", "a"}
    v3 := make(map[string]uint32)
    var v4 string = "hello"
    fmt.Println(v4)
    v6 := GenS1{F0: 90, F1: true}
    switch ("hello" + (v4 + v4)) {
    case "foo bar":
        switch 4294967295 {
        case 0:
            if false {
                var v7 bool = false
                v4 = ""
                v8 := GenS1{F0: -128, F1: (42 >= 1000)}
                _ = v7
                _ = v8
            }
        }
    case "":
        switch v4 {
        case "a":
            var v9 bool = (790 <= 32767)
            _ = v9
        default:
            var v10 uint8 = 2
            _ = v10
        }
    }
    switch 676 {
    case -301:
        delete(v3, "hello")
    case 0:
        v4 = v4
    }
    v4 = (("foo bar" + (("hello" + v4) + v4)) + v4)
    if (!((42 <= -2147483648) || (22 >= 2147483647))) {
        if false {
            var v11 bool = (1000 < -1)
            fmt.Println(v4)
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS2{F0: 1, F1: 1}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        v4 = "foo bar"
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint8 = 196
        _ = fi14
        _ = v15
    }
    v16 := GenS3{F0: 32767}
    var v17 uint16 = 42
    if (true && (false && ((!true) && true))) {
        v4 = "a"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
