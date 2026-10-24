package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 uint32) string {
    v7 := GenS1{F0: -362, F1: false, F2: 2}
    fmt.Println(p1, p0)
    v8 := []int{}
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (v9 + v9)
}
func genFn1(p0 int) string {
    for fi10 := 0; fi10 < 20; fi10++ {
        if (1 >= 284) {
            p0 = -41
            p0 = fi10
        } else {
            var v11 uint8 = 158
            _ = v11
        }
        _ = fi10
    }
    switch "" {
    case "\n\t":
        var v12 string = "foo bar"
        _ = v12
    }
    v13 := GenS1{F0: (p0 / ((p0) | 1)), F1: true, F2: 735}
    v14 := []int{(p0 | p0), 1}
    _ = p0
    _ = v13
    _ = v14
    return ("hello" + ("a" + ("\n\t" + ("a" + ""))))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{150, 9223372036854775807, -291}
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    v6 := make(map[int]bool)
    v15 := GenS2{F0: ("a" + "\n\t")}
    v6[2147483647] = (1 < 255)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int16 = -32768
        fmt.Println(v17)
        var v18 int32 = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int64 = -788
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int32 = -1
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 uint = 912
            var v24 uint8 = 250
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v19)
    v19 = v19
    v25 := GenS2{F0: (("a" + ("foo bar" + "\n\t")) + ("\n\t" + (((("" + ("hello" + "hello")) + ("a" + "foo bar")) + "foo bar") + ("hello" + ("a" + "foo bar")))))}
    for fi26 := 0; fi26 < 20; fi26++ {
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 uint16 = 255
            var v29 float64 = -2.25
            v30 := v4[("" + "")]
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        switch fi26 {
        case 799:
            fmt.Println(v19, fi26)
        case -2147483648:
            var v31 string = "a"
            _ = v31
        case 332:
            var v32 int = 127
            _ = v32
        }
        fmt.Println(fi26)
        _ = fi26
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        if (1 != 0) {
            var v34 uint64 = 635
            var v35 int16 = 687
            v19 += 9223372036854775807
            _ = v34
            _ = v35
        }
        v36 := v3["hello"]
        v37 := GenS1{F0: (fi33 >> ((404) & 63)), F1: false, F2: 0}
        _ = fi33
        _ = v36
        _ = v37
    }
    v19 = (v19 + v19)
    fmt.Println(v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
