package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 int16; F1 int64; F2 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 string; F2 bool }
type GenS9 struct { F0 int8; F1 uint }
func genFn0() int16 {
    if true {
        var v7 uint8 = 127
        v8 := GenS1{F0: "a", F1: 42, F2: 564}
        _ = v7
        _ = v8
    } else {
        if (!false) {
            var v9 uint8 = 2
            if ((^(v9)) == v9) {
                v9 -= v9
                var v10 uint64 = 4294967295
                _ = v10
            }
            _ = v9
        }
        var v11 float64 = (1e10 + ((-2.25 / 1.0) + float64(((3.14159 * float32(1.5)) - (1.5 * (3.14159 / 1e10))))))
        var v12 uint = 42
        _ = v11
        _ = v12
    }
    var v13 int8 = -2
    _ = v13
    return 1
}
func genFn1() uint16 {
    v14 := []int{-9223372036854775808, 2147483647, -2147483648}
    v15 := GenS2{F0: 0, F1: 42, F2: (0 < -458)}
    var v16 string = "hello"
    v17 := GenS3{F0: ((v16 + v16) + ("hello" + v16)), F1: 193}
    if false {
        if (!false) {
            if (!false) {
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := GenS4{F0: 42, F1: 373}
                    _ = fi18
                    _ = v19
                }
                v16 = "foo bar"
                if ((!false) || true) {
                    v14 = append(v14, 412)
                    if (905 >= 2147483648) {
                        v16 = ((v16 + v16) + (v16 + v16))
                        switch (v16 + (v16 + ((v16 + v16) + "hello"))) {
                        case "a":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        case "\n\t":
                            v16 = v16
                        }
                    }
                }
            } else {
                v16 = (v16 + (v16 + v16))
                if (-1 >= -171) {
                    v16 = v16
                    v20 := GenS4{F0: 32, F1: 18446744073709551615}
                    _ = v20
                } else {
                    var v21 uint8 = 224
                    _ = v21
                }
            }
            v22 := GenS1{F0: v16, F1: 100, F2: -9223372036854775808}
            if true {
                fmt.Println(v16)
            } else {
                if (!false) {
                    var v23 string = v16
                    _ = v23
                }
            }
            _ = v22
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v16 = (v16 + v16)
            _ = fi24
        }
    } else {
        for fi25 := 0; fi25 < 20; fi25++ {
            v16 = ""
            _ = fi25
        }
        v16 = (v16 + ("\n\t" + (v16 + ((v16 + ((("" + "\n\t") + ("foo bar" + v16)) + "hello")) + v16))))
    }
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 299
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 32767
    v3 := []uint32{42, 2147483647, 319, 2}
    var v4 uint8 = 127
    v5 := []uint32{2147483647, 701, 990}
    v6 := make(map[uint32]bool)
    v4 -= v4
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS5{F0: genFn0(), F1: 42}
        _ = fi26
        _ = v27
    }
    fmt.Println(v4, v4)
    v2 += v2
    delete(v1, "")
    fmt.Println(v2, v4)
    var v28 uint = 4294967295
    if (2 < 65535) {
        var v29 bool = ("hello" < "hello")
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(v2, v28)
            var v32 int32 = 37
            _ = fi31
            _ = v32
        }
        if (1 == -286) {
            if (!false) {
                _ = len(v5)
                v33 := GenS6{F0: 127}
                _ = v33
            } else {
                switch fi30 {
                case 338:
                    v34 := v6[1]
                    _ = v34
                }
            }
        }
        v4 += v4
        _ = fi30
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 int = v2
        for fi37 := 0; fi37 < 12; fi37++ {
            if (776 > 271) {
                if (-1 < 1) {
                    switch fi37 {
                    case 2147483647:
                        var v38 uint = (^(v28))
                        _ = v38
                    }
                    v39 := GenS7{F0: 0}
                    if (2 < 65535) {
                        var v40 uint64 = 25
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        switch (("hello" + (("\n\t" + (("a" + "foo bar") + "a")) + "a")) + "a") {
                        case "hello":
                            v41 := GenS5{F0: 42, F1: fi37}
                            _ = v41
                        }
                        _ = v40
                    }
                    if (-71 < -59) {
                        v2 += -1
                    } else {
                        _ = len(v5)
                        var v42 int16 = 32767
                        _ = v42
                    }
                    if false {
                        _ = len(v1)
                        v43 := GenS8{F0: 99, F1: ("\n\t" + "hello"), F2: ((true || true) && false)}
                        _ = v43
                    }
                    _ = v39
                }
                if false {
                    if false {
                        v4 += v4
                        fmt.Println(v4, v28)
                        var v44 int16 = 42
                        _ = v44
                    }
                    v45 := v1[""]
                    v45 = ""
                    _ = v45
                } else {
                    v4 += (v4 / ((129) | 1))
                }
            } else {
                v4 = 1
                fmt.Println(v28, v36, v2)
                var v46 int = (fi35 & (v36 ^ (fi37 + (v2 * (fi35 & v2)))))
                _ = v46
            }
            _ = fi37
        }
        var v47 uint8 = v4
        _ = fi35
        _ = v36
        _ = v47
    }
    for fi48 := 0; fi48 < 4; fi48++ {
        v28 = v28
        var v49 uint16 = 365
        _ = fi48
        _ = v49
    }
    switch 32767 {
    case -6:
        v2 = 269
    default:
        for fi50 := 0; fi50 < 8; fi50++ {
            var v51 uint = (v28 ^ (v28 >> (((v28 ^ (v28 / ((18446744073709551615) | 1)))) & 63)))
            var v52 int = v2
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    fmt.Println(v4)
    v53 := GenS9{F0: 95, F1: v28}
    switch "foo bar" {
    case "":
        delete(v6, 2147483647)
    }
    fmt.Println(v4, v4)
    v54 := GenS9{F0: -59, F1: v28}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
