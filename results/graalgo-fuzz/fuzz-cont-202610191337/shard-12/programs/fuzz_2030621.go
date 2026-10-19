package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int8 }
func genFn0(p0 int32) uint64 {
    var v5 float64 = 1.5
    v5 = 1.5
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 -= p0
        var v7 uint16 = 255
        v5 = float64(1e10)
        _ = fi6
        _ = v7
    }
    fmt.Println(v5, p0)
    _ = p0
    _ = v5
    return 2
}

func main() {
    var v1 int = 25
    var v2 float64 = 0.1
    v3 := []string{"a", "foo bar", "a", "\n\t"}
    var v4 string = (("hello" + "\n\t") + "foo bar")
    var v8 bool = (true || false)
    var v9 uint16 = 774
    for fi10 := 0; fi10 < 16; fi10++ {
        if false {
            var v11 int16 = 816
            for fi12 := 0; fi12 < 4; fi12++ {
                if false {
                    if v8 {
                        v3 = append(v3, "a")
                        v13 := GenS1{F0: 26, F1: 42, F2: 597}
                        fmt.Println(fi12)
                        _ = v13
                    }
                } else {
                    fmt.Println(fi12)
                    v14 := GenS2{F0: -7, F1: 1, F2: -14}
                    _ = v14
                }
                v3 = append(v3, v4)
                _ = fi12
            }
            _ = v11
        }
        var v15 uint8 = 2
        switch v1 {
        case -1:
            if len(v3) > 0 {
                v3[0] = "hello"
            }
        case 32768:
            fmt.Println(v15, v8)
        case -2147483648:
            var v16 float64 = v2
            _ = v16
        }
        _ = fi10
        _ = v15
    }
    if v8 {
        switch "\n\t" {
        case "":
            switch (v1 ^ -683) {
            case 128:
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS1{F0: 0, F1: 2, F2: 4294967295}
                    v19 := GenS3{F0: v8}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case 32768:
                var v20 uint8 = 255
                _ = v20
            }
        case "foo bar":
            var v21 uint16 = (v9 - (v9 - (v9 & v9)))
            _ = v21
        default:
            if v8 {
                v4 = (v4 + "hello")
            }
        }
    } else {
        for fi22 := 0; fi22 < 15; fi22++ {
            v8 = v8
            v3 = append(v3, ((("a" + ("" + (v4 + v4))) + v4) + (v4 + v4)))
            v23 := GenS1{F0: -1, F1: 191, F2: 2}
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        if ((0.1 > -2.25) || v8) {
            var v25 bool = v8
            _ = v25
        }
        _ = fi24
    }
    v26 := GenS4{F0: (v1 >> ((1) & 63))}
    fmt.Println(v4, v1, v2)
    v2 -= float64(-2.25)
    _ = len(v3)
    fmt.Println(v8)
    _ = len(v3)
    for fi27 := 0; fi27 < 7; fi27++ {
        switch ((v4 + ("" + "hello")) + v4) {
        case "a":
            _ = len(v3)
        case "hello":
            fmt.Println(v2)
        case "":
            var v28 uint16 = 201
            _ = v28
        default:
            v3 = append(v3, "hello")
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi27
    }
    v29 := GenS5{F0: -32768, F1: 1}
    var v30 int8 = -128
    for fi31 := 0; fi31 < 20; fi31++ {
        v2 += v2
        _ = fi31
    }
    v32 := GenS3{F0: (1e10 < (1.5 * 0.1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}
