package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int16; F2 uint64 }

func main() {
    var v1 uint64 = 42
    var v2 int8 = 1
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 int8 = v2
    if true {
        switch ("hello" + "hello") {
        case "":
            var v6 int = 786
            _ = v6
        case "\n\t":
            v7 := GenS1{F0: true}
            _ = v7
        case "hello":
            if false {
                v8 := GenS1{F0: false}
                for fi9 := 0; fi9 < 2; fi9++ {
                    fmt.Println(v5, v2)
                    _ = fi9
                }
                if false {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        for fi11 := 0; fi11 < 15; fi11++ {
                            fmt.Println(v2)
                            delete(v4, fi10)
                            _ = fi11
                        }
                        switch "hello" {
                        case "a":
                            if true {
                                v12 := GenS1{F0: false}
                                _ = v12
                            } else {
                                v13 := GenS2{F0: (^(v1))}
                                _ = v13
                            }
                        default:
                            fmt.Println(fi10)
                        }
                        _ = fi10
                    }
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v15 := GenS1{F0: (0.0 < float64(0.1))}
                        _ = fi14
                        _ = v15
                    }
                } else {
                    if false {
                        v5 = 127
                        var v16 int16 = 1000
                        _ = v16
                    }
                    switch (("foo bar" + "") + ("foo bar" + "a")) {
                    case "foo bar":
                        v2 = v2
                    case "a":
                        v17 := v4[-586]
                        _ = v17
                    }
                }
                _ = v8
            } else {
                _ = len(v4)
                var v18 string = (("a" + (("" + "") + "")) + "\n\t")
                if false {
                    v19 := GenS3{F0: -102}
                    _ = v19
                } else {
                    var v20 uint64 = 18446744073709551615
                    _ = v20
                }
                _ = v18
            }
        }
        var v21 int64 = -1
        _ = v21
    }
    delete(v4, 32767)
    v1 -= 2
    v22 := GenS4{F0: "hello", F1: -703, F2: v1}
    v23 := GenS4{F0: "\n\t", F1: 32767, F2: v1}
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int64 = 569
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        delete(v3, 699)
        v5 -= 42
        _ = fi26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v2 = -128
        _ = len(v3)
        _ = fi27
    }
    v2 = (v2 | (^(v5)))
    v28 := []int{4294967295, 9223372036854775807}
    v2 -= -128
    v2 = v5
    v29 := v4[-9223372036854775808]
    var v30 uint = 4294967295
    var v31 uint8 = 1
    v30 -= v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
