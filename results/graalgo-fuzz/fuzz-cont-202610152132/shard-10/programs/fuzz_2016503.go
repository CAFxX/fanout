package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint8; F1 int64; F2 int16 }
type GenS4 struct { F0 int; F1 uint8; F2 uint16 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int) uint64 {
    switch 32768 {
    case 42:
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                for fi8 := 0; fi8 < 20; fi8++ {
                    fmt.Println(p1, fi7)
                    var v9 uint64 = 0
                    fmt.Println(v9, p1)
                    _ = fi8
                    _ = v9
                }
                fmt.Println(p1, fi7, fi7)
                v10 := []int{}
                _ = v10
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                switch 32767 {
                case 0:
                    var v12 int = (p1 & 128)
                    _ = v12
                default:
                    v13 := []int{fi11}
                    if (1 < -972) {
                        v14 := GenS1{F0: 1, F1: 42}
                        fmt.Println(fi7)
                        _ = len(v13)
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 1, F1: -32768}
                        var v16 uint16 = 42
                        _ = v15
                        _ = v16
                    }
                    var v17 int16 = -32768
                    _ = v13
                    _ = v17
                }
                p0 = p0
                _ = fi11
            }
            _ = fi7
        }
    case 32768:
        v18 := GenS2{F0: 1, F1: ("\n\t" + "a")}
        _ = v18
    case 128:
        v19 := []int{p1, p1, (p1 >> ((65535) & 63)), p1}
        _ = v19
    default:
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(p0, p0, p1)
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := []int{127, fi20, (fi21 % (((fi21 & (fi20 * fi21))) | 1)), 56}
                var v23 uint16 = 948
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := []int{(fi20 / ((p1) | 1))}
            _ = fi20
            _ = v24
        }
    }
    v25 := GenS3{F0: (p0 / ((255) | 1)), F1: -1, F2: 42}
    var v26 int64 = 9223372036854775807
    var v27 uint8 = (^(p0))
    p0 -= (p0 << ((42) & 63))
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    _ = v27
    return 18
}
func genFn1(p0 int64, p1 string) bool {
    if false {
        for fi28 := 0; fi28 < 10; fi28++ {
            if true {
                v29 := []int{}
                _ = v29
            }
            var v30 uint = 2
            _ = fi28
            _ = v30
        }
    } else {
        fmt.Println(p1, p1)
        p0 = p0
        v31 := []int{}
        _ = v31
    }
    p0 -= p0
    for fi32 := 0; fi32 < 2; fi32++ {
        v33 := GenS4{F0: (fi32 / ((fi32) | 1)), F1: 128, F2: 255}
        for fi34 := 0; fi34 < 20; fi34++ {
            p0 = (^(p0))
            v35 := []int{fi32}
            p0 += (^(p0))
            p0 -= p0
            var v36 uint8 = 128
            _ = fi34
            _ = v35
            _ = v36
        }
        p1 = "a"
        _ = fi32
        _ = v33
    }
    switch ((p1 + "") + "\n\t") {
    case "foo bar":
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 string = ((p1 + p1) + ((("a" + p1) + p1) + ((("foo bar" + p1) + "") + p1)))
            v38 = (((((p1 + ("a" + "")) + v38) + v38) + ("foo bar" + (p1 + ((p1 + ("\n\t" + "")) + "foo bar")))) + "hello")
            _ = fi37
            _ = v38
        }
    default:
        if true {
            var v39 float32 = (3.14159 * 0.1)
            var v40 float64 = (0.1 * float64((0.1 - (-2.25 / (v39 / 1.0)))))
            v41 := GenS3{F0: 253, F1: (p0 >> ((0) & 63)), F2: -32768}
            _ = v39
            _ = v40
            _ = v41
        } else {
            var v42 string = (((((p1 + "") + p1) + "") + p1) + p1)
            v43 := GenS5{F0: "", F1: -2147483648}
            _ = v42
            _ = v43
        }
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]bool)
    var v3 float64 = 3.14159
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    var v6 uint64 = 346
    v44 := []int{int(v1), 2147483647, -381}
    fmt.Println(v3)
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    switch (((("foo bar" + "") + "") + "a") + ("\n\t" + "hello")) {
    case "a":
        delete(v2, "\n\t")
    case "":
        fmt.Println(v6)
    case "foo bar":
        if true {
            var v45 float64 = v3
            _ = v45
        } else {
            delete(v5, "foo bar")
            v6 += (v6 + v6)
            v46 := GenS3{F0: 1, F1: -1, F2: 1000}
            _ = v46
        }
    default:
        v3 = 1e10
    }
    switch -2147483648 {
    case -185:
        fmt.Println(v1, v3)
    case 32767:
        v6 = v6
    case 9223372036854775807:
        if (!false) {
            for fi47 := 0; fi47 < 10; fi47++ {
                var v48 uint8 = 203
                fmt.Println(v3, v3, v48)
                _ = fi47
                _ = v48
            }
            var v49 int32 = 822
            _ = v49
        }
    }
    v3 = (v3 - ((1.5 - v3) + 0.1))
    v6 -= v6
    for fi50 := 0; fi50 < 10; fi50++ {
        var v51 int16 = 42
        _ = fi50
        _ = v51
    }
    var v52 uint = 2
    for fi53 := 0; fi53 < 13; fi53++ {
        var v54 int = fi53
        _ = fi53
        _ = v54
    }
    fmt.Println(v3, v3, v52)
    _ = len(v4)
    if (1e10 == 0.0) {
        var v55 int64 = 1
        _ = v55
    }
    v44 = append(v44, -52)
    v56 := GenS6{F0: genFn0(228, 531)}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v52
    _ = v56
    _ = v6
}
