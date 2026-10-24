package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint16, p2 uint32) uint16 {
    v4 := GenS1{F0: (p1 & uint16(p0)), F1: p1, F2: (608 == 0)}
    p0 = p0
    v5 := GenS1{F0: 255, F1: 507, F2: true}
    switch -258 {
    case 4294967295:
        if true {
            v6 := GenS1{F0: 32767, F1: (p1 % ((32767) | 1)), F2: true}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 4; fi7++ {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 bool = true
                    var v10 uint = (p0 - (^(p0)))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
            if true {
                switch 32768 {
                case 32767:
                    var v11 string = ""
                    _ = v11
                case 17:
                    var v12 uint64 = 1
                    _ = v12
                case 128:
                    v13 := []int{339, 127}
                    _ = v13
                }
            } else {
                var v14 float64 = 0.1
                var v15 int64 = int64(p1)
                if ((("\n\t" + "hello") + ((("" + ("" + (("hello" + ((("\n\t" + "") + "hello") + "\n\t")) + ("" + "hello")))) + ("a" + "")) + "hello")) <= "\n\t") {
                    v15 -= (v15 << ((p0) & 63))
                } else {
                    if false {
                        var v16 uint16 = p1
                        _ = v16
                    }
                }
                _ = v14
                _ = v15
            }
            var v17 int = 0
            _ = v17
        }
    case 1:
        v18 := GenS1{F0: (p1 / (((p1 - (p1 % (((p1 / ((p1) | 1))) | 1)))) | 1)), F1: 1, F2: (42 > 240)}
        _ = v18
    case 0:
        v19 := []int{}
        _ = v19
    }
    v20 := []int{127, 369, -799, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v4
    _ = v5
    return 65535
}
func genFn1(p0 bool, p1 string, p2 uint16) string {
    var v21 float64 = 0.0
    for fi22 := 0; fi22 < 19; fi22++ {
        switch 32768 {
        case 9223372036854775807:
            var v23 float32 = (((1.5 * 1.5) - ((1e10 - 3.14159) / -2.25)) * ((0.1 * 0.1) + (-2.25 + 1.5)))
            _ = v23
        case 127:
            v24 := []int{}
            _ = v24
        }
        _ = fi22
    }
    fmt.Println(p1, v21)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return p1
}

func main() {
    v1 := []string{(("foo bar" + "") + "hello"), "hello"}
    var v2 uint32 = 0
    v3 := []uint32{v2, 1, v2}
    v3 = append(v3, (v2 % ((v2) | 1)))
    for fi25 := 0; fi25 < 12; fi25++ {
        v2 -= (v2 - 4294967295)
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS1{F0: 1, F1: genFn0(148, 0, 2147483647), F2: true}
            v28 := GenS2{F0: 830, F1: 65535}
            if false {
                v29 := GenS3{F0: 9223372036854775807, F1: (((!(!(42 >= 1))) || false) && false), F2: 503}
                var v30 bool = true
                v31 := GenS2{F0: uint(v2), F1: 2}
                _ = v29
                _ = v30
                _ = v31
            } else {
                v2 += v2
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    v2 = (^(v2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (((0.0 / 1.0) == 1.5) && false) {
        v2 = (v2 * (^(v2)))
    } else {
        var v32 string = "foo bar"
        var v33 int = 1
        _ = v32
        _ = v33
    }
    if (472 > 796) {
        fmt.Println(v2)
        if false {
            var v34 float32 = (1.5 - (0.1 / 1.0))
            for fi35 := 0; fi35 < 1; fi35++ {
                switch (^(fi35)) {
                case 0:
                    var v36 string = "\n\t"
                    _ = v36
                }
                v2 = (v2 ^ v2)
                v34 = float32(-2.25)
                _ = fi35
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v34
        } else {
            v2 = v2
            v2 -= v2
            var v37 uint32 = 2147483647
            _ = v37
        }
        fmt.Println(v2)
    }
    switch "hello" {
    case "":
        for fi38 := 0; fi38 < 9; fi38++ {
            v39 := GenS4{F0: v2}
            v40 := GenS2{F0: 4294967295, F1: genFn0(2, 1, v2)}
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    var v41 uint8 = 127
    if true {
        fmt.Println(v41, v41)
        var v42 int16 = 1000
        fmt.Println(v41, v42, v2)
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
}
