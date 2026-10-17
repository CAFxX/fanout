package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
func genFn0(p0 int32) uint16 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{(fi5 | (-(fi5))), -1, fi5, -9223372036854775808}
        p0 = (p0 + -48)
        switch fi5 {
        case 215:
            p0 -= (p0 & (p0 | -1))
        case 1:
            if false {
                v7 := GenS1{F0: 18446744073709551615}
                p0 -= p0
                v6 = append(v6, fi5)
                _ = v7
            } else {
                var v8 uint8 = 128
                var v9 int32 = (p0 + 0)
                v10 := GenS1{F0: 918}
                _ = v10
                _ = v8
                _ = v9
            }
        }
        _ = fi5
        _ = v6
    }
    var v11 int8 = -71
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v11 = v11
            _ = fi13
        }
        var v14 float32 = 3.14159
        _ = fi12
        _ = v14
    }
    v15 := GenS1{F0: 4294967295}
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 bool = ((fi16 << ((4294967295) & 63)) >= fi16)
        p0 = p0
        for fi18 := 0; fi18 < 15; fi18++ {
            v11 -= -128
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v11
    _ = v15
    return 1
}
func genFn1(p0 string, p1 int) int64 {
    p1 = p1
    var v19 float32 = 1e10
    fmt.Println(v19, p1, p1)
    var v20 float32 = v19
    switch 592 {
    case -62:
        if false {
            v21 := GenS2{F0: 0}
            _ = v21
        }
    case -1:
        var v22 int8 = -101
        _ = v22
    default:
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 9223372036854775807
}

func main() {
    var v1 int32 = -1
    var v2 bool = false
    var v3 int = -918
    var v4 uint16 = 1
    switch ("a" + "\n\t") {
    case "a":
        var v23 int8 = 127
        _ = v23
    case "foo bar":
        var v24 string = "hello"
        _ = v24
    }
    switch (("\n\t" + ("foo bar" + (("hello" + "foo bar") + "foo bar"))) + "\n\t") {
    case "\n\t":
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 int32 = (^(v1))
            var v27 uint32 = 823
            var v28 string = (((((("hello" + "a") + ("hello" + "foo bar")) + ("a" + "\n\t")) + "hello") + "foo bar") + "a")
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
    case "a":
        var v29 uint16 = genFn0((^(v1)))
        _ = v29
    case "foo bar":
        if false {
            v30 := GenS2{F0: 42}
            _ = v30
        } else {
            v31 := GenS3{F0: v4, F1: 2, F2: ""}
            _ = v31
        }
    default:
        fmt.Println(v4)
    }
    var v32 uint = 451
    var v33 int = v3
    v34 := []int{127, (-(v3)), v3, (^(v3))}
    v1 = (v1 >> (((^(v32))) & 63))
    fmt.Println(v33)
    var v35 int32 = (v1 / (((v1 & (v1 | (v1 ^ (v1 << ((v32) & 63)))))) | 1))
    var v36 int64 = 42
    fmt.Println(v3, v2)
    v34 = append(v34, -985)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
