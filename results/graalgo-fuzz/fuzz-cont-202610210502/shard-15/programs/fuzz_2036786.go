package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 int }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0(p0 string, p1 string, p2 int) int16 {
    var v4 int64 = -9223372036854775808
    v5 := GenS1{F0: 783, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1() uint {
    var v6 uint64 = 409
    var v7 uint16 = 65535
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []string{"a", "a", "\n\t", "\n\t"}
    v2 := []string{"\n\t", ((((("foo bar" + ("a" + "")) + (("\n\t" + "a") + ("foo bar" + ("" + "hello")))) + ("a" + "")) + "") + ("foo bar" + "hello")), "a", "hello"}
    v3 := []uint32{4294967295, 0, 0, 4294967295}
    v8 := GenS2{F0: 117, F1: 282}
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            var v10 uint8 = 2
            fmt.Println(fi9)
            _ = v10
        } else {
            _ = len(v2)
            var v11 uint8 = 128
            _ = v11
        }
        _ = fi9
    }
    var v12 int16 = -612
    var v13 string = ""
    v14 := GenS1{F0: 1, F1: 102}
    switch ("\n\t" + v13) {
    case "foo bar":
        fmt.Println(v12, v13)
    }
    var v15 uint16 = 32767
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    var v16 int64 = 0
    if true {
        if (!false) {
            fmt.Println(v13)
            v17 := GenS1{F0: 2, F1: -20}
            _ = v17
        }
    } else {
        var v18 int64 = (-(v16))
        _ = v18
    }
    fmt.Println(v15)
    var v19 int32 = 0
    switch v13 {
    case "hello":
        v19 = (v19 << ((847) & 63))
    default:
        _ = len(v2)
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        switch (v13 + ((("a" + "a") + v13) + v13)) {
        case "hello":
            var v21 int = int(v16)
            _ = v21
        case "a":
            var v22 uint = 4294967295
            _ = v22
        case "foo bar":
            var v23 uint8 = 53
            _ = v23
        }
        switch (("" + v13) + (("hello" + v13) + "\n\t")) {
        case "a":
            v12 += (v12 + (v12 | v12))
        case "foo bar":
            v24 := GenS3{F0: 899, F1: 1, F2: -2147483648}
            _ = v24
        case "hello":
            switch (fi20 / (((-(fi20))) | 1)) {
            case 42:
                v25 := GenS4{F0: v12, F1: 0}
                _ = v25
            case 2147483647:
                v2 = append(v2, "")
            }
        default:
            v12 += genFn0(("" + v13), "", (-(fi20)))
        }
        _ = fi20
    }
    v12 += (v12 | 167)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
