package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint16) int {
    var v5 string = "\n\t"
    v6 := []int{2147483647, 751, 32768, 777}
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = (p0 - (^(p0)))
        var v8 uint16 = p0
        var v9 int64 = -502
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 bool, p2 string) uint16 {
    for fi10 := 0; fi10 < 5; fi10++ {
        p1 = true
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint64 = 2
            _ = fi11
            _ = v12
        }
        p1 = true
        _ = fi10
    }
    switch p2 {
    case "foo bar":
        v13 := []int{2147483647, 9223372036854775807, 0}
        _ = v13
    case "\n\t":
        var v14 uint16 = 445
        _ = v14
    }
    p2 = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := []string{(("a" + (("" + "foo bar") + "\n\t")) + ""), "", ("a" + ("hello" + ("foo bar" + ("" + (("a" + "foo bar") + ("" + "foo bar"))))))}
    var v15 float32 = 1.5
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v15 = (0.0 * 0.0)
    var v16 uint16 = 65535
    for fi17 := 0; fi17 < 17; fi17++ {
        if false {
            v18 := GenS1{F0: "foo bar"}
            _ = len(v4)
            _ = v18
        } else {
            if true {
                _ = len(v4)
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 int64 = 935
                    var v21 uint32 = 2147483648
                    v20 -= -9223372036854775808
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                var v22 float64 = -2.25
                _ = v22
            }
            v16 -= (v16 & v16)
        }
        if true {
            var v23 uint16 = (^(v16))
            var v24 uint64 = 4294967295
            _ = v23
            _ = v24
        }
        _ = fi17
    }
    v15 = v15
    v16 += genFn1(1, ((42 < 42) || false), "foo bar")
    v16 = (v16 | (^(v16)))
    var v25 string = "hello"
    if (4294967295 != 0) {
        switch 0 {
        case 42:
            v25 = ("hello" + "\n\t")
        default:
            var v26 float32 = 0.1
            _ = v26
        }
        var v27 int = 128
        _ = v27
    }
    v28 := v2[65535]
    v29 := GenS1{F0: ""}
    delete(v3, genFn0(genFn1(42, (!true), v25)))
    v28 = "\n\t"
    v30 := GenS1{F0: "foo bar"}
    var v31 int = 32767
    fmt.Println(v15, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
