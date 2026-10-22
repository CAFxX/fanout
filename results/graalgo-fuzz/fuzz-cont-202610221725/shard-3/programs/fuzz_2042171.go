package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS5 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 int64) int32 {
    var v4 uint64 = 42
    v5 := []int{0, -487, 332, 4294967295}
    if true {
        v6 := GenS1{F0: 2, F1: 1000, F2: 127}
        switch "foo bar" {
        case "foo bar":
            v4 -= 1
        case "":
            var v7 bool = false
            _ = v7
        case "hello":
            fmt.Println(p0)
        }
        _ = v6
    } else {
        if true {
            v4 += (v4 | 0)
            if false {
                var v8 bool = true
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            switch (fi9 >> ((162) & 63)) {
            case 640:
                p0 = (p0 - (p0 << ((514) & 63)))
            case 32768:
                fmt.Println(v4, p0)
            case 34:
                var v10 uint16 = 65535
                _ = v10
            default:
                if (true || (65535 == 1)) {
                    fmt.Println(p0)
                }
            }
            var v11 int64 = p0
            _ = fi9
            _ = v11
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v12 uint16 = 42
        _ = v12
    default:
        switch 2147483647 {
        case 1:
            v4 = 4294967295
        default:
            if (!(!true)) {
                if true {
                    fmt.Println(v4)
                    if (!(1e10 != -2.25)) {
                        var v13 uint16 = 255
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            for fi15 := 0; fi15 < 10; fi15++ {
                                v5 = append(v5, (^(fi15)))
                                v16 := GenS1{F0: 65, F1: 32767, F2: 0}
                                _ = fi15
                                _ = v16
                            }
                            _ = fi14
                        }
                    }
                }
                fmt.Println(p0)
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS1{F0: 1, F1: 42, F2: 101}
                    var v19 uint64 = 18446744073709551615
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            } else {
                _ = len(v5)
                p0 -= (p0 ^ 1)
            }
        }
    }
    var v20 uint8 = 128
    _ = p0
    _ = v20
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint16, p1 int8, p2 int64) uint {
    p1 += p1
    fmt.Println(p1, p2, p0)
    if (!true) {
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := []int{fi21, (^(fi21)), (fi21 * fi21), fi21}
            v23 := GenS2{F0: p1}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v24 := []int{}
    if (float64(0.0) > 3.14159) {
        fmt.Println(p1)
        v25 := GenS3{F0: 1, F1: false}
        if len(v24) > 0 {
            v24[0] = 1
        }
        _ = v25
    } else {
        p0 += (p0 * p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    v3 := make(map[string]uint32)
    var v26 string = ""
    if (255 >= 65535) {
        v26 = ((("" + (v26 + v26)) + "\n\t") + v26)
        if (!(577 > -9223372036854775808)) {
            v27 := GenS4{F0: 4294967295, F1: 32767, F2: 973}
            if false {
                for fi28 := 0; fi28 < 3; fi28++ {
                    v26 = v26
                    _ = fi28
                }
                delete(v1, "a")
            } else {
                fmt.Println(v2, v26)
                var v29 float64 = float64(0.1)
                _ = v29
            }
            _ = v27
        } else {
            if true {
                if (true && ((779 > 0) && (!(v2 != v2)))) {
                    fmt.Println(v2, v26)
                    if (!((42 == -1) && (false && (v26 == "foo bar")))) {
                        var v30 int16 = 42
                        _ = v30
                    }
                }
                var v31 uint = 660
                _ = v31
            }
        }
    } else {
        var v32 uint8 = 255
        var v33 string = v26
        var v34 bool = true
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 int16 = 0
    v36 := GenS5{F0: v35, F1: (^(v2)), F2: 65535}
    v35 = (-(v35))
    switch ((((v26 + (v26 + ((v26 + v26) + (v26 + v26)))) + (v26 + "")) + v26) + "a") {
    case "foo bar":
        v2 -= 2147483647
    case "a":
        var v37 int32 = 0
        _ = v37
    }
    fmt.Println(v35)
    var v38 uint = 255
    var v39 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v39
}
