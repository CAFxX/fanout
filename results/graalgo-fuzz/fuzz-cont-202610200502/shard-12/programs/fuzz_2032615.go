package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0(p0 bool) string {
    v6 := GenS1{F0: (p0 || (!true)), F1: 4294967295}
    p0 = (!true)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint64 = 0
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 = true
            _ = fi9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if false {
                var v11 uint32 = 42
                var v12 int16 = 32767
                _ = v11
                _ = v12
            } else {
                fmt.Println(p0)
                v13 := []int{42, 42}
                fmt.Println(fi10)
                _ = v13
            }
            if p0 {
                v14 := []int{109}
                _ = v14
            }
            _ = fi10
        }
        _ = fi7
        _ = v8
    }
    v15 := GenS1{F0: false, F1: 18446744073709551615}
    _ = p0
    _ = v15
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -1
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    v5 := make(map[int]int)
    var v16 int = 32768
    for fi17 := 0; fi17 < 5; fi17++ {
        if true {
            v18 := GenS1{F0: (("" + "a") > "a"), F1: 2}
            v19 := GenS1{F0: (((((1e10 * -2.25) - ((-2.25 / 0.1) + (((3.14159 + -2.25) + 3.14159) - -2.25))) + 3.14159) != (0.1 + 0.1)) || (2 != 0)), F1: 42}
            switch "\n\t" {
            case "foo bar":
                if false {
                    v16 += 2147483647
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi17, fi17)
                }
            case "a":
                v20 := GenS1{F0: ((true || true) && (true || true)), F1: 82}
                _ = v20
            case "hello":
                v16 = (v16 + fi17)
            default:
                fmt.Println(v2)
            }
            _ = v18
            _ = v19
        } else {
            if true {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    var v21 bool = (!(885 <= 720))
                    _ = v21
                default:
                    var v22 int = fi17
                    _ = v22
                }
                var v23 uint16 = 42
                _ = v23
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                v25 := GenS2{F0: 42, F1: 1, F2: 0}
                _ = fi24
                _ = v25
            }
        }
        _ = fi17
    }
    v4 = append(v4, 0)
    v2 -= 100
    fmt.Println(v2, v2)
    v26 := v5[v16]
    if ((563 > 2) && true) {
        if (!false) {
            v26 += (^(v26))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        for fi27 := 0; fi27 < 3; fi27++ {
            for fi28 := 0; fi28 < 8; fi28++ {
                v16 = v16
                v29 := GenS1{F0: (!true), F1: 0}
                _ = fi28
                _ = v29
            }
            switch 32767 {
            case 166:
                if (!(false || true)) {
                    if (32767 == 1000) {
                        fmt.Println(v2, fi27)
                        fmt.Println(fi27, v26)
                    } else {
                        var v30 uint16 = 745
                        var v31 uint64 = 656
                        _ = v30
                        _ = v31
                    }
                    delete(v3, 65535)
                    v5[fi27] = (v26 >> ((4294967295) & 63))
                }
            case 852:
                switch "foo bar" {
                case "a":
                    v16 -= v16
                default:
                    _ = len(v4)
                }
            case 1:
                var v32 int = v26
                _ = v32
            default:
                delete(v5, (v16 | -2147483648))
            }
            v16 = v16
            _ = fi27
        }
        var v33 float32 = 0.1
        _ = v33
    }
    v26 += (v26 - v26)
    var v34 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
