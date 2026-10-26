package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string; F1 int64 }
func genFn0(p0 uint64, p1 string, p2 bool) int {
    v4 := []int{4294967295, -404}
    var v5 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -79
}

func main() {
    var v1 uint64 = 935
    var v2 int = 42
    var v3 int64 = 9223372036854775807
    v3 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{(-(v2)), v2, fi6, fi6}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 0, F1: 65535}
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 42
    var v11 string = "a"
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    switch v2 {
    case 9223372036854775807:
        fmt.Println(v3, v1, v1)
    case -314:
        if (!(true && (v2 < (v2 | (v2 >> ((912) & 63)))))) {
            for fi14 := 0; fi14 < 4; fi14++ {
                switch v2 {
                case 380:
                    v15 := GenS2{F0: -9223372036854775808}
                    _ = v15
                }
                v16 := []int{(-(fi14)), fi14, (fi14 >> ((1) & 63))}
                var v17 int64 = v3
                _ = fi14
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint16 = 449
                v1 -= v1
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 float32 = (1.5 / (((0.0 / 1e10) * -2.25) - ((-2.25 - ((1e10 / (((1.5 / (1.5 + 1.5)) - 1.5) / float32((1.5 * -2.25)))) / -2.25)) * (1e10 - 1.5))))
                    var v22 uint = 4294967295
                    v2 += (fi20 | int(v22))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
            v11 = v11
        } else {
            if ((((true || false) || ((float32(((3.14159 + -2.25) + 1e10)) >= 0.0) || (0.1 >= 0.0))) || true) && true) {
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 bool = false
                    _ = fi23
                    _ = v24
                }
            }
            fmt.Println(v3)
        }
    }
    var v25 int16 = 782
    fmt.Println(v3, v25)
    v26 := []int{v2, -689}
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(fi27, v25)
        for fi28 := 0; fi28 < 3; fi28++ {
            switch "hello" {
            case "a":
                var v29 string = "\n\t"
                _ = v29
            case "\n\t":
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
            case "":
                v1 = (v1 % (((v1 << ((18446744073709551615) & 63))) | 1))
            default:
                v1 -= uint64(v10)
            }
            _ = fi28
        }
        var v30 string = "foo bar"
        _ = fi27
        _ = v30
    }
    fmt.Println(v1, v10, v2)
    v31 := GenS1{F0: v10, F1: 1}
    switch v11 {
    case "a":
        for fi32 := 0; fi32 < 14; fi32++ {
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 int64 = v3
                var v35 string = "\n\t"
                if false {
                    var v36 float32 = 1e10
                    _ = v36
                } else {
                    fmt.Println(fi33, v35)
                    v2 += (v2 % ((fi33) | 1))
                    var v37 uint = 0
                    if (true || false) {
                        v25 = int16(v1)
                    } else {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        if (!((42 <= 1) || false)) {
                            v34 = (v3 | (v3 >> (((v37 * (^(v37)))) & 63)))
                            var v38 int16 = v25
                            _ = v38
                        }
                    }
                    v39 := GenS3{F0: 1, F1: (v37 + v37)}
                    v26 = append(v26, (-(fi33)))
                    switch ("hello" + ((v35 + (v11 + "foo bar")) + v11)) {
                    case "a":
                        v40 := GenS4{F0: ((v35 + v11) + ((((v35 + v11) + v11) + v11) + ("" + (v11 + (v11 + ""))))), F1: (v3 / ((v3) | 1))}
                        _ = v40
                    case "foo bar":
                        var v41 string = ((v11 + v11) + ((((v35 + v35) + (v11 + "\n\t")) + v35) + v11))
                        _ = v41
                    default:
                        var v42 string = (((v35 + "") + ("foo bar" + "")) + "a")
                        _ = v42
                    }
                    _ = v37
                    _ = v39
                }
                _ = fi33
                _ = v34
                _ = v35
            }
            for fi43 := 0; fi43 < 3; fi43++ {
                v10 = 4294967295
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = fi43
            }
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            _ = fi32
        }
    case "":
        var v44 int64 = int64(v2)
        _ = v44
    case "foo bar":
        var v45 uint8 = 127
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
}
