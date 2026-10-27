package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 string; F2 bool }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int32; F1 string }
func genFn0(p0 uint32) int16 {
    v7 := GenS1{F0: 1, F1: 100, F2: -547}
    v8 := []int{-62, -1, 32768, int(p0)}
    p0 -= (p0 ^ 2)
    var v9 bool = (((true && (((3.14159 - 1e10) + -2.25) >= ((1.5 - 3.14159) + 0.1))) && ((!(((p0 | 71) > p0) && (((42 == uint16(p0)) || true) && true))) && (-1 == 42))) || (4294967295 <= 0))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -32768
}

func main() {
    var v1 int8 = -128
    v2 := []int{127, 32767}
    var v3 int = -1
    var v4 uint64 = 1
    var v5 uint16 = 1
    var v6 uint16 = v5
    var v10 uint16 = (v5 << ((22) & 63))
    for fi11 := 0; fi11 < 18; fi11++ {
        v5 = (v5 / ((1) | 1))
        v2 = append(v2, v3)
        fmt.Println(v4)
        _ = fi11
    }
    var v12 uint8 = 1
    v5 = v6
    v13 := GenS2{F0: "a"}
    switch 9223372036854775807 {
    case 32768:
        var v14 uint32 = 249
        _ = v14
    default:
        fmt.Println(v5)
    }
    switch v3 {
    case -78:
        var v15 uint8 = (v12 + v12)
        _ = v15
    }
    v10 -= v10
    var v16 uint16 = (v6 + 561)
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 float64 = 3.14159
        v3 = (fi17 & v3)
        if (!false) {
            if true {
                v2 = append(v2, (^(v3)))
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(v4, fi17)
                    v10 -= 42
                    v18 -= 3.14159
                    _ = fi19
                }
                var v20 int8 = v1
                _ = v20
            } else {
                v10 = (^(v6))
                fmt.Println(v10, v4, v18)
            }
            fmt.Println(v5, v6)
            if false {
                v16 += 65535
                var v21 float32 = -2.25
                _ = v21
            } else {
                v10 -= (v6 % ((v10) | 1))
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    var v22 float64 = v18
                    _ = v22
                case "hello":
                    v23 := GenS2{F0: "foo bar"}
                    _ = v23
                case "":
                    v1 -= (v1 ^ v1)
                default:
                    var v24 float64 = v18
                    _ = v24
                }
                if true {
                    var v25 float32 = 3.14159
                    _ = v25
                }
            }
        } else {
            switch "foo bar" {
            case "":
                v26 := GenS3{F0: 4294967295, F1: ("" + ("a" + (("a" + "hello") + "hello"))), F2: true}
                _ = v26
            case "a":
                var v27 uint16 = 255
                _ = v27
            case "foo bar":
                v6 -= v6
            }
            if (1 <= 9223372036854775807) {
                for fi28 := 0; fi28 < 20; fi28++ {
                    var v29 bool = true
                    _ = fi28
                    _ = v29
                }
            }
            v3 -= (^(fi17))
        }
        _ = fi17
        _ = v18
    }
    var v30 int32 = -2147483648
    var v31 bool = true
    v32 := GenS4{F0: 0, F1: v10}
    if len(v2) > 0 {
        v2[0] = (^(v3))
    }
    v33 := GenS3{F0: 4294967295, F1: "a", F2: v31}
    if v31 {
        switch v3 {
        case 1:
            v34 := GenS5{F0: v3, F1: 65535}
            _ = v34
        case 4294967295:
            var v35 float32 = (((1.5 - 3.14159) / 0.1) * 0.0)
            _ = v35
        case -358:
            for fi36 := 0; fi36 < 8; fi36++ {
                v37 := GenS3{F0: 576, F1: "a", F2: ((-9223372036854775808 != 9223372036854775807) || false)}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi36
                _ = v37
            }
        }
        v38 := GenS6{F0: (v30 ^ v30), F1: ("foo bar" + "hello")}
        for fi39 := 0; fi39 < 12; fi39++ {
            v40 := GenS6{F0: v30, F1: ((("" + "\n\t") + ("\n\t" + "a")) + "")}
            _ = fi39
            _ = v40
        }
        _ = v38
    } else {
        var v41 bool = (v12 >= v12)
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
