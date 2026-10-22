package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint }
func genFn0() string {
    var v5 uint64 = 2
    v6 := GenS1{F0: (v5 % ((117) | 1)), F1: v5}
    var v7 uint8 = 255
    if false {
        var v8 bool = true
        if v8 {
            switch -724 {
            case -38:
                if v8 {
                    v9 := []int{}
                    switch ("" + "") {
                    case "":
                        v9 = append(v9, -1)
                    default:
                        v10 := GenS1{F0: 2, F1: (^(v5))}
                        _ = v10
                    }
                    _ = v9
                } else {
                    v11 := GenS2{F0: v5, F1: 0}
                    var v12 bool = v8
                    _ = v11
                    _ = v12
                }
            case -2147483648:
                switch ("" + "foo bar") {
                case "hello":
                    var v13 uint = 790
                    _ = v13
                case "a":
                    for fi14 := 0; fi14 < 13; fi14++ {
                        if v8 {
                            v8 = v8
                            switch (fi14 + (fi14 >> ((65535) & 63))) {
                            case -9223372036854775808:
                                v15 := []int{(^(fi14)), (fi14 << ((18446744073709551615) & 63))}
                                _ = v15
                            }
                        } else {
                            v16 := GenS1{F0: 896, F1: v5}
                            v7 = v7
                            for fi17 := 0; fi17 < 17; fi17++ {
                                v18 := []int{42}
                                if len(v18) > 0 {
                                    v18[0] = 395
                                }
                                var v19 bool = (v8 && v8)
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            _ = v16
                        }
                        _ = fi14
                    }
                case "foo bar":
                    var v20 bool = v8
                    _ = v20
                }
            }
            var v21 int64 = -9223372036854775808
            var v22 string = (("hello" + ((("" + ("a" + "\n\t")) + "a") + "\n\t")) + "hello")
            _ = v21
            _ = v22
        }
        v5 += v5
        _ = v8
    } else {
        var v23 string = "a"
        v24 := []int{int(v5), 9223372036854775807, 881, 9223372036854775807}
        fmt.Println(v5, v5)
        _ = v23
        _ = v24
    }
    _ = v5
    _ = v6
    _ = v7
    return (("\n\t" + ("hello" + "foo bar")) + "")
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint64 = 499
    v3 := make(map[string]int)
    var v4 int64 = 9223372036854775807
    v2 = v2
    for fi25 := 0; fi25 < 18; fi25++ {
        v2 = v2
        var v26 uint8 = 42
        v27 := GenS3{F0: 18446744073709551615}
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 18446744073709551615, F1: 42}
    if (false && (0.0 < (0.0 + 3.14159))) {
        v29 := GenS1{F0: v2, F1: 42}
        switch (("a" + ("foo bar" + (((("hello" + ("a" + ((("hello" + "\n\t") + ("\n\t" + ("a" + ("\n\t" + (("\n\t" + "") + "\n\t"))))) + "hello"))) + "hello") + "\n\t") + "foo bar"))) + ("hello" + "a")) {
        case "hello":
            var v30 uint32 = 65535
            _ = v30
        case "a":
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 uint32 = 2147483648
                var v33 int32 = -885
                _ = fi31
                _ = v32
                _ = v33
            }
        case "foo bar":
            var v34 float32 = (float32(-2.25) + -2.25)
            _ = v34
        }
        _ = v29
    } else {
        fmt.Println(v4)
        v3[""] = 32768
    }
    v2 -= 4294967295
    var v35 int16 = -768
    v2 -= (v2 << ((255) & 63))
    var v36 uint16 = 255
    var v37 int64 = v4
    switch 445 {
    case 32767:
        if ((!false) && true) {
            fmt.Println(v37)
            fmt.Println(v36, v36, v4)
        } else {
            var v38 uint16 = v36
            v2 -= 403
            _ = v38
        }
    case 127:
        var v39 int64 = (-(v4))
        _ = v39
    }
    fmt.Println(v2)
    if false {
        var v40 int64 = v37
        _ = v40
    } else {
        var v41 int = 127
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
