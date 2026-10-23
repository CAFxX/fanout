package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint32, p1 int32) int8 {
    v7 := []int{482, -1}
    v8 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    v3 := []string{(((("a" + ((("" + "hello") + ("a" + "hello")) + "a")) + "\n\t") + (("\n\t" + "") + "hello")) + "a"), (("a" + ("\n\t" + "\n\t")) + "\n\t"), "foo bar"}
    v4 := []int{449, -17}
    v5 := []int{2147483647, -685, 326, 9223372036854775807}
    v6 := make(map[string]int)
    if false {
        if true {
            var v9 uint8 = 159
            v9 = v9
            if true {
                var v10 float64 = -2.25
                var v11 uint64 = 41
                _ = v10
                _ = v11
            } else {
                var v12 uint32 = 2147483648
                _ = v12
            }
            _ = v9
        }
    } else {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int32 = -1
            _ = fi13
            _ = v14
        }
        switch "\n\t" {
        case "\n\t":
            _ = len(v6)
        case "hello":
            var v15 int8 = -1
            _ = v15
        case "a":
            var v16 uint32 = 0
            _ = v16
        }
    }
    var v17 bool = ((!(832 < 2147483648)) && false)
    switch -1 {
    case 0:
        if (42 < -1) {
            _ = len(v4)
        } else {
            v17 = (!v17)
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 int = fi18
                delete(v2, "foo bar")
                _ = fi18
                _ = v19
            }
            v17 = (!(4294967295 >= 1))
        }
    case 127:
        var v20 int64 = 42
        _ = v20
    }
    delete(v1, 42)
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 int8 = -115
        _ = fi21
        _ = v22
    }
    fmt.Println(v17)
    switch "foo bar" {
    case "foo bar":
        v17 = v17
    case "\n\t":
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 uint8 = 128
            _ = fi23
            _ = v24
        }
    default:
        var v25 uint16 = 2
        _ = v25
    }
    v26 := GenS1{F0: true}
    switch 961 {
    case 431:
        v17 = ((v17 && v17) || v17)
    case 146:
        v17 = false
    }
    if (42 != -1) {
        v27 := GenS1{F0: (802 >= 42)}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
