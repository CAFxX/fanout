package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int8; F1 string; F2 uint32 }

func main() {
    var v1 int16 = 555
    var v2 int8 = 1
    v3 := []string{"\n\t", "hello", ("foo bar" + "foo bar"), "a"}
    v4 := GenS1{F0: 0, F1: 120, F2: 302}
    var v5 uint = 65535
    switch "foo bar" {
    case "foo bar":
        v5 = v5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint64 = 2
        v7 = (^(v7))
        var v8 uint = (v5 | (^(v5)))
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v1, v2)
    var v9 uint8 = 128
    if true {
        if len(v3) > 0 {
            v3[0] = "a"
        }
    } else {
        v9 -= (^(v9))
        v10 := GenS2{F0: -2147483648, F1: 2147483647}
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS3{F0: (v2 <= 100), F1: -225}
            fmt.Println(fi11, v9, v1)
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    v13 := GenS2{F0: 32768, F1: 803}
    if true {
        v14 := GenS1{F0: 888, F1: (v5 ^ (v5 ^ (v5 | (v5 ^ v5)))), F2: 2}
        _ = v14
    } else {
        var v15 bool = (-482 != 1)
        _ = v15
    }
    if (uint16(v5) < 255) {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int16 = (-(v1))
            v9 = 0
            var v18 uint16 = 464
            _ = fi16
            _ = v17
            _ = v18
        }
        if false {
            if (2147483647 <= 708) {
                v5 = v5
                fmt.Println(v9)
            } else {
                var v19 uint64 = 2
                _ = v19
            }
            var v20 string = ""
            if false {
                for fi21 := 0; fi21 < 15; fi21++ {
                    switch (-(fi21)) {
                    case 127:
                        _ = len(v3)
                    case -540:
                        v9 += (^(v9))
                    case -9223372036854775808:
                        fmt.Println(v20, v1)
                    default:
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v2 = (v2 / ((v2) | 1))
                            v20 = ("hello" + v20)
                            var v23 int8 = v2
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = fi21
                }
            } else {
                var v24 uint32 = 2147483647
                fmt.Println(v5, v24, v9)
                if true {
                    v2 = v2
                    fmt.Println(v20, v9, v5)
                } else {
                    var v25 bool = (42 <= (^(v5)))
                    _ = v25
                }
                _ = v24
            }
            _ = v20
        }
    } else {
        var v26 uint8 = 2
        v1 = 0
        v27 := GenS2{F0: 9223372036854775807, F1: 2147483648}
        _ = v26
        _ = v27
    }
    v5 = 65535
    switch 1 {
    case 32768:
        v2 = 122
    case -943:
        if false {
            v3 = append(v3, "foo bar")
            v2 += 46
        } else {
            v2 = 0
            for fi28 := 0; fi28 < 1; fi28++ {
                v1 += int16(v9)
                _ = fi28
            }
            for fi29 := 0; fi29 < 12; fi29++ {
                if false {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v31 := GenS3{F0: false, F1: -733}
                        v32 := GenS1{F0: 4294967295, F1: (v5 - (v5 << ((v5) & 63))), F2: 299}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS1{F0: 18446744073709551615, F1: (v5 << ((18446744073709551615) & 63)), F2: 199}
                    _ = v33
                }
                _ = fi29
            }
        }
    case 127:
        if false {
            if (18446744073709551615 >= 42) {
                switch "\n\t" {
                case "a":
                    if true {
                        v34 := GenS4{F0: 100, F1: "\n\t", F2: 605}
                        var v35 uint32 = 2
                        _ = v34
                        _ = v35
                    }
                }
            } else {
                var v36 uint64 = 884
                _ = v36
            }
            v1 = int16(v9)
            if (346 < 1) {
                var v37 uint8 = v9
                var v38 uint = 255
                _ = v37
                _ = v38
            }
        }
    default:
        v9 = (v9 * (^(v9)))
    }
    v3 = append(v3, (("hello" + (("hello" + "\n\t") + ("\n\t" + ("foo bar" + ("a" + "foo bar"))))) + "a"))
    v2 = (v2 * (^(v2)))
    if false {
        fmt.Println(v2)
        var v39 int8 = v2
        _ = v39
    }
    v5 -= (v5 << ((384) & 63))
    fmt.Println(v5, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
