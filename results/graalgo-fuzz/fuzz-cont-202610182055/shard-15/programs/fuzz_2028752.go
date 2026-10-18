package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int8 }
func genFn0() uint16 {
    var v4 uint32 = 2
    v5 := []int{-627, 97, 32768, -1}
    var v6 int8 = -128
    var v7 int64 = 588
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 int8, p1 int8) int32 {
    if (!(0 >= 0)) {
        p1 += (p0 ^ (p0 >> ((65535) & 63)))
    }
    switch 226 {
    case 1:
        p0 = (p1 - p0)
    case 32767:
        p1 = (p1 >> ((2) & 63))
    case 128:
        var v8 uint8 = 211
        _ = v8
    default:
        p0 = (p1 % ((p1) | 1))
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []int{-408, 127, 32768}
    var v2 uint32 = 2147483647
    var v3 uint8 = 91
    if ((v3 % (((v3 * 1)) | 1)) < (^(v3))) {
        v9 := GenS1{F0: false, F1: 4294967295}
        _ = v9
    }
    if (808 <= 2147483647) {
        v10 := GenS1{F0: false, F1: 4294967295}
        var v11 int = -635
        switch (("" + "") + "foo bar") {
        case "hello":
            if true {
                if (0.0 > -2.25) {
                    var v12 uint32 = 0
                    if false {
                        var v13 uint64 = 331
                        _ = v13
                    } else {
                        fmt.Println(v11, v2, v12)
                        var v14 uint32 = 2147483647
                        v15 := GenS2{F0: v3, F1: v3}
                        _ = v14
                        _ = v15
                    }
                    if (true || (!(false && (!((v11 ^ 42) == (v11 ^ (v11 | v11))))))) {
                        fmt.Println(v2, v11)
                        for fi16 := 0; fi16 < 1; fi16++ {
                            for fi17 := 0; fi17 < 15; fi17++ {
                                switch ((("a" + ("\n\t" + "\n\t")) + (("foo bar" + "") + "a")) + ("\n\t" + "hello")) {
                                case "foo bar":
                                    v12 = 0
                                default:
                                    v3 = (v3 / ((v3) | 1))
                                }
                                _ = fi17
                            }
                            _ = fi16
                        }
                    }
                    _ = v12
                } else {
                    v18 := GenS3{F0: (v2 > 4294967295)}
                    _ = len(v1)
                    _ = v18
                }
                v3 += v3
                var v19 int16 = 1
                _ = v19
            }
        case "\n\t":
            v3 = v3
        default:
            if true {
                v2 -= 65535
                for fi20 := 0; fi20 < 18; fi20++ {
                    fmt.Println(fi20, fi20, v11)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi20
                }
            } else {
                for fi21 := 0; fi21 < 9; fi21++ {
                    v22 := GenS3{F0: true}
                    var v23 int32 = -1
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = len(v1)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        _ = v10
        _ = v11
    }
    switch (("hello" + "foo bar") + "hello") {
    case "hello":
        v24 := GenS4{F0: (635 != 32768), F1: -128}
        _ = v24
    }
    v25 := GenS3{F0: (true || true)}
    switch "" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v26 string = ((("hello" + "\n\t") + "") + ("a" + ("hello" + ("" + "hello"))))
        _ = v26
    }
    v2 = 65535
    for fi27 := 0; fi27 < 17; fi27++ {
        v2 = 2
        v1 = append(v1, (fi27 / (((fi27 & fi27)) | 1)))
        _ = fi27
    }
    if (1 > genFn1(100, 100)) {
        var v28 string = "a"
        var v29 int16 = -1
        _ = v28
        _ = v29
    }
    switch "hello" {
    case "hello":
        fmt.Println(v3)
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
    }
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
