package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 int64; F1 string; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS9 struct { F0 uint16; F1 string }
func genFn0(p0 string) uint16 {
    if true {
        if (-274 <= 32767) {
            var v7 uint32 = 2147483648
            v8 := []int{2147483647, 4294967295, 9223372036854775807, 0}
            fmt.Println(v7, p0)
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 270, F1: ("a" + p0), F2: 128}
        _ = v9
    } else {
        switch p0 {
        case "hello":
            var v10 uint8 = 160
            _ = v10
        case "\n\t":
            v11 := GenS2{F0: -357, F1: (("hello" + ("\n\t" + "hello")) + p0), F2: 1}
            _ = v11
        default:
            v12 := []int{}
            _ = v12
        }
        if ((32767 == -321) || true) {
            v13 := GenS3{F0: 2147483648}
            v14 := GenS4{F0: ((p0 + (p0 + (p0 + ((((("a" + (((p0 + p0) + "hello") + (((("hello" + (p0 + p0)) + ((((p0 + (p0 + "")) + "\n\t") + "hello") + p0)) + p0) + ""))) + "hello") + (p0 + "")) + p0) + "")))) + p0), F1: 4294967295}
            if true {
                for fi15 := 0; fi15 < 7; fi15++ {
                    switch (p0 + p0) {
                    case "":
                        p0 = "foo bar"
                    default:
                        var v16 int = fi15
                        _ = v16
                    }
                    _ = fi15
                }
            } else {
                v17 := []int{}
                v18 := GenS5{F0: 18446744073709551615}
                _ = len(v17)
                _ = v17
                _ = v18
            }
            _ = v13
            _ = v14
        }
    }
    var v19 uint8 = 0
    v20 := GenS6{F0: 1}
    for fi21 := 0; fi21 < 12; fi21++ {
        switch (p0 + "hello") {
        case "hello":
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 uint16 = 255
                var v24 uint = 1
                _ = fi22
                _ = v23
                _ = v24
            }
        case "":
            v25 := GenS7{F0: 1, F1: 1}
            _ = v25
        default:
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(fi21, p0)
                if (!false) {
                    v27 := []int{128, (fi21 << ((0) & 63))}
                    switch (("" + (p0 + (("" + "\n\t") + (p0 + (p0 + ("\n\t" + (p0 + p0))))))) + "foo bar") {
                    case "":
                        fmt.Println(fi21, fi26)
                    case "a":
                        v19 = (v19 % ((v19) | 1))
                    }
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                    _ = v27
                }
                p0 = p0
                _ = fi26
            }
        }
        v28 := GenS8{F0: -362, F1: 81, F2: 1}
        var v29 uint = 0
        _ = fi21
        _ = v28
        _ = v29
    }
    _ = p0
    _ = v19
    _ = v20
    return 32767
}
func genFn1(p0 uint32, p1 int16) string {
    for fi30 := 0; fi30 < 3; fi30++ {
        var v31 uint16 = genFn0(("hello" + "foo bar"))
        var v32 int32 = -2147483648
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int64 = -1
    _ = p0
    _ = p1
    _ = v33
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 172
    var v3 uint = 158
    var v4 int8 = -64
    var v5 uint = 405
    v6 := make(map[string]bool)
    v2 = (v2 / ((v2) | 1))
    v34 := v6[""]
    v4 = (-(v4))
    v35 := GenS9{F0: 284, F1: "a"}
    var v36 uint = v5
    delete(v6, (((("a" + ((("" + ("hello" + (("\n\t" + "a") + "hello"))) + "hello") + ((("" + "") + "\n\t") + ("hello" + "a")))) + ("" + "\n\t")) + ("\n\t" + (("hello" + "") + ("\n\t" + ("foo bar" + "hello"))))) + "\n\t"))
    v5 = (v5 * v36)
    for fi37 := 0; fi37 < 6; fi37++ {
        if (!v34) {
            v34 = (802 != uint32(v2))
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
