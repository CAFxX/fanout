package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 bool; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint8 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint8 }
func genFn0(p0 uint, p1 int) string {
    if true {
        fmt.Println(p1)
        fmt.Println(p1, p0)
    }
    switch (p1 & p1) {
    case -2147483648:
        for fi4 := 0; fi4 < 10; fi4++ {
            var v5 uint64 = 1
            v6 := GenS1{F0: p0, F1: 128}
            _ = fi4
            _ = v5
            _ = v6
        }
    case 127:
        var v7 uint16 = 42
        _ = v7
    case 2147483647:
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: p0, F1: -1}
            fmt.Println(p1, p0, p0)
            p0 += 537
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    return "\n\t"
}
func genFn1(p0 uint16, p1 int16, p2 bool) bool {
    if false {
        var v10 uint16 = p0
        v11 := []int{128, -9223372036854775808, 1}
        _ = v10
        _ = v11
    }
    v12 := []int{}
    if false {
        var v13 string = "foo bar"
        if (0 != 1) {
            var v14 uint16 = (p0 >> ((2) & 63))
            p2 = (0.0 >= 1.5)
            _ = v14
        } else {
            if ("\n\t" < v13) {
                v15 := GenS1{F0: 18446744073709551615, F1: 2147483647}
                for fi16 := 0; fi16 < 5; fi16++ {
                    if p2 {
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        default:
                            var v17 uint8 = 227
                            _ = v17
                        }
                        if (p1 < 1000) {
                            var v18 float64 = -2.25
                            if p2 {
                                p1 -= (-(p1))
                            } else {
                                v12 = append(v12, (fi16 >> ((415) & 63)))
                                v18 = 0.1
                            }
                            _ = v18
                        } else {
                            if false {
                                var v19 int8 = -1
                                var v20 int16 = p1
                                _ = v19
                                _ = v20
                            }
                            var v21 bool = p2
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 16; fi22++ {
                            v23 := GenS2{F0: p2, F1: 2147483647, F2: 2147483648}
                            _ = fi22
                            _ = v23
                        }
                    } else {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        fmt.Println(p2, p2)
                    }
                    var v24 uint16 = p0
                    _ = fi16
                    _ = v24
                }
                _ = v15
            } else {
                if (580 >= 925) {
                    v13 = ("" + "\n\t")
                    var v25 uint16 = p0
                    var v26 bool = p2
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = v13
    } else {
        switch 2147483647 {
        case 2147483647:
            p2 = false
        case 128:
            v27 := GenS3{F0: p0, F1: 128, F2: 139}
            _ = v27
        case 32767:
            if (false && p2) {
                var v28 int8 = -72
                p2 = p2
                _ = v28
            } else {
                _ = len(v12)
                var v29 uint = 369
                v30 := GenS2{F0: p2, F1: -887, F2: 2147483647}
                _ = v29
                _ = v30
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}

func main() {
    var v1 float32 = (1.5 + (0.0 + (-2.25 + 1.5)))
    var v2 uint16 = 160
    var v3 string = "hello"
    v2 -= (v2 % ((v2) | 1))
    var v31 uint16 = v2
    v31 += 238
    v32 := GenS4{F0: (v2 + 42), F1: (880 < 4294967295), F2: 99}
    if ((2 <= v31) || (!true)) {
        for fi33 := 0; fi33 < 20; fi33++ {
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := []int{(fi33 >> ((255) & 63)), (fi33 & 32768)}
                var v36 uint8 = 7
                var v37 int = 382
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            _ = fi33
        }
        var v38 uint32 = 568
        _ = v38
    } else {
        v31 += (v2 - 0)
        v2 = v31
        var v39 uint64 = 2
        _ = v39
    }
    var v40 uint = 18446744073709551615
    for fi41 := 0; fi41 < 3; fi41++ {
        switch (("" + v3) + "") {
        case "a":
            var v42 int16 = -1
            _ = v42
        }
        _ = fi41
    }
    v1 -= (0.0 / (v1 + 3.14159))
    if true {
        v43 := []int{-877, 32767}
        var v44 uint8 = 255
        for fi45 := 0; fi45 < 16; fi45++ {
            v44 += (^(v44))
            _ = fi45
        }
        _ = v43
        _ = v44
    }
    v40 -= (v40 % ((4294967295) | 1))
    if false {
        var v46 uint64 = 0
        var v47 int8 = 100
        _ = v46
        _ = v47
    }
    var v48 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v40
    _ = v48
}
