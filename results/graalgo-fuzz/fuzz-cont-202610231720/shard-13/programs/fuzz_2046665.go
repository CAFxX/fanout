package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint }

func main() {
    var v1 uint32 = 0
    var v2 string = "\n\t"
    var v3 uint = 42
    var v4 int64 = 0
    v5 := GenS1{F0: false, F1: -2147483648}
    v6 := GenS2{F0: v3, F1: 42}
    v7 := GenS1{F0: true, F1: -269}
    switch 0 {
    case 0:
        var v8 int8 = 100
        _ = v8
    case 128:
        var v9 uint16 = 810
        _ = v9
    case -2147483648:
        fmt.Println(v3)
    default:
        var v10 bool = (1 <= 1)
        _ = v10
    }
    var v11 bool = false
    if (-9223372036854775808 <= v4) {
        v12 := GenS2{F0: v3, F1: 2}
        var v13 int32 = -877
        _ = v12
        _ = v13
    }
    switch 0 {
    case -773:
        var v14 uint64 = 798
        _ = v14
    case -97:
        if (65535 == 65535) {
            for fi15 := 0; fi15 < 6; fi15++ {
                if v11 {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        switch v2 {
                        case "a":
                            var v17 int8 = 42
                            _ = v17
                        case "foo bar":
                            var v18 uint32 = 4294967295
                            _ = v18
                        }
                        _ = fi16
                    }
                    var v19 int8 = 100
                    _ = v19
                }
                var v20 int16 = 42
                _ = fi15
                _ = v20
            }
            v21 := GenS3{F0: 2147483647, F1: v3, F2: 612}
            _ = v21
        } else {
            fmt.Println(v2, v4)
            if v11 {
                if v11 {
                    switch 1 {
                    case 42:
                        if ((!v11) || v11) {
                            v22 := GenS1{F0: (v11 && ("" <= "hello")), F1: 0}
                            var v23 int8 = 42
                            var v24 uint = v3
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                } else {
                    v25 := GenS1{F0: v11, F1: 4294967295}
                    _ = v25
                }
                v26 := []int{9223372036854775807, -705, -2147483648}
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 float64 = (0.1 + 1e10)
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 int16 = -761
            v31 := GenS3{F0: (v1 + 0), F1: v3, F2: v3}
            var v32 uint32 = (^(v1))
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        var v33 float64 = (v28 / 0.1)
        _ = fi27
        _ = v28
        _ = v33
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 string = ""
        if false {
            v4 -= v4
            for fi36 := 0; fi36 < 14; fi36++ {
                if (v11 || (false && ((v3 >> ((4294967295) & 63)) >= (^(v3))))) {
                    var v37 uint32 = 0
                    _ = v37
                }
                _ = fi36
            }
        }
        v2 = (((("\n\t" + "a") + ("\n\t" + ((v2 + v35) + v35))) + v35) + ("a" + "foo bar"))
        _ = fi34
        _ = v35
    }
    if (!v11) {
        v2 = (("hello" + (v2 + "")) + (("\n\t" + (("hello" + v2) + "foo bar")) + v2))
    }
    v38 := []int{32768}
    var v39 float64 = 0.0
    fmt.Println(v1, v2, v2)
    fmt.Println(v3)
    v40 := GenS3{F0: v1, F1: v3, F2: v3}
    fmt.Println(v39)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
