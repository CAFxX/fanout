package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0() uint {
    var v6 int16 = 1000
    fmt.Println(v6)
    _ = v6
    return 500
}
func genFn1(p0 uint64) int {
    if false {
        var v7 int = 2147483647
        v8 := []int{(v7 + 128), v7, v7}
        var v9 string = "\n\t"
        _ = v7
        _ = v8
        _ = v9
    } else {
        if (375 <= 1000) {
            var v10 int8 = 1
            var v11 uint64 = (p0 ^ p0)
            _ = v10
            _ = v11
        } else {
            if ((240 != 126) && false) {
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := GenS1{F0: -32768, F1: 54}
                    var v14 int64 = -375
                    switch ((("" + ("" + ("foo bar" + "hello"))) + ("hello" + "\n\t")) + "foo bar") {
                    case "foo bar":
                        v15 := []int{(fi12 << ((909) & 63))}
                        _ = v15
                    case "a":
                        v14 = 42
                    default:
                        v14 = (v14 | -1)
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                fmt.Println(p0)
                p0 = (p0 & (p0 >> ((1) & 63)))
            }
            var v16 uint = 594
            _ = v16
        }
        var v17 int32 = -2147483648
        v17 -= (v17 + (v17 ^ -2147483648))
        _ = v17
    }
    switch (("foo bar" + "foo bar") + "hello") {
    case "":
        var v18 uint16 = 258
        _ = v18
    case "hello":
        var v19 uint16 = 42
        _ = v19
    case "\n\t":
        p0 += (p0 << ((genFn0()) & 63))
    }
    _ = p0
    return 32768
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"hello", "\n\t", ((("" + ("\n\t" + (("hello" + "") + ""))) + "hello") + ("hello" + ("hello" + "hello")))}
    v3 := make(map[string]bool)
    var v4 int64 = 42
    var v5 float32 = (3.14159 + float32(1.5))
    var v20 uint64 = 4294967295
    _ = len(v3)
    var v21 uint16 = 255
    var v22 int = 42
    if false {
        v4 = (v4 / ((int64(v21)) | 1))
        delete(v3, (("a" + "hello") + (("foo bar" + ("a" + ("" + ""))) + ("" + "hello"))))
    } else {
        fmt.Println(v21, v4, v21)
        v20 = (v20 % ((v20) | 1))
        v23 := GenS1{F0: -1, F1: -938}
        _ = v23
    }
    var v24 uint32 = 1
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 14; fi26++ {
            if false {
                fmt.Println(v21)
                fmt.Println(fi25, v22)
            }
            v24 += v24
            var v27 float32 = v5
            _ = fi26
            _ = v27
        }
        if false {
            v28 := GenS2{F0: (v24 / ((v24) | 1)), F1: -868, F2: 2147483647}
            _ = v28
        }
        _ = fi25
    }
    var v29 int = v22
    if (!((-495 > 42) || (((!(v21 == v21)) || false) && true))) {
        v29 = genFn1(v20)
    } else {
        _ = len(v1)
    }
    var v30 int32 = -2147483648
    v31 := GenS3{F0: v20}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
