package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint32; F2 int }
type GenS5 struct { F0 uint16; F1 string }
func genFn0(p0 int, p1 string) uint8 {
    v7 := []int{p0, p0, (p0 & p0), p0}
    if len(v7) > 0 {
        v7[0] = p0
    }
    v8 := GenS1{F0: 128, F1: true, F2: 2}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int32, p1 int8) int64 {
    var v9 int = 492
    fmt.Println(p1, p1)
    v10 := GenS2{F0: 1}
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 bool = (p0 != (-(p0)))
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 float32 = 0.1
    var v2 float32 = (((v1 - (v1 / v1)) + (v1 / float32((((-2.25 + 3.14159) - ((1e10 - 1e10) + (1e10 + 1e10))) - (float64(float32(-2.25)) + -2.25))))) / v1)
    var v3 uint32 = 4294967295
    var v4 int8 = 11
    var v5 int64 = 1
    v6 := make(map[int]int)
    for fi13 := 0; fi13 < 16; fi13++ {
        switch "hello" {
        case "\n\t":
            switch ("" + ("" + (("\n\t" + "\n\t") + "foo bar"))) {
            case "hello":
                var v14 uint = 2
                _ = v14
            case "":
                fmt.Println(v4)
            }
        case "hello":
            v15 := GenS2{F0: 255}
            _ = v15
        case "a":
            if ((((0.0 != 0.0) && (65535 >= 42)) && false) && false) {
                for fi16 := 0; fi16 < 9; fi16++ {
                    fmt.Println(v2)
                    if (42 >= genFn0((fi13 & int(v5)), "foo bar")) {
                        v17 := []int{}
                        v18 := GenS1{F0: 100, F1: (89 <= 1000), F2: 127}
                        _ = v17
                        _ = v18
                    }
                    var v19 float32 = v1
                    _ = fi16
                    _ = v19
                }
            }
        }
        _ = fi13
    }
    fmt.Println(v3, v1)
    v4 += v4
    fmt.Println(v4, v5)
    switch ("a" + "hello") {
    case "hello":
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 uint8 = 127
            switch "hello" {
            case "\n\t":
                var v22 uint8 = v21
                _ = v22
            case "":
                var v23 uint32 = v3
                _ = v23
            case "a":
                if true {
                    v21 = 2
                    for fi24 := 0; fi24 < 15; fi24++ {
                        if true {
                            var v25 int8 = (-(v4))
                            _ = v25
                        }
                        _ = fi24
                    }
                    if (true || true) {
                        var v26 uint = 2
                        v27 := []int{fi20, 0}
                        for fi28 := 0; fi28 < 6; fi28++ {
                            var v29 int8 = -128
                            v30 := GenS2{F0: (v21 - (v21 << ((953) & 63)))}
                            var v31 uint16 = uint16(v3)
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = v26
                        _ = v27
                    } else {
                        for fi32 := 0; fi32 < 11; fi32++ {
                            if false {
                                var v33 uint64 = 305
                                v34 := v6[(fi32 / ((fi32) | 1))]
                                fmt.Println(v21, fi20)
                                _ = v33
                                _ = v34
                            } else {
                                fmt.Println(v2, v4)
                            }
                            v35 := GenS2{F0: (v21 - v21)}
                            switch "" {
                            case "foo bar":
                                v36 := GenS3{F0: "foo bar"}
                                _ = v36
                            }
                            _ = fi32
                            _ = v35
                        }
                    }
                }
            default:
                var v37 int32 = 0
                _ = v37
            }
            _ = fi20
            _ = v21
        }
    case "":
        if (!(!true)) {
            var v38 float64 = 1.5
            _ = v38
        } else {
            v39 := GenS3{F0: "foo bar"}
            _ = v39
        }
    }
    if (-2147483648 > 32767) {
        v1 = v1
        if false {
            var v40 uint64 = uint64(v3)
            _ = v40
        } else {
            for fi41 := 0; fi41 < 9; fi41++ {
                fmt.Println(v4, v3)
                var v42 uint8 = 1
                v43 := GenS4{F0: (838 == 187), F1: v3, F2: fi41}
                _ = fi41
                _ = v42
                _ = v43
            }
        }
        var v44 int64 = (v5 | v5)
        _ = v44
    }
    var v45 int32 = 2147483647
    for fi46 := 0; fi46 < 8; fi46++ {
        delete(v6, -1)
        v47 := GenS5{F0: 65535, F1: ("foo bar" + "\n\t")}
        _ = fi46
        _ = v47
    }
    var v48 float32 = v1
    if false {
        fmt.Println(v5)
        v49 := []int{841, int(v5), -2147483648}
        var v50 string = "hello"
        _ = v49
        _ = v50
    } else {
        var v51 bool = true
        v52 := GenS2{F0: 1}
        _ = v51
        _ = v52
    }
    var v53 int64 = (v5 & v5)
    v54 := []int{9223372036854775807}
    if false {
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
