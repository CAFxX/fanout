package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 string; F2 bool }
type GenS4 struct { F0 uint32 }
func genFn0() int32 {
    v6 := GenS1{F0: 1, F1: 42, F2: 42}
    var v7 int64 = 9223372036854775807
    switch -1 {
    case -808:
        var v8 int16 = -1
        _ = v8
    case 32768:
        v9 := GenS1{F0: 121, F1: (v7 + v7), F2: 42}
        _ = v9
    case 137:
        if true {
            v10 := []int{9223372036854775807, 1, -2147483648}
            var v11 int32 = -2147483648
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 15; fi12++ {
                if true {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 uint8 = 6
                        var v15 int = -314
                        fmt.Println(v15)
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v7 = (v7 << ((255) & 63))
                    v16 := GenS1{F0: -43, F1: v7, F2: 42}
                    _ = v16
                } else {
                    var v17 string = "a"
                    v18 := GenS2{F0: v7}
                    v7 = (v7 & (v7 ^ v7))
                    _ = v17
                    _ = v18
                }
                _ = fi12
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                v7 -= -1
                _ = fi19
            }
            switch (("foo bar" + (("\n\t" + ("" + "hello")) + "")) + (("a" + ((((("foo bar" + "") + "foo bar") + "foo bar") + ("a" + "\n\t")) + (("\n\t" + "") + "\n\t"))) + "foo bar")) {
            case "foo bar":
                switch -650 {
                case 4294967295:
                    var v20 int32 = 1
                    _ = v20
                }
            }
        }
    }
    _ = v6
    _ = v7
    return 22
}
func genFn1() uint64 {
    var v21 uint32 = 9
    v22 := GenS1{F0: 19, F1: -753, F2: 0}
    v23 := GenS3{F0: 1, F1: "\n\t", F2: (!(1 > -11))}
    var v24 bool = false
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 1
}

func main() {
    v1 := []string{("hello" + ("a" + "foo bar")), ("" + "hello"), ("\n\t" + "hello")}
    var v2 uint32 = 2147483648
    var v3 int16 = 0
    var v4 uint16 = 138
    v5 := []string{"hello", "foo bar", "hello"}
    v25 := GenS3{F0: 127, F1: ("\n\t" + ""), F2: true}
    for fi26 := 0; fi26 < 12; fi26++ {
        _ = len(v5)
        v2 = (v2 | v2)
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int8 = 42
            fmt.Println(v4)
            v28 = 127
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    var v29 uint16 = (^(v4))
    switch (("hello" + ("" + ((("hello" + "hello") + ("a" + "foo bar")) + (("foo bar" + ("" + "\n\t")) + ("" + "\n\t"))))) + "hello") {
    case "a":
        v29 = (v4 - v29)
    default:
        v30 := GenS1{F0: -1, F1: -1, F2: 0}
        _ = v30
    }
    var v31 string = "a"
    v32 := GenS4{F0: v2}
    var v33 int64 = 0
    switch "hello" {
    case "\n\t":
        fmt.Println(v3, v33, v4)
    case "hello":
        var v34 int16 = (^(v3))
        _ = v34
    case "":
        _ = len(v1)
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        if false {
            var v36 uint32 = v2
            _ = v36
        }
        if (!false) {
            var v37 int16 = v3
            _ = v37
        }
        _ = fi35
    }
    if len(v1) > 0 {
        v1[0] = v31
    }
    v33 -= v33
    v4 -= 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
