package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int }
func genFn0(p0 int) int {
    var v6 uint32 = 653
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint64 = 928
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int8 = 42
            v11 := []int{2147483647, fi7}
            if len(v11) > 0 {
                v11[0] = (fi9 << ((18446744073709551615) & 63))
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: ("hello" + ("" + ""))}
        _ = fi7
        _ = v12
        _ = v8
    }
    if (1 >= 0) {
        v6 = 89
        v13 := []int{-9223372036854775808, (p0 | (p0 - (p0 + p0))), 32767}
        _ = v13
    } else {
        v6 = (v6 | 4294967295)
        if (3.14159 == 3.14159) {
            if true {
                v14 := []int{}
                if false {
                    v15 := GenS1{F0: ("a" + "foo bar")}
                    _ = v15
                } else {
                    p0 += (p0 | (p0 | (p0 << ((245) & 63))))
                    v16 := GenS1{F0: "foo bar"}
                    _ = v16
                }
                _ = v14
            } else {
                fmt.Println(p0)
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := []int{}
                    v6 = (v6 / ((v6) | 1))
                    var v19 int8 = 0
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := []int{p0, (p0 >> ((284) & 63))}
                _ = v20
            }
            switch p0 {
            case -541:
                if true {
                    var v21 uint16 = 255
                    v22 := []int{}
                    var v23 int = (p0 * (p0 / (((p0 + p0)) | 1)))
                    _ = v21
                    _ = v22
                    _ = v23
                }
            case 128:
                var v24 float64 = (1.5 - float64((3.14159 - 1.5)))
                _ = v24
            case -2147483648:
                p0 += p0
            default:
                var v25 int8 = 100
                _ = v25
            }
        } else {
            fmt.Println(p0, v6)
        }
    }
    var v26 float64 = 1e10
    _ = p0
    _ = v26
    _ = v6
    return p0
}

func main() {
    var v1 int = 127
    var v2 int16 = 1000
    var v3 bool = true
    var v4 int64 = -1
    v5 := []string{("" + "foo bar"), "a", "a", "\n\t"}
    switch "hello" {
    case "":
        var v27 uint8 = 128
        _ = v27
    default:
        for fi28 := 0; fi28 < 12; fi28++ {
            v2 = (v2 - 32767)
            var v29 uint16 = 255
            _ = fi28
            _ = v29
        }
    }
    fmt.Println(v4, v3)
    var v30 int64 = v4
    v31 := GenS2{F0: 143, F1: genFn0((v1 + v1)), F2: v3}
    var v32 string = "hello"
    for fi33 := 0; fi33 < 18; fi33++ {
        fmt.Println(v32, v2, v2)
        if (0.0 >= (0.1 * 1e10)) {
            for fi34 := 0; fi34 < 4; fi34++ {
                v35 := GenS2{F0: 254, F1: genFn0((-(fi34))), F2: false}
                v36 := GenS1{F0: "a"}
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = fi33
    }
    fmt.Println(v4, v1, v2)
    var v37 int16 = 1000
    var v38 int = (^(v1))
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := GenS3{F0: (-(v1))}
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        var v42 int = fi41
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        for fi44 := 0; fi44 < 5; fi44++ {
            v3 = (!((v2 + -1) != 1000))
            _ = fi44
        }
        switch 3 {
        case -218:
            var v45 int64 = (v30 % (((v4 >> ((18446744073709551615) & 63))) | 1))
            _ = v45
        case 1:
            switch ("a" + v32) {
            case "hello":
                fmt.Println(v4, v38, v38)
            }
        }
        _ = fi43
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        if v3 {
            v47 := GenS1{F0: "hello"}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v47
        } else {
            v1 = -9223372036854775808
            v48 := GenS4{F0: 2147483647, F1: fi46}
            if (2147483647 == 42) {
                fmt.Println(v32, v32)
                var v49 int32 = 42
                _ = v49
            }
            _ = v48
        }
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
