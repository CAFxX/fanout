package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint) uint64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: true}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 bool = (true && ("\n\t" != ""))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: false}
    var v11 string = (((("foo bar" + (("" + "a") + "a")) + (("a" + ("hello" + "")) + "")) + "a") + "hello")
    _ = p0
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 uint16) uint8 {
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int32 = 494
        if false {
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := []int{fi12}
                var v16 uint32 = 2147483648
                v17 := GenS1{F0: 531, F1: (false || (false || false))}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := []int{fi12, fi12}
            _ = v18
        }
        if true {
            if (("\n\t" + "foo bar") != "a") {
                fmt.Println(v13, v13, v13)
            }
            for fi19 := 0; fi19 < 19; fi19++ {
                p0 = (p0 % (((p0 ^ 1)) | 1))
                p0 += p0
                _ = fi19
            }
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 int32 = (v13 * v13)
                if (-2.25 <= 0.1) {
                    v22 := []int{fi12}
                    _ = v22
                } else {
                    var v23 int32 = (v21 + -547)
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            switch "hello" {
            case "\n\t":
                v24 := []int{(-(fi12))}
                _ = v24
            case "foo bar":
                if false {
                    p0 = p0
                    if (1 < 2) {
                        for fi25 := 0; fi25 < 11; fi25++ {
                            v13 = v13
                            fmt.Println(v13, fi12)
                            _ = fi25
                        }
                        var v26 uint32 = 4294967295
                        _ = v26
                    }
                }
            case "hello":
                var v27 float32 = (3.14159 / -2.25)
                _ = v27
            default:
                var v28 int8 = -111
                _ = v28
            }
        }
        _ = fi12
        _ = v13
    }
    var v29 float64 = 0.0
    var v30 uint8 = 174
    v31 := GenS1{F0: 4294967295, F1: false}
    v29 -= float64(float32(((v29 - 1.5) - -2.25)))
    _ = p0
    _ = v29
    _ = v30
    _ = v31
    return (v30 | 122)
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 2
    var v3 int16 = -32768
    var v4 int16 = v3
    v5 := []int{2147483647, 128, -601}
    var v32 int64 = 86
    var v33 uint64 = 4294967295
    v2 = (v2 & 492)
    v34 := GenS2{F0: 0}
    var v35 bool = false
    v1 = ""
    var v36 int8 = 0
    v36 = (v36 / ((v36) | 1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v37 := GenS1{F0: 2, F1: (!(32767 == v4))}
    v36 = (v36 | v36)
    v38 := GenS1{F0: (^(v33)), F1: v35}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
