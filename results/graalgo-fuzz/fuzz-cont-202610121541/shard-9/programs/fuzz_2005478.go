package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 127
    v3 := make(map[uint32]string)
    var v4 uint8 = 0
    if false {
        v4 = (v4 << ((18446744073709551615) & 63))
    }
    delete(v3, 2147483647)
    var v5 uint16 = 1
    v6 := GenS1{F0: v5, F1: 103}
    if true {
        var v7 int8 = v2
        switch (("" + "a") + ("hello" + ("\n\t" + "hello"))) {
        case "\n\t":
            var v8 int = 127
            _ = v8
        case "a":
            fmt.Println(v2, v2)
        }
        _ = v7
    }
    var v9 float64 = (3.14159 * 3.14159)
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int64 = 815
        v12 := GenS1{F0: (^(v5)), F1: fi10}
        for fi13 := 0; fi13 < 6; fi13++ {
            if true {
                switch ("" + ("\n\t" + ("hello" + ("\n\t" + "foo bar")))) {
                case "a":
                    v14 := GenS1{F0: (v5 >> ((0) & 63)), F1: (fi10 >> ((uint(v5)) & 63))}
                    _ = v14
                }
            }
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v4 = v4
    v4 = (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
