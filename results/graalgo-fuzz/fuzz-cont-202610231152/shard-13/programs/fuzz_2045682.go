package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int; F1 int16; F2 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 int32, p1 int64) uint32 {
    v5 := GenS1{F0: -2147483648, F1: -1}
    v6 := GenS1{F0: -240, F1: -175}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1() uint64 {
    if false {
        if true {
            var v7 uint64 = 71
            var v8 string = ""
            _ = v7
            _ = v8
        }
        var v9 uint32 = 159
        v9 -= v9
        _ = v9
    }
    var v10 uint64 = 4294967295
    _ = v10
    return v10
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    var v3 uint = 0
    var v4 int8 = 111
    switch 350 {
    case -644:
        for fi11 := 0; fi11 < 6; fi11++ {
            v1 = v1
            for fi12 := 0; fi12 < 18; fi12++ {
                switch (("a" + "hello") + "foo bar") {
                case "a":
                    _ = len(v2)
                default:
                    var v13 int64 = -199
                    _ = v13
                }
                fmt.Println(v3)
                fmt.Println(v4, fi11, v1)
                _ = fi12
            }
            _ = fi11
        }
    }
    v3 -= v3
    fmt.Println(v1, v3)
    var v14 float32 = 0.0
    fmt.Println(v3)
    v3 -= v3
    var v15 uint64 = 1
    var v16 int32 = 1
    switch "" {
    case "foo bar":
        if (v1 && (0 == 42)) {
            var v17 uint = (v3 + (v3 / ((822) | 1)))
            _ = v17
        }
    }
    v15 += (^(v15))
    var v18 bool = v1
    v19 := GenS1{F0: 127, F1: -1}
    switch "\n\t" {
    case "hello":
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := []int{648, (-(fi20))}
            _ = fi20
            _ = v21
        }
    case "\n\t":
        v2[("a" + "\n\t")] = true
    case "":
        var v22 uint32 = 2147483648
        _ = v22
    }
    if (v18 && true) {
        switch -1 {
        case -9223372036854775808:
            var v23 int = 127
            _ = v23
        case 0:
            v24 := GenS1{F0: 379, F1: 42}
            _ = v24
        case 2147483647:
            if v18 {
                v25 := []int{127}
                v26 := GenS1{F0: int(v16), F1: -729}
                _ = v25
                _ = v26
            } else {
                if (!v18) {
                    fmt.Println(v14)
                } else {
                    v27 := GenS1{F0: 128, F1: 596}
                    switch -1 {
                    case -957:
                        for fi28 := 0; fi28 < 8; fi28++ {
                            if v1 {
                                for fi29 := 0; fi29 < 4; fi29++ {
                                    fmt.Println(v15, v18, v15)
                                    var v30 int32 = -1
                                    _ = fi29
                                    _ = v30
                                }
                            } else {
                                v31 := GenS2{F0: (fi28 >> (((v3 | v3)) & 63)), F1: 42, F2: 1000}
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    v1 = (255 != 42)
                                    _ = fi32
                                }
                                for fi33 := 0; fi33 < 15; fi33++ {
                                    v3 = uint(v15)
                                    if true {
                                        v1 = v1
                                        var v34 int = fi28
                                        _ = v34
                                    }
                                    _ = fi33
                                }
                                _ = v31
                            }
                            for fi35 := 0; fi35 < 13; fi35++ {
                                if v1 {
                                    v15 = 0
                                    v36 := GenS1{F0: (fi35 + (fi28 + fi35)), F1: 1000}
                                    _ = v36
                                }
                                var v37 uint = 0
                                v38 := []int{fi35, (fi35 * int(v16))}
                                _ = fi35
                                _ = v37
                                _ = v38
                            }
                            _ = fi28
                        }
                    case 42:
                        v39 := []int{9223372036854775807, int(v15), -784}
                        _ = v39
                    }
                    fmt.Println(v14, v16, v1)
                    _ = v27
                }
            }
        default:
            v4 -= 1
        }
        v40 := GenS1{F0: -795, F1: -659}
        switch "\n\t" {
        case "":
            v14 = v14
        case "hello":
            var v41 bool = (-128 >= (v4 << ((18446744073709551615) & 63)))
            _ = v41
        default:
            if v1 {
                for fi42 := 0; fi42 < 7; fi42++ {
                    var v43 uint32 = 2
                    v44 := GenS1{F0: (fi42 & int(v4)), F1: -1}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            } else {
                var v45 int64 = 9223372036854775807
                fmt.Println(v45)
                _ = len(v2)
                _ = v45
            }
        }
        _ = v40
    } else {
        if v1 {
            v46 := GenS3{F0: (^(v3)), F1: 42}
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
