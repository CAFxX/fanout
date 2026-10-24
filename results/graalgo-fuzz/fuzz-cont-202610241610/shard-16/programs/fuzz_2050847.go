package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint) string {
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := []int{(^(fi4))}
            p0 = p0
            var v7 uint8 = 128
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    if ((!true) || true) {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 bool = ("foo bar" > ("" + "a"))
            v10 := []int{}
            fmt.Println(p0)
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(p0)
    } else {
        var v11 uint8 = 42
        _ = v11
    }
    p0 = 1
    v12 := []int{4294967295, 127, 54, -290}
    _ = p0
    _ = v12
    return ("hello" + "a")
}
func genFn1(p0 int32, p1 int32, p2 int8) int8 {
    v13 := GenS1{F0: "\n\t"}
    var v14 bool = (2 != 0)
    var v15 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 45
}

func main() {
    var v1 string = ""
    v2 := []uint32{2147483648, 532, 721, 2147483648}
    v3 := make(map[string]bool)
    switch 128 {
    case 2147483647:
        if true {
            v16 := GenS1{F0: v1}
            v1 = (v1 + ("\n\t" + "a"))
            if false {
                if false {
                    v17 := v3["foo bar"]
                    _ = v17
                }
                v18 := GenS2{F0: 255}
                var v19 uint64 = 444
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 9; fi20++ {
                    _ = len(v3)
                    _ = fi20
                }
                if true {
                    switch ((("\n\t" + v1) + "hello") + v1) {
                    case "a":
                        v21 := GenS2{F0: 2}
                        _ = v21
                    default:
                        fmt.Println(v1)
                    }
                }
            }
            _ = v16
        }
    default:
        for fi22 := 0; fi22 < 12; fi22++ {
            v1 = v1
            v3["foo bar"] = false
            _ = fi22
        }
    }
    var v23 float32 = float32(3.14159)
    var v24 float32 = v23
    v25 := GenS2{F0: 165}
    _ = len(v3)
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            v23 -= v23
            _ = fi27
        }
        _ = fi26
    }
    var v28 float64 = 1e10
    var v29 string = ("hello" + "a")
    v30 := v3[("hello" + v1)]
    if (1 == 42) {
        switch v29 {
        case "foo bar":
            var v31 int16 = 1
            _ = v31
        case "\n\t":
            for fi32 := 0; fi32 < 12; fi32++ {
                v1 = v29
                var v33 int8 = 42
                _ = fi32
                _ = v33
            }
        case "a":
            v1 = "hello"
        }
        if ((v28 - ((float64(v24) - v28) + (v28 / ((v28 * (0.0 - (v28 * float64(float32(3.14159))))) / float64(v23))))) >= 0.1) {
            for fi34 := 0; fi34 < 12; fi34++ {
                _ = len(v3)
                var v35 float32 = v24
                _ = fi34
                _ = v35
            }
        } else {
            v24 += 0.0
            fmt.Println(v29, v28, v23)
        }
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 uint32 = 357
        v37 += v37
        v29 = v1
        _ = fi36
        _ = v37
    }
    _ = len(v3)
    fmt.Println(v24)
    if (v30 || v30) {
        fmt.Println(v30)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
