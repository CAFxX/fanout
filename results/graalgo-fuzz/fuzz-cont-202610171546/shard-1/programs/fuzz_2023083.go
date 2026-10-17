package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int8; F1 int8; F2 string }
type GenS6 struct { F0 uint8; F1 int; F2 uint32 }
type GenS7 struct { F0 string; F1 bool }
func genFn0(p0 int8, p1 uint16) int8 {
    var v5 string = ("foo bar" + "a")
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint32 = 4294967295
        v8 := GenS1{F0: 252}
        switch (fi6 >> ((42) & 63)) {
        case 128:
            if (!(87 != 427)) {
                var v9 uint = 2
                if (-1 <= 42) {
                    v7 -= (v7 % ((566) | 1))
                    var v10 float64 = 3.14159
                    _ = v10
                }
                _ = v9
            }
        case 784:
            for fi11 := 0; fi11 < 8; fi11++ {
                switch 42 {
                case -2147483648:
                    fmt.Println(v7)
                case 2147483647:
                    v12 := GenS1{F0: 65535}
                    _ = v12
                case -1:
                    v13 := []int{}
                    _ = v13
                }
                v14 := []int{(fi6 | 2147483647), 520, 781, (fi11 - fi11)}
                _ = fi11
                _ = v14
            }
        case 9223372036854775807:
            v15 := GenS1{F0: 4294967295}
            _ = v15
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = p0
    switch -545 {
    case 9223372036854775807:
        fmt.Println(p0)
    case 2147483647:
        v16 := GenS2{F0: p1}
        _ = v16
    case 32768:
        var v17 int = 127
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := []string{"hello", ("hello" + ""), (("" + "a") + "foo bar"), "hello"}
    v2 := []int{2147483647}
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    var v18 float64 = 1e10
    if (0 < 1) {
        if false {
            if len(v1) > 0 {
                v1[0] = "a"
            }
            for fi19 := 0; fi19 < 19; fi19++ {
                switch (("a" + "hello") + ("" + "foo bar")) {
                case "\n\t":
                    fmt.Println(fi19)
                }
                v20 := GenS3{F0: 6, F1: int16(fi19), F2: 596}
                if true {
                    v21 := GenS4{F0: 129, F1: (("" + "hello") + "foo bar")}
                    v22 := GenS1{F0: 18446744073709551615}
                    _ = v21
                    _ = v22
                } else {
                    switch "a" {
                    case "a":
                        var v23 float32 = -2.25
                        _ = v23
                    }
                    var v24 string = "hello"
                    _ = v24
                }
                _ = fi19
                _ = v20
            }
            v18 = float64(0.0)
        }
        v25 := v3[-9223372036854775808]
        v2 = append(v2, 2147483647)
        _ = v25
    }
    v26 := GenS5{F0: -107, F1: -1, F2: ("hello" + "a")}
    v27 := GenS6{F0: 217, F1: 975, F2: 2147483648}
    var v28 int32 = 42
    var v29 uint = 170
    for fi30 := 0; fi30 < 10; fi30++ {
        v28 -= (v28 & v28)
        _ = fi30
    }
    if (42 == 100) {
        var v31 uint64 = 1
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 bool = true
            v33 = v33
            _ = fi32
            _ = v33
        }
        _ = v31
    } else {
        for fi34 := 0; fi34 < 10; fi34++ {
            v35 := GenS2{F0: 255}
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v4[fi36] = 431
        v37 := GenS6{F0: 1, F1: 465, F2: 1}
        _ = fi36
        _ = v37
    }
    v28 = (v28 ^ (v28 ^ v28))
    var v38 int16 = 924
    var v39 uint32 = 54
    v40 := GenS7{F0: ("a" + "\n\t"), F1: true}
    var v41 int32 = v28
    v28 -= int32(v29)
    v42 := GenS5{F0: 100, F1: -29, F2: "hello"}
    var v43 bool = (!(("" + "hello") > ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
