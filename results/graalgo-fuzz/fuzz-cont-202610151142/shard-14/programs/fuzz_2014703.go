package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint8 }
func genFn0(p0 string, p1 uint64) string {
    var v5 uint = 255
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: 1}
            _ = fi6
            _ = v7
        }
    } else {
        switch -18 {
        case 2147483647:
            var v8 uint16 = 42
            _ = v8
        case -2147483648:
            if (-128 < -30) {
                if ((v5 & (v5 | (v5 % (((v5 - v5)) | 1)))) < v5) {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        switch (p0 + p0) {
                        case "hello":
                            if false {
                                v10 := GenS1{F0: 1000}
                                _ = v10
                            }
                        case "a":
                            var v11 bool = false
                            _ = v11
                        default:
                            var v12 uint16 = 65535
                            _ = v12
                        }
                        var v13 int16 = -1
                        p0 = ("hello" + "hello")
                        _ = fi9
                        _ = v13
                    }
                    v14 := []int{-884, 68, -29}
                    var v15 uint64 = 42
                    _ = v14
                    _ = v15
                } else {
                    v16 := []int{-9223372036854775808, 9223372036854775807, 0, 4294967295}
                    _ = v16
                }
                v17 := []int{-9223372036854775808, -1}
                _ = v17
            } else {
                v18 := []int{}
                fmt.Println(p1)
                fmt.Println(p0, p1, v5)
                _ = v18
            }
        case 127:
            for fi19 := 0; fi19 < 2; fi19++ {
                p0 = "foo bar"
                _ = fi19
            }
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -128
    v3 := make(map[int]string)
    var v4 int32 = 1
    if ((("a" + (("hello" + ("\n\t" + ("hello" + ("\n\t" + ("\n\t" + ""))))) + "")) == ("\n\t" + (("\n\t" + ("foo bar" + "hello")) + "foo bar"))) && (false && true)) {
        v4 = 677
    } else {
        if false {
            v4 = v4
        } else {
            var v20 uint16 = 421
            var v21 int64 = -97
            v22 := []int{32768, -246, 42}
            _ = v20
            _ = v21
            _ = v22
        }
    }
    switch -210 {
    case 42:
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 int8 = (v2 - v2)
            _ = fi23
            _ = v24
        }
    case 9223372036854775807:
        v4 = v4
    case 1:
        var v25 int16 = 32767
        _ = v25
    default:
        fmt.Println(v4)
    }
    var v26 int = 4294967295
    var v27 int32 = v4
    if (42 != 447) {
        switch "\n\t" {
        case "a":
            v28 := []int{}
            _ = v28
        case "":
            var v29 uint16 = 0
            _ = v29
        case "\n\t":
            if true {
                switch (("foo bar" + "") + "hello") {
                case "":
                    var v30 float32 = 3.14159
                    _ = v30
                default:
                    _ = len(v1)
                }
                var v31 uint64 = 4294967295
                _ = v31
            } else {
                for fi32 := 0; fi32 < 3; fi32++ {
                    var v33 float64 = 1.5
                    _ = fi32
                    _ = v33
                }
                v27 -= (v27 * -677)
            }
        }
        v4 = v4
    } else {
        if true {
            switch int(v4) {
            case 511:
                v34 := GenS2{F0: true, F1: (v4 & (v27 - -941)), F2: 128}
                _ = v34
            }
            var v35 uint32 = 2147483647
            _ = len(v1)
            _ = v35
        } else {
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 uint = 42
                switch ("hello" + "foo bar") {
                case "foo bar":
                    var v38 bool = true
                    _ = v38
                case "":
                    var v39 int8 = v2
                    _ = v39
                case "a":
                    var v40 int64 = -9223372036854775808
                    _ = v40
                }
                _ = fi36
                _ = v37
            }
            v41 := GenS2{F0: (("" != "") && (true || ("\n\t" != ""))), F1: v4, F2: 0}
            fmt.Println(v26)
            _ = v41
        }
    }
    if ((true || false) && ("foo bar" < ("\n\t" + "a"))) {
        fmt.Println(v2, v26, v26)
        var v42 uint16 = 2
        _ = v42
    } else {
        var v43 uint64 = 2
        if ((("\n\t" + "\n\t") <= "") && true) {
            v4 += v27
        }
        fmt.Println(v26, v2, v26)
        _ = v43
    }
    var v44 uint32 = 42
    for fi45 := 0; fi45 < 14; fi45++ {
        v44 += v44
        v4 = v4
        _ = fi45
    }
    var v46 string = (("\n\t" + (("foo bar" + "a") + "foo bar")) + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v46
}
