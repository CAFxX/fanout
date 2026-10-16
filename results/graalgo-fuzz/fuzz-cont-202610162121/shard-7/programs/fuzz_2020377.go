package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2
    var v3 uint32 = (^(v2))
    var v4 uint64 = 2
    v5 := []string{("hello" + "hello"), ("a" + ("a" + ("foo bar" + "a"))), "a"}
    if true {
        var v6 int8 = -63
        var v7 string = "foo bar"
        _ = v6
        _ = v7
    }
    v8 := v1[32767]
    fmt.Println(v4, v2)
    v3 -= v3
    switch 500 {
    case 170:
        var v9 bool = v8
        _ = v9
    case 32768:
        v4 = (v4 % ((1) | 1))
    case 700:
        var v10 bool = v8
        _ = v10
    }
    switch -1 {
    case 4294967295:
        var v11 uint64 = 0
        _ = v11
    }
    if false {
        v12 := v1[185]
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(v3, v12)
                _ = fi14
            }
            var v15 uint16 = 0
            if len(v5) > 0 {
                v5[0] = ("foo bar" + "hello")
            }
            _ = fi13
            _ = v15
        }
        fmt.Println(v8, v8)
        _ = v12
    } else {
        switch 32768 {
        case 1:
            fmt.Println(v4, v3)
        case 457:
            var v16 bool = (233 != v3)
            _ = v16
        case 0:
            switch ("" + ("foo bar" + "hello")) {
            case "\n\t":
                v8 = v8
            case "":
                v2 -= 1
            default:
                var v17 int16 = 1
                _ = v17
            }
        }
        fmt.Println(v4, v3, v8)
        var v18 int = -2147483648
        _ = v18
    }
    v4 -= (v4 | v4)
    v3 = (v2 << ((18446744073709551615) & 63))
    fmt.Println(v2, v2, v8)
    switch 4294967295 {
    case 821:
        if (2147483647 != 1) {
            var v19 string = (("\n\t" + "") + ("foo bar" + ((("a" + "hello") + ("hello" + "a")) + "a")))
            v20 := GenS1{F0: (v4 / ((v4) | 1))}
            _ = v19
            _ = v20
        }
    }
    var v21 uint64 = v4
    v21 -= (v4 + v4)
    if v8 {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 int8 = -1
            if v8 {
                v5 = append(v5, (("a" + "") + "a"))
                v3 -= 225
            }
            if (false && v8) {
                v3 = v2
                switch ("\n\t" + "a") {
                case "":
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if false {
                            if len(v5) > 0 {
                                v5[0] = ("" + ("a" + "hello"))
                            }
                            v2 = (^(v2))
                        }
                        var v25 uint32 = (v2 >> ((785) & 63))
                        _ = fi24
                        _ = v25
                    }
                }
                switch ("hello" + "hello") {
                case "foo bar":
                    if false {
                        switch fi22 {
                        case 449:
                            var v26 int64 = 1
                            _ = v26
                        default:
                            switch "a" {
                            case "foo bar":
                                var v27 uint8 = 1
                                _ = v27
                            case "\n\t":
                                v5 = append(v5, "foo bar")
                            }
                        }
                        v28 := v1[fi22]
                        var v29 uint64 = 42
                        _ = v28
                        _ = v29
                    } else {
                        v21 += (^(v21))
                    }
                case "\n\t":
                    switch (-(fi22)) {
                    case 32768:
                        v30 := v1[828]
                        _ = v30
                    }
                case "a":
                    v21 += v4
                }
            }
            _ = fi22
            _ = v23
        }
        v31 := GenS1{F0: (^(v4))}
        _ = len(v5)
        _ = v31
    }
    if v8 {
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := v1[(-(fi32))]
            for fi34 := 0; fi34 < 4; fi34++ {
                v35 := v1[-9223372036854775808]
                fmt.Println(v35)
                v36 := GenS1{F0: (^(v21))}
                _ = fi34
                _ = v35
                _ = v36
            }
            if true {
                var v37 int64 = -1
                v4 = (v21 + v4)
                fmt.Println(fi32, v4, v21)
                _ = v37
            } else {
                for fi38 := 0; fi38 < 20; fi38++ {
                    v2 -= v3
                    _ = fi38
                }
                fmt.Println(v4, v3, v21)
                switch 127 {
                case 4294967295:
                    fmt.Println(v33, v33)
                case 9223372036854775807:
                    v39 := v1[(fi32 % ((0) | 1))]
                    _ = v39
                default:
                    var v40 int8 = -53
                    _ = v40
                }
            }
            _ = fi32
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
