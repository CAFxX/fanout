package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0() int8 {
    var v4 uint32 = 1
    if (false || ((!false) && (0 != 42))) {
        switch 127 {
        case 4294967295:
            fmt.Println(v4)
        case 181:
            v4 = (v4 >> ((42) & 63))
        }
    } else {
        if (!false) {
            if true {
                var v5 int32 = 0
                v6 := []int{-2147483648, 4294967295, 42}
                _ = v5
                _ = v6
            } else {
                fmt.Println(v4)
                v4 -= v4
                v4 = (^(v4))
            }
            v4 = v4
            if (-498 == 0) {
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := []int{(fi7 + (-(fi7)))}
                    fmt.Println(v4, fi7)
                    _ = fi7
                    _ = v8
                }
            }
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int8 = -128
            _ = fi9
            _ = v10
        }
        var v11 int64 = -1
        _ = v11
    }
    _ = v4
    return 102
}
func genFn1(p0 int, p1 uint32) int16 {
    p1 -= 2147483647
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: (p1 % ((p1) | 1)), F1: 0, F2: 128}
            v15 := GenS1{F0: p1, F1: 9223372036854775807, F2: 85}
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 float64 = (3.14159 / ((0.1 / 1.5) - 0.0))
        _ = fi12
        _ = v16
    }
    var v17 int8 = 0
    if (42 != 65535) {
        var v18 uint = 989
        var v19 float64 = 1.5
        _ = v18
        _ = v19
    }
    var v20 bool = false
    _ = p0
    _ = p1
    _ = v17
    _ = v20
    return 1
}

func main() {
    v1 := []string{}
    var v2 string = ("\n\t" + "\n\t")
    v3 := []int{636, 685, -419}
    fmt.Println(v2)
    switch "hello" {
    case "hello":
        for fi21 := 0; fi21 < 17; fi21++ {
            v2 = "\n\t"
            var v22 uint64 = 18446744073709551615
            _ = fi21
            _ = v22
        }
    case "a":
        v23 := GenS1{F0: 284, F1: -1, F2: 128}
        _ = v23
    }
    v2 = v2
    if ((genFn0() != genFn0()) && (32767 >= 1)) {
        if (!true) {
            if false {
                var v24 int16 = 42
                _ = v24
            } else {
                v2 = v2
                fmt.Println(v2)
            }
            v2 = v2
            if (954 == 0) {
                switch "" {
                case "":
                    var v25 int = 4294967295
                    _ = v25
                case "hello":
                    var v26 bool = false
                    _ = v26
                case "foo bar":
                    var v27 int8 = 95
                    _ = v27
                }
            }
        } else {
            v2 = v2
            var v28 uint16 = 948
            _ = v28
        }
    }
    v2 = v2
    v2 = (v2 + "hello")
    v29 := GenS2{F0: 1}
    fmt.Println(v2)
    var v30 uint64 = 42
    var v31 string = v2
    v31 = (("" + (v31 + ("a" + "hello"))) + v31)
    v31 = "a"
    switch v2 {
    case "a":
        switch 37 {
        case 32768:
            if ((((0.0 * -2.25) / 1.0) * (3.14159 + float32(1e10))) <= 3.14159) {
                _ = len(v1)
            } else {
                v30 -= 368
            }
        }
    case "hello":
        for fi32 := 0; fi32 < 17; fi32++ {
            v33 := GenS1{F0: 0, F1: -851, F2: 255}
            _ = len(v1)
            _ = fi32
            _ = v33
        }
    }
    if (-2147483648 >= 506) {
        if (false && false) {
            if (0.0 <= 1.5) {
                for fi34 := 0; fi34 < 1; fi34++ {
                    var v35 uint8 = 42
                    _ = len(v3)
                    v2 = v31
                    _ = fi34
                    _ = v35
                }
                v30 = v30
                var v36 float64 = float64(0.1)
                _ = v36
            } else {
                fmt.Println(v31)
            }
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 int32 = -1
            fmt.Println(v30, v31, v2)
            switch v2 {
            case "foo bar":
                v30 = v30
            case "\n\t":
                fmt.Println(v2, v30)
            case "a":
                var v39 int8 = 114
                _ = v39
            }
            _ = fi37
            _ = v38
        }
    }
    var v40 int8 = 100
    v41 := GenS2{F0: 128}
    var v42 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v40
    _ = v41
    _ = v42
}
