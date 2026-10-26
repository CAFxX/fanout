package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 bool }
func genFn0(p0 string, p1 int64) uint64 {
    var v5 uint = 4294967295
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 bool = (true || true)
        fmt.Println(v7, v5)
        fmt.Println(v7)
        _ = fi6
        _ = v7
    }
    v8 := []int{32767, 151}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 455
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint = 455
    v3 := []string{"hello", "hello", "a"}
    var v4 int64 = -1
    _ = len(v3)
    var v9 int8 = 100
    switch 9223372036854775807 {
    case 32768:
        if (false && (1 > 1000)) {
            v2 -= v2
            var v10 float64 = 0.0
            if false {
                v9 += v9
            } else {
                v10 -= (0.1 * float64(((-2.25 * (float32(-2.25) - ((1.5 / (1.5 + (float32(v10) + (1e10 / -2.25)))) / ((-2.25 - float32(v10)) * (3.14159 / (0.1 * 1.5)))))) * 1e10)))
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 uint = v2
                    _ = fi11
                    _ = v12
                }
            }
            _ = v10
        } else {
            var v13 uint16 = 32767
            v14 := GenS1{F0: (v4 + v4)}
            var v15 string = "foo bar"
            _ = v13
            _ = v14
            _ = v15
        }
    case 9223372036854775807:
        _ = len(v3)
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    default:
        switch 128 {
        case 2147483647:
            if ((!(9223372036854775807 <= -1)) || true) {
                v16 := GenS1{F0: 42}
                for fi17 := 0; fi17 < 18; fi17++ {
                    switch (("" + (("hello" + "a") + ("\n\t" + ""))) + "\n\t") {
                    case "a":
                        if (v9 <= v9) {
                            for fi18 := 0; fi18 < 1; fi18++ {
                                var v19 uint64 = 42
                                _ = fi18
                                _ = v19
                            }
                            fmt.Println(v2)
                        }
                    case "\n\t":
                        v4 = (v4 % (((v4 & v4)) | 1))
                    }
                    fmt.Println(v2, fi17)
                    _ = fi17
                }
                v9 -= (v9 | (v9 << (((v2 & 0)) & 63)))
                _ = v16
            }
        case 85:
            v2 -= 42
        default:
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = (("hello" + "foo bar") + "foo bar")
    }
    v2 = (v2 / ((v2) | 1))
    if false {
        switch -571 {
        case 1:
            v9 -= v9
        case 744:
            if len(v3) > 0 {
                v3[0] = ("hello" + (("foo bar" + "a") + "foo bar"))
            }
        case 0:
            switch (("hello" + "hello") + "") {
            case "":
                v2 = 549
            case "hello":
                var v20 uint = (^(v2))
                _ = v20
            case "a":
                fmt.Println(v1, v4)
            }
        default:
            fmt.Println(v4)
        }
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        _ = len(v3)
        switch fi22 {
        case 637:
            v23 := GenS2{F0: "foo bar"}
            _ = v23
        case -9223372036854775808:
            if (!true) {
                _ = len(v3)
                v24 := GenS2{F0: "foo bar"}
                var v25 int = (fi22 / ((829) | 1))
                _ = v24
                _ = v25
            } else {
                if true {
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                }
                _ = len(v3)
                for fi26 := 0; fi26 < 1; fi26++ {
                    fmt.Println(v1, fi26)
                    _ = fi26
                }
            }
        case 975:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            for fi27 := 0; fi27 < 2; fi27++ {
                var v28 bool = true
                _ = fi27
                _ = v28
            }
        }
        _ = fi22
    }
    v3 = append(v3, "\n\t")
    v29 := GenS3{F0: 2147483647, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
