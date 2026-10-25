package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint32) int8 {
    v6 := GenS1{F0: 42, F1: uint8(p0), F2: "\n\t"}
    var v7 uint32 = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -41
}
func genFn1(p0 bool, p1 int) uint {
    if true {
        switch ("foo bar" + "foo bar") {
        case "":
            if true {
                var v8 int8 = 1
                v9 := GenS2{F0: 65535, F1: 2}
                _ = v8
                _ = v9
            } else {
                var v10 float64 = float64(-2.25)
                v10 = v10
                _ = v10
            }
        case "a":
            v11 := GenS3{F0: "foo bar", F1: 9223372036854775807}
            _ = v11
        case "hello":
            p0 = (65535 > 0)
        }
        v12 := GenS1{F0: -8, F1: 69, F2: "\n\t"}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS1{F0: 0, F1: 1, F2: "hello"}
            var v15 uint16 = 1
            fmt.Println(p0, v15)
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := []int{p1}
        _ = v16
    }
    switch "foo bar" {
    case "":
        switch "hello" {
        case "hello":
            var v17 bool = false
            _ = v17
        case "\n\t":
            if (!((-1 <= -71) || p0)) {
                var v18 int8 = 0
                var v19 uint8 = 193
                _ = v18
                _ = v19
            } else {
                v20 := []int{p1, p1}
                var v21 uint16 = 426
                fmt.Println(v21)
                _ = v20
                _ = v21
            }
        }
    case "\n\t":
        p0 = (2147483647 != -2147483648)
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        p0 = (!((("\n\t" + ("" + (("hello" + (("hello" + "") + "\n\t")) + "foo bar"))) + ((("a" + "a") + "") + "foo bar")) == ("foo bar" + "a")))
        _ = fi22
    }
    switch ("\n\t" + "\n\t") {
    case "":
        var v23 bool = ((255 > 526) && p0)
        _ = v23
    case "foo bar":
        v24 := GenS4{F0: 0}
        _ = v24
    case "\n\t":
        v25 := GenS3{F0: "\n\t", F1: 1}
        _ = v25
    default:
        p0 = p0
    }
    v26 := []int{}
    _ = p0
    _ = p1
    _ = v26
    return 18446744073709551615
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]int)
    var v3 int32 = -970
    v4 := make(map[uint32]int)
    v5 := make(map[string]int)
    v27 := []int{}
    var v28 float32 = float32(1e10)
    v3 += v3
    var v29 float64 = 1.5
    var v30 int32 = v3
    v1 -= (v1 ^ v1)
    var v31 string = ("" + "foo bar")
    if ((((v31 + v31) + "") > "hello") && true) {
        if true {
            if false {
                v29 = v29
                v30 = (-(v3))
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            }
            var v32 int32 = (v30 % ((v30) | 1))
            v29 = 0.0
            _ = v32
        }
        for fi33 := 0; fi33 < 17; fi33++ {
            v30 = -160
            _ = fi33
        }
        var v34 int16 = -1
        _ = v34
    }
    v28 -= float32(((v29 + 1e10) - v29))
    if false {
        v29 += (v29 + 0.1)
        _ = len(v27)
        var v35 int64 = (v1 & (v1 | v1))
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
