package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int) uint8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(fi5, fi5)
        var v6 int32 = -1
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: 42, F2: -1}
    v8 := GenS2{F0: -1}
    v9 := []int{p0, p0}
    for fi10 := 0; fi10 < 11; fi10++ {
        v9 = append(v9, p0)
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint64 = 18446744073709551615
            if false {
                p0 = -715
                var v13 float32 = 0.1
                var v14 uint32 = 0
                _ = v13
                _ = v14
            }
            var v15 string = "a"
            _ = fi11
            _ = v12
            _ = v15
        }
        var v16 int = p0
        _ = fi10
        _ = v16
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint32) uint8 {
    var v17 int = -9223372036854775808
    p0 += (p0 - p0)
    p0 -= (p0 & p0)
    v18 := GenS1{F0: 947, F1: 819, F2: 0}
    _ = p0
    _ = v17
    _ = v18
    return 255
}

func main() {
    var v1 int8 = -6
    var v2 int8 = (v1 + v1)
    v3 := make(map[uint32]int)
    var v4 int64 = -9223372036854775808
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            if true {
                switch ("hello" + "") {
                case "foo bar":
                    var v21 float64 = (1e10 + -2.25)
                    var v22 int = fi19
                    fmt.Println(v1, v1)
                    _ = v21
                    _ = v22
                }
            }
            _ = fi20
        }
        _ = fi19
    }
    if true {
        v23 := GenS2{F0: 1000}
        if ((42 < 2) || true) {
            switch 32767 {
            case 1:
                var v24 bool = (2 <= 1000)
                _ = v24
            case 128:
                var v25 uint64 = 42
                _ = v25
            case 32767:
                _ = len(v3)
            }
            var v26 string = ("\n\t" + "a")
            v27 := v3[482]
            _ = v26
            _ = v27
        }
        v1 = (v1 * 67)
        _ = v23
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        delete(v3, 2147483647)
        if (506 > 734) {
            if ((v2 + v1) == v2) {
                var v29 int32 = 11
                if true {
                    var v30 float64 = -2.25
                    _ = v30
                }
                _ = v29
            }
        }
        _ = fi28
    }
    var v31 float32 = (1.5 - ((3.14159 * 1e10) / ((1e10 + ((0.1 + 0.0) - (1.5 + (3.14159 - 0.0)))) + -2.25)))
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 5; fi33++ {
            switch ("hello" + "a") {
            case "":
                v34 := GenS1{F0: 2, F1: 563, F2: (v4 + (v4 - 42))}
                _ = v34
            default:
                if false {
                    var v35 int8 = 100
                    v36 := []int{-59, fi32, fi32, fi32}
                    switch "hello" {
                    case "hello":
                        v1 -= (-(v35))
                    case "\n\t":
                        var v37 uint16 = 0
                        _ = v37
                    }
                    _ = v35
                    _ = v36
                }
            }
            v1 = v1
            var v38 uint8 = 2
            _ = fi33
            _ = v38
        }
        var v39 uint = 255
        v40 := GenS2{F0: 32767}
        _ = fi32
        _ = v39
        _ = v40
    }
    v41 := []int{406, 42, -942}
    var v42 uint8 = 127
    var v43 int32 = -127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
