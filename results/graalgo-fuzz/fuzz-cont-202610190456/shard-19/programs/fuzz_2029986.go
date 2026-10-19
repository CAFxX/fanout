package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint16) int {
    for fi4 := 0; fi4 < 5; fi4++ {
        p0 += 42
        if (fi4 <= fi4) {
            var v5 int16 = 32767
            _ = v5
        }
        p0 = p0
        _ = fi4
    }
    var v6 uint32 = 98
    if true {
        fmt.Println(v6)
    }
    _ = p0
    _ = v6
    return -2147483648
}

func main() {
    v1 := []string{"hello"}
    var v2 float32 = ((0.1 - 0.0) - 1.5)
    var v3 uint16 = 32767
    var v7 float64 = (((0.0 / 3.14159) / 0.1) / 1.0)
    v8 := GenS1{F0: 0, F1: 0, F2: 523}
    v9 := GenS2{F0: v3, F1: 0}
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 string = ""
            if true {
                if (1 <= 4294967295) {
                    if true {
                        v7 = v7
                        v7 -= (v7 + 0.1)
                    }
                    var v13 uint = 955
                    _ = v13
                }
                if ("foo bar" > (v12 + ("a" + ("hello" + v12)))) {
                    var v14 int16 = 0
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v14
                }
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if false {
        var v15 int16 = 305
        _ = v15
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 string = "foo bar"
            for fi18 := 0; fi18 < 9; fi18++ {
                v3 += (v3 << ((4294967295) & 63))
                var v19 int8 = 127
                v17 = v17
                _ = fi18
                _ = v19
            }
            if (1 < 72) {
                _ = len(v1)
            } else {
                var v20 string = v17
                if false {
                    var v21 string = "a"
                    for fi22 := 0; fi22 < 2; fi22++ {
                        v23 := GenS3{F0: 0, F1: 1}
                        _ = fi22
                        _ = v23
                    }
                    v24 := GenS2{F0: v3, F1: -4}
                    _ = v21
                    _ = v24
                }
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v3, v3)
        var v25 int64 = -921
        _ = v25
    }
    if (!true) {
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 float64 = v7
            if false {
                v28 := GenS2{F0: v3, F1: -1}
                v27 = 0.1
                v29 := GenS3{F0: 303, F1: 0}
                _ = v28
                _ = v29
            }
            v30 := GenS3{F0: 0, F1: int16(v3)}
            _ = fi26
            _ = v27
            _ = v30
        }
        switch ("\n\t" + "hello") {
        case "foo bar":
            var v31 float32 = v2
            _ = v31
        case "\n\t":
            v3 += (v3 + (v3 + (v3 ^ v3)))
        case "":
            fmt.Println(v2, v3, v7)
        default:
            fmt.Println(v2, v2)
        }
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 string = "a"
            for fi34 := 0; fi34 < 9; fi34++ {
                v33 = v33
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v35 := GenS1{F0: 44, F1: 86, F2: 929}
                _ = len(v1)
                _ = fi34
                _ = v35
            }
            _ = fi32
            _ = v33
        }
    }
    var v36 uint32 = 42
    var v37 float32 = (v2 + float32(0.0))
    if (1 >= 42) {
        v38 := GenS1{F0: 626, F1: uint(v3), F2: 518}
        _ = v38
    }
    if len(v1) > 0 {
        v1[0] = ("a" + ("" + (("hello" + ("a" + "hello")) + "\n\t")))
    }
    if false {
        for fi39 := 0; fi39 < 20; fi39++ {
            if len(v1) > 0 {
                v1[0] = ("\n\t" + "")
            }
            _ = fi39
        }
        v2 -= (v2 / ((v2 * v2) * v2))
    }
    if (-1 >= 0) {
        v36 = 0
        for fi40 := 0; fi40 < 2; fi40++ {
            v3 += v3
            _ = fi40
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v7
    _ = v8
    _ = v9
}
