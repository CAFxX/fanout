package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 string; F1 int64; F2 uint8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint8) uint8 {
    p0 = p0
    if true {
        switch ("foo bar" + "a") {
        case "a":
            v7 := GenS1{F0: 473}
            _ = v7
        case "foo bar":
            p1 = p1
        default:
            var v8 string = ("hello" + "hello")
            _ = v8
        }
        v9 := GenS1{F0: -1}
        _ = v9
    } else {
        if false {
            v10 := []int{173, 32767, -676, 128}
            _ = v10
        } else {
            var v11 float64 = ((-2.25 * (-2.25 / (1.5 + ((0.1 * 1e10) / (1e10 * 3.14159))))) - 0.0)
            _ = v11
        }
        var v12 int = -16
        fmt.Println(p0, v12)
        _ = v12
    }
    var v13 uint64 = 42
    _ = p0
    _ = p1
    _ = v13
    return (p1 | 157)
}

func main() {
    v1 := []uint32{2147483647, 65535}
    v2 := make(map[string]bool)
    var v3 int32 = 424
    var v4 float32 = -2.25
    var v5 float32 = ((v4 / (float32((1.5 * 1e10)) + ((((v4 / 0.1) + 0.1) * 1e10) - 1e10))) * v4)
    var v6 bool = true
    v14 := GenS2{F0: (1000 != 1), F1: 0}
    if false {
        switch 128 {
        case -407:
            v5 -= v5
        case -9223372036854775808:
            _ = len(v2)
        case 1:
            delete(v2, ("" + "foo bar"))
        }
        v3 -= v3
    } else {
        v1 = append(v1, 738)
    }
    var v15 uint32 = 2147483648
    v5 -= (((v4 * float32(0.0)) / (v5 * v4)) + v4)
    v16 := GenS1{F0: -1}
    if ((-11 < v3) && (!v6)) {
        var v17 uint16 = 32767
        switch 1 {
        case -708:
            v2[(("foo bar" + "\n\t") + "")] = v6
        case 0:
            var v18 uint32 = 690
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := v2["foo bar"]
            v6 = false
            _ = fi19
            _ = v20
        }
        _ = v17
    } else {
        var v21 int32 = (v3 + -403)
        v6 = false
        var v22 bool = ((v3 * v21) > 42)
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: ((v6 || (v6 || true)) && v6), F1: 42}
    fmt.Println(v3)
    var v24 int64 = 42
    v24 -= v24
    if v6 {
        switch int(v15) {
        case 276:
            var v25 float32 = v5
            _ = v25
        case 996:
            switch "foo bar" {
            case "hello":
                v26 := GenS3{F0: (3.14159 < ((3.14159 + float64(v4)) - 3.14159)), F1: v3}
                _ = v26
            case "\n\t":
                var v27 int = -1
                _ = v27
            }
        }
    } else {
        switch ("a" + ((("a" + "foo bar") + "") + "foo bar")) {
        case "a":
            var v28 int32 = (v3 + (-(v3)))
            _ = v28
        }
    }
    _ = len(v2)
    if v6 {
        if v6 {
            v29 := GenS1{F0: v3}
            fmt.Println(v5, v5)
            switch 32768 {
            case 0:
                var v30 int = -9223372036854775808
                _ = v30
            case -2147483648:
                v31 := GenS4{F0: ((("a" + "hello") + "hello") + "hello"), F1: 42, F2: 186}
                _ = v31
            case -73:
                var v32 float32 = v5
                _ = v32
            }
            _ = v29
        } else {
            if v6 {
                v6 = v6
                v4 = ((v5 / 3.14159) - float32(float64((0.1 - v5))))
            }
            v4 -= 1e10
        }
        var v33 string = "hello"
        var v34 uint16 = 255
        _ = v33
        _ = v34
    }
    delete(v2, "a")
    v35 := v2[(("hello" + "") + "")]
    v15 += 0
    for fi36 := 0; fi36 < 16; fi36++ {
        v37 := GenS5{F0: -48}
        _ = fi36
        _ = v37
    }
    var v38 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
