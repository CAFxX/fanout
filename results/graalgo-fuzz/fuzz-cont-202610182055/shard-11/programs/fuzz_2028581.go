package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int8; F2 int16 }
func genFn0() int {
    var v7 string = ("foo bar" + (("a" + (("" + "a") + ("foo bar" + (((("a" + ("hello" + "foo bar")) + "hello") + ("foo bar" + (("hello" + "a") + ""))) + "foo bar")))) + (("a" + ("foo bar" + "\n\t")) + ("" + "hello"))))
    v8 := GenS1{F0: 1, F1: 2147483647, F2: 1}
    var v9 float32 = -2.25
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint8, p1 bool) int {
    p0 -= 0
    if p1 {
        v10 := []int{}
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint8 = 128
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    v13 := []int{128, -1}
    var v14 bool = (272 < 472)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -288
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]uint32)
    var v3 uint8 = 0
    var v4 uint32 = 42
    var v5 int16 = -1
    v6 := make(map[string]uint32)
    v15 := GenS1{F0: 130, F1: (v4 - 65535), F2: 9223372036854775807}
    if (1 >= 42) {
        var v16 int8 = 1
        if false {
            v4 = (v4 - v4)
        } else {
            v16 -= -19
        }
        _ = v16
    }
    var v17 uint = 18446744073709551615
    if (127 == 667) {
        v5 += (v5 ^ 42)
        if (418 != uint16(v17)) {
            fmt.Println(v17, v3, v17)
            v4 = (^(v4))
            var v18 uint32 = 2147483648
            _ = v18
        }
        if ("\n\t" == "") {
            v19 := GenS2{F0: v17, F1: true}
            _ = v19
        } else {
            v20 := GenS3{F0: 1}
            fmt.Println(v3, v17, v4)
            _ = v20
        }
    }
    if true {
        fmt.Println(v4, v3, v3)
        v3 += (v3 - 99)
        var v21 uint64 = 2
        _ = v21
    } else {
        v22 := GenS4{F0: (v5 & (v5 - (v5 ^ 799))), F1: -128, F2: v5}
        delete(v6, "a")
        var v23 int32 = -330
        _ = v22
        _ = v23
    }
    var v24 float64 = (((3.14159 - (1e10 - (float64(-2.25) + 1e10))) - 0.0) * 1e10)
    v25 := v2[32767]
    if true {
        var v26 float64 = 3.14159
        _ = v26
    } else {
        v6[("foo bar" + "hello")] = (v4 | v4)
        if ("\n\t" == "a") {
            if false {
                v27 := GenS2{F0: v17, F1: false}
                var v28 int = 9223372036854775807
                var v29 uint32 = v4
                _ = v27
                _ = v28
                _ = v29
            }
        }
        fmt.Println(v24)
    }
    if (!false) {
        switch ("hello" + "hello") {
        case "foo bar":
            v5 = 1000
        case "\n\t":
            v24 = (0.1 * -2.25)
        case "hello":
            var v30 uint8 = (v3 + v3)
            _ = v30
        }
    }
    var v31 int = int(v5)
    switch (v31 | v31) {
    case 1:
        if ((3.14159 / v24) <= (v24 / v24)) {
            if ((32767 >= 806) || false) {
                fmt.Println(v3)
                var v32 bool = true
                _ = v32
            }
        } else {
            if true {
                v5 += v5
                v33 := v6["a"]
                for fi34 := 0; fi34 < 8; fi34++ {
                    switch genFn1(127, true) {
                    case 5:
                        var v35 string = "foo bar"
                        _ = v35
                    case -699:
                        delete(v2, genFn0())
                    case 0:
                        v36 := v2[int(v17)]
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 8; fi37++ {
                        var v38 int = genFn0()
                        v38 = fi37
                        _ = fi37
                        _ = v38
                    }
                    _ = fi34
                }
                _ = v33
            } else {
                if false {
                    var v39 uint = v17
                    _ = v39
                }
            }
            if true {
                switch "" {
                case "":
                    if true {
                        fmt.Println(v25)
                        var v40 string = "hello"
                        var v41 float64 = 1.5
                        _ = v40
                        _ = v41
                    }
                default:
                    v5 = v5
                }
            }
            _ = len(v6)
        }
    case 42:
        for fi42 := 0; fi42 < 1; fi42++ {
            v24 = v24
            _ = len(v1)
            _ = fi42
        }
    default:
        v24 -= (3.14159 + v24)
    }
    var v43 int64 = 9223372036854775807
    if (uint8(v17) != (^(v3))) {
        fmt.Println(v43)
        v5 = -673
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
