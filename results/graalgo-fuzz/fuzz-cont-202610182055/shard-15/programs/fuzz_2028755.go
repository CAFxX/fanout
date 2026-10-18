package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint64) int {
    var v6 int = 99
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 bool = (false && (false || false))
        var v9 int32 = -2147483648
        v10 := GenS1{F0: 0}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch (("" + "a") + "") {
    case "hello":
        v11 := GenS2{F0: "\n\t"}
        _ = v11
    default:
        if (2147483647 == 1) {
            var v12 uint32 = 65535
            _ = v12
        }
    }
    _ = p0
    _ = v6
    return (v6 >> ((399) & 63))
}
func genFn1(p0 uint64) bool {
    v13 := GenS3{F0: 2, F1: 255}
    var v14 uint16 = 223
    _ = p0
    _ = v13
    _ = v14
    return true
}

func main() {
    v1 := []string{"foo bar", "", ""}
    v2 := make(map[string]bool)
    var v3 int32 = -2147483648
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    var v15 int32 = v3
    if false {
        fmt.Println(v3, v3)
        var v16 int16 = -32768
        if ((3.14159 <= ((3.14159 * -2.25) + 0.0)) || false) {
            var v17 uint16 = 2
            if (!true) {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v3 -= (v15 & (v3 % ((610) | 1)))
                    v5[1] = ("" + ("" + ("hello" + "hello")))
                    _ = fi18
                }
            } else {
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := GenS3{F0: v17, F1: 42}
                    v17 -= v17
                    _ = fi19
                    _ = v20
                }
                var v21 uint64 = 359
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
                _ = v21
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v17
        } else {
            if false {
                var v22 uint16 = 255
                v23 := GenS3{F0: v22, F1: 0}
                _ = v22
                _ = v23
            }
            var v24 uint64 = 18446744073709551615
            _ = v24
        }
        _ = v16
    }
    fmt.Println(v15)
    var v25 float32 = float32(1e10)
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int8 = 1
        if (-32768 >= -1) {
            var v28 uint16 = 567
            if (false && (0.0 < 1.5)) {
                var v29 int32 = 586
                fmt.Println(fi26, v15, fi26)
                _ = v29
            }
            if true {
                v3 -= (v15 | (v15 << ((18446744073709551615) & 63)))
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 int32 = (v15 - 1)
                    _ = fi30
                    _ = v31
                }
            } else {
                for fi32 := 0; fi32 < 1; fi32++ {
                    if true {
                        v28 -= (^(v28))
                        _ = len(v1)
                        v3 += v3
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                    var v33 uint32 = 2147483647
                    v34 := GenS2{F0: ("\n\t" + "hello")}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            _ = v28
        } else {
            delete(v5, 82)
            switch "foo bar" {
            case "a":
                fmt.Println(v3, v15, v25)
            case "":
                if (!(32767 <= 1)) {
                    var v35 int = 0
                    v15 = v3
                    _ = v35
                }
            default:
                var v36 uint = 8
                _ = v36
            }
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v15, v25)
    v15 = (^(v15))
    var v37 float64 = -2.25
    fmt.Println(v15)
    for fi38 := 0; fi38 < 9; fi38++ {
        if false {
            var v39 bool = (1 >= -32768)
            v40 := GenS4{F0: 466, F1: 0}
            _ = v39
            _ = v40
        }
        _ = fi38
    }
    if true {
        if false {
            fmt.Println(v37, v25, v3)
            var v41 int16 = 825
            _ = v41
        }
    } else {
        var v42 bool = (("hello" + "\n\t") < "")
        _ = v42
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        switch (fi43 >> ((634) & 63)) {
        case -49:
            var v44 uint8 = 0
            _ = v44
        case 873:
            switch "a" {
            case "a":
                var v45 uint32 = 381
                _ = v45
            case "":
                v46 := GenS2{F0: ("hello" + ("" + ((("\n\t" + "") + (("\n\t" + ("hello" + "a")) + (("hello" + ("foo bar" + "hello")) + ("\n\t" + "")))) + (("hello" + "") + ("foo bar" + "\n\t")))))}
                _ = v46
            default:
                v2["a"] = ((1 >= 246) || false)
            }
        case 128:
            switch 1 {
            case 128:
                v1 = append(v1, ("foo bar" + ""))
            case 860:
                if true {
                    var v47 int32 = (v3 & (v3 - v3))
                    v37 = float64(v25)
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 9; fi48++ {
                        if false {
                            v49 := GenS2{F0: "a"}
                            _ = v49
                        }
                        v50 := GenS5{F0: 65535}
                        _ = fi48
                        _ = v50
                    }
                }
            }
        }
        if true {
            v37 += -2.25
            v25 += ((v25 + ((float32(float64(((v25 + v25) - (v25 / (float32(-2.25) + v25))))) - v25) - v25)) / 3.14159)
        }
        var v51 float64 = v37
        _ = fi43
        _ = v51
    }
    fmt.Println(v15, v37)
    v37 = v37
    for fi52 := 0; fi52 < 17; fi52++ {
        if (!(18446744073709551615 != 278)) {
            fmt.Println(v15)
            v37 += 0.1
        } else {
            v53 := GenS3{F0: 65535, F1: 103}
            if (true || false) {
                fmt.Println(v3)
                v54 := GenS4{F0: 255, F1: 83}
                _ = v54
            } else {
                fmt.Println(v3, fi52)
                v55 := GenS4{F0: 2, F1: 100}
                var v56 float32 = (((0.1 + (((v25 - 0.0) - float32(0.1)) * float32(-2.25))) + (0.0 / float32(v37))) * ((v25 / v25) - -2.25))
                _ = v55
                _ = v56
            }
            _ = v53
        }
        _ = fi52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
