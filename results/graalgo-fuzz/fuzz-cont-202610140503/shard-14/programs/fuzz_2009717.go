package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8, p1 int64) bool {
    switch 204 {
    case 9223372036854775807:
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint32 = 0
            v9 := []int{4294967295, int(p1), fi7}
            v10 := GenS1{F0: v8}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        var v11 int16 = 32767
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        p0 += p0
        v13 := []int{}
        v14 := GenS2{F0: false, F1: 2}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int64 = p1
    var v16 uint64 = 1
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return false
}
func genFn1(p0 string, p1 int8, p2 uint16) uint8 {
    p2 = p2
    p2 = (^(p2))
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = (("\n\t" + "hello") + ("\n\t" + ("\n\t" + "")))
    var v3 uint32 = 687
    v4 := []uint32{v3}
    var v5 bool = false
    var v6 int32 = 0
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            if len(v4) > 0 {
                v4[0] = v3
            }
            _ = fi18
        }
        v2 = (v2 + ("\n\t" + "\n\t"))
        _ = fi17
    }
    fmt.Println(v2, v3, v5)
    v1[(v2 + (v2 + v2))] = v5
    for fi19 := 0; fi19 < 7; fi19++ {
        if true {
            switch (fi19 * fi19) {
            case 4294967295:
                v20 := GenS1{F0: v3}
                _ = v20
            case 156:
                v2 = ("a" + "\n\t")
            case 32767:
                var v21 uint64 = 18446744073709551615
                _ = v21
            default:
                var v22 int32 = v6
                _ = v22
            }
        } else {
            var v23 int = fi19
            _ = v23
        }
        var v24 uint8 = 42
        _ = fi19
        _ = v24
    }
    var v25 int32 = (v6 & v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v26 uint = 4294967295
    var v27 uint64 = 732
    if false {
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 3; fi29++ {
                if len(v4) > 0 {
                    v4[0] = (v3 & (v3 / ((v3) | 1)))
                }
                _ = fi29
            }
            v6 = v25
            var v30 float32 = float32(1.5)
            _ = fi28
            _ = v30
        }
        v1[(v2 + (v2 + v2))] = v5
    } else {
        if ((v3 / ((42) | 1)) != 4294967295) {
            var v31 bool = (!(32768 != -10))
            v25 = v6
            _ = v31
        } else {
            fmt.Println(v6)
            var v32 float64 = 3.14159
            _ = v32
        }
        _ = len(v1)
    }
    var v33 uint16 = 65535
    if (((v3 < 485) || ("a" <= "foo bar")) || false) {
        fmt.Println(v3, v25)
        var v34 uint8 = 0
        var v35 uint8 = uint8(v27)
        _ = v34
        _ = v35
    }
    var v36 int = -365
    switch (v36 ^ (v36 | (v36 + (v36 + (^(v36)))))) {
    case -1:
        v37 := GenS2{F0: v5, F1: uint32(v26)}
        _ = v37
    case -998:
        delete(v1, v2)
    }
    v38 := v1["a"]
    switch v2 {
    case "a":
        for fi39 := 0; fi39 < 14; fi39++ {
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 int64 = 1
                _ = fi40
                _ = v41
            }
            var v42 uint16 = v33
            _ = fi39
            _ = v42
        }
    default:
        if ((9223372036854775807 != 42) && v5) {
            switch 9223372036854775807 {
            case 1:
                v5 = (((-2.25 * 1.5) / 1e10) != 1.5)
            default:
                var v43 float32 = 0.1
                _ = v43
            }
            var v44 uint = v26
            var v45 bool = v5
            _ = v44
            _ = v45
        } else {
            v3 = v3
        }
    }
    var v46 uint64 = 846
    v47 := GenS2{F0: ((("a" + (("" + ((v2 + v2) + "foo bar")) + "hello")) + "") == v2), F1: uint32(v26)}
    for fi48 := 0; fi48 < 2; fi48++ {
        if (v38 && (false || true)) {
            _ = len(v4)
            for fi49 := 0; fi49 < 17; fi49++ {
                fmt.Println(v2, v36)
                _ = fi49
            }
        } else {
            var v50 uint = 64
            _ = v50
        }
        v4 = append(v4, 1)
        for fi51 := 0; fi51 < 17; fi51++ {
            v38 = v38
            _ = fi51
        }
        _ = fi48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v36
    _ = v38
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
