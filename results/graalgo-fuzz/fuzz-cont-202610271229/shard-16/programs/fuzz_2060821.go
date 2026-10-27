package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint32, p1 uint, p2 int16) uint32 {
    if (!((!true) && (!(("\n\t" != "a") || true)))) {
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                p2 = (-(p2))
                v8 := []int{(fi6 * 2147483647)}
                if true {
                    p2 = p2
                    var v9 int32 = 42
                    _ = v9
                } else {
                    v10 := GenS1{F0: 42, F1: (p0 << ((255) & 63)), F2: (p1 * 65535)}
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        var v11 int = -2147483648
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    var v14 uint32 = (p0 % (((p0 % ((42) | 1))) | 1))
    var v15 uint64 = 511
    fmt.Println(p0, p0)
    switch ("a" + ("a" + ("" + "a"))) {
    case "":
        switch 32767 {
        case 0:
            var v16 int16 = (p2 + p2)
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    var v3 bool = (true || false)
    var v4 int16 = -32768
    var v5 int16 = (v4 % (((v4 * (v4 - (^(v4))))) | 1))
    v5 = (v5 - v4)
    switch 4294967295 {
    case 133:
        v5 = v5
    case -2147483648:
        var v17 bool = true
        _ = v17
    }
    v2 = ((278 > 65535) || (false || (!(!((!true) || ((v4 ^ v5) < v5))))))
    if v2 {
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS2{F0: 127}
            var v20 int8 = 100
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS2{F0: 127}
        _ = v21
    } else {
        switch ("hello" + "hello") {
        case "":
            var v22 int16 = v5
            _ = v22
        case "hello":
            v2 = ((42 >= genFn0(2147483648, 18446744073709551615, v5)) && v2)
        case "foo bar":
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 float32 = ((1.5 + ((1e10 * 0.1) / 0.1)) / 1e10)
                _ = fi23
                _ = v24
            }
        }
        var v25 uint64 = 4294967295
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        if (65535 != 255) {
            switch fi26 {
            case -2147483648:
                if true {
                    var v27 float64 = 0.1
                    v28 := []int{}
                    _ = v27
                    _ = v28
                }
            case 675:
                if v3 {
                    var v29 uint64 = 1
                    for fi30 := 0; fi30 < 9; fi30++ {
                        if (v3 || (!(0 != 0))) {
                            if v3 {
                                fmt.Println(fi26, v3)
                                v31 := GenS1{F0: (v29 & (v29 ^ (^(v29)))), F1: 65535, F2: 0}
                                _ = v31
                            } else {
                                var v32 uint32 = 2147483648
                                switch (fi30 * 42) {
                                case 32767:
                                    _ = len(v1)
                                }
                                fmt.Println(fi26)
                                _ = v32
                            }
                            v1["hello"] = false
                            v33 := GenS1{F0: v29, F1: 2147483648, F2: 255}
                            _ = v33
                        }
                        _ = fi30
                    }
                    _ = v29
                }
            default:
                var v34 uint32 = 2
                _ = v34
            }
        }
        v4 -= -32768
        _ = fi26
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        v35 := GenS2{F0: 42}
        _ = v35
    }
    v4 = v4
    v2 = (!((^(v5)) == 42))
    fmt.Println(v3, v5)
    var v36 uint32 = 2
    switch (("foo bar" + "hello") + "foo bar") {
    case "foo bar":
        switch 28 {
        case 127:
            fmt.Println(v3, v3, v3)
        }
    case "a":
        var v37 int8 = 0
        _ = v37
    case "hello":
        var v38 int16 = (-(v5))
        _ = v38
    default:
        v39 := []int{int(v36)}
        _ = v39
    }
    v36 = v36
    fmt.Println(v5, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
