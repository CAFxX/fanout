package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 string }
func genFn0() int32 {
    var v6 int = 32768
    var v7 int64 = 69
    v6 = v6
    if (599 >= 4294967295) {
        v8 := GenS1{F0: v6}
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (("a" + "\n\t") + "a") {
            case "hello":
                for fi10 := 0; fi10 < 17; fi10++ {
                    v11 := GenS1{F0: (v6 % (((^(v6))) | 1))}
                    var v12 uint64 = 657
                    v6 += fi10
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                v6 = (fi13 << ((2) & 63))
                v14 := GenS2{F0: -602, F1: 128, F2: -128}
                v6 -= 127
                v15 := GenS3{F0: "foo bar"}
                v16 := []int{128}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi9
        }
        _ = v8
    }
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint) uint {
    var v17 string = ("\n\t" + "hello")
    v18 := []int{32767, 565, 32768, -2147483648}
    p0 = 4294967295
    v19 := GenS4{F0: 579}
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 int32 = 0
    v5 := make(map[string]uint32)
    if (0.0 == (3.14159 + 1e10)) {
        v20 := GenS5{F0: 165}
        var v21 uint8 = 255
        switch "foo bar" {
        case "hello":
            fmt.Println(v4, v4)
        case "":
            var v22 uint64 = 208
            _ = v22
        }
        _ = v20
        _ = v21
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(v4, v4)
            v24 := GenS5{F0: 530}
            fmt.Println(v4, fi23)
            _ = fi23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS6{F0: fi25, F1: ""}
        v4 = (v4 & (v4 | (v4 % ((v4) | 1))))
        delete(v5, "\n\t")
        _ = fi25
        _ = v26
    }
    var v27 uint64 = 42
    if true {
        if true {
            var v28 float64 = 3.14159
            _ = v28
        } else {
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
            if (0.0 == 1e10) {
                v4 = genFn0()
                v4 -= genFn0()
            }
        }
    } else {
        v29 := GenS6{F0: -2147483648, F1: "a"}
        var v30 float64 = float64(1e10)
        _ = v29
        _ = v30
    }
    if true {
        v27 = (v27 >> ((genFn1(1)) & 63))
        v31 := GenS2{F0: -32768, F1: 42, F2: 0}
        v27 -= v27
        _ = v31
    } else {
        v27 -= v27
    }
    switch "a" {
    case "foo bar":
        v32 := GenS5{F0: 42}
        _ = v32
    case "a":
        if false {
            fmt.Println(v4, v4)
        } else {
            var v33 string = ("" + ("a" + ("\n\t" + "")))
            var v34 int8 = 1
            var v35 int16 = -1
            _ = v33
            _ = v34
            _ = v35
        }
    case "":
        if false {
            v4 = (v4 ^ v4)
        }
    default:
        var v36 int32 = 1
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        fmt.Println(fi37)
        switch "\n\t" {
        case "":
            v4 += -2147483648
        case "hello":
            v2 = append(v2, (fi37 << ((uint(v27)) & 63)))
        case "\n\t":
            var v38 int64 = -9223372036854775808
            _ = v38
        }
        if false {
            if true {
                var v39 float64 = 3.14159
                var v40 uint = 781
                _ = len(v1)
                _ = v39
                _ = v40
            }
        }
        _ = fi37
    }
    fmt.Println(v4, v27)
    for fi41 := 0; fi41 < 18; fi41++ {
        var v42 uint32 = 0
        v5[("" + "foo bar")] = v42
        _ = fi41
        _ = v42
    }
    var v43 int = 449
    v44 := GenS4{F0: 65535}
    for fi45 := 0; fi45 < 12; fi45++ {
        v46 := GenS5{F0: 2147483648}
        fmt.Println(v43)
        var v47 uint16 = 561
        _ = fi45
        _ = v46
        _ = v47
    }
    var v48 uint64 = (v27 % (((v27 << ((genFn1(2)) & 63))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v5
}
