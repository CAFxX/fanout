package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }

func main() {
    var v1 int16 = 904
    var v2 uint32 = 313
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    var v5 int8 = 0
    var v6 int16 = v1
    switch 128 {
    case -842:
        v7 := []int{-9223372036854775808, 32768}
        _ = v7
    }
    if false {
        v1 = v6
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint16 = 747
            for fi10 := 0; fi10 < 20; fi10++ {
                switch int(v6) {
                case 32767:
                    v6 += v6
                }
                if true {
                    v11 := GenS1{F0: "hello", F1: 0, F2: fi10}
                    var v12 bool = (2147483648 == 597)
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS1{F0: "foo bar", F1: 255, F2: -9223372036854775808}
                    fmt.Println(v6, v2)
                    v14 := GenS1{F0: "hello", F1: 128, F2: 32767}
                    _ = v13
                    _ = v14
                }
                switch (fi10 ^ 498) {
                case 4294967295:
                    if true {
                        if false {
                            fmt.Println(fi8, v6, v5)
                        }
                        var v15 int = (fi10 - (-(fi10)))
                        _ = v15
                    }
                case 9223372036854775807:
                    v4[v2] = (^(fi8))
                case 2147483647:
                    var v16 int16 = v1
                    _ = v16
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    }
    v17 := []int{127, 32768}
    var v18 uint64 = 18446744073709551615
    var v19 uint64 = v18
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 float64 = 1.5
        _ = fi20
        _ = v21
    }
    fmt.Println(v19, v1, v2)
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS1{F0: ("foo bar" + "\n\t"), F1: 42, F2: fi22}
        _ = fi22
        _ = v23
    }
    var v24 int16 = v6
    v25 := v4[(^(v2))]
    if true {
        v6 += (-(v6))
        v26 := GenS1{F0: "\n\t", F1: 128, F2: v25}
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 15; fi28++ {
            v18 = 4294967295
            _ = fi28
        }
        _ = len(v17)
        _ = len(v3)
        _ = fi27
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v25 = (^(v25))
        v24 += v24
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
