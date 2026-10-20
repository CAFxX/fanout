package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 string; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 int64; F2 uint16 }
func genFn0() int8 {
    var v4 bool = true
    v5 := []int{42, 2147483647, 1}
    _ = v4
    _ = v5
    return -69
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    fmt.Println(v1)
    var v6 uint64 = 0
    if true {
        v6 = (v6 - 385)
        var v7 int32 = -2147483648
        v8 := GenS1{F0: -299, F1: 65535, F2: 4294967295}
        _ = v7
        _ = v8
    }
    var v9 int8 = -16
    fmt.Println(v9, v6, v9)
    fmt.Println(v6)
    for fi10 := 0; fi10 < 9; fi10++ {
        v1 = 378
        _ = fi10
    }
    switch "a" {
    case "hello":
        v9 = 55
    case "":
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int16 = 1000
            _ = fi11
            _ = v12
        }
    case "foo bar":
        v9 = -1
    default:
        switch -970 {
        case 39:
            v13 := []int{127, 9223372036854775807, 9223372036854775807}
            _ = v13
        case 2147483647:
            var v14 int = 32768
            _ = v14
        case 906:
            var v15 string = "a"
            _ = v15
        default:
            v6 = v6
        }
    }
    var v16 int = -263
    v17 := GenS1{F0: -154, F1: 1, F2: 400}
    if (false && ((^(v1)) >= (v1 / (((v1 >> ((877) & 63))) | 1)))) {
        v6 += (v6 - v6)
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := GenS1{F0: 1, F1: 255, F2: 65535}
            fmt.Println(v6, v1)
            _ = fi18
            _ = v19
        }
        var v20 int8 = (v9 * (v9 << ((808) & 63)))
        _ = v20
    } else {
        v9 = v9
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS2{F0: 1, F1: 9223372036854775807, F2: genFn0()}
        var v23 bool = (v6 != (v6 | v6))
        _ = fi21
        _ = v22
        _ = v23
    }
    switch 0 {
    case -1:
        if (-32768 >= 42) {
            var v24 int8 = (-(v9))
            fmt.Println(v9, v9)
            switch "a" {
            case "hello":
                fmt.Println(v6, v16)
            default:
                var v25 float32 = (-2.25 / 1.0)
                _ = v25
            }
            _ = v24
        }
    case 42:
        var v26 int16 = -1
        _ = v26
    case 32767:
        switch "" {
        case "hello":
            var v27 int32 = -500
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 8; fi28++ {
                for fi29 := 0; fi29 < 12; fi29++ {
                    if ((("" + ("a" + "")) + ("" + ("hello" + ""))) < "a") {
                        if (true || (true && true)) {
                            switch "" {
                            case "a":
                                v30 := GenS1{F0: 0, F1: 0, F2: 812}
                                _ = v30
                            case "foo bar":
                                var v31 uint16 = (v1 + (v1 * 65535))
                                _ = v31
                            }
                            v32 := GenS3{F0: (("foo bar" + "\n\t") + "hello"), F1: 42, F2: v6}
                            v33 := GenS2{F0: 42, F1: -1, F2: 42}
                            _ = v32
                            _ = v33
                        }
                        if false {
                            v34 := []int{0, (fi29 % (((fi28 & (fi28 >> ((42) & 63)))) | 1)), fi28}
                            v35 := GenS1{F0: 1, F1: 830, F2: 15}
                            _ = v34
                            _ = v35
                        } else {
                            if true {
                                fmt.Println(v1)
                                var v36 int64 = -9223372036854775808
                                _ = v36
                            }
                            v37 := GenS3{F0: ("" + "hello"), F1: 2147483647, F2: (v6 % ((42) | 1))}
                            _ = v37
                        }
                        fmt.Println(fi29, fi29, fi28)
                    } else {
                        var v38 float64 = 1.5
                        _ = v38
                    }
                    switch (fi28 >> ((18446744073709551615) & 63)) {
                    case 128:
                        if false {
                            var v39 int = fi28
                            v1 -= (v1 >> ((2) & 63))
                            var v40 float64 = ((((0.1 + 0.1) + 1.5) * (0.0 / 1.5)) / 1.0)
                            var v41 int64 = 995
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    case -87:
                        v42 := []int{}
                        _ = v42
                    default:
                        fmt.Println(fi29, fi28, v9)
                    }
                    _ = fi29
                }
                if false {
                    v43 := GenS4{F0: ("\n\t" + "\n\t"), F1: 42, F2: (v1 / ((v1) | 1))}
                    _ = v43
                } else {
                    v16 = (v16 + (v16 | 711))
                }
                _ = fi28
            }
        case "":
            fmt.Println(v1, v6)
        }
    }
    var v44 string = ("" + ("a" + ("hello" + ("\n\t" + "a"))))
    if (v44 > "") {
        fmt.Println(v16)
        if true {
            var v45 float32 = 0.0
            _ = v45
        } else {
            fmt.Println(v1, v1, v9)
            v16 -= v16
            fmt.Println(v9, v1, v9)
        }
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        v47 := []int{(^(fi46)), fi46, (v16 + fi46), v16}
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v44
    _ = v6
    _ = v9
}
