package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint32) int {
    if (((!false) || (!(565 <= 255))) || ((1 >= 18446744073709551615) && (!((42 > -1) && (!true))))) {
        p1 -= 4294967295
    } else {
        if (541 <= 42) {
            fmt.Println(p1, p1)
            for fi4 := 0; fi4 < 18; fi4++ {
                v5 := []int{fi4, fi4, fi4, fi4}
                var v6 string = ""
                _ = fi4
                _ = v5
                _ = v6
            }
        }
    }
    p1 -= p1
    p1 = 0
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 0
    var v3 int32 = -2147483648
    v7 := GenS1{F0: 1}
    if (!(!false)) {
        fmt.Println(v3)
        switch 42 {
        case 42:
            var v8 int64 = 9223372036854775807
            _ = v8
        case 2147483647:
            for fi9 := 0; fi9 < 11; fi9++ {
                switch genFn0(0, 42) {
                case 0:
                    var v10 float64 = 3.14159
                    _ = v10
                case -9223372036854775808:
                    v11 := []int{(fi9 / ((fi9) | 1))}
                    _ = v11
                case 127:
                    v12 := GenS1{F0: 100}
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v3 = v3
                        fmt.Println(fi13)
                        var v14 uint = 1
                        _ = fi13
                        _ = v14
                    }
                }
                _ = fi9
            }
        default:
            v15 := []int{}
            _ = v15
        }
    }
    var v16 int64 = 1
    switch 2147483647 {
    case 42:
        switch v1 {
        case "":
            for fi17 := 0; fi17 < 11; fi17++ {
                if true {
                    v2 = (^(v2))
                    if (!(-2.25 > (((-2.25 / 1.5) - 0.0) - 1e10))) {
                        var v18 uint8 = 0
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v20 := []int{(fi19 & (fi17 % (((fi19 / ((genFn0(1, 420)) | 1))) | 1))), 2147483647, fi19, (fi19 ^ fi17)}
                            _ = fi19
                            _ = v20
                        }
                        v2 = 1
                        _ = v18
                    }
                    v21 := []int{(fi17 % ((fi17) | 1))}
                    _ = v21
                } else {
                    var v22 int64 = (v16 % ((v16) | 1))
                    _ = v22
                }
                _ = fi17
            }
        default:
            v23 := GenS2{F0: 0}
            _ = v23
        }
    case 32767:
        if (!(!(!false))) {
            v3 -= -389
            switch "" {
            case "\n\t":
                v24 := []int{genFn0(-466, 1), genFn0((v16 & -647), 245), 1, int(v3)}
                _ = v24
            default:
                v1 = ("hello" + v1)
            }
            var v25 uint = 207
            _ = v25
        } else {
            var v26 int32 = (^(v3))
            fmt.Println(v3, v26, v16)
            v2 -= v2
            _ = v26
        }
    case -1:
        v27 := []int{genFn0(-1, 656), -441}
        _ = v27
    }
    v28 := []int{0, -9223372036854775808, -245}
    var v29 bool = (0 != 42)
    v1 = "hello"
    if len(v28) > 0 {
        v28[0] = -2147483648
    }
    var v30 uint8 = 127
    switch v1 {
    case "":
        if v29 {
            var v31 bool = (!v29)
            var v32 uint8 = v30
            _ = v31
            _ = v32
        }
    }
    var v33 float32 = (0.0 / 3.14159)
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v7
}
