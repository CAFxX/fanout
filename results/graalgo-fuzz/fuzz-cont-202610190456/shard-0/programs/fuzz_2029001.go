package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint32; F1 bool }
func genFn0() uint {
    v7 := GenS1{F0: 42}
    var v8 uint64 = 2
    var v9 string = "foo bar"
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int16, p1 uint16, p2 string) bool {
    fmt.Println(p1, p1)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return (p1 != (p1 - (p1 + 42)))
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]string)
    var v3 float64 = (-2.25 - v1)
    v4 := []string{("a" + "foo bar"), "hello", "a", ("hello" + ("foo bar" + "hello"))}
    v5 := make(map[int]bool)
    var v6 float32 = 3.14159
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 float32 = ((3.14159 + 0.1) / v6)
            var v13 uint8 = 128
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    if ((-2147483648 <= 42) && (false && true)) {
        var v14 int = 0
        _ = v14
    } else {
        v6 = (v6 + v6)
    }
    v4 = append(v4, ("\n\t" + "foo bar"))
    v4 = append(v4, ("" + "foo bar"))
    v1 = v1
    var v15 uint64 = 42
    v6 = v6
    v2[2] = ("\n\t" + (("\n\t" + "foo bar") + (("foo bar" + "") + "a")))
    var v16 float32 = -2.25
    switch 42 {
    case 127:
        var v17 uint64 = (v15 | v15)
        _ = v17
    default:
        v2[42] = ("a" + (("" + "\n\t") + "\n\t"))
    }
    v5[781] = (((v16 / (v16 / ((((v6 / float32(v3)) / v6) + (v6 - (((v6 * v6) / v6) + 0.0))) * v6))) == (v16 - (1.5 + v6))) && (true && (!(!(!true)))))
    v18 := GenS2{F0: 32767, F1: -936}
    for fi19 := 0; fi19 < 15; fi19++ {
        if true {
            var v20 int16 = -981
            v21 := GenS3{F0: 359, F1: true}
            for fi22 := 0; fi22 < 5; fi22++ {
                v15 = v15
                var v23 float32 = -2.25
                v6 += ((v6 - (3.14159 * v23)) + 3.14159)
                _ = fi22
                _ = v23
            }
            _ = v20
            _ = v21
        }
        var v24 int32 = -1
        if (!(true && (0 < 0))) {
            v24 = v24
            v4 = append(v4, ((("hello" + ("a" + ("\n\t" + ("foo bar" + "foo bar")))) + "hello") + ""))
        }
        _ = fi19
        _ = v24
    }
    var v25 uint64 = v15
    fmt.Println(v1, v15)
    v6 -= 1.5
    for fi26 := 0; fi26 < 19; fi26++ {
        _ = len(v4)
        v25 = 42
        _ = fi26
    }
    v27 := GenS1{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
