package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint }
type GenS3 struct { F0 bool; F1 int32; F2 bool }
type GenS4 struct { F0 uint8; F1 uint16 }
func genFn0(p0 string) int16 {
    if (5 != 1) {
        if false {
            p0 = "\n\t"
        } else {
            switch ((p0 + p0) + p0) {
            case "foo bar":
                v5 := GenS1{F0: 0, F1: -2147483648, F2: (!((false && true) || true))}
                _ = v5
            }
            fmt.Println(p0)
        }
        if true {
            if true {
                p0 = ("" + (((p0 + (p0 + p0)) + "") + "foo bar"))
                var v6 uint = 18446744073709551615
                _ = v6
            }
            fmt.Println(p0)
        } else {
            fmt.Println(p0)
        }
    } else {
        var v7 bool = ((4294967295 != 0) && (true && ((!true) || true)))
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(fi8, p0)
            _ = fi8
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: 255, F1: 32767, F2: 255}
        p0 = p0
        var v11 uint = 757
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: false, F1: 1, F2: (true && true)}
    var v13 int32 = -2147483648
    var v14 string = ("foo bar" + p0)
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 32767
}
func genFn1(p0 uint16, p1 int64) int16 {
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := []int{fi15, fi15, 128, (^(fi15))}
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(p0, fi15)
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    var v18 uint32 = 760
    var v19 uint16 = p0
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return -1
}

func main() {
    var v1 float64 = 3.14159
    var v2 int16 = 786
    v3 := make(map[int]uint32)
    var v4 uint = 42
    v2 = v2
    var v20 bool = (false || false)
    var v21 uint16 = 797
    for fi22 := 0; fi22 < 14; fi22++ {
        switch "foo bar" {
        case "hello":
            v4 = (v4 / ((v4) | 1))
        default:
            v2 -= (v2 - 965)
        }
        _ = fi22
    }
    v2 -= (v2 * v2)
    switch 32768 {
    case -9223372036854775808:
        v2 = -1
    case 502:
        v23 := []int{9223372036854775807, 4294967295, 2147483647, -2147483648}
        _ = v23
    }
    v24 := GenS2{F0: 160, F1: v21, F2: 255}
    if (!(128 < 32767)) {
        var v25 string = (("foo bar" + "a") + ("" + "hello"))
        v26 := GenS1{F0: 42, F1: 2147483647, F2: (4294967295 == 65535)}
        _ = v25
        _ = v26
    } else {
        v20 = (0 > 128)
    }
    var v27 uint16 = 511
    v20 = false
    fmt.Println(v2, v27, v27)
    var v28 uint32 = 65535
    v29 := GenS1{F0: 9223372036854775807, F1: 1, F2: (("foo bar" + "hello") < ("" + ("\n\t" + "hello")))}
    for fi30 := 0; fi30 < 11; fi30++ {
        v20 = v20
        _ = fi30
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 int16 = v2
            var v34 int32 = -2147483648
            v28 = v28
            _ = fi32
            _ = v33
            _ = v34
        }
        v27 -= (v21 ^ (v27 << (((v4 % ((v4) | 1))) & 63)))
        _ = fi31
    }
    _ = len(v3)
    v1 = (v1 * v1)
    if (!(!true)) {
        v35 := GenS4{F0: 127, F1: v27}
        fmt.Println(v27)
        _ = v35
    } else {
        fmt.Println(v28, v21, v28)
        v36 := GenS4{F0: 13, F1: (v27 | 993)}
        if ((65535 >= (v27 ^ (^(v27)))) || v20) {
            v21 += v21
            _ = len(v3)
            switch ("\n\t" + "hello") {
            case "foo bar":
                var v37 int32 = 42
                _ = v37
            case "":
                _ = len(v3)
            case "\n\t":
                var v38 int8 = 105
                _ = v38
            }
        } else {
            var v39 uint64 = 0
            switch (("" + "hello") + "") {
            case "foo bar":
                v2 = v2
            case "\n\t":
                switch 589 {
                case -1:
                    var v40 float32 = -2.25
                    _ = v40
                }
            }
            _ = v39
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
