package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int32 = 42
    var v2 int8 = 127
    var v3 uint = 1
    v4 := []string{"foo bar", "hello", ("hello" + "\n\t")}
    v5 := []int{-9223372036854775808}
    _ = len(v5)
    v3 = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 := GenS1{F0: 573}
    if true {
        if (!false) {
            var v7 string = "foo bar"
            _ = v7
        }
    }
    fmt.Println(v1, v1)
    if (!(!true)) {
        var v8 bool = false
        _ = v8
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
