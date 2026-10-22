package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint; F1 uint32; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 string; F1 int }
func genFn0(p0 int16, p1 int64) uint32 {
    v4 := GenS1{F0: 42, F1: 32768}
    p0 += p0
    if (42 != -128) {
        var v5 int = -849
        _ = v5
    }
    v6 := []int{-9223372036854775808, -412, -11, -2147483648}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 0.0
    var v3 uint32 = 148
    v7 := GenS1{F0: 127, F1: 9223372036854775807}
    if (((-1 != 1) || ((true || false) || true)) && false) {
        var v8 float32 = 1.5
        _ = v8
    } else {
        v9 := []int{4294967295, -1}
        var v10 string = "foo bar"
        if true {
            if (-1 <= -1) {
                var v11 int64 = 1
                v10 = (v10 + ("hello" + ""))
                _ = v11
            }
            for fi12 := 0; fi12 < 2; fi12++ {
                if true {
                    var v13 uint64 = 4294967295
                    _ = v13
                }
                fmt.Println(v2)
                v14 := GenS1{F0: 100, F1: (fi12 + (fi12 + fi12))}
                _ = fi12
                _ = v14
            }
            var v15 uint32 = 1
            _ = v15
        }
        _ = v10
        _ = v9
    }
    if true {
        if ((!((1e10 != (float32((0.0 + 0.1)) / (v2 * v2))) && (!(112 >= 2)))) && false) {
            var v16 bool = false
            if (42 > (v1 | (v1 * (v1 | v1)))) {
                if ("\n\t" < "foo bar") {
                    v16 = true
                    fmt.Println(v1, v1)
                } else {
                    var v17 uint64 = 4294967295
                    var v18 int32 = 1
                    _ = v17
                    _ = v18
                }
                var v19 uint16 = v1
                _ = v19
            } else {
                var v20 int64 = 42
                if v16 {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        fmt.Println(v16, v20, fi21)
                        _ = fi21
                    }
                    v22 := []int{32768}
                    if v16 {
                        var v23 int16 = int16(v1)
                        v24 := GenS2{F0: 0, F1: (v3 - v3), F2: 4294967295}
                        _ = v23
                        _ = v24
                    }
                    _ = v22
                } else {
                    v25 := GenS3{F0: "foo bar"}
                    for fi26 := 0; fi26 < 10; fi26++ {
                        v20 = v20
                        if (1e10 <= 0.1) {
                            var v27 uint = 255
                            for fi28 := 0; fi28 < 10; fi28++ {
                                v3 = v3
                                var v29 uint64 = 2
                                var v30 int8 = int8(v29)
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            _ = v27
                        }
                        _ = fi26
                    }
                    _ = v25
                }
                _ = v20
            }
            _ = v16
        }
        var v31 int = -994
        _ = v31
    }
    var v32 uint64 = 535
    fmt.Println(v1, v32)
    if ((18446744073709551615 > 4294967295) || (!false)) {
        var v33 uint = 42
        v1 += (^(v1))
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        if ("" < "foo bar") {
            var v35 int8 = int8(v32)
            if false {
                switch (((("\n\t" + "foo bar") + "") + "") + ((("" + "foo bar") + "\n\t") + "a")) {
                case "foo bar":
                    var v36 int = -2147483648
                    _ = v36
                case "a":
                    var v37 string = "hello"
                    _ = v37
                case "":
                    fmt.Println(v1)
                default:
                    v38 := []int{(fi34 / (((fi34 << ((27) & 63))) | 1)), 1, (fi34 << ((846) & 63))}
                    _ = v38
                }
            } else {
                v39 := GenS4{F0: -467}
                v40 := []int{(-(fi34)), -2147483648, 32767, (fi34 << ((613) & 63))}
                switch "hello" {
                case "\n\t":
                    fmt.Println(v32)
                default:
                    switch 42 {
                    case 271:
                        if (float32(1e10) != 0.1) {
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                        }
                    default:
                        switch ("" + "hello") {
                        case "\n\t":
                            fmt.Println(v2, v35, fi34)
                        case "hello":
                            var v41 uint64 = uint64(v1)
                            _ = v41
                        }
                    }
                }
                _ = v39
                _ = v40
            }
            _ = v35
        }
        _ = fi34
    }
    switch "foo bar" {
    case "a":
        switch -2147483648 {
        case 4294967295:
            v42 := []int{32768, 32768}
            _ = v42
        case 32767:
            v43 := GenS5{F0: 196}
            _ = v43
        case 128:
            if (1000 < 1000) {
                v44 := []int{}
                v45 := GenS3{F0: (("hello" + "a") + (("\n\t" + "foo bar") + "a"))}
                if len(v44) >= 2 {
                    v44 = v44[1:len(v44)]
                }
                _ = v44
                _ = v45
            }
        default:
            if false {
                fmt.Println(v3, v3, v32)
                if false {
                    v46 := []int{-2147483648, int(v32)}
                    if true {
                        for fi47 := 0; fi47 < 8; fi47++ {
                            v48 := GenS6{F0: 32767, F1: v32}
                            _ = fi47
                            _ = v48
                        }
                    }
                    var v49 uint16 = (v1 >> ((18446744073709551615) & 63))
                    _ = v46
                    _ = v49
                } else {
                    fmt.Println(v2)
                    v50 := []int{0, 2147483647}
                    fmt.Println(v2, v3, v1)
                    _ = v50
                }
            }
        }
    case "":
        if false {
            v51 := []int{127, -801, 32768}
            v32 = 0
            _ = v51
        }
    case "\n\t":
        v1 = (^(v1))
    default:
        v2 += (0.0 + -2.25)
    }
    v3 -= uint32(v32)
    for fi52 := 0; fi52 < 14; fi52++ {
        v53 := []int{(fi52 - (fi52 >> ((522) & 63))), (fi52 % ((-1) | 1)), 731, (^(fi52))}
        _ = fi52
        _ = v53
    }
    if false {
        if false {
            v54 := GenS5{F0: v1}
            _ = v54
        } else {
            v55 := GenS7{F0: v1, F1: 186}
            v56 := GenS5{F0: v1}
            if (1.5 <= v2) {
                v3 -= v3
                v57 := GenS2{F0: 65535, F1: (v3 - (v3 >> ((967) & 63))), F2: 1}
                _ = v57
            }
            _ = v55
            _ = v56
        }
    } else {
        v58 := GenS8{F0: ""}
        _ = v58
    }
    switch ("hello" + "") {
    case "\n\t":
        v59 := GenS9{F0: 1, F1: 42}
        _ = v59
    case "a":
        v60 := GenS10{F0: "\n\t", F1: int(v3)}
        _ = v60
    case "foo bar":
        v2 = v2
    }
    var v61 string = ("a" + ((("hello" + ("foo bar" + ("\n\t" + "\n\t"))) + "") + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v61
    _ = v7
}
