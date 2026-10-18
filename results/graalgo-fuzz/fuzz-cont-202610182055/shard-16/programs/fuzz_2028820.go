package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16, p1 int8, p2 string) uint16 {
    p0 = p0
    p1 = 42
    for fi6 := 0; fi6 < 6; fi6++ {
        switch (("" + "\n\t") + ((p2 + "\n\t") + p2)) {
        case "foo bar":
            v7 := GenS1{F0: ((p0 - p0) > (p0 - 1)), F1: (p1 & (^(p1)))}
            _ = v7
        case "a":
            v8 := GenS2{F0: p1}
            _ = v8
        case "hello":
            v9 := []int{(fi6 & fi6), 2147483647, fi6, 212}
            _ = v9
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint64 = 139
        switch fi10 {
        case -9223372036854775808:
            v12 := []int{(fi10 >> ((2) & 63)), (fi10 - int(v11)), (fi10 & (fi10 >> ((18446744073709551615) & 63)))}
            _ = v12
        case 42:
            v13 := GenS2{F0: p1}
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1(p0 uint32) string {
    var v14 string = ("foo bar" + "a")
    fmt.Println(p0, p0)
    _ = p0
    _ = v14
    return ((v14 + "hello") + "\n\t")
}

func main() {
    var v1 uint = 138
    v2 := make(map[string]uint32)
    var v3 uint16 = 65535
    var v4 uint8 = 121
    v5 := make(map[string]bool)
    if true {
        v15 := []int{657, -800, 127, 213}
        for fi16 := 0; fi16 < 9; fi16++ {
            switch fi16 {
            case 32768:
                v1 = (v1 & (v1 * 65535))
            }
            _ = fi16
        }
        _ = v15
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        fmt.Println(v4, v4, v4)
        switch 508 {
        case -187:
            v3 = (v3 >> ((v1) & 63))
        case 2147483647:
            v18 := GenS3{F0: (v4 ^ (^(v4)))}
            _ = v18
        default:
            fmt.Println(v1, v4)
        }
        _ = fi17
    }
    var v19 int = 32768
    var v20 string = "hello"
    v19 = (v19 << ((219) & 63))
    if (127 <= -47) {
        var v21 uint32 = 2
        for fi22 := 0; fi22 < 20; fi22++ {
            v3 += v3
            _ = fi22
        }
        _ = v21
    } else {
        v3 += v3
        var v23 uint8 = (v4 - v4)
        _ = v23
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 string = (((("\n\t" + "") + ("" + v20)) + "") + (v20 + v20))
        if (((341 < -622) || true) || ((-(fi24)) == (v19 + (-(fi24))))) {
            v26 := GenS1{F0: false, F1: 0}
            v3 = 762
            _ = len(v2)
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    v27 := v2[("hello" + "a")]
    v28 := []int{(v19 + (^(v19))), (v19 / ((-975) | 1)), -258, v19}
    fmt.Println(v27, v19, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
