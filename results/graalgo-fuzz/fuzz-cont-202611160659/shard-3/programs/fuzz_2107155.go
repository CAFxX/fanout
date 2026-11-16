package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int8; F2 int32 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint8 }
func genFn0(p0 string, p1 uint) int16 {
    v4 := GenS1{F0: 0}
    if true {
        for fi5 := 0; fi5 < 17; fi5++ {
            switch fi5 {
            case 1:
                p0 = ((p0 + p0) + (p0 + "a"))
            }
            var v6 uint8 = 0
            _ = fi5
            _ = v6
        }
        p1 -= p1
        var v7 uint64 = 300
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1(p0 string, p1 int32) int64 {
    if true {
        p1 -= p1
        var v8 int16 = -1
        _ = v8
    }
    switch p0 {
    case "hello":
        v9 := []int{1, 502, 32767}
        _ = v9
    default:
        var v10 int64 = 0
        _ = v10
    }
    var v11 uint32 = 0
    _ = p0
    _ = p1
    _ = v11
    return 250
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    v12 := GenS2{F0: 820, F1: 4294967295, F2: 1}
    var v13 string = (v1 + (v1 + (("a" + ("" + v1)) + v1)))
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int32 = 886
        var v16 float32 = (-2.25 + -2.25)
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 1000}
    if (!true) {
        v13 = ("" + "hello")
        var v18 int32 = -1
        var v19 uint = 65535
        _ = v18
        _ = v19
    } else {
        v20 := GenS1{F0: genFn0((v1 + ("\n\t" + "\n\t")), 2)}
        var v21 uint = 970
        v22 := GenS3{F0: "hello"}
        _ = v20
        _ = v21
        _ = v22
    }
    switch 638 {
    case -9223372036854775808:
        v23 := GenS4{F0: 127, F1: 100, F2: 0}
        _ = v23
    }
    if (false || ((42 >= 808) && true)) {
        v24 := []int{4294967295, -364, 127}
        var v25 uint64 = 118
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := GenS5{F0: 76, F1: 2}
            fmt.Println(fi26, v13, v1)
            _ = fi26
            _ = v27
        }
    }
    v1 = v1
    var v28 int16 = 0
    var v29 uint16 = 827
    var v30 uint8 = 127
    var v31 uint8 = v30
    if false {
        switch v13 {
        case "":
            v32 := GenS6{F0: v29}
            _ = v32
        case "\n\t":
            v33 := []int{2147483647, 547}
            _ = v33
        case "hello":
            fmt.Println(v30, v30, v30)
        }
        switch v1 {
        case "foo bar":
            var v34 bool = false
            _ = v34
        case "":
            var v35 int64 = -1
            _ = v35
        case "a":
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 int = 4294967295
                fmt.Println(v29, v29, v1)
                v37 += 32767
                _ = fi36
                _ = v37
            }
        default:
            var v38 float64 = 0.0
            _ = v38
        }
        v39 := GenS7{F0: 1, F1: (^(v31))}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
