package main

import (
    "fmt"
)


func main() {
    var v1 bool = (true || false)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    fmt.Println(v1)
    var v4 uint = 4294967295
    v3[("hello" + "")] = ""
    v1 = v1
    for fi5 := 0; fi5 < 5; fi5++ {
        _ = len(v2)
        _ = fi5
    }
    var v6 int32 = -991
    v6 = v6
    var v7 uint8 = 128
    var v8 int32 = (v6 >> (((v4 - (v4 & 18446744073709551615))) & 63))
    if (v4 <= v4) {
        v1 = (0 < 42)
    } else {
        var v9 float32 = 0.1
        v3["foo bar"] = ((("a" + ("\n\t" + "foo bar")) + "hello") + "")
        var v10 int16 = 1000
        _ = v10
        _ = v9
    }
    v11 := v2[("hello" + "foo bar")]
    v6 = v8
    var v12 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
