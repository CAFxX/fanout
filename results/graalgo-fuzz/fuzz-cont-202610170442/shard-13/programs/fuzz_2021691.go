package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint; F2 uint16 }
type GenS7 struct { F0 int32; F1 int64; F2 uint8 }
func genFn0(p0 int, p1 int16) bool {
    v4 := GenS1{F0: 17, F1: 191}
    p1 -= p1
    v5 := []int{}
    if (true && (((1 >= 100) || (!(!(4294967295 > 161)))) || (p0 <= (p0 - 32767)))) {
        fmt.Println(p0, p1)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float32 = 1.5
    v3 := []string{("a" + ("hello" + "hello")), "", "hello", ("\n\t" + (((("\n\t" + ("a" + "")) + "foo bar") + "foo bar") + (("a" + "\n\t") + "foo bar")))}
    v6 := GenS2{F0: 0, F1: "a"}
    var v7 uint8 = 127
    switch ("foo bar" + "\n\t") {
    case "a":
        var v8 uint32 = (v1 ^ v1)
        _ = v8
    case "hello":
        v9 := GenS3{F0: v7, F1: 1}
        _ = v9
    }
    v10 := GenS2{F0: 42, F1: "a"}
    v11 := GenS4{F0: 32767, F1: 792}
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            v1 = (v1 + v1)
            v7 = v7
            var v14 int32 = 42
            _ = fi13
            _ = v14
        }
        v15 := GenS5{F0: 547, F1: v1}
        for fi16 := 0; fi16 < 16; fi16++ {
            if (784 != 42) {
                switch (("" + "\n\t") + "") {
                case "":
                    v2 -= v2
                case "foo bar":
                    var v17 int16 = 1000
                    _ = v17
                case "a":
                    switch "foo bar" {
                    case "foo bar":
                        var v18 bool = (!true)
                        _ = v18
                    default:
                        fmt.Println(fi16)
                    }
                default:
                    if false {
                        v7 = v7
                        if len(v3) > 0 {
                            v3[0] = ("" + (((((("a" + "") + "foo bar") + "\n\t") + "foo bar") + "a") + "a"))
                        }
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    } else {
                        if true {
                            v19 := GenS6{F0: 0, F1: 2, F2: 1}
                            _ = v19
                        }
                    }
                    var v20 int64 = 42
                    if false {
                        v20 = (-(v20))
                        var v21 int16 = -667
                        fmt.Println(v1, v2, v2)
                        _ = v21
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ("foo bar" + (((((("a" + "\n\t") + "") + "\n\t") + ("" + "a")) + ("hello" + ("hello" + (("foo bar" + ("" + "a")) + (("\n\t" + "") + ""))))) + ("hello" + ("" + ("\n\t" + "")))))
                        }
                    }
                    _ = v20
                }
            }
            _ = fi16
        }
        _ = fi12
        _ = v15
    }
    if false {
        var v22 string = "a"
        v23 := GenS6{F0: 1, F1: 18446744073709551615, F2: 65535}
        _ = v22
        _ = v23
    } else {
        v24 := GenS7{F0: -2147483648, F1: 698, F2: v7}
        _ = v24
    }
    var v25 uint16 = 53
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v6
    _ = v7
}
