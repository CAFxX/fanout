package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int64) uint8 {
    var v7 bool = true
    var v8 uint = 18446744073709551615
    v9 := []int{631, -1, 9223372036854775807}
    var v10 string = "foo bar"
    switch "" {
    case "a":
        var v11 int = -1
        _ = v11
    case "":
        switch (((((("" + v10) + v10) + v10) + (v10 + (((v10 + v10) + ((v10 + v10) + ("\n\t" + ((v10 + (((("\n\t" + v10) + v10) + "a") + v10)) + "a")))) + ((((v10 + v10) + "hello") + (v10 + ("a" + "hello"))) + v10)))) + ("hello" + "hello")) + "a") {
        case "a":
            v7 = (0.0 < (float64((float32(0.1) * -2.25)) - 1.5))
        case "hello":
            v7 = v7
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                v10 = v10
                _ = fi12
            }
        }
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint64, p1 uint32, p2 uint8) int {
    v13 := GenS1{F0: uint16(p1), F1: (p1 >> ((uint(p2)) & 63))}
    p0 = (p0 + p0)
    v14 := []int{108, 32767, 42, -659}
    p2 -= genFn0(0)
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 232
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    v3 := make(map[string]bool)
    v4 := []string{("a" + ("foo bar" + "")), "hello", "a", ("foo bar" + ("a" + (((("a" + ((("\n\t" + "hello") + "") + "\n\t")) + "\n\t") + "") + "foo bar")))}
    var v5 uint8 = 1
    v6 := make(map[uint32]string)
    v15 := v3["hello"]
    v16 := GenS2{F0: 1}
    v17 := v3[(("" + ("" + ("foo bar" + ("foo bar" + "hello")))) + ("a" + (((("" + ((("" + "\n\t") + "") + (("hello" + "a") + "hello"))) + ("a" + "\n\t")) + ("hello" + "\n\t")) + "\n\t")))]
    fmt.Println(v17, v5, v5)
    var v18 uint16 = 0
    var v19 int64 = 9223372036854775807
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 uint64 = 689
        _ = fi20
        _ = v21
    }
    v5 += v5
    if true {
        v19 -= 0
        var v22 string = "hello"
        var v23 int64 = v19
        _ = v22
        _ = v23
    } else {
        var v24 uint8 = (^(v5))
        if len(v4) > 0 {
            v4[0] = (("foo bar" + ("" + "\n\t")) + "hello")
        }
        v25 := GenS2{F0: 264}
        _ = v24
        _ = v25
    }
    switch 32768 {
    case 9223372036854775807:
        var v26 string = "a"
        _ = v26
    case 1:
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 uint = 42
            v6[265] = "\n\t"
            v29 := GenS3{F0: 1, F1: (fi27 ^ -462)}
            _ = fi27
            _ = v28
            _ = v29
        }
    case 4294967295:
        v5 += 136
    default:
        var v30 string = (("hello" + "") + ("\n\t" + "foo bar"))
        _ = v30
    }
    var v31 float64 = (float64(0.1) - 1e10)
    if true {
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 uint8 = uint8(v18)
            switch (fi32 % (((fi32 & (fi32 << ((18) & 63)))) | 1)) {
            case 4294967295:
                v33 = v33
            }
            var v34 float32 = 1e10
            _ = fi32
            _ = v33
            _ = v34
        }
        v5 = 2
        var v35 int8 = 9
        _ = v35
    }
    v2 = (v31 != float64(float32((-2.25 - ((v31 - (float64((0.0 * (((float32(-2.25) * 3.14159) / 1.5) / -2.25))) * (1.5 + -2.25))) / float64(float32(((float64(1.5) / (-2.25 * v31)) * 0.1))))))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
