package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint32 = 2147483647
    var v2 bool = false
    var v3 string = "\n\t"
    v4 := []uint32{950, v1}
    switch int(v1) {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := GenS1{F0: 42}
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int8 = 42
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: -1}
            _ = fi5
            _ = v6
            _ = v9
        }
    }
    v10 := GenS1{F0: 0}
    for fi11 := 0; fi11 < 3; fi11++ {
        if true {
            var v12 string = "hello"
            _ = v12
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(v2)
                _ = fi13
            }
            fmt.Println(fi11)
        }
        var v14 string = ("hello" + (v3 + ((("foo bar" + "foo bar") + v3) + "\n\t")))
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v3, v3)
        switch 303 {
        case -9223372036854775808:
            fmt.Println(v2, v1, v3)
        case 904:
            var v16 uint16 = 32767
            _ = v16
        case -281:
            fmt.Println(v2, v2)
        }
        var v17 int32 = 42
        _ = fi15
        _ = v17
    }
    _ = len(v4)
    var v18 uint16 = 682
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int8 = 127
        _ = fi19
        _ = v20
    }
    var v21 string = v3
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
