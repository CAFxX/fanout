package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint64) uint {
    v4 := []int{0}
    for fi5 := 0; fi5 < 6; fi5++ {
        v4 = append(v4, (-(fi5)))
        _ = fi5
    }
    v6 := GenS1{F0: 0, F1: -1, F2: "\n\t"}
    switch (("hello" + (("" + "foo bar") + "\n\t")) + "\n\t") {
    case "\n\t":
        var v7 string = "a"
        _ = v7
    default:
        if (!(-128 >= -29)) {
            if ((false && false) && false) {
                if ((!true) || (9223372036854775807 != 460)) {
                    var v8 int32 = 0
                    _ = v8
                }
                for fi9 := 0; fi9 < 5; fi9++ {
                    fmt.Println(fi9, p0)
                    _ = fi9
                }
            }
            p0 -= 933
        } else {
            p0 += 943
            switch 4294967295 {
            case 128:
                var v10 int64 = 947
                _ = v10
            case -408:
                v11 := GenS1{F0: 127, F1: -1, F2: ((("hello" + "\n\t") + (("foo bar" + "a") + "\n\t")) + "a")}
                _ = v11
            }
            var v12 string = "\n\t"
            _ = v12
        }
    }
    _ = p0
    _ = v4
    _ = v6
    return 794
}
func genFn1(p0 uint8) int64 {
    v13 := GenS1{F0: 0, F1: -821, F2: "foo bar"}
    p0 -= 1
    p0 = 46
    var v14 string = "hello"
    _ = p0
    _ = v13
    _ = v14
    return -224
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 1.5
    var v3 bool = ((!(78 <= 1)) && (-128 >= -1))
    v15 := GenS2{F0: 100, F1: 638}
    var v16 float32 = -2.25
    if v3 {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint8 = 127
            if (100 != 42) {
                if (-955 <= 316) {
                    var v19 uint = genFn0(2)
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 int = (fi17 * fi17)
                        v22 := GenS1{F0: 1, F1: -693, F2: "hello"}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS1{F0: 0, F1: 827, F2: ""}
                    _ = v19
                    _ = v23
                }
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 int16 = 32767
                    v26 := GenS3{F0: 38}
                    var v27 bool = (false && (741 >= -1))
                    var v28 uint64 = uint64(v18)
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            } else {
                var v29 uint64 = 18446744073709551615
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 uint = uint(v18)
                    v18 = (v18 - (v18 & (^(v18))))
                    v32 := GenS1{F0: -42, F1: 2147483647, F2: (("foo bar" + "\n\t") + "\n\t")}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = v29
            }
            _ = fi17
            _ = v18
        }
        v2 = (v16 * (v16 * ((1.5 * (((float32((-2.25 / 3.14159)) * (0.0 - v16)) / float32((1.5 * -2.25))) / (float32(3.14159) * v2))) + 1e10)))
    }
    v33 := GenS1{F0: 42, F1: 1, F2: "hello"}
    var v34 int32 = 42
    var v35 uint64 = 18446744073709551615
    v35 -= 18446744073709551615
    fmt.Println(v34, v2)
    var v36 uint16 = 42
    fmt.Println(v3)
    v37 := []int{4294967295, 9223372036854775807, 1}
    v36 = (^(v36))
    if ((v35 & (v35 << ((96) & 63))) <= v35) {
        for fi38 := 0; fi38 < 10; fi38++ {
            switch (fi38 ^ fi38) {
            case 4294967295:
                fmt.Println(v3)
            case 2147483647:
                v35 += (v35 % ((v35) | 1))
            case 864:
                fmt.Println(v36)
            default:
                if true {
                    if len(v37) >= 2 {
                        v37 = v37[1:len(v37)]
                    }
                    for fi39 := 0; fi39 < 8; fi39++ {
                        switch 127 {
                        case -9223372036854775808:
                            switch "a" {
                            case "foo bar":
                                v40 := GenS4{F0: 508, F1: genFn1(0)}
                                _ = v40
                            case "a":
                                v2 -= (float32(1.5) - ((float32(1e10) + 3.14159) / v2))
                            case "\n\t":
                                var v41 uint64 = 42
                                _ = v41
                            }
                        default:
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                        }
                        v42 := GenS5{F0: (^(v36)), F1: 2147483647}
                        _ = fi39
                        _ = v42
                    }
                }
            }
            _ = fi38
        }
    } else {
        if (("hello" < "hello") || (-69 != 127)) {
            delete(v1, "a")
        } else {
            var v43 int64 = 1
            if v3 {
                v2 = v16
            }
            _ = v43
        }
        v44 := GenS4{F0: 1, F1: 1}
        _ = v44
    }
    switch 2147483647 {
    case -1:
        if (!(!v3)) {
            var v45 float32 = v16
            v36 = (v36 + (^(v36)))
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 11; fi46++ {
        v34 = (v34 - (v34 % ((0) | 1)))
        v47 := GenS3{F0: 65535}
        v16 = v2
        _ = fi46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
