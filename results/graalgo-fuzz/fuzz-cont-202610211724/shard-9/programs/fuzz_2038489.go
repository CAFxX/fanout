package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint8 }
func genFn0(p0 int32) uint64 {
    v4 := GenS1{F0: 9223372036854775807, F1: 138, F2: 479}
    v5 := GenS2{F0: ("" + "a")}
    var v6 uint = 4294967295
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 uint64) uint8 {
    switch -2147483648 {
    case 2147483647:
        if ((true && (!(42 != 18446744073709551615))) && false) {
            switch "\n\t" {
            case "a":
                for fi7 := 0; fi7 < 6; fi7++ {
                    v8 := []int{}
                    _ = fi7
                    _ = v8
                }
            default:
                v9 := []int{}
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := GenS1{F0: 42, F1: 4294967295, F2: 65535}
                _ = fi10
                _ = v11
            }
        }
    case 32767:
        p0 = genFn0(2147483647)
    }
    switch "foo bar" {
    case "\n\t":
        v12 := GenS1{F0: 9223372036854775807, F1: 42, F2: 255}
        _ = v12
    case "":
        p0 += p0
    }
    fmt.Println(p0)
    var v13 bool = false
    var v14 uint = 42
    _ = p0
    _ = v13
    _ = v14
    return 80
}

func main() {
    var v1 float64 = 0.1
    var v2 int = 2147483647
    v3 := make(map[uint32]string)
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint64 = 1
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(fi15)
            var v18 float32 = (((3.14159 + 1.5) - (-2.25 - 0.0)) - 1.5)
            var v19 string = ("foo bar" + "foo bar")
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    v20 := GenS1{F0: -185, F1: 0, F2: 255}
    var v21 string = "\n\t"
    delete(v3, 2147483648)
    v2 = (v2 | v2)
    var v22 bool = false
    if false {
        fmt.Println(v2, v21)
        var v23 uint64 = 598
        _ = v23
    } else {
        var v24 bool = (!(v22 && v22))
        _ = v24
    }
    v1 += (0.0 + v1)
    for fi25 := 0; fi25 < 7; fi25++ {
        if v22 {
            v3[0] = "foo bar"
            v22 = (!v22)
            v26 := []int{}
            _ = v26
        } else {
            v2 = (v2 - fi25)
            switch -171 {
            case 4294967295:
                fmt.Println(v2, fi25, v2)
            }
            v1 -= (1.5 + ((v1 / 1.0) * 3.14159))
        }
        delete(v3, 1)
        for fi27 := 0; fi27 < 13; fi27++ {
            fmt.Println(v21, v22)
            v28 := v3[1]
            _ = fi27
            _ = v28
        }
        _ = fi25
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        fmt.Println(v1)
        _ = fi29
    }
    v3[42] = "hello"
    var v30 float64 = v1
    switch v21 {
    case "\n\t":
        v31 := []int{}
        _ = v31
    }
    if true {
        var v32 int32 = 2147483647
        v33 := GenS2{F0: v21}
        if true {
            var v34 int8 = 100
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS3{F0: v32}
                if (-973 > 2) {
                    v37 := GenS4{F0: v34, F1: 0}
                    _ = v37
                }
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
}
