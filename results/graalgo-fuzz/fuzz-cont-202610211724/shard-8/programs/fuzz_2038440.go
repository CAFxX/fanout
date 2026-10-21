package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }

func main() {
    v1 := []string{("hello" + "a"), "\n\t"}
    var v2 uint16 = 32767
    var v3 int64 = -9223372036854775808
    var v4 int64 = v3
    v1 = append(v1, "hello")
    fmt.Println(v3)
    if ((43 <= (^(v2))) && false) {
        var v5 int16 = 1000
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: fi6}
            var v8 uint64 = 4294967295
            _ = fi6
            _ = v7
            _ = v8
        }
        v2 += v2
        _ = v5
    }
    v3 = (v4 << ((1) & 63))
    v4 = (v4 % ((1) | 1))
    v3 += 1
    v9 := GenS1{F0: 4294967295}
    switch "\n\t" {
    case "":
        v2 = (v2 % ((v2) | 1))
    case "hello":
        v10 := GenS1{F0: 42}
        _ = v10
    }
    v3 += (v4 % (((^(v3))) | 1))
    v11 := GenS2{F0: (("foo bar" + ("a" + "a")) + "foo bar")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
