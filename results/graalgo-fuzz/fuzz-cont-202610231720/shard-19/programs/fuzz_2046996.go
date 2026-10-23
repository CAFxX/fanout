package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 string; F2 int }
func genFn0(p0 uint32, p1 bool) uint64 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint64 = 856
        var v9 uint16 = 137
        v8 -= (v8 & (v8 + 4294967295))
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{1, 4294967295, 0}
    for fi11 := 0; fi11 < 9; fi11++ {
        p1 = false
        _ = fi11
    }
    var v12 uint64 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return v12
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = float32((((3.14159 / 1e10) + -2.25) - (((3.14159 * -2.25) - (0.1 * 1e10)) + (((((0.1 / 1.0) - 0.0) + 0.0) - 3.14159) + 1e10))))
    var v3 int32 = -788
    v4 := []int{int(v3)}
    var v5 int32 = v3
    v6 := make(map[uint32]bool)
    fmt.Println(v3, v1, v2)
    var v13 int = 127
    v3 -= (v5 ^ -110)
    _ = len(v4)
    v14 := GenS1{F0: (v3 + (v5 & v5)), F1: 234, F2: 9223372036854775807}
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS2{F0: 0, F1: "a", F2: (-(v13))}
            v18 := GenS1{F0: -891, F1: 65535, F2: 42}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    v1 = (v1 >> ((4294967295) & 63))
    if false {
        v19 := v6[1]
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 string = (("" + (("" + "\n\t") + "hello")) + "foo bar")
            if (!v19) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    fmt.Println(v1, v3, v21)
                    var v23 int = 32768
                    if true {
                        v23 = -242
                        v2 = (v2 / 0.1)
                        var v24 float64 = 0.0
                        _ = v24
                    } else {
                        v25 := GenS2{F0: 42, F1: "", F2: (v23 | v13)}
                        v26 := GenS1{F0: v3, F1: 348, F2: -9223372036854775808}
                        _ = v25
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                }
                v21 = v21
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    } else {
        if (1000 > 969) {
            fmt.Println(v1)
        }
        for fi27 := 0; fi27 < 1; fi27++ {
            v13 -= 32768
            _ = fi27
        }
    }
    if ((((("" + ("hello" + "foo bar")) + ("hello" + "")) + "\n\t") + ("foo bar" + "foo bar")) != "") {
        switch 9223372036854775807 {
        case 0:
            v28 := GenS1{F0: v5, F1: 42, F2: 0}
            _ = v28
        case 562:
            v29 := GenS1{F0: (-(v3)), F1: 950, F2: -900}
            _ = v29
        case 32768:
            var v30 bool = true
            _ = v30
        }
        if ((!false) || (-2.25 >= float32((-2.25 - -2.25)))) {
            for fi31 := 0; fi31 < 16; fi31++ {
                fmt.Println(v5, v1)
                _ = fi31
            }
            v13 += (v13 / ((v13) | 1))
        }
    } else {
        var v32 int = (v13 & v13)
        _ = v32
    }
    v3 -= v5
    delete(v6, 42)
    var v33 int64 = -9223372036854775808
    v3 += (v5 >> ((686) & 63))
    var v34 int64 = (v33 | -881)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
