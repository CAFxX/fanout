package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
func genFn0(p0 string, p1 int) bool {
    v7 := []int{p1, p1, (^(p1))}
    var v8 uint64 = 282
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (!(153 <= 1))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 42
    var v5 uint64 = 44
    var v6 bool = true
    var v9 uint = 2
    var v10 uint = 83
    fmt.Println(v9, v10, v4)
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: -922}
        v13 := GenS2{F0: 0, F1: (fi11 << ((v9) & 63)), F2: 0}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 823}
    v10 += v10
    var v15 string = ""
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint = v9
        switch ((v15 + ("\n\t" + ("hello" + v15))) + "") {
        case "":
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 int = (^(fi18))
                _ = fi18
                _ = v19
            }
        default:
            var v20 uint16 = 418
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v21 := GenS2{F0: 42, F1: 1, F2: -1}
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int = (^(fi22))
        v9 -= 42
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 int64 = -9223372036854775808
        if v6 {
            fmt.Println(v5, v9, v5)
            if ((v6 || (!((v5 << (((v9 ^ v9)) & 63)) > 18446744073709551615))) && v6) {
                fmt.Println(v6)
            } else {
                for fi26 := 0; fi26 < 12; fi26++ {
                    if (!(-2.25 <= (-2.25 / 1.5))) {
                        if (!(0 < 1)) {
                            v27 := v3[(v4 & uint32(v9))]
                            _ = v27
                        } else {
                            v28 := GenS2{F0: -1, F1: (fi24 + -1), F2: -125}
                            _ = v28
                        }
                        if (v6 || (215 != 1)) {
                            v4 = 750
                            v15 = "foo bar"
                        }
                    }
                    _ = fi26
                }
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi24
        _ = v25
    }
    v29 := GenS1{F0: -32768}
    if v6 {
        v30 := GenS2{F0: 903, F1: 1, F2: -2147483648}
        _ = len(v1)
        switch "a" {
        case "foo bar":
            v15 = ("\n\t" + v15)
        }
        _ = v30
    }
    var v31 int8 = -1
    for fi32 := 0; fi32 < 16; fi32++ {
        v5 = v5
        if v6 {
            var v33 float64 = 3.14159
            v5 -= 342
            v34 := GenS1{F0: 0}
            _ = v33
            _ = v34
        }
        _ = fi32
    }
    var v35 int16 = 32767
    var v36 uint64 = 51
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
