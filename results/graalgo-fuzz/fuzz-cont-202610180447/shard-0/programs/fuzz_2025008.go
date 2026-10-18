package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 bool; F1 string; F2 uint16 }
func genFn0(p0 uint64) int8 {
    switch "foo bar" {
    case "foo bar":
        var v7 int32 = 0
        _ = v7
    case "a":
        v8 := []int{391, -2147483648, 9223372036854775807}
        _ = v8
    default:
        v9 := []int{}
        _ = v9
    }
    v10 := []int{32767, -965, 127}
    if (true && true) {
        v11 := GenS1{F0: "foo bar"}
        var v12 uint64 = (p0 << ((65535) & 63))
        for fi13 := 0; fi13 < 5; fi13++ {
            p0 -= (p0 >> ((42) & 63))
            if (false && (2 <= 65535)) {
                v14 := GenS1{F0: (("hello" + "\n\t") + ("a" + "foo bar"))}
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 string = (("foo bar" + (("a" + "foo bar") + (("foo bar" + "") + ("a" + "")))) + (("\n\t" + "\n\t") + ("" + "foo bar")))
                    p0 = (v12 - (p0 << ((339) & 63)))
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                v12 += 178
                switch "hello" {
                case "hello":
                    switch fi13 {
                    case -997:
                        if (255 >= 1) {
                            var v17 int16 = 42
                            if (397 < 4294967295) {
                                var v18 float32 = 1e10
                                p0 -= (v12 << ((0) & 63))
                                _ = v18
                            }
                            v19 := GenS1{F0: "hello"}
                            _ = v17
                            _ = v19
                        }
                    }
                case "\n\t":
                    var v20 int = -2147483648
                    _ = v20
                default:
                    var v21 uint16 = 0
                    _ = v21
                }
            }
            fmt.Println(fi13)
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v10
    return -128
}
func genFn1(p0 int16) int16 {
    var v22 uint = 65535
    p0 -= -32768
    for fi23 := 0; fi23 < 14; fi23++ {
        if true {
            v24 := GenS1{F0: ("foo bar" + "\n\t")}
            switch "a" {
            case "hello":
                v25 := GenS1{F0: "a"}
                _ = v25
            default:
                v26 := []int{int(p0), fi23}
                _ = v26
            }
            _ = v24
        } else {
            var v27 uint32 = 304
            var v28 int16 = 42
            _ = v27
            _ = v28
        }
        _ = fi23
    }
    if true {
        if false {
            var v29 int = 578
            if (true || true) {
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS1{F0: ""}
                    _ = fi30
                    _ = v31
                }
                var v32 uint16 = 2
                _ = v32
            } else {
                fmt.Println(v22, v22, p0)
                var v33 int16 = 1000
                switch (("" + "a") + ("" + "")) {
                case "\n\t":
                    fmt.Println(v33)
                default:
                    v22 -= 31
                }
                _ = v33
            }
            _ = v29
        } else {
            v34 := GenS2{F0: ("a" > ("\n\t" + "foo bar")), F1: 0}
            _ = v34
        }
        v35 := []int{}
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        switch fi36 {
        case 127:
            var v37 int64 = 1
            _ = v37
        default:
            v38 := []int{fi36, 0, 128}
            _ = v38
        }
        switch (fi36 % (((fi36 & int(v22))) | 1)) {
        case 32767:
            var v39 float64 = -2.25
            _ = v39
        case 0:
            fmt.Println(p0)
        case 4294967295:
            if true {
                fmt.Println(v22, v22, v22)
                v40 := GenS2{F0: ("\n\t" <= (("" + "\n\t") + "foo bar")), F1: 42}
                _ = v40
            } else {
                p0 -= (p0 - (p0 >> ((0) & 63)))
            }
        }
        v41 := GenS2{F0: (0 >= 0), F1: 0}
        _ = fi36
        _ = v41
    }
    _ = p0
    _ = v22
    return -407
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (-2.25 - 1e10)
    var v3 uint8 = 1
    v4 := make(map[int]int)
    var v5 float64 = 1e10
    v6 := make(map[int]bool)
    fmt.Println(v3)
    v2 = (((0.0 * (float32(3.14159) - 1e10)) + 1.5) + (v2 * float32(3.14159)))
    v42 := []int{128, 9223372036854775807}
    v43 := GenS3{F0: 127, F1: v3, F2: 2}
    v44 := GenS4{F0: true, F1: "hello", F2: 599}
    var v45 uint = 0
    switch 36 {
    case 110:
        fmt.Println(v3)
    case -9223372036854775808:
        v2 = float32(0.0)
    case 128:
        switch "a" {
        case "foo bar":
            if len(v42) >= 2 {
                v42 = v42[1:len(v42)]
            }
        }
    }
    var v46 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
