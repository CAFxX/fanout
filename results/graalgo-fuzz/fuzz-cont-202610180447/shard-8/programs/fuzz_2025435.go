package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32; F1 int64; F2 int }

func main() {
    v1 := []int{128, 42}
    v2 := []uint32{2, 2, 42, 65535}
    v3 := []uint32{42, 0, 0, 42}
    var v4 uint32 = 65535
    v5 := []string{"foo bar", "a", "a", ""}
    var v6 uint64 = 1
    v4 -= (^(v4))
    for fi7 := 0; fi7 < 19; fi7++ {
        switch (("hello" + ("hello" + "a")) + ("" + "\n\t")) {
        case "a":
            if ("foo bar" > "\n\t") {
                var v8 int64 = 42
                var v9 uint8 = 98
                _ = v8
                _ = v9
            } else {
                switch fi7 {
                case 128:
                    switch (fi7 & (fi7 & (fi7 & fi7))) {
                    case 9223372036854775807:
                        fmt.Println(v6, fi7)
                    case 0:
                        fmt.Println(v4)
                    case -9223372036854775808:
                        v10 := GenS1{F0: 0}
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 5; fi11++ {
                            if (false || (42 > -32768)) {
                                var v12 uint = 42
                                v4 -= 42
                                var v13 int32 = 89
                                if (true || ((42 != -1) && (1 <= 1))) {
                                    v12 = (v12 >> ((v12) & 63))
                                    fmt.Println(v6, v4)
                                    var v14 uint8 = 91
                                    _ = v14
                                } else {
                                    switch (fi11 + (fi7 + 32768)) {
                                    case -9223372036854775808:
                                        v15 := GenS2{F0: fi11, F1: "hello"}
                                        _ = v15
                                    }
                                    v4 = (v4 & uint32(v6))
                                    v12 -= 0
                                }
                                fmt.Println(v4, fi7, fi7)
                                _ = v12
                                _ = v13
                            }
                            var v16 int16 = -599
                            fmt.Println(v16)
                            v6 = (v6 << ((752) & 63))
                            _ = fi11
                            _ = v16
                        }
                    }
                case -1:
                    switch (fi7 / ((32767) | 1)) {
                    case 9223372036854775807:
                        var v17 int8 = 100
                        _ = v17
                    case -564:
                        if false {
                            if false {
                                for fi18 := 0; fi18 < 9; fi18++ {
                                    var v19 bool = (!true)
                                    if (42 >= 1000) {
                                        v20 := GenS3{F0: 1, F1: -487, F2: -98}
                                        var v21 int32 = -1
                                        _ = v20
                                        _ = v21
                                    } else {
                                        var v22 int16 = 42
                                        v19 = ((v6 * 1) == v6)
                                        if true {
                                            v6 -= v6
                                        } else {
                                            var v23 int = (fi18 ^ fi18)
                                            _ = v23
                                        }
                                        _ = v22
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                                if true {
                                    for fi24 := 0; fi24 < 2; fi24++ {
                                        v4 = v4
                                        _ = fi24
                                    }
                                    var v25 uint64 = (^(v6))
                                    _ = v25
                                }
                            }
                        } else {
                            for fi26 := 0; fi26 < 14; fi26++ {
                                fmt.Println(v4)
                                var v27 uint32 = 957
                                _ = fi26
                                _ = v27
                            }
                            var v28 bool = (-100 > 1)
                            if (v28 || v28) {
                                var v29 uint8 = 1
                                var v30 uint16 = 762
                                _ = v29
                                _ = v30
                            } else {
                                var v31 float32 = (-2.25 / 1e10)
                                _ = v31
                            }
                            _ = v28
                        }
                    case 128:
                        if (-2147483648 >= 1) {
                            var v32 bool = false
                            _ = v32
                        } else {
                            var v33 string = ""
                            v33 = ((v33 + v33) + v33)
                            v4 = 157
                            _ = v33
                        }
                    }
                default:
                    var v34 float64 = float64(float32((1.5 + ((1e10 / 1.5) + (float64(1e10) / float64(-2.25))))))
                    _ = v34
                }
            }
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = (("foo bar" + (("a" + ("hello" + "\n\t")) + ((("\n\t" + "foo bar") + "hello") + "hello"))) + "hello")
            }
        case "foo bar":
            _ = len(v1)
        default:
            v6 = 42
        }
        var v35 bool = true
        _ = fi7
        _ = v35
    }
    var v36 int8 = -128
    var v37 int64 = -1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v37 = (v37 >> ((0) & 63))
    var v38 float32 = ((1e10 * 0.1) + 3.14159)
    v39 := GenS2{F0: 127, F1: ""}
    var v40 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
