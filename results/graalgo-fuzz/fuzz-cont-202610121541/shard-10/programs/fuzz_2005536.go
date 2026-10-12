package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 bool }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0(p0 int16) uint32 {
    switch ("" + "hello") {
    case "":
        p0 = p0
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint16 = 0
        _ = fi6
        _ = v7
    }
    var v8 int32 = 0
    v8 += (v8 * v8)
    v9 := GenS1{F0: "a", F1: (!(1 != 0)), F2: true}
    _ = p0
    _ = v8
    _ = v9
    return 358
}
func genFn1(p0 int64, p1 int) int32 {
    v10 := GenS1{F0: "", F1: false, F2: false}
    v11 := GenS1{F0: "\n\t", F1: true, F2: false}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := []int{4294967295, 32767}
    v4 := []uint32{0, 2}
    v5 := []int{265, 32767, -1, 1}
    var v12 uint64 = 749
    _ = len(v4)
    for fi13 := 0; fi13 < 1; fi13++ {
        v12 += 1
        _ = fi13
    }
    v12 -= v12
    v14 := v2[481]
    var v15 uint = 4294967295
    switch "a" {
    case "a":
        fmt.Println(v12, v14)
    case "":
        for fi16 := 0; fi16 < 2; fi16++ {
            delete(v1, "\n\t")
            var v17 uint32 = 2147483647
            _ = len(v1)
            _ = fi16
            _ = v17
        }
    case "\n\t":
        if (v14 && ((v14 && (!false)) && true)) {
            v14 = (!v14)
            v18 := GenS2{F0: -9223372036854775808, F1: 127}
            fmt.Println(v14, v14, v15)
            _ = v18
        } else {
            var v19 int64 = 1
            var v20 bool = true
            var v21 uint8 = 42
            _ = v19
            _ = v20
            _ = v21
        }
    default:
        v12 += v12
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            v15 = 18446744073709551615
            fmt.Println(v14, v14, v14)
            _ = fi23
        }
        fmt.Println(v15, v14)
        v12 += uint64(v15)
        _ = fi22
    }
    var v24 uint32 = 42
    for fi25 := 0; fi25 < 18; fi25++ {
        switch "a" {
        case "foo bar":
            for fi26 := 0; fi26 < 3; fi26++ {
                fmt.Println(v24, fi25)
                v27 := GenS1{F0: (("\n\t" + "hello") + ("hello" + "")), F1: ((v14 && (!((42 > 42) && ((v24 >> ((0) & 63)) > (^(v24)))))) && v14), F2: v14}
                _ = fi26
                _ = v27
            }
        case "a":
            var v28 int32 = 1
            _ = v28
        case "":
            fmt.Println(v14)
        default:
            v24 -= uint32(v15)
        }
        _ = fi25
    }
    if v14 {
        var v29 uint64 = v12
        _ = v29
    } else {
        var v30 bool = v14
        switch -2147483648 {
        case 2147483647:
            v2[(v24 << ((904) & 63))] = v14
        }
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        fmt.Println(v14)
        fmt.Println(v14, v14)
        if false {
            var v32 uint32 = (v24 - v24)
            _ = v32
        } else {
            var v33 string = "foo bar"
            for fi34 := 0; fi34 < 3; fi34++ {
                v35 := GenS1{F0: v33, F1: (v14 || (0 != 0)), F2: ((float32(float64(1e10)) - 0.0) == 1.5)}
                _ = fi34
                _ = v35
            }
            _ = v33
        }
        _ = fi31
    }
    v12 = v12
    if (1 != -655) {
        v14 = false
        var v36 uint32 = (^(v24))
        v37 := v2[(v36 - (v24 % (((v24 >> ((v15) & 63))) | 1)))]
        _ = v36
        _ = v37
    } else {
        if (0.0 != 0.1) {
            if false {
                delete(v1, ("a" + "\n\t"))
                var v38 int32 = -2147483648
                _ = v38
            }
            var v39 uint32 = v24
            _ = v39
        }
        var v40 int16 = -158
        v2[(v24 - genFn0(1000))] = v14
        _ = v40
    }
    if v14 {
        var v41 bool = v14
        v42 := GenS1{F0: ("" + "hello"), F1: true, F2: v41}
        if (0 >= -136) {
            var v43 int8 = 42
            fmt.Println(v14, v14)
            for fi44 := 0; fi44 < 7; fi44++ {
                var v45 uint8 = 42
                _ = fi44
                _ = v45
            }
            _ = v43
        } else {
            for fi46 := 0; fi46 < 18; fi46++ {
                v41 = v14
                _ = fi46
            }
        }
        _ = v41
        _ = v42
    } else {
        var v47 int = -2147483648
        fmt.Println(v14, v12)
        _ = v47
    }
    fmt.Println(v15, v15, v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
