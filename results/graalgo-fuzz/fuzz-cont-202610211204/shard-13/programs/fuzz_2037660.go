package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint; F2 uint64 }
func genFn0() uint {
    var v4 int8 = 127
    if true {
        var v5 string = "foo bar"
        _ = v5
    } else {
        fmt.Println(v4)
    }
    v4 = v4
    _ = v4
    return 255
}

func main() {
    v1 := []uint32{}
    var v2 string = "a"
    var v3 uint8 = 141
    switch 2147483647 {
    case -9223372036854775808:
        v6 := GenS1{F0: 1, F1: (v2 + ("\n\t" + "foo bar"))}
        _ = v6
    }
    switch v2 {
    case "a":
        var v7 float32 = 1.5
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: 1, F1: "foo bar"}
        _ = v8
    case "hello":
        v9 := GenS2{F0: 32767}
        _ = v9
    default:
        if (-77 > 1) {
            fmt.Println(v2, v2)
            if false {
                for fi10 := 0; fi10 < 6; fi10++ {
                    if (1.5 >= 1e10) {
                        if true {
                            var v11 uint16 = 841
                            _ = v11
                        }
                        for fi12 := 0; fi12 < 3; fi12++ {
                            if false {
                                fmt.Println(fi10, fi12)
                                var v13 int64 = 42
                                var v14 int8 = -1
                                _ = v13
                                _ = v14
                            }
                            v2 = (v2 + v2)
                            _ = fi12
                        }
                    }
                    if false {
                        var v15 float64 = 0.0
                        _ = v15
                    }
                    v1 = append(v1, 42)
                    _ = fi10
                }
            }
        }
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint8 = v3
        var v18 int16 = -1
        v19 := GenS3{F0: 133}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    switch ("foo bar" + v2) {
    case "":
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    case "\n\t":
        fmt.Println(v3, v2)
    }
    v20 := GenS4{F0: 82, F1: 18446744073709551615, F2: 4294967295}
    switch (v2 + (v2 + "")) {
    case "hello":
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 bool = (0 != 42)
            var v23 uint64 = 18446744073709551615
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        if true {
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS2{F0: 65535}
                var v26 string = ""
                _ = fi24
                _ = v25
                _ = v26
            }
            if (false && (0 != 0)) {
                v2 = v2
                for fi27 := 0; fi27 < 19; fi27++ {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 uint16 = 1
                        fmt.Println(v3)
                        var v30 string = "\n\t"
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi27
                }
            }
            var v31 bool = (0.0 < 3.14159)
            _ = v31
        } else {
            switch 1 {
            case 2147483647:
                if (1e10 <= float64((-2.25 - 0.0))) {
                    switch 128 {
                    case 32767:
                        var v32 uint = 0
                        _ = v32
                    case 1:
                        var v33 int8 = 0
                        _ = v33
                    case 0:
                        v3 -= v3
                    default:
                        for fi34 := 0; fi34 < 4; fi34++ {
                            for fi35 := 0; fi35 < 10; fi35++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v2)
                                _ = fi35
                            }
                            _ = fi34
                        }
                    }
                    if ((610 == 0) || false) {
                        var v36 int64 = 42
                        var v37 float32 = 3.14159
                        if len(v1) > 0 {
                            v1[0] = 776
                        }
                        _ = v36
                        _ = v37
                    } else {
                        v3 -= v3
                        v2 = v2
                        var v38 uint16 = 711
                        _ = v38
                    }
                    if len(v1) > 0 {
                        v1[0] = 146
                    }
                }
            case -856:
                v1 = append(v1, 887)
            default:
                v39 := GenS3{F0: (v3 >> ((65535) & 63))}
                _ = v39
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    if false {
        v2 = v2
    } else {
        switch ((v2 + v2) + (("foo bar" + "\n\t") + "")) {
        case "foo bar":
            for fi40 := 0; fi40 < 5; fi40++ {
                if ((!(0.1 >= -2.25)) || false) {
                    v3 += 167
                } else {
                    if true {
                        var v41 int8 = 75
                        _ = v41
                    }
                    switch v2 {
                    case "hello":
                        var v42 float64 = (0.0 * 1e10)
                        _ = v42
                    case "foo bar":
                        v1 = append(v1, 2)
                    }
                    var v43 float32 = float32(1.5)
                    _ = v43
                }
                _ = fi40
            }
        case "\n\t":
            v44 := GenS2{F0: 744}
            _ = v44
        case "":
            fmt.Println(v2, v2)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
