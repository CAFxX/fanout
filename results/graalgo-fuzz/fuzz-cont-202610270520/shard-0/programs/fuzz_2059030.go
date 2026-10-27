package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8) bool {
    if false {
        if true {
            var v7 int16 = 148
            _ = v7
        } else {
            fmt.Println(p0)
        }
        v8 := GenS1{F0: 128}
        p0 -= (^(p0))
        _ = v8
    }
    v9 := GenS1{F0: 32767}
    var v10 int8 = 9
    _ = p0
    _ = v10
    _ = v9
    return (!true)
}

func main() {
    v1 := []string{("" + ("hello" + "hello")), ("hello" + "foo bar"), "hello", ("\n\t" + "\n\t")}
    var v2 bool = (4294967295 <= 2147483647)
    var v3 uint16 = 32767
    v4 := make(map[string]bool)
    v5 := []uint32{}
    var v6 uint8 = 2
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "")
    }
    v6 = 127
    var v11 int32 = 0
    v3 -= 2
    fmt.Println(v11, v6)
    var v12 uint16 = (v3 ^ v3)
    fmt.Println(v2, v2, v11)
    var v13 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
