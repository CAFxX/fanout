package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 string; F2 int }
func genFn0(p0 int64) string {
    p0 += (p0 % (((p0 << ((255) & 63))) | 1))
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 471}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        if (!true) {
            var v7 float64 = (((((1.5 / 1.0) * 1.5) - 0.0) - 1.5) + (1e10 + 0.0))
            _ = v7
        }
        v8 := []int{32767}
        if ((1 <= -1) && (!true)) {
            var v9 uint32 = 468
            for fi10 := 0; fi10 < 18; fi10++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch ("a" + "a") {
                case "foo bar":
                    var v11 uint8 = 29
                    _ = v11
                default:
                    var v12 int32 = -2147483648
                    fmt.Println(v12, v9, fi10)
                    fmt.Println(p0, v12)
                    _ = v12
                }
                _ = fi10
            }
            _ = v9
        } else {
            p0 = (p0 % (((p0 & (p0 - p0))) | 1))
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 uint = 948
                var v15 int16 = 1
                var v16 float32 = 3.14159
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        _ = fi6
        _ = v8
    }
    if (false && false) {
        var v17 int8 = -100
        _ = v17
    }
    _ = p0
    return ("a" + (("hello" + "foo bar") + ("" + "a")))
}
func genFn1(p0 int32, p1 string, p2 string) uint32 {
    var v18 float64 = float64(-2.25)
    var v19 int16 = 760
    if (65535 < 165) {
        v20 := []int{2147483647, -2147483648}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 837
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 906
    var v3 float32 = 1e10
    var v21 int = -9223372036854775808
    if (1e10 > (1.5 * (0.1 + 3.14159))) {
        var v22 uint64 = 2
        v23 := []int{}
        var v24 bool = false
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint64 = 18446744073709551615
        var v27 float32 = (v3 - (float32(((((3.14159 - 1e10) * 0.1) - -2.25) / (0.1 + 1.5))) - float32(3.14159)))
        if (1e10 == ((0.1 * 1e10) / 1.5)) {
            var v28 uint64 = 1
            _ = v28
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    v21 = v21
    v2 = (^(v2))
    for fi29 := 0; fi29 < 9; fi29++ {
        if true {
            if false {
                if true {
                    if (98 <= 4294967295) {
                        var v30 string = ""
                        _ = v30
                    } else {
                        var v31 uint16 = 65535
                        v3 = ((-2.25 + 1e10) - (v3 - v3))
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 8; fi32++ {
                        v2 += (v2 & v2)
                        var v33 uint = 153
                        fmt.Println(fi29, v21)
                        _ = fi32
                        _ = v33
                    }
                } else {
                    v21 = v21
                    var v34 float32 = 1.5
                    v35 := GenS1{F0: 2}
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 16; fi36++ {
                    switch 526 {
                    case 32767:
                        v37 := GenS1{F0: 1}
                        _ = v37
                    default:
                        v38 := GenS1{F0: 292}
                        _ = v38
                    }
                    v3 -= 0.1
                    _ = fi36
                }
            }
        }
        for fi39 := 0; fi39 < 12; fi39++ {
            v2 = 17
            v40 := GenS1{F0: 60}
            v2 -= (^(v2))
            var v41 int64 = 0
            v2 = v2
            var v42 int16 = -32768
            v43 := []int{fi29, (^(v21)), (fi29 | fi39)}
            _ = fi39
            _ = v40
            _ = v41
            _ = v42
            _ = v43
        }
        v44 := GenS1{F0: 65535}
        _ = fi29
        _ = v44
    }
    var v45 uint16 = 65535
    var v46 float32 = 1.5
    for fi47 := 0; fi47 < 12; fi47++ {
        v2 = (v2 - (v2 + v2))
        switch "a" {
        case "a":
            var v48 float32 = (float32((0.1 * (-2.25 - (3.14159 / -2.25)))) + 3.14159)
            _ = v48
        case "hello":
            for fi49 := 0; fi49 < 17; fi49++ {
                v50 := GenS2{F0: "\n\t", F1: ("a" + ("hello" + "\n\t")), F2: fi47}
                v2 -= (v2 & (^(v2)))
                _ = fi49
                _ = v50
            }
        default:
            var v51 int64 = 9223372036854775807
            _ = v51
        }
        _ = fi47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v45
    _ = v46
}
