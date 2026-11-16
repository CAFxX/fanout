package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int32; F2 int32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int64, p1 int) uint8 {
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(p0, p1)
        var v6 uint32 = 2
        if (!(v6 > v6)) {
            v7 := GenS1{F0: v6}
            var v8 uint64 = 42
            if true {
                fmt.Println(p1, v8, v8)
                v6 = v6
                v9 := GenS1{F0: (v6 + v6)}
                _ = v9
            }
            _ = v7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p0)
    var v10 string = "a"
    p0 -= -552
    p1 = 32767
    _ = p0
    _ = p1
    _ = v10
    return 79
}

func main() {
    var v1 bool = (!(!false))
    v2 := []uint32{4294967295, 4294967295, 2147483647, 4294967295}
    v3 := make(map[string]int)
    v4 := []uint32{493, 682}
    var v11 bool = v1
    v1 = (0 > 42)
    for fi12 := 0; fi12 < 16; fi12++ {
        switch "a" {
        case "\n\t":
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS2{F0: v1, F1: -1, F2: 815}
                var v15 int8 = 13
                _ = fi13
                _ = v14
                _ = v15
            }
        case "":
            for fi16 := 0; fi16 < 10; fi16++ {
                v17 := GenS3{F0: 255, F1: 0}
                switch (^(fi12)) {
                case -9223372036854775808:
                    if (("a" + "a") == ("hello" + "foo bar")) {
                        fmt.Println(v11, v1)
                    } else {
                        fmt.Println(fi16)
                    }
                case 389:
                    if v11 {
                        v18 := v3[""]
                        _ = v18
                    }
                case -1:
                    fmt.Println(fi12, fi16, fi16)
                }
                _ = fi16
                _ = v17
            }
        case "foo bar":
            v19 := GenS4{F0: 1, F1: 65535, F2: 2}
            _ = v19
        default:
            v20 := GenS5{F0: 481}
            _ = v20
        }
        _ = fi12
    }
    v21 := GenS4{F0: 2, F1: 1, F2: 174}
    switch "\n\t" {
    case "":
        if (32767 >= 42) {
            v1 = v11
            for fi22 := 0; fi22 < 6; fi22++ {
                for fi23 := 0; fi23 < 6; fi23++ {
                    fmt.Println(fi23)
                    var v24 uint32 = 2147483648
                    v1 = (1 == 4294967295)
                    _ = fi23
                    _ = v24
                }
                var v25 bool = (2 <= 4294967295)
                _ = fi22
                _ = v25
            }
        } else {
            v11 = (((!false) || v1) || (665 == 604))
            v11 = (0 <= 852)
            if len(v4) > 0 {
                v4[0] = 442
            }
        }
    }
    v11 = (!(!(!(1 <= 9223372036854775807))))
    if v1 {
        var v26 float32 = -2.25
        _ = v26
    }
    v1 = ((127 < -971) || ((("hello" + ("foo bar" + ("foo bar" + "hello"))) + "\n\t") >= "foo bar"))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
