package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 int) bool {
    p0 = p0
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 568}
        v8 := GenS1{F0: p0}
        var v9 int16 = int16(p0)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{(p0 - p0), 0, p0, p0}
    _ = p0
    _ = v10
    return true
}
func genFn1() uint16 {
    if (32767 <= 32767) {
        var v11 uint32 = 2147483648
        if false {
            v11 = 2147483648
            v12 := []int{}
            _ = v12
        } else {
            var v13 int32 = 1
            _ = v13
        }
        v11 = (v11 % ((v11) | 1))
        _ = v11
    } else {
        var v14 int16 = -32768
        _ = v14
    }
    var v15 uint = 42
    _ = v15
    return 581
}

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 2
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    v5 := []string{"foo bar", (("hello" + "\n\t") + "foo bar"), "", "\n\t"}
    switch "a" {
    case "foo bar":
        v2 = (v2 & (v2 + v2))
    case "hello":
        v1 -= float64(0.0)
    }
    if false {
        for fi16 := 0; fi16 < 15; fi16++ {
            if (!(2 <= 681)) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v17 float64 = float64((1.5 / -2.25))
                _ = v17
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                if true {
                    var v19 uint16 = 65535
                    switch 520 {
                    case -1:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    v19 += (v19 % (((^(v19))) | 1))
                    _ = v19
                }
                v2 = v2
                _ = fi18
            }
            if false {
                v2 = 345
                v1 -= (float64(0.1) / v1)
                var v20 bool = (true || true)
                _ = v20
            }
            _ = fi16
        }
    }
    v21 := v4[9223372036854775807]
    for fi22 := 0; fi22 < 7; fi22++ {
        v21 = v21
        v23 := GenS1{F0: 91}
        _ = fi22
        _ = v23
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v21 = v21
        fmt.Println(v21)
        _ = fi24
    }
    v2 += 4294967295
    var v25 uint8 = 2
    var v26 bool = v21
    var v27 uint32 = 2147483648
    v28 := GenS2{F0: ("foo bar" + "foo bar")}
    _ = len(v5)
    v29 := GenS3{F0: -508, F1: v2}
    v25 -= 2
    for fi30 := 0; fi30 < 6; fi30++ {
        if true {
            for fi31 := 0; fi31 < 13; fi31++ {
                if (0.1 <= 1e10) {
                    v21 = ((!v26) || (!v26))
                    switch "hello" {
                    case "\n\t":
                        var v32 int = fi31
                        var v33 uint8 = v25
                        if false {
                            var v34 string = "a"
                            _ = v34
                        }
                        _ = v32
                        _ = v33
                    default:
                        var v35 uint64 = v2
                        v25 = 42
                        _ = v35
                    }
                } else {
                    var v36 bool = true
                    var v37 string = ((("a" + "hello") + "hello") + (("foo bar" + ("hello" + ("hello" + "\n\t"))) + "\n\t"))
                    switch fi31 {
                    case 4294967295:
                        var v38 int16 = 0
                        _ = v38
                    default:
                        if len(v5) > 0 {
                            v5[0] = (v37 + v37)
                        }
                    }
                    v39 := GenS3{F0: 32767, F1: 921}
                    _ = v36
                    _ = v37
                    _ = v39
                }
                _ = fi31
            }
            if v21 {
                if (v26 && (112 > 42)) {
                    v40 := GenS1{F0: fi30}
                    _ = v40
                } else {
                    var v41 bool = v21
                    v5 = append(v5, "")
                    v1 = -2.25
                    _ = v41
                }
            } else {
                v42 := GenS2{F0: ""}
                _ = v42
            }
            v2 = v2
        }
        _ = fi30
    }
    switch 4294967295 {
    case 0:
        if v26 {
            delete(v3, ("a" + (("" + "a") + ("foo bar" + ""))))
            v1 += (float64((3.14159 * 0.0)) - (1e10 + (1e10 - ((1e10 - 1e10) - (v1 / (((float64(-2.25) / (v1 - (v1 / (((float64((3.14159 + 0.0)) - (v1 + (v1 + v1))) - v1) / 1e10)))) + v1) + v1))))))
            fmt.Println(v27)
        } else {
            var v43 int16 = 927
            _ = v43
        }
    case 2147483647:
        for fi44 := 0; fi44 < 7; fi44++ {
            if len(v5) > 0 {
                v5[0] = ("\n\t" + "foo bar")
            }
            switch "foo bar" {
            case "foo bar":
                v45 := GenS4{F0: (v25 << ((4294967295) & 63)), F1: 14, F2: 4294967295}
                _ = v45
            case "\n\t":
                var v46 uint16 = genFn1()
                _ = v46
            default:
                v47 := GenS2{F0: "foo bar"}
                _ = v47
            }
            _ = fi44
        }
    case 9223372036854775807:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
