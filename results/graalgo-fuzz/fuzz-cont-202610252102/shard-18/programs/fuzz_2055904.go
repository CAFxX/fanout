package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 uint32) string {
    p0 = (p0 & p0)
    fmt.Println(p0)
    if false {
        if (p0 >= 983) {
            p0 += 450
            var v5 float64 = (((0.1 - float64((-2.25 * 1.5))) * (float64(1.5) - 0.0)) - 1.5)
            _ = v5
        } else {
            var v6 int8 = 127
            var v7 bool = (255 < 42)
            v6 = -103
            _ = v6
            _ = v7
        }
        fmt.Println(p0)
        v8 := []int{-2147483648}
        _ = v8
    } else {
        var v9 int = 1
        _ = v9
    }
    _ = p0
    return ("" + "\n\t")
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    if (0 <= 42) {
        var v10 float32 = (((float32((3.14159 * (1e10 + float64(-2.25)))) * 3.14159) + (1e10 - 0.1)) / 1e10)
        if (false || ((!(1 >= 0)) && (false || false))) {
            delete(v3, 843)
        } else {
            v11 := GenS1{F0: 262, F1: ("a" + (("" + "hello") + "a")), F2: ""}
            var v12 int = 42
            fmt.Println(v10)
            _ = v11
            _ = v12
        }
        v10 = (v10 * 1.5)
        _ = v10
    } else {
        var v13 int32 = -2147483648
        v13 -= (-(v13))
        fmt.Println(v13)
        _ = v13
    }
    if true {
        if (!(161 != 127)) {
            var v14 int64 = 919
            var v15 string = "hello"
            _ = v14
            _ = v15
        }
    }
    var v16 int64 = 9223372036854775807
    for fi17 := 0; fi17 < 2; fi17++ {
        v16 = v16
        var v18 uint = 426
        if false {
            v18 = v18
        } else {
            var v19 float64 = 1e10
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := GenS1{F0: (v16 << ((v18) & 63)), F1: "a", F2: "hello"}
                var v22 float64 = -2.25
                if (4294967295 <= 1) {
                    fmt.Println(v16, fi20)
                } else {
                    v18 = v18
                    switch (("foo bar" + "foo bar") + (("foo bar" + (("\n\t" + "a") + "a")) + "\n\t")) {
                    case "foo bar":
                        v16 = (v16 / ((v16) | 1))
                    }
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    var v23 int = 420
    v16 -= (v16 | (v16 << ((42) & 63)))
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS2{F0: 42, F1: 42, F2: 92}
        var v26 int16 = 827
        var v27 uint8 = 127
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    if false {
        var v28 int32 = 42
        v29 := GenS2{F0: 44, F1: -32768, F2: 2}
        fmt.Println(v16, v28)
        _ = v28
        _ = v29
    }
    switch v23 {
    case -499:
        if true {
            fmt.Println(v23, v23)
            var v30 string = ("a" + "hello")
            _ = v30
        } else {
            v31 := GenS3{F0: v16, F1: "hello"}
            _ = v31
        }
    case -2147483648:
        v32 := v4[v23]
        _ = v32
    case 0:
        v16 -= (v16 >> ((1) & 63))
    }
    v33 := []int{(^(v23)), 128, 676}
    v16 = v16
    if (-128 != 127) {
        v34 := v1[""]
        v16 += (v16 / ((v16) | 1))
        if false {
            v16 = v16
            _ = len(v1)
            var v35 int8 = -1
            _ = v35
        } else {
            fmt.Println(v34)
            if (!(587 == 469)) {
                if false {
                    v23 = (v23 * v23)
                    v23 += 128
                } else {
                    fmt.Println(v34, v23)
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    if false {
                        fmt.Println(v16, v23)
                        if true {
                            v23 += v23
                            var v36 string = v34
                            _ = v36
                        }
                    } else {
                        v23 += 728
                        fmt.Println(v23, v16, v23)
                    }
                }
                v16 += 1
                v37 := v3[v23]
                _ = v37
            } else {
                v1[(v34 + v34)] = ("hello" + ("hello" + v34))
                delete(v1, ("\n\t" + (v34 + v34)))
                v16 += 136
            }
        }
        _ = v34
    }
    v23 -= v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
}
