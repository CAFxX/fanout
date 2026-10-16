package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int32; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 bool; F1 int8; F2 int8 }
func genFn0(p0 uint64) uint {
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = (p0 ^ p0)
        _ = fi7
    }
    v8 := GenS1{F0: 2147483648, F1: p0}
    _ = p0
    _ = v8
    return 65535
}
func genFn1() int64 {
    v9 := GenS2{F0: 4294967295}
    for fi10 := 0; fi10 < 6; fi10++ {
        if false {
            if false {
                fmt.Println(fi10)
                if true {
                    fmt.Println(fi10)
                    fmt.Println(fi10)
                } else {
                    var v11 int = -846
                    var v12 uint64 = 515
                    _ = v11
                    _ = v12
                }
                v13 := GenS3{F0: 255, F1: 2147483647, F2: false}
                _ = v13
            }
            switch (fi10 * -1) {
            case -9223372036854775808:
                var v14 int8 = 1
                _ = v14
            default:
                fmt.Println(fi10)
            }
        } else {
            fmt.Println(fi10)
            v15 := []int{32767}
            _ = v15
        }
        switch "\n\t" {
        case "\n\t":
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := []int{(fi16 >> ((18446744073709551615) & 63))}
                var v18 int64 = 9223372036854775807
                if false {
                    if false {
                        fmt.Println(fi16, v18)
                        var v19 int32 = 1
                        _ = v19
                    } else {
                        if len(v17) > 0 {
                            v17[0] = (fi16 ^ (^(fi16)))
                        }
                        var v20 int64 = (v18 + 267)
                        _ = v20
                    }
                    if len(v17) > 0 {
                        v17[0] = (-(fi16))
                    }
                    v18 = -742
                } else {
                    var v21 uint16 = 394
                    var v22 int16 = 1
                    var v23 string = ("foo bar" + "foo bar")
                    _ = v21
                    _ = v22
                    _ = v23
                }
                switch 128 {
                case 42:
                    v18 += v18
                case 128:
                    if (true || false) {
                        v18 -= (v18 + 996)
                        fmt.Println(fi10, v18, fi16)
                        v24 := GenS4{F0: (416 <= -435)}
                        _ = v24
                    }
                default:
                    v25 := GenS5{F0: "foo bar", F1: 723}
                    _ = v25
                }
                _ = fi16
                _ = v17
                _ = v18
            }
        case "hello":
            if true {
                var v26 uint32 = 0
                fmt.Println(fi10, fi10)
                v27 := []int{fi10}
                _ = v26
                _ = v27
            }
        }
        _ = fi10
    }
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 uint16 = 497
    var v4 float64 = 0.1
    var v5 int32 = 0
    var v6 float64 = float64(-2.25)
    if false {
        v28 := GenS6{F0: ((26 <= 127) && (!false)), F1: 127, F2: -37}
        _ = v28
    } else {
        v6 = v4
        v29 := GenS3{F0: 0, F1: v5, F2: ((65535 >= 4294967295) && false)}
        _ = v29
    }
    if false {
        for fi30 := 0; fi30 < 12; fi30++ {
            if true {
                var v31 int32 = v5
                var v32 int16 = -32768
                _ = v31
                _ = v32
            } else {
                v5 += (v5 % ((v5) | 1))
                switch -9223372036854775808 {
                case -9223372036854775808:
                    if true {
                        v3 = (v3 | 0)
                        var v33 uint32 = 0
                        v34 := []int{}
                        _ = v33
                        _ = v34
                    } else {
                        v3 = v3
                        v35 := GenS4{F0: (691 > 2)}
                        for fi36 := 0; fi36 < 6; fi36++ {
                            var v37 string = (("foo bar" + "foo bar") + ((("hello" + ("" + "\n\t")) + ("hello" + "\n\t")) + ""))
                            var v38 string = ""
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = v35
                    }
                case -860:
                    var v39 string = "\n\t"
                    _ = v39
                default:
                    var v40 string = "hello"
                    _ = v40
                }
                switch (^(fi30)) {
                case 1:
                    fmt.Println(fi30)
                default:
                    v6 += v4
                }
            }
            _ = fi30
        }
    }
    var v41 int32 = (v5 / (((-(v5))) | 1))
    fmt.Println(v4, v41, v4)
    for fi42 := 0; fi42 < 20; fi42++ {
        _ = len(v2)
        fmt.Println(v5)
        _ = fi42
    }
    v5 = (v5 & v5)
    fmt.Println(v41, v5, v5)
    if true {
        for fi43 := 0; fi43 < 11; fi43++ {
            var v44 string = "hello"
            switch 32767 {
            case 709:
                fmt.Println(v3, v6)
            }
            _ = fi43
            _ = v44
        }
    } else {
        v45 := v2[798]
        v46 := []int{}
        var v47 int32 = (v41 % (((v41 + v41)) | 1))
        _ = v45
        _ = v46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
