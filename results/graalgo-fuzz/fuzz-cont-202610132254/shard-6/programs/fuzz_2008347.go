package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int, p1 uint64, p2 bool) uint16 {
    switch "a" {
    case "\n\t":
        p2 = (int32(p0) <= 450)
    case "a":
        var v4 string = "foo bar"
        _ = v4
    default:
        v5 := []int{(p0 / (((-(p0))) | 1)), (p0 + p0), -301, (p0 * p0)}
        _ = v5
    }
    v6 := GenS1{F0: uint32(p1)}
    var v7 int = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint = 4294967295
    var v2 int16 = -32768
    var v3 int8 = 1
    var v8 int = -735
    var v9 uint16 = genFn0((^(v8)), 136, true)
    v10 := []int{}
    v11 := GenS1{F0: 1}
    switch ("\n\t" + "a") {
    case "\n\t":
        if len(v10) > 0 {
            v10[0] = (-(v8))
        }
    case "a":
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                if false {
                    var v14 uint16 = v9
                    _ = v14
                }
                v15 := GenS2{F0: 42}
                v16 := GenS1{F0: 577}
                _ = fi13
                _ = v15
                _ = v16
            }
            _ = fi12
        }
    case "hello":
        switch "\n\t" {
        case "":
            v9 += 94
        case "hello":
            fmt.Println(v3)
        default:
            var v17 float64 = 3.14159
            _ = v17
        }
    }
    _ = len(v10)
    if ("a" < ((("a" + "a") + ("" + (("" + ("" + ((("\n\t" + "\n\t") + "foo bar") + ((("a" + ("hello" + "a")) + ("\n\t" + "hello")) + "foo bar")))) + "\n\t"))) + "foo bar")) {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        var v18 bool = true
        _ = v18
    } else {
        switch "foo bar" {
        case "foo bar":
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(v1, v1)
                switch ("hello" + "") {
                case "\n\t":
                    var v20 uint64 = 2
                    _ = v20
                case "a":
                    v21 := GenS2{F0: 65535}
                    _ = v21
                case "foo bar":
                    for fi22 := 0; fi22 < 15; fi22++ {
                        switch -105 {
                        case -9223372036854775808:
                            v10 = append(v10, fi19)
                        default:
                            v23 := GenS3{F0: (v3 + v3)}
                            _ = v23
                        }
                        fmt.Println(v9)
                        var v24 int64 = -622
                        _ = fi22
                        _ = v24
                    }
                }
                _ = fi19
            }
        case "a":
            fmt.Println(v8, v3)
        case "":
            v25 := GenS1{F0: 42}
            _ = v25
        }
        v9 += (v9 / ((255) | 1))
    }
    v26 := GenS4{F0: -122, F1: (65535 == 0)}
    for fi27 := 0; fi27 < 10; fi27++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
