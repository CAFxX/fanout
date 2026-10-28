package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
func genFn0(p0 int32, p1 int16) bool {
    v7 := GenS1{F0: true, F1: ("" + "")}
    if ("foo bar" == "") {
        v8 := GenS1{F0: ((p1 / ((466) | 1)) != 32767), F1: "a"}
        _ = v8
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return ((0.0 * 3.14159) >= (3.14159 + 0.0))
}
func genFn1() int32 {
    switch -9223372036854775808 {
    case -2147483648:
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := []int{(fi9 << ((0) & 63)), (fi9 - fi9), fi9}
            _ = fi9
            _ = v10
        }
    case 4294967295:
        switch "\n\t" {
        case "foo bar":
            for fi11 := 0; fi11 < 7; fi11++ {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := []int{fi11}
                    v14 := GenS1{F0: (false || (((128 == 0) && false) && (-571 != 1))), F1: ("a" + "foo bar")}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    if ((fi11 - (fi11 ^ fi11)) > -2147483648) {
                        v16 := []int{fi15, (^(fi11))}
                        _ = v16
                    }
                    var v17 int8 = -1
                    _ = fi15
                    _ = v17
                }
                switch "foo bar" {
                case "foo bar":
                    var v18 int16 = 174
                    _ = v18
                case "\n\t":
                    v19 := []int{1, 9223372036854775807, fi11}
                    _ = v19
                }
                _ = fi11
            }
        case "hello":
            var v20 int32 = 0
            _ = v20
        default:
            switch 648 {
            case 4294967295:
                v21 := GenS1{F0: true, F1: (("a" + (("hello" + ("" + "")) + "a")) + "hello")}
                _ = v21
            case 2147483647:
                var v22 float32 = 3.14159
                _ = v22
            case 128:
                var v23 int16 = 42
                _ = v23
            default:
                var v24 uint32 = 65535
                _ = v24
            }
        }
    case 128:
        var v25 int = 670
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint = 1
            fmt.Println(v28, v28, v28)
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    return 2147483647
}

func main() {
    var v1 uint8 = 110
    var v2 int32 = -1
    var v3 uint = 4294967295
    v4 := make(map[uint32]int)
    var v5 bool = true
    var v6 int = 32767
    var v29 float64 = -2.25
    v29 = v29
    v30 := v4[uint32(v1)]
    _ = len(v4)
    switch ("hello" + ("a" + "hello")) {
    case "":
        v6 -= (v30 % ((v6) | 1))
    case "\n\t":
        if v5 {
            var v31 int8 = 1
            _ = v31
        } else {
            v32 := []int{(v30 & (-(v30))), v6, -9223372036854775808, 128}
            switch "" {
            case "a":
                _ = len(v32)
            default:
                if true {
                    v2 = 0
                    var v33 bool = v5
                    _ = v33
                }
            }
            fmt.Println(v2)
            _ = v32
        }
    default:
        var v34 float64 = v29
        _ = v34
    }
    if true {
        v3 += uint(v1)
        v35 := []int{}
        var v36 float64 = 0.1
        _ = v35
        _ = v36
    }
    v37 := GenS2{F0: 862, F1: 65535, F2: -1}
    var v38 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
