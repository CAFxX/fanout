package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 string) int {
    switch -664 {
    case 1:
        for fi5 := 0; fi5 < 4; fi5++ {
            switch -1 {
            case -1:
                var v6 int = fi5
                _ = v6
            case 2147483647:
                if (1 != -977) {
                    if false {
                        var v7 uint64 = 407
                        var v8 uint8 = 255
                        _ = v7
                        _ = v8
                    }
                    v9 := GenS1{F0: 42, F1: 42}
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                if (32767 >= 42) {
                    p0 = ("a" + p0)
                    var v11 int32 = 0
                    v12 := GenS1{F0: 164, F1: (v11 | (v11 >> ((549) & 63)))}
                    var v13 int8 = 127
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    var v14 uint = 1
                    p0 = p0
                    if false {
                        v14 = (v14 << (((^(v14))) & 63))
                        p0 = p0
                    }
                    _ = v14
                }
                _ = fi10
            }
            p0 = "foo bar"
            _ = fi5
        }
    case -9223372036854775808:
        var v15 uint32 = 2147483647
        _ = v15
    }
    p0 = "foo bar"
    p0 = (p0 + ((p0 + p0) + "\n\t"))
    var v16 uint64 = 511
    _ = p0
    _ = v16
    return -9223372036854775808
}
func genFn1(p0 int8, p1 uint, p2 uint8) uint64 {
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint16 = 946
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(p1)
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    var v20 uint8 = 48
    var v21 uint16 = 65535
    v22 := GenS2{F0: genFn0("hello")}
    v23 := GenS3{F0: 42, F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    return 188
}

func main() {
    var v1 float64 = (-2.25 - 0.0)
    v2 := []uint32{0, 0, 4294967295}
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    for fi24 := 0; fi24 < 20; fi24++ {
        if (42 > 42) {
            v1 = float64(1e10)
            var v25 uint = 431
            v26 := GenS3{F0: 800, F1: "a"}
            _ = v25
            _ = v26
        } else {
            if false {
                var v27 uint64 = 42
                _ = v27
            } else {
                v1 -= v1
                v1 += v1
                fmt.Println(v1)
            }
        }
        var v28 int16 = -32768
        _ = fi24
        _ = v28
    }
    delete(v4, "\n\t")
    var v29 int = 2147483647
    v1 -= v1
    v1 += v1
    for fi30 := 0; fi30 < 7; fi30++ {
        v4["hello"] = genFn0("hello")
        _ = fi30
    }
    switch (v29 * v29) {
    case 1:
        v1 += -2.25
    case 127:
        v31 := v4["hello"]
        _ = v31
    }
    fmt.Println(v29, v1)
    fmt.Println(v1)
    switch ("\n\t" + ("\n\t" + "a")) {
    case "\n\t":
        v29 += v29
    case "hello":
        v1 = (0.1 / ((v1 + (v1 * ((v1 - v1) / (v1 * v1)))) / 1.0))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
