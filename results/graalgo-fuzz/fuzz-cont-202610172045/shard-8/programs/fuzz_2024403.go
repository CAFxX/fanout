package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 uint64, p1 uint16) uint {
    fmt.Println(p0, p0)
    var v7 bool = true
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 820
    var v3 uint32 = 65535
    var v4 uint64 = 4294967295
    var v5 bool = true
    var v6 int = 32767
    v3 = v3
    var v8 uint8 = 127
    v9 := GenS1{F0: 0, F1: v6}
    if ((^(v3)) >= v3) {
        switch "foo bar" {
        case "a":
            switch v6 {
            case -301:
                v3 = v3
            case 128:
                var v10 int32 = 0
                _ = v10
            }
        }
        var v11 int64 = v2
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 string = (("hello" + ("foo bar" + ("foo bar" + "foo bar"))) + "\n\t")
            var v14 uint = 905
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        delete(v1, ("hello" + ""))
    }
    var v15 int32 = -752
    for fi16 := 0; fi16 < 9; fi16++ {
        if v5 {
            var v17 string = "a"
            fmt.Println(v2)
            _ = v17
        } else {
            v18 := GenS2{F0: (^(v2)), F1: 42}
            delete(v1, "foo bar")
            _ = v18
        }
        v19 := v1[("\n\t" + "hello")]
        v19 = ((v5 && v19) && v19)
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v6 += int(v15)
        v4 = v4
        _ = fi20
    }
    switch "hello" {
    case "foo bar":
        v5 = v5
    case "":
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 int = fi21
            _ = fi21
            _ = v22
        }
    }
    delete(v1, (("a" + "\n\t") + "a"))
    fmt.Println(v3, v3, v6)
    switch "foo bar" {
    case "\n\t":
        var v23 uint = 18446744073709551615
        _ = v23
    case "foo bar":
        v24 := GenS2{F0: (^(v2)), F1: -32768}
        _ = v24
    case "":
        v25 := GenS2{F0: -9223372036854775808, F1: -204}
        _ = v25
    default:
        v6 -= v6
    }
    v26 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
