package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 int) uint64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint16 = 0
        if (false && (!(!true))) {
            if (0 > 42) {
                v9 := []int{(-(fi7))}
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 float32 = float32(((-2.25 / (float64(0.0) - 1e10)) + 0.1))
                    p0 -= 127
                    fmt.Println(v11)
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: (v8 / ((901) | 1)), F1: p0, F2: -280}
                _ = v12
                _ = v9
            }
            v13 := GenS1{F0: v8, F1: fi7, F2: 1000}
            _ = v13
        } else {
            var v14 int64 = 9223372036854775807
            _ = v14
        }
        _ = fi7
        _ = v8
    }
    p0 = -9223372036854775808
    p0 += p0
    switch (^(p0)) {
    case 9223372036854775807:
        var v15 float32 = (0.1 + 0.0)
        _ = v15
    case -2147483648:
        if false {
            v16 := GenS2{F0: false, F1: (!true), F2: 100}
            if false {
                for fi17 := 0; fi17 < 20; fi17++ {
                    var v18 bool = (true || ("\n\t" < ("\n\t" + "hello")))
                    p0 += p0
                    if v18 {
                        for fi19 := 0; fi19 < 2; fi19++ {
                            p0 += fi17
                            _ = fi19
                        }
                    } else {
                        if v18 {
                            var v20 uint = 151
                            var v21 bool = (!v18)
                            v18 = v18
                            _ = v20
                            _ = v21
                        } else {
                            switch (fi17 * (-(p0))) {
                            case -265:
                                var v22 float64 = 1e10
                                _ = v22
                            }
                        }
                        v23 := GenS2{F0: (1 <= -102), F1: v18, F2: 100}
                        for fi24 := 0; fi24 < 6; fi24++ {
                            p0 = (fi24 % ((fi17) | 1))
                            _ = fi24
                        }
                        _ = v23
                    }
                    _ = fi17
                    _ = v18
                }
            } else {
                if ((308 < 0) || true) {
                    fmt.Println(p0)
                    var v25 uint64 = 427
                    _ = v25
                } else {
                    p0 -= (p0 + (p0 ^ (p0 ^ p0)))
                    fmt.Println(p0)
                }
            }
            fmt.Println(p0)
            _ = v16
        }
    case 1:
        for fi26 := 0; fi26 < 6; fi26++ {
            for fi27 := 0; fi27 < 10; fi27++ {
                fmt.Println(fi26, p0)
                fmt.Println(fi27, fi26, p0)
                fmt.Println(fi27, p0, fi26)
                _ = fi27
            }
            p0 = (fi26 >> ((548) & 63))
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 int64 = 9223372036854775807
                _ = fi28
                _ = v29
            }
            _ = fi26
        }
    }
    _ = p0
    return 2
}
func genFn1() int32 {
    var v30 float32 = -2.25
    var v31 string = "foo bar"
    _ = v30
    _ = v31
    return 2147483647
}

func main() {
    v1 := []int{32767}
    var v2 int64 = 156
    var v3 bool = ("" <= "")
    v4 := make(map[string]string)
    var v5 float64 = -2.25
    var v6 bool = true
    v3 = (!("" == ("hello" + "hello")))
    v32 := GenS2{F0: v3, F1: (!(!v6)), F2: 42}
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint = 352
        v4["hello"] = "foo bar"
        _ = fi33
        _ = v34
    }
    v5 -= -2.25
    delete(v4, "foo bar")
    v35 := GenS3{F0: 4294967295}
    var v36 int32 = -259
    var v37 uint64 = 207
    switch (("" + "hello") + "\n\t") {
    case "hello":
        v36 = v36
    case "a":
        v38 := GenS3{F0: 380}
        _ = v38
    case "foo bar":
        v36 = 42
    default:
        v2 += (v2 % ((int64(v37)) | 1))
    }
    fmt.Println(v36, v6, v36)
    v5 = v5
    v39 := GenS1{F0: 42, F1: -204, F2: 1000}
    fmt.Println(v3, v6)
    v2 = -988
    _ = len(v1)
    var v40 int = 590
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
