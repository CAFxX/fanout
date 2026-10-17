package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 bool) string {
    var v7 float32 = 1e10
    switch -2147483648 {
    case 595:
        var v8 uint8 = 1
        _ = v8
    case 128:
        for fi9 := 0; fi9 < 19; fi9++ {
            switch (("\n\t" + "hello") + ("" + ((("foo bar" + "\n\t") + ("foo bar" + "\n\t")) + "\n\t"))) {
            case "a":
                fmt.Println(p0, p0)
            default:
                v10 := GenS1{F0: ""}
                _ = v10
            }
            var v11 float32 = (v7 - (-2.25 - float32(((0.1 + 1e10) - (3.14159 * 1e10)))))
            _ = fi9
            _ = v11
        }
    }
    if (139 != 255) {
        fmt.Println(v7)
        var v12 string = ("\n\t" + ("" + ("hello" + "a")))
        _ = v12
    } else {
        if (-9223372036854775808 == 42) {
            switch 42 {
            case 4294967295:
                v13 := GenS1{F0: ""}
                _ = v13
            case -683:
                v14 := GenS1{F0: ""}
                _ = v14
            case -1:
                v15 := []int{-168}
                _ = v15
            default:
                var v16 float64 = 0.1
                _ = v16
            }
        }
        switch 1 {
        case 9223372036854775807:
            fmt.Println(v7, p0)
        default:
            v17 := []int{128, 9223372036854775807, -2147483648, 128}
            _ = v17
        }
    }
    _ = p0
    _ = v7
    return ""
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []int{1, -2147483648, 42, 2147483647}
    v3 := []uint32{}
    var v4 uint64 = 2
    var v5 int64 = 1
    v6 := []int{-2147483648, 2147483647, 0, 0}
    var v18 int64 = (v5 ^ (v5 >> ((581) & 63)))
    v5 = (v5 & (v18 | v18))
    if true {
        switch ("" + "a") {
        case "foo bar":
            v1 = (^(v1))
        case "\n\t":
            var v19 string = ""
            _ = v19
        case "a":
            v1 = v1
        default:
            var v20 int32 = -1
            _ = v20
        }
        v21 := GenS1{F0: ("" + ("\n\t" + "hello"))}
        _ = v21
    }
    var v22 uint64 = (v4 >> ((18446744073709551615) & 63))
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            v3 = append(v3, (^(v1)))
            _ = fi24
        }
        if false {
            v18 -= (^(v5))
        } else {
            switch ("a" + (("" + "\n\t") + "")) {
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            for fi25 := 0; fi25 < 3; fi25++ {
                v26 := GenS2{F0: v1}
                if (true || (((fi25 & fi25) == fi23) || ((-2.25 - ((3.14159 - 1e10) / -2.25)) != 3.14159))) {
                    var v27 string = ("foo bar" + "hello")
                    _ = len(v2)
                    v5 = -703
                    _ = v27
                } else {
                    v28 := GenS3{F0: -1, F1: 255}
                    var v29 int32 = 0
                    _ = v28
                    _ = v29
                }
                v4 = 42
                _ = fi25
                _ = v26
            }
            v30 := GenS2{F0: v1}
            _ = v30
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            if true {
                var v32 int32 = 1
                var v33 int32 = -2147483648
                _ = v32
                _ = v33
            }
            var v34 int = fi23
            var v35 bool = true
            _ = fi31
            _ = v34
            _ = v35
        }
        _ = fi23
    }
    if true {
        _ = len(v6)
        var v36 uint32 = v1
        var v37 uint64 = (v4 | (^(v4)))
        _ = v36
        _ = v37
    }
    var v38 float32 = float32(0.0)
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 float32 = v38
        switch (fi39 / ((733) | 1)) {
        case 554:
            v1 = v1
        }
        _ = fi39
        _ = v40
    }
    var v41 float32 = v38
    v38 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
