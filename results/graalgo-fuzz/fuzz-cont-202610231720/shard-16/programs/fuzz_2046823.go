package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint16, p2 bool) uint32 {
    if (true || (!((!(!(-2.25 == (((1e10 * (((0.1 * -2.25) * (-2.25 + 0.1)) * ((3.14159 - 0.0) * ((1e10 - 3.14159) - -2.25)))) * 1e10) * 3.14159)))) && (!(-147 > -970))))) {
        var v7 uint16 = p1
        var v8 int32 = 0
        _ = v7
        _ = v8
    } else {
        switch 9223372036854775807 {
        case 32768:
            v9 := GenS1{F0: ((p2 && (p0 || false)) || (!(!((411 < 18446744073709551615) || (!((127 >= 191) || (-32768 < -602))))))), F1: 212}
            _ = v9
        case 834:
            for fi10 := 0; fi10 < 13; fi10++ {
                p0 = (false || true)
                p0 = true
                var v11 string = ""
                _ = fi10
                _ = v11
            }
        case 742:
            fmt.Println(p1, p2)
        }
        v12 := GenS1{F0: (2147483647 < 65535), F1: 2}
        if false {
            if p2 {
                var v13 int32 = 42
                v14 := []int{32768, 127, 4294967295, 4294967295}
                p2 = p2
                _ = v13
                _ = v14
            } else {
                switch "a" {
                case "hello":
                    v15 := []int{4294967295, 1, 9223372036854775807, 833}
                    _ = v15
                default:
                    v16 := []int{-2147483648}
                    _ = v16
                }
                if true {
                    p1 = 0
                    v17 := GenS1{F0: p0, F1: 59}
                    _ = v17
                }
                for fi18 := 0; fi18 < 3; fi18++ {
                    p2 = p0
                    _ = fi18
                }
            }
            p1 = (p1 << ((255) & 63))
        } else {
            fmt.Println(p2, p0)
            var v19 uint8 = 128
            _ = v19
        }
        _ = v12
    }
    v20 := GenS2{F0: 0}
    var v21 bool = false
    for fi22 := 0; fi22 < 8; fi22++ {
        switch "" {
        case "hello":
            for fi23 := 0; fi23 < 18; fi23++ {
                p2 = p0
                if (true && (!(0 == 127))) {
                    var v24 int16 = -1
                    v25 := []int{(^(fi22)), fi22, 2147483647}
                    _ = v24
                    _ = v25
                }
                v21 = false
                switch ("foo bar" + ((("hello" + ("" + "")) + "") + (((("hello" + "a") + "foo bar") + ("hello" + ("\n\t" + ((("hello" + ("\n\t" + "foo bar")) + "a") + "foo bar")))) + (("a" + "") + "\n\t")))) {
                case "\n\t":
                    var v26 float32 = ((0.1 + (-2.25 / (3.14159 / 0.1))) * 1.5)
                    _ = v26
                case "foo bar":
                    var v27 int64 = -9223372036854775808
                    _ = v27
                }
                fmt.Println(p0, p2)
                _ = fi23
            }
        case "":
            v28 := []int{}
            _ = v28
        }
        for fi29 := 0; fi29 < 13; fi29++ {
            fmt.Println(p2, v21, p0)
            _ = fi29
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 437
}

func main() {
    var v1 int64 = 158
    var v2 int = 1
    var v3 int64 = (v1 / ((v1) | 1))
    var v4 uint64 = 251
    var v5 uint64 = v4
    var v6 float64 = 3.14159
    fmt.Println(v5, v1)
    var v30 float64 = 1.5
    v5 -= (v5 / ((v4) | 1))
    v2 -= 2147483647
    v31 := []int{(v2 | 42)}
    v30 = v6
    if (157 == 0) {
        v4 = 18446744073709551615
        var v32 int32 = 0
        _ = v32
    } else {
        var v33 int8 = -128
        _ = v33
    }
    switch 32768 {
    case 127:
        v2 = 4294967295
    case 0:
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
    }
    v6 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
