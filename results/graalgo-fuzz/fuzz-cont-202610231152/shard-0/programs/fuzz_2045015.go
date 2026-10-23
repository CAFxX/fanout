package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
type GenS4 struct { F0 string; F1 int }
func genFn0(p0 string) uint64 {
    p0 = "a"
    v6 := GenS1{F0: 438}
    p0 = "hello"
    _ = p0
    _ = v6
    return 882
}

func main() {
    var v1 uint8 = 143
    var v2 uint32 = 2
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    var v5 int32 = int32(v2)
    switch 98 {
    case 42:
        v7 := GenS2{F0: 0, F1: v2, F2: (v1 | v1)}
        _ = v7
    case 2147483647:
        var v8 string = ""
        _ = v8
    }
    if (32768 != 2147483647) {
        var v9 uint8 = (v1 << ((2) & 63))
        switch ("hello" + "foo bar") {
        case "":
            var v10 int16 = 1
            _ = v10
        }
        var v11 uint64 = 508
        _ = v11
        _ = v9
    } else {
        v12 := GenS3{F0: 666, F1: 255, F2: "\n\t"}
        v1 -= 42
        delete(v3, "\n\t")
        _ = v12
    }
    if true {
        var v13 int64 = 386
        _ = v13
    }
    v14 := GenS3{F0: 0, F1: v1, F2: (((("" + "a") + ((((("foo bar" + (("hello" + "\n\t") + "hello")) + "") + ("\n\t" + ("" + ("foo bar" + ("foo bar" + "hello"))))) + "") + ("hello" + ("a" + ("hello" + "\n\t"))))) + "hello") + "a")}
    v2 = v2
    if false {
        var v15 uint16 = 487
        var v16 uint16 = v15
        switch ("" + ("hello" + ("\n\t" + "foo bar"))) {
        case "hello":
            switch "\n\t" {
            case "":
                var v17 int16 = 1
                _ = v17
            }
        }
        _ = v15
        _ = v16
    } else {
        v18 := GenS4{F0: ("foo bar" + "\n\t"), F1: 32767}
        var v19 uint16 = 1
        _ = v18
        _ = v19
    }
    fmt.Println(v2, v2, v1)
    fmt.Println(v2)
    var v20 string = ("\n\t" + (("foo bar" + ("foo bar" + "hello")) + ""))
    var v21 uint = 386
    var v22 uint64 = genFn0("a")
    var v23 uint = (v21 ^ (v21 * (v21 % (((^(v21))) | 1))))
    fmt.Println(v21)
    var v24 uint16 = 203
    v24 -= v24
    v2 = 2147483647
    v25 := GenS3{F0: 605, F1: uint8(v2), F2: (v20 + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
