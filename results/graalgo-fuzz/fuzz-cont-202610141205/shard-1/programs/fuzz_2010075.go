package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0() int64 {
    switch "\n\t" {
    case "":
        switch 1 {
        case 2147483647:
            switch "\n\t" {
            case "foo bar":
                var v6 uint16 = 651
                _ = v6
            }
        }
    }
    v7 := GenS1{F0: (-917 == 32768)}
    v8 := GenS2{F0: 0, F1: ((("hello" + "\n\t") + (("foo bar" + (("" + "foo bar") + ((("a" + "") + "a") + ("foo bar" + "foo bar")))) + "a")) + ("hello" + "\n\t")), F2: -1}
    if false {
        var v9 bool = true
        _ = v9
    }
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 uint16) int64 {
    switch 0 {
    case 42:
        var v10 bool = (71 > 65535)
        _ = v10
    case -121:
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: true}
            _ = fi11
            _ = v12
        }
    default:
        p0 = 2
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint64 = 558
        var v15 string = "foo bar"
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(fi16, p0, fi13)
            v14 -= v14
            v17 := []int{}
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        switch "hello" {
        case "hello":
            v19 := GenS3{F0: 2147483647, F1: genFn0(), F2: 42}
            _ = v19
        case "":
            v20 := []int{fi18}
            _ = v20
        case "foo bar":
            p0 += p0
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            switch (("a" + "") + (("a" + "hello") + "a")) {
            case "":
                if (!true) {
                    p0 = p0
                    v22 := GenS3{F0: 42, F1: -288, F2: 42}
                    _ = v22
                }
            case "hello":
                if true {
                    switch "hello" {
                    case "\n\t":
                        p0 = (p0 ^ (p0 | 255))
                    }
                }
            case "foo bar":
                v23 := []int{}
                _ = v23
            }
            var v24 uint16 = p0
            _ = fi21
            _ = v24
        }
        _ = fi18
    }
    var v25 uint16 = (p0 % ((42) | 1))
    p0 = (v25 >> ((65535) & 63))
    _ = p0
    _ = v25
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 42
    var v3 uint = 1
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    v26 := []int{-800}
    for fi27 := 0; fi27 < 1; fi27++ {
        fmt.Println(v3, v3)
        _ = len(v5)
        var v28 int32 = 0
        _ = fi27
        _ = v28
    }
    var v29 uint16 = 1
    v29 = (v29 / (((^(v29))) | 1))
    switch -9223372036854775808 {
    case -334:
        v30 := GenS1{F0: (!(4294967295 <= 4294967295))}
        _ = v30
    case 42:
        fmt.Println(v29, v2, v3)
    case -2147483648:
        switch ("" + ("foo bar" + "a")) {
        case "":
            var v31 uint8 = 127
            _ = v31
        default:
            delete(v5, "hello")
        }
    }
    v3 = v3
    v32 := v4[9223372036854775807]
    v33 := GenS2{F0: v2, F1: ("" + ""), F2: -1}
    var v34 uint8 = 127
    switch -2147483648 {
    case 32767:
        v2 -= -32768
    default:
        delete(v4, -1)
    }
    v35 := v4[0]
    v36 := GenS1{F0: (true && ((4294967295 != 547) && true))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
