package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 bool; F2 string }
type GenS5 struct { F0 string }
func genFn0(p0 uint16) uint32 {
    for fi4 := 0; fi4 < 2; fi4++ {
        p0 = 1
        var v5 uint16 = (^(p0))
        switch ("" + "\n\t") {
        case "\n\t":
            p0 = v5
        case "":
            for fi6 := 0; fi6 < 4; fi6++ {
                v7 := []int{(^(fi6)), (fi4 * (fi6 >> ((930) & 63)))}
                v5 -= p0
                v5 = (p0 - 32767)
                _ = fi6
                _ = v7
            }
        case "a":
            switch "a" {
            case "foo bar":
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    p0 += p0
                case "":
                    if false {
                        fmt.Println(fi4, v5, p0)
                        for fi8 := 0; fi8 < 20; fi8++ {
                            p0 = p0
                            _ = fi8
                        }
                    } else {
                        for fi9 := 0; fi9 < 18; fi9++ {
                            v5 += p0
                            _ = fi9
                        }
                    }
                default:
                    var v10 uint8 = 170
                    _ = v10
                }
            default:
                var v11 uint8 = 1
                _ = v11
            }
        default:
            v12 := GenS1{F0: ""}
            _ = v12
        }
        _ = fi4
        _ = v5
    }
    if true {
        v13 := GenS1{F0: (("\n\t" + "foo bar") + "foo bar")}
        v14 := GenS1{F0: ("hello" + ((("\n\t" + ("hello" + "a")) + ("a" + (("" + "\n\t") + ("\n\t" + "\n\t")))) + "\n\t"))}
        v15 := GenS2{F0: 1}
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 2147483647}
    _ = p0
    _ = v16
    return 245
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 string = (((("foo bar" + ("" + (((("hello" + "hello") + "") + ("a" + "hello")) + "\n\t"))) + "foo bar") + ("a" + "foo bar")) + "a")
    fmt.Println(v3)
    v17 := v1[42]
    var v18 int64 = -1
    v19 := GenS3{F0: 2}
    if true {
        var v20 string = (v3 + "\n\t")
        for fi21 := 0; fi21 < 1; fi21++ {
            v3 = v20
            v18 -= (^(v18))
            var v22 string = "\n\t"
            _ = fi21
            _ = v22
        }
        v20 = "\n\t"
        _ = v20
    } else {
        var v23 string = ("a" + v3)
        if ((^(v18)) >= (v18 ^ v18)) {
            switch v3 {
            case "\n\t":
                switch (v17 >> ((4294967295) & 63)) {
                case -378:
                    v17 = int(v18)
                case -2147483648:
                    var v24 uint8 = 2
                    _ = v24
                default:
                    v25 := GenS4{F0: 32767, F1: false, F2: v23}
                    _ = v25
                }
            }
        } else {
            if false {
                v26 := GenS5{F0: "hello"}
                v3 = ("hello" + (("a" + "hello") + ("foo bar" + v3)))
                v17 = v17
                _ = v26
            }
            v27 := []int{v17}
            var v28 int16 = -911
            _ = v27
            _ = v28
        }
        _ = v23
    }
    if true {
        var v29 int64 = v18
        _ = v29
    } else {
        var v30 int = 32767
        v30 -= (v30 - v30)
        _ = v30
    }
    var v31 string = "foo bar"
    switch -9223372036854775808 {
    case -9223372036854775808:
        if (false && ((!true) || true)) {
            fmt.Println(v31, v31)
            switch v17 {
            case -1:
                v18 -= 662
            }
        }
    case -585:
        if false {
            v32 := GenS2{F0: -1}
            _ = v32
        } else {
            var v33 uint16 = 2
            if false {
                var v34 int32 = 2147483647
                _ = v34
            }
            var v35 float32 = 0.0
            _ = v33
            _ = v35
        }
    case 9223372036854775807:
        switch v3 {
        case "\n\t":
            v36 := []int{-849, v17, 4294967295}
            _ = v36
        case "a":
            switch v31 {
            case "\n\t":
                for fi37 := 0; fi37 < 19; fi37++ {
                    v38 := []int{(fi37 ^ fi37), (fi37 << ((65535) & 63)), (fi37 + (fi37 & (v17 & fi37))), int(v18)}
                    v39 := v1[2]
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            case "a":
                v40 := GenS3{F0: 18446744073709551615}
                _ = v40
            }
        case "foo bar":
            v17 = (v17 % ((195) | 1))
        }
    }
    v3 = v31
    fmt.Println(v31, v17)
    for fi41 := 0; fi41 < 16; fi41++ {
        fmt.Println(v3, v17)
        fmt.Println(v3, v18, v17)
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
}
