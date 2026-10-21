package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() string {
    var v6 bool = (42 >= 9223372036854775807)
    var v7 string = "hello"
    switch v7 {
    case "hello":
        v8 := []int{}
        _ = v8
    }
    var v9 uint32 = 55
    _ = v6
    _ = v7
    _ = v9
    return v7
}

func main() {
    var v1 int32 = 42
    var v2 uint16 = 1
    var v3 string = "a"
    v4 := []uint32{}
    var v5 uint64 = 4294967295
    v3 = ((v3 + (v3 + (("" + ((v3 + (v3 + v3)) + v3)) + "hello"))) + v3)
    v10 := GenS1{F0: "\n\t"}
    var v11 uint16 = (v2 << ((4294967295) & 63))
    if false {
        v3 = (("" + "hello") + "hello")
        var v12 int8 = 42
        _ = v12
    }
    v2 = v11
    var v13 int32 = v1
    var v14 int16 = 175
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if true {
        _ = len(v4)
        for fi15 := 0; fi15 < 8; fi15++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi15
        }
    } else {
        v11 = v2
        var v16 uint32 = 65535
        _ = v16
    }
    v17 := GenS1{F0: ((v3 + v3) + "foo bar")}
    _ = len(v4)
    v11 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
