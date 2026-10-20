package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 42
    var v3 uint32 = 779
    var v4 uint64 = v2
    for fi5 := 0; fi5 < 7; fi5++ {
        if true {
            v4 = v2
        }
        _ = fi5
    }
    if false {
        if true {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 string = "foo bar"
                var v8 int64 = -110
                v3 += v3
                _ = fi6
                _ = v7
                _ = v8
            }
            fmt.Println(v3, v4, v4)
        }
        var v9 float32 = float32(-2.25)
        _ = v9
    } else {
        v4 = (v2 | 4294967295)
        v10 := v1["\n\t"]
        v10 = v10
        _ = v10
    }
    switch ((("hello" + "") + ("foo bar" + "\n\t")) + ((("" + ("a" + "a")) + "hello") + "\n\t")) {
    case "\n\t":
        v2 = 18446744073709551615
    case "foo bar":
        var v11 uint16 = 32767
        _ = v11
    }
    v4 -= (v4 - 47)
    switch ("\n\t" + ("a" + "a")) {
    case "foo bar":
        v12 := GenS1{F0: 0, F1: 0}
        _ = v12
    }
    var v13 uint = 1
    v4 = v2
    switch 2147483647 {
    case 968:
        v14 := v1[""]
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
