package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int16, p1 int8) int64 {
    var v5 int = 2147483647
    v6 := GenS1{F0: p1}
    p1 += p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []string{(((("hello" + "foo bar") + "\n\t") + "") + ("\n\t" + ("\n\t" + (("\n\t" + "a") + "\n\t"))))}
    var v2 int = 90
    var v3 int64 = -469
    var v4 bool = false
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if (1e10 > 1.5) {
                var v9 uint8 = 42
                _ = v9
            } else {
                var v10 uint64 = 18446744073709551615
                _ = v10
            }
            fmt.Println(v3)
            _ = fi8
        }
        _ = fi7
    }
    var v11 int64 = v3
    v11 += (v3 >> ((313) & 63))
    v12 := GenS2{F0: "a"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: -40}
        v4 = (!(924 >= 0))
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: ("a" + "")}
    var v16 int64 = (v3 << ((1) & 63))
    if v4 {
        fmt.Println(v11, v4)
    }
    switch "hello" {
    case "foo bar":
        for fi17 := 0; fi17 < 20; fi17++ {
            if v4 {
                v18 := GenS2{F0: ("foo bar" + "foo bar")}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v19 uint = 0
                _ = v18
                _ = v19
            } else {
                v16 = (v3 ^ 42)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v20 uint32 = 2147483648
                _ = v20
            }
            _ = fi17
        }
    case "a":
        var v21 int64 = (v3 + v16)
        _ = v21
    }
    _ = len(v1)
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS2{F0: "a"}
        for fi24 := 0; fi24 < 11; fi24++ {
            if v4 {
                switch "foo bar" {
                case "":
                    v25 := GenS3{F0: 0, F1: 2}
                    _ = v25
                case "foo bar":
                    v16 = 0
                case "\n\t":
                    var v26 uint16 = 65535
                    _ = v26
                }
                _ = len(v1)
                v16 = (v3 | v11)
            } else {
                fmt.Println(v4)
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
            var v27 uint = 723
            fmt.Println(v2, v2)
            _ = fi24
            _ = v27
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            v11 = (v11 ^ (v11 + genFn0(0, -1)))
            v1 = append(v1, "\n\t")
            _ = fi28
        }
        _ = fi22
        _ = v23
    }
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
