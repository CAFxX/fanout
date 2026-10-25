package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32, p1 int64) uint8 {
    fmt.Println(p0)
    var v5 int = -297
    v6 := []int{(v5 % ((v5) | 1)), v5, (v5 + (v5 ^ -9223372036854775808)), (^(v5))}
    switch 9223372036854775807 {
    case 1:
        p1 += -920
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int32, p1 bool, p2 int16) uint {
    var v7 uint = 255
    var v8 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (v7 % (((v7 | v7)) | 1))
}

func main() {
    v1 := []string{"a"}
    var v2 float64 = 1.5
    var v3 int16 = -271
    var v4 uint64 = 42
    if ((4294967295 <= v4) && (-44 > 9223372036854775807)) {
        v9 := GenS1{F0: 936}
        fmt.Println(v3)
        var v10 float32 = 1e10
        _ = v10
        _ = v9
    }
    fmt.Println(v2, v3, v3)
    var v11 int8 = 100
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, ((("hello" + ("a" + ("foo bar" + "hello"))) + ("" + "foo bar")) + ("foo bar" + "\n\t")))
    fmt.Println(v3)
    if len(v1) > 0 {
        v1[0] = (("" + ((("\n\t" + (((("foo bar" + "") + "a") + "a") + "")) + ("\n\t" + "hello")) + (("a" + "hello") + "a"))) + "foo bar")
    }
    var v12 uint8 = 33
    v2 = ((v2 / float64(1.5)) / float64((3.14159 + float32((v2 + float64(-2.25))))))
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            switch "\n\t" {
            case "":
                v2 += 0.0
            case "\n\t":
                _ = len(v1)
            case "hello":
                v12 -= v12
            }
            v3 -= -1
            var v15 uint16 = 0
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if true {
        switch 32767 {
        case 128:
            fmt.Println(v12)
        default:
            var v16 float32 = -2.25
            _ = v16
        }
    } else {
        var v17 uint16 = uint16(v12)
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(v11)
        v19 := GenS1{F0: genFn1(-58, (65535 != 1), (v3 - 318))}
        _ = fi18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
