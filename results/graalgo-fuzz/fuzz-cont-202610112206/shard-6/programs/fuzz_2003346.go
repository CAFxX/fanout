package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 string; F2 bool }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
func genFn0(p0 bool, p1 int8) int16 {
    p1 = 42
    v5 := []int{-2147483648}
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint = 946
        var v9 float32 = (0.1 + -2.25)
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 -= (p1 % (((p1 | p1)) | 1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1() bool {
    var v10 uint = 65535
    v11 := GenS1{F0: 2, F1: 9223372036854775807, F2: 1}
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            v10 = v10
            v10 = 42
            _ = fi13
        }
        v10 = 709
        v14 := GenS2{F0: -456, F1: "a", F2: ("foo bar" >= ("" + "hello"))}
        _ = fi12
        _ = v14
    }
    _ = v10
    _ = v11
    return ((false && false) || (-842 > -125))
}

func main() {
    var v1 uint64 = 42
    v2 := []uint32{2, 4294967295, 2147483647}
    v3 := make(map[string]int)
    var v4 uint = 2
    for fi15 := 0; fi15 < 4; fi15++ {
        if ((9223372036854775807 != 836) || true) {
            v1 -= v1
            var v16 uint64 = 4294967295
            var v17 int32 = 1
            _ = v16
            _ = v17
        } else {
            var v18 uint32 = 2
            fmt.Println(v4)
            for fi19 := 0; fi19 < 16; fi19++ {
                fmt.Println(v18, fi15)
                _ = fi19
            }
            _ = v18
        }
        switch ("a" + "\n\t") {
        case "a":
            v20 := GenS2{F0: 1, F1: "a", F2: false}
            _ = v20
        case "\n\t":
            v2 = append(v2, 65535)
        }
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(fi15, fi15)
        } else {
            v21 := GenS3{F0: 10, F1: (fi15 + fi15), F2: true}
            v1 -= 458
            _ = v21
        }
        _ = fi15
    }
    v22 := GenS2{F0: 9223372036854775807, F1: "a", F2: true}
    if (!(255 > 65535)) {
        var v23 string = "hello"
        v23 = "\n\t"
        fmt.Println(v4, v1)
        _ = v23
    }
    if (127 != 181) {
        v1 -= v1
        if (true && ((false || false) && (v4 != (v4 - (v4 << ((v4) & 63)))))) {
            var v24 int64 = 9223372036854775807
            _ = v24
        }
        if (v4 == 4294967295) {
            var v25 uint64 = (v1 & (v1 << ((v4) & 63)))
            v25 -= v1
            switch 9223372036854775807 {
            case 32768:
                delete(v3, ("a" + "foo bar"))
            case 42:
                v26 := GenS1{F0: 0, F1: 1, F2: 1}
                _ = v26
            default:
                if false {
                    var v27 int32 = -621
                    var v28 uint32 = 1
                    _ = v27
                    _ = v28
                } else {
                    v4 -= (v4 / ((v4) | 1))
                }
            }
            _ = v25
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                v1 = 80
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi29
            }
        }
    } else {
        var v30 uint = v4
        var v31 uint16 = 78
        var v32 uint = v30
        _ = v30
        _ = v31
        _ = v32
    }
    v4 = 1
    if ("\n\t" >= ((("\n\t" + (("\n\t" + ("" + ("hello" + "foo bar"))) + (("\n\t" + (((("hello" + "foo bar") + "\n\t") + "") + "foo bar")) + ("\n\t" + "")))) + "hello") + (("" + "a") + ""))) {
        var v33 float64 = float64((-2.25 * 3.14159))
        var v34 float32 = 3.14159
        _ = v33
        _ = v34
    }
    v35 := GenS2{F0: -1, F1: "hello", F2: false}
    switch -404 {
    case 32767:
        fmt.Println(v4)
    case 549:
        var v36 uint32 = 0
        _ = v36
    case 2147483647:
        v4 = v4
    }
    var v37 int = 2147483647
    var v38 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
}
