package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 string }
func genFn0() uint {
    var v6 float32 = -2.25
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := []int{(fi7 + (fi7 >> ((18446744073709551615) & 63))), (fi7 & 4294967295), fi7}
            var v9 int16 = 32767
            v9 += 936
            _ = fi7
            _ = v8
            _ = v9
        }
        if ((!(-42 < 32)) && true) {
            var v10 bool = false
            v11 := []int{2147483647, -263}
            if false {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                switch 4294967295 {
                case -1:
                    _ = len(v11)
                }
                var v12 uint16 = 373
                _ = v12
            } else {
                v11 = append(v11, 127)
                var v13 int = 1
                if true {
                    v10 = (42 <= 42)
                }
                _ = v13
            }
            _ = v10
            _ = v11
        }
    }
    _ = v6
    return 0
}
func genFn1(p0 int16) int16 {
    switch "" {
    case "a":
        v14 := GenS1{F0: 149}
        _ = v14
    case "\n\t":
        if false {
            v15 := GenS1{F0: 2}
            for fi16 := 0; fi16 < 6; fi16++ {
                for fi17 := 0; fi17 < 9; fi17++ {
                    p0 = p0
                    if (127 < 128) {
                        v18 := []int{(^(fi16)), (fi17 ^ 9223372036854775807)}
                        p0 = (p0 / (((p0 ^ p0)) | 1))
                        var v19 uint64 = 2
                        _ = v18
                        _ = v19
                    } else {
                        p0 = p0
                    }
                    var v20 int64 = 1
                    _ = fi17
                    _ = v20
                }
                if (((1.5 + (-2.25 * 3.14159)) == float32((float64(3.14159) / float64(-2.25)))) && false) {
                    v21 := []int{fi16, fi16, 4294967295}
                    _ = v21
                } else {
                    v22 := GenS1{F0: 847}
                    _ = v22
                }
                _ = fi16
            }
            _ = v15
        } else {
            var v23 uint32 = 2
            var v24 bool = (!false)
            var v25 float32 = (0.0 / 1.5)
            _ = v23
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        switch 764 {
        case 2147483647:
            fmt.Println(p0, p0)
        case 0:
            var v27 uint32 = 2147483648
            _ = v27
        case 4294967295:
            v28 := []int{(fi26 - (fi26 << ((18446744073709551615) & 63)))}
            _ = v28
        }
        p0 += (p0 | (p0 >> ((395) & 63)))
        _ = fi26
    }
    _ = p0
    return p0
}

func main() {
    var v1 int = 128
    var v2 uint64 = 42
    v3 := []int{(v1 + 562)}
    var v4 uint16 = 452
    var v5 uint32 = 1
    v2 = uint64(v4)
    fmt.Println(v2, v1)
    var v29 string = ("\n\t" + "hello")
    v30 := GenS1{F0: 143}
    v31 := GenS1{F0: 18446744073709551615}
    v32 := GenS1{F0: v2}
    v2 += (v2 - (^(v2)))
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 int8 = -128
        v35 := GenS2{F0: 0, F1: v29}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
