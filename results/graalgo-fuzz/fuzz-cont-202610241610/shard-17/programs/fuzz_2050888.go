package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0(p0 int32) string {
    var v7 int = 128
    fmt.Println(p0, p0)
    v8 := GenS1{F0: 42, F1: int16(p0), F2: 42}
    _ = p0
    _ = v7
    _ = v8
    return ((("a" + (((("foo bar" + ("a" + ("" + (("foo bar" + "\n\t") + "\n\t")))) + "foo bar") + "hello") + ("" + ""))) + ("" + "\n\t")) + (("\n\t" + ("hello" + ("hello" + "\n\t"))) + ("\n\t" + (("hello" + "a") + ("hello" + "a")))))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 152
    var v3 int = -31
    v4 := make(map[int]int)
    var v5 uint32 = 659
    v6 := make(map[string]string)
    var v9 int8 = 127
    v10 := v1["\n\t"]
    if (!("hello" != "\n\t")) {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int32 = 0
            _ = fi11
            _ = v12
        }
        var v13 bool = (!false)
        _ = v13
    } else {
        if false {
            for fi14 := 0; fi14 < 3; fi14++ {
                fmt.Println(fi14, v2, v9)
                v15 := GenS1{F0: (v2 * 1), F1: 32767, F2: -790}
                var v16 uint = (v2 | 79)
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            var v17 uint16 = 255
            var v18 int32 = -2147483648
            _ = v17
            _ = v18
        }
    }
    v19 := GenS2{F0: 2}
    v20 := GenS2{F0: v5}
    switch (-(v10)) {
    case 9223372036854775807:
        if false {
            if false {
                v21 := v1["\n\t"]
                _ = v21
            } else {
                if (1000 != -32768) {
                    if true {
                        fmt.Println(v2, v3, v10)
                        var v22 int = (v3 >> ((v2) & 63))
                        v3 += v22
                        _ = v22
                    }
                    v4[v10] = (v3 & (v10 - 32767))
                }
                v3 = (^(v10))
                v23 := GenS2{F0: 65535}
                _ = v23
            }
        }
    case 32767:
        var v24 int = (v10 / ((v3) | 1))
        _ = v24
    default:
        for fi25 := 0; fi25 < 7; fi25++ {
            v3 = 0
            if (0 != 9) {
                var v26 uint64 = 491
                switch (("" + "foo bar") + ("hello" + "foo bar")) {
                case "a":
                    fmt.Println(fi25, v9)
                case "\n\t":
                    if false {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            switch "hello" {
                            case "a":
                                fmt.Println(fi25)
                            case "\n\t":
                                v3 -= (fi27 * fi27)
                            case "":
                                v10 -= (fi25 | 4294967295)
                            default:
                                v28 := GenS1{F0: (v2 << (((v2 % (((v2 << ((uint(v5)) & 63))) | 1))) & 63)), F1: 42, F2: -1}
                                _ = v28
                            }
                            var v29 bool = (fi27 == (fi27 ^ v10))
                            _ = fi27
                            _ = v29
                        }
                    } else {
                        var v30 uint16 = 42
                        v10 -= 2147483647
                        _ = v30
                    }
                case "foo bar":
                    v2 = v2
                default:
                    var v31 float64 = 0.0
                    _ = v31
                }
                for fi32 := 0; fi32 < 2; fi32++ {
                    v33 := GenS3{F0: v2}
                    _ = fi32
                    _ = v33
                }
                _ = v26
            }
            var v34 int64 = 42
            _ = fi25
            _ = v34
        }
    }
    var v35 uint64 = uint64(v5)
    var v36 string = ("hello" + ("a" + ((((("a" + "foo bar") + "hello") + "") + "a") + "")))
    switch v36 {
    case "hello":
        if ((v5 >= (v5 / ((v5) | 1))) && (true || false)) {
            var v37 uint16 = 255
            if (((float32((1e10 - 1.5)) * 0.0) - (0.0 / (-2.25 - 1.5))) == 0.1) {
                v38 := GenS1{F0: v2, F1: 1, F2: 42}
                _ = v38
            }
            v39 := []int{(-(v3)), 696, v3}
            _ = v37
            _ = v39
        } else {
            v40 := GenS4{F0: (v5 ^ (v5 * v5)), F1: (false && (3.14159 >= -2.25))}
            _ = v40
        }
    }
    v9 = (v9 - (v9 - (v9 / ((int8(v35)) | 1))))
    var v41 bool = false
    v42 := GenS1{F0: (^(v2)), F1: -32768, F2: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v9
}
