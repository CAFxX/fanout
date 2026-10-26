package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint8) uint8 {
    fmt.Println(p0)
    v4 := GenS1{F0: p0, F1: 31}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 bool = ((p0 >> ((1) & 63)) > (p0 * p0))
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = 88
    v3 := make(map[uint32]int)
    if true {
        v1 = (^(v1))
        var v8 float64 = 0.0
        _ = v8
    } else {
        v9 := GenS1{F0: genFn0(v2), F1: (v1 * 2147483647)}
        v2 -= genFn0(237)
        _ = v9
    }
    if false {
        if (1 >= -32768) {
            for fi10 := 0; fi10 < 20; fi10++ {
                if false {
                    switch (fi10 << ((18446744073709551615) & 63)) {
                    case -404:
                        v1 = v1
                    }
                } else {
                    switch "" {
                    case "a":
                        for fi11 := 0; fi11 < 5; fi11++ {
                            v12 := GenS1{F0: (v2 * 127), F1: v1}
                            _ = fi11
                            _ = v12
                        }
                    case "foo bar":
                        fmt.Println(fi10, v1)
                    }
                }
                _ = fi10
            }
            var v13 string = ("\n\t" + "a")
            v14 := GenS2{F0: -1}
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS1{F0: (^(v2)), F1: v1}
                var v17 uint64 = 0
                _ = fi15
                _ = v16
                _ = v17
            }
            switch -877 {
            case -1:
                var v18 float64 = 1e10
                _ = v18
            default:
                v19 := []int{9223372036854775807, int(v2), -55, 506}
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v1, v1, v1)
            var v21 uint16 = 42
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS2{F0: 127}
        var v24 float64 = 1e10
        _ = fi22
        _ = v23
        _ = v24
    }
    if (!false) {
        v25 := GenS2{F0: 100}
        v26 := GenS3{F0: 0}
        _ = v25
        _ = v26
    } else {
        v27 := []int{-9223372036854775808, 32767, 1}
        v1 = v1
        _ = len(v27)
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            if true {
                v30 := GenS3{F0: 216}
                _ = v30
            }
            v31 := GenS4{F0: 18446744073709551615}
            var v32 string = "\n\t"
            if (-2.25 <= -2.25) {
                var v33 float32 = (1e10 / 1.0)
                _ = v33
            } else {
                v1 += v1
                v32 = ""
                v32 = "hello"
            }
            v32 = "a"
            _ = fi29
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        if true {
            var v35 uint32 = uint32(v2)
            var v36 int8 = 0
            _ = v35
            _ = v36
        }
        var v37 uint64 = uint64(v2)
        var v38 int = fi34
        _ = fi34
        _ = v37
        _ = v38
    }
    if true {
        v1 += v1
        var v39 float64 = 1.5
        fmt.Println(v1, v1, v2)
        _ = v39
    }
    var v40 float64 = (1e10 * (0.1 + (-2.25 * 1.5)))
    switch -1 {
    case 0:
        v2 -= v2
    case -9223372036854775808:
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := []int{}
            _ = fi41
            _ = v42
        }
    }
    v2 = 42
    switch "foo bar" {
    case "a":
        for fi43 := 0; fi43 < 2; fi43++ {
            if (false && (42 <= 0)) {
                v44 := GenS2{F0: 127}
                _ = v44
            } else {
                var v45 uint64 = 2
                var v46 int = fi43
                _ = v45
                _ = v46
            }
            if (1 < 0) {
                v47 := []int{}
                _ = v47
            }
            fmt.Println(fi43, fi43, v1)
            _ = fi43
        }
    case "foo bar":
        var v48 int8 = -101
        _ = v48
    }
    v1 = v1
    v40 += (-2.25 / v40)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
}
