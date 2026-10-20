package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint64; F1 bool; F2 int32 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint32 }

func main() {
    var v1 uint16 = 49
    v2 := []uint32{294, 234, 65535, 838}
    v3 := make(map[string]bool)
    var v4 uint16 = v1
    v4 = (v1 % ((32767) | 1))
    if ((883 <= 2) && (!false)) {
        if false {
            if (!(true && (!(-113 > 1)))) {
                fmt.Println(v4, v4)
            } else {
                var v5 uint64 = 1
                _ = v5
            }
        }
        v1 -= v4
        v6 := GenS1{F0: -1, F1: (("" + "a") + "foo bar"), F2: 322}
        _ = v6
    } else {
        if (!(true || false)) {
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 float64 = 0.1
                var v9 int8 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        v10 := GenS2{F0: 62, F1: ((("" + (("" + "hello") + "foo bar")) + ("" + (("" + ("foo bar" + "")) + ("" + ((("" + "\n\t") + (((("a" + "hello") + (("hello" + "") + "")) + "a") + "\n\t")) + (("foo bar" + "hello") + "foo bar")))))) + "foo bar")}
        v1 = (v4 << ((4294967295) & 63))
        _ = v10
    }
    _ = len(v2)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS3{F0: 278, F1: false, F2: 0}
        _ = fi11
        _ = v12
    }
    var v13 uint16 = v1
    fmt.Println(v13, v1)
    if len(v2) > 0 {
        v2[0] = 710
    }
    if false {
        var v14 int32 = 0
        v15 := GenS4{F0: (v14 | (v14 << ((205) & 63))), F1: 0, F2: 752}
        _ = v14
        _ = v15
    } else {
        v4 = v1
        v3["hello"] = false
        fmt.Println(v4, v4, v13)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
