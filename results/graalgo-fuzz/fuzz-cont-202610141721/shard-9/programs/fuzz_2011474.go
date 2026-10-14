package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 int8, p1 int64) string {
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 bool, p1 uint) int32 {
    p0 = false
    var v8 int32 = -938
    if (p0 || (-128 > -128)) {
        var v9 string = ((("" + "foo bar") + (("a" + ("a" + "hello")) + ("foo bar" + ("hello" + "hello")))) + "foo bar")
        v9 = (v9 + "a")
        for fi10 := 0; fi10 < 9; fi10++ {
            v9 = v9
            _ = fi10
        }
        _ = v9
    } else {
        v11 := []int{}
        fmt.Println(p0, p0, v8)
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                fmt.Println(fi13, fi12, p1)
                _ = fi13
            }
            _ = fi12
        }
        _ = v11
    }
    if ("hello" < ("hello" + ((((("hello" + ("\n\t" + (("hello" + ("hello" + ("foo bar" + ""))) + "hello"))) + (("" + ("\n\t" + "")) + ("foo bar" + "hello"))) + ("" + "\n\t")) + ("hello" + ("hello" + "hello"))) + ("hello" + "\n\t")))) {
        switch "a" {
        case "\n\t":
            v14 := GenS1{F0: p1}
            _ = v14
        }
        p0 = (1 != 128)
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 string = ("foo bar" + "foo bar")
            p0 = p0
            var v17 int32 = v8
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 int32 = v8
        switch (((("hello" + "\n\t") + "\n\t") + "\n\t") + "hello") {
        case "hello":
            if true {
                fmt.Println(p0)
                fmt.Println(v19, p0, v8)
                v20 := GenS2{F0: 1}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v22 := []int{(-(fi18)), fi18, (-(fi18))}
                    v8 = 273
                    var v23 uint16 = 255
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
        case "":
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 string = ("a" + "\n\t")
                v26 := []int{fi18}
                fmt.Println(fi24, fi24)
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        switch ("hello" + "") {
        case "hello":
            v27 := GenS3{F0: p1, F1: uint16(p1)}
            _ = v27
        }
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    var v1 uint8 = 223
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := make(map[uint32]string)
    var v5 int64 = 1
    var v28 uint64 = 18446744073709551615
    var v29 float32 = (float32(0.0) + 0.1)
    switch ("hello" + "foo bar") {
    case "\n\t":
        v5 -= v5
    default:
        v5 += 0
    }
    var v30 int8 = 127
    v30 = (v30 << ((554) & 63))
    fmt.Println(v5)
    var v31 uint32 = 2
    v30 += (v30 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
