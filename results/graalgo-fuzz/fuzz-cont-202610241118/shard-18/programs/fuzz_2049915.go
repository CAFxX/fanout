package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{128, 1, 127}
    var v3 uint16 = 65535
    fmt.Println(v3)
    fmt.Println(v3)
    if true {
        var v4 float32 = 3.14159
        _ = v4
    }
    if false {
        fmt.Println(v3)
    }
    fmt.Println(v3)
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            v3 -= (v3 % ((v3) | 1))
            _ = fi6
        }
        _ = fi5
    }
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(v3, fi7)
            if len(v2) > 0 {
                v2[0] = (fi7 ^ fi7)
            }
            _ = fi7
        }
        var v8 uint32 = 775
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(fi9, v3, v8)
            if false {
                fmt.Println(v3, v8)
            } else {
                var v10 bool = false
                _ = v10
            }
            for fi11 := 0; fi11 < 9; fi11++ {
                switch ((("foo bar" + "") + "\n\t") + "") {
                case "\n\t":
                    delete(v1, 0)
                }
                switch ("a" + ("hello" + ("\n\t" + "a"))) {
                case "a":
                    v8 = v8
                }
                _ = fi11
            }
            _ = fi9
        }
        _ = v8
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(fi12)
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS1{F0: 127, F1: 42}
                switch (fi12 - 128) {
                case -992:
                    _ = len(v1)
                case 127:
                    if true {
                        var v15 bool = true
                        if len(v2) > 0 {
                            v2[0] = (fi13 ^ (fi13 << ((4294967295) & 63)))
                        }
                        _ = v15
                    }
                    _ = len(v2)
                default:
                    v3 = (v3 << ((0) & 63))
                }
                v16 := GenS2{F0: 4294967295}
                _ = fi13
                _ = v14
                _ = v16
            }
            _ = fi12
        }
        var v17 bool = true
        _ = v17
    }
    v3 = 1
    fmt.Println(v3)
    if false {
        var v18 uint32 = 961
        _ = v18
    } else {
        var v19 string = ("\n\t" + "hello")
        _ = v19
    }
    v20 := v1[42]
    for fi21 := 0; fi21 < 9; fi21++ {
        if ((v20 + "a") < "hello") {
            v20 = ((v20 + "\n\t") + ((v20 + v20) + v20))
        } else {
            v22 := v1[2]
            _ = v22
        }
        var v23 uint32 = 101
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS3{F0: 1, F1: 2147483647, F2: (128 == (fi24 >> ((2) & 63)))}
            var v26 int16 = 36
            var v27 int64 = -9223372036854775808
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi21
        _ = v23
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 int = 32767
    delete(v1, 4294967295)
    v3 = (v3 - v3)
    var v29 uint16 = (v3 << ((65535) & 63))
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 int64 = -9223372036854775808
        _ = fi30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
}
