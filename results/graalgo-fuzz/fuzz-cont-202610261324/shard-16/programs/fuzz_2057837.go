package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint64; F1 bool; F2 int64 }
func genFn0() bool {
    var v6 float64 = (3.14159 * -2.25)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(^(fi7))}
        v6 = 1.5
        var v9 uint8 = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 331, F1: 22, F2: 428}
    _ = v10
    _ = v6
    return false
}
func genFn1(p0 int) int {
    var v11 float64 = 0.0
    v12 := []int{(p0 % (((p0 + p0)) | 1)), 32768, p0}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v12 = append(v12, (p0 << ((509) & 63)))
    var v13 uint32 = 65535
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return (p0 & 2147483647)
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = (true && true)
    v3 := make(map[string]bool)
    var v4 uint32 = 2147483648
    v5 := make(map[string]int)
    v14 := GenS2{F0: 0, F1: (v1 % (((v1 - 21)) | 1)), F2: 18446744073709551615}
    var v15 int8 = 0
    v2 = ("a" == "\n\t")
    for fi16 := 0; fi16 < 9; fi16++ {
        switch (^(fi16)) {
        case 42:
            fmt.Println(v1, v2, v4)
        case -893:
            if true {
                v1 = v1
                v4 = v4
                v17 := GenS1{F0: 65535, F1: 128, F2: uint16(v4)}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 int8 = v15
                    if (v2 && ((4294967295 < (v4 ^ v4)) && true)) {
                        if (2 < 65535) {
                            if v2 {
                                if v2 {
                                    v4 = 4294967295
                                }
                                var v20 uint32 = (v4 / ((v4) | 1))
                                _ = v20
                            } else {
                                if (0.0 != -2.25) {
                                    v21 := GenS3{F0: "", F1: -32768}
                                    switch ((("a" + "hello") + (((("hello" + "a") + ("foo bar" + "")) + (("foo bar" + ("\n\t" + "foo bar")) + ((("" + "") + "\n\t") + "\n\t"))) + "a")) + "a") {
                                    case "foo bar":
                                        fmt.Println(fi16, v4, v19)
                                    case "":
                                        var v22 uint16 = (v1 - (v1 | (v1 & v1)))
                                        _ = v22
                                    }
                                    _ = v21
                                }
                                delete(v5, ((("\n\t" + ((((("hello" + "a") + "a") + ("foo bar" + "hello")) + "\n\t") + "a")) + ("" + "a")) + "\n\t"))
                                var v23 uint = 255
                                _ = v23
                            }
                            v24 := []int{fi18, (fi16 & (^(fi16)))}
                            _ = v24
                        } else {
                            var v25 int = (fi16 >> ((uint(v1)) & 63))
                            _ = v25
                        }
                        switch (("" + "hello") + "foo bar") {
                        case "a":
                            var v26 int = 243
                            _ = v26
                        case "hello":
                            if (2147483647 <= 420) {
                                if false {
                                    if (-1 < int32(v19)) {
                                        v19 += (v15 * v19)
                                    } else {
                                        switch (fi16 << ((395) & 63)) {
                                        case 127:
                                            v4 = (v4 ^ v4)
                                        default:
                                            v2 = v2
                                        }
                                    }
                                    var v27 uint64 = 605
                                    _ = v27
                                }
                                var v28 float32 = (3.14159 * 1.5)
                                v19 = -126
                                _ = v28
                            } else {
                                v29 := GenS2{F0: 1, F1: v1, F2: 42}
                                switch 32768 {
                                case 22:
                                    if (65535 <= uint(v1)) {
                                        var v30 uint = 243
                                        if (uint(v1) > v30) {
                                            var v31 uint = (v30 + (v30 & v30))
                                            v32 := []int{}
                                            _ = v31
                                            _ = v32
                                        }
                                        v30 -= v30
                                        v33 := []int{(fi18 + -9223372036854775808), fi18, fi18}
                                        _ = v30
                                        _ = v33
                                    }
                                default:
                                    fmt.Println(fi16, v2)
                                }
                                _ = v29
                            }
                        case "\n\t":
                            if v2 {
                                v2 = v2
                                v2 = true
                            } else {
                                fmt.Println(v19, v1)
                                v34 := GenS2{F0: -940, F1: (^(v1)), F2: 18446744073709551615}
                                _ = v34
                            }
                        }
                        v4 += v4
                    }
                    _ = fi18
                    _ = v19
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS3{F0: ("\n\t" + (("a" + "a") + ("\n\t" + ("" + (("foo bar" + "") + "\n\t"))))), F1: 112}
                    var v37 int8 = (v15 + 127)
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 2; fi38++ {
                    v39 := v3[(("hello" + "") + ("" + ("foo bar" + "a")))]
                    v5["a"] = fi16
                    _ = fi38
                    _ = v39
                }
            }
        case -9223372036854775808:
            v40 := GenS3{F0: "\n\t", F1: 637}
            _ = v40
        }
        v2 = false
        _ = fi16
    }
    fmt.Println(v1, v1, v4)
    var v41 bool = v2
    v42 := GenS4{F0: 18446744073709551615, F1: v41, F2: 707}
    fmt.Println(v1, v2, v4)
    v2 = v41
    if v41 {
        v43 := GenS2{F0: -14, F1: 255, F2: 4294967295}
        var v44 bool = (128 <= -2147483648)
        _ = v43
        _ = v44
    } else {
        v2 = (v2 || false)
    }
    var v45 uint32 = (^(v4))
    switch 127 {
    case 42:
        var v46 string = "a"
        _ = v46
    default:
        v1 = 32767
    }
    delete(v5, "")
    v2 = (v2 && (!(!v41)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v5
}
