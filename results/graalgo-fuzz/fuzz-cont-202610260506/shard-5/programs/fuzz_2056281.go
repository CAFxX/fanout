package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 string = "foo bar"
    v4 := make(map[uint32]uint32)
    fmt.Println(v3)
    for fi5 := 0; fi5 < 16; fi5++ {
        if (((v3 + "\n\t") >= v3) || (3.14159 > 1.5)) {
            v2[v3] = ("foo bar" + "hello")
            for fi6 := 0; fi6 < 14; fi6++ {
                v3 = v3
                v2["\n\t"] = (((v3 + (("foo bar" + v3) + (v3 + v3))) + v3) + ((("hello" + (v3 + v3)) + v3) + "hello"))
                _ = fi6
            }
            var v7 int = fi5
            _ = v7
        } else {
            v3 = v3
            fmt.Println(fi5, v3)
            v3 = "a"
        }
        if true {
            fmt.Println(v3, v3)
        }
        v8 := GenS1{F0: 2147483648, F1: 42}
        _ = fi5
        _ = v8
    }
    fmt.Println(v3)
    switch ("hello" + ("\n\t" + v3)) {
    case "foo bar":
        v3 = "foo bar"
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v3 = v3
        delete(v1, v3)
        _ = fi9
    }
    if false {
        _ = len(v1)
        for fi10 := 0; fi10 < 10; fi10++ {
            v3 = ""
            v11 := []int{(fi10 * fi10)}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi10
            _ = v11
        }
    } else {
        var v12 int64 = -1
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint = 42
            var v15 int32 = 1
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: 874, F1: 255}
        _ = v12
        _ = v16
    }
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
