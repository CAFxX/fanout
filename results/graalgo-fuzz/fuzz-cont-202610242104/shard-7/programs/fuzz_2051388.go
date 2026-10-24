package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 int8, p1 int64, p2 uint32) int {
    if ((p2 >= (^(p2))) && true) {
        switch 0 {
        case 32767:
            p2 = 2
        case 127:
            switch -9223372036854775808 {
            case 9223372036854775807:
                if (4294967295 > (p2 * 2147483647)) {
                    v4 := []int{-494, -876, 127, -959}
                    fmt.Println(p2)
                    for fi5 := 0; fi5 < 9; fi5++ {
                        v6 := GenS1{F0: 42, F1: 2147483647, F2: 228}
                        p2 = (p2 & 2)
                        _ = fi5
                        _ = v6
                    }
                    _ = v4
                } else {
                    switch "foo bar" {
                    case "":
                        var v7 string = ("" + "\n\t")
                        _ = v7
                    }
                    if true {
                        switch ((("\n\t" + ("" + "\n\t")) + "a") + "") {
                        case "":
                            v8 := []int{128, 32768}
                            _ = v8
                        case "foo bar":
                            p0 = p0
                        }
                        var v9 int = 0
                        if false {
                            switch 869 {
                            case 485:
                                var v10 bool = ((42 > 42) || false)
                                _ = v10
                            case -2147483648:
                                var v11 float64 = -2.25
                                _ = v11
                            }
                        }
                        _ = v9
                    } else {
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 int64 = p1
                            if (((!false) && true) && (1.5 == 0.0)) {
                                fmt.Println(v13)
                                p2 = (p2 & (p2 & (p2 % ((p2) | 1))))
                            }
                            _ = fi12
                            _ = v13
                        }
                    }
                }
            }
        case 42:
            var v14 uint32 = p2
            _ = v14
        }
        var v15 bool = true
        fmt.Println(p2)
        _ = v15
    } else {
        var v16 uint8 = 16
        v16 += v16
        _ = v16
    }
    v17 := []int{-888}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 778
}

func main() {
    var v1 bool = (("\n\t" + "") >= (("foo bar" + "\n\t") + "hello"))
    var v2 uint32 = 36
    var v3 uint = 129
    if (false || (4294967295 == -503)) {
        var v18 uint64 = 42
        v19 := []int{574, 42, -9223372036854775808}
        v20 := GenS2{F0: v3, F1: "foo bar"}
        _ = v18
        _ = v19
        _ = v20
    } else {
        if true {
            if v1 {
                v1 = (1000 >= uint64(v2))
                if v1 {
                    fmt.Println(v1, v3)
                    switch -141 {
                    case 0:
                        v21 := []int{9223372036854775807, -636, -517}
                        _ = v21
                    case 322:
                        v2 += v2
                    }
                    v3 = v3
                }
            }
            v1 = v1
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint8 = 2
            var v24 int8 = 127
            fmt.Println(v24)
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := []int{fi25, (fi25 >> ((v3) & 63)), (fi25 & fi25), (fi25 >> ((18446744073709551615) & 63))}
            _ = fi25
            _ = v26
        }
    }
    var v27 bool = v1
    v28 := GenS2{F0: (v3 - (v3 + (v3 / ((v3) | 1)))), F1: "\n\t"}
    v29 := []int{}
    v30 := GenS2{F0: (^(v3)), F1: ((("hello" + "a") + (("a" + ("foo bar" + "")) + "\n\t")) + ("foo bar" + "\n\t"))}
    var v31 float32 = 3.14159
    if v27 {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        v32 := GenS2{F0: v3, F1: "hello"}
        v27 = (9223372036854775807 == 1)
        _ = v32
    }
    switch "\n\t" {
    case "\n\t":
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
