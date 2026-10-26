package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint16 = 768
    var v2 string = ""
    var v3 uint64 = 1
    switch 42 {
    case 606:
        var v4 int64 = -1
        _ = v4
    }
    v5 := []int{-968, 32767}
    switch (v2 + "foo bar") {
    case "\n\t":
        if true {
            var v6 int32 = 2147483647
            var v7 uint8 = 255
            _ = v6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint8 = 128
        if len(v5) > 0 {
            v5[0] = int(v9)
        }
        var v10 int64 = -1
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = (v2 + (("foo bar" + "") + ((v2 + v2) + ("a" + v2))))
    if false {
        v11 := GenS1{F0: 65535}
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int16 = 30
            var v14 uint16 = (v1 | v1)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    fmt.Println(v2)
    v15 := GenS1{F0: 586}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
}
