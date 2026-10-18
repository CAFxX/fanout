package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 string, p1 int8, p2 string) uint {
    var v6 float64 = 1e10
    var v7 int = 32768
    p0 = "foo bar"
    v8 := GenS1{F0: "", F1: -32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 uint8) bool {
    var v9 int32 = 978
    fmt.Println(p0)
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint = genFn0("hello", -128, "foo bar")
        if (42 < 42) {
            if false {
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 int32 = v9
                    p0 = 42
                    _ = fi12
                    _ = v13
                }
            }
            v9 = (v9 * v9)
        }
        v14 := GenS1{F0: "", F1: 1}
        _ = fi10
        _ = v11
        _ = v14
    }
    v15 := []int{}
    _ = p0
    _ = v15
    _ = v9
    return false
}

func main() {
    var v1 uint = 65535
    var v2 int64 = -9223372036854775808
    var v3 int64 = v2
    var v4 uint64 = 1
    v5 := []uint32{2147483648}
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: ("a" + ("a" + "foo bar")), F1: 1000}
        v1 -= (v1 & (v1 / ((v1) | 1)))
        switch ("foo bar" + "a") {
        case "hello":
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(v1)
                _ = fi18
            }
        case "foo bar":
            for fi19 := 0; fi19 < 10; fi19++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v20 int64 = -9223372036854775808
                var v21 int = fi16
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        _ = fi16
        _ = v17
    }
    var v22 int = 2147483647
    fmt.Println(v4, v4)
    v23 := GenS2{F0: v22, F1: false}
    fmt.Println(v3, v3, v2)
    switch "a" {
    case "":
        fmt.Println(v3, v4, v1)
    case "foo bar":
        var v24 uint16 = 255
        _ = v24
    case "hello":
        if ((2147483647 < 411) && (v22 <= -2147483648)) {
            v25 := GenS3{F0: 0, F1: false}
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 int = v22
                _ = fi26
                _ = v27
            }
            var v28 int64 = v2
            _ = v25
            _ = v28
        } else {
            var v29 bool = true
            fmt.Println(v22, v3, v3)
            _ = v29
        }
    }
    var v30 uint = 2
    var v31 uint32 = 2147483647
    if false {
        v31 = 2147483647
        var v32 string = "a"
        var v33 bool = (686 < (^(v4)))
        _ = v32
        _ = v33
    } else {
        v34 := GenS3{F0: 189, F1: true}
        _ = v34
    }
    var v35 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
