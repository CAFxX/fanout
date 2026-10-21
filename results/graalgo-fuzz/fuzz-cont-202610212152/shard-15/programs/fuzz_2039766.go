package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int8 {
    var v4 float64 = 0.0
    var v5 int = -914
    var v6 int64 = -1
    v7 := GenS1{F0: 773}
    fmt.Println(v5)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 uint32) uint64 {
    if (true || false) {
        if (-1 == -865) {
            switch "foo bar" {
            case "":
                switch 1 {
                case -964:
                    v8 := []int{-158, -1}
                    _ = v8
                default:
                    var v9 string = "hello"
                    _ = v9
                }
            case "foo bar":
                for fi10 := 0; fi10 < 4; fi10++ {
                    if (82 > 127) {
                        p0 -= -975
                        p1 = (^(p1))
                    }
                    _ = fi10
                }
            case "a":
                switch "\n\t" {
                case "hello":
                    switch -1 {
                    case 9223372036854775807:
                        var v11 int64 = -9223372036854775808
                        _ = v11
                    case 1:
                        v12 := []int{}
                        _ = v12
                    }
                case "":
                    var v13 float32 = 1.5
                    _ = v13
                case "foo bar":
                    if (115 > 32768) {
                        var v14 int = 32768
                        p1 = (^(p1))
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 2; fi15++ {
                            v16 := GenS1{F0: -1}
                            fmt.Println(p1)
                            _ = fi15
                            _ = v16
                        }
                        var v17 int16 = (p0 ^ (p0 / ((p0) | 1)))
                        v18 := GenS1{F0: -959}
                        _ = v17
                        _ = v18
                    }
                }
            default:
                p0 = (p0 ^ p0)
            }
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            p1 += p1
            v21 := []int{fi20, fi19, (fi19 | (fi19 ^ (fi20 | fi19))), (fi19 * (fi19 + (fi20 ^ 543)))}
            if (((3.14159 / (1.5 * float32(-2.25))) - (0.1 + 1e10)) != -2.25) {
                if len(v21) > 0 {
                    v21[0] = fi19
                }
                var v22 uint8 = 2
                _ = v22
            } else {
                var v23 uint = 557
                _ = v23
            }
            _ = fi20
            _ = v21
        }
        fmt.Println(p1, fi19)
        p0 += (p0 * (p0 | p0))
        _ = fi19
    }
    fmt.Println(p0, p0)
    var v24 uint16 = 2
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS1{F0: 1}
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v24
    return 42
}

func main() {
    var v1 uint8 = 250
    var v2 uint16 = 32767
    var v3 int8 = 88
    fmt.Println(v1, v3)
    fmt.Println(v1)
    v3 += v3
    v27 := GenS1{F0: 1}
    for fi28 := 0; fi28 < 7; fi28++ {
        for fi29 := 0; fi29 < 5; fi29++ {
            v2 = v2
            var v30 string = "a"
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    switch "\n\t" {
    case "\n\t":
        v3 = 98
    }
    if true {
        if ((v3 ^ 1) < (v3 | 127)) {
            if (1.5 < -2.25) {
                v31 := GenS1{F0: -1}
                _ = v31
            }
            v3 = genFn0()
        }
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 uint32 = 942
            v34 := []int{(^(fi32))}
            v1 = v1
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 int16 = int16(v3)
    v2 = (v2 ^ (v2 | v2))
    v35 = 232
    v1 = v1
    switch 32767 {
    case 4294967295:
        var v36 int16 = v35
        _ = v36
    case -2147483648:
        v37 := []int{832, 42, 128}
        _ = v37
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v35
}
