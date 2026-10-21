package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 uint16, p1 int8, p2 uint8) int8 {
    var v6 float64 = (0.1 * (-2.25 * -2.25))
    v7 := GenS1{F0: p1}
    switch "foo bar" {
    case "":
        v8 := GenS1{F0: (-(p1))}
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        if (42 < -1) {
            if false {
                p1 = (-(p1))
            }
            if false {
                for fi10 := 0; fi10 < 14; fi10++ {
                    if false {
                        if (!true) {
                            var v11 uint16 = 2
                            switch "foo bar" {
                            case "a":
                                var v12 uint = 4294967295
                                _ = v12
                            }
                            var v13 int64 = -681
                            _ = v11
                            _ = v13
                        }
                    } else {
                        v14 := GenS2{F0: "\n\t", F1: 42}
                        var v15 uint = 4294967295
                        fmt.Println(p2, p0, p0)
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                }
            } else {
                fmt.Println(p2, p2, fi9)
                for fi16 := 0; fi16 < 1; fi16++ {
                    if (!("hello" <= "\n\t")) {
                        v17 := []int{fi16, (-(fi9)), (fi16 ^ (-(fi9))), fi9}
                        _ = v17
                    }
                    _ = fi16
                }
            }
            v6 = v6
        } else {
            if true {
                p1 += p1
            }
            var v18 uint = 42
            _ = v18
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 int16) uint8 {
    var v19 int16 = p0
    v20 := []int{-706}
    var v21 uint = 65535
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    return 127
}

func main() {
    var v1 bool = ("a" != "a")
    var v2 int16 = 32767
    var v3 uint = 42
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    fmt.Println(v3, v3, v2)
    if (v1 && false) {
        var v22 float32 = 0.0
        _ = v22
    } else {
        v1 = true
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            if v1 {
                var v23 uint64 = 421
                _ = v23
            }
        case "foo bar":
            fmt.Println(v2, v2)
        }
    }
    var v24 string = "hello"
    v3 = (v3 >> (((v3 & (v3 % (((^(v3))) | 1)))) & 63))
    var v25 int64 = 984
    v1 = (42 >= 1)
    switch 32768 {
    case 32767:
        fmt.Println(v25, v25)
    }
    var v26 float64 = float64(0.1)
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS2{F0: (v24 + v24), F1: 2}
        switch 353 {
        case 42:
            v24 = (v24 + v24)
        case -419:
            _ = len(v4)
        case 2147483647:
            v24 = ((((v24 + "a") + ((v24 + v24) + v24)) + (v24 + "")) + v24)
        }
        fmt.Println(v26, v2)
        _ = fi27
        _ = v28
    }
    fmt.Println(v3, v25, v25)
    var v29 string = v24
    v24 = "a"
    v30 := []int{1}
    v30 = append(v30, -501)
    for fi31 := 0; fi31 < 19; fi31++ {
        v24 = (v29 + v24)
        var v32 bool = ((v3 * v3) > 529)
        _ = fi31
        _ = v32
    }
    var v33 uint32 = 1
    var v34 string = v24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
