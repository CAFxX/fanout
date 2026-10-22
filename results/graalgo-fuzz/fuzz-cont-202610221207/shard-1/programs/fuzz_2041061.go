package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int64; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint16 }
func genFn0() uint32 {
    switch ("foo bar" + ("hello" + ("a" + "a"))) {
    case "foo bar":
        v4 := GenS1{F0: -498}
        _ = v4
    case "a":
        switch 4294967295 {
        case -2147483648:
            switch "a" {
            case "a":
                if false {
                    var v5 float64 = (1e10 / 1.5)
                    v5 += (1e10 - 0.0)
                    v6 := GenS1{F0: -1}
                    _ = v5
                    _ = v6
                }
            }
        default:
            var v7 float32 = 0.0
            _ = v7
        }
    }
    var v8 uint32 = 2147483647
    _ = v8
    return v8
}
func genFn1(p0 int16) int64 {
    v9 := GenS1{F0: 589}
    var v10 uint16 = 256
    switch -2147483648 {
    case 115:
        switch 4294967295 {
        case 2147483647:
            switch -789 {
            case 4294967295:
                var v11 float64 = (1.5 / 1.0)
                _ = v11
            case 2147483647:
                for fi12 := 0; fi12 < 16; fi12++ {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        var v14 uint8 = 0
                        v15 := GenS1{F0: (^(fi12))}
                        v10 = v10
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = fi12
                }
            case 1:
                fmt.Println(p0, p0)
            default:
                for fi16 := 0; fi16 < 2; fi16++ {
                    if true {
                        var v17 uint64 = 42
                        v17 += v17
                        var v18 uint = 2
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(fi16)
                    _ = fi16
                }
            }
        }
    case 9223372036854775807:
        fmt.Println(p0, p0)
    case -9223372036854775808:
        v10 -= 32767
    default:
        v19 := []int{127, 42, int(v10), -9223372036854775808}
        _ = v19
    }
    fmt.Println(v10, p0)
    if false {
        if false {
            v20 := GenS1{F0: 128}
            _ = v20
        }
        v21 := GenS1{F0: 169}
        _ = v21
    }
    _ = p0
    _ = v10
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 2
    v3 := []uint32{}
    if true {
        var v22 float64 = 3.14159
        var v23 int64 = 1
        _ = v22
        _ = v23
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if ((0.0 / 3.14159) == float64(0.1)) {
            v2 = (v2 ^ v2)
            v2 = 128
        } else {
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint8 = v2
                v25 = (v2 - (v2 + v25))
                _ = fi24
                _ = v25
            }
            v2 = v2
            if ((!true) || (402 < 42)) {
                fmt.Println(v1, v2)
            } else {
                var v26 float64 = 0.0
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(v1)
            _ = fi28
        }
        var v29 int = fi27
        _ = fi27
        _ = v29
    }
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    var v30 uint16 = 255
    switch "a" {
    case "a":
        var v31 bool = (!false)
        _ = v31
    }
    fmt.Println(v30, v2, v30)
    var v32 int = 2147483647
    switch "" {
    case "foo bar":
        v1 -= (v1 + 42)
    default:
        var v33 uint16 = (^(v30))
        _ = v33
    }
    v32 += (v32 | (v32 << ((v1) & 63)))
    v34 := GenS1{F0: (^(v32))}
    if (true && ((65535 <= 4294967295) && true)) {
        switch v32 {
        case -9223372036854775808:
            if ((uint32(v2) > 1) && false) {
                var v35 uint64 = 928
                fmt.Println(v30, v35, v32)
                v2 -= v2
                _ = v35
            }
        case 813:
            if true {
                for fi36 := 0; fi36 < 18; fi36++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi36
                }
            } else {
                v37 := GenS2{F0: v32}
                if len(v3) > 0 {
                    v3[0] = uint32(v2)
                }
                fmt.Println(v32, v1)
                _ = v37
            }
        }
        if len(v3) > 0 {
            v3[0] = 0
        }
        fmt.Println(v2, v32, v30)
    } else {
        if (1 == 18446744073709551615) {
            if (3.14159 != (1e10 - 3.14159)) {
                v32 = v32
                fmt.Println(v1)
            }
            _ = len(v3)
            v38 := GenS1{F0: (-(v32))}
            _ = v38
        } else {
            v32 = 4294967295
        }
        var v39 float32 = -2.25
        if (v32 >= (v32 % ((v32) | 1))) {
            v3 = append(v3, 2147483647)
            var v40 uint8 = v2
            var v41 uint64 = 1
            _ = v40
            _ = v41
        }
        _ = v39
    }
    fmt.Println(v1, v32, v32)
    var v42 uint16 = v30
    v43 := GenS3{F0: (^(v1)), F1: -622, F2: (v2 - uint8(v30))}
    v44 := GenS4{F0: 69, F1: v30}
    var v45 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
