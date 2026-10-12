package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
func genFn0(p0 int, p1 uint8, p2 uint32) uint16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 float64 = 0.1
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint32 = p2
            v9 := []int{(fi5 * (fi7 * fi5)), 384, fi5}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v10 := []int{42, p0, p0}
    _ = len(v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}
func genFn1() int16 {
    v11 := GenS1{F0: 1}
    if true {
        v12 := GenS1{F0: 122}
        var v13 uint16 = genFn0(-941, 2, 65535)
        switch "a" {
        case "\n\t":
            if (507 < 1000) {
                v14 := GenS1{F0: 2}
                _ = v14
            } else {
                v15 := GenS2{F0: 42, F1: -2147483648, F2: 1}
                switch (("foo bar" + "foo bar") + ("foo bar" + (("" + ("hello" + "\n\t")) + "foo bar"))) {
                case "":
                    for fi16 := 0; fi16 < 8; fi16++ {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            v18 := GenS2{F0: 42, F1: -1, F2: 68}
                            _ = fi17
                            _ = v18
                        }
                        var v19 int = 32767
                        switch (v19 - -938) {
                        case 32767:
                            var v20 int64 = 42
                            _ = v20
                        case -2147483648:
                            var v21 uint8 = 2
                            _ = v21
                        }
                        _ = fi16
                        _ = v19
                    }
                case "a":
                    var v22 uint64 = 560
                    _ = v22
                }
                _ = v15
            }
        default:
            fmt.Println(v13)
        }
        _ = v12
        _ = v13
    }
    var v23 string = "a"
    v24 := []int{-906, -694}
    _ = v11
    _ = v23
    _ = v24
    return 0
}

func main() {
    var v1 float64 = 0.1
    var v2 bool = false
    v3 := make(map[string]bool)
    v4 := []uint32{118, 0}
    v1 += 1e10
    var v25 int16 = 20
    var v26 int64 = -9223372036854775808
    v1 -= v1
    for fi27 := 0; fi27 < 15; fi27++ {
        fmt.Println(v25)
        _ = fi27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 uint64 = 685
        var v30 int64 = v26
        v4 = append(v4, 356)
        _ = fi28
        _ = v29
        _ = v30
    }
    v1 = ((v1 / 1.0) * (3.14159 - float64(1e10)))
    fmt.Println(v1, v1, v1)
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 string = "a"
        v26 = -722
        if ((false || (!(((^(fi31)) >= fi31) || v2))) || true) {
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 string = (v32 + v32)
                _ = fi33
                _ = v34
            }
            v32 = v32
            v35 := GenS1{F0: 32767}
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
