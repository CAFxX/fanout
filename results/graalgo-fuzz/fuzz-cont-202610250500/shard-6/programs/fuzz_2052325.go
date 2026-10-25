package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
func genFn0(p0 int32, p1 uint8) bool {
    fmt.Println(p1, p1)
    p0 = -976
    switch -1 {
    case 88:
        if (!(-244 != 9223372036854775807)) {
            v6 := []int{}
            _ = v6
        }
    case 32767:
        if false {
            switch ((("hello" + (("a" + ("" + (("hello" + "a") + ((("foo bar" + "hello") + "\n\t") + "foo bar")))) + (("a" + ("\n\t" + "hello")) + (("" + "hello") + ("" + ""))))) + "\n\t") + ("\n\t" + "")) {
            case "\n\t":
                fmt.Println(p0, p0)
            case "a":
                v7 := []int{-9223372036854775808, -2147483648, 622, -1}
                _ = v7
            }
        } else {
            var v8 int16 = 42
            var v9 float32 = -2.25
            v8 += 1
            _ = v8
            _ = v9
        }
    }
    p1 += p1
    var v10 int16 = -675
    _ = p0
    _ = p1
    _ = v10
    return true
}
func genFn1(p0 uint8, p1 bool, p2 int16) int64 {
    p2 = p2
    var v11 uint8 = 2
    switch "" {
    case "a":
        v12 := []int{32768, 1, 32768}
        _ = v12
    case "hello":
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    var v14 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v14
    return 807
}

func main() {
    v1 := []int{-713}
    var v2 int32 = -2147483648
    v3 := []uint32{}
    v4 := []int{-9223372036854775808, -85}
    var v5 int64 = -619
    var v15 uint = 0
    switch "foo bar" {
    case "a":
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(fi16, v5)
            var v17 int8 = -128
            var v18 int = (fi16 * fi16)
            _ = fi16
            _ = v17
            _ = v18
        }
    case "hello":
        var v19 int16 = 0
        _ = v19
    default:
        var v20 int = int(v15)
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        switch ("\n\t" + (("foo bar" + "") + "\n\t")) {
        case "":
            switch (fi21 ^ (fi21 & fi21)) {
            case 4294967295:
                v22 := GenS1{F0: ("\n\t" + ""), F1: 65535}
                _ = v22
            case 9223372036854775807:
                _ = len(v1)
            }
        case "hello":
            fmt.Println(v2, v15, v5)
        case "foo bar":
            switch int(v5) {
            case 42:
                v23 := GenS1{F0: "", F1: 65535}
                _ = v23
            case 128:
                var v24 int64 = v5
                _ = v24
            }
        default:
            v5 = v5
        }
        var v25 uint8 = 49
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi21
        _ = v25
    }
    _ = len(v1)
    v2 = -1
    v26 := GenS1{F0: "", F1: 501}
    if true {
        var v27 float32 = float32(float64(1.5))
        var v28 uint16 = 255
        if true {
            var v29 bool = (128 != 1)
            _ = v29
        }
        _ = v27
        _ = v28
    } else {
        v30 := GenS2{F0: -1, F1: 1, F2: -1}
        _ = len(v1)
        _ = v30
    }
    var v31 uint8 = 0
    var v32 int32 = v2
    if (649 >= 2) {
        var v33 uint16 = 32767
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 int16 = 1000
            v32 = -2147483648
            v36 := GenS2{F0: (v2 >> (((v15 | (v15 << (((v15 << ((v15) & 63))) & 63)))) & 63)), F1: (v33 ^ v33), F2: 0}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    }
    if true {
        v37 := GenS2{F0: v32, F1: 560, F2: -1}
        v2 -= (v2 << ((v15) & 63))
        _ = v37
    } else {
        v31 = v31
        switch 32768 {
        case 32768:
            for fi38 := 0; fi38 < 20; fi38++ {
                for fi39 := 0; fi39 < 4; fi39++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch (fi38 & fi38) {
                    case 776:
                        fmt.Println(v31, fi38)
                    default:
                        fmt.Println(v32)
                    }
                    fmt.Println(v15, v5, v5)
                    _ = fi39
                }
                v40 := GenS1{F0: "foo bar", F1: 775}
                _ = fi38
                _ = v40
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
