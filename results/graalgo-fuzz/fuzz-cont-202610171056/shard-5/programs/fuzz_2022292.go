package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint8; F2 string }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0() uint8 {
    switch (("\n\t" + "\n\t") + (("hello" + ((("\n\t" + ("\n\t" + "\n\t")) + ("hello" + "a")) + "foo bar")) + "a")) {
    case "hello":
        if false {
            var v6 float64 = 0.0
            _ = v6
        } else {
            v7 := GenS1{F0: 81, F1: 2, F2: 91}
            v8 := []int{4294967295, -1}
            var v9 uint64 = 2
            _ = v7
            _ = v8
            _ = v9
        }
    case "foo bar":
        v10 := []int{-290, 4294967295, 706, 32767}
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int8 = -107
            _ = fi11
            _ = v12
        }
    default:
        if (2 < 2) {
            var v13 bool = (!false)
            v13 = v13
            _ = v13
        } else {
            if (42 <= 18446744073709551615) {
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(fi14)
                    switch 127 {
                    case 559:
                        if true {
                            v15 := []int{(fi14 << ((617) & 63)), (fi14 >> ((2) & 63)), -894, (fi14 | fi14)}
                            v16 := GenS2{F0: -1, F1: 128, F2: ("\n\t" + "")}
                            _ = v15
                            _ = v16
                        }
                    default:
                        for fi17 := 0; fi17 < 4; fi17++ {
                            var v18 int16 = 0
                            _ = fi17
                            _ = v18
                        }
                    }
                    _ = fi14
                }
                v19 := GenS1{F0: -1, F1: 4294967295, F2: 42}
                _ = v19
            } else {
                var v20 uint16 = 965
                _ = v20
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 int = (fi21 ^ 0)
                if (84 >= 2) {
                    switch ((("" + ("hello" + ("\n\t" + ((("hello" + "\n\t") + ("foo bar" + "a")) + "\n\t")))) + "") + "\n\t") {
                    case "\n\t":
                        v23 := GenS3{F0: 0, F1: 2}
                        _ = v23
                    case "":
                        if ("hello" == "") {
                            v24 := GenS1{F0: 9223372036854775807, F1: 279, F2: 1}
                            var v25 int = (^(fi21))
                            _ = v24
                            _ = v25
                        } else {
                            v22 -= 32768
                            var v26 int = 4294967295
                            _ = v26
                        }
                    default:
                        var v27 int16 = 32767
                        _ = v27
                    }
                }
                fmt.Println(fi21, v22)
                _ = fi21
                _ = v22
            }
        }
    }
    var v28 string = ("" + "")
    v29 := GenS1{F0: 9223372036854775807, F1: 42, F2: 2}
    _ = v28
    _ = v29
    return 255
}

func main() {
    var v1 int32 = -625
    var v2 bool = (false && (-128 <= 100))
    v3 := []uint32{139, 1}
    v4 := make(map[string]int)
    var v5 int8 = -91
    v30 := v4["hello"]
    for fi31 := 0; fi31 < 6; fi31++ {
        switch (("foo bar" + ("\n\t" + "a")) + "hello") {
        case "":
            _ = len(v3)
        case "a":
            fmt.Println(v30)
        case "foo bar":
            switch "a" {
            case "a":
                v32 := GenS2{F0: 1, F1: 42, F2: ("\n\t" + ("a" + ("foo bar" + ("" + "a"))))}
                _ = v32
            case "hello":
                for fi33 := 0; fi33 < 9; fi33++ {
                    v34 := GenS1{F0: 90, F1: 1, F2: 42}
                    _ = len(v4)
                    _ = fi33
                    _ = v34
                }
            case "\n\t":
                delete(v4, "hello")
            default:
                v30 -= (v30 * (fi31 | (fi31 - fi31)))
            }
        }
        var v35 float64 = ((3.14159 - (-2.25 * (3.14159 * float64(0.0)))) - float64(0.1))
        _ = fi31
        _ = v35
    }
    fmt.Println(v30, v30, v5)
    v2 = false
    var v36 uint64 = 4294967295
    var v37 bool = v2
    v4[("hello" + (((("hello" + "") + "") + "a") + (("hello" + "\n\t") + ("" + (("" + ("" + "\n\t")) + "foo bar")))))] = (v30 / ((42) | 1))
    var v38 uint16 = 879
    v39 := GenS3{F0: 65535, F1: (v36 / ((1) | 1))}
    if ((v1 < (v1 + (^(v1)))) || false) {
        var v40 uint16 = (v38 >> ((18446744073709551615) & 63))
        var v41 int64 = -9223372036854775808
        fmt.Println(v30)
        _ = v40
        _ = v41
    }
    v38 = (^(v38))
    v1 -= 42
    var v42 uint = uint(v38)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
