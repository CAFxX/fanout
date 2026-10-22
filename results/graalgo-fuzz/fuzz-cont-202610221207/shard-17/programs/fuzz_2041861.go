package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint16; F2 uint8 }
func genFn0(p0 string, p1 int) int64 {
    var v6 uint32 = 42
    fmt.Println(p0)
    var v7 uint8 = 255
    v8 := GenS1{F0: 42, F1: v6}
    v9 := []int{p1, 127, 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 int16, p1 int16) int8 {
    fmt.Println(p0, p1)
    if false {
        var v10 int64 = genFn0(("" + "\n\t"), 127)
        _ = v10
    }
    var v11 int16 = (p1 - p1)
    v12 := GenS1{F0: 545, F1: 926}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := []string{("" + ("" + "\n\t"))}
    var v2 int64 = -526
    v3 := make(map[string]bool)
    v4 := []int{}
    var v5 int8 = -13
    switch "foo bar" {
    case "a":
        switch (("a" + ("a" + ("foo bar" + ((("" + "hello") + "hello") + "")))) + ("foo bar" + (("foo bar" + "a") + ((("" + "hello") + (("" + "foo bar") + "hello")) + "a")))) {
        case "hello":
            var v13 uint16 = 255
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 3; fi14++ {
                v2 = v2
                _ = fi14
            }
        default:
            var v15 uint64 = 0
            _ = v15
        }
    }
    v1 = append(v1, "foo bar")
    var v16 uint16 = 65535
    var v17 uint8 = 200
    for fi18 := 0; fi18 < 13; fi18++ {
        if true {
            for fi19 := 0; fi19 < 5; fi19++ {
                fmt.Println(v17, v16, fi18)
                _ = len(v1)
                if true {
                    if (false || (!(0 == 42))) {
                        var v20 int64 = v2
                        if false {
                            v21 := v3["\n\t"]
                            _ = v21
                        }
                        _ = v20
                    }
                }
                _ = fi19
            }
        }
        var v22 float64 = 0.0
        _ = fi18
        _ = v22
    }
    var v23 float32 = ((3.14159 * -2.25) - (1.5 * (3.14159 / 1.5)))
    v5 = v5
    if true {
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                _ = len(v3)
                switch fi25 {
                case -9223372036854775808:
                    v5 -= (v5 - genFn1(1, -26))
                }
                var v26 string = ""
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        v27 := GenS2{F0: (65535 > uint32(v17))}
        _ = v27
    } else {
        v16 = uint16(v17)
        fmt.Println(v17, v23)
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 float64 = 1e10
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: "foo bar", F1: v16, F2: (v17 + (v17 ^ v17))}
    switch -1 {
    case 32768:
        fmt.Println(v16, v17, v17)
    case -2147483648:
        var v31 uint64 = 851
        _ = v31
    }
    v23 = 1.5
    fmt.Println(v2)
    var v32 uint8 = (v17 ^ 128)
    if false {
        delete(v3, ("hello" + ("a" + "\n\t")))
        var v33 int32 = 2147483647
        v17 += 128
        _ = v33
    } else {
        var v34 uint = 65535
        v35 := v3[("" + "hello")]
        if (32767 == -32768) {
            var v36 uint8 = (v32 ^ (v32 + v32))
            _ = v36
        } else {
            if v35 {
                v37 := GenS2{F0: (18446744073709551615 == 716)}
                _ = len(v3)
                _ = v37
            } else {
                for fi38 := 0; fi38 < 5; fi38++ {
                    if v35 {
                        var v39 uint = (v34 - 255)
                        _ = v39
                    }
                    v40 := GenS3{F0: ("a" + (("\n\t" + ("foo bar" + "a")) + ("a" + "a"))), F1: (v16 / (((v16 ^ (^(v16)))) | 1)), F2: (v17 * (v32 << (((v34 - (v34 & v34))) & 63)))}
                    _ = fi38
                    _ = v40
                }
                if ((v17 ^ v32) >= (v17 * (v32 | v17))) {
                    v3["foo bar"] = false
                    var v41 int = -2147483648
                    if ((v32 + v32) != v17) {
                        v35 = v35
                    }
                    _ = v41
                }
                v23 = 1e10
            }
        }
        _ = v34
        _ = v35
    }
    v42 := GenS1{F0: 1, F1: 1}
    v2 += (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v42
    _ = v5
}
