package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 string, p2 string) uint32 {
    var v6 int64 = -1
    if false {
        for fi7 := 0; fi7 < 18; fi7++ {
            if (0.1 > (-2.25 / 1.0)) {
                if false {
                    fmt.Println(p1, fi7, fi7)
                } else {
                    var v8 int = fi7
                    v9 := GenS1{F0: p0, F1: "\n\t", F2: 42}
                    var v10 bool = ((v6 - v6) < v6)
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
            _ = fi7
        }
        v11 := []int{563, -604, 127, -59}
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 string = "\n\t"
            p2 = p1
            p2 = "foo bar"
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 875
}
func genFn1(p0 int16) uint8 {
    var v14 int32 = 2147483647
    switch "" {
    case "\n\t":
        var v15 int8 = 127
        _ = v15
    }
    v14 -= v14
    v16 := GenS1{F0: false, F1: ("hello" + (("foo bar" + "\n\t") + "foo bar")), F2: 871}
    var v17 bool = false
    _ = p0
    _ = v14
    _ = v16
    _ = v17
    return 30
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 uint16 = 528
    v4 := []string{"a", ("\n\t" + ("hello" + "")), "a", "\n\t"}
    var v5 int64 = -1
    var v18 int16 = 1000
    var v19 float32 = 3.14159
    if true {
        var v20 int = -106
        _ = v20
    } else {
        var v21 uint8 = v1
        if true {
            var v22 int = 32767
            v23 := v2[("a" + "")]
            for fi24 := 0; fi24 < 10; fi24++ {
                fmt.Println(v22)
                _ = fi24
            }
            _ = v22
            _ = v23
        }
        var v25 int32 = 80
        _ = v21
        _ = v25
    }
    fmt.Println(v3, v5)
    if true {
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 int32 = -264
            _ = fi26
            _ = v27
        }
    }
    v19 = 1e10
    for fi28 := 0; fi28 < 9; fi28++ {
        switch (-(fi28)) {
        case 32767:
            var v29 int32 = 42
            _ = v29
        case 127:
            var v30 float32 = float32((-2.25 - 1.5))
            _ = v30
        default:
            v31 := v2["\n\t"]
            _ = v31
        }
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 int32 = 2147483647
            v19 = (1e10 * (((-2.25 - (v19 + (v19 * v19))) * v19) / v19))
            _ = fi32
            _ = v33
        }
        if true {
            for fi34 := 0; fi34 < 18; fi34++ {
                var v35 uint8 = (v1 | v1)
                v35 = v35
                v5 += v5
                _ = fi34
                _ = v35
            }
            if true {
                switch fi28 {
                case 1:
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                case 0:
                    v18 += v18
                }
                v36 := GenS2{F0: -120}
                _ = v36
            }
            var v37 int16 = -32768
            _ = v37
        }
        _ = fi28
    }
    var v38 float64 = 1e10
    if len(v4) > 0 {
        v4[0] = (("a" + "foo bar") + "")
    }
    v38 = (3.14159 * 3.14159)
    v38 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
