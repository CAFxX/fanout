package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    var v2 int = 2147483647
    v3 := []string{"foo bar", "\n\t", "foo bar", "a"}
    v4 := []string{("" + "foo bar"), (("a" + "\n\t") + (("a" + ("" + "a")) + "")), ("\n\t" + "")}
    var v5 int8 = 42
    var v6 bool = (42 >= 255)
    delete(v1, 4294967295)
    var v7 int16 = -32768
    v5 = (v5 - 100)
    var v8 string = "\n\t"
    v7 += 926
    var v9 uint8 = 60
    if len(v3) > 0 {
        v3[0] = (v8 + (v8 + v8))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
