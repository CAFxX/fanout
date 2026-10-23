package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 int16 }
func genFn0() int64 {
    v7 := GenS1{F0: 556, F1: (42 >= 42), F2: 42}
    v8 := []int{32767, 544}
    switch "" {
    case "foo bar":
        if len(v8) > 0 {
            v8[0] = 127
        }
    default:
        var v9 int32 = 1
        _ = v9
    }
    v10 := GenS1{F0: 16, F1: true, F2: -548}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v10
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 string, p1 int8, p2 string) uint16 {
    p2 = p2
    if true {
        var v11 uint = 228
        fmt.Println(v11, p1)
        p0 = (("a" + "hello") + "\n\t")
        _ = v11
    } else {
        if true {
            v12 := []int{}
            p0 = "a"
            var v13 uint8 = 110
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 42, F1: (0 <= 1), F2: -378}
            v15 := GenS2{F0: 0}
            _ = v14
            _ = v15
        }
    }
    var v16 float64 = 0.1
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 32767
}

func main() {
    var v1 uint64 = 2
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    var v4 uint64 = v1
    var v5 int16 = 931
    var v6 bool = false
    var v17 uint64 = (v4 * 18446744073709551615)
    v18 := []int{-270}
    var v19 int32 = -891
    var v20 int8 = -106
    fmt.Println(v5, v20, v1)
    for fi21 := 0; fi21 < 2; fi21++ {
        v4 += 1
        var v22 uint16 = 2
        var v23 float64 = 0.1
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 = 0
    var v24 float64 = 0.1
    switch int(v2) {
    case 4294967295:
        v5 -= (v5 << ((255) & 63))
    case -120:
        if (32767 >= 2) {
            v3[1] = ((v24 * 0.1) <= v24)
            v25 := GenS2{F0: v5}
            var v26 int64 = -1
            _ = v25
            _ = v26
        } else {
            if v6 {
                v1 = (v1 % (((v17 % ((v17) | 1))) | 1))
                fmt.Println(v20, v1, v17)
            }
            var v27 int8 = (-(v20))
            var v28 int32 = v19
            _ = v27
            _ = v28
        }
    case -453:
        if true {
            if false {
                v2 = (v2 >> ((434) & 63))
            }
            switch -217 {
            case 9223372036854775807:
                v3[696] = (1 == 100)
            case 4294967295:
                if len(v18) > 0 {
                    v18[0] = -2147483648
                }
            default:
                v29 := GenS1{F0: genFn0(), F1: v6, F2: v19}
                _ = v29
            }
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 bool = (v6 || ("a" <= ("a" + ("\n\t" + "hello"))))
                _ = fi30
                _ = v31
            }
        }
    }
    var v32 float32 = (1.5 - 3.14159)
    switch ("foo bar" + "foo bar") {
    case "":
        for fi33 := 0; fi33 < 7; fi33++ {
            fmt.Println(v4)
            _ = fi33
        }
    case "a":
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 uint16 = 32767
            if false {
                fmt.Println(v5, v17, v24)
                switch "" {
                case "":
                    var v36 int64 = int64(v5)
                    _ = v36
                case "a":
                    for fi37 := 0; fi37 < 13; fi37++ {
                        if v6 {
                            _ = len(v18)
                        }
                        _ = fi37
                    }
                case "hello":
                    v2 = genFn0()
                }
            }
            switch "hello" {
            case "\n\t":
                fmt.Println(v4)
            case "hello":
                v19 += v19
            }
            _ = fi34
            _ = v35
        }
    case "hello":
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 bool = v6
            _ = fi38
            _ = v39
        }
    default:
        v40 := GenS2{F0: v5}
        _ = v40
    }
    v41 := GenS2{F0: v5}
    if len(v18) > 0 {
        v18[0] = 127
    }
    fmt.Println(v20, v6, v17)
    v42 := GenS2{F0: v5}
    switch ("a" + "") {
    case "foo bar":
        var v43 bool = ((193 >= 127) && v6)
        _ = v43
    case "\n\t":
        switch ("foo bar" + "a") {
        case "\n\t":
            for fi44 := 0; fi44 < 7; fi44++ {
                _ = len(v18)
                switch ("" + "a") {
                case "\n\t":
                    v45 := GenS1{F0: (v2 >> ((18446744073709551615) & 63)), F1: ((("\n\t" + "\n\t") + "") > "a"), F2: v19}
                    _ = v45
                }
                _ = fi44
            }
        }
    default:
        v19 = v19
    }
    delete(v3, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
