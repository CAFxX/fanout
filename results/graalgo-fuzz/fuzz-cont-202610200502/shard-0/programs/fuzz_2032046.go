package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 uint32) uint64 {
    v4 := []int{-1, 32768, 42}
    for fi5 := 0; fi5 < 19; fi5++ {
        if (42 < 0) {
            fmt.Println(fi5)
            var v6 int32 = -1
            _ = v6
        } else {
            var v7 uint8 = 42
            var v8 string = "hello"
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    p0 = 1
    p0 = 2147483647
    _ = p0
    _ = v4
    return 0
}

func main() {
    v1 := []uint32{65535, 9, 653, 232}
    var v2 uint16 = 679
    var v3 int = 42
    var v9 float64 = 1e10
    for fi10 := 0; fi10 < 18; fi10++ {
        switch "a" {
        case "a":
            switch (fi10 << ((63) & 63)) {
            case 42:
                for fi11 := 0; fi11 < 20; fi11++ {
                    if false {
                        var v12 uint32 = 1
                        v12 = 243
                        if true {
                            var v13 uint = 861
                            _ = v13
                        } else {
                            var v14 uint8 = 0
                            fmt.Println(v2, fi10, v9)
                            v15 := GenS1{F0: 1000, F1: 42}
                            _ = v14
                            _ = v15
                        }
                        _ = v12
                    }
                    var v16 int64 = 404
                    var v17 string = "foo bar"
                    _ = fi11
                    _ = v16
                    _ = v17
                }
            case 1:
                switch 263 {
                case 1:
                    if false {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v1 = append(v1, uint32(v2))
                            v9 -= v9
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                            _ = fi18
                        }
                    }
                }
            default:
                fmt.Println(v2, v9)
            }
        case "":
            v9 = v9
        default:
            var v19 float32 = 0.0
            _ = v19
        }
        if ((fi10 + fi10) != (fi10 | (fi10 - fi10))) {
            v3 += fi10
            v20 := GenS1{F0: -1, F1: 42}
            var v21 uint32 = 618
            _ = v20
            _ = v21
        }
        _ = fi10
    }
    switch "foo bar" {
    case "":
        if true {
            var v22 float64 = (float64(-2.25) + 1e10)
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(v22)
                for fi24 := 0; fi24 < 14; fi24++ {
                    _ = len(v1)
                    v25 := GenS2{F0: "a", F1: "\n\t"}
                    var v26 int64 = -361
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
            v27 := GenS2{F0: (("hello" + "hello") + ""), F1: ""}
            _ = v22
            _ = v27
        } else {
            var v28 int64 = -630
            _ = v28
        }
    default:
        var v29 string = "hello"
        _ = v29
    }
    var v30 uint8 = 89
    if (true && false) {
        var v31 string = "a"
        if (!(v3 >= v3)) {
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 int8 = 42
                _ = len(v1)
                _ = fi32
                _ = v33
            }
        }
        _ = v31
    }
    var v34 int32 = -584
    v2 = v2
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if (false || true) {
            v1 = append(v1, 2147483648)
            fmt.Println(v9, v30)
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int8 = -1
                var v37 uint64 = 1
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        fmt.Println(v2, v3)
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            var v40 uint8 = 177
            _ = fi39
            _ = v40
        }
        var v41 uint64 = 2
        _ = fi38
        _ = v41
    }
    var v42 float32 = 3.14159
    if false {
        v3 -= v3
    }
    v42 -= v42
    v43 := GenS1{F0: 1000, F1: 4294967295}
    var v44 int32 = (v34 - v34)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v42
    _ = v43
    _ = v44
    _ = v9
}
