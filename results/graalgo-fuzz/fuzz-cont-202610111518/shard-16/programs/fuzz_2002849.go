package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int64 }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 int8) string {
    for fi6 := 0; fi6 < 17; fi6++ {
        switch -9223372036854775808 {
        case -2147483648:
            var v7 string = ("" + (("hello" + "foo bar") + "foo bar"))
            _ = v7
        case 128:
            if (!true) {
                fmt.Println(fi6)
            }
        case 32767:
            var v8 float64 = 3.14159
            _ = v8
        default:
            if false {
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 int32 = 1
                    switch (fi6 * (fi6 | (fi6 >> ((255) & 63)))) {
                    case -235:
                        p0 += p0
                    }
                    v11 := GenS1{F0: 290, F1: "a", F2: 8}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        }
        if false {
            fmt.Println(fi6, p0)
        } else {
            fmt.Println(p0)
            if (!(!(!(((-22 >= p0) || false) || (61 > 367))))) {
                if true {
                    p0 = p0
                }
            } else {
                var v12 string = "foo bar"
                _ = v12
            }
        }
        _ = fi6
    }
    if (true || false) {
        var v13 string = ("foo bar" + "hello")
        for fi14 := 0; fi14 < 8; fi14++ {
            p0 += (p0 * 127)
            v15 := GenS1{F0: (fi14 >> ((255) & 63)), F1: (v13 + ((v13 + v13) + v13)), F2: -174}
            p0 = (p0 % ((42) | 1))
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v16 uint64 = 0
        v17 := []int{}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(p0)
        var v19 int = fi18
        _ = fi18
        _ = v19
    }
    fmt.Println(p0)
    _ = p0
    return ("a" + "hello")
}
func genFn1(p0 uint16, p1 int) uint8 {
    p0 = (p0 / ((2) | 1))
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS1{F0: p1, F1: "", F2: 0}
        _ = fi20
        _ = v21
    }
    var v22 uint64 = 1
    _ = p0
    _ = p1
    _ = v22
    return 3
}

func main() {
    var v1 uint32 = 425
    v2 := []string{"a", "", "foo bar", ("" + "\n\t")}
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint = 312
    fmt.Println(v1, v1)
    var v23 int = 42
    var v24 uint = 2
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 bool = true
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := GenS1{F0: (-(fi27)), F1: "\n\t", F2: 0}
        if (true && (42 >= 207)) {
            v29 := GenS2{F0: 0, F1: "a"}
            switch -867 {
            case 42:
                var v30 int8 = 28
                _ = v30
            case -1:
                v1 = (v1 | 1)
            case -136:
                v1 = v1
            default:
                var v31 int = v23
                _ = v31
            }
            _ = v29
        } else {
            switch "hello" {
            case "":
                v5 = uint(v1)
            default:
                var v32 uint32 = v1
                _ = v32
            }
            if false {
                v2 = append(v2, "hello")
                fmt.Println(v5, fi27, v24)
            }
        }
        _ = fi27
        _ = v28
    }
    v5 = 236
    fmt.Println(v1)
    var v33 uint32 = (^(v1))
    fmt.Println(v1, v5)
    fmt.Println(v24, v23)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
