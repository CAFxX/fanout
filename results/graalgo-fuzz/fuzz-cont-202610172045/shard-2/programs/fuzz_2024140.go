package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 string; F1 int64 }
func genFn0(p0 bool) uint {
    p0 = (!(!(4294967295 == 4294967295)))
    var v7 uint32 = 0
    p0 = (true && ((false || false) || false))
    var v8 int16 = 247
    var v9 int8 = 100
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 79
}
func genFn1(p0 int, p1 bool) int32 {
    for fi10 := 0; fi10 < 19; fi10++ {
        p0 = 0
        _ = fi10
    }
    switch "foo bar" {
    case "\n\t":
        var v11 bool = p1
        _ = v11
    case "foo bar":
        v12 := GenS1{F0: -2147483648}
        _ = v12
    case "a":
        if ((-2147483648 == -594) && true) {
            if true {
                var v13 float64 = (1e10 - -2.25)
                if p1 {
                    var v14 uint64 = 2
                    p1 = ("hello" != "foo bar")
                    _ = v14
                } else {
                    v15 := GenS2{F0: 589, F1: 255}
                    v16 := GenS3{F0: ((("" + ("" + (("" + "hello") + "hello"))) + ("" + "")) + "\n\t"), F1: (p1 || ((!(127 >= -83)) && ((2 == 255) && (false && (float32(1e10) <= float32(v13)))))), F2: 127}
                    _ = v15
                    _ = v16
                }
                if ((!p1) || true) {
                    var v17 uint32 = 4294967295
                    fmt.Println(p0)
                    v18 := GenS4{F0: 37, F1: 42}
                    _ = v17
                    _ = v18
                } else {
                    v13 = v13
                }
                _ = v13
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                p0 = p0
                p1 = (-802 > 42)
                _ = fi19
            }
        } else {
            var v20 int64 = 0
            _ = v20
        }
    }
    var v21 float64 = ((0.1 - 1.5) - -2.25)
    _ = p0
    _ = p1
    _ = v21
    return 2147483647
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 1
    var v3 int8 = 127
    var v4 int32 = -276
    var v5 uint32 = (v2 ^ 2147483647)
    v6 := make(map[string]bool)
    v22 := GenS4{F0: v3, F1: 1000}
    var v23 float32 = ((-2.25 * 0.1) * 3.14159)
    var v24 bool = (0.0 < ((1.5 + 3.14159) * 1.5))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v6)
    if v24 {
        if true {
            if (!(("foo bar" + "foo bar") != ("hello" + "a"))) {
                var v25 int64 = -1
                v24 = v24
                v26 := GenS2{F0: 2, F1: 190}
                _ = v25
                _ = v26
            } else {
                var v27 bool = ((v2 ^ 1) <= v2)
                var v28 uint8 = 0
                _ = v27
                _ = v28
            }
            var v29 int32 = (v4 ^ genFn1(128, (true || ("" > "hello"))))
            _ = v29
        } else {
            var v30 string = "foo bar"
            v6["foo bar"] = true
            v31 := v6[(v30 + ((v30 + (v30 + v30)) + v30))]
            _ = v30
            _ = v31
        }
        _ = len(v1)
    } else {
        if ("hello" < "foo bar") {
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 float64 = -2.25
                if v24 {
                    if (v24 && ("hello" >= ("\n\t" + "foo bar"))) {
                        fmt.Println(v2, v4, v2)
                        if len(v1) > 0 {
                            v1[0] = (v5 << ((1) & 63))
                        }
                        v33 += (((-2.25 + 1.5) - float64((3.14159 * 0.0))) + (v33 / 1.0))
                    } else {
                        _ = len(v6)
                    }
                    v34 := GenS5{F0: "", F1: -9223372036854775808}
                    _ = v34
                }
                _ = fi32
                _ = v33
            }
        } else {
            fmt.Println(v5)
        }
        if true {
            v4 += (v4 ^ (v4 >> ((958) & 63)))
            v35 := GenS5{F0: "foo bar", F1: -9223372036854775808}
            var v36 string = "\n\t"
            _ = v35
            _ = v36
        }
    }
    v24 = v24
    v24 = (3.14159 <= 0.1)
    v6[("foo bar" + "")] = (!(!(false && ((157 > 10) || (v24 || v24)))))
    var v37 int32 = v4
    v3 += v3
    var v38 int8 = (v3 ^ 42)
    var v39 uint = 217
    if len(v1) > 0 {
        v1[0] = (v2 - (^(v5)))
    }
    v24 = ((!(true && ((0 <= -1) || ((v2 / ((v2) | 1)) <= v5)))) && v24)
    v5 += 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
