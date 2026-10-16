package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    var v1 int64 = 90
    var v2 int8 = 111
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    v5 := GenS1{F0: 1}
    v6 := GenS1{F0: 1}
    var v7 uint32 = 2147483647
    var v8 float64 = 1e10
    var v9 uint8 = 127
    delete(v4, v7)
    v4[65535] = (false && false)
    for fi10 := 0; fi10 < 13; fi10++ {
        v8 += v8
        _ = fi10
    }
    v7 = (^(v7))
    if (false || (uint16(v9) < 65535)) {
        v11 := GenS2{F0: (("\n\t" + "\n\t") + "foo bar"), F1: (v1 & v1)}
        if false {
            v1 = 42
            fmt.Println(v9)
        } else {
            fmt.Println(v1, v9)
        }
        _ = v11
    } else {
        if false {
            var v12 uint8 = v9
            var v13 int64 = (v1 + (v1 % ((v1) | 1)))
            if (true && (true || false)) {
                var v14 int16 = 461
                _ = v14
            } else {
                var v15 uint = 255
                _ = v15
            }
            _ = v12
            _ = v13
        } else {
            var v16 string = ((("hello" + ("foo bar" + (("foo bar" + ("\n\t" + ("a" + "a"))) + "hello"))) + (("\n\t" + ("hello" + "hello")) + (("a" + (("hello" + "hello") + "")) + "foo bar"))) + "\n\t")
            _ = v16
        }
    }
    v1 = v1
    var v17 int32 = 711
    v1 = -1
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            delete(v4, 65535)
            var v19 int16 = 1
            fmt.Println(v9)
            _ = fi18
            _ = v19
        }
        var v20 int16 = -1
        _ = v20
    } else {
        var v21 string = ("" + ("\n\t" + "hello"))
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
