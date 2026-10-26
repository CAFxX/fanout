package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 uint, p1 uint8, p2 uint) string {
    var v6 uint8 = p1
    if (false || ((v6 % ((v6) | 1)) <= v6)) {
        var v7 int8 = -1
        v8 := GenS1{F0: 12}
        _ = v7
        _ = v8
    } else {
        if (!(-352 == 2147483647)) {
            p1 = 127
        } else {
            var v9 bool = false
            _ = v9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS2{F0: p0, F1: p1}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "\n\t"
}
func genFn1(p0 int, p1 int32, p2 uint16) uint {
    p1 = (-(p1))
    v12 := []int{(^(p0)), (^(p0))}
    if (true && ((true && false) || true)) {
        for fi13 := 0; fi13 < 20; fi13++ {
            fmt.Println(p2, fi13)
            p2 -= p2
            _ = fi13
        }
    }
    var v14 int64 = -1
    v15 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v14
    _ = v15
    return 774
}

func main() {
    var v1 bool = true
    var v2 float64 = 3.14159
    var v3 uint = 0
    var v4 bool = v1
    var v5 uint16 = 2
    switch 404 {
    case 1:
        fmt.Println(v4, v5, v3)
    }
    var v16 uint32 = 892
    fmt.Println(v5, v16, v1)
    var v17 int16 = 1
    if true {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 float32 = ((0.1 + 0.0) - 0.0)
            v16 -= 2
            _ = fi18
            _ = v19
        }
        v20 := []int{981}
        if v4 {
            _ = len(v20)
        } else {
            var v21 int64 = -618
            for fi22 := 0; fi22 < 17; fi22++ {
                v16 -= v16
                _ = fi22
            }
            _ = v21
        }
        _ = v20
    }
    var v23 int32 = -2147483648
    v1 = v1
    var v24 bool = (((((("" + "hello") + ("\n\t" + "")) + "\n\t") + ("hello" + "hello")) + ("foo bar" + "\n\t")) > "\n\t")
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := []int{}
        _ = fi25
        _ = v26
    }
    v27 := []int{}
    if ((float64(1e10) < v2) && false) {
        if v4 {
            var v28 float32 = (((3.14159 - -2.25) + 0.0) - -2.25)
            _ = v28
        } else {
            _ = len(v27)
            v4 = true
            for fi29 := 0; fi29 < 2; fi29++ {
                v24 = (!true)
                if ((832 != 487) || v4) {
                    v17 -= 0
                    var v30 int64 = -354
                    _ = v30
                }
                v23 += (v23 - -775)
                _ = fi29
            }
        }
        if false {
            switch "a" {
            case "\n\t":
                var v31 uint = 18446744073709551615
                _ = v31
            case "":
                if true {
                    var v32 int = 32768
                    _ = v32
                }
            case "hello":
                var v33 uint = genFn1(4294967295, v23, uint16(v16))
                _ = v33
            default:
                _ = len(v27)
            }
            if true {
                v23 = v23
                var v34 int32 = (^(v23))
                _ = v34
            }
        }
        var v35 float64 = (v2 - 1.5)
        _ = v35
    } else {
        fmt.Println(v3, v5)
        for fi36 := 0; fi36 < 14; fi36++ {
            switch 679 {
            case 32768:
                fmt.Println(v2, v4)
            case 928:
                var v37 int8 = 0
                _ = v37
            case 127:
                if len(v27) > 0 {
                    v27[0] = (-(fi36))
                }
            }
            switch fi36 {
            case -657:
                var v38 int = fi36
                _ = v38
            case -2147483648:
                v39 := GenS1{F0: -361}
                _ = v39
            case 2147483647:
                v40 := GenS3{F0: (v16 ^ uint32(v5)), F1: uint8(v5), F2: 0}
                _ = v40
            default:
                switch 9223372036854775807 {
                case 1:
                    v3 -= 65535
                case -1:
                    fmt.Println(v4)
                default:
                    var v41 int64 = 42
                    _ = v41
                }
            }
            _ = fi36
        }
        v2 = float64((0.1 + 0.0))
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    var v42 uint = 2
    v4 = ((1.5 * -2.25) <= 3.14159)
    v23 = v23
    var v43 int8 = 6
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
