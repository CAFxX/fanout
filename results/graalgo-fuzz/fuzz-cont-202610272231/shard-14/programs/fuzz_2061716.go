package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16) bool {
    p0 -= (p0 + (p0 * p0))
    v7 := []int{42}
    if len(v7) > 0 {
        v7[0] = -354
    }
    if true {
        switch "hello" {
        case "\n\t":
            var v8 int16 = 1
            _ = v8
        case "foo bar":
            fmt.Println(p0)
        default:
            if true {
                v9 := GenS1{F0: 88, F1: 4294967295}
                _ = v9
            }
        }
        var v10 int = -196
        switch (v10 * (v10 & (-(v10)))) {
        case -1:
            for fi11 := 0; fi11 < 12; fi11++ {
                v10 -= 42
                p0 -= 105
                var v12 uint8 = 127
                _ = fi11
                _ = v12
            }
        case 2147483647:
            fmt.Println(p0)
        }
        _ = v10
    }
    p0 = (p0 & (p0 & (p0 * p0)))
    _ = p0
    _ = v7
    return false
}
func genFn1(p0 uint32, p1 uint8) int {
    if false {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 string = ""
            var v15 string = v14
            if (!false) {
                v16 := GenS1{F0: 73, F1: 234}
                _ = v16
            } else {
                var v17 uint16 = uint16(p1)
                v15 = "foo bar"
                _ = v17
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        p0 = 651
    }
    v18 := GenS1{F0: -87, F1: 572}
    switch -414 {
    case -892:
        if false {
            fmt.Println(p1, p1)
            for fi19 := 0; fi19 < 6; fi19++ {
                switch "a" {
                case "foo bar":
                    for fi20 := 0; fi20 < 8; fi20++ {
                        p0 = p0
                        _ = fi20
                    }
                case "a":
                    var v21 uint8 = 97
                    _ = v21
                case "hello":
                    v22 := []int{}
                    _ = v22
                }
                p0 = p0
                _ = fi19
            }
        }
    case 9223372036854775807:
        var v23 int32 = 42
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v18
    return -762
}

func main() {
    var v1 int8 = -128
    var v2 bool = (false && true)
    v3 := make(map[string]string)
    v4 := []int{0, 244, -1, 9223372036854775807}
    v5 := []string{"", ""}
    var v6 string = ""
    if false {
        if v2 {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 int64 = -528
                fmt.Println(v6, v6, v6)
                _ = fi24
                _ = v25
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v6 = ("hello" + v6)
            switch 751 {
            case 32768:
                v3["\n\t"] = "hello"
            case 42:
                v26 := GenS1{F0: (v1 % ((12) | 1)), F1: 505}
                _ = v26
            case 1:
                var v27 int32 = -2147483648
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 uint = 374
            v1 -= 42
            _ = fi28
            _ = v29
        }
        var v30 float64 = (1.5 * (1e10 * (-2.25 + (0.0 * 0.1))))
        _ = v30
    }
    var v31 int32 = 42
    v32 := v3[(((v6 + ("hello" + v6)) + v6) + ("\n\t" + "hello"))]
    for fi33 := 0; fi33 < 4; fi33++ {
        v32 = v32
        v2 = v2
        v1 -= v1
        _ = fi33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        v6 = (v6 + (v32 + ((("" + ("\n\t" + ((v32 + (v32 + v32)) + "a"))) + ("a" + (v32 + "a"))) + (("" + v32) + (v6 + v32)))))
        _ = fi34
    }
    v32 = "foo bar"
    v35 := v3[(v32 + "")]
    var v36 uint32 = 4294967295
    switch 32768 {
    case 128:
        var v37 float64 = float64((1e10 / (-2.25 / float32(((0.0 / 0.1) - (((((0.0 + (((1.5 - 0.0) * float64(-2.25)) - 1.5)) / ((((1.5 - 1e10) * 0.1) * (3.14159 - ((0.0 * 1.5) / 1e10))) / 1.5)) / float64(((float32(((0.1 / float64(3.14159)) + float64(0.1))) + 0.1) / 1e10))) + float64(1e10)) * 1e10))))))
        _ = v37
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 uint64 = 260
        _ = fi38
        _ = v39
    }
    v40 := GenS1{F0: -128, F1: 0}
    if len(v5) > 0 {
        v5[0] = ("foo bar" + v35)
    }
    var v41 string = v6
    v42 := GenS2{F0: 311}
    var v43 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
