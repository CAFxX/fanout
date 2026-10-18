package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    v3 := make(map[string]int)
    var v4 int = 32768
    var v5 uint64 = 0
    var v6 int = (v4 << ((0) & 63))
    if (!(42 != 1)) {
        var v7 uint = 299
        _ = v7
    } else {
        v6 -= (-(v6))
        var v8 int32 = -1
        var v9 int64 = 42
        _ = v8
        _ = v9
    }
    v5 += 0
    var v10 float32 = 1.5
    var v11 int32 = -2147483648
    v11 = v11
    v10 = 3.14159
    v6 = int(v2)
    switch ("" + ("" + (("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) + "\n\t"))) {
    case "hello":
        v12 := GenS1{F0: 42}
        _ = v12
    }
    fmt.Println(v10, v5, v4)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int32 = v11
        _ = fi13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
