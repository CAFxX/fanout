package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 uint64) uint16 {
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            if (1 != -1) {
                switch fi6 {
                case 128:
                    var v7 string = "a"
                    _ = v7
                case 77:
                    p0 = ("a" + ((((p0 + (p0 + p0)) + (p0 + "hello")) + p0) + (p0 + ("foo bar" + ("\n\t" + ((("" + "hello") + p0) + "a"))))))
                case 42:
                    switch (fi6 % (((fi6 / (((fi6 * (^(fi6)))) | 1))) | 1)) {
                    case -82:
                        p1 -= 924
                    }
                default:
                    var v8 uint64 = (^(p1))
                    _ = v8
                }
                p0 = (p0 + "hello")
                if false {
                    var v9 uint = 0
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 uint8 = 146
                        _ = fi10
                        _ = v11
                    }
                    var v12 uint = v9
                    _ = v12
                    _ = v9
                } else {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        p0 = "\n\t"
                        v14 := GenS1{F0: ((905 == 32767) || ((p1 ^ p1) >= p1)), F1: 1000, F2: 2147483648}
                        _ = fi13
                        _ = v14
                    }
                }
            }
            v15 := GenS2{F0: 2}
            var v16 int64 = -497
            _ = v15
            _ = v16
        } else {
            switch 2147483647 {
            case -413:
                var v17 int64 = -194
                _ = v17
            default:
                var v18 uint = 18446744073709551615
                _ = v18
            }
        }
        if (!(2147483647 >= -759)) {
            v19 := []int{fi6, (fi6 + fi6)}
            v5 -= (-2.25 * (1e10 - (1e10 + v5)))
            _ = v19
        }
        p0 = "foo bar"
        _ = fi6
    }
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 uint16 = 32767
    v4 -= (v4 + 2)
    fmt.Println(v4)
    delete(v1, "")
    var v20 bool = true
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS2{F0: v4}
        for fi23 := 0; fi23 < 17; fi23++ {
            if (4294967295 <= 42) {
                v24 := []int{fi21, 32768, fi23, fi21}
                switch fi23 {
                case 632:
                    var v25 int = fi23
                    _ = v25
                }
                _ = v24
            } else {
                _ = len(v3)
                v4 -= v4
            }
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    var v26 int16 = -730
    v20 = ((!((!(180 != 136)) && v20)) || true)
    if v20 {
        if ("hello" <= ((("\n\t" + "a") + "foo bar") + "")) {
            for fi27 := 0; fi27 < 3; fi27++ {
                v26 -= v26
                fmt.Println(v4, v4)
                v28 := []int{fi27}
                _ = fi27
                _ = v28
            }
            v29 := []int{-103, 1, 32767, -2147483648}
            v30 := GenS1{F0: (!(2 >= 0)), F1: 1, F2: 2147483648}
            _ = v29
            _ = v30
        }
        switch "foo bar" {
        case "a":
            v20 = true
        case "\n\t":
            fmt.Println(v4, v26)
        default:
            var v31 float32 = 1e10
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            if v20 {
                var v33 uint32 = 273
                _ = v33
            } else {
                for fi34 := 0; fi34 < 5; fi34++ {
                    _ = len(v2)
                    _ = fi34
                }
            }
            _ = fi32
        }
    }
    var v35 bool = (42 < 18)
    if v20 {
        v36 := []int{}
        _ = v36
    } else {
        var v37 string = "hello"
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v35
    _ = v4
}
