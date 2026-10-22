package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int; F2 int16 }
type GenS5 struct { F0 int32; F1 int16; F2 int16 }
func genFn0() int {
    v7 := GenS1{F0: 409}
    if true {
        v8 := []int{146, 0, -2147483648}
        _ = v8
    }
    var v9 uint = 1
    _ = v7
    _ = v9
    return 194
}
func genFn1(p0 string) int32 {
    switch -1 {
    case -426:
        p0 = (p0 + p0)
    case 42:
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := []int{genFn0(), genFn0(), genFn0()}
                var v13 uint16 = 32767
                fmt.Println(v13)
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(fi10)
            _ = fi10
        }
    case -9223372036854775808:
        var v14 int32 = 2147483647
        _ = v14
    default:
        switch "\n\t" {
        case "a":
            for fi15 := 0; fi15 < 20; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    if true {
                        var v17 float64 = float64(float32(-2.25))
                        fmt.Println(v17, p0, fi15)
                        var v18 float32 = float32(((1.5 * 3.14159) / (float64((((3.14159 - 3.14159) - -2.25) / ((((float32(1.5) + ((1.5 / 3.14159) - 1.5)) * (0.1 / 1.5)) + 3.14159) - 1e10))) * -2.25)))
                        _ = v17
                        _ = v18
                    } else {
                        p0 = ((p0 + (p0 + (p0 + ((((p0 + ("hello" + ((((((p0 + p0) + (((p0 + "\n\t") + p0) + p0)) + ((p0 + p0) + "")) + (p0 + ("foo bar" + p0))) + "foo bar") + p0))) + "foo bar") + p0) + p0)))) + (((p0 + p0) + p0) + "foo bar"))
                        v19 := GenS2{F0: -1}
                        if (4294967295 != 42) {
                            p0 = p0
                            p0 = ""
                        } else {
                            var v20 int64 = 1
                            v21 := []int{0, genFn0(), (fi16 / ((-2147483648) | 1)), (-(fi15))}
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    }
                    p0 = ""
                    fmt.Println(p0, fi16, p0)
                    fmt.Println(fi15, fi16, fi16)
                    _ = fi16
                }
                var v22 uint8 = 235
                p0 = (p0 + p0)
                _ = fi15
                _ = v22
            }
        }
    }
    var v23 int16 = -32768
    v24 := GenS1{F0: 4294967295}
    var v25 float32 = 3.14159
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 uint = 255
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = v23
    _ = v24
    _ = v25
    return 1
}

func main() {
    var v1 uint16 = 255
    var v2 float32 = 1e10
    v3 := []int{9223372036854775807}
    v4 := []string{"", "hello", "hello"}
    v5 := make(map[string]int)
    v6 := []string{"hello"}
    var v28 uint8 = 127
    fmt.Println(v1, v28, v1)
    if true {
        _ = len(v5)
        v29 := GenS2{F0: 0}
        _ = v29
    }
    v30 := v5[("hello" + "a")]
    if true {
        v31 := GenS3{F0: "foo bar"}
        if false {
            var v32 string = ("a" + ((("\n\t" + "\n\t") + "\n\t") + ((("a" + "") + "a") + "hello")))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v33 := GenS4{F0: (v32 + v32), F1: (v30 & (v30 + v30)), F2: 32767}
            _ = v32
            _ = v33
        } else {
            for fi34 := 0; fi34 < 19; fi34++ {
                _ = len(v4)
                v30 -= 128
                v35 := GenS4{F0: "\n\t", F1: (fi34 | (-(fi34))), F2: 0}
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 12; fi36++ {
                fmt.Println(v28, fi36, v2)
                if false {
                    var v37 int32 = 42
                    v28 = (^(v28))
                    var v38 int = (v30 / ((fi36) | 1))
                    _ = v37
                    _ = v38
                } else {
                    v39 := GenS4{F0: "\n\t", F1: 1, F2: -1}
                    var v40 int32 = 42
                    _ = v39
                    _ = v40
                }
                _ = fi36
            }
            switch (((("\n\t" + "") + ("foo bar" + "foo bar")) + "foo bar") + "") {
            case "\n\t":
                var v41 uint64 = 18446744073709551615
                _ = v41
            case "":
                for fi42 := 0; fi42 < 3; fi42++ {
                    var v43 float32 = 1.5
                    fmt.Println(v43, v28, v28)
                    _ = fi42
                    _ = v43
                }
            case "foo bar":
                if true {
                    v44 := GenS3{F0: "a"}
                    v2 = -2.25
                    v28 += 0
                    _ = v44
                } else {
                    v45 := GenS5{F0: 42, F1: 519, F2: 512}
                    for fi46 := 0; fi46 < 15; fi46++ {
                        for fi47 := 0; fi47 < 16; fi47++ {
                            v30 += 9223372036854775807
                            v48 := GenS3{F0: ("a" + "hello")}
                            _ = fi47
                            _ = v48
                        }
                        _ = fi46
                    }
                    v28 = (^(v28))
                    _ = v45
                }
            }
        }
        for fi49 := 0; fi49 < 5; fi49++ {
            for fi50 := 0; fi50 < 20; fi50++ {
                switch (fi49 % ((v30) | 1)) {
                case 9223372036854775807:
                    v28 -= v28
                case 32767:
                    if true {
                        switch -608 {
                        case 352:
                            delete(v5, (("" + "hello") + "foo bar"))
                        }
                        if (float32(0.1) > 1e10) {
                            v5[(("" + "\n\t") + ("\n\t" + "a"))] = (^(v30))
                        }
                        v28 += (v28 << ((255) & 63))
                    }
                case -9223372036854775808:
                    v1 = (v1 >> ((2) & 63))
                }
                if (!((false || true) || false)) {
                    v51 := GenS1{F0: 4294967295}
                    v52 := GenS5{F0: genFn1("\n\t"), F1: 42, F2: 1}
                    _ = v51
                    _ = v52
                } else {
                    v28 = (v28 | 255)
                    v53 := v5["\n\t"]
                    _ = v53
                }
                var v54 uint8 = (v28 + v28)
                _ = fi50
                _ = v54
            }
            _ = fi49
        }
        _ = v31
    }
    v55 := GenS4{F0: ((("hello" + ("foo bar" + "a")) + "foo bar") + "a"), F1: (v30 / (((v30 * 0)) | 1)), F2: 0}
    fmt.Println(v28, v2, v1)
    v2 -= (((v2 * 0.1) * (1e10 * 1e10)) * (v2 - (v2 * (v2 - v2))))
    v3 = append(v3, v30)
    v28 = (v28 - (v28 + (v28 & 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
