package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32) uint64 {
    var v4 float32 = (0.0 * 1.5)
    v5 := GenS1{F0: (116 < 42), F1: -92}
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint32) int64 {
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            fmt.Println(fi6, p0)
            v7 := GenS1{F0: (-7 == 0), F1: -44}
            p0 = (^(p0))
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int64 = 42
            switch fi8 {
            case 42:
                switch "\n\t" {
                case "a":
                    v10 := []int{132, fi8, (fi8 ^ fi8)}
                    _ = v10
                default:
                    v9 = 9223372036854775807
                }
            }
            _ = fi8
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS2{F0: 42}
        fmt.Println(p0, fi11)
        _ = fi11
        _ = v12
    }
    p0 -= (p0 ^ (^(p0)))
    _ = p0
    return -95
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := []string{"a", "hello", "a"}
    var v13 int64 = genFn1(2147483647)
    var v14 int64 = v13
    fmt.Println(v13, v14)
    switch 675 {
    case -837:
        var v15 int32 = -704
        _ = v15
    case 42:
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    default:
        v14 = -1
    }
    v16 := GenS1{F0: ((genFn1(42) != v13) || (42 > 42)), F1: 1}
    fmt.Println(v13)
    for fi17 := 0; fi17 < 1; fi17++ {
        if (0 <= 32767) {
            var v18 float64 = 0.0
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            _ = v18
        }
        _ = len(v3)
        _ = fi17
    }
    switch "\n\t" {
    case "a":
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 float64 = -2.25
            _ = fi19
            _ = v20
        }
    case "foo bar":
        switch (((("a" + "hello") + "hello") + ("" + "\n\t")) + "a") {
        case "a":
            fmt.Println(v14)
        case "foo bar":
            fmt.Println(v14, v13)
        }
    case "":
        switch "a" {
        case "\n\t":
            var v21 uint32 = 42
            _ = v21
        case "a":
            for fi22 := 0; fi22 < 10; fi22++ {
                v14 = genFn1(945)
                v13 += v14
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 uint64 = 18446744073709551615
                    fmt.Println(v14, fi22, v24)
                    var v25 uint32 = 42
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
        }
    }
    fmt.Println(v13)
    if (!true) {
        v13 = (v13 << ((319) & 63))
    }
    var v26 bool = (4294967295 == 910)
    if v26 {
        if (1 != -1) {
            switch ("\n\t" + (("foo bar" + "") + ("hello" + "hello"))) {
            case "a":
                v26 = (("hello" + "hello") >= "\n\t")
            case "":
                var v27 int64 = (v13 ^ (v14 / ((v13) | 1)))
                _ = v27
            case "hello":
                var v28 int = -509
                _ = v28
            }
            var v29 int32 = 2147483647
            _ = v29
        } else {
            v13 = (^(v14))
            var v30 string = ("a" + "")
            switch ("\n\t" + "hello") {
            case "foo bar":
                if false {
                    fmt.Println(v14)
                }
            case "hello":
                if false {
                    fmt.Println(v14, v13)
                }
            default:
                var v31 uint = 18446744073709551615
                _ = v31
            }
            _ = v30
        }
        var v32 uint = 255
        var v33 float32 = 3.14159
        _ = v32
        _ = v33
    } else {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = len(v3)
    }
    v26 = ("\n\t" < "")
    var v34 float32 = 1e10
    _ = len(v1)
    var v35 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v35
}
