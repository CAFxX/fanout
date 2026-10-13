package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 612
    var v3 uint32 = (v2 / ((v2) | 1))
    v3 += v2
    v3 = v3
    var v4 uint16 = 1
    v1["\n\t"] = ("foo bar" + ("a" + "hello"))
    if ((v4 * 220) == v4) {
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: 32767}
            var v7 uint8 = 0
            v8 := GenS1{F0: -1}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 42}
        var v10 float32 = 1.5
        _ = v10
        _ = v9
    } else {
        var v11 bool = ((("foo bar" + "") < ("a" + "foo bar")) && (2 <= 2))
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float32 = 1.5
        _ = fi12
        _ = v13
    }
    v14 := []int{42, 32768}
    switch "hello" {
    case "\n\t":
        fmt.Println(v2, v2)
    case "":
        if (-32768 <= 903) {
            v15 := GenS2{F0: v2}
            var v16 bool = false
            _ = v15
            _ = v16
        } else {
            v2 = v2
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 uint = 427
                var v19 string = "foo bar"
                delete(v1, ("foo bar" + v19))
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
